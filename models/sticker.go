package models

import (
	"encoding/json"
	"linecrmapi/configs"
	"linecrmapi/structs"
)

func GetStickerList(filter structs.ObjPayloadSearchSticker, isCount bool, dt *[]structs.Stickers) (err error) {
	query := configs.DB1.Table("stickers")
	if *filter.StkActiveId == 1 {
		query = query.Joins("JOIN invoices ON stickers.invoice_id = invoices.id")
	}
	// else if *filter.StkActiveId == 2 {
	// 	query = query.Joins("LEFT JOIN invoices ON stickers.invoice_id = invoices.id")
	// }

	query = query.
		Joins("JOIN customers ON stickers.customer_id = customers.id").
		Joins("JOIN users ON stickers.user_id = users.id").
		Where("stickers.sticker_is_del = 0").
		// Where("invoices.shop_id = ? OR stickers.shop_id = ?", filter.Shop_id, filter.Shop_id)
		Where("stickers.shop_id = ?", filter.Shop_id)

	if isCount {
		if *filter.StkActiveId == 2 {
			query = query.Select("stickers.*, '' AS inv_code, customers.ctm_id, customers.ctm_fname, customers.ctm_lname, customers.ctm_fname_en, customers.ctm_lname_en, customers.ctm_prefix, users.user_fullname, users.user_fullname_en")
		} else {
			query = query.Select("stickers.*, invoices.inv_code, customers.ctm_id, customers.ctm_fname, customers.ctm_lname, customers.ctm_fname_en, customers.ctm_lname_en, customers.ctm_prefix, users.user_fullname, users.user_fullname_en")
		}
	} else {
		query = query.Select("stickers.shop_id")
	}

	if searchText := *filter.Search_text; searchText != "" {
		if *filter.StkActiveId == 2 {
			query = query.Where(
				"customers.ctm_id LIKE ? OR CONCAT(customers.ctm_fname_en,' ',customers.ctm_lname_en) LIKE ? OR CONCAT(customers.ctm_fname,' ',customers.ctm_lname) LIKE ?",
				"%"+searchText+"%", "%"+searchText+"%", "%"+searchText+"%",
			)
		} else {
			query = query.Where(
				"invoices.inv_code LIKE ? OR invoices.inv_fullname LIKE ? OR customers.ctm_id LIKE ? OR CONCAT(customers.ctm_fname_en,' ',customers.ctm_lname_en) LIKE ? OR CONCAT(customers.ctm_fname,' ',customers.ctm_lname) LIKE ?",
				"%"+searchText+"%", "%"+searchText+"%", "%"+searchText+"%", "%"+searchText+"%", "%"+searchText+"%",
			)
		}
	}

	if *filter.StkActiveId == 2 {
		query = query.Where("stickers.sticker_active_id = ?", *filter.StkActiveId)
		query = query.Where("stickers.invoice_id IS NULL")
	} else {
		query = query.Where("stickers.sticker_active_id = ?", *filter.StkActiveId)
	}

	if *filter.User_id != -1 {
		query = query.Where("stickers.user_id = ?", *filter.User_id)
	}

	if searchDate := *filter.Search_date; searchDate != "" {
		query = query.Where("stickers.sticker_modify LIKE ?", "%"+searchDate+"%")
	}

	if *filter.StkActiveId == 2 {
		query = query.Order("stickers.id DESC")
	} else {
		query = query.Group("stickers.invoice_id").Order("stickers.sticker_modify DESC")
	}

	if isCount {
		offset := filter.ActivePage * filter.PerPage
		query = query.Limit(filter.PerPage).Offset(offset)
	}

	return query.Scan(&dt).Error

}

