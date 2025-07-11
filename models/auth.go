package models

import (
	"encoding/json"
	"linecrmapi/configs"
	"linecrmapi/structs"
	"time"
)

func GetCustomerOnlineByEmail(co_email string, result *structs.CustomerOnline) (err error) {
	query := configs.DB1.Table("customer_onlines")
    query = query.Select("customer_onlines.*")
    query = query.Where("customer_onlines.co_email = ?", co_email)
    query = query.Where("customer_onlines.co_is_active = 1")
    query = query.Where("customer_onlines.co_is_del = 0")
    query = query.Where("customer_onlines.co_otp = ?", "")
    query = query.Order("customer_onlines.id DESC")
    query = query.Limit(1)
    if err := query.Scan(&result).Error; err != nil {
        return err
    }
    return nil
}

func GetCustomerOnlineByLineId(line_id string, result *structs.CustomerOnline) (err error) {
	query := configs.DB1.Table("customer_onlines")
	query = query.Select("customer_onlines.*")
	query = query.Where("customer_onlines.co_line_id = ?", line_id)
	query = query.Where("customer_onlines.co_is_active = 1")
	query = query.Where("customer_onlines.co_is_del = 0")
	query = query.Where("customer_onlines.co_otp = ?", "")
	query = query.Order("customer_onlines.id DESC")
	query = query.Limit(1)
	if err := query.Scan(&result).Error; err != nil {
		return err
	}
	return nil
}

func GetCustomerOnlineEmailPassword(co_email string, co_password string, result *structs.CustomerOnline) (err error) {
	query := configs.DB1.Table("customer_onlines")
	query = query.Select("customer_onlines.*")
	query = query.Where("customer_onlines.co_email = ?", co_email)
	query = query.Where("customer_onlines.co_password = ?", co_password)
	query = query.Where("customer_onlines.co_is_active = 1")
	query = query.Where("customer_onlines.co_is_del = 0")
	// query = query.Where("customer_onlines.co_otp = ?", "")
	query = query.Order("customer_onlines.id DESC")
	query = query.Limit(1)
	if err := query.Scan(&result).Error; err != nil {
		return err
	}
	return nil
}

func GetCustomerOnlineTelCheck(co_tel string, result *structs.CustomerOnline) (err error) {
	query := configs.DB1.Table("customer_onlines")
	query = query.Select("customer_onlines.*")
	query = query.Where("customer_onlines.co_tel = ?", co_tel)
	query = query.Where("customer_onlines.co_is_active = 1")
	query = query.Where("customer_onlines.co_is_del = 0")
	// query = query.Where("customer_onlines.co_otp = ?", "")
	query = query.Order("customer_onlines.id DESC")
	query = query.Limit(1)
	if err := query.Scan(&result).Error; err != nil {
		return err
	}
	return nil
}

func CheckCustomerOnlineLogin(line_id string, citizen string, tel string, result *structs.CustomerOnline) (err error) {
	query := configs.DB1.Table("customer_onlines")
	query = query.Select("customer_onlines.*")
	query = query.Where("customer_onlines.co_line_id = ?", line_id)
	query = query.Where("customer_onlines.co_citizen_id = ?", citizen)
	query = query.Where("customer_onlines.co_tel = ?", tel)
	query = query.Where("customer_onlines.co_is_active = 1")
	query = query.Where("customer_onlines.co_is_del = 0")
	query = query.Where("customer_onlines.co_otp != ?", "")
	query = query.Order("customer_onlines.id DESC")
	query = query.Limit(1)
	if err := query.Scan(&result).Error; err != nil {
		return err
	}
	return nil
}

func GetCustomerOnlineByOtp(line_id string, tel string, otp string, result *structs.CustomerOnline) (err error) {
	query := configs.DB1.Table("customer_onlines")
	query = query.Select("customer_onlines.*")
	query = query.Where("customer_onlines.co_line_id = ?", line_id)
	query = query.Where("customer_onlines.co_tel = ?", tel)
	query = query.Where("customer_onlines.co_otp = ?", otp)
	query = query.Where("customer_onlines.co_is_active = 1")
	query = query.Where("customer_onlines.co_is_del = 0")
	query = query.Where("customer_onlines.co_otp_expire > ?", time.Now().Format("2006-01-02 15:04:05"))
	query = query.Order("customer_onlines.id DESC")
	query = query.Limit(1)
	if err := query.Scan(&result).Error; err != nil {
		return err
	}
	return nil
}

func GetCustomerFirstByCitizenId(citizen_id string, result *Customer) (err error) {
	query := configs.DB1.Table("customers")
	query = query.Select("customers.*")
	query = query.Joins("INNER JOIN shops ON shops.id = customers.shop_id")
	query = query.Where("customers.ctm_citizen_id = ?", citizen_id)
	query = query.Where("customers.ctm_is_active = 1")
	query = query.Where("customers.ctm_is_del = 0")
	query = query.Where("shops.shop_status_id = 1")
	query = query.Where("shops.shop_expire > ?", time.Now().Format("2006-01-02"))
	query = query.Order("shops.id ASC")
	query = query.Limit(1)
	if err := query.Scan(&result).Error; err != nil {
		return err
	}
	return nil
}

