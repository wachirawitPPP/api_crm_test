package models

import (
	"encoding/json"
	"linecrmapi/configs"
	"linecrmapi/structs"
)

func AddLogUser(logU *structs.LogUser) (err error) {
	query := configs.DBL1.Table("log_user").Create(&logU)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetUserByIdShop(user_id int, shopId int, user *User) (err error) {
	query := configs.DB1.Table("users")
	query = query.Select("users.*, user_shops.shop_id, shops.shop_name, shops.shop_name_en,	shop_roles.role_id,	shop_roles.role_name_th, shop_roles.role_name_en")
	query = query.Joins("LEFT JOIN user_shops ON user_shops.user_id = users.id")
	query = query.Joins("LEFT JOIN shops ON user_shops.shop_id = shops.id")
	query = query.Joins("LEFT JOIN shop_roles ON user_shops.shop_role_id = shop_roles.id")
	query = query.Where("users.id = ?", user_id)
	query = query.Where("user_shops.shop_id = ?", shopId)
	query = query.Where("users.user_is_active = 1")
	if err := query.Scan(&user).Error; err != nil {
		// configs.DB1.Close()
		return err
	}
	return nil
}

func GetUserById(user_id int, shopId int, user *User) (err error) {
	query := configs.DB1.Table("users")
	query = query.Select("users.*, user_shops.shop_id, shops.shop_name, shops.shop_name_en,	shop_roles.role_id,	shop_roles.role_name_th, shop_roles.role_name_en,timesets.timeset_open,timesets.timeset_close")
	query = query.Joins("JOIN user_shops ON user_shops.user_id = users.id")
	query = query.Joins("JOIN shops ON user_shops.shop_id = shops.id")
	query = query.Joins("JOIN shop_roles ON user_shops.shop_role_id = shop_roles.id")
	query = query.Joins("LEFT JOIN timesets ON shops.id = timesets.shop_id")
	query = query.Where("users.id = ?", user_id)
	query = query.Where("user_shops.shop_id = ?", shopId)
	query = query.Where("users.user_is_active = 1")
	if err := query.Scan(&user).Error; err != nil {
		// configs.DB1.Close()
		return err
	}
	return nil
}

func GetCheckUserById(user_id int, user *User) (err error) {
	query := configs.DB1.Table("users")
	query = query.Select("users.*")
	query = query.Where("users.id = ?", user_id)
	query = query.Where("users.user_is_active = 1")
	if err := query.Scan(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByShopId(user_id int, shop_id int, user *User) (err error) {
	query := configs.DB1.Table("users")
	query = query.Select("users.*, user_shops.shop_id, shops.shop_name, shops.shop_name_en,	shop_roles.id AS shop_role_id, shop_roles.role_id,	shop_roles.role_name_th, shop_roles.role_name_en")
	query = query.Joins("INNER JOIN user_shops ON user_shops.user_id = users.id")
	query = query.Joins("INNER JOIN shops ON user_shops.shop_id = shops.id")
	query = query.Joins("INNER JOIN shop_roles ON user_shops.shop_role_id = shop_roles.id")
	query = query.Where("users.id = ?", user_id)
	query = query.Where("user_shops.shop_id = ?", shop_id)
	query = query.Where("users.user_is_active = 1")
	if err := query.Scan(&user).Error; err != nil {
		// configs.DB1.Close()
		return err
	}
	return nil
}

func GetUserByEmail(email string, user *structs.ProfileUser) (err error) {
	query := configs.DB1.Table("users")
	query = query.Where("user_email = ?", email)
	if err := query.Scan(&user).Error; err != nil {
		return err
	}
	return nil
}

func VerifiedUser(user_id int, data *structs.ObjUserVerified) (err error) {
	query := configs.DB1.Table("users")
	query = query.Where("ID = ?", user_id)
	query = query.Model(&data)
	query = query.Updates(&data)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func ChangePwdUser(user_id int, data *structs.ObjChangePwdUser) (err error) {
	query := configs.DB1.Table("users")
	query = query.Where("ID = ?", user_id)
	query = query.Model(&data)
	query = query.Updates(&data)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func AddVFUser(uVF *structs.VerifiedUser) (err error) {
	query := configs.DB1.Table("user_verified").Create(&uVF)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetVFUserByToken(token string, user *structs.VerifiedUserFull) (err error) {
	query := configs.DB1.Table("user_verified")
	query = query.Select("user_verified.*, users.*")
	query = query.Joins("JOIN users on users.ID = user_verified.user_id")
	query = query.Where("user_verified.uvf_token = ?", token)
	query = query.Where("users.user_is_activate = 0")
	if err := query.Find(&user).Error; err != nil {
		return err
	}
	return nil
}

func AddFPUser(uVF *structs.ForgetPwdStamp) (err error) {
	query := configs.DB1.Table("user_forget_pwd").Create(&uVF)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user_id int, data *structs.ObjUpdateUser) (err error) {
	query := configs.DB1.Table("users")
	query = query.Where("users.id = ?", user_id)
	query = query.Model(&data)
	query = query.Updates(&data)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateUserSettime(user_id int, data *structs.ObjUpdateUserSettime) (err error) {
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)

	query := configs.DB1.Table("users")
	query = query.Where("users.id  = ?", user_id)
	query = query.Model(&data)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateUFPUser(user_id int, data *structs.ObjUpdateExpireStatus) (err error) {
	query := configs.DB1.Table("users")
	query = query.Where("users.id = ?", user_id)
	query = query.Model(&data)
	query = query.Updates(&data)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetFPUserByToken(token string, user *structs.ForgetPwdUser) (err error) {
	query := configs.DB1.Table("user_forget_pwd")
	query = query.Select("user_forget_pwd.*, u.user_email as user_email, u.user_is_activate as user_is_activate")
	query = query.Joins("JOIN user u on u.user_id = user_forget_pwd.user_id")
	query = query.Where("ufp_token = ?", token)
	query = query.Where("user_is_activate = 1")
	query = query.Where("user_forget_pwd.ufp_is_expired = 0")
	if err := query.Find(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetFPUserByPWD(pwd string, user *User) (err error) {
	query := configs.DB1.Table("users")
	query = query.Where("password = ?", pwd)
	query = query.Where("user_is_activate = 1")
	if err := query.Find(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetDay(data *[]structs.Days) (err error) {
	query := configs.DB1.Table("days")
	query = query.Select("days.*")
	query = query.Order("days.id ASC")
	if err := query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetSettime(day_id int, user_id int, shop_id int, data *structs.UserSettime) (err error) {
	query := configs.DB1.Table("user_settime")
	query = query.Select("user_settime.*")
	query = query.Where("user_settime.day_id = ?", day_id)
	query = query.Where("user_settime.user_id = ?", user_id)
	query = query.Where("user_settime.shop_id = ?", shop_id)
	if err := query.Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func AddSettime(data *structs.UserSettime) (err error) {
	query := configs.DB1.Table("user_settime").Create(&data)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateSettime(id int, data *structs.ObjUpdateUserSettimeList) (err error) {
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)

	query := configs.DB1.Table("user_settime")
	query = query.Where("user_settime.id = ?", id)
	query = query.Model(&data)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func DeleteSettime(user_id int, day_id int, data *structs.UserSettime) (err error) {
	query := configs.DB1.Table("user_settime")
	query = query.Where("user_settime.user_id = ?", user_id)
	query = query.Where("user_settime.day_id = ?", day_id)
	query = query.Delete(&data)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func SearchSelectUser(objPayload *structs.ObjPayloadSearchSelectUser, objQuery *[]structs.ObjResponseSearchSelectUser) error {
	query := configs.DB1.Table("user_shops")
	query = query.Select("users.id, users.user_email, users.user_image, users.user_fullname, users.user_fullname_en")
	query = query.Joins("LEFT JOIN users ON users.id = user_shops.user_id")
	query = query.Where("user_shops.shop_id = ?", objPayload.ShopId)
	query = query.Where("users.user_is_active = ?", 1)
	query = query.Where("users.user_email LIKE ? OR users.user_fullname LIKE ? OR users.user_tel LIKE ?", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%")
	query = query.Order("users.user_fullname ASC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

// popup
func GetPopupShow(notInIds []int, objQuery *[]Popup) error {
	query := configs.DB1.Table("popup")
	query = query.Select("popup.*")
	query = query.Where("popup.popup_status_id = ?", 1)
	if len(notInIds) > 0 {
		query = query.Where("popup.popup_id NOT IN ?", notInIds)
	}
	query = query.Find(&objQuery)
	return query.Error
}

func GetPopupUser(user_id int, objQuery *[]PopupUser) error {
	query := configs.DB1.Table("popup_users")
	query = query.Select("popup_users.*")
	query = query.Where("popup_users.user_id = ?", user_id)
	query = query.Find(&objQuery)
	return query.Error
}

func CreatePopupUser(data *PopupUser) (err error) {
	query := configs.DB1.Table("popup_users").Create(&data)
	return query.Error
}

// tx user
func UpdateOrderUser(orderId int, data User) error {
	itf := map[string]interface{}{
		"user_id": data.ID,
	}
	query := configs.DB1.Table("orders")
	query = query.Where("orders.id = ?", orderId)
	query = query.Updates(itf)
	return query.Error
}

func UpdateTaxUser(taxId int, data User) error {
	itf := map[string]interface{}{
		"user_id": data.ID,
	}
	query := configs.DB1.Table("taxs")
	query = query.Where("taxs.id = ?", taxId)
	query = query.Updates(itf)
	return query.Error
}

func UpdateQueueUser(queueId int, data User) error {
	itf := map[string]interface{}{
		"user_id": data.ID,
	}
	query := configs.DB1.Table("queues")
	query = query.Where("queues.id = ?", queueId)
	query = query.Updates(itf)
	return query.Error
}

func UpdateOpdUser(opdId int, data User) error {
	itf := map[string]interface{}{
		"user_id": data.ID,
	}
	query := configs.DB1.Table("opds")
	query = query.Where("opds.id = ?", opdId)
	query = query.Updates(itf)
	return query.Error
}

func UpdateInvoiceUser(invoiceId int, data User) error {
	itf := map[string]interface{}{
		"user_id": data.ID,
	}
	query := configs.DB1.Table("invoices")
	query = query.Where("invoices.id = ?", invoiceId)
	query = query.Updates(itf)
	return query.Error
}

func UpdateReceiptUser(receiptId int, data User) error {
	itf := map[string]interface{}{
		"user_id": data.ID,
	}
	query := configs.DB1.Table("receipts")
	query = query.Where("receipts.id = ?", receiptId)
	query = query.Updates(itf)
	return query.Error
}

func GetUserEmail(user_id int, data *structs.ObjUserEmail) (err error) {
	query := configs.DB1.Table("users")
	query = query.Select("users.user_email")
	query = query.Where("users.id = ?", user_id)
	if err := query.Scan(&data).Error; err != nil {
		// configs.DB1.Close()
		return err
	}
	return nil
}
