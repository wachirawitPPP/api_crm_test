package models

import (
	"linecrmapi/configs"
	"linecrmapi/structs"

	"gorm.io/gorm"
)

func GetItemProductList(filter structs.ObjPayloaItem, isCount bool, product *[]structs.ItemProduct) error {
	query := configs.DB1.Table("products")
	if isCount {
		query = query.Select(`
			products.id AS product_id, 
			products.*,shops.shop_code,shops.shop_name,shops.shop_name_en
		`)
	} else {
		query = query.Select("products.id")
	}
	query = query.Joins("INNER JOIN shops ON products.shop_id = shops.id")
	query = query.Where("products.pd_is_line = ?", 1).
		Where("products.pd_is_active = ?", 1).
		Where("products.pd_is_del = ?", 0)

	query = applyItemProductFilters(query, filter)
	query = query.Order("products.pd_code ASC")
	return query.Scan(product).Error
}

func applyItemProductFilters(query *gorm.DB, filter structs.ObjPayloaItem) *gorm.DB {
	if filter.Search != nil && *filter.Search != "" {
		query = query.Where(
			`products.pd_code LIKE ? OR 
			 products.pd_name LIKE ? OR 
			 products.pd_barcode LIKE ?`,
			"%"+*filter.Search+"%",
			"%"+*filter.Search+"%",
			"%"+*filter.Search+"%",
		)
	}
	return query
}

