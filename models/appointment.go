package models

import (
	"encoding/json"
	"linecrmapi/configs"
	"linecrmapi/structs"
)

func AppointmentSearch(filter structs.PayloadSearchAppointment, isCount bool, apList *[]structs.AppointmentList) (err error) {
	query := configs.DB1.Table("appointments")
	query = query.Select("appointments.*,customers.ctm_id,customers.ctm_fname,customers.ctm_lname,customers.ctm_fname_en,customers.ctm_lname_en,users.user_fullname,users.user_fullname_en,shop_roles.role_name_th,shop_roles.role_name_en")
	query = query.Joins("LEFT JOIN customers ON customers.id = appointments.customer_id")
	query = query.Joins("LEFT JOIN users ON users.id = appointments.user_id")
	query = query.Joins("JOIN user_shops ON user_shops.user_id = users.id")
	query = query.Joins("JOIN shop_roles ON shop_roles.id = user_shops.shop_role_id")
	query = query.Where("shop_roles.shop_id = ?", filter.Shop_id)
	query = query.Where("appointments.shop_id = ?", filter.Shop_id)
	query = query.Where("appointments.ap_is_del = 0")
	if *filter.Search != "" {
		query = query.Where("appointments.ap_topic LIKE '%" + *filter.Search + "%' OR customers.ctm_id LIKE '%" + *filter.Search + "%' OR appointments.customer_fullname LIKE '%" + *filter.Search + "%' OR appointments.ap_comment LIKE '%" + *filter.Search + "%' OR users.user_fullname LIKE '%" + *filter.Search + "%' OR users.user_fullname_en LIKE '%" + *filter.Search + "%' ")
	}
	if *filter.Date != "" {
		query = query.Where("appointments.ap_datetime LIKE '%" + *filter.Date + "%'")
	}
	if *filter.Type != "" {
		query = query.Where("appointments.ap_type = ?", *filter.Type)
	}
	if *filter.OpdType != "" {
		query = query.Where("appointments.ap_opd_type = ?", *filter.OpdType)
	}
	if *filter.Is_active != "" {
		query = query.Where("appointments.ap_status_id = ?", *filter.Is_active)
	}

	if isCount == true {
		offset := filter.ActivePage * filter.PerPage
		query = query.Limit(filter.PerPage)
		query = query.Offset(offset)
	}
	if *filter.Date != "" {
		query = query.Order("appointments.ap_datetime ASC")
	} else {
		query = query.Order("appointments.ap_datetime DESC")
		query = query.Order("appointments.id DESC")
	}
	if err = query.Scan(&apList).Error; err != nil {
		return err
	}
	return nil
}

