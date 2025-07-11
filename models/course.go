package models

import (
	"encoding/json"
	"linecrmapi/configs"
	"linecrmapi/structs"
)

func GetCourseList(filter structs.ObjPayloadSearchCourse, isCount bool, courses *[]structs.CourseList) (err error) {
	query := configs.DB1.Table("courses")
	query = query.Select("courses.*, categorys.category_type_id, categorys.category_name,	category_types.category_type_th")
	query = query.Joins("JOIN shops ON courses.shop_id = shops.id")
	query = query.Joins("JOIN categorys ON courses.category_id = categorys.id")
	query = query.Joins("JOIN category_types ON categorys.category_type_id = category_types.id")
	query = query.Where("courses.shop_id = ?", filter.Shop_id)
	query = query.Where("courses.course_is_del = 0")
	query = query.Where("categorys.category_is_del = 0")

	if *filter.Search != "" {
		query = query.Where("courses.course_code LIKE '%" + *filter.Search + "%' OR courses.course_name LIKE '%" + *filter.Search + "%'")
	}
	if *filter.Is_active != "" {
		query = query.Where("courses.course_is_active = ?", *filter.Is_active)
	}

	if filter.CategoryId != -1 {
		query = query.Where("courses.category_id = ?", filter.CategoryId)
	}
	if filter.CourseTypeId != -1 {
		query = query.Where("courses.course_type_id = ?", filter.CourseTypeId)
	}

	query = query.Order("LENGTH(courses.course_code) ASC")
	query = query.Order("courses.course_code ASC")
	if isCount == true {
		offset := filter.ActivePage * filter.PerPage
		query = query.Limit(filter.PerPage)
		query = query.Offset(offset)
	}

	if err = query.Scan(&courses).Error; err != nil {
		return err
	}
	return nil
}

func GetCourseDetail(courseId int, Shop_id int, course *structs.CourseDetail) (err error) {
	query := configs.DB1.Table("courses")
	query = query.Select("courses.*, categorys.category_type_id, categorys.category_name,	category_types.category_type_th, course_sets.id AS course_set_id")
	query = query.Joins("JOIN categorys ON courses.category_id = categorys.id")
	query = query.Joins("JOIN category_types ON categorys.category_type_id = category_types.id")
	query = query.Joins("LEFT JOIN course_sets ON courses.id = course_sets.course_id")
	query = query.Where("courses.id = ?", courseId)
	query = query.Where("courses.course_is_del = 0")
	query = query.Where("categorys.category_is_del = 0")
	query = query.Where("courses.shop_id = ?", Shop_id)
	if err = query.Scan(&course).Error; err != nil {
		return err
	}
	return nil
}

func GetCourseProduct(courseId int, Shop_id int, cp *[]structs.Course_product) (err error) {
	query := configs.DB1.Table("course_products")
	query = query.Select("course_products.*, categorys.category_name, categorys.category_is_del, category_types.category_type_th, products.pd_type_id, products.pd_code, products.pd_name, products.pd_code_acc, products.pd_name_acc, products.pd_is_active, products.pd_is_del, product_units.unit_id, product_units.pu_rate, product_stores.pds_cost AS pu_price, product_units.pu_is_del, ref_units.u_name")
	query = query.Joins("JOIN courses ON 	course_products.course_id = courses.id")
	query = query.Joins("JOIN products ON course_products.product_id = products.id")
	query = query.Joins("JOIN product_units ON product_units.product_id = products.id")
	query = query.Joins("JOIN product_stores ON products.id = product_stores.product_id")
	query = query.Joins("JOIN shop_stores ON product_stores.shop_store_id = shop_stores.id")
	query = query.Joins("JOIN ref_units ON product_units.unit_id = ref_units.id")
	query = query.Joins("JOIN categorys ON products.category_id = categorys.id")
	query = query.Joins("JOIN category_types ON categorys.category_type_id = category_types.id")
	query = query.Where("course_products.course_id = ?", courseId)
	query = query.Where("product_units.pu_rate = 1")
	query = query.Where("course_products.cp_is_active = 1")
	query = query.Where("course_products.cp_is_del = 0")
	query = query.Where("categorys.category_is_del = 0")
	query = query.Where("products.pd_is_active = 1")
	query = query.Where("products.pd_is_del = 0")
	query = query.Where("product_units.pu_is_del = 0")
	query = query.Where("courses.shop_id = ?", Shop_id)
	query = query.Where("shop_stores.shop_id = ?", Shop_id)
	query = query.Where("shop_stores.ss_type_id = ?", 1)
	if err = query.Scan(&cp).Error; err != nil {
		return err
	}
	return nil
}