func AddCustomerOnline(data *structs.CustomerOnline) (err error) {
	query := configs.DB1.Table("customer_onlines").Create(&data)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}
func SyncAccountWithEmail(co_id int, data *structs.SyncAccountWithEmail) (err error) {
	var inInterface map[string]interface{}
    in, _ := json.Marshal(&data)
    json.Unmarshal(in, &inInterface)
    query := configs.DB1.Table("customer_onlines")
    query = query.Where("customer_onlines.id = ?", co_id)
    query = query.Updates(&inInterface)
    if err = query.Error; err != nil {
        return err
    }
    return nil
}

func UpdateCustomerOnline(co_id int, data *structs.CustomerOnlineUpdate) (err error) {
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query := configs.DB1.Table("customer_onlines")
	query = query.Where("customer_onlines.id = ?", co_id)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}
func UpdateOtpCustomerOnline(co_id int, data *structs.CustomerOnlineOtpUpdate) (err error) {
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query := configs.DB1.Table("customer_onlines")
	query = query.Where("customer_onlines.id = ?", co_id)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateCustomerOnlineSync(co_id int, data *structs.CustomerOnlineUpdateSync) (err error) {
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query := configs.DB1.Table("customer_onlines")
	query = query.Where("customer_onlines.id = ?", co_id)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateCustomerOnlineTel(co_id int, data *structs.CustomerOnlineUpdateTel) (err error) {
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query := configs.DB1.Table("customer_onlines")
	query = query.Where("customer_onlines.id = ?", co_id)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CreateLogOauth(objCreate *structs.ObjLogCustomerLogin) error {
	query := configs.DBL1.Table("log_customer_login").Create(&objCreate)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetCustomerOnlineByLineId2(line_id string, result *structs.CustomerOnline) (err error) {
	query := configs.DB1.Table("customer_onlines")
	query = query.Select("customer_onlines.*")
	query = query.Where("customer_onlines.co_line_id = ?", line_id)
	query = query.Where("customer_onlines.co_is_active = 1")
	query = query.Where("customer_onlines.co_is_del = 0")
	query = query.Order("customer_onlines.id DESC")
	query = query.Limit(1)
	if err := query.Scan(&result).Error; err != nil {
		return err
	}
	return nil
}

func DeleteCustomerOnline(co_id int) (err error) {
	inInterface := make(map[string]interface{})
	inInterface["co_is_del"] = 1
	inInterface["co_update"] = time.Now().Format("2006-01-02 15:04:05")
	query := configs.DB1.Table("customer_onlines")
	query = query.Where("customer_onlines.id = ?", co_id)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetCustomerOnlineById(co_id int, result *structs.CustomerOnline) (err error) {
	query := configs.DB1.Table("customer_onlines")
	query = query.Select("customer_onlines.*")
	query = query.Where("customer_onlines.id = ?", co_id)
	query = query.Where("customer_onlines.co_is_active = 1")
	query = query.Where("customer_onlines.co_is_del = 0")
	// query = query.Where("customer_onlines.co_otp = ?", "")
	query = query.Order("customer_onlines.id DESC")
	query = query.Limit(1)
	if err := query.Scan(&result).Error; err != nil {
		return err
	}
	return nil
}

func CheckMail(co_email string, result *[]structs.CustomerOnline) (err error) {
	query := configs.DB1.Table("customer_onlines")
	query = query.Where("co_email = ?", co_email)
	query = query.Where("co_is_del = 0")
	if err := query.Scan(&result).Error; err != nil {
		return err
	}
	return nil
}

func CheckCitizen(co_citizen_id string, result *[]structs.CustomerOnline) (err error) {
	query := configs.DB1.Table("customer_onlines")
	query = query.Where("co_citizen_id = ?", co_citizen_id)
	query = query.Where("co_is_del = 0")
	if err := query.Scan(&result).Error; err != nil {
		return err
	}
	return nil
}

func CheckTel(co_tel string, result *[]structs.CustomerOnline) (err error) {
	query := configs.DB1.Table("customer_onlines")
	query = query.Where("co_tel = ?", co_tel)
	query = query.Where("co_is_del = 0")
	if err := query.Scan(&result).Error; err != nil {
		return err
	}
	return nil
}

func GetCustomerOnlineByOtpRegister(co_email string, co_tel string, otp string, result *structs.CustomerOnline) (err error) {
	query := configs.DB1.Table("customer_onlines")
	query = query.Select("customer_onlines.*")
	query = query.Where("customer_onlines.co_email = ?", co_email)
	query = query.Where("customer_onlines.co_tel = ?", co_tel)
	query = query.Where("customer_onlines.co_otp = ?", otp)
	query = query.Where("customer_onlines.co_is_active = 1")
	query = query.Where("customer_onlines.co_is_del = 0")
	query = query.Where("customer_onlines.co_otp_expire > ?", time.Now().Format("2006-01-02 15:04:05"))
	query = query.Order("customer_onlines.id DESC")
	query = query.Limit(1)
	if err := query.Scan(&result).Error; err != nil {
		return err
	}
	return nil
}

func GetCustomerOnlineByOtpTelOnlyRegister(co_tel string, otp string, result *structs.CustomerOnline) (err error) {
	query := configs.DB1.Table("customer_onlines")
	query = query.Select("customer_onlines.*")
	query = query.Where("customer_onlines.co_tel = ?", co_tel)
	query = query.Where("customer_onlines.co_otp = ?", otp)
	query = query.Where("customer_onlines.co_is_active = 1")
	query = query.Where("customer_onlines.co_is_del = 0")
	query = query.Where("customer_onlines.co_otp_expire > ?", time.Now().Format("2006-01-02 15:04:05"))
	query = query.Order("customer_onlines.id DESC")
	query = query.Limit(1)
	if err := query.Scan(&result).Error; err != nil {
		return err
	}
	return nil
}

func GetCustomerOnlineByOtpSync(co_id int, otp string, result *structs.CustomerOnline) (err error) {
	query := configs.DB1.Table("customer_onlines")
	query = query.Select("customer_onlines.*")
	query = query.Where("customer_onlines.id = ?", co_id)
	query = query.Where("customer_onlines.co_otp = ?", otp)
	query = query.Where("customer_onlines.co_is_active = 1")
	query = query.Where("customer_onlines.co_is_del = 0")
	query = query.Where("customer_onlines.co_otp_expire > ?", time.Now().Format("2006-01-02 15:04:05"))
	query = query.Order("customer_onlines.id DESC")
	query = query.Limit(1)
	if err := query.Scan(&result).Error; err != nil {
		return err
	}
	return nil
}
func GetCustomerOnlineByEmailOtpKeyRegister(co_email string, co_otp_key string, otp string, result *structs.CustomerOnline) (err error) {
	query := configs.DB1.Table("customer_onlines")
	query = query.Select("customer_onlines.*")
	query = query.Where("customer_onlines.co_email = ?", co_email)
	query = query.Where("customer_onlines.co_otp_key = ?", co_otp_key)
	query = query.Where("customer_onlines.co_otp = ?", otp)
	query = query.Where("customer_onlines.co_is_active = 1")
	query = query.Where("customer_onlines.co_is_del = 0")
	query = query.Where("customer_onlines.co_otp_expire > ?", time.Now().Format("2006-01-02 15:04:05"))
	query = query.Order("customer_onlines.id DESC")
	query = query.Limit(1)
	if err := query.Scan(&result).Error; err != nil {
		return err
	}
	return nil
}

func GetCustomerOnlineByCitizenId (citizen_id string, result *structs.CustomerOnline) (err error) {
	query := configs.DB1.Table("customer_onlines")
    query = query.Select("customer_onlines.*")
    query = query.Where("customer_onlines.co_citizen_id =?", citizen_id)
    query = query.Where("customer_onlines.co_is_active = 1")
    query = query.Where("customer_onlines.co_is_del = 0")
    query = query.Order("customer_onlines.id DESC")
    query = query.Limit(1)
    if err := query.Scan(&result).Error; err != nil {
        return err
    }
    return nil
}
func GetCustomerLockShop(citizen_id string, shop_id int, result *Customer) (err error) {
	query := configs.DB1.Table("customers")
	query = query.Select("customers.*")
	query = query.Joins("INNER JOIN shops ON shops.id = customers.shop_id")
	query = query.Where("customers.ctm_citizen_id = ?", citizen_id)
	query = query.Where("customers.ctm_is_active = 1")
	query = query.Where("customers.ctm_is_del = 0")
	query = query.Where("shops.shop_status_id = 1")
	query = query.Where("shops.id = ?", shop_id)
	query = query.Where("shops.shop_expire > ?", time.Now().Format("2006-01-02"))
	query = query.Order("shops.id ASC")
	query = query.Limit(1)
	if err := query.Scan(&result).Error; err != nil {
		return err
	}
	return nil
}

func TestRemoveAccountWithLine(line_id string, data *structs.TestRemoveAccountWithLineId) (err error) {

 var inInterface map[string]interface{}
 in, _ := json.Marshal(&data)
 json.Unmarshal(in, &inInterface)
 query := configs.DB1.Table("customer_onlines")
 query = query.Where("customer_onlines.co_line_id = ?", line_id)
 query = query.Updates(&inInterface)

 if err = query.Error; err != nil {
  return err
 }
 return nil
}
