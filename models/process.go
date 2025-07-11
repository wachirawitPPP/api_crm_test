package models

import (
	"encoding/json"
	"linecrmapi/configs"
	"linecrmapi/structs"
	"time"
)

func GetProcessReceiptDetail(rcId int, data *[]structs.ProcessReceiptDetail) (err error) {
	query := configs.DB1.Table("receipt_details")
	query = query.Select("receipt_details.*")
	query = query.Where("receipt_details.receipt_id", rcId)
	query = query.Where("receipt_details.recd_is_active = ?", 1)
	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetProcessReceiptId(receipt_id int, dt *structs.ProcessReceipt) (err error) {
	query := configs.DB1.Table("receipts")
	query = query.Select("receipts.*")
	query = query.Where("receipts.id = ?", receipt_id)
	if err = query.Scan(&dt).Error; err != nil {
		return err
	}
	return nil
}

func GetProcessReceiptDetailCheck(receipt_id int, data *[]structs.ProcessReceiptDetail) (err error) {
	query := configs.DB1.Table("receipt_details")
	query = query.Select("receipt_details.*")
	query = query.Where("receipt_details.receipt_id = ?", receipt_id)
	query = query.Where("receipt_details.recd_type_id = ?", 1)
	query = query.Where("receipt_details.recd_is_active = ?", 1)
	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetProcessReceiptDetailCheckProduct(receipt_id int, checking_id int, data *[]structs.ProcessReceiptDetail) (err error) {
	query := configs.DB1.Table("receipt_details")
	query = query.Select("receipt_details.*")
	query = query.Where("receipt_details.receipt_id = ?", receipt_id)
	query = query.Where("receipt_details.checking_id = ?", checking_id)
	query = query.Where("receipt_details.recd_type_id = ?", 5)
	query = query.Where("receipt_details.recd_is_active = ?", 1)
	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func CreateChecks(objQuery *structs.ProcessCheck) (err error) {
	query := configs.DB1.Table("checks").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CreateChecksProduct(objQuery *structs.ProcessCheckProduct) (err error) {
	query := configs.DB1.Table("check_products").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetProcessCheckingChecksId(checkingId int, checking *structs.ProcessCheckChecking) (err error) {
	query := configs.DB1.Table("checkings")
	query = query.Select("checkings.*")
	query = query.Where("checkings.id = ?", checkingId)
	if err = query.Scan(&checking).Error; err != nil {
		return err
	}
	return nil
}

func GetProcessReceiptDetailCourse(receipt_id int, data *[]structs.ProcessReceiptDetail) (err error) {
	query := configs.DB1.Table("receipt_details")
	query = query.Select("receipt_details.*")
	query = query.Where("receipt_details.receipt_id = ?", receipt_id)
	query = query.Where("receipt_details.recd_type_id = ?", 2)
	query = query.Where("receipt_details.recd_is_active = ?", 1)
	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetProcessReceiptDetailCourseProduct(receipt_id int, course_id int, data *[]structs.ProcessReceiptDetail) (err error) {
	query := configs.DB1.Table("receipt_details")
	query = query.Select("receipt_details.*")
	query = query.Where("receipt_details.receipt_id = ?", receipt_id)
	query = query.Where("receipt_details.course_id = ?", course_id)
	query = query.Where("receipt_details.recd_type_id = ?", 5)
	query = query.Where("receipt_details.recd_is_active = ?", 1)
	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetProcessReceiptDetailProduct(receipt_id int, data *[]structs.ProcessReceiptDetail) (err error) {
	query := configs.DB1.Table("receipt_details")
	query = query.Select("receipt_details.*")
	query = query.Where("receipt_details.receipt_id = ?", receipt_id)
	query = query.Where("receipt_details.recd_type_id = ?", 3)
	query = query.Where("receipt_details.recd_is_active = ?", 1)
	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetCourseServicesId(coursesId int, course *structs.ProcessCourseService) (err error) {
	query := configs.DB1.Table("courses")
	query = query.Select("courses.*")
	query = query.Where("courses.id = ?", coursesId)
	if err = query.Scan(&course).Error; err != nil {
		return err
	}
	return nil
}

func CreateServiceProduct(objQuery *structs.ProcessServiceProduct) (err error) {
	query := configs.DB1.Table("service_products").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CreateServices(objQuery *structs.ProcessService) (err error) {
	query := configs.DB1.Table("services").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CreateServiceProductUsed(objQuery *structs.ProcessServiceProductUsed) (err error) {
	query := configs.DB1.Table("service_product_useds").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CreateServicesUsed(objQuery *structs.ProcessServiceUsed) (err error) {
	query := configs.DB1.Table("service_useds").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateService(service_id int, data *structs.ProcessServiceUpdate) (err error) {
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

func GetProcessCourseProduct(course_id int, cp *[]structs.ProcessCourseProductSet) (err error) {
	query := configs.DB1.Table("course_products")
	query = query.Select("course_products.*")
	query = query.Where("course_products.course_id = ?", course_id)
	query = query.Where("course_products.cp_is_active = ?", 1)
	query = query.Where("course_products.cp_is_del = ?", 0)
	if err = query.Scan(&cp).Error; err != nil {
		return err
	}
	return nil
}

func GetProcessProductSetId(product_id int, shop_id int, product *structs.ProcessProductSet) (err error) {
	query := configs.DB1.Table("products")
	query = query.Select("products.id AS id,products.id AS product_id,product_stores.id AS product_store_id, 1 AS pu_amount,product_units.id AS product_units_id,product_stores.pds_cost,products.pd_code,products.pd_name,product_units.pu_rate,ref_units.u_name,product_shop_prices.psp_price_opd,product_shop_prices.psp_price_ipd")
	query = query.Joins("INNER JOIN product_stores ON products.id = product_stores.product_id")
	query = query.Joins("INNER JOIN shop_stores ON product_stores.shop_store_id = shop_stores.id")
	query = query.Joins("INNER JOIN shops ON shop_stores.shop_id = shops.id")
	query = query.Joins("INNER JOIN product_units ON products.id = product_units.product_id")
	query = query.Joins("INNER JOIN ref_units ON product_units.unit_id = ref_units.id")
	query = query.Joins("INNER JOIN product_shop_prices ON product_units.id = product_shop_prices.product_unit_id")
	query = query.Where("products.id = ?", product_id)
	query = query.Where("product_units.pu_rate = ?", 1)
	query = query.Where("product_units.pu_is_del = ?", 0)
	query = query.Where("shop_stores.ss_type_id = ?", 1)
	query = query.Where("shop_stores.shop_id = ?", shop_id)
	query = query.Where("product_shop_prices.shop_id = ?", shop_id)
	// query = query.Where("product_shop_prices.psp_is_default = ?", 1)
	query = query.Where("product_shop_prices.psp_is_del = ?", 0)
	query = query.Where("products.pd_is_active = ?", 1)
	query = query.Where("products.pd_is_del = ?", 0)
	if err = query.Scan(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProcessProductStoreOrder(product_id int, product_stores_id int, product *structs.ProcessProductStoreOrder) (err error) {
	query := configs.DB1.Table("product_store_orders")
	query = query.Select("product_store_orders.*")
	query = query.Joins("INNER JOIN product_stores ON product_store_orders.product_store_id = product_stores.id")
	query = query.Where("product_stores.product_id = ?", product_id)
	query = query.Where("product_store_orders.product_store_id = ?", product_stores_id)
	query = query.Where("product_store_orders.pdso_is_active = ?", 1)
	query = query.Where("product_store_orders.pdso_total > ?", 0)
	query = query.Where("product_store_orders.pdso_expire > ?", time.Now().Format("2006-01-02")) // ไม่หมดอายุ
	query = query.Order("product_store_orders.pdso_expire ASC")
	query = query.Limit(1)
	if err = query.Scan(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProcessProductStoreOrderLast(product_id int, product_stores_id int, product *structs.ProcessProductStoreOrder) (err error) {
	query := configs.DB1.Table("product_store_orders")
	query = query.Select("product_store_orders.*")
	query = query.Joins("INNER JOIN product_stores ON product_store_orders.product_store_id = product_stores.id")
	query = query.Where("product_stores.product_id = ?", product_id)
	query = query.Where("product_store_orders.product_store_id = ?", product_stores_id)
	query = query.Where("product_store_orders.pdso_is_active = ?", 1)
	query = query.Where("product_store_orders.pdso_expire > ?", time.Now().Format("2006-01-02")) // ไม่หมดอายุ
	query = query.Order("product_store_orders.pdso_expire DESC")
	query = query.Limit(1)
	if err = query.Scan(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProcessProductStoreOrderId(product_store_order_id int, product *structs.ProcessProductStoreOrder) (err error) {
	query := configs.DB1.Table("product_store_orders")
	query = query.Select("product_store_orders.*")
	query = query.Joins("INNER JOIN product_stores ON product_store_orders.product_store_id = product_stores.id")
	query = query.Where("product_store_orders.id = ?", product_store_order_id)
	query = query.Where("product_store_orders.pdso_is_active = ?", 1)
	query = query.Where("product_store_orders.pdso_total > ?", 0)
	query = query.Where("product_store_orders.pdso_expire > ?", time.Now().Format("2006-01-02")) // ไม่หมดอายุ
	if err = query.Scan(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProcessProductStoreOrders(product_store_order_id int, product *structs.ProcessProductStoreOrder) (err error) {
	query := configs.DB1.Table("product_store_orders")
	query = query.Select("product_store_orders.*")
	query = query.Where("product_store_orders.id = ?", product_store_order_id)
	if err = query.Scan(product).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProductStoreOrders(product_store_orders_id int, data *structs.ProcessProductStoreOrder) (err error) {
	query := configs.DB1.Table("product_store_orders")
	query = query.Where("product_store_orders.id = ?", product_store_orders_id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	delete(inInterface, "pdso_expire")
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetProcessProductStoreId(product_store_id int, product *structs.ProcessProductStore) (err error) {
	query := configs.DB1.Table("product_stores")
	query = query.Select("product_stores.*")
	query = query.Where("product_stores.id = ?", product_store_id)
	if err = query.Scan(product).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProductStore(product_store_id int, data *structs.ProcessProductStore) (err error) {
	query := configs.DB1.Table("product_stores")
	query = query.Where("product_stores.id = ?", product_store_id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CreateProductStoreHistory(objQuery *structs.ProcessProductStoreHistory) (err error) {
	query := configs.DB1.Table("product_store_historys").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CreateProductStoreHistoryExp(objQuery *structs.ProcessProductStoreHistoryExp) (err error) {
	query := configs.DB1.Table("product_store_historys").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CreateProductSticker(objQuery *structs.ProcessProductSticker) (err error) {
	query := configs.DB1.Table("stickers").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetProcessReceiptDetailCoin(receipt_id int, data *[]structs.ProcessReceiptDetail) (err error) {
	query := configs.DB1.Table("receipt_details")
	query = query.Select("receipt_details.*")
	query = query.Where("receipt_details.receipt_id = ?", receipt_id)
	query = query.Where("receipt_details.recd_type_id = ?", 4)
	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetProcessCustomerCoin(customer_id int, customer *structs.ProcessCoin) (err error) {
	query := configs.DB1.Table("customers")
	query = query.Select("customers.id,customers.ctm_coin,customers.ctm_update")
	query = query.Where("customers.id = ?", customer_id)
	if err = query.Scan(&customer).Error; err != nil {
		return err
	}
	return nil
}

func GetProcessCoinId(coin_id int, data *structs.CoinCustomerGroupId) (err error) {
	query := configs.DB1.Table("coins")
	query = query.Select("coins.id,coins.customer_group_id")
	query = query.Where("coins.id = ?", coin_id)
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetProcessProductTypeId(product_id int, customer *structs.ProcessProductType) (err error) {
	query := configs.DB1.Table("products")
	query = query.Select("products.id,products.pd_type_id,products.pd_name_acc")
	query = query.Where("products.id = ?", product_id)
	if err = query.Scan(&customer).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProcessCustomerCoin(customer_id int, data *structs.ProcessCoin) (err error) {
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

func UpdateProcessCustomerCoinCancel(customer_id int, data *structs.ProcessCoinCancelUpdate) (err error) {
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

func GetProcessQueueById(queueId int, objQuery *structs.ProcessQueue) error {
	query := configs.DB1.Table("queues")
	query = query.Select("queues.*")
	query = query.Where("queues.id = ?", queueId)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func UpdateQueueProcessId(queue_id int, data *structs.QueueProcessUpdate) (err error) {
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

func GetProcessInvoiceById(invoice_id int, objQuery *structs.ProcessInvoice) error {
	query := configs.DB1.Table("invoices")
	query = query.Select("invoices.*")
	query = query.Where("invoices.id = ?", invoice_id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func UpdateInvoiceProcessId(invoice_id int, data *structs.InvoiceProcessUpdate) (err error) {
	// query := configs.DB1.Table("invoices")
	// query = query.Where("invoices.id = ?", invoice_id)
	// query = query.Model(&data)
	// query = query.Updates(&data)
	// if err = query.Error; err != nil {
	// 	return err
	// }
	// return nil
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

func GetReceiptProductStoreHistory(receipt_id int, objQuery *[]structs.ProcessProductStoreHistory) (err error) {
	query := configs.DB1.Table("product_store_historys")
	query = query.Select("product_store_historys.*")
	query = query.Where("product_store_historys.receipt_id = ?", receipt_id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func UpdateCaneclCheckProductId(receipt_id int, data *structs.CheckProductUpdate) (err error) {
	query := configs.DB1.Table("check_products")
	query = query.Where("check_products.receipt_id = ?", receipt_id)
	query = query.Where("check_products.chkp_is_active != ?", 0)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateCaneclCheckId(receipt_id int, data *structs.CheckUpdate) (err error) {
	query := configs.DB1.Table("checks")
	query = query.Where("checks.receipt_id = ?", receipt_id)
	query = query.Where("checks.chk_is_active != ?", 0)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateCaneclServiceProductUsedId(receipt_id int, data *structs.ServiceProductUsedUpdate) (err error) {
	query := configs.DB1.Table("service_product_useds")
	query = query.Where("service_product_useds.receipt_id = ?", receipt_id)
	query = query.Where("service_product_useds.serpu_is_active != ?", 0)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateCaneclServiceProductId(receipt_id int, data *structs.ServiceProductUpdate) (err error) {
	query := configs.DB1.Table("service_products")
	query = query.Where("service_products.receipt_id = ?", receipt_id)
	query = query.Where("service_products.serp_is_active != ?", 0)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateCaneclServiceUsedId(receipt_id int, data *structs.ServiceUsedUpdate) (err error) {
	query := configs.DB1.Table("service_useds")
	query = query.Where("service_useds.receipt_id = ?", receipt_id)
	query = query.Where("service_useds.seru_is_active != ?", 0)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateCaneclServiceId(receipt_id int, data *structs.ServiceUpdate) (err error) {
	query := configs.DB1.Table("services")
	query = query.Where("services.receipt_id = ?", receipt_id)
	query = query.Where("services.ser_is_active != ?", 0)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateInvoiceCancelId(invoice_id int, data *structs.InvoiceProcessUpdate) (err error) {
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

func UpdateCaneclStickerId(receipt_id int, data *structs.StickerUpdate) (err error) {
	query := configs.DB1.Table("stickers")
	query = query.Where("stickers.receipt_id = ?", receipt_id)
	query = query.Where("stickers.sticker_active_id != ?", 0)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetProcessProductStoreOrderExpire(data *[]structs.ProcessProductStoreOrderExpire) (err error) {
	query := configs.DB1.Table("product_store_orders")
	query = query.Select("product_store_orders.*")
	// query = query.Joins("INNER JOIN product_stores ON product_store_orders.product_store_id = product_stores.id")
	query = query.Where("product_store_orders.pdso_is_active = ?", 1)
	query = query.Where("product_store_orders.pdso_total > ?", 0)
	query = query.Where("product_store_orders.pdso_expire <= ?", time.Now().Format("2006-01-02")) // หมดอายุแล้ว
	// query = query.Limit(1)
	if err = query.Scan(data).Error; err != nil {
		return err
	}
	return nil
}

func GetProcessProductStoreOrderExpireHistory(data *[]structs.ProcessProductStoreOrderExpire) (err error) {
	query := configs.DB1.Table("product_store_orders")
	query = query.Select("product_store_orders.*")
	query = query.Where("product_store_orders.pdso_is_active = ?", 1)
	query = query.Where("product_store_orders.pdso_expire <= '2024-06-24'")
	// query = query.Where("product_store_orders.pdso_expire = '2024-05-23'") // หมดอายุแล้ว
	// query = query.Limit(1)
	if err = query.Scan(data).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProductStoreOrdersExpire(product_store_orders_id int, data *structs.ProcessProductStoreOrderExpire) (err error) {
	query := configs.DB1.Table("product_store_orders")
	query = query.Where("product_store_orders.id = ?", product_store_orders_id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetProcessProductStoreExpire(product_store_id int, product *structs.ProcessProductStoreExpire) (err error) {
	query := configs.DB1.Table("product_stores")
	query = query.Select("product_stores.*,shop_stores.shop_id")
	query = query.Joins("INNER JOIN shop_stores ON product_stores.shop_store_id = shop_stores.id")
	query = query.Where("product_stores.id = ?", product_store_id)
	if err = query.Scan(product).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProductStoreExpire(product_store_id int, data *structs.UpdateProcessProductStoreExpire) (err error) {
	query := configs.DB1.Table("product_stores")
	query = query.Where("product_stores.id = ?", product_store_id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetProductIdProcess(product_id int, data *structs.ProductIdProcess) (err error) {
	query := configs.DB1.Table("products")
	query = query.Select("products.id,products.pd_code,products.pd_name")
	query = query.Where("products.id = ?", product_id)
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetProcessCustomerId(customer_id int, customer *structs.CustomerIdProduct) (err error) {
	query := configs.DB1.Table("customers")
	query = query.Select("customers.id,customers.ctm_prefix,customers.ctm_fname,customers.ctm_lname,customers.ctm_fname_en,customers.ctm_lname_en,customers.ctm_gender,customers.ctm_tel")
	query = query.Where("customers.id = ?", customer_id)
	if err = query.Scan(&customer).Error; err != nil {
		return err
	}
	return nil
}

// func GetProductRateId(shop_id int, product_id int, product_unit_id int, data *structs.ProductRate) (err error) {
// 	query := configs.DB1.Table("product_units")
// 	query = query.Select("product_units.pu_rate")
// 	query = query.Joins("INNER JOIN product_shops ON product_units.product_id = product_shops.product_id")
// 	query = query.Where("product_units.product_id = ?", product_id)
// 	query = query.Where("product_units.id = ?", product_unit_id)
// 	query = query.Where("product_shops.shop_id = ?", shop_id)
// 	query = query.Where("product_units.pu_is_del = ?", 0)
// 	if err = query.Scan(&data).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

func GetProductRate1(shop_id int, product_id int, data *structs.ProductRate) (err error) {
	query := configs.DB1.Table("product_units")
	query = query.Select("products.id,products.pd_type_id,products.pd_name_acc,product_units.pu_rate,ref_units.u_name,ref_units.u_name_en")
	query = query.Joins("INNER JOIN product_shops ON product_units.product_id = product_shops.product_id")
	query = query.Joins("INNER JOIN ref_units ON product_units.unit_id = ref_units.id")
	query = query.Joins("INNER JOIN products ON product_units.product_id = products.id")
	query = query.Where("product_units.product_id = ?", product_id)
	query = query.Where("product_units.pu_rate = ?", 1)
	query = query.Where("product_shops.shop_id = ?", shop_id)
	query = query.Where("product_units.pu_is_del = ?", 0)
	query = query.Where(" ref_units.u_is_del = ?", 0)
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetProcessServiceQueueCourse(receipt_detail_id int, data *structs.ProcessServiceQueueCourse) (err error) {
	query := configs.DB1.Table("receipt_details")
	query = query.Select("queue_courses.id AS queue_course_id")
	query = query.Joins("INNER JOIN invoice_details ON receipt_details.invoice_detail_id = invoice_details.id")
	query = query.Joins("INNER JOIN order_details ON invoice_details.order_detail_id = order_details.id")
	query = query.Joins("INNER JOIN queue_courses ON order_details.queue_course_id = queue_courses.id")
	query = query.Where("receipt_details.id = ?", receipt_detail_id)
	if err = query.Scan(data).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProcessServiceQueueCourseUpdate(queue_course_id int, data *structs.ProcessServiceQueueCourseUpdate) (err error) {
	query := configs.DB1.Table("queue_courses")
	query = query.Where("queue_courses.id = ?", queue_course_id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateCaneclQueueCourseId(receipt_id int) (err error) {
	query := configs.DB1.Table("queue_courses")
	query = query.Where("queue_courses.receipt_id = ?", receipt_id)
	query = query.Updates(map[string]interface{}{"service_id": nil, "receipt_id": nil, "quec_modify": time.Now().Format("2006-01-02 15:04:05")})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CreateCoinHistory(objQuery *structs.AddCoinHistory) (err error) {
	query := configs.DB1.Table("coin_historys").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CreatePointHistory(objQuery *structs.AddPointHistory) (err error) {
	query := configs.DB1.Table("point_historys").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetProcessInvoiceDetailById(receipt_id int, receipt_detail_id int, data *structs.ProcessInvoiceDetailById) (err error) {
	query := configs.DB1.Table("receipts")
	query = query.Select("receipts.invoice_id,receipt_details.invoice_detail_id")
	query = query.Joins("INNER JOIN receipt_details ON receipts.id = receipt_details.receipt_id")
	query = query.Where("receipts.id = ?", receipt_id)
	query = query.Where("receipt_details.id = ?", receipt_detail_id)
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetProcessCheck(invoice_detail_id int, checking_id int, data *structs.ProcessCheck) (err error) {
	query := configs.DB1.Table("checks")
	query = query.Select("checks.*")
	query = query.Where("checks.invoice_detail_id = ?", invoice_detail_id)
	query = query.Where("checks.checking_id = ?", checking_id)
	query = query.Where("checks.chk_is_active = ?", 1)
	if err = query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCheckInvoice(check_id int, receipt_id int, receipt_detail_id int) (err error) {
	query := configs.DB1.Table("checks")
	query = query.Where("checks.id = ?", check_id)
	query = query.Where("checks.chk_is_active = ?", 1)
	query = query.Updates(map[string]interface{}{"checks.receipt_id": receipt_id, "checks.receipt_detail_id": receipt_detail_id, "checks.chk_update": time.Now().Format("2006-01-02 15:04:05")})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateCaneclCheckReceiptId(receipt_id int, data *structs.CheckUpdate) (err error) {
	query := configs.DB1.Table("checks")
	query = query.Where("checks.receipt_id = ?", receipt_id)
	query = query.Where("checks.invoice_id IS NULL")
	query = query.Where("checks.chk_is_active != ?", 0)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateCancalCheckReceipt(invoice_id int) (err error) {
	query := configs.DB1.Table("checks")
	query = query.Where("checks.invoice_id = ?", invoice_id)
	query = query.Where("checks.chk_is_active != ?", 0)
	query = query.Updates(map[string]interface{}{"checks.receipt_id": nil, "checks.receipt_detail_id": nil, "checks.chk_update": time.Now().Format("2006-01-02 15:04:05")})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}
