package models

import (
	"encoding/json"
	"linecrmapi/configs"
	"linecrmapi/structs"
)

func GetServiceQueueItemCourseId(course_id int, objQuery *structs.ServiceQueueItemCourse) (err error) {
	query := configs.DB1.Table("courses")
	query = query.Select("courses.*")
	query = query.Where("courses.id = ?", course_id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CreateServiceQueueItemCourse(objQuery *structs.ServiceQueueCourse) (err error) {
	query := configs.DB1.Table("queue_courses").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetQueueItemCourseServiceProducts(service_id int, cp *[]structs.ServiceProduct) (err error) {
	query := configs.DB1.Table("service_products")
	query = query.Select("service_products.*")
	query = query.Where("service_products.service_id = ?", service_id)
	query = query.Where("service_products.serp_is_active = ?", 1)
	if err = query.Scan(&cp).Error; err != nil {
		return err
	}
	return nil
}

func CreateServiceQueueItemProduct(objQuery *structs.ServiceQueueProduct) (err error) {
	query := configs.DB1.Table("queue_products").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetServiceQueueProductUse(queue_id int, objQuery *[]structs.ServiceQueueProduct) (err error) {
	query := configs.DB1.Table("queue_products")
	query = query.Select("queue_products.*")
	query = query.Where("queue_products.queue_id = ?", queue_id)
	query = query.Where("queue_products.quep_type_id = ?", 2)
	query = query.Where("queue_products.quep_is_active = ?", 1)
	query = query.Order("queue_products.course_id ASC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetServiceQueueCourse(queue_id int, objQuery *[]structs.ServiceQueueCourse) (err error) {
	query := configs.DB1.Table("queue_courses")
	query = query.Select("queue_courses.*,receipts.rec_code")
	query = query.Where("queue_courses.queue_id = ?", queue_id)
	query = query.Where("queue_courses.quec_is_active = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetServiceQueueCourseReceipt(queue_id int, objQuery *[]structs.ServiceQueueCourseReceipt) (err error) {
	query := configs.DB1.Table("queue_courses")
	query = query.Select("queue_courses.*,receipts.rec_code,services.ser_lock_drug")
	query = query.Joins("JOIN receipts ON queue_courses.receipt_id = receipts.id")
	query = query.Joins("JOIN services ON queue_courses.service_id = services.id")
	query = query.Where("queue_courses.queue_id = ?", queue_id)
	query = query.Where("queue_courses.quec_is_active = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetServiceQueueItemProductId(product_id int, objQuery *structs.ServiceQueueItemProduct) (err error) {
	query := configs.DB1.Table("products")
	query = query.Select("products.*")
	query = query.Where("products.id = ?", product_id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetServiceQueueItemProductStoresCost(product_stores_id int, product *structs.ServiceQueueProductStoresCost) (err error) {
	query := configs.DB1.Table("product_stores")
	query = query.Select("product_stores.id,product_stores.pds_cost")
	query = query.Where("product_stores.id = ?", product_stores_id)
	if err = query.Scan(product).Error; err != nil {
		return err
	}
	return nil
}

func GetServiceQueueItemProductIdSet(product_id int, product *[]structs.ServiceQueueItemProductSet) (err error) {
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

func CheckServiceQueueProductId(queue_id int, product_id int, objQuery *structs.ServiceQueueProduct) (err error) {
	query := configs.DB1.Table("queue_products")
	query = query.Select("queue_products.*")
	query = query.Where("queue_products.queue_id = ?", queue_id)
	query = query.Where("queue_products.service_id IS NULL")
	query = query.Where("queue_products.product_id = ?", product_id)
	query = query.Where("queue_products.quep_type_id = ?", 2)
	query = query.Where("queue_products.quep_is_active = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetServiceQueueItemProductSetId(product_id int, shop_id int, product *structs.ServiceQueueProductSet) (err error) {
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

func GetServiceQueueCourseItem(id int, objQuery *structs.ServiceQueueCourse) (err error) {
	query := configs.DB1.Table("queue_courses")
	query = query.Select("queue_courses.*")
	query = query.Where("queue_courses.id = ?", id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CheckServiceQueueCourseId(queue_id int, service_id int, course_id int, objQuery *structs.ServiceQueueCourse) (err error) {
	query := configs.DB1.Table("queue_courses")
	query = query.Select("queue_courses.*")
	query = query.Where("queue_courses.queue_id = ?", queue_id)
	query = query.Where("queue_courses.service_id = ?", service_id)
	query = query.Where("queue_courses.course_id = ?", course_id)
	query = query.Where("queue_courses.quec_is_active = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func UpdateServiceQueueItemCourse(objQuery *structs.ServiceQueueCourseUpdate) (err error) {
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

func GetServiceQueueCourseProductItem(id int, objQuery *[]structs.ServiceQueueProduct) (err error) {
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

func UpdateServiceQueueItemProduct(objQuery *structs.ServiceQueueProductUpdate) (err error) {
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

func GetServiceQueueProductItem(id int, objQuery *structs.ServiceQueueProduct) (err error) {
	query := configs.DB1.Table("queue_products")
	query = query.Select("queue_products.*")
	query = query.Where("queue_products.id = ?", id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetServiceQueueCourseId(queue_course_id int, objQuery *structs.ServiceQueueCourse) (err error) {
	query := configs.DB1.Table("queue_courses")
	query = query.Select("queue_courses.*")
	query = query.Where("queue_courses.id = ?", queue_course_id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func UpdateServiceQueue(queueId int, objQuery *structs.ServiceQueue) (err error) {
	query := configs.DB1.Table("queues")
	query = query.Where("id = ?", queueId)
	query = query.Updates(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetServiceQueueById(queueId int, objQuery *structs.ServiceQueue) error {
	query := configs.DB1.Table("queues")
	query = query.Select("queues.*")
	query = query.Where("queues.id = ?", queueId)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetServiceCourse(queue_id int, objQuery *[]structs.ServiceQueueCourse) (err error) {
	query := configs.DB1.Table("queue_courses")
	query = query.Select("queue_courses.*")
	query = query.Where("queue_courses.queue_id = ?", queue_id)
	query = query.Where("queue_courses.quec_is_active = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CreateQueueServicesUsed(objQuery *structs.ServiceUsed) (err error) {
	query := configs.DB1.Table("service_useds").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateQueueService(service_id int, data *structs.ProcessServiceQueueUpdate) (err error) {
	// query := configs.DB1.Table("services")
	// query = query.Where("services.id = ?", service_id)
	// query = query.Model(&data)
	// query = query.Updates(&data)
	// if err = query.Error; err != nil {
	// 	return err
	// }
	// return nil
	query := configs.DB1.Table("services")
	query = query.Where("services.id = ?", service_id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetServiceProduct(queue_id int, service_id int, objQuery *[]structs.ServiceQueueProduct) (err error) {
	query := configs.DB1.Table("queue_products")
	query = query.Select("queue_products.*")
	query = query.Where("queue_products.queue_id = ?", queue_id)
	query = query.Where("queue_products.service_id = ?", service_id)
	query = query.Where("queue_products.quep_is_active = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetServiceQueueProduct(queue_id int, objQuery *[]structs.ServiceQueueProduct) (err error) {
	query := configs.DB1.Table("queue_products")
	query = query.Select("queue_products.*")
	query = query.Where("queue_products.queue_id = ?", queue_id)
	query = query.Where("queue_products.service_id IS NULL")
	query = query.Where("queue_products.quep_is_active = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CreateQueueServiceProductUsed(objQuery *structs.ServiceProductUsed) (err error) {
	query := configs.DB1.Table("service_product_useds").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CreateQueueServiceProductUsedNotService(objQuery *structs.ServiceProductUsedNotSevice) (err error) {
	query := configs.DB1.Table("service_product_useds").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateQueueServiceProduct(service_product_id int, data *structs.ProcessServiceProductQueueUpdate) (err error) {
	// query := configs.DB1.Table("service_products")
	// query = query.Where("service_products.id = ?", service_product_id)
	// query = query.Model(&data)
	// query = query.Updates(&data)
	// if err = query.Error; err != nil {
	// 	return err
	// }
	// return nil
	query := configs.DB1.Table("service_products")
	query = query.Where("service_products.id = ?", service_product_id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetProcessCourseProductSet(course_id int, cp *[]structs.ServiceQueueCourseProductSet) (err error) {
	query := configs.DB1.Table("course_products")
	query = query.Select("course_products.*,products.pd_name")
	query = query.Joins("INNER JOIN products ON course_products.product_id = products.id")
	query = query.Where("course_products.course_id = ?", course_id)
	query = query.Where("course_products.cp_is_active = ?", 1)
	query = query.Where("course_products.cp_is_del = ?", 0)
	if err = query.Scan(&cp).Error; err != nil {
		return err
	}
	return nil
}

func CreateQueueServiceProduct(objQuery *structs.ServiceProduct) (err error) {
	query := configs.DB1.Table("service_products").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateQueueServiceId(queue_id int, data *structs.QueueServiceUpdate) (err error) {
	// query := configs.DB1.Table("queues")
	// query = query.Where("queues.id = ?", queue_id)
	// query = query.Model(&data)
	// query = query.Updates(&data)
	// if err = query.Error; err != nil {
	// 	return err
	// }
	// return nil
	query := configs.DB1.Table("queues")
	query = query.Where("queues.id = ?", queue_id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetServiceCourseId(service_id int, course_id int, objQuery *structs.Service) (err error) {
	query := configs.DB1.Table("services")
	query = query.Select("services.*")
	query = query.Where("services.id = ?", service_id)
	query = query.Where("services.course_id = ?", course_id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetServiceId(service_id int, objQuery *structs.ServiceList) (err error) {
	query := configs.DB1.Table("services")
	query = query.Select("services.*")
	query = query.Where("services.id = ?", service_id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueItemCourseServiceProductId(service_product_id int, cp *structs.ServiceProduct) (err error) {
	query := configs.DB1.Table("service_products")
	query = query.Select("service_products.*")
	query = query.Where("service_products.id = ?", service_product_id)
	// query = query.Where("service_products.serp_is_active = ?", 1)
	if err = query.Scan(&cp).Error; err != nil {
		return err
	}
	return nil
}

func CreateServicesTranferAdjust(objQuery *structs.Service) (err error) {
	query := configs.DB1.Table("services").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateServiceTranfer(service_id int, data *structs.ServiceUpdateTranfer) (err error) {
	// query := configs.DB1.Table("services")
	// query = query.Where("services.id = ?", service_id)
	// query = query.Model(&data)
	// query = query.Updates(&data)
	// if err = query.Error; err != nil {
	// 	return err
	// }
	// return nil
	query := configs.DB1.Table("services")
	query = query.Where("services.id = ?", service_id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateServiceAdjust(service_id int, data *structs.ServiceUpdateAdjust) (err error) {
	// query := configs.DB1.Table("services")
	// query = query.Where("services.id = ?", service_id)
	// query = query.Model(&data)
	// query = query.Updates(&data)
	// if err = query.Error; err != nil {
	// 	return err
	// }
	// return nil
	query := configs.DB1.Table("services")
	query = query.Where("services.id = ?", service_id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CreateServiceProductTranferAdjust(objQuery *structs.ServiceProduct) (err error) {
	query := configs.DB1.Table("service_products").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateServiceProductTranfer(service_product_id int, data *structs.ServiceProductUpdateTranfer) (err error) {
	// query := configs.DB1.Table("service_products")
	// query = query.Where("service_products.id = ?", service_product_id)
	// query = query.Model(&data)
	// query = query.Updates(&data)
	// if err = query.Error; err != nil {
	// 	return err
	// }
	// return nil
	query := configs.DB1.Table("service_products")
	query = query.Where("service_products.id = ?", service_product_id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateServiceProductAdjust(service_product_id int, data *structs.ServiceProductUpdateAdjust) (err error) {
	// query := configs.DB1.Table("service_products")
	// query = query.Where("service_products.id = ?", service_product_id)
	// query = query.Model(&data)
	// query = query.Updates(&data)
	// if err = query.Error; err != nil {
	// 	return err
	// }
	// return nil
	query := configs.DB1.Table("service_products")
	query = query.Where("service_products.id = ?", service_product_id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetQueueItemCourseServiceProductsList(service_id int, cp *[]structs.ServiceProductList) (err error) {
	query := configs.DB1.Table("service_products")
	query = query.Select("service_products.*")
	query = query.Where("service_products.service_id = ?", service_id)
	query = query.Where("service_products.serp_is_active = ?", 1)
	if err = query.Scan(&cp).Error; err != nil {
		return err
	}
	return nil
}

func GetServiceList(filter structs.ObjPayloadSearchService, isCount bool, dt *[]structs.ServiceSearch, customerId int) (err error) {
	query := configs.DB1.Table("services")
	if isCount == true {
		query = query.Select("services.*,receipts.shop_id AS shop_id,customers.shop_id AS customer_shop_id, receipts.rec_code, receipts.queue_id, courses.course_amount, courses.course_ipd, courses.course_opd, courses.course_cost")
	} else {
		query = query.Select("services.id")
	}
	query = query.Joins("INNER JOIN courses ON services.course_id = courses.id")
	query = query.Joins("INNER JOIN receipts ON services.receipt_id = receipts.id")
	query = query.Joins("INNER JOIN customers ON receipts.customer_id = customers.id")
	query = query.Where("services.customer_id = ?", customerId)
	// query = query.Where("services.shop_id = ?", filter.Shop_id)
	if *filter.Search != "" {
		query = query.Where("services.ser_code LIKE '%" + *filter.Search + "%' OR receipts.rec_code LIKE '%" + *filter.Search + "%' OR services.ser_name LIKE '%" + *filter.Search + "%'")
	}
	// query = query.Where("services.ser_is_active != ? ", 0)
	query = query.Where("services.ser_is_active = ? ", 1)
	// query = query.Where("services.ser_exp_date IS NULL OR services.ser_exp_date > NOW()")
	if isCount == true {
		offset := filter.ActivePage * filter.PerPage
		query = query.Limit(filter.PerPage)
		query = query.Offset(offset)
	}
	query = query.Order("services.ser_datetime DESC")
	query = query.Order("services.ser_code DESC")
	if err = query.Find(&dt).Error; err != nil {
		return err
	}
	return nil
}

func GetServiceUsedList(service_id int, cp *[]structs.ServiceUsedList) (err error) {
	query := configs.DB1.Table("service_useds")
	query = query.Select("service_useds.*,shops.shop_name,users.user_fullname,queues.que_code")
	query = query.Joins("INNER JOIN queues ON service_useds.queue_id = queues.id")
	query = query.Joins("INNER JOIN shops ON shops.id = queues.shop_id")
	query = query.Joins("INNER JOIN users ON users.id = queues.user_id")
	query = query.Where("service_useds.service_id = ?", service_id)
	query = query.Where("service_useds.seru_is_active = ?", 1)
	query = query.Order("service_useds.seru_date DESC")
	if err = query.Scan(&cp).Error; err != nil {
		return err
	}
	return nil
}

func AddLogServices(log *structs.LogServices) (err error) {
	query := configs.DBL1.Table("log_services").Create(&log)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func AddLogServiceProducts(log *structs.LogServiceProducts) (err error) {
	query := configs.DBL1.Table("log_service_products").Create(&log)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateServiceProductUsed(queue_id int, data *structs.UpdateServiceProductUsed) (err error) {
	query := configs.DB1.Table("service_product_useds")
	query = query.Where("service_product_useds.queue_id = ?", queue_id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateServiceUsed(queue_id int, data *structs.UpdateServiceUsed) (err error) {
	query := configs.DB1.Table("service_useds")
	query = query.Where("service_useds.queue_id = ?", queue_id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateCaneclQueueServiceId(queue_id int, data *structs.QueueServiceUpdate) (err error) {
	query := configs.DB1.Table("queues")
	query = query.Where("queues.id = ?", queue_id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetServiceQueueServiceProductUsed(queue_id int, objQuery *[]structs.ServiceProductUsed) (err error) {
	query := configs.DB1.Table("service_product_useds")
	query = query.Select("service_product_useds.*")
	query = query.Where("service_product_useds.queue_id = ?", queue_id)
	query = query.Where("service_product_useds.serpu_is_active = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetServiceQueueServiceUsed(queue_id int, objQuery *[]structs.ServiceUsed) (err error) {
	query := configs.DB1.Table("service_useds")
	query = query.Select("service_useds.*")
	query = query.Where("service_useds.queue_id = ?", queue_id)
	query = query.Where("service_useds.seru_is_active = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

// func GetQueueProductStoreHistory(queue_id int, objQuery *[]structs.ProcessProductStoreHistory) (err error) {
// 	query := configs.DB1.Table("product_store_historys")
// 	query = query.Select("product_store_historys.*")
// 	query = query.Where("product_store_historys.queue_id = ?", queue_id)
// 	query = query.Find(&objQuery)
// 	if query.Error != nil {
// 		return query.Error
// 	}
// 	return nil
// }

func GetQueService(queId int, data *structs.PrintService) (err error) {
	query := configs.DB1.Table("queues")
	query = query.Select("queues.que_code, queues.ctm_fullname, queues.que_user_fullname, queues.ctm_fullname_en, queues.que_note, queues.que_datetime AS que_create, shops.*")
	query = query.Where("queues.id = ?", queId)
	query = query.Joins("INNER JOIN shops ON queues.shop_id = shops.id")
	query = query.Find(&data)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetServiceCourseForPrint(queId int, data *[]structs.PrintServiceCourse) (err error) {
	query := configs.DB1.Table("services")
	query = query.Select("services.ser_customer_id, services.customer_id, CONCAT(CA.ctm_prefix,' ',CA.ctm_fname, '  ' ,CA.ctm_lname) AS use_customer_fullname, CONCAT(CA.ctm_fname_en, '  ' ,CA.ctm_lname_en) AS use_customer_fullname_en, CONCAT(CB.ctm_prefix,' ',CB.ctm_fname, '  ' ,CB.ctm_lname) AS ser_customer_fullname, CONCAT(CB.ctm_fname_en, '  ' ,CB.ctm_lname_en) AS ser_customer_fullname_en, queue_courses.queue_id, services.ser_code, services.ser_name, queue_courses.quec_qty, queue_courses.quec_unit, services.ser_use_date, receipts.rec_code, queue_courses.quec_modify")
	query = query.Where("queue_courses.queue_id = ?", queId)
	query = query.Where("services.ser_is_active >= ?", 1)
	query = query.Where("queue_courses.quec_is_active = ?", 1)
	query = query.Joins("INNER JOIN customers AS CA ON services.customer_id = CA.id")
	query = query.Joins("INNER JOIN customers AS CB ON services.ser_customer_id = CB.id")
	query = query.Joins("INNER JOIN queue_courses ON queue_courses.service_id = services.id")
	query = query.Joins("LEFT JOIN receipts ON services.receipt_id = receipts.id")
	query = query.Find(&data)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetServiceProductForPrint(queId int, data *[]structs.PrintServiceProduct) (err error) {
	query := configs.DB1.Table("queue_products")
	query = query.Select("queue_products.queue_id, queue_products.quep_code, queue_products.quep_name, queue_products.quep_qty, queue_products.quep_unit, services.ser_code, services.ser_name, services.ser_is_active, queue_products.quep_is_active")
	query = query.Where("queue_products.queue_id = ?", queId)
	query = query.Where("queue_products.quep_is_active = ?", 1)
	query = query.Joins("LEFT JOIN services ON queue_products.service_id = services.id")
	query = query.Find(&data)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

// func GetServiceQueueProductAll(queue_id int, objQuery *[]structs.QueueProduct) (err error) {
// 	query := configs.DB1.Table("queue_products")
// 	query = query.Select("queue_products.*")
// 	query = query.Where("queue_products.queue_id = ?", queue_id)
// 	query = query.Where("queue_products.quep_is_active = ?", 1)
// 	query = query.Find(&objQuery)
// 	if query.Error != nil {
// 		return query.Error
// 	}
// 	return nil
// }
