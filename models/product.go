package models

import (
	"encoding/json"
	"fmt"
	"linecrmapi/configs"
	"linecrmapi/structs"
	"time"
)

func ProductSearch(filter structs.PayloadSearchProduct, isCount bool, product *[]structs.ProductList) (err error) {
	query := configs.DB1.Table("products")
	query = query.Select("products.*, categorys.category_type_id, categorys.category_name,	category_types.category_type_th, product_shops.user_id,product_shops.topical_id")
	query = query.Joins("JOIN categorys ON products.category_id = categorys.id")
	query = query.Joins("JOIN category_types ON categorys.category_type_id = category_types.id")
	query = query.Joins("LEFT JOIN product_shops ON product_shops.product_id = products.id &  product_shops.shop_id = ?", filter.Shop_id)
	query = query.Where("products.shop_id = ?", filter.Shop_m_id)
	query = query.Where("products.pd_is_del = 0")
	query = query.Where("categorys.category_is_del = 0")
	query = query.Where("products.pd_type_id IN ?", []int{1, 2})
	if *filter.Search != "" {
		query = query.Where("products.pd_name LIKE '%" + *filter.Search + "%' OR products.pd_code LIKE '%" + *filter.Search + "%' OR products.pd_barcode LIKE '%" + *filter.Search + "%' ")
	}
	if *filter.Is_active != "" {
		query = query.Where("products.pd_is_active = ?", *filter.Is_active)
	}
	if filter.CategoryId != -1 {
		query = query.Where("products.category_id = ?", filter.CategoryId)
	}
	if filter.PdTypeId != -1 {
		query = query.Where("products.pd_type_id = ?", filter.PdTypeId)
	}
	if isCount == true {
		offset := filter.ActivePage * filter.PerPage
		query = query.Limit(filter.PerPage)
		query = query.Offset(offset)
	}
	query = query.Order("LENGTH(products.pd_code) ASC")
	query = query.Order("products.pd_code ASC")
	if err = query.Scan(&product).Error; err != nil {
		return err
	}
	return nil
}

func GetProductAll(Shop_m_id int, product *[]structs.GetProductAll) (err error) {
	query := configs.DB1.Table("products")
	query = query.Select("products.id")
	query = query.Where("products.shop_id = ?", Shop_m_id)
	query = query.Where("products.pd_is_del = 0")
	query = query.Where("products.pd_type_id IN ?", []int{1, 2})
	query = query.Order("products.id ASC")
	if err = query.Scan(&product).Error; err != nil {
		return err
	}
	return nil
}
func GetProductDetail(productId int, shopId int, product *structs.ProductDetail) (err error) {
	query := configs.DB1.Table("products")
	query = query.Select("products.*,categorys.category_type_id,categorys.category_name,category_types.category_type_th,product_shops.topical_id,product_shops.user_id,topicals.topical_name,users.user_fullname,users.user_fullname_en")
	query = query.Joins("JOIN categorys ON products.category_id = categorys.id")
	query = query.Joins("JOIN category_types ON categorys.category_type_id = category_types.id")
	query = query.Joins("JOIN product_shops ON product_shops.product_id = products.id")
	query = query.Joins("LEFT JOIN users ON users.id = product_shops.user_id")
	query = query.Joins("LEFT JOIN topicals ON topicals.id = product_shops.topical_id")
	query = query.Where("products.id = ?", productId)
	query = query.Where("product_shops.shop_id = ?", shopId)
	query = query.Where("products.pd_is_del = 0")
	query = query.Where("categorys.category_is_del = 0")
	if err = query.Scan(&product).Error; err != nil {
		return err
	}
	return nil
}

func CheckProductAdd(shopId int, product_name string, product_code string, product_barcode string, product *[]structs.ProductByID) (err error) {
	query := configs.DB1.Table("products")
	query = query.Select("products.*")
	query = query.Where("products.shop_id = ?", shopId)
	query = query.Where("products.pd_is_del = 0")
	wh := ""
	if product_name != "" {
		wh += "products.pd_name = '" + product_name + "'"
	}
	if product_code != "" {
		if wh != "" {
			wh += " OR "
		}
		wh += "products.pd_code = '" + product_code + "'"
	}
	if product_barcode != "" {
		if wh != "" {
			wh += " OR "
		}
		wh += "products.pd_barcode = '" + product_barcode + "'"
	}
	if wh != "" {
		query = query.Where(wh)
	}
	if err = query.Scan(&product).Error; err != nil {
		return err
	}
	return nil
}