func GetCourseSubList(courseId int, Shop_id int, cs *[]structs.CourseSubList) (err error) {
	query := configs.DB1.Table("course_sets")
	query = query.Select("course_lists.id, course_lists.course_set_id, course_lists.course_list_opd, course_lists.course_list_ipd, course_sets.course_id, courses.shop_id, courses.category_id, courses.course_type_id, courses.course_code, courses.course_name, courses.course_unit,courses.course_opd,courses.course_ipd ,courses.course_cost ,courses.course_fee_df ,courses.course_fee ,courses.course_is_active ,courses.course_is_del ,courses.course_create ,courses.course_update, categorys.category_type_id, categorys.category_name,	category_types.category_type_th")
	query = query.Joins("JOIN course_lists ON course_sets.id = course_lists.course_set_id")
	query = query.Joins("JOIN courses ON course_lists.course_id = courses.id")
	query = query.Joins("JOIN categorys ON courses.category_id = categorys.id")
	query = query.Joins("JOIN category_types ON categorys.category_type_id = category_types.id")
	query = query.Where("course_sets.course_id = ?", courseId)
	query = query.Where("courses.shop_id = ?", Shop_id)
	if err = query.Scan(&cs).Error; err != nil {
		return err
	}
	return nil
}

func GetCourseSubListId(courseId int, Shop_id int, cs *[]structs.CourseSubListId) (err error) {
	query := configs.DB1.Table("course_sets")
	query = query.Select("course_lists.course_id")
	query = query.Joins("JOIN course_lists ON course_sets.id = course_lists.course_set_id")
	query = query.Joins("JOIN courses ON course_lists.course_id = courses.id")
	query = query.Where("course_sets.course_id = ?", courseId)
	query = query.Where("courses.shop_id = ?", Shop_id)
	if err = query.Scan(&cs).Error; err != nil {
		return err
	}
	return nil
}

func AddCourseProduct(dataCo *structs.CourseAction, dataCp *[]structs.Obj_Course_product) (err error) {
	tx := configs.DB1.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	// add courses
	if err = tx.Table("courses").Create(&dataCo).Error; err != nil {
		tx.Rollback()
		return err
	}

	// add course products
	for _, cp := range *dataCp {
		cp.Course_id = dataCo.Id
		cp.Cp_is_active = 1
		cp.Cp_is_del = 0
		if err = tx.Table("course_products").Create(&cp).Error; err != nil {
			tx.Rollback()
			return
		}
	}

	tx.Commit()
	return nil
}

func AddCourseList(dataCo *structs.CourseAction, dataCs *[]structs.Obj_Course_set) (err error) {
	tx := configs.DB1.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	// add courses
	if err = tx.Table("courses").Create(&dataCo).Error; err != nil {
		tx.Rollback()
		return err
	}
	// add course set
	if dataCo.Course_type_id == 2 {
		cs := structs.Obj_Course_set{
			Id:        0,
			Course_id: dataCo.Id,
		}
		if err = tx.Table("course_sets").Create(&cs).Error; err != nil {
			tx.Rollback()
			return
		}
	}

	// add course products
	// for _, cl := range *dataCs {
	// 	cs := structs.Obj_Course_set{
	// 		Id:        0,
	// 		Course_id: cl.Course_id,
	// 	}
	// 	if err = tx.Table("course_sets").Create(&cs).Error; err != nil {
	// 		tx.Rollback()
	// 		return
	// 	}

	// 	cl := structs.Obj_Course_list{
	// 		Id:            0,
	// 		Course_set_id: cs.Id,
	// 		Course_id:     dataCo.Id,
	// 	}
	// 	if err = tx.Table("course_lists").Create(&cl).Error; err != nil {
	// 		tx.Rollback()
	// 		return
	// 	}
	// }

	tx.Commit()
	return nil
}

// func UpdateCourse(course_id int, data *structs.CourseAction) (err error) {
// 	query := configs.DB1.Table("courses")
// 	if data.Course_is_active == 0 {
// 		query = query.Where("courses.id = ?", course_id)
// 		query = query.Model(&data)
// 		query = query.Updates(map[string]interface{}{"courses.course_is_active": 0})
// 	}
// 	query = query.Where("id = ?", course_id)
// 	query = query.Model(&data)
// 	query = query.Updates(&data)
// 	if data.Course_lock_drug == 0 {
// 		query = query.Updates(map[string]interface{}{"courses.course_lock_drug": 0})
// 	}
// 	if data.Course_use_date == 0 {
// 		query = query.Updates(map[string]interface{}{"courses.course_use_date": 0})
// 	}
// 	if data.Course_exp_date == 0 {
// 		query = query.Updates(map[string]interface{}{"courses.course_exp_date": 0})
// 	}
// 	if err = query.Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