func GetShopStoreByIdType1(shopId int, objResponse *ShopStore) error {
	query := configs.DB1.Table("shop_stores")
	query = query.Where("ss_is_active = ?", 1)
	query = query.Where("ss_type_id = ?", 1)
	query = query.Where("shop_id = ?", shopId)
	query = query.Find(&objResponse)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetItemProductSetId(Product_id int, Shop_id int, product *[]structs.ItemProductSub) (err error) {
	query := configs.DB1.Table("products")
	query = query.Select("products.id AS id,products.id AS product_id,product_stores.id AS product_store_id, 1 AS pu_amount,product_units.id AS product_units_id,product_stores.pds_cost,products.pd_code,products.pd_name,products.drug_direction,product_units.pu_rate,topicals.id AS topical_id,products.drug_direction,ref_units.u_name,product_shop_prices.psp_price_opd,product_shop_prices.psp_price_ipd, product_shop_prices.psp_price_ofc, product_shop_prices.psp_price_lgo, product_shop_prices.psp_price_ucs, product_shop_prices.psp_price_sss, product_shop_prices.psp_price_nhs, product_shop_prices.psp_price_ssi ,products.pd_image_1,products.pd_image_2,products.pd_image_3,products.pd_image_4,products.pd_detail")
	query = query.Joins("INNER JOIN product_stores ON products.id = product_stores.product_id")
	query = query.Joins("INNER JOIN shop_stores ON product_stores.shop_store_id = shop_stores.id")
	query = query.Joins("INNER JOIN shops ON shop_stores.shop_id = shops.id")
	query = query.Joins("INNER JOIN product_units ON products.id = product_units.product_id")
	query = query.Joins("INNER JOIN ref_units ON product_units.unit_id = ref_units.id")
	query = query.Joins("INNER JOIN product_shop_prices ON product_units.id = product_shop_prices.product_unit_id")
	query = query.Joins("LEFT JOIN product_shops ON product_shops.product_id = products.id")
	query = query.Joins("LEFT JOIN topicals ON topicals.id = product_shops.topical_id AND topicals.topical_is_active = 1 AND topicals.topical_is_del = 0")
	query = query.Where("products.id = ?", Product_id)
	query = query.Where("shop_stores.ss_type_id = ?", 1)
	query = query.Where("product_units.pu_rate = ?", 1)
	query = query.Where("product_units.pu_is_del = ?", 0)
	query = query.Where("product_shop_prices.psp_is_del = ?", 0)
	// query = query.Where("product_shop_prices.psp_is_default = ?", 1)
	query = query.Where("shop_stores.shop_id = ?", Shop_id)
	query = query.Where("product_shop_prices.shop_id = ?", Shop_id)
	query = query.Where("product_shops.shop_id = ?", Shop_id)
	if err = query.Scan(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProductStoreBalance(ShopStoreId int, product_id int, objQuery *structs.ObjQueryProductStoreBalance) error {
	query := configs.DB1.Table("product_stores")
	query = query.Select("IFNULL(SUM( product_store_orders.pdso_total ),0) AS pds_balance")
	query = query.Joins("LEFT JOIN product_store_orders on product_store_orders.product_store_id = product_stores.id")
	query = query.Joins("LEFT JOIN products on products.id = product_stores.product_id")
	query = query.Where("product_stores.pds_is_del = ?", 0)
	query = query.Where("product_stores.shop_store_id = ?", ShopStoreId)
	query = query.Where("product_stores.product_id = ?", product_id)
	query = query.Where("product_stores.pds_is_active = ?", 1)
	query = query.Group("product_stores.id,product_stores.product_id")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetItemProductUnitList(Product_id int, Shop_id int, unit *[]structs.ItemProductUnit) (err error) {
	query := configs.DB1.Table("products")
	query = query.Select("products.id AS id,products.id AS product_id,1 AS pu_amount,product_units.id AS product_units_id,products.pd_code,products.pd_name,product_units.pu_rate,ref_units.u_name,product_shop_prices.psp_price_opd,product_shop_prices.psp_price_ipd, product_shop_prices.psp_price_ofc, product_shop_prices.psp_price_lgo, product_shop_prices.psp_price_ucs, product_shop_prices.psp_price_sss, product_shop_prices.psp_price_nhs, product_shop_prices.psp_price_ssi")
	query = query.Joins("INNER JOIN product_units ON products.id = product_units.product_id")
	query = query.Joins("INNER JOIN ref_units ON product_units.unit_id = ref_units.id")
	query = query.Joins("INNER JOIN product_shop_prices ON product_units.id = product_shop_prices.product_unit_id")
	query = query.Where("product_shop_prices.shop_id = ?", Shop_id)
	query = query.Where("products.id = ?", Product_id)
	query = query.Where("product_units.pu_is_del = ?", 0)
	query = query.Where("product_shop_prices.psp_is_del = ?", 0)
	if err = query.Scan(&unit).Error; err != nil {
		return err
	}
	return nil
}

func GetItemTopicalId(topical_id int, objQuery *structs.ItemTopical) (err error) {
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

// Course /////////////////////////////////////////////////////////////////////////////////////////////////////////
func GetItemCourseList(filter structs.ObjPayloaItem, isCount bool, courses *[]structs.ItemCourse) (err error) {
	query := configs.DB1.Table("courses")
	query = query.Select("courses.*,shops.shop_code,shops.shop_name,shops.shop_name_en, categorys.category_type_id, categorys.category_name,	category_types.category_type_th")
	query = query.Joins("JOIN shops ON courses.shop_id = shops.id")
	query = query.Joins("JOIN categorys ON courses.category_id = categorys.id")
	query = query.Joins("JOIN category_types ON categorys.category_type_id = category_types.id")
	query = query.Where("courses.course_is_line = ?", 1)
	query = query.Where("courses.course_is_active = 1")
	query = query.Where("courses.course_is_del = 0")
	query = query.Where("categorys.category_is_del = 0")

	if *filter.Search != "" {
		query = query.Where("courses.course_code LIKE '%" + *filter.Search + "%' OR courses.course_name LIKE '%" + *filter.Search + "%'")
	}

	query = query.Order("courses.course_code ASC")
	if err = query.Scan(&courses).Error; err != nil {
		return err
	}
	return nil
}

func GetItemCourseProduct(course_id int, shop_id int, cp *[]structs.ItemProductSubSet) (err error) {
	query := configs.DB1.Table("course_products")
	query = query.Select("products.id AS id,products.id AS product_id,product_stores.id AS product_store_id, course_products.cp_amount AS pu_amount,product_units.id AS product_units_id,product_stores.pds_cost,products.pd_code,products.pd_name,product_units.pu_rate,ref_units.u_name,product_shop_prices.psp_price_opd,product_shop_prices.psp_price_ipd, product_shop_prices.psp_price_ofc, product_shop_prices.psp_price_lgo, product_shop_prices.psp_price_ucs, product_shop_prices.psp_price_sss, product_shop_prices.psp_price_nhs, product_shop_prices.psp_price_ssi")
	query = query.Joins("INNER JOIN products ON products.id = course_products.product_id")
	query = query.Joins("INNER JOIN product_stores ON products.id = product_stores.product_id")
	query = query.Joins("INNER JOIN shop_stores ON product_stores.shop_store_id = shop_stores.id")
	query = query.Joins("INNER JOIN shops ON shop_stores.shop_id = shops.id")
	query = query.Joins("INNER JOIN product_units ON products.id = product_units.product_id")
	query = query.Joins("INNER JOIN ref_units ON product_units.unit_id = ref_units.id")
	query = query.Joins("INNER JOIN product_shop_prices ON product_units.id = product_shop_prices.product_unit_id")
	query = query.Where("product_units.pu_rate = ?", 1)
	query = query.Where("product_units.pu_is_del = ?", 0)
	query = query.Where("shop_stores.ss_type_id = ?", 1)
	query = query.Where("shop_stores.shop_id = ?", shop_id)
	query = query.Where("product_shop_prices.shop_id = ?", shop_id)
	query = query.Where("product_shop_prices.psp_is_del = ?", 0)
	query = query.Where("products.pd_is_active = ?", 1)
	query = query.Where("products.pd_is_del = ?", 0)
	query = query.Where("course_products.course_id = ?", course_id)
	query = query.Where("course_products.cp_is_active = ?", 1)
	query = query.Where("course_products.cp_is_del = ?", 0)
	if err = query.Scan(&cp).Error; err != nil {
		return err
	}
	return nil
}

func GetItemProductSetList(Product_id int, Shop_id int, product *[]structs.ItemProductSet) (err error) {
	query := configs.DB1.Table("product_sets")
	query = query.Select("product_lists.product_id AS id,topicals.id AS topical_id,product_lists.product_amount,product_lists.product_list_opd,product_lists.product_list_ipd,products.pd_image_1,products.pd_image_2,products.pd_image_3,products.pd_image_4,products.pd_detail")
	query = query.Joins("JOIN product_lists ON product_sets.id = product_lists.product_set_id")
	query = query.Joins("JOIN products ON product_lists.product_id = products.id")

	query = query.Joins("LEFT JOIN product_shops ON product_shops.product_id = products.id")
	query = query.Joins("LEFT JOIN topicals ON topicals.id = product_shops.topical_id AND topicals.topical_is_active = 1 AND topicals.topical_is_del = 0")

	query = query.Where("product_sets.product_id = ?", Product_id)
	query = query.Where("product_sets.shop_id = ?", Shop_id)
	query = query.Where("product_shops.shop_id = ?", Shop_id)
	query = query.Where("products.pd_is_active = ?", 1)
	query = query.Where("products.pd_is_del = ?", 0)
	query = query.Order("products.pd_code ASC")
	if err = query.Scan(&product).Error; err != nil {
		return err
	}
	return nil
}

func GetItemCourseIdSet(course_id int, course *[]structs.ItemCourseSub) (err error) {
	query := configs.DB1.Table("course_sets")
	query = query.Select("courses.*,course_lists.course_list_qtyset,course_lists.course_list_opd,course_lists.course_list_ipd")
	query = query.Joins("JOIN course_lists ON course_sets.id = course_lists.course_set_id")
	query = query.Joins("JOIN courses ON course_lists.course_id = courses.id")
	query = query.Where("course_sets.course_id = ?", course_id)
	if err = query.Scan(&course).Error; err != nil {
		return err
	}
	return nil
}

// Checking /////////////////////////////////////////////////////////////////////////////////////////////////////////
func GetItemCheckingList(filter structs.ObjPayloaItem, isCount bool, checking *[]structs.ItemChecking) error {
	query := configs.DB1.Table("checkings").
		Joins("JOIN shops ON checkings.shop_id = shops.id").
		Joins("JOIN categorys ON checkings.category_id = categorys.id").
		Joins("JOIN category_types ON categorys.category_type_id = category_types.id")
	if isCount {
		query = query.Select(`
			checkings.*,shops.shop_code,shops.shop_name,shops.shop_name_en, 
			1 AS checking_amount,
			categorys.category_type_id, 
			categorys.category_name,
			category_types.category_type_th
		`)
	} else {
		query = query.Select("checkings.id")
	}
	query = query.Where("checkings.checking_is_line = ?", 1).
		Where("checkings.checking_is_active = ?", 1).
		Where("checkings.checking_is_del = ?", 0).
		Where("categorys.category_is_del = ?", 0)
	if filter.Search != nil && *filter.Search != "" {
		query = query.Where(
			"checkings.checking_name LIKE ? OR checkings.checking_code LIKE ?",
			"%"+*filter.Search+"%",
			"%"+*filter.Search+"%",
		)
	}
	query = query.Order("checkings.checking_code DESC, checkings.checking_create DESC")
	return query.Scan(checking).Error
}

func GetItemCheckingProduct(checking_id int, shop_id int, cp *[]structs.ItemProductSubSet) (err error) {
	query := configs.DB1.Table("checking_products")
	query = query.Select("products.id AS id,checking_products.cip_amount AS 'qty_set',CONCAT(checkings.checking_code,':',checkings.checking_name) AS label,products.id AS product_id,product_stores.id AS product_store_id, checking_products.cip_amount AS 'pu_amount',product_units.id AS product_units_id,product_stores.pds_cost,products.pd_code,products.pd_name,product_units.pu_rate,ref_units.u_name,product_shop_prices.psp_price_opd,product_shop_prices.psp_price_ipd,topicals.id AS topical_id,topicals.topical_name,topicals.topical_detail,products.drug_direction, product_shop_prices.psp_price_ofc, product_shop_prices.psp_price_lgo, product_shop_prices.psp_price_ucs, product_shop_prices.psp_price_sss, product_shop_prices.psp_price_nhs, product_shop_prices.psp_price_ssi")
	query = query.Joins("INNER JOIN checkings ON checkings.id = checking_products.checking_id")
	query = query.Joins("INNER JOIN products ON products.id = checking_products.product_id")
	query = query.Joins("INNER JOIN product_stores ON products.id = product_stores.product_id")
	query = query.Joins("INNER JOIN shop_stores ON product_stores.shop_store_id = shop_stores.id")
	query = query.Joins("INNER JOIN shops ON shop_stores.shop_id = shops.id")
	query = query.Joins("INNER JOIN product_units ON products.id = product_units.product_id")
	query = query.Joins("INNER JOIN ref_units ON product_units.unit_id = ref_units.id")
	query = query.Joins("INNER JOIN product_shop_prices ON product_units.id = product_shop_prices.product_unit_id")
	query = query.Joins("LEFT JOIN product_shops ON product_shops.product_id = products.id")
	query = query.Joins("LEFT JOIN topicals ON topicals.id = product_shops.topical_id AND topicals.topical_is_active = 1 AND topicals.topical_is_del = 0")
	query = query.Where("product_units.pu_rate = ?", 1)
	query = query.Where("product_units.pu_is_del = ?", 0)
	query = query.Where("shop_stores.ss_type_id = ?", 1)
	query = query.Where("shop_stores.shop_id = ?", shop_id)
	query = query.Where("product_shop_prices.shop_id = ?", shop_id)
	query = query.Where("product_shops.shop_id = ?", shop_id)
	query = query.Where("product_shop_prices.psp_is_del = ?", 0)
	query = query.Where("products.pd_is_active = ?", 1)
	query = query.Where("products.pd_is_del = ?", 0)
	query = query.Where("checking_products.checking_id = ?", checking_id)
	query = query.Where("checking_products.cip_is_active = ?", 1)
	query = query.Where("checking_products.cip_is_del = ?", 0)
	if err = query.Scan(&cp).Error; err != nil {
		return err
	}
	return nil
}

func GetItemCheckingIdSet(checking_id int, checking *[]structs.ItemCheckingSub) (err error) {
	query := configs.DB1.Table("checking_sets")
	query = query.Select("checkings.*,checking_lists.checking_list_ipd,checking_lists.checking_list_opd")
	query = query.Joins("JOIN checking_lists ON checking_sets.id = checking_lists.checking_set_id")
	query = query.Joins("JOIN checkings ON checking_lists.checking_id = checkings.id")
	query = query.Where("checking_sets.checking_id = ?", checking_id)
	if err = query.Scan(&checking).Error; err != nil {
		return err
	}
	return nil
}