func CheckProductUpdate(shopId int, product_id int, product_barcode string, product *[]structs.ProductByID) (err error) {
	query := configs.DB1.Table("products")
	query = query.Select("products.*")
	query = query.Where("products.shop_id = ?", shopId)
	query = query.Where("products.pd_is_del = 0")
	query = query.Not("products.id = ?", product_id)
	if product_barcode != "" {
		query = query.Where("products.pd_barcode = ?", product_barcode)
	}
	if err = query.Scan(&product).Error; err != nil {
		return err
	}

	return nil
}

func CheckProductStore(shop_id int, product_id int, product *[]structs.OdjProductStore) (err error) {
	query := configs.DB1.Table("product_stores")
	query = query.Select("product_stores.*,shop_stores.shop_id")
	query = query.Joins("JOIN shop_stores ON shop_stores.id = product_stores.shop_store_id")
	query = query.Where("shop_stores.shop_id = ?", shop_id)
	query = query.Where("product_stores.product_id = ?", product_id)
	query = query.Where("product_stores.pds_is_del = 0")
	if err = query.Scan(&product).Error; err != nil {
		return err
	}
	return nil
}

func GetProductCategory(Shop_id int, Type_id int, cate *[]structs.Category) (err error) {
	query := configs.DB1.Table("categorys")
	query = query.Select("categorys.id,categorys.shop_id,categorys.category_type_id,categorys.category_name, category_types.category_type_th, category_types.category_type_en")
	query = query.Joins("JOIN category_types ON categorys.category_type_id = category_types.id")
	query = query.Where("categorys.category_is_del = 0")
	query = query.Where("categorys.shop_id = ?", Shop_id)
	query = query.Where("categorys.category_type_id = ?", Type_id)
	if err = query.Scan(&cate).Error; err != nil {
		return err
	}
	return nil
}