func UpdateCourse(course_id int, data *structs.CourseAction) (err error) {
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)

	if *data.Course_image == "" {
		delete(inInterface, "course_image")
	}

	query := configs.DB1.Table("courses")
	query = query.Where("id = ?", course_id)
	query = query.Model(&data)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateCourseProduct(course_id int, data *[]structs.Obj_Course_product) (err error) {
	tx := configs.DB1.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	for _, cp := range *data {
		if cp.Id > 0 {
			if err = tx.Table("course_products").Where("id = ?", cp.Id).Updates(&cp).Error; err != nil {
				tx.Rollback()
				return
			}
		} else {
			cp.Course_id = course_id
			cp.Cp_is_active = 1
			cp.Cp_is_del = 0
			if err = tx.Table("course_products").Create(&cp).Error; err != nil {
				tx.Rollback()
				return
			}
		}
	}
	tx.Commit()
	return nil
}

func UpdateCourseList(Course_set_id int, data *[]structs.Obj_Course_list) (err error) {
	for _, cl := range *data {
		if cl.Id == 0 {
			cl := structs.Obj_Course_list{
				Id:              0,
				Course_set_id:   Course_set_id,
				Course_id:       cl.Course_id,
				Course_list_opd: cl.Course_list_opd,
				Course_list_ipd: cl.Course_list_ipd,
			}
			if err = configs.DB1.Table("course_lists").Create(&cl).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func GetCourseListNotSet(filter structs.ObjPayloadSearchCourseList, isCount bool, course *[]structs.CourseNotSetList, course_id []int) (err error) {
	query := configs.DB1.Table("courses")
	query = query.Select("courses.*, categorys.category_type_id, categorys.category_name,	category_types.category_type_th")
	query = query.Joins("JOIN categorys ON courses.category_id = categorys.id")
	query = query.Joins("JOIN category_types ON categorys.category_type_id = category_types.id")
	query = query.Where("courses.shop_id = ?", filter.Shop_id)
	query = query.Where("courses.course_is_del = 0")
	query = query.Where("categorys.category_is_del = 0")
	query = query.Where("courses.course_type_id != 2")
	if len(course_id) > 0 {
		query = query.Where("courses.id NOT IN ?", course_id)
	}
	if *filter.Search != "" {
		query = query.Where("courses.course_name LIKE '%" + *filter.Search + "%' OR courses.course_code LIKE '%" + *filter.Search + "%' ")
	}

	query = query.Where("courses.course_is_active = ?", 1)
	if isCount == true {
		offset := filter.ActivePage * filter.PerPage
		query = query.Limit(filter.PerPage)
		query = query.Offset(offset)
	}

	if err = query.Scan(&course).Error; err != nil {
		return err
	}
	return nil
}

func GetCourseProductList(filter structs.ObjPayloadSearchProductCourse, isCount bool, product *[]structs.Product) (err error) {
	query := configs.DB1.Table("products")
	query = query.Select("product_stores.pds_cost AS pu_price, categorys.category_name, categorys.category_is_del, category_types.category_type_th, products.id, products.shop_id, products.category_id, products.pd_type_id, products.pd_code, products.pd_name, products.pd_code_acc, products.pd_name_acc, products.pd_is_active, products.pd_is_del, product_units.unit_id, product_units.pu_rate, product_units.pu_is_del, ref_units.u_name")
	query = query.Joins("JOIN product_units ON product_units.product_id = products.id")
	query = query.Joins("JOIN product_stores ON products.id = product_stores.product_id")
	query = query.Joins("JOIN shop_stores ON product_stores.shop_store_id = shop_stores.id")
	query = query.Joins("JOIN ref_units ON product_units.unit_id = ref_units.id")
	query = query.Joins("JOIN categorys ON products.category_id = categorys.id")
	query = query.Joins("JOIN category_types ON categorys.category_type_id = category_types.id")
	query = query.Where("shop_stores.shop_id = ?", filter.Shop_id)
	query = query.Where("shop_stores.ss_type_id = ?", 1)
	query = query.Where("product_units.pu_rate = 1")
	query = query.Where("products.pd_is_active = 1")
	query = query.Where("products.pd_is_del = 0")
	query = query.Where("product_units.pu_is_del = 0")
	query = query.Where("categorys.category_is_del = 0")

	if *filter.Search != "" {
		query = query.Where("products.pd_code LIKE '%" + *filter.Search + "%' OR products.pd_name LIKE '%" + *filter.Search + "%' ")
	}

	if isCount == true {
		offset := filter.ActivePage * filter.PerPage
		query = query.Limit(filter.PerPage)
		query = query.Offset(offset)
	}

	if err = query.Scan(&product).Error; err != nil {
		return err
	}
	return nil
}

func GetCourseSubListById(courselistsId int, cl *structs.CourseSubList) (err error) {
	query := configs.DB1.Table("course_lists")
	query = query.Select("course_lists.id, course_lists.course_set_id, course_sets.course_id, courses.shop_id, courses.category_id, courses.course_type_id, courses.course_code, courses.course_name, courses.course_unit,courses.course_opd,courses.course_ipd ,courses.course_cost ,courses.course_fee_df ,courses.course_fee ,courses.course_is_active ,courses.course_is_del ,courses.course_create ,courses.course_update, categorys.category_type_id, categorys.category_name,	category_types.category_type_th")
	query = query.Joins("JOIN course_sets ON course_lists.course_set_id = course_sets.id")
	query = query.Joins("JOIN courses ON course_sets.course_id = courses.id")
	query = query.Joins("JOIN categorys ON courses.category_id = categorys.id")
	query = query.Joins("JOIN category_types ON categorys.category_type_id = category_types.id")
	query = query.Where("course_lists.id = ?", courselistsId)
	if err = query.Scan(&cl).Error; err != nil {
		return err
	}
	return nil
}

func GetCourseProductById(Id int, cp *structs.Course_product) (err error) {
	query := configs.DB1.Table("course_products")
	query = query.Select("course_products.*, products.pd_type_id, products.pd_code, products.pd_name, products.pd_code_acc, products.pd_name_acc")
	query = query.Joins("JOIN products ON course_products.product_id = products.id")
	query = query.Where("course_products.id = ?", Id)
	if err = query.Scan(&cp).Error; err != nil {
		return err
	}
	return nil
}

func DeleteCourse(Id int, obj *structs.CourseDetail) (err error) {
	query := configs.DB1.Table("courses")
	query = query.Where("courses.id = ?", Id)
	query = query.Model(&obj)
	query = query.Updates(map[string]interface{}{"courses.course_is_del": 1})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func DeleteCourseList(Id int, course_set_id int, course_id int) (err error) {
	tx := configs.DB1.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	var cl structs.Obj_Course_list
	if err = tx.Table("course_lists").Where("course_lists.id = ?", Id).Delete(&cl).Error; err != nil {
		tx.Rollback()
		return
	}
	// var cs structs.Obj_Course_set
	// if err = tx.Table("course_sets").Where("course_sets.id = ?", course_set_id).Where("course_sets.course_id = ?", course_id).Delete(&cs).Error; err != nil {
	// 	tx.Rollback()
	// 	return
	// }
	tx.Commit()
	return nil
}

func DeleteCourseProduct(Id int, course_id int, data *structs.Course_product) (err error) {
	query := configs.DB1.Table("course_products")
	query = query.Where("course_products.id = ?", Id)
	query = query.Where("course_products.course_id = ?", course_id)
	query = query.Delete(&data)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func AddLogCourse(log *structs.LogCourse) (err error) {
	query := configs.DBL1.Table("log_course").Create(&log)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetCourseDocNoData(ShopId int, data *structs.DocCourse) (err error) {
	query := configs.DB1.Table("doc_settings")
	query = query.Select("doc_settings.shop_id, doc_settings.course_id_default, doc_settings.course_number_default, doc_settings.course_number_digit, doc_settings.course_type")
	query = query.Where("doc_settings.shop_id = ?", ShopId)
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCourseDocno(ShopId int, data *structs.DocCourse) (err error) {
	query := configs.DB1.Table("doc_settings")
	query = query.Where("doc_settings.shop_id = ?", ShopId)
	query = query.Model(&data)
	query = query.Updates(map[string]interface{}{"doc_settings.course_number_default": data.Course_number_default})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetCourseSetList(Id int, cp *structs.Obj_Course_list) (err error) {
	query := configs.DB1.Table("course_lists")
	query = query.Select("course_lists.id, course_lists.course_set_id, SUM(course_lists.course_list_opd) AS course_list_opd, SUM(course_lists.course_list_ipd) AS course_list_ipd")
	query = query.Where("course_lists.course_set_id = ?", Id)
	if err = query.Scan(&cp).Error; err != nil {
		return err
	}
	return nil
}

// import
func GetCourseCheckImport(shopId int, courseCode string, objQuery *structs.ObjQueryCourse) (err error) {
	query := configs.DB1.Table("courses")
	query = query.Select("courses.*")
	query = query.Where("courses.course_is_del = 0")
	query = query.Where("courses.shop_id = ?", shopId)
	query = query.Where("courses.course_code = ?", courseCode)
	if err = query.Scan(&objQuery).Error; err != nil {
		return err
	}
	return nil
}

func CreateCourseBatch(objCreate *[]Course) (err error) {

	tx := configs.DB1.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.CreateInBatches(&objCreate, 100).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil

}

// Set Id 0
func AddCourseSet(data *structs.Obj_Course_set) (err error) {
	query := configs.DB1.Table("course_sets").Create(&data)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}
