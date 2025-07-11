package models

import (
	"linecrmapi/configs"
	"linecrmapi/structs"
	"time"
)

func AddLogInvoice(log *structs.LogInvoice) (err error) {
	query := configs.DBL1.Table("log_invoices").Create(&log)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetOrderId(orderId int, dt *structs.InvoiceOrder) (err error) {
	query := configs.DB1.Table("orders")
	query = query.Select("orders.*")
	query = query.Where("orders.id = ?", orderId)
	if err = query.Scan(&dt).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderDetailId(orderId int, dt *[]structs.InvoiceOrderDetail) (err error) {
	query := configs.DB1.Table("order_details")
	query = query.Select("order_details.*")
	query = query.Where("order_details.order_id = ?", orderId)
	query = query.Where("order_details.ord_is_active = ?", 1)
	if err = query.Scan(&dt).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderInvoiceTags(orderId int, dt *[]structs.OrderTags) (err error) {
	query := configs.DB1.Table("order_tags")
	query = query.Select("order_tags.*, tags.tag_name")
	query = query.Where("order_tags.order_id = ?", orderId)

	query = query.Joins("INNER JOIN tags ON order_tags.tags_id = tags.id")

	if err = query.Scan(&dt).Error; err != nil {
		return err
	}
	return nil
}

func AddInvoice(dataH *structs.InvoiceOrder, dataD *[]structs.InvoiceOrderDetail, Code string, User_id int, dpmId int) (id int, err error) {
	tx := configs.DB1.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return 0, err
	}

	// add checkings
	objH := structs.Invoice{
		Shop_id:              dataH.Shop_id,
		User_id:              User_id,
		Customer_id:          dataH.Customer_id,
		Customer_online_id:   dataH.Customer_online_id,
		Queue_id:             dataH.Queue_id,
		Order_id:             dataH.Id,
		Inv_code:             Code,
		Inv_fullname:         dataH.Or_fullname,
		Inv_tel:              dataH.Or_tel,
		Inv_email:            dataH.Or_email,
		Inv_address:          dataH.Or_address,
		Inv_district:         dataH.Or_district,
		Inv_amphoe:           dataH.Or_amphoe,
		Inv_province:         dataH.Or_province,
		Inv_zipcode:          dataH.Or_zipcode,
		Inv_comment:          dataH.Or_comment,
		Inv_total_price:      dataH.Or_total_price,
		Inv_discount:         dataH.Or_discount,
		Inv_befor_vat:        dataH.Or_befor_vat,
		Tax_type_id:          dataH.Tax_type_id,
		Tax_rate:             dataH.Tax_rate,
		Inv_vat:              dataH.Or_vat,
		Inv_total:            dataH.Or_total,
		Inv_pay:              0,
		Inv_pay_total:        0,
		Inv_is_active:        dataH.Or_is_active,
		Inv_tele_code:        dataH.Or_tele_code,
		Inv_datetime:         dataH.Or_datetime,
		Inv_create:           time.Now().Format("2006-01-02 15:04:05"),
		Inv_update:           time.Now().Format("2006-01-02 15:04:05"),
		DpmId:                dpmId,
		Inv_discount_type_id: dataH.Or_discount_type_id,
		Inv_discount_item:    dataH.Or_discount_item,
		Inv_discount_value:   dataH.Or_discount_value,
		Inv_eclaim_id:        dataH.Or_eclaim_id,
		Inv_eclaim_rate:      dataH.Or_eclaim_rate,
		Inv_eclaim_over:      dataH.Or_eclaim_over,
		Inv_eclaim_total:     dataH.Or_eclaim_total,
	}

	if err = tx.Table("invoices").Create(&objH).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	//add subs
	var addList []structs.InvoiceDetail
	for _, sub := range *dataD {
		objSub := structs.InvoiceDetail{
			Invoice_id:            objH.Id,
			Course_id:             sub.Course_id,
			Checking_id:           sub.Checking_id,
			Product_id:            sub.Product_id,
			Product_store_id:      sub.Product_store_id,
			Product_unit_id:       sub.Product_unit_id,
			Coin_id:               sub.Coin_id,
			Room_id:               sub.Room_id,
			Queue_id:              sub.Queue_id,
			Order_detail_id:       sub.Id,
			Invd_type_id:          sub.Ord_type_id,
			Invd_code:             sub.Ord_code,
			Invd_name:             sub.Ord_name,
			Invd_qty:              sub.Ord_qty,
			Invd_set_qty:          sub.Ord_set_qty,
			Invd_limit_qty:        sub.Ord_limit_qty,
			Invd_rate:             sub.Ord_rate,
			Topical_id:            sub.Topical_id,
			Invd_topical:          sub.Ord_topical,
			Invd_direction:        sub.Ord_direction,
			Invd_unit:             sub.Ord_unit,
			Invd_cost:             sub.Ord_cost,
			Invd_price:            sub.Ord_price,
			Invd_discount:         sub.Ord_discount,
			Invd_amount:           sub.Ord_amount,
			Tax_type_id:           sub.Tax_type_id,
			Tax_rate:              sub.Tax_rate,
			Invd_vat:              sub.Ord_vat,
			Invd_total:            sub.Ord_total,
			Invd_is_set:           sub.Ord_is_set,
			Invd_id_set:           sub.Ord_id_set,
			Category_eclaim_id:    sub.Category_eclaim_id,
			Invd_is_active:        1,
			Invd_modify:           time.Now().Format("2006-01-02 15:04:05"),
			Invd_eclaim:           sub.Ord_eclaim,
			Invd_discount_type_id: sub.Ord_discount_type_id,
			Invd_discount_item:    sub.Ord_discount_item,
		}
		addList = append(addList, objSub)
	}

	if len(addList) > 0 {
		if err = tx.Table("invoice_details").CreateInBatches(&addList, 200).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	Orderupdate := structs.OrderInvoiceUpdate{
		Or_is_active: 2,
		Or_update:    time.Now().Format("2006-01-02 15:04:05"),
	}
	if err = tx.Table("orders").Where("orders.id = ?", dataH.Id).Model(&Orderupdate).Updates(&Orderupdate).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()
	return objH.Id, nil
}

func CancelInvoice(invoice_id int, userId int, data *structs.Invoice) (err error) {
	query := configs.DB1.Table("invoices")
	query = query.Where("id = ?", invoice_id)
	query = query.Model(&data)
	query = query.Updates(map[string]interface{}{"invoices.inv_is_active": 0, "invoices.user_id_cancel": userId, "invoices.inv_update": time.Now().Format("2006-01-02 15:04:05")})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetInvoiceDocNoData(ShopId int, data *structs.DocInvoice) (err error) {
	query := configs.DB1.Table("doc_settings")
	query = query.Select("doc_settings.shop_id, doc_settings.invoice_id_default, doc_settings.invoice_number_default, doc_settings.invoice_number_digit, doc_settings.invoice_type")
	query = query.Where("doc_settings.shop_id = ?", ShopId)
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func UpdateInvoiceDocno(ShopId int, data *structs.DocInvoice) (err error) {
	query := configs.DB1.Table("doc_settings")
	query = query.Where("doc_settings.shop_id = ?", ShopId)
	query = query.Model(&data)
	query = query.Updates(map[string]interface{}{"doc_settings.invoice_number_default": data.Invoice_number_default})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetInvoiceId(invoice_id int, dt *structs.Invoice) (err error) {
	query := configs.DB1.Table("invoices")
	query = query.Select("invoices.*")
	query = query.Where("invoices.id = ?", invoice_id)
	if err = query.Scan(&dt).Error; err != nil {
		return err
	}
	return nil
}

func GetInvoiceList(filter structs.ObjPayloadSearchInvoice, isCount bool, dt *[]structs.InvoiceList) error {
	// Define select fields based on isCount
	selectFields := "invoices.id"
	if isCount {
		selectFields = `
			invoices.*,
			customers.shop_id AS ctm_shop_id,
			shops.shop_name AS ctm_shop_name,
			queues.que_code,
			users.user_fullname,
			users.user_fullname_en,
			customers.ctm_id,
			customers.ctm_fname,
			customers.ctm_lname,
			customers.ctm_fname_en,
			customers.ctm_lname_en,
			queues.que_code,
			receipts.id AS receipt_id`
	}

	// Build base query with joins
	query := configs.DB1.Table("invoices").
		Select(selectFields).
		Joins("INNER JOIN customers ON invoices.customer_id = customers.id").
		Joins("INNER JOIN users ON invoices.user_id = users.id").
		Joins("LEFT JOIN queues ON invoices.queue_id = queues.id").
		Joins(`LEFT JOIN receipts ON receipts.invoice_id = invoices.id 
			   AND receipts.rec_is_active = 1 
			   AND receipts.rec_is_process = 1`).
		Joins("LEFT JOIN shops ON customers.shop_id = shops.id")

	// Add base shop filter
	query = query.Where("invoices.shop_id = ?", filter.Shop_id)

	// Add conditional filters
	if *filter.Inv_datetime != "" {
		query = query.Where("DATE(invoices.inv_datetime) = ?", *filter.Inv_datetime)
	}

	if *filter.Customer_id != "" {
		query = query.Where("invoices.customer_id = ?", *filter.Customer_id)
	}

	// Handle invoice active status
	if *filter.Inv_is_active != "" {
		activeStatus := 0
		switch *filter.Inv_is_active {
		case "1":
			activeStatus = 1
		case "2":
			activeStatus = 2
		}
		query = query.Where("invoices.inv_is_active = ?", activeStatus)
	}

	// Handle search with parameterized query
	if *filter.Search != "" {
		searchTerm := "%" + *filter.Search + "%"
		query = query.Where(
			`queues.que_code LIKE ? OR 
			 invoices.inv_code LIKE ? OR 
			 customers.ctm_id LIKE ? OR 
			 invoices.inv_fullname LIKE ? OR 
			 customers.ctm_fname LIKE ? OR 
			 customers.ctm_lname LIKE ? OR 
			 customers.ctm_fname_en LIKE ? OR 
			 customers.ctm_lname_en LIKE ?`,
			searchTerm, searchTerm, searchTerm, searchTerm,
			searchTerm, searchTerm, searchTerm, searchTerm,
		)
	}

	// Add pagination if counting
	if isCount {
		offset := filter.ActivePage * filter.PerPage
		query = query.Limit(filter.PerPage).Offset(offset)
	}

	// Add ordering and execute query
	return query.Order("invoices.inv_create DESC").Find(&dt).Error
}

func GetInvoiceDetail(invoice_id int, shopId int, dt *structs.InvoiceId) (err error) {
	query := configs.DB1.Table("invoices")
	query = query.Select("invoices.*, users.user_fullname, users.user_fullname_en, customers.ctm_fname, customers.ctm_lname, customers.ctm_fname_en, customers.ctm_lname_en, queues.que_code")
	query = query.Where("invoices.id = ?", invoice_id)
	query = query.Where("invoices.shop_id = ?", shopId)
	query = query.Joins("INNER JOIN customers ON invoices.customer_id = customers.id")
	query = query.Joins("INNER JOIN users ON invoices.user_id = users.id")
	query = query.Joins("LEFT JOIN queues ON invoices.queue_id = queues.id")

	if err = query.Scan(&dt).Error; err != nil {
		return err
	}
	return nil
}

func GetInvoiceSub(invoice_id int, dt *[]structs.InvoiceSub) (err error) {
	query := configs.DB1.Table("invoice_details")
	query = query.Select("invoice_details.*, rooms.room_code, rooms.room_th, rooms.room_en, room_types.room_type_th, room_types.room_type_en, courses.course_code, courses.course_name, courses.course_unit, checkings.checking_code, checkings.checking_name, checkings.checking_unit, products.pd_code, products.pd_name, ref_units.u_name, ref_units.u_name_en")
	query = query.Where("invoice_details.invd_is_active = 1")
	query = query.Where("invoice_details.invoice_id = ?", invoice_id)
	query = query.Joins("LEFT JOIN rooms ON invoice_details.room_id = rooms.id")
	query = query.Joins("LEFT JOIN room_types ON rooms.room_type_id = room_types.id")
	query = query.Joins("LEFT JOIN courses ON invoice_details.course_id = courses.id")
	query = query.Joins("LEFT JOIN checkings ON invoice_details.checking_id = checkings.id")
	query = query.Joins("LEFT JOIN products ON invoice_details.product_id = products.id")
	query = query.Joins("LEFT JOIN product_units ON invoice_details.product_unit_id = product_units.id")
	query = query.Joins("LEFT JOIN ref_units ON product_units.unit_id = ref_units.id")

	if err = query.Scan(&dt).Error; err != nil {
		return err
	}
	return nil
}

func UpdateInvoiceOrderCancal(order_id int) (err error) {
	query := configs.DB1.Table("orders")
	query = query.Where("orders.id = ?", order_id)
	query = query.Updates(map[string]interface{}{"orders.or_is_active": 1, "orders.or_update": time.Now().Format("2006-01-02 15:04:05")})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetCategoryEclaimCheck(id int, data *structs.CategoryEclaimId) (err error) {
	query := configs.DB1.Table("checkings")
	query = query.Select("category_eclaim_id")
	query = query.Where("checkings.id = ?", id)
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetCategoryEclaimCourse(id int, data *structs.CategoryEclaimId) (err error) {
	query := configs.DB1.Table("courses")
	query = query.Select("category_eclaim_id")
	query = query.Where("courses.id = ?", id)
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetCategoryEclaimProduct(product_id int, product_unit_id int, shop_id int, data *structs.CategoryEclaimId) (err error) {
	query := configs.DB1.Table("products")
	query = query.Select("product_shop_prices.category_eclaim_id")
	query = query.Joins("INNER JOIN product_units ON products.id = product_units.product_id")
	query = query.Joins("INNER JOIN product_shop_prices ON product_units.id = product_shop_prices.product_unit_id")
	query = query.Where("products.id = ?", product_id)
	query = query.Where("product_units.id = ?", product_unit_id)
	query = query.Where("product_shop_prices.shop_id = ?", shop_id)
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

// Add Sticker
func GetStickerInvoiceDetailProduct(invoice_id int, data *[]structs.StickerInvoiceDetail) (err error) {
	query := configs.DB1.Table("invoice_details")
	query = query.Select("invoice_details.*")
	query = query.Where("invoice_details.invoice_id = ?", invoice_id)
	query = query.Where("invoice_details.invd_is_active = ?", 1)
	query = query.Where("invoice_details.invd_type_id IN (3,5)")
	query = query.Order("invoice_details.invd_type_id,invoice_details.id")
	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func CreateProductStickerInvoice(objQuery *structs.ProcessProductStickerInvoice) (err error) {
	query := configs.DB1.Table("stickers").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateCancalStickerInvoice(invoice_id int) (err error) {
	query := configs.DB1.Table("stickers")
	query = query.Where("stickers.invoice_id = ?", invoice_id)
	query = query.Updates(map[string]interface{}{"stickers.sticker_is_del": 1, "stickers.sticker_modify": time.Now().Format("2006-01-02 15:04:05")})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetInvoiceDetailCheck(invoice_id int, data *[]structs.InvoiceDetail) (err error) {
	query := configs.DB1.Table("invoice_details")
	query = query.Select("invoice_details.*")
	query = query.Where("invoice_details.invoice_id = ?", invoice_id)
	query = query.Where("invoice_details.invd_type_id = ?", 1)
	query = query.Where("invoice_details.invd_is_active = ?", 1)
	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetCheckingChecksId(checkingId int, checking *structs.CheckChecking) (err error) {
	query := configs.DB1.Table("checkings")
	query = query.Select("checkings.*")
	query = query.Where("checkings.id = ?", checkingId)
	if err = query.Scan(&checking).Error; err != nil {
		return err
	}
	return nil
}

func CreateCheck(objQuery *structs.InvoiceCheck) (id int, err error) {
	query := configs.DB1.Table("checks").Create(&objQuery)
	if err = query.Error; err != nil {
		return 0, err
	}
	return objQuery.Id, nil
}

func UpdateCancalCheckInvoice(invoice_id int) (err error) {
	query := configs.DB1.Table("checks")
	query = query.Where("checks.invoice_id = ?", invoice_id)
	query = query.Updates(map[string]interface{}{"checks.chk_is_active": 0, "checks.chk_update": time.Now().Format("2006-01-02 15:04:05")})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateInvoiceDetails(invoice_detail_id int, check_id int) (err error) {
	query := configs.DB1.Table("invoice_details")
	query = query.Where("invoice_details.id = ?", invoice_detail_id)
	query = query.Updates(map[string]interface{}{"invoice_details.check_id": check_id, "invoice_details.invd_modify": time.Now().Format("2006-01-02 15:04:05")})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetQueueDateTime(queue_id int, data *structs.GetQueueDateTime) error {
	query := configs.DB1.Table("queues")
	query = query.Select("queues.que_datetime")
	query = query.Where("queues.id = ?", queue_id)
	query = query.Find(&data)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
