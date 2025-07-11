package models

import (
	"encoding/json"
	"fmt"
	"linecrmapi/configs"
	"linecrmapi/structs"
	"math"
	"time"
)

func AddLogReceipt(log *structs.LogReceipt) (err error) {
	query := configs.DBL1.Table("log_receipts").Create(&log)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetReceiptInvoiceId(invoiceId int, dt *structs.ReceiptInvoice) (err error) {
	query := configs.DB1.Table("invoices")
	query = query.Select("invoices.*")
	query = query.Where("invoices.id = ?", invoiceId)
	if err = query.Scan(&dt).Error; err != nil {
		return err
	}
	return nil
}

func GetInvoiceDetailId(invoiceId int, dt *[]structs.ReceiptInvoiceDetail) (err error) {
	query := configs.DB1.Table("invoice_details")
	query = query.Select("invoice_details.*")
	query = query.Where("invoice_details.invoice_id = ?", invoiceId)
	query = query.Where("invoice_details.invd_is_active = ?", 1)
	if err = query.Scan(&dt).Error; err != nil {
		return err
	}
	return nil
}

func GetReceiptList(filter structs.ObjPayloadSearchReceipt, isCount bool, dt *[]structs.ReceiptList) (err error) {
	query := configs.DB1.Table("receipts")
	if isCount == true {
		query = query.Select("receipts.*,customers.shop_id AS ctm_shop_id,shops.shop_name AS ctm_shop_name, receipts.rec_create AS rec_datetime, shops.shop_name, users.user_fullname,users.user_fullname_en, customers.ctm_id, customers.ctm_fname, customers.ctm_lname, customers.ctm_fname_en, customers.ctm_lname_en, queues.que_code, invoices.inv_code")
	} else {
		query = query.Select("receipts.id")
	}
	if *filter.Rec_pay_datetime != "" {
		query = query.Where("DATE(receipts.rec_pay_datetime) = ?", *filter.Rec_pay_datetime)
	}
	if *filter.Customer_id != "" {
		query = query.Where("receipts.customer_id = ?", filter.Customer_id)
	}
	query = query.Where("receipts.shop_id = ?", filter.Shop_id)
	if *filter.Rec_is_active != "" {
		if *filter.Rec_is_active == "1" {
			query = query.Where("receipts.rec_is_active = 1")
		} else if *filter.Rec_is_active == "2" {
			query = query.Where("receipts.rec_is_active = 2")
		} else {
			query = query.Where("receipts.rec_is_active = 0")
		}
	}

	if *filter.Search != "" {
		query = query.Where("queues.que_code LIKE '%" + *filter.Search + "%' OR customers.ctm_id LIKE '%" + *filter.Search + "%' OR receipts.rec_code LIKE '%" + *filter.Search + "%' OR invoices.inv_code LIKE '%" + *filter.Search + "%' OR receipts.rec_fullname LIKE '%" + *filter.Search + "%' OR customers.ctm_fname LIKE '%" + *filter.Search + "%' OR customers.ctm_lname LIKE '%" + *filter.Search + "%' OR customers.ctm_fname_en LIKE '%" + *filter.Search + "%' OR customers.ctm_lname_en LIKE '%" + *filter.Search + "%'")
	}

	query = query.Joins("INNER JOIN customers ON receipts.customer_id = customers.id")
	query = query.Joins("INNER JOIN users ON receipts.user_id = users.id")
	query = query.Joins("INNER JOIN invoices ON receipts.invoice_id = invoices.id")
	query = query.Joins("LEFT JOIN queues ON receipts.queue_id = queues.id")
	// query = query.Joins("INNER JOIN shops ON receipts.shop_id = shops.id")
	query = query.Joins("LEFT JOIN shops ON customers.shop_id = shops.id")
	if isCount == true {
		offset := filter.ActivePage * filter.PerPage
		query = query.Limit(filter.PerPage)
		query = query.Offset(offset)
	}

	query = query.Order("receipts.rec_code DESC")
	query = query.Order("receipts.rec_create DESC")
	query = query.Order("receipts.rec_is_active DESC")
	if err = query.Find(&dt).Error; err != nil {
		return err
	}
	return nil
}

func GetReceipById(rcId int, shopId int, data *structs.ReceiptPrint, customerId int) (err error) {
	query := configs.DB1.Table("receipts")
	query = query.Select("receipts.*, account_lists.acl_code, account_lists.acl_name, receipts.rec_create AS rec_datetime, users.user_fullname as rec_user_fullname, users.user_fullname_en as rec_user_fullname_en, customers.ctm_fname, customers.ctm_lname, customers.ctm_fname_en, customers.ctm_lname_en, queues.que_code, invoices.inv_code")
	query = query.Joins("INNER JOIN customers ON receipts.customer_id = customers.id")
	query = query.Joins("INNER JOIN users ON receipts.user_id = users.id")
	query = query.Joins("INNER JOIN invoices ON receipts.invoice_id = invoices.id")
	query = query.Joins("INNER JOIN doc_settings ON receipts.shop_id = doc_settings.shop_id")
	query = query.Joins("LEFT JOIN queues ON receipts.queue_id = queues.id")
	query = query.Joins("LEFT JOIN account_lists ON receipts.account_list_id = account_lists.id")
	query = query.Where("receipts.id", rcId)
	query = query.Where("receipts.shop_id", shopId)
	query = query.Where("receipts.customer_id", customerId)
	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetUserCancel(userId int, dt *structs.UserCancel) (err error) {
	query := configs.DB1.Table("users")
	query = query.Select("users.user_fullname, users.user_fullname_en")
	query = query.Where("users.id = ?", userId)
	if err = query.Scan(&dt).Error; err != nil {
		return err
	}
	return nil
}

func GetReceiptDetail(rcId int, data *[]structs.ReceiptDetail) (err error) {
	query := configs.DB1.Table("receipt_details")
	query = query.Select("receipt_details.*, ref_units.u_name_en, services.ser_exp_date")
	query = query.Joins("LEFT JOIN product_units ON receipt_details.product_unit_id = product_units.id")
	query = query.Joins("LEFT JOIN ref_units ON product_units.unit_id = ref_units.id")
	query = query.Joins("LEFT JOIN services ON services.course_id = receipt_details.course_id")
	query = query.Where("receipt_details.receipt_id", rcId)
	query = query.Group("receipt_details.id")

	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetReceiptDocNoData(ShopId int, data *structs.DocReceipt) (err error) {
	query := configs.DB1.Table("doc_settings")
	query = query.Select("doc_settings.shop_id, doc_settings.receipt_id_default, doc_settings.receipt_number_default, doc_settings.receipt_number_digit, doc_settings.receipt_type")
	query = query.Where("doc_settings.shop_id = ?", ShopId)
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetReceiptId(receipt_id int, dt *structs.Receipt) (err error) {
	query := configs.DB1.Table("receipts")
	query = query.Select("receipts.*")
	query = query.Where("receipts.id = ?", receipt_id)
	if err = query.Scan(&dt).Error; err != nil {
		return err
	}
	return nil
}

func GetReceiptFileCheckBYId(receipt_id int, dt *structs.ReceiptFileCheck) (err error) {
	query := configs.DB1.Table("receipts")
	query = query.Select("receipts.*, shops.shop_code, customers.ctm_id")
	query = query.Joins("INNER JOIN customers ON receipts.customer_id = customers.id")
	query = query.Joins("INNER JOIN shops ON receipts.shop_id = shops.id")
	query = query.Where("receipts.id = ?", receipt_id)
	if err = query.Scan(&dt).Error; err != nil {
		return err
	}
	return nil
}

func CancelReceipt(receipt_id int, userId int, data *structs.Receipt) (err error) {
	query := configs.DB1.Table("receipts")
	query = query.Where("id = ?", receipt_id)
	query = query.Model(&data)
	query = query.Updates(map[string]interface{}{"receipts.rec_is_active": 0, "receipts.user_id_cancel": userId, "receipts.rec_update": time.Now().Format("2006-01-02 15:04:05")})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateFileReceipt(receipt_id int, size int, file string, data *structs.Receipt) (err error) {
	query := configs.DB1.Table("receipts")
	query = query.Where("id = ?", receipt_id)
	query = query.Model(&data)
	query = query.Updates(map[string]interface{}{"receipts.rec_file": file, "receipts.rec_file_size": size, "receipts.rec_update": time.Now().Format("2006-01-02 15:04:05")})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateReceiptInvoiceCancal(invoice_id int, data *structs.CancelInvoiceUpdate) (err error) {
	query := configs.DB1.Table("invoices")
	query = query.Where("invoices.id = ?", invoice_id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateReceiptQueueCancal(queue_id int) (err error) {
	query := configs.DB1.Table("queues")
	query = query.Where("queues.id = ?", queue_id)
	query = query.Updates(map[string]interface{}{"queues.que_status_id": 3, "queues.que_update": time.Now().Format("2006-01-02 15:04:05")})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetReceiptDetailCheck(receipt_id int, data *[]structs.ReceiptDetail) (err error) {
	query := configs.DB1.Table("receipt_details")
	query = query.Select("receipt_details.*")
	query = query.Where("receipt_details.receipt_id = ?", receipt_id)
	query = query.Where("receipt_details.recd_type_id = ?", 1)
	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetReceiptDetailCheckProduct(receipt_id int, checking_id int, data *[]structs.ReceiptDetail) (err error) {
	query := configs.DB1.Table("receipt_details")
	query = query.Select("receipt_details.*")
	query = query.Where("receipt_details.receipt_id = ?", receipt_id)
	query = query.Where("receipt_details.checking_id = ?", checking_id)
	query = query.Where("receipt_details.recd_type_id = ?", 5)
	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetReceiptDetailCourse(receipt_id int, data *[]structs.ReceiptDetail) (err error) {
	query := configs.DB1.Table("receipt_details")
	query = query.Select("receipt_details.*")
	query = query.Where("receipt_details.receipt_id = ?", receipt_id)
	query = query.Where("receipt_details.recd_type_id = ?", 2)
	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetReceiptDetailCourseProduct(receipt_id int, course_id int, data *[]structs.ReceiptDetail) (err error) {
	query := configs.DB1.Table("receipt_details")
	query = query.Select("receipt_details.*")
	query = query.Where("receipt_details.receipt_id = ?", receipt_id)
	query = query.Where("receipt_details.course_id = ?", course_id)
	query = query.Where("receipt_details.recd_type_id = ?", 5)
	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func UpdateReceiptDocno(ShopId int, data *structs.DocReceipt) (err error) {
	query := configs.DB1.Table("doc_settings")
	query = query.Where("doc_settings.shop_id = ?", ShopId)
	query = query.Model(&data)
	query = query.Updates(map[string]interface{}{"doc_settings.receipt_number_default": data.Receipt_number_default})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetShopReceiptById(shopId int, objResponse *structs.ReceiptShop) error {
	query := configs.DB1.Table("shops")
	query = query.Select("shops.*,currency.currency_symbol, doc_settings.show_product_id, doc_settings.show_product_th, doc_settings.show_product_en, doc_settings.show_course_check_id, doc_settings.show_course_check_th, doc_settings.show_course_check_en,doc_settings.show_date_id, doc_settings.show_page_id, doc_settings.print_th, doc_settings.print_en, doc_settings.print_la, doc_settings.print_a4, doc_settings.print_ca, doc_settings.print_a5, doc_settings.print_80, doc_settings.invoice_comment_id, doc_settings.receipt_comment_id, doc_settings.tax_comment_id, doc_settings.purchase_comment_id, doc_settings.transfer_comment_id, doc_settings.invoice_comment, doc_settings.receipt_comment, doc_settings.tax_comment, doc_settings.purchase_comment, doc_settings.transfer_comment, doc_settings.invoice_copy, doc_settings.receipt_copy, doc_settings.tax_copy, doc_settings.purchase_copy, doc_settings.transfer_copy, doc_settings.sticker_font_size, doc_settings.sticker_width, doc_settings.sticker_height, doc_settings.sticker_show_name, doc_settings.sticker_show_address, doc_settings.sticker_show_tel, doc_settings.sticker_show_date, doc_settings.sticker_show_expdate, doc_settings.sticker_show_detail")
	query = query.Joins("INNER JOIN doc_settings ON doc_settings.shop_id = shops.id")
	query = query.Joins("INNER JOIN currency ON currency.id = shops.currency_id")
	query = query.Where("shops.id = ?", shopId)
	query = query.Find(&objResponse)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func UpdateReceipt() {

	// Invoiceupdate := structs.InvoiceReceiptUpdate{
	// 	Inv_pay:       payload.Rec_pay,
	// 	Inv_is_active: 2,
	// 	Inv_update:    time.Now().Format("2006-01-02 15:04:05"),
	// }
	// if err = tx.Table("invoices").Where("invoices.id = ?", dataH.Id).Model(&Invoiceupdate).Updates(&Invoiceupdate).Error; err != nil {
	// 	tx.Rollback()
	// 	return 0, err
	// }

	// Queueupdate := structs.ReceiptQueueUpdate{
	// 	Que_status_id: 4,
	// 	Que_update:    time.Now().Format("2006-01-02 15:04:05"),
	// }

	// if *dataH.Queue_id > 0 {
	// 	if err = tx.Table("queues").Where("queues.id = ?", *dataH.Queue_id).Model(&Queueupdate).Updates(&Queueupdate).Error; err != nil {
	// 		tx.Rollback()
	// 		return 0, err
	// 	}
	// }
}

// tom code
func GetInvoiceById(invId int, objQuery *Invoice) error {
	query := configs.DB1.Table("invoices")
	query = query.Select("invoices.*")
	query = query.Where("invoices.id = ?", invId)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetInvoiceDetailById(invId int, objQuery *[]InvoiceDetail) error {
	query := configs.DB1.Table("invoice_details")
	query = query.Select("invoice_details.*")
	query = query.Where("invoice_details.invoice_id = ?", invId)
	query = query.Where("invoice_details.invd_is_active = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetReceiptDocSetting(shopId int, objQuery *structs.ObjQueryReceiptDocSetting) error {
	query := configs.DB1.Table("doc_settings")
	query = query.Select("receipt_id_default, receipt_number_default, receipt_number_digit, receipt_type")
	query = query.Where("shop_id = ?", shopId)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetReceiptLast(invId int, objQuery *Receipt) error {
	query := configs.DB1.Table("receipts")
	query = query.Select("receipts.*")
	query = query.Where("receipts.invoice_id = ?", invId)
	query = query.Where("receipts.rec_is_active != ?", 0)
	query = query.Order("receipts.rec_create DESC")
	query = query.Limit(1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetReceiptHistory(invId int, objQuery *[]Receipt) error {
	query := configs.DB1.Table("receipts")
	query = query.Select("receipts.*")
	query = query.Where("receipts.invoice_id = ?", invId)
	query = query.Where("receipts.rec_is_active != ?", 0)
	query = query.Order("receipts.rec_create ASC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetReceiptShopById(shopId int, objQuery *Shop) error {
	query := configs.DB1.Table("shops")
	query = query.Select("shops.*")
	query = query.Where("shops.id = ?", shopId)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetReceiptCustomerById(ctmId int, objQuery *Customer) error {
	query := configs.DB1.Table("customers")
	query = query.Select("customers.*")
	query = query.Where("customers.id = ?", ctmId)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func UpdateReceiptDocSetting(shopId int, nextNumberDefault int, data *DocSetting) error {
	query := configs.DB1.Table("doc_settings")
	query = query.Where("shop_id = ?", shopId)
	query = query.Model(&data)
	query = query.Updates(map[string]interface{}{"receipt_number_default": nextNumberDefault})
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CreateReceiptTx(objInvoice *Invoice, objInvoiceDetail *[]InvoiceDetail, objPayload *structs.ObjPayloadAddReceipt, objShop *Shop, objCustomer *Customer, recCode string, recIsProcess int, recPeriod int, recCodeNext int, invIsActive int, invPayTotal float64, invIsGetPoint int, invGetPointAmount int) (int, error) {

	tx := configs.DB1.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return 0, err
	}

	// add receipt
	var RecBalance float64 = 0
	var RecPayTotal float64 = 0
	var RecPointGive int = 0
	var RecPointUsed int = 0
	RecPayTotal = invPayTotal
	RecBalance = objInvoice.InvTotal - RecPayTotal
	if invIsGetPoint == 1 {
		RecPointGive = invGetPointAmount
	}
	if objPayload.RecPaymentType == 6 {
		RecPointUsed = int(math.Ceil(objPayload.RecPay) * float64(objShop.ShopPointUseRate))
	}
	objAddReceipt := Receipt{
		ShopId:             objInvoice.ShopId,
		UserId:             objPayload.RecUserId,
		CustomerId:         objInvoice.CustomerId,
		CustomerOnlineId:   objInvoice.CustomerOnlineId,
		QueueId:            objInvoice.QueueId,
		InvoiceId:          objInvoice.ID,
		AccountListId:      objPayload.AccountListId,
		RecCode:            recCode,
		RecFullname:        objInvoice.InvFullname,
		RecTel:             objInvoice.InvTel,
		RecEmail:           objInvoice.InvEmail,
		RecAddress:         objInvoice.InvAddress,
		RecDiscountTypeId:  objInvoice.Inv_discount_type_id,
		Rec_discount_item:  &objInvoice.Inv_discount_item,
		Rec_discount_value: &objInvoice.Inv_discount_value,
		RecDistrict:        objInvoice.InvDistrict,
		RecAmphoe:          objInvoice.InvAmphoe,
		RecProvince:        objInvoice.InvProvince,
		RecZipcode:         objInvoice.InvZipcode,
		RecComment:         objInvoice.InvComment,
		RecTotalPrice:      objInvoice.InvTotalPrice,
		RecDiscount:        objInvoice.InvDiscount,
		RecBeforVat:        objInvoice.InvBeforVat,
		TaxTypeId:          objInvoice.TaxTypeId,
		TaxRate:            objInvoice.TaxRate,
		RecVat:             objInvoice.InvVat,
		RecTotal:           objInvoice.InvTotal,
		RecPaymentType:     objPayload.RecPaymentType,
		RecTypeId:          objPayload.RecTypeId,
		RecPeriod:          recPeriod,
		RecPay:             objPayload.RecPay,
		RecBalance:         RecBalance,
		RecPayTotal:        RecPayTotal,
		RecPayDatetime:     objPayload.RecPayDatetime,
		RecDescription:     objPayload.RecDescription,
		RecAccount:         objPayload.RecAccount,
		RecUserId:          objPayload.RecUserId,
		RecUserFullname:    objPayload.RecUserFullname,
		RecPointGive:       RecPointGive,
		RecIsProcess:       recIsProcess,
		RecPointUsed:       RecPointUsed,
		RecIsActive:        1,
		RecCreate:          time.Now().Format("2006-01-02 15:04:05"),
		RecUpdate:          time.Now().Format("2006-01-02 15:04:05"),
		DpmId:              objPayload.DpmId,
		Rec_eclaim_id:      objInvoice.Inv_eclaim_id,
		Rec_eclaim_rate:    objInvoice.Inv_eclaim_rate,
		Rec_eclaim_over:    objInvoice.Inv_eclaim_over,
		Rec_eclaim_total:   objInvoice.Inv_eclaim_total,
	}

	if err := tx.Table("receipts").Create(&objAddReceipt).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return 0, err
	}

	// add receipt detail
	if len(*objInvoiceDetail) > 0 {
		objAddReceiptDetail := []ReceiptDetail{}
		for _, item := range *objInvoiceDetail {
			objAddReceiptDetail = append(objAddReceiptDetail, ReceiptDetail{
				ReceiptId:             objAddReceipt.ID,
				CourseId:              item.CourseId,
				CheckingId:            item.CheckingId,
				ProductId:             item.ProductId,
				ProductStoreId:        item.ProductStoreId,
				ProductUnitId:         item.ProductUnitId,
				CoinId:                item.CoinId,
				RoomId:                item.RoomId,
				QueueId:               item.QueueId,
				InvoiceDetailId:       item.ID,
				RecdTypeId:            item.InvdTypeId,
				RecdCode:              item.InvdCode,
				RecdName:              item.InvdName,
				RecdQty:               item.InvdQty,
				RecdSetQty:            item.InvdSetQty,
				RecdLimitQty:          item.InvdLimitQty,
				RecdRate:              item.InvdRate,
				TopicalId:             item.TopicalId,
				RecdTopical:           item.InvdTopical,
				RecdDirection:         item.InvdDirection,
				RecdUnit:              item.InvdUnit,
				RecdCost:              item.InvdCost,
				RecdPrice:             item.InvdPrice,
				RecdDiscount:          item.InvdDiscount,
				RecdAmount:            item.InvdAmount,
				TaxTypeId:             item.TaxTypeId,
				TaxRate:               item.TaxRate,
				RecdVat:               item.InvdVat,
				RecdTotal:             item.InvdTotal,
				RecdIsActive:          1,
				RecdModify:            time.Now().Format("2006-01-02 15:04:05"),
				Recd_eclaim:           item.Invd_eclaim,
				Recd_discount_type_id: item.Invd_discount_type_id,
				Recd_discount_item:    item.Invd_discount_item,
			})
		}

		objLen := len(objAddReceiptDetail)
		if objLen > 24 {
			objLen = 24
		}

		if err := tx.Table("receipt_details").CreateInBatches(&objAddReceiptDetail, objLen).Error; err != nil {
			fmt.Println("error = " + err.Error())
			tx.Rollback()
			return 0, err
		}
	}

	// update invoice
	var objInvoice_ Invoice
	if err := tx.Table("invoices").Model(&objInvoice_).Where("id = ?", objInvoice.ID).Updates(map[string]interface{}{"inv_pay_total": invPayTotal}).Error; err != nil {
		fmt.Println("error inv1 = " + err.Error())
		tx.Rollback()
		return 0, err
	}

	if objPayload.RecTypeId == 2 {
		var objInvoice__ Invoice
		if err := tx.Table("invoices").Model(&objInvoice__).Where("id = ?", objInvoice.ID).Updates(map[string]interface{}{"inv_deposit": objPayload.RecPay}).Error; err != nil {
			fmt.Println("error inv2 = " + err.Error())
			tx.Rollback()
			return 0, err
		}
	}

	// update doc setting
	var objDocSetting DocSetting
	if err := tx.Table("doc_settings").Model(&objDocSetting).Where("shop_id = ?", objInvoice.ShopId).Updates(map[string]interface{}{"receipt_number_default": recCodeNext}).Error; err != nil {
		fmt.Println("error doc_settings = " + err.Error())
		tx.Rollback()
		return 0, err
	}

	if objPayload.RecPaymentType == 5 {
		ctmCoin := objCustomer.CtmCoin - objPayload.RecPay
		var objCustomer_ Customer
		if err := tx.Table("customers").Model(&objCustomer_).Where("id = ?", objInvoice.CustomerId).Updates(map[string]interface{}{"ctm_coin": ctmCoin}).Error; err != nil {
			fmt.Println("error customers = " + err.Error())
			tx.Rollback()
			return 0, err
		}
		// Add coin history
		AddCoinHistory := structs.AddCoinHistory{
			Shop_id:     objInvoice.ShopId,
			Customer_id: objInvoice.CustomerId,
			Receipt_id:  objAddReceipt.ID,
			Rec_code:    recCode,
			Ch_forward:  float64(objCustomer.CtmCoin),
			Ch_amount:   float64(objPayload.RecPay * (-1)),
			Ch_total:    float64(ctmCoin),
			Ch_comment:  "Add Receipt: Used Coin History",
			Ch_create:   time.Now().Format("2006-01-02 15:04:05"),
		}
		if err := tx.Table("coin_historys").Create(&AddCoinHistory).Error; err != nil {
			fmt.Println("error coin_historys = " + err.Error())
			tx.Rollback()
			return 0, err
		}
	}

	if objPayload.RecPaymentType == 6 {
		ctmPoint := int(float64(objCustomer.CtmPoint) - (math.Ceil(objPayload.RecPay) * float64(objShop.ShopPointUseRate)))
		var objCustomer__ Customer
		if err := tx.Table("customers").Model(&objCustomer__).Where("id = ?", objInvoice.CustomerId).Updates(map[string]interface{}{"ctm_point": ctmPoint}).Error; err != nil {
			fmt.Println("error customers2 = " + err.Error())
			tx.Rollback()
			return 0, err
		}
		// Add point history
		AddPointHistory := structs.AddPointHistory{
			Shop_id:     objInvoice.ShopId,
			Customer_id: objInvoice.CustomerId,
			Receipt_id:  objAddReceipt.ID,
			Rec_code:    recCode,
			Ph_forward:  float64(objCustomer.CtmPoint),
			Ph_amount:   (math.Ceil(float64(objPayload.RecPay)) * float64(objShop.ShopPointUseRate)) * float64(-1),
			Ph_total:    float64(ctmPoint),
			Ph_comment:  "Add Receipt: Used Point History",
			Ph_create:   time.Now().Format("2006-01-02 15:04:05"),
		}
		if err := tx.Table("point_historys").Create(&AddPointHistory).Error; err != nil {
			fmt.Println("error point_historys = " + err.Error())
			tx.Rollback()
			return 0, err
		}
	}

	if invIsActive == 2 {
		var objInvoice__ Invoice
		if err := tx.Table("invoices").Model(&objInvoice__).Where("id = ?", objInvoice.ID).Updates(map[string]interface{}{"inv_is_active": invIsActive}).Error; err != nil {
			fmt.Println("error invoices2 = " + err.Error())
			tx.Rollback()
			return 0, err
		}
		if invIsGetPoint == 1 {
			ctmPoint := objCustomer.CtmPoint + invGetPointAmount
			var objCustomer___ Customer
			if err := tx.Table("customers").Model(&objCustomer___).Where("id = ?", objInvoice.CustomerId).Updates(map[string]interface{}{"ctm_point": ctmPoint}).Error; err != nil {
				fmt.Println("error customers3 = " + err.Error())
				tx.Rollback()
				return 0, err
			}
			if err := tx.Table("invoices").Model(&objInvoice__).Where("id = ?", objInvoice.ID).Updates(map[string]interface{}{"inv_point_give": invGetPointAmount}).Error; err != nil {
				fmt.Println("error invoices3 = " + err.Error())
				tx.Rollback()
				return 0, err
			}
			// Add point history
			AddPointHistory := structs.AddPointHistory{
				Shop_id:     objInvoice.ShopId,
				Customer_id: objInvoice.CustomerId,
				Receipt_id:  objAddReceipt.ID,
				Rec_code:    recCode,
				Ph_forward:  float64(objCustomer.CtmPoint),
				Ph_amount:   math.Ceil(float64(invGetPointAmount)),
				Ph_total:    float64(ctmPoint),
				Ph_comment:  "Add Receipt: Add Point History",
				Ph_create:   time.Now().Format("2006-01-02 15:04:05"),
			}
			if err := tx.Table("point_historys").Create(&AddPointHistory).Error; err != nil {
				fmt.Println("error point_historys2 = " + err.Error())
				tx.Rollback()
				return 0, err
			}
		}
	}

	tx.Commit()

	return objAddReceipt.ID, nil
}

func UpdatePointCustomerId(customer_id int, data *structs.PointCustomerUpdate) (err error) {
	query := configs.DB1.Table("customers")
	query = query.Where("customers.id = ?", customer_id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateCoinCustomerId(customer_id int, data *structs.CoinCustomerUpdate) (err error) {
	query := configs.DB1.Table("customers")
	query = query.Where("customers.id = ?", customer_id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CheckQueueReceiptServiceUsed(receipt_id int, queue_id int, data *[]structs.ServiceUsed) (err error) {
	query := configs.DB1.Table("service_useds")
	query = query.Select("service_useds.*")
	query = query.Where("service_useds.receipt_id = ?", receipt_id)
	query = query.Where("service_useds.queue_id != ?", queue_id)
	query = query.Where("service_useds.seru_is_active = ?", 1)
	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func CheckReceiptServiceUsed(receipt_id int, data *[]structs.ServiceUsed) (err error) {
	query := configs.DB1.Table("service_useds")
	query = query.Select("service_useds.*")
	query = query.Where("service_useds.receipt_id = ?", receipt_id)
	query = query.Where("service_useds.seru_is_active = ?", 1)
	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetInvoiceDetailProduct(invoice_id int, data *[]structs.InvoiceDetail) (err error) {
	query := configs.DB1.Table("invoice_details")
	query = query.Select("invoice_details.*")
	query = query.Where("invoice_details.invoice_id = ?", invoice_id)
	query = query.Where("invoice_details.invd_is_active = ?", 1)
	query = query.Where("(invoice_details.invd_type_id = 3 || invoice_details.invd_type_id = 5)")
	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func CheckMaxReceipt(invoice_id int, data *structs.MaxReceipt) (err error) {
	query := configs.DB1.Table("receipts")
	query = query.Select("MAX(receipts.rec_period) AS rec_period_max")
	query = query.Where("receipts.invoice_id = ?", invoice_id)
	query = query.Where("receipts.rec_is_active != ?", 0)
	query = query.Group("receipts.invoice_id")
	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetInvoiceDetailProductCheck(invoice_id int, data *[]structs.InvoiceDetailCheck) error {
	return configs.DB1.Table("invoice_details").
		Select("invoice_details.id, invoice_details.product_id, invoice_details.product_store_id, invoice_details.invd_code, invoice_details.invd_name, SUM(invoice_details.invd_qty) AS invd_qty").
		Where("invoice_details.invoice_id = ? AND invoice_details.invd_is_active = ?",
			invoice_id, 1).
		Where("(invoice_details.invd_type_id = 3 OR invoice_details.invd_type_id = 5)").
		Group("invoice_details.product_store_id").
		Find(data).Error
}

func GetProcessProductStoreIdCheck(product_store_id int, product *structs.ProcessProductStore) (err error) {
	query := configs.DB1.Table("product_stores")
	query = query.Select("product_stores.*")
	query = query.Where("product_stores.id = ?", product_store_id)
	if err = query.Scan(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProcessProductStoreOrderCheck(product_id int, product_stores_id int, product *structs.ProcessProductStoreOrderCheck) (err error) {
	query := configs.DB1.Table("product_store_orders")
	query = query.Select("product_store_orders.id,SUM(product_store_orders.pdso_total) AS pdso_total")
	query = query.Joins("INNER JOIN product_stores ON product_store_orders.product_store_id = product_stores.id")
	query = query.Where("product_stores.product_id = ?", product_id)
	query = query.Where("product_store_orders.product_store_id = ?", product_stores_id)
	query = query.Where("product_store_orders.pdso_is_active = ?", 1)
	query = query.Where("product_store_orders.pdso_expire > ?", time.Now().Format("2006-01-02")) // ไม่หมดอายุ
	if err = query.Scan(product).Error; err != nil {
		return err
	}
	return nil
}

func GetInvoiceDetailProductCheckCourse(invoice_id int, product_store_id int, data *structs.InvoiceDetailCheckCourse) (err error) {
	query := configs.DB1.Table("invoice_details")
	query = query.Select("invoice_details.id,SUM( invoice_details.invd_qty ) AS invd_qty ")
	query = query.Joins("INNER JOIN invoices ON invoice_details.invoice_id = invoices.id ")
	query = query.Where("invoice_details.invoice_id = ?", invoice_id)
	query = query.Where("invoice_details.product_store_id = ?", product_store_id)
	query = query.Where("invoice_details.invd_is_active = ?", 1)
	query = query.Where("invoice_details.course_id IS NOT NULL")
	query = query.Where("invoice_details.queue_id IS NULL")
	if err = query.Scan(data).Error; err != nil {
		return err
	}
	return nil
}
