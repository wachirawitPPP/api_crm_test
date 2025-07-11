package models

import (
	"encoding/json"
	"linecrmapi/configs"
	"linecrmapi/structs"
	"time"
)

func GetQueueItemProductId(shop_id int, product_id int, objQuery *structs.QueueItemProduct) (err error) {
	query := configs.DB1.Table("products")
	query = query.Select("products.*,topicals.id AS topical_id")
	query = query.Joins("LEFT JOIN product_shops ON product_shops.product_id = products.id")
	query = query.Joins("LEFT JOIN topicals ON topicals.id = product_shops.topical_id AND topicals.topical_is_active = 1 AND topicals.topical_is_del = 0")
	query = query.Where("products.id = ?", product_id)
	query = query.Where("product_shops.shop_id = ?", shop_id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueItemSetProductId(product_id int, objQuery *structs.QueueItemProduct) (err error) {
	query := configs.DB1.Table("products")
	query = query.Select("products.*")
	query = query.Where("products.id = ?", product_id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueItemProductSet(shop_id int, product_id int, objQuery *structs.QueueItemProduct) (err error) {
	query := configs.DB1.Table("products")
	query = query.Select("products.*")
	query = query.Where("products.id = ?", product_id)
	query = query.Where("products.shop_id = ?", shop_id)
	query = query.Where("products.pd_type_id = ?", 3)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueItemProductIdSet(product_id int, product *[]structs.QueueItemProductSet) (err error) {
	query := configs.DB1.Table("product_sets")
	query = query.Select("product_lists.*")
	query = query.Joins("JOIN product_lists ON product_sets.id = product_lists.product_set_id")
	query = query.Joins("JOIN products ON product_lists.product_id = products.id")
	query = query.Where("product_sets.product_id = ?", product_id)
	query = query.Where("products.pd_is_active = ?", 1)
	query = query.Where("products.pd_is_del = ?", 0)
	if err = query.Scan(&product).Error; err != nil {
		return err
	}
	return nil
}

func GetQueueItemProductSetId(product_id int, shop_id int, product *structs.QueueProductSet) (err error) {
	query := configs.DB1.Table("products")
	query = query.Select("products.id AS id,products.id AS product_id,product_stores.id AS product_store_id, 1 AS pu_amount,product_units.id AS product_units_id,product_stores.pds_cost,products.pd_code,products.pd_name,product_units.pu_rate,ref_units.u_name,product_shop_prices.psp_price_opd,product_shop_prices.psp_price_ipd,topicals.id AS topical_id,products.drug_direction")
	query = query.Joins("INNER JOIN product_stores ON products.id = product_stores.product_id")
	query = query.Joins("INNER JOIN shop_stores ON product_stores.shop_store_id = shop_stores.id")
	query = query.Joins("INNER JOIN shops ON shop_stores.shop_id = shops.id")
	query = query.Joins("INNER JOIN product_units ON products.id = product_units.product_id")
	query = query.Joins("INNER JOIN ref_units ON product_units.unit_id = ref_units.id")
	query = query.Joins("INNER JOIN product_shop_prices ON product_units.id = product_shop_prices.product_unit_id")

	query = query.Joins("LEFT JOIN product_shops ON product_shops.product_id = products.id")
	query = query.Joins("LEFT JOIN topicals ON topicals.id = product_shops.topical_id AND topicals.topical_is_active = 1 AND topicals.topical_is_del = 0")

	query = query.Where("products.id = ?", product_id)
	query = query.Where("product_units.pu_rate = ?", 1)
	query = query.Where("product_units.pu_is_del = ?", 0)
	query = query.Where("shop_stores.ss_type_id = ?", 1)
	query = query.Where("shop_stores.shop_id = ?", shop_id)
	query = query.Where("product_shop_prices.shop_id = ?", shop_id)
	query = query.Where("product_shops.shop_id = ?", shop_id)
	// query = query.Where("product_shop_prices.psp_is_default = ?", 1)
	query = query.Where("product_shop_prices.psp_is_del = ?", 0)
	query = query.Where("products.pd_is_active = ?", 1)
	query = query.Where("products.pd_is_del = ?", 0)

	if err = query.Scan(product).Error; err != nil {
		return err
	}
	return nil
}

func CreateQueueItemProduct(objQuery *structs.QueueProduct) (err error) {
	query := configs.DB1.Table("queue_products").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CreateQueueItemProductList(objQuery *[]structs.QueueProduct) (err error) {
	query := configs.DB1.Table("queue_products").CreateInBatches(&objQuery, 200)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetQueueItemCourseId(course_id int, objQuery *structs.QueueItemCourse) (err error) {
	query := configs.DB1.Table("courses")
	query = query.Select("courses.*")
	query = query.Where("courses.id = ?", course_id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetTopicalId(topical_id int, objQuery *structs.QueueItemTopical) (err error) {
	query := configs.DB1.Table("topicals")
	query = query.Select("topicals.*")
	query = query.Where("topicals.id = ?", topical_id)
	query = query.Where("topicals.topical_type_id = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueItemCourseProduct(course_id int, cp *[]structs.QueueCourseProductSet) (err error) {
	query := configs.DB1.Table("course_products")
	query = query.Select("course_products.*")
	query = query.Joins("JOIN products ON course_products.product_id = products.id")
	query = query.Where("course_products.course_id = ?", course_id)
	query = query.Where("course_products.cp_is_active = ?", 1)
	query = query.Where("course_products.cp_is_del = ?", 0)
	query = query.Where("products.pd_is_del = ?", 0)
	query = query.Where("products.pd_is_active = ?", 1)
	if err = query.Scan(&cp).Error; err != nil {
		return err
	}
	return nil
}

func GetQueueItemCourseIdListProduct(course_id []int, cp *[]structs.QueueCourseProductSetByCourseList) (err error) {
	query := configs.DB1.Table("course_products")
	query = query.Select("course_products.*,courses.course_code,courses.course_name,products.pd_code,products.pd_name")
	query = query.Joins("JOIN products ON course_products.product_id = products.id")
	query = query.Joins("JOIN courses ON course_products.course_id = courses.id")
	query = query.Where("course_products.course_id IN ?", course_id)
	query = query.Where("course_products.cp_is_active = ?", 1)
	query = query.Where("course_products.cp_is_del = ?", 0)
	query = query.Where("products.pd_is_del = ?", 0)
	query = query.Where("products.pd_is_active = ?", 1)
	if err = query.Scan(&cp).Error; err != nil {
		return err
	}
	return nil
}

func GetProductInShopByCode(shop_id int, product_code []string, product *[]structs.OdjProductStore) (err error) {
	query := configs.DB1.Table("product_stores as ps ")
	query = query.Select("ss.ss_name,ps.*")
	query = query.Joins("JOIN shop_stores AS ss ON ss.id = ps.shop_store_id")
	query = query.Where("ss.shop_id = ? AND ps.pds_is_active = 1 AND ps.pds_is_del = 0", shop_id)
	query = query.Where("product_code IN ?", product_code)
	query = query.Group("ps.product_id")
	if err = query.Scan(&product).Error; err != nil {
		return err
	}
	return nil
}

func GetCheckingInShopByCode(shop_id int, checking_code []string, item *[]structs.CheckingList) (err error) {
	query := configs.DB1.Table("checkings")
	query = query.Where("shop_id = ?", shop_id)
	query = query.Where("checking_code IN ?", checking_code)
	query = query.Where("checking_is_del = ?", 0)
	query = query.Where("checking_is_active = ?", 1)
	if err = query.Scan(&item).Error; err != nil {
		return err
	}
	return nil
}

func GetCourseInShopByCode(shop_id int, course_code []string, item *[]structs.CourseList) (err error) {
	query := configs.DB1.Table("courses")
	query = query.Where("shop_id = ?", shop_id)
	query = query.Where("course_code IN ?", course_code)
	query = query.Where("course_is_del = ?", 0)
	query = query.Where("course_is_active = ?", 1)
	if err = query.Scan(&item).Error; err != nil {
		return err
	}
	return nil
}

func GetCourseProductByID(course_id int, product_id int, item *[]structs.CourseList) (err error) {
	query := configs.DB1.Table("course_products")
	query = query.Where("course_id = ? AND product_id = ?", course_id, product_id)
	query = query.Where("cp_is_del = ?", 0)
	query = query.Where("cp_is_active = ?", 1)
	if err = query.Find(&item).Error; err != nil {
		return err
	}
	return nil
}

func GetCourseProductDeleteByCode(shop_id int, course_code string, product_code []string, item *[]structs.QueueCourseProductSetByCourseList) (err error) {
	query := configs.DB1.Table("course_products")
	query = query.Select("course_products.*,courses.course_code,courses.course_name,products.pd_code,products.pd_name")
	query = query.Joins("JOIN products ON course_products.product_id = products.id")
	query = query.Joins("JOIN courses ON course_products.course_id = courses.id")
	query = query.Where("courses.shop_id = ?", shop_id)
	query = query.Where("courses.course_code = ?", course_code)
	query = query.Where("products.pd_code NOT IN ?", product_code)
	query = query.Where("course_products.cp_is_active = ?", 1)
	query = query.Where("course_products.cp_is_del = ?", 0)
	query = query.Where("products.pd_is_del = ?", 0)
	query = query.Where("products.pd_is_active = ?", 1)
	if err = query.Find(&item).Error; err != nil {
		return err
	}
	return nil
}

func GetCourseProductById2(shop_id int, course_id int, product_id int, item *[]structs.QueueCourseProductSetByCourseList) (err error) {
	query := configs.DB1.Table("course_products")
	query = query.Select("course_products.*,courses.course_code,courses.course_name,products.pd_code,products.pd_name")
	query = query.Joins("JOIN products ON course_products.product_id = products.id")
	query = query.Joins("JOIN courses ON course_products.course_id = courses.id")
	query = query.Where("courses.shop_id = ?", shop_id)
	query = query.Where("courses.id = ?", course_id)
	query = query.Where("products.id = ?", product_id)
	query = query.Where("course_products.cp_is_active = ?", 1)
	query = query.Where("course_products.cp_is_del = ?", 0)
	query = query.Where("products.pd_is_del = ?", 0)
	query = query.Where("products.pd_is_active = ?", 1)
	if err = query.Find(&item).Error; err != nil {
		return err
	}
	return nil
}

func GetQueueItemCourseIdSet(course_id int, course *[]structs.QueueItemCourseSet) (err error) {
	query := configs.DB1.Table("course_sets")
	query = query.Select("courses.*,course_lists.course_list_qtyset, course_lists.course_list_opd ,course_lists.course_list_ipd")
	query = query.Joins("JOIN course_lists ON course_sets.id = course_lists.course_set_id")
	query = query.Joins("JOIN courses ON course_lists.course_id = courses.id")
	query = query.Where("course_sets.course_id = ?", course_id)
	if err = query.Scan(&course).Error; err != nil {
		return err
	}
	return nil
}

func CreateQueueItemCourse(objQuery *structs.QueueCourse) (err error) {
	query := configs.DB1.Table("queue_courses").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CreateQueueItemCourseList(objQuery *[]structs.QueueCourse) (err error) {
	query := configs.DB1.Table("queue_courses").CreateInBatches(&objQuery, 200)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetQueueItemCheckingId(checking_id int, objQuery *structs.QueueItemChecking) (err error) {
	query := configs.DB1.Table("checkings")
	query = query.Select("checkings.*")
	query = query.Where("checkings.id = ?", checking_id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueItemCheckingProduct(checking_id int, cp *[]structs.QueueCheckingProductSet) (err error) {
	query := configs.DB1.Table("checking_products")
	query = query.Select("checking_products.*")
	query = query.Joins("JOIN products ON checking_products.product_id = products.id")
	query = query.Where("checking_products.checking_id = ?", checking_id)
	query = query.Where("checking_products.cip_is_active = ?", 1)
	query = query.Where("checking_products.cip_is_del = ?", 0)
	query = query.Where("products.pd_is_del = ?", 0)
	query = query.Where("products.pd_is_active = ?", 1)
	if err = query.Scan(&cp).Error; err != nil {
		return err
	}
	return nil
}

func GetQueueItemCheckingIdSet(checking_id int, checking *[]structs.QueueItemCheckingSet) (err error) {
	query := configs.DB1.Table("checking_sets")
	query = query.Select("checkings.*,checking_lists.checking_list_opd ,checking_lists.checking_list_ipd")
	query = query.Joins("JOIN checking_lists ON checking_sets.id = checking_lists.checking_set_id")
	query = query.Joins("JOIN checkings ON checking_lists.checking_id = checkings.id")
	query = query.Where("checking_sets.checking_id = ?", checking_id)
	if err = query.Scan(&checking).Error; err != nil {
		return err
	}
	return nil
}

func CreateQueueItemChecking(objQuery *structs.QueueChecking) (err error) {
	query := configs.DB1.Table("queue_checkings").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CreateQueueItemCheckingList(objQuery *[]structs.QueueChecking) (err error) {
	query := configs.DB1.Table("queue_checkings").CreateInBatches(&objQuery, 200)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetQueueProduct(queue_id int, objQuery *[]structs.QueueProduct2) (err error) {
	query := configs.DB1.Table("queue_products AS qp")
	query = query.Select("qp.*,rd.dpm_name,rd.dpm_name_en,invoices.id AS inv_id,invoices.inv_code,invoices.inv_datetime ")
	query = query.Joins("LEFT JOIN ref_departments AS rd ON rd.id = qp.dpm_id ")
	query = query.Joins("LEFT JOIN invoices ON invoices.id = qp.quep_id_ref")
	query = query.Where("qp.queue_id = ?", queue_id)
	query = query.Where("qp.quep_type_id = ?", 1)
	query = query.Where("qp.quep_is_active = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueProductUse(queue_id int, objQuery *[]structs.QueueProduct2) (err error) {
	query := configs.DB1.Table("queue_products")
	query = query.Select("queue_products.*,rd.dpm_name,rd.dpm_name_en,invoices.id AS inv_id,invoices.inv_code,invoices.inv_datetime")
	query = query.Joins("INNER JOIN ref_departments AS rd ON rd.id = queue_products.dpm_id ")
	query = query.Joins("LEFT JOIN invoices ON invoices.id = queue_products.quep_id_ref")
	query = query.Where("queue_products.queue_id = ?", queue_id)
	query = query.Where("queue_products.quep_type_id = ?", 2)
	query = query.Where("queue_products.quep_is_active = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CheckQueueProductId(queue_id int, product_id int, objQuery *structs.QueueProduct) (err error) {
	query := configs.DB1.Table("queue_products")
	query = query.Select("queue_products.*")
	query = query.Where("queue_products.queue_id = ?", queue_id)
	query = query.Where("queue_products.product_id = ?", product_id)
	query = query.Where("queue_products.quep_type_id = ?", 1)
	query = query.Where("queue_products.quep_is_active = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueCourse(queue_id int, objQuery *[]structs.QueueCourse2) (err error) {
	query := configs.DB1.Table("queue_courses")
	query = query.Select("queue_courses.*,rd.dpm_name,rd.dpm_name_en,invoices.id AS inv_id,invoices.inv_code,invoices.inv_datetime ")
	query = query.Joins("LEFT JOIN ref_departments AS rd ON rd.id = queue_courses.dpm_id")
	query = query.Joins("LEFT JOIN invoices ON invoices.id = queue_courses.quec_id_ref ")
	query = query.Where("queue_courses.queue_id = ?", queue_id)
	query = query.Where("queue_courses.quec_is_active = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CheckQueueCourseId(queue_id int, course_id int, objQuery *structs.QueueCourse) (err error) {
	query := configs.DB1.Table("queue_courses")
	query = query.Select("queue_courses.*")
	query = query.Where("queue_courses.queue_id = ?", queue_id)
	query = query.Where("queue_courses.course_id = ?", course_id)
	query = query.Where("queue_courses.quec_is_active = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueChecking(queue_id int, objQuery *[]structs.QueueChecking2) (err error) {
	query := configs.DB1.Table("queue_checkings")
	query = query.Select("queue_checkings.*,rd.dpm_name,rd.dpm_name_en,invoices.id AS inv_id,invoices.inv_code,invoices.inv_datetime")
	query = query.Joins("LEFT JOIN ref_departments AS rd ON rd.id = queue_checkings.dpm_id")
	query = query.Joins("LEFT JOIN invoices ON invoices.id = queue_checkings.queci_id_ref")
	query = query.Where("queue_checkings.queue_id = ?", queue_id)
	query = query.Where("queue_checkings.queci_is_active = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CheckQueueCheckingId(queue_id int, checking_id int, objQuery *structs.QueueChecking) (err error) {
	query := configs.DB1.Table("queue_checkings")
	query = query.Select("queue_checkings.*")
	query = query.Where("queue_checkings.queue_id = ?", queue_id)
	query = query.Where("queue_checkings.checking_id = ?", checking_id)
	query = query.Where("queue_checkings.queci_is_active = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueCheckingItem(id int, objQuery *structs.QueueChecking) (err error) {
	query := configs.DB1.Table("queue_checkings")
	query = query.Select("queue_checkings.*")
	query = query.Where("queue_checkings.id = ?", id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueProductId(queue_product_id int, objQuery *structs.QueueProduct) (err error) {
	query := configs.DB1.Table("queue_products")
	query = query.Select("queue_products.*")
	query = query.Where("queue_products.id = ?", queue_product_id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueProductItem(id int, objQuery *structs.QueueProduct) (err error) {
	query := configs.DB1.Table("queue_products")
	query = query.Select("queue_products.*")
	query = query.Where("queue_products.id = ?", id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueCourseId(queue_course_id int, objQuery *structs.QueueCourse) (err error) {
	query := configs.DB1.Table("queue_courses")
	query = query.Select("queue_courses.*")
	query = query.Where("queue_courses.id = ?", queue_course_id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueCourseItem(id int, objQuery *structs.QueueCourse) (err error) {
	query := configs.DB1.Table("queue_courses")
	query = query.Select("queue_courses.*")
	query = query.Where("queue_courses.id = ?", id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueCheckingId(queue_checking_id int, objQuery *structs.QueueChecking) (err error) {
	query := configs.DB1.Table("queue_checkings")
	query = query.Select("queue_checkings.*")
	query = query.Where("queue_checkings.id = ?", queue_checking_id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueCheckingProductItem(id int, objQuery *[]structs.QueueProduct2) (err error) {
	query := configs.DB1.Table("queue_products")
	query = query.Select("queue_products.*")
	query = query.Where("queue_products.queue_checking_id = ?", id)
	query = query.Where("queue_products.quep_is_active = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueCourseProductItem(id int, objQuery *[]structs.QueueProduct2) (err error) {
	query := configs.DB1.Table("queue_products")
	query = query.Select("queue_products.*")
	query = query.Where("queue_products.queue_course_id = ?", id)
	query = query.Where("queue_products.quep_is_active = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CancelQueueChecking(id int, data *structs.QueueChecking) (err error) {
	query := configs.DB1.Table("queue_checkings")
	query = query.Where("id = ?", id)
	query = query.Model(&data)
	query = query.Updates(map[string]interface{}{"queue_checkings.queci_is_active": 0, "queue_checkings.queci_modify": time.Now().Format("2006-01-02 15:04:05")})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CancelQueueCourse(id int, data *structs.QueueCourse) (err error) {
	query := configs.DB1.Table("queue_courses")
	query = query.Where("id = ?", id)
	query = query.Model(&data)
	query = query.Updates(map[string]interface{}{"queue_courses.quec_is_active": 0, "queue_courses.quec_modify": time.Now().Format("2006-01-02 15:04:05")})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CancelQueueCheckingProduct(id int, data *[]structs.QueueProduct2) (err error) {
	query := configs.DB1.Table("queue_products")
	query = query.Where("queue_checking_id = ?", id)
	query = query.Model(&data)
	query = query.Updates(map[string]interface{}{"queue_products.quep_is_active": 0, "queue_products.quep_modify": time.Now().Format("2006-01-02 15:04:05")})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CancelQueueCourseProduct(id int, data *[]structs.QueueProduct2) (err error) {
	query := configs.DB1.Table("queue_products")
	query = query.Where("queue_course_id = ?", id)
	query = query.Model(&data)
	query = query.Updates(map[string]interface{}{"queue_products.quep_is_active": 0, "queue_products.quep_modify": time.Now().Format("2006-01-02 15:04:05")})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CancelQueueProduct(id int, data *structs.QueueProduct) (err error) {
	query := configs.DB1.Table("queue_products")
	query = query.Where("id = ?", id)
	query = query.Model(&data)
	query = query.Updates(map[string]interface{}{"queue_products.quep_is_active": 0, "queue_products.quep_modify": time.Now().Format("2006-01-02 15:04:05")})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func DeleteQueueItemProduct(Id int, objQuery *structs.QueueProduct) (err error) {
	query := configs.DB1.Table("queue_products")
	query = query.Where("id = ?", Id)
	query = query.Delete(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func DeleteQueueItemCourse(Id int, objQuery *structs.QueueCourse) (err error) {
	query := configs.DB1.Table("queue_courses")
	query = query.Where("id = ?", Id)
	query = query.Delete(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func DeleteQueueItemChecking(Id int, objQuery *structs.QueueChecking) (err error) {
	query := configs.DB1.Table("queue_checkings")
	query = query.Where("id = ?", Id)
	query = query.Delete(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateQueueItemProduct(objQuery *structs.QueueProductUpdate) (err error) {
	query := configs.DB1.Table("queue_products")
	query = query.Where("id = ?", objQuery.Id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&objQuery)
	json.Unmarshal(in, &inInterface)

	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateQueueItemCourse(objQuery *structs.QueueCourseUpdate) (err error) {
	query := configs.DB1.Table("queue_courses")
	query = query.Where("id = ?", objQuery.Id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&objQuery)
	json.Unmarshal(in, &inInterface)

	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateQueueItemChecking(objQuery *structs.QueueCheckingUpdate) (err error) {
	query := configs.DB1.Table("queue_checkings")
	query = query.Where("id = ?", objQuery.Id)

	var inInterface map[string]interface{}
	in, _ := json.Marshal(&objQuery)
	json.Unmarshal(in, &inInterface)

	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetQueueItemProductStoresCost(product_stores_id int, product *structs.QueueProductStoresCost) (err error) {
	query := configs.DB1.Table("product_stores")
	query = query.Select("product_stores.id,product_stores.pds_cost")
	query = query.Where("product_stores.id = ?", product_stores_id)
	if err = query.Scan(product).Error; err != nil {
		return err
	}
	return nil
}

func UpdateQueueItemProductDirection(objQuery *structs.QueueDirection) (err error) {
	query := configs.DB1.Table("queue_products")
	query = query.Where("id = ?", objQuery.Id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&objQuery)
	json.Unmarshal(in, &inInterface)

	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetQueueItemTopicalNoti(product_id int, tpd_amount float64, objQuery *structs.QueueTopicalNoti) (err error) {
	query := configs.DB1.Table("topical_products")
	query = query.Select("topical_products.id,topical_products.product_id,topical_products.tpd_amount,topicals.topical_name,topicals.topical_detail")
	query = query.Joins("JOIN topicals ON topical_products.topical_id = topicals.id")
	query = query.Where("topical_products.product_id = ?", product_id)
	query = query.Where("topical_products.tpd_amount <= ?", tpd_amount)
	query = query.Where("topicals.topical_is_active = ?", 1)
	query = query.Where("topicals.topical_is_del = ?", 0)
	query = query.Order("topical_products.tpd_amount DESC")
	query = query.Limit(1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueId(queue_id int, objQuery *structs.GetQueue) (err error) {
	query := configs.DB1.Table("queues")
	query = query.Select("queues.*")
	query = query.Where("queues.id = ?", queue_id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetMaxQueue(shop_id int, customer_id int, objQuery *structs.QueueMax) (err error) {
	query := configs.DB1.Table("queues")
	query = query.Select("MAX(queues.id) AS id")
	query = query.Where("queues.customer_id = ?", customer_id)
	query = query.Where("queues.shop_id = ?", shop_id)
	query = query.Where("queues.que_status_id = ?", 4)
	query = query.Where("queues.que_type_id = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetInvoiceByCustomer(shop_id int, customer_id int, limit int, objQuery *[]structs.QueueInvoiceList) (err error) {
	query := configs.DB1.Table("invoices AS iv")
	query = query.Select("iv.id AS inv_id,iv.queue_id,iv.inv_code,iv.inv_datetime,iv.dpm_id,rd.dpm_name,rd.dpm_name_en")
	query = query.Joins("LEFT JOIN ref_departments AS rd ON rd.id = iv.dpm_id")
	query = query.Where("iv.customer_id = ?", customer_id)
	query = query.Where("iv.shop_id = ?", shop_id)
	query = query.Where("iv.queue_id IS NOT NULL AND iv.inv_is_active = 2")
	query = query.Order("iv.inv_datetime DESC")
	query = query.Limit(limit)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetCheckingByQueueList(shop_id int, customer_id int, limit int, objQuery *[]structs.QueueChecking2) (err error) {
	query := configs.DB1.Table("queue_checkings AS qc")
	query = query.Select("qc.*,iv.inv_id,iv.inv_code,iv.inv_datetime")
	query = query.Joins(`INNER JOIN (  
	SELECT
      queue_id AS id,
      id AS inv_id,
	  inv_code,
      inv_datetime 
  	FROM
    	invoices 
  	WHERE
		customer_id = ? 
		AND shop_id = ? 
		AND inv_is_active = 2 
		AND queue_id IS NOT NULL 
  	ORDER BY
    	inv_datetime DESC 
	LIMIT ?) AS iv ON qc.queue_id = iv.id `, customer_id, shop_id, limit)
	query = query.Where("qc.queci_is_active = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetCoursesByQueueList(shop_id int, customer_id int, limit int, objQuery *[]structs.QueueCourse2) (err error) {
	query := configs.DB1.Table("queue_courses AS qc")
	query = query.Select("qc.*,iv.inv_id,iv.inv_code,iv.inv_datetime")
	query = query.Joins(`INNER JOIN (  
	SELECT
      queue_id AS id,
      id AS inv_id,
	  inv_code,
      inv_datetime 
  	FROM
    	invoices 
  	WHERE
		customer_id = ? 
		AND shop_id = ? 
		AND inv_is_active = 2 
		AND queue_id IS NOT NULL 
  	ORDER BY
    	inv_datetime DESC 
	LIMIT ?) AS iv ON qc.queue_id = iv.id `, customer_id, shop_id, limit)
	query = query.Where("qc.quec_is_active = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetProductByQueueList(shop_id int, customer_id int, limit int, objQuery *[]structs.QueueProduct2) (err error) {
	query := configs.DB1.Table("queue_products AS qc")
	query = query.Select("qc.*,iv.inv_id,iv.inv_code,iv.inv_datetime")
	query = query.Joins(`INNER JOIN (  
	SELECT
      queue_id AS id,
      id AS inv_id,
	  inv_code,
      inv_datetime 
  	FROM
    	invoices 
  	WHERE
		customer_id = ? 
		AND shop_id = ? 
		AND inv_is_active = 2 
		AND queue_id IS NOT NULL 
  	ORDER BY
    	inv_datetime DESC 
	LIMIT ?) AS iv ON qc.queue_id = iv.id `, customer_id, shop_id, limit)
	query = query.Where("qc.quep_is_active = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
func GetQueueItemCostTotalCheck(queue_id int, product_store_id int, checking_id int, objQuery *structs.CostTotal) (err error) {
	query := configs.DB1.Table("product_store_historys")
	query = query.Select("SUM(product_store_historys.pdsh_out * product_store_orders.pdso_cost) AS cost_total")
	query = query.Joins("INNER JOIN product_store_orders ON product_store_historys.product_store_order_id = product_store_orders.id")
	query = query.Joins("INNER JOIN check_products ON product_store_historys.check_product_id = check_products.id")
	query = query.Where("product_store_historys.queue_id = ?", queue_id)
	query = query.Where("product_store_historys.product_store_id = ?", product_store_id)
	query = query.Where("check_products.checking_id = ?", checking_id)
	query = query.Where("product_store_historys.pdsh_out_id != 0")
	query = query.Where("check_products.chkp_is_active = 1")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueItemCostTotalServiec(queue_id int, product_store_id int, course_id int, objQuery *structs.CostTotal) (err error) {
	query := configs.DB1.Table("product_store_historys")
	query = query.Select("SUM(product_store_historys.pdsh_out * product_store_orders.pdso_cost) AS cost_total")
	query = query.Joins("INNER JOIN product_store_orders ON product_store_historys.product_store_order_id = product_store_orders.id")
	query = query.Joins("INNER JOIN service_product_useds ON product_store_historys.service_product_used_id = service_product_useds.id")
	query = query.Where("product_store_historys.queue_id = ?", queue_id)
	query = query.Where("product_store_historys.product_store_id = ?", product_store_id)
	query = query.Where("service_product_useds.course_id = ?", course_id)
	query = query.Where("product_store_historys.pdsh_out_id != 0")
	query = query.Where("service_product_useds.serpu_is_active = 1")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueItemCostTotalServiecNotCourse(queue_id int, product_store_id int, queue_products_id int, objQuery *structs.CostTotal) (err error) {
	query := configs.DB1.Table("product_store_historys")
	query = query.Select("SUM(product_store_historys.pdsh_out * product_store_orders.pdso_cost) AS cost_total")
	query = query.Joins("INNER JOIN product_store_orders ON product_store_historys.product_store_order_id = product_store_orders.id")
	query = query.Joins("INNER JOIN queue_products ON product_store_historys.queue_id = queue_products.queue_id AND product_store_historys.product_store_id = queue_products.product_store_id")
	query = query.Where("product_store_historys.queue_id = ?", queue_id)
	query = query.Where("product_store_historys.product_store_id = ?", product_store_id)
	query = query.Where("queue_products.id = ?", queue_products_id)
	query = query.Where("product_store_historys.pdsh_out_id != 0")
	query = query.Where("queue_products.quep_is_active = 1")
	query = query.Where("product_store_historys.service_product_used_id IS NULL")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