func GetShopByMotherId(shopId int, shopMId int, objResponse *[]structs.GetShopByMotherId) error {
	query := configs.DB1.Table("shops")
	query = query.Select("shops.id as shop_id")
	query = query.Where("shops.id != ?", shopId)
	query = query.Where("shops.shop_mother_id = ?", shopMId)
	query = query.Find(&objResponse)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetProductUnit(productId int, shopId int, pu *[]structs.ProductUnitList) (err error) {
	query := configs.DB1.Table("product_units")
	query = query.Select("product_units.*, product_units.id AS product_units_id, product_shop_prices.category_eclaim_id, product_shop_prices.psp_price_opd, product_shop_prices.psp_price_ipd, product_shop_prices.psp_price_ofc, product_shop_prices.psp_price_lgo, product_shop_prices.psp_price_ucs, product_shop_prices.psp_price_sss, product_shop_prices.psp_price_nhs, product_shop_prices.psp_price_ssi,ref_units.u_name,ref_units.u_name_en")
	query = query.Joins("JOIN ref_units ON ref_units.id = product_units.unit_id")
	query = query.Joins("LEFT JOIN product_shop_prices ON product_shop_prices.product_unit_id = product_units.id")
	query = query.Where("product_units.product_id = ?", productId)
	query = query.Where("product_shop_prices.shop_id = ?", shopId)
	query = query.Where("product_units.pu_is_del = 0")
	query = query.Order("product_units.pu_rate ASC")
	if err = query.Scan(&pu).Error; err != nil {
		return err
	}
	return nil
}

func CheckProductShopPrice(shopId int, productUnitId int, psp *[]structs.ProductShopPrice) (err error) {
	query := configs.DB1.Table("product_shop_prices")
	query = query.Select("product_shop_prices.*")
	query = query.Where("product_shop_prices.shop_id = ?", shopId)
	query = query.Where("product_shop_prices.product_unit_id = ?", productUnitId)
	query = query.Where("product_shop_prices.psp_is_del = 0")
	if err = query.Scan(&psp).Error; err != nil {
		return err
	}
	return nil
}

func AddProductShopPrice(psp *structs.ProductShopPriceAction) (err error) {
	query := configs.DB1
	if err = query.Table("product_shop_prices").Create(&psp).Error; err != nil {
		return err
	}
	return nil
}

func GetUserProduct(shopId int, up *[]structs.UserProduct) (err error) {
	query := configs.DB1.Table("users")
	query = query.Select("users.*")
	query = query.Joins("JOIN user_shops ON user_shops.user_id = users.id")
	query = query.Where("user_shops.shop_id = ?", shopId)
	query = query.Where("user_shops.us_is_active = 1")
	query = query.Where("user_shops.us_invite = 2")
	query = query.Where("users.user_is_active = 1")
	query = query.Order("users.user_fullname ASC")
	if err = query.Scan(&up).Error; err != nil {
		return err
	}
	return nil
}

func GetProductTopical(shopId int, pt *[]structs.ProductTopical) (err error) {
	query := configs.DB1.Table("topicals")
	query = query.Select("topicals.id,topicals.topical_name,topicals.topical_detail")
	query = query.Where("topicals.shop_id = ?", shopId)
	query = query.Where("topicals.topical_is_active = 1")
	query = query.Where("topicals.topical_is_del = 0")
	query = query.Where("topicals.topical_type_id = 1")
	query = query.Order("topicals.topical_name ASC")
	if err = query.Scan(&pt).Error; err != nil {
		return err
	}
	return nil
}

func SearchProductTopical(filter structs.ObjPayloadSearch, pt *[]structs.ProductTopical) (err error) {
	query := configs.DB1.Table("topicals")
	query = query.Select("topicals.id,topicals.topical_name,topicals.topical_detail")
	query = query.Where("topicals.shop_id = ?", filter.Shop_id)

	if *filter.Search != "" {
		query = query.Where("topicals.topical_name LIKE '%" + *filter.Search + "%' ")
	}

	query = query.Where("topicals.topical_is_active = 1")
	query = query.Where("topicals.topical_is_del = 0")
	query = query.Where("topicals.topical_type_id = 1")
	query = query.Order("topicals.topical_name ASC")
	if err = query.Scan(&pt).Error; err != nil {
		return err
	}
	return nil
}

func AddProduct(dataPo *structs.ProductAction, dataPu *[]structs.ObjProductUnit) (err error) {
	tx := configs.DB1.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	// add products
	if err = tx.Table("products").Create(&dataPo).Error; err != nil {
		tx.Rollback()
		fmt.Println(err)
		return err
	}

	// add products unit
	for _, pu := range *dataPu {
		var pua structs.ProductUnitAction
		pua.ProductId = dataPo.Id
		pua.UnitId = pu.UnitId
		pua.PuRate = pu.PuRate
		pua.PuIsDel = 0
		pua.PuCreate = time.Now().Format("2006-01-02 15:04:05")
		pua.PuUpdate = time.Now().Format("2006-01-02 15:04:05")
		if err = tx.Table("product_units").Create(&pua).Error; err != nil {
			tx.Rollback()
			fmt.Println(err)
			return err
		}

		// add products price
		var psp structs.ProductShopPriceAction
		psp.ShopId = dataPo.ShopId
		psp.ProductUnitId = pua.Id
		psp.PspPriceOpd = *pu.PspPriceOpd
		psp.PspPriceIpd = *pu.PspPriceIpd
		psp.PspPriceOfc = *pu.PspPriceOfc
		psp.PspPriceLgo = *pu.PspPriceLgo
		psp.PspPriceUcs = *pu.PspPriceUcs
		psp.PspPriceSss = *pu.PspPriceSss
		psp.PspPriceNhs = *pu.PspPriceNhs
		psp.PspPriceSsi = *pu.PspPriceSsi
		psp.PspIsDefault = 0
		psp.PspIsDel = 0
		psp.PspCreate = time.Now().Format("2006-01-02 15:04:05")
		psp.PspUpdate = time.Now().Format("2006-01-02 15:04:05")

		var Category_eclaim_id *int
		if *pu.CategoryEclaimId == 0 {
			Category_eclaim_id = nil
		} else {
			Category_eclaim_id = pu.CategoryEclaimId
		}
		psp.CategoryEclaimId = Category_eclaim_id

		if err = tx.Table("product_shop_prices").Create(&psp).Error; err != nil {
			tx.Rollback()
			fmt.Println(err)
			return err
		}
	}

	tx.Commit()
	return nil
}

func UpdateProduct(shop_id int, product_id int, data *structs.ProductAction) (err error) {

	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	delete(inInterface, "shop_id")
	delete(inInterface, "id")
	delete(inInterface, "pd_create")

	query := configs.DB1.Table("products")
	if data.PdIsActive == 0 {
		query = query.Where("products.shop_id = ?", shop_id)
		query = query.Where("products.id = ?", product_id)
		query = query.Model(&data)
		query = query.Updates(map[string]interface{}{"products.pd_is_active": 0})
	}
	query = query.Where("id = ?", product_id)
	query = query.Model(&data)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func AddProductShop(dataPs *structs.ProductShopAction) (err error) {
	query := configs.DB1.Table("product_shops").Create(&dataPs)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetProductShopByPDSP(productId int, shopId int, ps *[]structs.ProductShopAction) (err error) {
	query := configs.DB1.Table("product_shops")
	query = query.Select("product_shops.*")
	query = query.Where("product_shops.product_id = ?", productId)
	query = query.Where("product_shops.shop_id = ?", shopId)
	if err = query.Scan(&ps).Error; err != nil {
		return err
	}
	return nil
}

func GetProductShop(productId int, shopId int, ps *structs.ProductShopAction) (err error) {
	query := configs.DB1.Table("product_shops")
	query = query.Select("product_shops.*")
	query = query.Where("product_shops.product_id = ?", productId)
	query = query.Where("product_shops.shop_id = ?", shopId)
	if err = query.Scan(&ps).Error; err != nil {
		return err
	}
	return nil
}

func GetProductShopByID(productId int, shopId int, ps *structs.ProductShop) (err error) {
	query := configs.DB1.Table("product_shops")
	query = query.Select("product_shops.*,")
	query = query.Joins("LEFT JOIN users ON users.id = product_shops.user_id")
	query = query.Joins("LEFT JOIN topicals ON topicals.id = product_shops.topical_id")
	query = query.Where("product_shops.product_id = ?", productId)
	query = query.Where("product_shops.shop_id = ?", shopId)
	if err = query.Scan(&ps).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProductShop(id int, dataPs *structs.ProductShopAction) (err error) {

	var inInterface map[string]interface{}
	in, _ := json.Marshal(&dataPs)
	json.Unmarshal(in, &inInterface)
	delete(inInterface, "id")
	delete(inInterface, "shop_id")
	delete(inInterface, "product_id")
	delete(inInterface, "ps_create")

	query := configs.DB1.Table("product_shops")
	query = query.Where("product_shops.id = ?", id)
	query = query.Model(&dataPs)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateProductUnit(shopId int, product_id int, data *[]structs.ObjProductUnit) (err error) {
	tx := configs.DB1.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	for _, pu := range *data {
		var pua structs.ProductUnitAction
		if pu.Id > 0 {
			pua.PuUpdate = time.Now().Format("2006-01-02 15:04:05")
			if err = tx.Table("product_units").Where("id = ?", pu.Id).Updates(&pua).Error; err != nil {
				tx.Rollback()
				return
			}

			// update products price
			var get_psp structs.ProductShopPriceAction
			query := configs.DB1.Table("product_shop_prices")
			query = query.Select("product_shop_prices.*")
			query = query.Where("product_shop_prices.shop_id = ?", shopId)
			query = query.Where("product_shop_prices.product_unit_id = ?", pu.Id)
			query = query.Where("product_shop_prices.psp_is_del = 0")
			if err = query.Scan(&get_psp).Error; err != nil {
				return err
			}

			if get_psp.Id > 0 {
				// var psp structs.ProductShopPriceAction
				// psp.PspPriceOpd = pu.PspPriceOpd
				// psp.PspPriceIpd = pu.PspPriceIpd
				// psp.PspUpdate = time.Now().Format("2006-01-02 15:04:05")

				var Category_eclaim_id *int
				if *pu.CategoryEclaimId == 0 {
					Category_eclaim_id = nil
				} else {
					Category_eclaim_id = pu.CategoryEclaimId
				}

				objQueryUpdateProductPrice := map[string]interface{}{
					"psp_price_opd":      pu.PspPriceOpd,
					"psp_price_ipd":      pu.PspPriceIpd,
					"psp_price_ofc":      pu.PspPriceOfc,
					"psp_price_lgo":      pu.PspPriceLgo,
					"psp_price_ucs":      pu.PspPriceUcs,
					"psp_price_sss":      pu.PspPriceSss,
					"psp_price_nhs":      pu.PspPriceNhs,
					"psp_price_ssi":      pu.PspPriceSsi,
					"category_eclaim_id": Category_eclaim_id,
					"psp_update":         time.Now().Format("2006-01-02 15:04:05"),
				}

				if err = tx.Table("product_shop_prices").Where("id = ?", get_psp.Id).Updates(&objQueryUpdateProductPrice).Error; err != nil {
					tx.Rollback()
					return
				}
			} else {
				// add products price
				var psp structs.ProductShopPriceAction
				psp.ShopId = shopId
				psp.ProductUnitId = pua.Id
				psp.PspPriceOpd = *pu.PspPriceOpd
				psp.PspPriceIpd = *pu.PspPriceIpd
				psp.PspPriceOfc = *pu.PspPriceOfc
				psp.PspPriceLgo = *pu.PspPriceLgo
				psp.PspPriceUcs = *pu.PspPriceUcs
				psp.PspPriceSss = *pu.PspPriceSss
				psp.PspPriceNhs = *pu.PspPriceNhs
				psp.PspPriceSsi = *pu.PspPriceSsi
				psp.PspIsDefault = 0
				psp.PspIsDel = 0
				psp.PspCreate = time.Now().Format("2006-01-02 15:04:05")
				psp.PspUpdate = time.Now().Format("2006-01-02 15:04:05")

				var Category_eclaim_id *int
				if *pu.CategoryEclaimId == 0 {
					Category_eclaim_id = nil
				} else {
					Category_eclaim_id = pu.CategoryEclaimId
				}
				psp.CategoryEclaimId = Category_eclaim_id

				if err = tx.Table("product_shop_prices").Create(&psp).Error; err != nil {
					tx.Rollback()
					return
				}
			}

		} else {
			pua.ProductId = product_id
			pua.UnitId = pu.UnitId
			pua.PuRate = pu.PuRate
			pua.PuIsDel = 0
			pua.PuCreate = time.Now().Format("2006-01-02 15:04:05")
			pua.PuUpdate = time.Now().Format("2006-01-02 15:04:05")
			if err = tx.Table("product_units").Create(&pua).Error; err != nil {
				tx.Rollback()
				return
			}

			// add products price
			var psp structs.ProductShopPriceAction
			psp.ShopId = shopId
			psp.ProductUnitId = pua.Id
			psp.PspPriceOpd = *pu.PspPriceOpd
			psp.PspPriceIpd = *pu.PspPriceIpd
			psp.PspPriceOfc = *pu.PspPriceOfc
			psp.PspPriceLgo = *pu.PspPriceLgo
			psp.PspPriceUcs = *pu.PspPriceUcs
			psp.PspPriceSss = *pu.PspPriceSss
			psp.PspPriceNhs = *pu.PspPriceNhs
			psp.PspPriceSsi = *pu.PspPriceSsi
			psp.PspIsDefault = 0
			psp.PspIsDel = 0
			psp.PspCreate = time.Now().Format("2006-01-02 15:04:05")
			psp.PspUpdate = time.Now().Format("2006-01-02 15:04:05")

			var Category_eclaim_id *int
			if *pu.CategoryEclaimId == 0 {
				Category_eclaim_id = nil
			} else {
				Category_eclaim_id = pu.CategoryEclaimId
			}
			psp.CategoryEclaimId = Category_eclaim_id

			if err = tx.Table("product_shop_prices").Create(&psp).Error; err != nil {
				tx.Rollback()
				return
			}
		}
	}
	tx.Commit()
	return nil
}

func GetProductShopPrice(shopId int, productUnitId int, pu *[]structs.ProductShopPriceAction) (err error) {
	query := configs.DB1.Table("product_shop_prices")
	query = query.Select("product_shop_prices.*")
	query = query.Where("product_shop_prices.shop_id = ?", shopId)
	query = query.Where("product_shop_prices.product_unit_id = ?", productUnitId)
	query = query.Where("product_shop_prices.psp_is_del = 0")
	if err = query.Scan(&pu).Error; err != nil {
		return err
	}
	return nil
}

func GetProductUnitById(productUnitId int, pua *structs.ProductUnitAction) (err error) {
	query := configs.DB1.Table("product_units")
	query = query.Select("product_units.*")
	query = query.Where("product_units.id = ?", productUnitId)
	query = query.Where("product_units.pu_is_del = 0")
	if err = query.Scan(&pua).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProduct(Id int, obj *structs.ProductDetail) (err error) {
	query := configs.DB1.Table("products")
	query = query.Where("products.id = ?", Id)
	query = query.Model(&obj)
	query = query.Updates(map[string]interface{}{"products.pd_is_del": 1})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func DeleteProductUnit(Id int, product_id int, obj *structs.ProductUnitAction) (err error) {
	query := configs.DB1.Table("product_units")
	query = query.Where("product_units.id = ?", Id)
	query = query.Where("product_units.product_id = ?", product_id)
	query = query.Model(&obj)
	query = query.Updates(map[string]interface{}{"product_units.pu_is_del": 1})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func AddLogProduct(log *structs.LogProduct) (err error) {
	query := configs.DBL1.Table("log_product").Create(&log)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetDocNoProduct(ShopId int, data *structs.DocNoProduct) (err error) {
	query := configs.DB1.Table("doc_settings")
	query = query.Select("doc_settings.shop_id, doc_settings.product_id_default, doc_settings.product_number_default, doc_settings.product_number_digit, doc_settings.product_type")
	query = query.Where("doc_settings.shop_id = ?", ShopId)
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func UpdateDocNoProduct(ShopId int, data *structs.DocNoProduct) (err error) {
	query := configs.DB1.Table("doc_settings")
	query = query.Where("doc_settings.shop_id = ?", ShopId)
	query = query.Model(&data)
	query = query.Updates(map[string]interface{}{"doc_settings.product_number_default": data.ProductNumberDefault})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetUnit(un *[]structs.Unit) (err error) {
	query := configs.DB1.Table("ref_units")
	query = query.Select("ref_units.id,ref_units.u_name,ref_units.u_name_en,ref_units.u_sort,ref_units.u_is_del")
	query = query.Where("ref_units.u_is_del = 0")
	query = query.Order("ref_units.u_sort ASC")

	if err = query.Scan(&un).Error; err != nil {
		return err
	}
	return nil
}

func GetUnitList(filter structs.PayloadSearchUnit, un *[]structs.Unit) (err error) {
	query := configs.DB1.Table("ref_units")
	query = query.Select("ref_units.id,ref_units.u_name,ref_units.u_name_en,ref_units.u_sort,ref_units.u_is_del")
	query = query.Where("ref_units.u_is_del = 0")
	if *filter.Search != "" {
		query = query.Where("ref_units.u_name LIKE '%" + *filter.Search + "%' OR ref_units.u_name_en LIKE '%" + *filter.Search + "%'")
	}
	query = query.Order("ref_units.u_sort ASC")
	if err = query.Scan(&un).Error; err != nil {
		return err
	}
	return nil
}

// import
func GetProductCheckImport(shopId int, pdCode string, pdBarcode string, objQuery *structs.ObjQueryProduct) (err error) {
	query := configs.DB1.Table("products")
	query = query.Select("products.*")
	query = query.Where("products.shop_id = ?", shopId)
	query = query.Where("products.pd_code = ?", pdCode)
	query = query.Where("products.pd_barcode = ?", pdBarcode)
	query = query.Where("products.pd_is_del = ?", 0)
	if err = query.Scan(&objQuery).Error; err != nil {
		return err
	}
	return nil
}

func GetCheckImportPdCode(shopId int, pdCodes []string, objQuery *[]structs.ObjQueryCheckImportProduct) error {
	query := configs.DB1.Table("products")
	query = query.Select("products.pd_code, products.pd_barcode")
	query = query.Where("products.shop_id = ?", shopId)
	query = query.Where("products.pd_code IN ?", pdCodes)
	query = query.Where("products.pd_is_del = ?", 0)
	query = query.Find(&objQuery)
	return query.Error
}

func GetCheckImportPdBarcode(shopId int, pdBarcodes []string, objQuery *[]structs.ObjQueryCheckImportProduct) error {
	query := configs.DB1.Table("products")
	query = query.Select("products.pd_code, products.pd_barcode")
	query = query.Where("products.shop_id = ?", shopId)
	query = query.Where("products.pd_barcode IN ?", pdBarcodes)
	query = query.Where("products.pd_is_del = ?", 0)
	query = query.Find(&objQuery)
	return query.Error
}

func GetCheckUnit(objQuery *[]structs.ObjQueryCheckUnit) error {
	query := configs.DB1.Table("ref_units")
	query = query.Select("ref_units.id, ref_units.u_name, ref_units.u_name_en")
	query = query.Where("ref_units.u_is_del = ?", 0)
	query = query.Find(&objQuery)
	return query.Error
}

func CreateProductImport(objCreate []structs.ObjCheckExcelProduct) (err error) {

	tx := configs.DB1.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for _, row := range objCreate {

		// product
		objCreateProduct := Product{
			ShopId:        row.ShopId,
			CategoryId:    row.CategoryId,
			PdTypeId:      row.PdTypeId,
			PdCode:        row.PdCode,
			PdName:        row.PdName,
			PdCodeAcc:     row.PdCodeAcc,
			PdNameAcc:     row.PdNameAcc,
			PdImage:       "",
			PdDescription: "",
			PdBarcode:     row.PdBarcode,
			PdIsSerial:    0,
			PdNarcotic2:   0,
			PdNarcotic3:   0,
			PdNarcotic4:   0,
			PdNarcotic5:   0,
			DrugDirection: row.DrugDirection,
			PdAmountNoti:  row.PdAmountNoti,
			PdExpireNoti:  row.PdExpireNoti,
			PdIsOver:      0,
			PdIsActive:    1,
			PdIsDel:       0,
			PdCreate:      time.Now().Format("2006-01-02 15:04:05"),
			PdUpdate:      time.Now().Format("2006-01-02 15:04:05"),
		}
		if err = tx.Table("products").Create(&objCreateProduct).Error; err != nil {
			tx.Rollback()
			return err
		}

		// product unit
		objCreateProductUnit := ProductUnit{
			ProductId: objCreateProduct.ID,
			UnitId:    row.UnitId,
			PuRate:    1,
			PuIsDel:   0,
			PuCreate:  time.Now().Format("2006-01-02 15:04:05"),
			PuUpdate:  time.Now().Format("2006-01-02 15:04:05"),
		}
		if err = tx.Table("product_units").Create(&objCreateProductUnit).Error; err != nil {
			tx.Rollback()
			return err
		}

		// product shop price
		objCreateProductShopPrice := ProductShopPrice{
			ShopId:        row.ShopId,
			ProductUnitId: objCreateProductUnit.ID,
			PspPriceOpd:   row.OpdPrice,
			PspPriceIpd:   row.IpdPrice,
			PspIsDefault:  0,
			PspIsDel:      0,
			PspCreate:     time.Now().Format("2006-01-02 15:04:05"),
			PspUpdate:     time.Now().Format("2006-01-02 15:04:05"),
		}
		if err = tx.Table("product_shop_prices").Create(&objCreateProductShopPrice).Error; err != nil {
			tx.Rollback()
			return err
		}

		// product shop
		objCreateProductShop := ProductShop{
			ShopId:    row.ShopId,
			ProductId: objCreateProduct.ID,
			PsCreate:  time.Now().Format("2006-01-02 15:04:05"),
			PsUpdate:  time.Now().Format("2006-01-02 15:04:05"),
		}
		if err = tx.Table("product_shops").Create(&objCreateProductShop).Error; err != nil {
			tx.Rollback()
			return err
		}

	}

	tx.Commit()

	return nil

}
