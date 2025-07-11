package models

import (
	"linecrmapi/configs"
	"linecrmapi/structs"
	"time"
)

func CreateOrderOnline(createObj *structs.OrderOnline) (err error) {
	tx := configs.DB1.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err = tx.Table("order_onlines").Create(&createObj).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func UpdateOrderOnline(updateObj *structs.UpdateOrderOnline) (err error) {
	tx := configs.DB1.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	updateObj.Or_update = time.Now().Format("2006-01-02 15:04:05")
	if err = tx.Table("order_onlines").Updates(&updateObj).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func PaymentOrderOnline(updateObj *structs.ObjPaymentOrderOnline) (err error) {
	tx := configs.DB1.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	update := make(map[string]interface{})
	update["or_update"] = time.Now().Format("2006-01-02 15:04:05")
	update["payment_status"] = updateObj.Payment_status
	update["payment_type_id"] = updateObj.Payment_type_id
	update["payment_ref"] = updateObj.Payment_ref
	if err = tx.Table("order_onlines").Where("id = ?", updateObj.Order_online_id).Updates(&update).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func CreateOrderOnlineDetail(createObj *[]structs.OrderOnlineDetail) (err error) {
	tx := configs.DB1.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err = tx.Table("order_online_detail").CreateInBatches(&createObj, 400).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func UpdateOrderOnlineDetail(updateObj *structs.UpdateOrderOnlineDetail) (err error) {
	tx := configs.DB1.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	updateObj.Ord_update = time.Now().Format("2006-01-02 15:04:05")
	if err = tx.Table("order_online_detail").Updates(&updateObj).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func SearchOrderOnline(filter structs.ObjPayloadSearchOrderOnline, order *[]structs.OrderOnline) (err error) {
	query := configs.DB1.Table("order_onlines")
	query = query.Where("order_onlines.or_is_del = 0")
	if filter.Customer_online_id > 0 {
		query = query.Where("order_onlines.customer_online_id  = ?", filter.Customer_online_id)
	}
	if filter.Order_online_id > 0 {
		query = query.Where("order_onlines.id = ?", filter.Order_online_id)
	}
	if filter.Search_text != "" {
		query = query.Where("order_onlines.or_code LIKE ? OR order_onlines.or_website_name LIKE ? OR order_onlines.or_fullname LIKE ? OR order_onlines.payment_ref LIKE ?", "%"+filter.Search_text+"%", "%"+filter.Search_text+"%", "%"+filter.Search_text+"%", "%"+filter.Search_text+"%")
	}
	if filter.Or_website_id > 0 {
		query = query.Where("order_onlines.or_website_id = ?", filter.Or_website_id)
	}
	if filter.Or_is_active >= 0 {
		query = query.Where("order_onlines.or_is_active = ?", filter.Or_is_active)
	}
	if filter.Payment_status > 0 {
		query = query.Where("order_onlines.payment_status = ?", filter.Payment_status)
	}
	if filter.Or_date != "" {
		query = query.Where("DATE(order_onlines.or_date) = ?", filter.Or_date)
	}
	// <pagination
	if filter.Per_page > 0 && filter.Active_page > 0 {
		offset := (filter.Active_page - 1) * filter.Per_page
		query = query.Limit(filter.Per_page)
		query = query.Offset(offset)
	}
	// pagination>
	query = query.Order("order_onlines.or_date DESC")
	if err = query.Find(&order).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderOnline(id int, order *structs.OrderOnline) (err error) {
	query := configs.DB1.Table("order_onlines")
	query = query.Where("order_onlines.id = ?", id)
	query = query.Where(" order_onlines.or_is_del = 0")
	if err = query.First(&order).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderOnlineDetailById(id int, order *[]structs.OrderOnlineDetail) (err error) {
	query := configs.DB1.Table("order_online_detail")
	query = query.Where("order_online_detail.order_online_id = ?", id)
	query = query.Where(" order_online_detail.ord_is_del = 0")
	if err = query.Find(&order).Error; err != nil {
		return err
	}
	return nil
}
func GetOrderOnlineDetailByIds(ids []int, ord_is_active int, order *[]structs.OrderOnlineDetail) (err error) {
	query := configs.DB1.Table("order_online_detail")
	query = query.Where("order_online_detail.order_online_id IN ?", ids)
	query = query.Where(" order_online_detail.ord_is_del = 0")
	if ord_is_active >= 0 {
		query = query.Where("order_online_detail.ord_is_active = 1 ")
	}
	if err = query.Find(&order).Error; err != nil {
		return err
	}
	return nil
}

func CancelOrderOnline(payload structs.ObjCancelOrderOnline) (err error) {
	query := configs.DB1.Table("order_onlines")
	query = query.Where("order_onlines.id = ?", payload.Order_online_id)
	if err = query.Updates(map[string]interface{}{"or_is_active": 0, "user_id_cancel": payload.User_id_cancel, "or_update": time.Now().Format("2006-01-02 15:04:05")}).Error; err != nil {
		return err
	}
	return nil
}

func CancelOrderOnlineDetail(id int) (err error) {
	query := configs.DB1.Table("order_online_detail")
	query = query.Where("order_online_detail.order_online_id = ?", id)
	if err = query.Updates(map[string]interface{}{"ord_is_active": 0, "ord_update": time.Now().Format("2006-01-02 15:04:05")}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteOrderOnlineDetail(id int) (err error) {
	query := configs.DB1.Table("order_online_detail")
	query = query.Where("order_online_detail.id = ?", id)
	if err = query.Updates(map[string]interface{}{"ord_is_del": 1, "ord_update": time.Now().Format("2006-01-02 15:04:05")}).Error; err != nil {
		return err
	}
	return nil
}

func GetCustomerOnline(co_id int, result *structs.CustomerOnline) (err error) {
	query := configs.DB1.Table("customer_onlines")
	query = query.Select("customer_onlines.*")
	query = query.Where("customer_onlines.id = ?", co_id)
	query = query.Where("customer_onlines.co_is_del = 0")
	if err := query.First(&result).Error; err != nil {
		return err
	}
	return nil
}

func GetCustomerOnlineInshop(shop int, co_citizen_id string, result *[]structs.ObjQueryGetCustomerPagination) (err error) {
	query := configs.DB1.Table("customers")
	query = query.Where("customers.shop_id = ? and customers.ctm_citizen_id = ?", shop, co_citizen_id)
	query = query.Where("customers.ctm_is_del = 0")
	if err := query.Find(&result).Error; err != nil {
		return err
	}
	return nil
}

func GetCustomerOrderOnlineHistory(customer_online_id int, order *[]structs.OrderOnline) (err error) {
	query := configs.DB1.Table("order_onlines")
	query = query.Where("order_onlines.or_is_del = 0 and order_onlines.customer_online_id = ?", customer_online_id)
	query = query.Order("order_onlines.or_date DESC")
	if err = query.Find(&order).Error; err != nil {
		return err
	}
	return nil
}