func GetStickerReceipt(invoice_id int, data *[]structs.StickerProduct) (err error) {
	query := configs.DB1.Table("stickers")
	query = query.Where("stickers.invoice_id = ?", invoice_id)
	query = query.Where("stickers.sticker_is_del = ?", 0)
	query = query.Where("stickers.sticker_active_id = ?", 1)
	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetStickerReceiptIdInArray(invoice_id []int, data *[]structs.StickerProduct) (err error) {
	query := configs.DB1.Table("stickers")
	query = query.Select("stickers.*,products.pd_code_acc,products.pd_name_acc")
	query = query.Joins("JOIN products ON products.id = stickers.product_id")
	query = query.Where("stickers.invoice_id IN ?", invoice_id)
	query = query.Where("stickers.sticker_is_del = ?", 0)
	query = query.Where("stickers.sticker_active_id = ?", 1)
	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetReceipInPrescriptionById(InvId int, data *structs.Prescription) (err error) {
	query := configs.DB1.Table("invoices")
	// query = query.Select("invoices.*, account_lists.acl_code, account_lists.acl_name, invoices.inv_create AS inv_datetime, customers.ctm_fname, customers.ctm_lname, customers.ctm_fname_en, customers.ctm_lname_en, queues.que_code, invoices.inv_code, invoices.queue_id, queues.que_user_fullname, users.user_fullname")
	query = query.Select("invoices.*, invoices.inv_create AS inv_datetime, customers.ctm_fname, customers.ctm_lname, customers.ctm_fname_en, customers.ctm_lname_en, queues.que_code, invoices.inv_code, invoices.queue_id, queues.que_user_fullname, users.user_fullname")
	query = query.Where("invoices.id", InvId)
	query = query.Joins("INNER JOIN customers ON invoices.customer_id = customers.id")
	// query = query.Joins("INNER JOIN invoices ON invoices.invoice_id = invoices.id")
	query = query.Joins("INNER JOIN doc_settings ON invoices.shop_id = doc_settings.shop_id")
	// query = query.Joins("LEFT JOIN account_lists ON invoices.account_list_id = account_lists.id")
	query = query.Joins("LEFT JOIN queues ON invoices.queue_id = queues.id")
	query = query.Joins("LEFT JOIN users ON queues.user_id = users.id")
	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetStickerInvoiceList(invoice_id int, data *[]structs.StickerDetail) (err error) {
	query := configs.DB1.Table("stickers")
	query = query.Select("stickers.*,invoices.inv_code, customers.ctm_id,customers.ctm_birthdate, customers.ctm_fname, customers.ctm_lname, customers.ctm_fname_en, customers.ctm_lname_en, customers.ctm_prefix, customers.ctm_birthdate, users.user_fullname, doc_settings.sticker_font_size, doc_settings.sticker_width, shops.shop_name, shops.shop_name_en, shops.shop_phone, doc_settings.sticker_height, doc_settings.sticker_show_name, doc_settings.sticker_show_address, doc_settings.sticker_show_tel, doc_settings.sticker_show_date, doc_settings.sticker_show_expdate, doc_settings.sticker_show_detail, invoice_details.invd_discount, invoice_details.invd_amount, invoice_details.invd_total, invoice_details.invd_vat, shops.shop_address, shops.shop_address_en, shops.shop_district, shops.shop_district_en, shops.shop_amphoe, shops.shop_amphoe_en, shops.shop_province, shops.shop_province_en, shops.shop_zipcode,	shops.shop_zipcode_en, shops.shop_license, shops.shop_tax,products.pd_code_acc,products.pd_name_acc,doc_settings.show_product_acc,shops.shop_image,doc_settings.show_shop_image,doc_settings.show_shop_image")
	query = query.Joins("JOIN invoices ON stickers.invoice_id = invoices.id")
	query = query.Joins("JOIN customers ON stickers.customer_id = customers.id")
	query = query.Joins("JOIN shops ON shops.id = invoices.shop_id")
	query = query.Joins("JOIN doc_settings ON doc_settings.shop_id = invoices.shop_id")
	query = query.Joins("JOIN users ON stickers.user_id = users.id")
	query = query.Joins("JOIN invoice_details ON invoice_details.id = stickers.invoice_detail_id")
	query = query.Joins("JOIN products ON products.id = stickers.product_id")
	query = query.Where("stickers.invoice_id = ?", invoice_id)
	query = query.Where("stickers.sticker_is_del = ?", 0)
	query = query.Where("stickers.sticker_active_id = ?", 1)
	// query = query.Where("stickers.invoice_id = ? OR stickers.order_id = ? AND stickers.sticker_is_del = ?", invoice_id)
	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetStickerNoInvoiceList(sticker_id int, actv_id int, data *[]structs.StickerDetail) (err error) {
	query := configs.DB1.Table("stickers")
	query = query.Select("stickers.*,invoices.inv_code, customers.ctm_id, customers.ctm_fname, customers.ctm_lname, customers.ctm_fname_en, customers.ctm_lname_en, customers.ctm_prefix, customers.ctm_birthdate, users.user_fullname, doc_settings.sticker_font_size, doc_settings.sticker_width, shops.shop_name, shops.shop_name_en, shops.shop_phone, doc_settings.sticker_height, doc_settings.sticker_show_name, doc_settings.sticker_show_address, doc_settings.sticker_show_tel, doc_settings.sticker_show_date, doc_settings.sticker_show_expdate, doc_settings.sticker_show_detail, invoice_details.invd_discount, invoice_details.invd_amount, invoice_details.invd_total, invoice_details.invd_vat, shops.shop_address, shops.shop_address_en, shops.shop_district, shops.shop_district_en, shops.shop_amphoe, shops.shop_amphoe_en, shops.shop_province, shops.shop_province_en, shops.shop_zipcode,	shops.shop_zipcode_en, shops.shop_license, shops.shop_tax,products.pd_code_acc,products.pd_name_acc,doc_settings.show_product_acc,shops.shop_image,doc_settings.show_shop_image,doc_settings.show_shop_image")
	query = query.Joins("LEFT JOIN invoices ON stickers.invoice_id = stickers.id")
	query = query.Joins("LEFT JOIN customers ON stickers.customer_id = customers.id")
	query = query.Joins("LEFT JOIN shops ON shops.id = stickers.shop_id")
	query = query.Joins("LEFT JOIN doc_settings ON doc_settings.shop_id = stickers.shop_id")
	query = query.Joins("LEFT JOIN users ON stickers.user_id = users.id")
	query = query.Joins("LEFT JOIN invoice_details ON invoice_details.id = stickers.invoice_detail_id")
	query = query.Joins("JOIN products ON products.id = stickers.product_id")
	query = query.Where("stickers.id = ?", sticker_id)
	query = query.Where("stickers.sticker_is_del = ?", 0)
	//query = query.Where("stickers.sticker_active_id = ?", 1)
	if actv_id == 2 {
		query = query.Where("stickers.sticker_active_id = ?", 2)
	} else {
		query = query.Where("stickers.sticker_active_id = ?", 1)
	}
	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetStickerReceiptInArray(invoice_id string, data *[]structs.StickerDetail) (err error) {
	query := configs.DB1.Table("stickers")
	query = query.Select("stickers.*,invoices.inv_code, customers.ctm_id,customers.ctm_birthdate, customers.ctm_fname, customers.ctm_lname, customers.ctm_fname_en, customers.ctm_lname_en, customers.ctm_prefix, users.user_fullname, doc_settings.sticker_font_size, doc_settings.sticker_width, shops.shop_name, shops.shop_name_en, shops.shop_phone, doc_settings.sticker_height, doc_settings.sticker_show_name, doc_settings.sticker_show_address, doc_settings.sticker_show_tel, doc_settings.sticker_show_date, doc_settings.sticker_show_expdate, doc_settings.sticker_show_detail, shops.shop_address, shops.shop_address_en, shops.shop_district, shops.shop_district_en,	shops.shop_amphoe, shops.shop_amphoe_en, shops.shop_province, shops.shop_province_en,shops.shop_zipcode,	shops.shop_zipcode_en,products.pd_code_acc,products.pd_name_acc,doc_settings.show_product_acc,shops.shop_image,doc_settings.show_shop_image,doc_settings.show_shop_image")
	query = query.Joins("JOIN invoices ON stickers.invoice_id = invoices.id")
	query = query.Joins("JOIN customers ON stickers.customer_id = customers.id")
	query = query.Joins("JOIN shops ON shops.id = invoices.shop_id")
	query = query.Joins("JOIN doc_settings ON doc_settings.shop_id = invoices.shop_id")
	query = query.Joins("JOIN users ON stickers.user_id = users.id")
	query = query.Joins("JOIN products ON products.id = stickers.product_id")
	query = query.Where("stickers.sticker_is_del = ?", 0)
	query = query.Where("stickers.sticker_active_id = ?", 1)
	query = query.Where("stickers.invoice_id in (" + invoice_id + ")")
	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetStickerById(Id int, data *structs.StickerDetail) (err error) {
	query := configs.DB1.Table("stickers")
	query = query.Select("stickers.*,invoices.inv_code, customers.ctm_id, customers.ctm_fname, customers.ctm_lname, customers.ctm_fname_en, customers.ctm_lname_en, customers.ctm_prefix, users.user_fullname")
	query = query.Joins("JOIN invoices ON stickers.invoice_id = invoices.id")
	query = query.Joins("JOIN customers ON stickers.customer_id = customers.id")
	query = query.Joins("JOIN users ON stickers.user_id = users.id")
	query = query.Where("stickers.id = ?", Id)
	query = query.Where("stickers.sticker_is_del = 0")
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetDocSettingSticker(ShopId int, obj *structs.StickerDocSetting) (err error) {
	query := configs.DB1.Table("doc_settings")
	query = query.Select("doc_settings.*")
	query = query.Where("doc_settings.shop_id = ?", ShopId)
	query = query.Scan(obj)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateDocSettingSticker(ShopId int, obj *structs.StickerDocSetting) (err error) {
	query := configs.DB1.Table("doc_settings")
	query = query.Where("doc_settings.shop_id = ?", ShopId)
	query = query.Model(&obj)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&obj)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func AddLogStk(log *structs.LogStk) (err error) {
	query := configs.DBL1.Table("log_stickers").Create(&log)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdatePrintStickerStatus(Id int) (err error) {
	query := configs.DB1.Table("stickers")
	query = query.Where("stickers.id = ?", Id)
	query = query.Updates(map[string]interface{}{"stickers.sticker_print_label": 1})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdatePrescriptionStatus(Id int) (err error) {
	query := configs.DB1.Table("stickers")
	query = query.Where("stickers.id = ?", Id)
	query = query.Updates(map[string]interface{}{"stickers.sticker_print_order": 1})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateSticker(dt *structs.UpdateSticker) (err error) {
	query := configs.DB1.Table("stickers")
	query = query.Where("stickers.id = ?", dt.Id)
	query = query.Updates(map[string]interface{}{"stickers.sticker_topical": dt.Sticker_topical, "stickers.sticker_direction": dt.Sticker_direction})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetStickerHeadList(invoice_id string, dt *[]structs.StickerDetail2) (err error) {
	query := configs.DB1.Table("stickers")
	query = query.Select("stickers.*,invoices.inv_code, customers.ctm_id, customers.ctm_fname, customers.ctm_lname, customers.ctm_fname_en, customers.ctm_lname_en, customers.ctm_prefix, customers.ctm_birthdate, users.user_fullname, doc_settings.sticker_font_size, doc_settings.sticker_width, doc_settings.show_product_acc, shops.shop_name,shops.shop_image, shops.shop_name_en, shops.shop_phone, doc_settings.sticker_height, doc_settings.sticker_show_name, doc_settings.sticker_show_address, doc_settings.sticker_show_tel, doc_settings.sticker_show_date, doc_settings.sticker_show_expdate, doc_settings.sticker_show_detail,doc_settings.show_shop_image, shops.shop_address, shops.shop_address_en, shops.shop_district, shops.shop_district_en, shops.shop_amphoe, shops.shop_amphoe_en, shops.shop_province, shops.shop_province_en, shops.shop_zipcode,	shops.shop_zipcode_en, shops.shop_license, shops.shop_tax")
	query = query.Joins("JOIN invoices ON stickers.invoice_id = invoices.id")
	query = query.Joins("JOIN customers ON stickers.customer_id = customers.id")
	query = query.Joins("JOIN shops ON shops.id = invoices.shop_id")
	query = query.Joins("JOIN doc_settings ON doc_settings.shop_id = invoices.shop_id")
	query = query.Joins("JOIN users ON stickers.user_id = users.id")
	query = query.Where("stickers.invoice_id in (" + invoice_id + ")")
	query = query.Where("stickers.sticker_is_del = 0")
	query = query.Group("stickers.invoice_id")
	query = query.Order("stickers.id DESC")
	if err = query.Scan(&dt).Error; err != nil {
		return err
	}
	return nil
}
func GetStickerHeadListNoInVoice(order_id string, dt *[]structs.StickerDetail2) (err error) {
	query := configs.DB1.Table("stickers")
	query = query.Select("stickers.*,invoices.inv_code, customers.ctm_id, customers.ctm_fname, customers.ctm_lname, customers.ctm_fname_en, customers.ctm_lname_en, customers.ctm_prefix, customers.ctm_birthdate, users.user_fullname, doc_settings.sticker_font_size, doc_settings.sticker_width, shops.shop_name, shops.shop_name_en,shops.shop_image, shops.shop_phone, doc_settings.sticker_height, doc_settings.sticker_show_name, doc_settings.sticker_show_address, doc_settings.sticker_show_tel, doc_settings.sticker_show_date, doc_settings.sticker_show_expdate, doc_settings.sticker_show_detail,doc_settings.show_shop_image,doc_settings.show_product_acc, shops.shop_address, shops.shop_address_en, shops.shop_district, shops.shop_district_en, shops.shop_amphoe, shops.shop_amphoe_en, shops.shop_province, shops.shop_province_en, shops.shop_zipcode,shops.shop_zipcode_en, shops.shop_license, shops.shop_tax")
	query = query.Joins("LEFT JOIN invoices ON stickers.invoice_id = invoices.id")
	query = query.Joins("LEFT JOIN customers ON stickers.customer_id = customers.id")
	query = query.Joins("LEFT JOIN shops ON shops.id = stickers.shop_id")
	query = query.Joins("LEFT JOIN doc_settings ON doc_settings.shop_id = stickers.shop_id")
	query = query.Joins("LEFT JOIN users ON stickers.user_id = users.id")
	query = query.Where("stickers.order_id in (" + order_id + ")")
	query = query.Where("stickers.sticker_is_del = 0")
	query = query.Group("stickers.invoice_id")
	query = query.Order("stickers.id DESC")
	if err = query.Scan(&dt).Error; err != nil {
		return err
	}
	return nil
}

func GetShopStickerById(shopId int, objResponse *structs.StickerShop) error {
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

func CreateStickerDrugLabel(objQuery *structs.PrintStickerDrugLabel) (err error) {
	query := configs.DB1.Table("stickers").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}