func AppointmentSearchCalendar(customer_id int, shop_id int, filter structs.PayloadCalendarAppointment, data *[]structs.AppointmentList) (err error) {
	query := configs.DB1.Table("appointments")
	query = query.Select("appointments.*,customers.ctm_fname,customers.ctm_lname,customers.ctm_fname_en,customers.ctm_lname_en,users.user_fullname,users.user_fullname_en,shop_roles.role_name_th,shop_roles.role_name_en")
	query = query.Joins("LEFT JOIN customers ON customers.id = appointments.customer_id")
	query = query.Joins("JOIN users ON users.id = appointments.user_id")
	query = query.Joins("JOIN user_shops ON user_shops.user_id = users.id")
	query = query.Joins("JOIN shop_roles ON shop_roles.id = user_shops.shop_role_id")
	query = query.Where("shop_roles.shop_id = ?", shop_id)
	query = query.Where("appointments.shop_id = ?", shop_id)
	query = query.Where("appointments.customer_id = ?", customer_id)
	query = query.Where("appointments.ap_is_del = 0")
	if *filter.Date_start != "" {
		query = query.Where("appointments.ap_datetime >= ?", *filter.Date_start+" 00:00:00")
	}
	if *filter.Date_end != "" {
		query = query.Where("appointments.ap_datetime <= ?", *filter.Date_end+" 23:59:59")
	}
	query = query.Where("appointments.ap_status_id != ?", 4)
	query = query.Order("appointments.ap_datetime ASC")
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func CheckAppointmentDupicate(userID int, datetime string, ap *[]structs.AppointmentDetail) (err error) {
	query := configs.DB1.Table("appointments")
	query = query.Select("appointments.*,customers.ctm_id,customers.ctm_fname,customers.ctm_lname,customers.ctm_fname_en,customers.ctm_lname_en,users.user_fullname,users.user_fullname_en")
	query = query.Joins("LEFT JOIN customers ON customers.id = appointments.customer_id")
	query = query.Joins("JOIN users ON users.id = appointments.user_id")
	query = query.Where("appointments.user_id = ?", userID)
	query = query.Where("appointments.ap_datetime = ?", datetime)
	query = query.Where("appointments.ap_is_del = 0")
	query = query.Order("appointments.ap_datetime DESC")
	if err = query.Scan(&ap).Error; err != nil {
		return err
	}
	return nil
}

func CheckAppointmentDupicateUpdate(userID int, datetime string, apID int, ap *[]structs.AppointmentDetail) (err error) {
	query := configs.DB1.Table("appointments")
	query = query.Select("appointments.*,customers.ctm_id,customers.ctm_fname,customers.ctm_lname,customers.ctm_fname_en,customers.ctm_lname_en,users.user_fullname,users.user_fullname_en")
	query = query.Joins("LEFT JOIN customers ON customers.id = appointments.customer_id")
	query = query.Joins("JOIN users ON users.id = appointments.user_id")
	query = query.Where("appointments.user_id = ?", userID)
	query = query.Where("appointments.ap_datetime = ?", datetime)
	query = query.Where("appointments.id != ?", apID)
	query = query.Where("appointments.ap_is_del = 0")
	query = query.Order("appointments.ap_datetime DESC")
	if err = query.Scan(&ap).Error; err != nil {
		return err
	}
	return nil
}

func GetAppointmentDetail(apID int, shopId int, ap *structs.AppointmentDetail) (err error) {
	query := configs.DB1.Table("appointments")
	query = query.Select("appointments.*, customers.ctm_id, customers.ctm_tel, customers.ctm_tel_2, customers.ctm_birthdate, customers.ctm_fname,customers.ctm_lname,customers.ctm_fname_en,customers.ctm_lname_en,users.user_fullname,users.user_fullname_en,shop_roles.role_name_th,shop_roles.role_name_en")
	query = query.Joins("LEFT JOIN customers ON customers.id = appointments.customer_id")
	query = query.Joins("JOIN users ON users.id = appointments.user_id")
	query = query.Joins("JOIN user_shops ON user_shops.user_id = users.id")
	query = query.Joins("JOIN shop_roles ON shop_roles.id = user_shops.shop_role_id")
	query = query.Where("shop_roles.shop_id = ?", shopId)
	query = query.Where("appointments.id = ?", apID)
	query = query.Where("appointments.shop_id = ?", shopId)
	query = query.Where("appointments.ap_is_del = 0")
	if err = query.Scan(&ap).Error; err != nil {
		return err
	}
	return nil
}

func GetAppointmentDetailList(customerId int, shopId int, ap *[]structs.AppointmentDetail) (err error) {
	query := configs.DB1.Table("appointments")
	query = query.Select("appointments.*,customers.ctm_id, customers.ctm_tel, customers.ctm_tel_2, customers.ctm_birthdate,customers.ctm_fname,customers.ctm_lname,customers.ctm_fname_en,customers.ctm_lname_en,users.user_fullname,users.user_fullname_en")
	query = query.Joins("LEFT JOIN customers ON customers.id = appointments.customer_id")
	query = query.Joins("JOIN users ON users.id = appointments.user_id")
	query = query.Joins("JOIN user_shops ON user_shops.user_id = users.id")
	// query = query.Joins("JOIN shop_roles ON shop_roles.id = user_shops.shop_role_id")
	// query = query.Where("shop_roles.shop_id = ?", shopId)
	query = query.Where("appointments.customer_id = ?", customerId)
	query = query.Where("appointments.shop_id = ?", shopId)
	query = query.Where("user_shops.shop_id = ?", shopId)
	query = query.Where("appointments.ap_is_del = 0")
	query = query.Order("appointments.ap_datetime DESC")
	if err = query.Scan(&ap).Error; err != nil {
		return err
	}
	return nil
}

func GetAppointmentTagList(apID int, tgs *[]structs.AppointmentTags) (err error) {
	query := configs.DB1.Table("appointment_tags")
	query = query.Select("appointment_tags.*")
	query = query.Where("appointment_tags.appointment_id = ?", apID)
	query = query.Order("appointment_tags.tag_id ASC")
	if err = query.Scan(&tgs).Error; err != nil {
		return err
	}
	return nil
}

func GetAppointmentByID(apID int, shopId int, ap *structs.AppointmentAction) (err error) {
	query := configs.DB1.Table("appointments")
	query = query.Select("appointments.*")
	query = query.Where("appointments.id = ?", apID)
	query = query.Where("appointments.shop_id = ?", shopId)
	query = query.Where("appointments.ap_is_del = 0")

	if err = query.Scan(&ap).Error; err != nil {
		return err
	}
	return nil
}

func AddAppointment(dataAP *structs.AppointmentAction) (err error) {
	tx := configs.DB1.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	// add appointments
	if err = tx.Table("appointments").Create(&dataAP).Error; err != nil {
		tx.Rollback()
		return err
	}

	if dataAP.CustomerID > 0 {
		if err = tx.Table("appointments").Where("appointments.id = ?", dataAP.ID).Updates(map[string]interface{}{"appointments.customer_id": dataAP.CustomerID}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func UpdateAppointment(shop_id int, ap_id int, data *structs.AppointmentAction) (err error) {

	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	delete(inInterface, "id")
	delete(inInterface, "shop_id")
	delete(inInterface, "ap_type")
	delete(inInterface, "ap_user_id")
	delete(inInterface, "customer_id")
	delete(inInterface, "customer_fullname")
	delete(inInterface, "ap_create")

	query := configs.DB1.Table("appointments")
	query = query.Where("appointments.shop_id = ?", shop_id)
	query = query.Where("appointments.id = ?", ap_id)
	// query = query.Model(&data)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateAppointmentStatus(shop_id int, ap_id int, data *structs.AppointmentStatusAction) (err error) {

	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	delete(inInterface, "id")
	delete(inInterface, "user_id")
	delete(inInterface, "customer_id")
	delete(inInterface, "ap_opd_type")
	delete(inInterface, "ap_create")
	query := configs.DB1.Table("appointments")
	query = query.Where("appointments.shop_id = ?", shop_id)
	query = query.Where("appointments.id = ?", ap_id)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func DeleteAppointment(Id int, obj *structs.AppointmentAction) (err error) {
	query := configs.DB1.Table("appointments")
	query = query.Where("appointments.id = ?", Id)
	query = query.Model(&obj)
	query = query.Updates(map[string]interface{}{"appointments.ap_is_del": 1})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CreateAppointmentTagBatch(objCreate *[]structs.AppointmentTags) error {
	query := configs.DB1.CreateInBatches(&objCreate, 24)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func DeleteAppointmentTag(Id int, obj *[]structs.AppointmentTags) (err error) {
	query := configs.DB1.Table("appointment_tags")
	query = query.Where("appointment_tags.appointment_id = ?", Id)
	query = query.Delete(&obj)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetAppointmentDoctor(shopId int, objQuery *[]structs.ObjResponseSearchAppointmentDoctor) error {
	query := configs.DB1.Table("user_shops")
	query = query.Select("users.id, users.user_email, users.user_fullname, users.user_fullname_en,shop_roles.role_name_th,shop_roles.role_name_en")
	query = query.Joins("JOIN users ON users.id = user_shops.user_id")
	query = query.Joins("JOIN shop_roles ON shop_roles.id = user_shops.shop_role_id")
	query = query.Where("user_shops.shop_id = ?", shopId)
	query = query.Where("shop_roles.role_id IN ?", []int{1, 2, 6, 8})
	query = query.Where("users.user_is_active = ?", 1)
	query = query.Order("users.user_fullname ASC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetAppointmentCustomer(shopId int, ac *[]structs.AppointmentCustomer) (err error) {
	query := configs.DB1.Table("customers")
	query = query.Select("customers.*")
	query = query.Where("customers.shop_id = ?", shopId)
	query = query.Where("customers.ctm_is_active = 1")
	query = query.Where("customers.ctm_is_del = 0")
	query = query.Order("customers.ctm_fname ASC")
	if err = query.Scan(&ac).Error; err != nil {
		return err
	}
	return nil
}

func GetAppointmentUser(shopId int, up *[]structs.AppointmentUser) (err error) {
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

func GetAppointmentTopic(shopId int, at *[]structs.AppointmentTopic) (err error) {
	query := configs.DB1.Table("topics")
	query = query.Select("topics.*")
	query = query.Where("topics.shop_id = ?", shopId)
	query = query.Where("topics.topic_is_del = 0")
	query = query.Order("topics.topic_th ASC")
	if err = query.Scan(&at).Error; err != nil {
		return err
	}
	return nil
}

func GetAppointmentTag(shopId int, at *[]structs.AppointmentTag) (err error) {
	query := configs.DB1.Table("tags")
	query = query.Select("tags.*")
	query = query.Where("tags.shop_id = ?", shopId)
	query = query.Where("tags.tag_type_id = 9")
	query = query.Where("tags.tag_is_del = 0")
	query = query.Order("tags.tag_name ASC")
	if err = query.Scan(&at).Error; err != nil {
		return err
	}
	return nil
}

func GetAppointmentTagById(appointmentId int, objQuery *[]structs.ObjQueryCustomerTag) error {
	query := configs.DB1.Table("appointment_tags")
	query = query.Select("appointment_tags.*, tags.tag_name, tag_types.tag_type_th, tag_types.tag_type_en")
	query = query.Joins("LEFT JOIN tags on tags.id = appointment_tags.tag_id")
	query = query.Joins("LEFT JOIN tag_types on tag_types.id = tags.tag_type_id")
	query = query.Where("customer_tags.appointment_id = ?", appointmentId)
	query = query.Where("tags.tag_type_id = ?", 9)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func AddLogAppointment(log *structs.LogAppointment) (err error) {
	query := configs.DBL1.Table("log_appointment").Create(&log)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetTimesetConfig(shopId int, data *structs.ShopTimeset) (err error) {
	query := configs.DB1.Table("timesets")
	query = query.Select("timesets.*")
	query = query.Where("timesets.shop_id = ?", shopId)
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetUserSettime(userId int, dayId int, data *[]structs.UserSettimeList) (err error) {
	query := configs.DB1.Table("user_settime")
	query = query.Select("user_settime.*")
	query = query.Where("user_settime.user_id = ?", userId)
	query = query.Where("user_settime.day_id = ?", dayId)
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetUserTimesetDoctor(shopId int, dateStart string, dateEnd string, data *[]structs.UserSettimeToDay) (err error) {
	query := configs.DB1.Table("user_settime")
	query = query.Select("user_settime.id, shop_roles.role_name_th, shop_roles.role_name_en, user_shops.shop_id,user_settime.user_id,users.user_fullname, users.user_fullname_en, user_settime.day_id,user_settime.time_start,user_settime.time_end,users.user_image,days.day_name_th,days.day_name_en")
	query = query.Joins("INNER JOIN users ON user_settime.user_id = users.id")
	query = query.Joins("INNER JOIN user_shops ON user_shops.user_id = users.id")
	query = query.Joins("INNER JOIN shop_roles ON user_shops.shop_role_id = shop_roles.id")
	query = query.Joins("INNER JOIN days ON user_settime.day_id = days.id")
	query = query.Where("user_settime.shop_id = ?", shopId)
	query = query.Where("user_shops.shop_id = ?", shopId)
	// query = query.Where("user_settime.time_start >= ?", dateStart)
	// query = query.Where("user_settime.time_end <= ?", dateEnd)
	query = query.Where("shop_roles.role_id IN (1, 2, 6, 8)")
	query = query.Order("users.id ASC")
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetAppointmentByDate(shopId int, dateStart string, dateEnd string, data *[]structs.AppointmentList) (err error) {
	query := configs.DB1.Table("appointments")
	query = query.Select("appointments.*, users.*")
	query = query.Joins("JOIN users ON users.id = appointments.user_id")
	query = query.Where("appointments.shop_id = ?", shopId)
	query = query.Where("appointments.ap_is_del = 0")
	query = query.Where("appointments.ap_datetime >= ?", dateStart)
	query = query.Where("appointments.ap_datetime <= ?", dateEnd)
	query = query.Where("appointments.ap_status_id = ?", 1)
	query = query.Order("appointments.ap_datetime ASC")
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetUserTimesetByDay(shopId int, day int, dateStart string, dateEnd string, data *[]structs.UserSettimeToDay) (err error) {
	query := configs.DB1.Table("user_settime")
	query = query.Select("user_settime.id, shop_roles.role_name_th, shop_roles.role_name_en, user_shops.shop_id, user_settime.user_id, users.user_fullname_en, users.user_fullname, user_settime.day_id, user_settime.time_start, user_settime.time_end, users.user_image, days.day_name_th, days.day_name_en")
	query = query.Joins("INNER JOIN users ON user_settime.user_id = users.id")
	query = query.Joins("INNER JOIN user_shops ON user_shops.user_id = users.id")
	query = query.Joins("INNER JOIN shop_roles ON user_shops.shop_role_id = shop_roles.id")
	query = query.Joins("INNER JOIN days ON user_settime.day_id = days.id")
	query = query.Where("user_settime.shop_id = ?", shopId)
	query = query.Where("user_shops.shop_id = ?", shopId)
	// query = query.Where("user_settime.time_start >= ?", dateStart)
	// query = query.Where("user_settime.time_end <= ?", dateEnd)
	query = query.Where("user_settime.day_id = ?", day)
	query = query.Where("shop_roles.role_id IN ?", []int{1, 2, 6, 8})
	query = query.Order("users.id ASC")
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetEmailByAppoinment(apId int, data *structs.AppointmentEmail) (err error) {
	query := configs.DB1.Table("appointments")
	query = query.Select("appointments.id, appointments.user_id, appointments.customer_id, appointments.shop_id, shops.shop_email,shops.shop_gid,shops.shop_gc_token, users.user_email, customers.ctm_email")
	query = query.Joins("JOIN users ON users.id = appointments.user_id")
	query = query.Joins("LEFT JOIN customers ON appointments.customer_id = customers.id")
	query = query.Joins("JOIN shops ON appointments.shop_id = shops.id")
	query = query.Where("appointments.id = ?", apId)
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetShopHoliday(shopId int, date string, data *[]structs.ShopHoliday) (err error) {
	query := configs.DB1.Table("holidays")
	query = query.Select("holidays.*")
	query = query.Where("holidays.shop_id = ?", shopId)
	query = query.Where("holidays.holiday_date = ?", date)
	query = query.Where("holidays.holiday_is_del = ?", 0)
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}
