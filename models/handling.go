package models

import (
	"linecrmapi/configs"
	"linecrmapi/structs"

	"gorm.io/gorm"
)

func GetHandling(invoiceId int, queueId int, objQuery *[]structs.Handling) (err error) {
	query := configs.DB1.Table("handlings")
	query = query.Select("handlings.*, users.user_fullname, users.user_fullname_en, roles.role_name_th AS role_name")
	query = query.Joins("LEFT JOIN users ON users.id = handlings.user_id")
	query = query.Joins("LEFT JOIN roles ON roles.id = handlings.role_id")
	query = query.Where("handlings.hand_is_del = ?", 0)
	if invoiceId != 0 {
		query = query.Where("handlings.invoice_id = ?", invoiceId)
	}
	if queueId != 0 {
		query = query.Where("handlings.queue_id = ?", queueId)
	}
	query = query.Find(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetHandlingUser(objPayload *structs.ObjPayloadGetHandlingUser, objQuery *[]structs.ObjResponseGetHandlingUser) error {
	query := configs.DB1.Table("user_shops")
	query = query.Select("user_shops.id AS user_shop_id, users.id AS user_id, users.user_email, users.user_fullname, users.user_fullname_en, shop_roles.role_id")
	query = query.Joins("LEFT JOIN users ON users.id = user_shops.user_id")
	query = query.Joins("LEFT JOIN shop_roles ON shop_roles.id = user_shops.shop_role_id")
	query = query.Where("user_shops.shop_id = ?", objPayload.ShopId)
	query = query.Where("user_shops.us_is_active = ?", 1)
	query = query.Where("users.user_is_active = ?", 1)
	query = query.Where("users.user_email LIKE ? OR users.user_fullname LIKE ? OR users.user_tel LIKE ?", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%")
	query = query.Order("users.user_fullname ASC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetHandlingUserById(userShopId int, objQuery *structs.ObjQueryUserHandling) (err error) {
	query := configs.DB1.Table("user_shops")
	query = query.Select("user_shops.*, shop_roles.role_id")
	query = query.Joins("LEFT JOIN shop_roles ON shop_roles.id = user_shops.shop_role_id")
	query = query.Where("user_shops.id = ?", userShopId)
	query = query.Find(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CreateHandlingBatch(objCreate *[]Handling) error {
	query := configs.DB1.CreateInBatches(&objCreate, 24)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func DeleteHandlingNotInIds(objPayload *structs.ObjPayloadUpdateHandling, notInIds []int) (err error) {
	objQuery := Handling{}
	query := configs.DB1.Table("handlings")
	query = query.Where("hand_is_del = ?", 0)
	if *objPayload.InvoiceId != 0 {
		query = query.Where("handlings.invoice_id = ?", objPayload.InvoiceId)
	}
	if *objPayload.QueueId != 0 {
		query = query.Where("handlings.queue_id = ?", objPayload.QueueId)
	}
	query = query.Where("id NOT IN ?", notInIds)
	query = query.Model(objQuery)
	query = query.Updates(map[string]interface{}{"hand_is_del": 1})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func ClearHandling(objPayload *structs.ObjPayloadUpdateHandling) (err error) {
	objQuery := Handling{}
	query := configs.DB1.Table("handlings")
	if *objPayload.InvoiceId != 0 {
		query = query.Where("handlings.invoice_id = ?", objPayload.InvoiceId)
	}
	if *objPayload.QueueId != 0 {
		query = query.Where("handlings.queue_id = ?", objPayload.QueueId)
	}
	query = query.Where("hand_is_del = ?", 0)
	query = query.Model(objQuery)
	query = query.Updates(map[string]interface{}{"hand_is_del": 1})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetCommissionById(commissionId int, objQuery *Commission) error {
	query := configs.DB1.Preload("CommissionLevels", func(db *gorm.DB) *gorm.DB {
		return db.Where("commission_level_is_del = ?", 0).Order("commission_level_rate DESC")
	})
	query = query.Where("commissions.commission_is_del = ?", 0)
	query = query.Where("commissions.id = ?", commissionId)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetFeeById(feeId int, objQuery *Fee) error {
	query := configs.DB1.Table("fees")
	query = query.Select("fees.*")
	query = query.Where("fees.fee_is_del = ?", 0)
	query = query.Where("fees.id = ?", feeId)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetReceiptType2(invoiceId int, objQuery *Receipt) error {
	query := configs.DB1.Table("receipts")
	query = query.Select("receipts.*")
	query = query.Where("receipts.rec_type_id = ?", 2)
	query = query.Where("receipts.rec_is_active != ?", 0)
	query = query.Where("receipts.invoice_id = ?", invoiceId)
	query = query.Limit(1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueTimeEnd_(queueId int, objQuery *structs.ObjQueryQueueTimeEnd) error {
	query := configs.DB1.Table("queues")
	query = query.Select("queues.id, queues.que_time_end")
	query = query.Where("queues.que_status_id = ?", 4)
	query = query.Where("queues.id = ?", queueId)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetInvoiceChecking_(invoiceId int, objQuery *[]structs.ObjQueryInvoiceChecking) error {
	query := configs.DB1.Table("invoices")
	query = query.Select("invoices.id AS invoice_id, invoice_details.id AS invoice_detail_id, invoice_details.invd_total, checkings.id AS checking_id, checkings.checking_fee_df, checkings.checking_fee_nr, checkings.checking_fee_tr, checkings.checking_fee, invoice_details.invd_qty")
	query = query.Joins("LEFT JOIN invoice_details ON invoice_details.invoice_id = invoices.id")
	query = query.Joins("LEFT JOIN checkings ON checkings.id = invoice_details.checking_id")
	query = query.Where("invoices.id = ?", invoiceId)
	query = query.Where("invoices.inv_is_active != ?", 0)
	query = query.Where("invoice_details.invd_type_id = ?", 1)
	query = query.Where("invoice_details.invd_is_active = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueService_(queueId int, objQuery *[]structs.ObjQueryQueueService) error {
	query := configs.DB1.Table("queues")
	query = query.Select("queues.id AS queue_id, queue_courses.id AS queue_course_id, services.id AS service_id, services.ser_price_total, courses.id AS course_id, courses.course_fee_df, courses.course_fee_nr, courses.course_fee_tr, courses.course_fee, queue_courses.quec_qty")
	query = query.Joins("LEFT JOIN queue_courses ON queue_courses.queue_id = queues.id")
	query = query.Joins("LEFT JOIN courses ON courses.id = queue_courses.course_id")
	query = query.Joins("LEFT JOIN services ON services.id = queue_courses.service_id")
	query = query.Where("queues.id = ?", queueId)
	query = query.Where("queues.que_status_id = ?", 4)
	query = query.Where("queue_courses.quec_is_active = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
