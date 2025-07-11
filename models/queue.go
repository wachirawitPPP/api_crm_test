package models

import (
	"encoding/json"
	"linecrmapi/configs"
	"linecrmapi/structs"

	"time"
)

func CountQueuePagination(objPayload *structs.ObjPayloadGetQueuePagination) (int64, error) {
	var count int64
	query := configs.DB1.Table("queues")
	query = query.Select("queues.id")
	query = query.Joins("LEFT JOIN customers ON customers.id = queues.customer_id")
	query = query.Joins("LEFT JOIN users ON users.id = queues.user_id")
	query = query.Joins("LEFT JOIN rooms ON rooms.id = queues.room_id")
	query = query.Joins("LEFT JOIN beds ON beds.id = queues.bed_id")
	query = query.Joins("LEFT JOIN invoices ON invoices.queue_id = queues.id AND invoices.inv_is_active != 0")
	query = query.Joins("LEFT JOIN shops ON customers.shop_id = shops.id")
	query = query.Where("queues.shop_id = ?", objPayload.ShopId)
	if *objPayload.QueAdmisId != -1 {
		query = query.Where("queues.que_admis_id = ?", objPayload.QueAdmisId)
	}
	if *objPayload.QueTypeId != -1 {
		query = query.Where("queues.que_type_id = ?", objPayload.QueTypeId)
	}
	if *objPayload.QueStatusId != -1 {
		query = query.Where("queues.que_status_id = ?", objPayload.QueStatusId)
	}
	if *objPayload.DoctorId != -1 {
		query = query.Where("queues.que_user_id = ?", objPayload.DoctorId)
	}
	if *objPayload.QueDate != "" {
		query = query.Where("DATE(queues.que_datetime) = ?", objPayload.QueDate)
	}
	if *objPayload.SearchText != "" {
		query = query.Where("queues.que_code LIKE '%" + *objPayload.SearchText + "%' OR queues.ctm_fullname LIKE '%" + *objPayload.SearchText + "%' OR queues.ctm_fullname_en LIKE '%" + *objPayload.SearchText + "%' OR CONCAT(customers.ctm_fname,' ',customers.ctm_lname) LIKE '%" + *objPayload.SearchText + "%' OR customers.ctm_id LIKE '%" + *objPayload.SearchText + "%'")
	}
	query = query.Count(&count)
	if query.Error != nil {
		return 0, query.Error
	}
	return count, nil
}

func GetQueuePagination(objPayload *structs.ObjPayloadGetQueuePagination, objQuery *[]structs.ObjQueryGetQueuePagination) error {
	query := configs.DB1.Table("queues")
	query = query.Select("queues.*, customers.shop_id AS ctm_shop_id, shops.shop_name AS ctm_shop_name, customers.ctm_id, customers.ctm_image, customers.ctm_prefix, customers.ctm_gender, customers.ctm_fname, customers.ctm_lname, customers.ctm_fname_en, customers.ctm_lname_en, rooms.room_code, rooms.room_th, rooms.room_en, beds.bed_code, invoices.id AS invoice_id")
	query = query.Joins("LEFT JOIN customers ON customers.id = queues.customer_id")
	query = query.Joins("LEFT JOIN users ON users.id = queues.user_id")
	query = query.Joins("LEFT JOIN rooms ON rooms.id = queues.room_id")
	query = query.Joins("LEFT JOIN beds ON beds.id = queues.bed_id")
	query = query.Joins("LEFT JOIN invoices ON invoices.queue_id = queues.id AND invoices.inv_is_active != 0")
	query = query.Joins("LEFT JOIN shops ON customers.shop_id = shops.id")
	query = query.Where("queues.shop_id = ?", objPayload.ShopId)
	if *objPayload.QueAdmisId != -1 {
		query = query.Where("queues.que_admis_id = ?", objPayload.QueAdmisId)
	}
	if *objPayload.QueTypeId != -1 {
		query = query.Where("queues.que_type_id = ?", objPayload.QueTypeId)
	}
	if *objPayload.QueStatusId != -1 {
		query = query.Where("queues.que_status_id = ?", objPayload.QueStatusId)
	}
	if *objPayload.DoctorId != -1 {
		query = query.Where("queues.que_user_id = ?", objPayload.DoctorId)
	}
	if *objPayload.QueDate != "" {
		query = query.Where("DATE(queues.que_datetime) = ?", objPayload.QueDate)
	}
	if *objPayload.SearchText != "" {
		query = query.Where("queues.que_code LIKE '%" + *objPayload.SearchText + "%' OR queues.ctm_fullname LIKE '%" + *objPayload.SearchText + "%' OR queues.ctm_fullname_en LIKE '%" + *objPayload.SearchText + "%' OR CONCAT(customers.ctm_fname,' ',customers.ctm_lname) LIKE '%" + *objPayload.SearchText + "%' OR customers.ctm_id LIKE '%" + *objPayload.SearchText + "%'")
	}
	// <pagination
	offset := objPayload.CurrentPage * objPayload.PerPage
	query = query.Limit(objPayload.PerPage)
	query = query.Offset(offset)
	// pagination>
	query = query.Order("queues.que_datetime DESC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueById(queueId int, objQuery *structs.ObjQueryGetQueue) error {
	query := configs.DB1.Table("queues")
	query = query.Select("queues.*, customers.ctm_id, customers.ctm_gender, customers.ctm_prefix, customers.ctm_fname, customers.ctm_lname, customers.ctm_fname_en, customers.ctm_lname_en, customers.ctm_image, customers.ctm_birthdate, customers.ctm_blood, customers.ctm_weight, customers.ctm_height, customers.ctm_waistline, customers.ctm_chest, customers.ctm_treatment_type, customers.ctm_health_comment, customers.ctm_disease, customers.ctm_mental_health, customers.ctm_allergic, users.user_fullname, ref_right_treatments.rt_code, ref_right_treatments.rt_name, rooms.room_code, rooms.room_en, rooms.room_th, beds.bed_code")
	query = query.Joins("LEFT JOIN customers ON customers.id = queues.customer_id")
	query = query.Joins("LEFT JOIN users ON users.id = queues.user_id")
	query = query.Joins("LEFT JOIN ref_right_treatments on ref_right_treatments.id = customers.right_treatment_id")
	query = query.Joins("LEFT JOIN rooms ON rooms.id = queues.room_id")
	query = query.Joins("LEFT JOIN beds ON beds.id = queues.bed_id")
	query = query.Where("queues.id = ?", queueId)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueLabel(shopId int, objQuery *[]structs.ObjQueryQueueLabel) error {

	today := time.Now().Truncate(24 * time.Hour)

	query := configs.DB1.Table("queues")
	query = query.Select("queues.*")
	query = query.Where("queues.que_status_id < ?", 4)
	query = query.Where("queues.shop_id = ?", shopId)
	query = query.Where("DATE(que_datetime) = ?", today.Format("2006-01-02"))
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueHistory(customerId int, queueId int, queueTypeId int, objQuery *[]structs.ObjQueryQueueHistory) error {
	query := configs.DB1.Table("queues")
	query = query.Select("queues.id AS que_id, queues.que_code, queues.que_datetime, users.user_fullname, users.user_fullname_en, queues.que_admis_id")
	query = query.Joins("LEFT JOIN users ON users.id = queues.user_id")
	query = query.Where("queues.customer_id = ?", customerId)
	query = query.Where("queues.id != ?", queueId)
	query = query.Where("queues.que_type_id = ?", queueTypeId)
	query = query.Limit(10)
	query = query.Order("queues.que_datetime DESC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func SearchQueueCustomer(objPayload *structs.ObjPayloadSearchQueueCustomer, objQuery *[]structs.ObjResponseSearchQueueCustomer) error {
	query := configs.DB1.Table("customers")
	query = query.Select("customers.id, customers.ctm_id, customers.ctm_fname, customers.ctm_lname, customers.ctm_fname_en, customers.ctm_lname_en")
	query = query.Where("customers.shop_mother_id = ?", objPayload.ShopId)
	// query = query.Where("customers.shop_id IN ?", objPayload.ShopIds)
	query = query.Where("customers.ctm_is_active = ?", 1)
	query = query.Where("customers.ctm_id LIKE ? OR customers.ctm_fname LIKE ? OR customers.ctm_lname LIKE ? OR customers.ctm_nname LIKE ? OR customers.ctm_nname LIKE ? OR customers.ctm_nname LIKE ? OR customers.ctm_fname_en LIKE ? OR customers.ctm_lname_en LIKE ? OR customers.ctm_citizen_id LIKE ? OR customers.ctm_passport_id LIKE ? OR CONCAT(customers.ctm_fname,' ',customers.ctm_lname) LIKE ?", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%")
	query = query.Order("customers.ctm_fname ASC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func SearchQueueDoctor(objPayload *structs.ObjPayloadSearchQueueDoctor, objQuery *[]structs.ObjResponseSearchQueueDoctor) error {
	query := configs.DB1.Table("user_shops")
	query = query.Select("users.id, users.user_email, users.user_fullname, users.user_fullname_en")
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

func SearchQueueDiagnostic(objPayload *structs.ObjPayloadSearchQueueDiagnostic, objQuery *[]structs.ObjResponseSearchQueueDiagnostic) error {
	query := configs.DB1.Table("diagnostics")
	query = query.Select("*")
	query = query.Where("diagnostics.shop_id = ?", objPayload.ShopId)
	query = query.Where("diagnostics.diagnostic_is_active = ?", 1)
	query = query.Where("diagnostics.diagnostic_is_del = ?", 0)
	query = query.Where("diagnostics.diagnostic_code LIKE ? OR diagnostics.diagnostic_th LIKE ? OR diagnostics.diagnostic_en LIKE ?", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%")
	query = query.Order("diagnostics.diagnostic_code ASC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueDocSetting(shopId int, data *structs.ObjQueryQueueDocSetting) (err error) {
	query := configs.DB1.Table("doc_settings")
	query = query.Select("shop_id, ipd_id_default, ipd_number_default, ipd_number_digit, ipd_type, opd_id_default, opd_number_default, opd_number_digit, opd_type, serve_id_default, serve_number_default, serve_number_digit, serve_type, cert_id_default, cert_number_default, cert_number_digit, cert_type, phrf_id_default, phrf_number_default, phrf_number_digit, phrf_type, sick_id_default, sick_number_default, sick_number_digit, sick_type")
	query = query.Where("shop_id = ?", shopId)
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func UpdateQueueDocSetting(shopId int, docKey int, docValue int, data *DocSetting) (err error) {
	query := configs.DB1.Table("doc_settings")
	query = query.Where("shop_id = ?", shopId)
	query = query.Model(&data)
	if docKey == 1 {
		query = query.Updates(map[string]interface{}{"ipd_number_default": docValue})
	} else if docKey == 2 {
		query = query.Updates(map[string]interface{}{"opd_number_default": docValue})
	} else if docKey == 3 {
		query = query.Updates(map[string]interface{}{"serve_number_default": docValue})
	} else if docKey == 4 {
		query = query.Updates(map[string]interface{}{"cert_number_default": docValue})
	} else if docKey == 5 {
		query = query.Updates(map[string]interface{}{"phrf_number_default": docValue})
	} else if docKey == 6 {
		query = query.Updates(map[string]interface{}{"sick_number_default": docValue})
	}
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CreateQueue(objQuery *Queue) (err error) {
	query := configs.DB1.Table("queues").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateQueue(queueId int, objQuery *Queue) (err error) {
	query := configs.DB1.Table("queues")
	query = query.Where("id = ?", queueId)
	query = query.Updates(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func ClearQueue(objQuery *Queue) (err error) {

	today := time.Now().Truncate(24 * time.Hour)

	query := configs.DB1.Table("queues")
	query = query.Where("que_admis_id = ?", 2)  // OPD
	query = query.Where("que_status_id = ?", 1) // รอ
	query = query.Where("DATE(que_datetime) < ?", today.Format("2006-01-02"))
	query = query.Updates(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetOpd(queueId int, objQuery *[]Opd) error {
	query := configs.DB1.Preload("OpdCustoms")
	query = query.Preload("OpdDiagnostics")
	query = query.Preload("User")
	query = query.Where("opds.opd_is_del = ?", 0)
	query = query.Where("opds.queue_id = ?", queueId)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetOpdById(opdId int, objQuery *Opd) error {
	query := configs.DB1.Preload("OpdCustoms")
	query = query.Preload("OpdDiagnostics")
	query = query.Preload("User")
	query = query.Where("opds.opd_is_del = ?", 0)
	query = query.Where("opds.id = ?", opdId)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
func GetOpdByIds(opdIds []int, objQuery *Opd) error {
	query := configs.DB1.Preload("OpdCustoms")
	query = query.Preload("OpdDiagnostics")
	query = query.Preload("User")
	query = query.Where("opds.opd_is_del = ?", 0)
	query = query.Where("opds.id IN ?", opdIds)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetOpdHistory(customerId int, queueId int, queueTypeId int, objQuery *[]structs.ObjQueryOpdHistory) error {
	query := configs.DB1.Table("opds")
	query = query.Select("opds.id AS opd_id, opds.opd_code, opds.opd_date, opds.opd_is_data, opds.opd_create, queues.id AS que_id, queues.que_code, queues.que_datetime, queues.que_admis_id, users.id AS user_id, users.user_fullname, users.user_fullname_en")
	query = query.Joins("LEFT JOIN queues ON queues.id = opds.queue_id")
	query = query.Joins("LEFT JOIN users ON users.id = opds.user_id")
	query = query.Where("queues.customer_id = ?", customerId)
	query = query.Where("queues.id != ?", queueId)
	query = query.Where("queues.que_type_id = ?", queueTypeId)
	query = query.Limit(10)
	query = query.Order("opds.opd_create DESC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CountOpdHistoryPagination(objPayload *structs.ObjPayloadGetOpdHistoryPagination) (int64, error) {
	var count int64
	query := configs.DB1.Table("opds")
	query = query.Select("opds.id")
	query = query.Joins("LEFT JOIN queues ON queues.id = opds.queue_id")
	query = query.Joins("LEFT JOIN shops ON shops.id = queues.shop_id")
	query = query.Where("opds.opd_is_del = ?", 0)
	query = query.Where("opds.customer_id = ?", objPayload.CustomerId)
	query = query.Where("queues.id != ?", objPayload.QueueId)
	query = query.Where("shops.shop_mother_id = ?", objPayload.ShopMotherId)
	if objPayload.SearchDate != "" {
		query = query.Where("opds.opd_date = ?", objPayload.SearchDate)
	}
	if objPayload.SearchText != "" {
		query = query.Where("opds.opd_code LIKE '%" + objPayload.SearchText + "%'")
	}
	query = query.Count(&count)
	if query.Error != nil {
		return 0, query.Error
	}
	return count, nil
}

func GetOpdHistoryPagination(objPayload *structs.ObjPayloadGetOpdHistoryPagination, objQuery *[]structs.ObjQueryOpdHistory) error {
	query := configs.DB1.Table("opds")
	query = query.Select("opds.id AS opd_id, opds.opd_code, opds.opd_date, opds.opd_is_data, opds.opd_create, queues.id AS que_id, queues.que_code, queues.que_datetime, queues.que_admis_id, users.id AS user_id, users.user_fullname, users.user_fullname_en")
	query = query.Joins("LEFT JOIN queues ON queues.id = opds.queue_id")
	query = query.Joins("LEFT JOIN shops ON shops.id = queues.shop_id")
	query = query.Joins("LEFT JOIN users ON users.id = opds.user_id")
	query = query.Where("opds.opd_is_del = ?", 0)
	query = query.Where("opds.customer_id = ?", objPayload.CustomerId)
	query = query.Where("queues.id != ?", objPayload.QueueId)
	query = query.Where("shops.shop_mother_id = ?", objPayload.ShopMotherId)
	if objPayload.SearchDate != "" {
		query = query.Where("opds.opd_date = ?", objPayload.SearchDate)
	}
	if objPayload.SearchText != "" {
		query = query.Where("opds.opd_code LIKE '%" + objPayload.SearchText + "%'")
	}
	// <pagination
	offset := objPayload.CurrentPage * objPayload.PerPage
	query = query.Limit(objPayload.PerPage)
	query = query.Offset(offset)
	// pagination>
	query = query.Order("opds.opd_create DESC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CreateOpd(objQuery *Opd) (err error) {
	query := configs.DB1.Table("opds").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateOpd(opdId int, objQuery *structs.ObjUpdateOpd) (err error) {
	query := configs.DB1.Table("opds")
	query = query.Where("id = ?", opdId)
	query = query.Updates(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CountOpdDiagnostic(opdId int) (int64, error) {
	var count int64
	query := configs.DB1.Table("opd_diagnostics")
	query = query.Select("opd_diagnostics.id")
	query = query.Where("opd_diagnostics.opd_id = ?", opdId)
	query = query.Count(&count)
	if query.Error != nil {
		return 0, query.Error
	}
	return count, nil
}

func CreateOpdDiagnosticBatch(objCreate *[]OpdDiagnostic) error {
	query := configs.DB1.CreateInBatches(&objCreate, 24)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func DeleteOpdDiagnosticNotInIds(opdId int, notInIds []int, objDelete *OpdDiagnostic) error {
	query := configs.DB1.Table("opd_diagnostics")
	query = query.Where("opd_diagnostics.opd_id = ?", opdId)
	if len(notInIds) > 0 {
		query = query.Where("opd_diagnostics.id NOT IN ?", notInIds)
	}
	query = query.Delete(&objDelete)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func ClearOpdDiagnostic(opdId int, objDelete *OpdDiagnostic) error {
	query := configs.DB1.Table("opd_diagnostics")
	query = query.Where("opd_diagnostics.opd_id = ?", opdId)
	query = query.Delete(&objDelete)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CountOpdCustom(opdId int) (int64, error) {
	var count int64
	query := configs.DB1.Table("opd_customs")
	query = query.Select("opd_customs.id")
	query = query.Where("opd_customs.opd_id = ?", opdId)
	query = query.Count(&count)
	if query.Error != nil {
		return 0, query.Error
	}
	return count, nil
}

func CreateOpdCustomBatch(objCreate *[]OpdCustom) error {
	query := configs.DB1.CreateInBatches(&objCreate, 24)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func UpdateOpdCustom(opdcId int, objQuery *OpdCustom) error {
	query := configs.DB1.Table("opd_customs")
	query = query.Where("id = ?", opdcId)
	query = query.Updates(&objQuery)
	return query.Error
}

func DeleteOpdCustomNotInIds(opdId int, notInIds []int, objDelete *OpdCustom) error {
	query := configs.DB1.Table("opd_customs")
	query = query.Where("opd_customs.opd_id = ?", opdId)
	if len(notInIds) > 0 {
		query = query.Where("opd_customs.id NOT IN ?", notInIds)
	}
	query = query.Delete(&objDelete)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func ClearOpdCustom(opdId int, objDelete *OpdCustom) error {
	query := configs.DB1.Table("opd_customs")
	query = query.Where("opd_customs.opd_id = ?", opdId)
	query = query.Delete(&objDelete)
	return query.Error
}

// file
func GetFile(queueId int, objQuery *[]QueueFile) error {
	query := configs.DB1.Table("queue_files")
	query = query.Select("*")
	query = query.Where("queue_files.queue_id = ?", queueId)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CreateFile(objQuery *QueueFile) (err error) {
	query := configs.DB1.Table("queue_files").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateFile(quefId int, objUpdate *QueueFile) (err error) {
	query := configs.DB1.Table("queue_files")
	query = query.Where("queue_files.id = ?", quefId)
	query = query.Updates(&objUpdate)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func DeleteFile(quefId int, objDelete *QueueFile) (err error) {
	query := configs.DB1.Table("queue_files")
	query = query.Where("queue_files.id = ?", quefId)
	query = query.Delete(&objDelete)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

// medical cert
func GetMedicalCert(queueId int, objQuery *[]structs.ObjQueryMedicalCert) error {
	query := configs.DB1.Table("medical_certs")
	query = query.Select("medical_certs.*, users.user_fullname, medical_cert_types.mdct_th, medical_cert_types.mdct_en, medical_cert_types.mdct_group_id, opds.opd_code, customers.ctm_prefix, customers.ctm_fname, customers.ctm_lname, customers.ctm_fname_en, customers.ctm_lname_en")
	query = query.Joins("LEFT JOIN users ON users.id = medical_certs.user_id")
	query = query.Joins("LEFT JOIN medical_cert_types ON medical_cert_types.id = medical_certs.medical_cert_type_id")
	query = query.Joins("LEFT JOIN opds ON opds.id = medical_certs.opd_id")
	query = query.Joins("LEFT JOIN customers ON customers.id = opds.customer_id")
	query = query.Where("medical_certs.mdc_is_del = ?", 0)
	query = query.Where("opds.queue_id = ?", queueId)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CreateMedicalCert(objQuery *MedicalCert) (err error) {
	query := configs.DB1.Table("medical_certs").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateMedicalCert(mdcId int, objUpdate *MedicalCert) (err error) {
	query := configs.DB1.Table("medical_certs")
	query = query.Where("medical_certs.id = ?", mdcId)
	query = query.Updates(&objUpdate)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func ClearMedicalCert(opdId int, objUpdate *MedicalCert) (err error) {
	query := configs.DB1.Table("medical_certs")
	query = query.Where("medical_certs.opd_id = ?", opdId)
	query = query.Updates(&objUpdate)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetQueueTag(queueId int, objQuery *[]structs.ObjQueryQueueTag) error {
	query := configs.DB1.Table("queue_tags")
	query = query.Select("queue_tags.*, tags.tag_name")
	query = query.Joins("LEFT JOIN tags ON tags.id = queue_tags.tag_id")
	query = query.Where("queue_tags.queue_id = ?", queueId)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CreateQueueTagBatch(objCreate *[]QueueTag) error {
	query := configs.DB1.CreateInBatches(&objCreate, 24)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func DeleteQueueTag(queTagId int, objDelete *QueueTag) (err error) {
	query := configs.DB1.Table("queue_tags")
	query = query.Where("queue_tags.id = ?", queTagId)
	query = query.Delete(&objDelete)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

// checking
func GetCheck(queueId int, objQuery *[]structs.ObjQueryCheck) error {
	query := configs.DB1.Table("checks")
	query = query.Select("checks.*, receipts.rec_code, shops.shop_name, shops.shop_phone, doc_settings.sticker_font_size, doc_settings.sticker_width, doc_settings.sticker_height, customers.ctm_id, customers.ctm_prefix, customers.ctm_fname, customers.ctm_lname, customers.ctm_gender, customers.ctm_birthdate, queues.que_code, directions.direction_name, directions.direction_detail, categorys.category_name")
	query = query.Joins("JOIN shops ON checks.shop_id = shops.id")
	query = query.Joins("JOIN customers ON checks.customer_id = customers.id")
	query = query.Joins("JOIN queues ON checks.queue_id = queues.id")
	query = query.Joins("JOIN doc_settings ON doc_settings.shop_id = shops.id")
	query = query.Joins("LEFT JOIN directions ON checks.direction_id = directions.id")
	query = query.Joins("INNER JOIN checkings ON checks.checking_id = checkings.id")
	query = query.Joins("INNER JOIN categorys ON checkings.category_id = categorys.id")
	query = query.Joins("LEFT JOIN receipts ON checks.receipt_id = receipts.id")
	query = query.Where("checks.chk_type_id != ?", 1)
	query = query.Where("checks.chk_is_active = ?", 1)
	query = query.Where("checks.queue_id = ?", queueId)
	query = query.Order("checks.id ASC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetOPDinReceipt(queId int, data *structs.XReyDetail) error {
	query := configs.DB1.Table("queues")
	query = query.Select("queues.id AS queue_id, queues.que_code, queues.que_note,  queues.que_directions, queues.customer_id, queues.shop_id")
	query = query.Where("queues.que_status_id = ?", 4)
	query = query.Where("queues.id = ?", queId)
	query = query.Find(&data)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetCheckById(chkId int, objQuery *Check) error {
	query := configs.DB1.Table("checks")
	query = query.Select("checks.*")
	query = query.Where("checks.id = ?", chkId)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetCheckOldById(chkId int, customerId int, checkingId int, objQuery *Check) error {
	query := configs.DB1.Table("checks")
	query = query.Select("checks.*")
	query = query.Where("checks.id < ?", chkId)
	query = query.Where("checks.customer_id = ?", customerId)
	query = query.Where("checks.checking_id = ?", checkingId)
	query = query.Order("checks.id DESC")
	query = query.Limit(1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func UpdateCheck(chkId int, objUpdate *Check) (err error) {
	query := configs.DB1.Table("checks")
	query = query.Where("id = ?", chkId)
	query = query.Updates(&objUpdate)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateCheckByInterface(chkId int, objUpdate map[string]interface{}) (err error) {
	query := configs.DB1.Table("checks")
	query = query.Where("id = ?", chkId)
	query = query.Updates(&objUpdate)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

// func CancelQueueIDChecking(id int, data *structs.QueueChecking) (err error) {
// 	query := configs.DB1.Table("queue_checkings")
// 	query = query.Where("queue_checkings.queue_id = ?", id)
// 	query = query.Model(&data)
// 	query = query.Updates(map[string]interface{}{"queue_checkings.queci_is_active": 0, "queue_checkings.queci_modify": time.Now().Format("2006-01-02 15:04:05")})
// 	if err = query.Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func CancelQueueIDCourse(id int, data *structs.QueueCourse) (err error) {
// 	query := configs.DB1.Table("queue_courses")
// 	query = query.Where("queue_courses.queue_id = ?", id)
// 	query = query.Model(&data)
// 	query = query.Updates(map[string]interface{}{"queue_courses.quec_is_active": 0, "queue_courses.quec_modify": time.Now().Format("2006-01-02 15:04:05")})
// 	if err = query.Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func CancelQueueIDProduct(id int, data *structs.QueueProduct) (err error) {
// 	query := configs.DB1.Table("queue_products")
// 	query = query.Where("queue_products.queue_id = ?", id)
// 	query = query.Model(&data)
// 	query = query.Updates(map[string]interface{}{"queue_products.quep_is_active": 0, "queue_products.quep_modify": time.Now().Format("2006-01-02 15:04:05")})
// 	if err = query.Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// service
func GetService(shopMId int, customerId int, Que_datetime string, objQuery *[]structs.ObjQueryService) error {
	query := configs.DB1.Table("services")
	query = query.Select("services.*, receipts.rec_code, courses.course_amount, courses.course_ipd, courses.course_opd, courses.course_cost")
	query = query.Joins("INNER JOIN courses ON services.course_id = courses.id")
	query = query.Joins("INNER JOIN receipts ON services.receipt_id = receipts.id")
	query = query.Where("services.ser_is_active = ? ", 1)
	query = query.Where("services.shop_mother_id = ?", shopMId)
	// query = query.Where("services.shop_id = ?", shopId)
	query = query.Where("services.customer_id = ?", customerId)
	query = query.Where("services.ser_datetime <= ?", Que_datetime)
	query = query.Where("services.ser_exp_date IS NULL OR services.ser_exp_date > NOW()")
	query = query.Order("services.ser_create DESC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

// lab x-ray
func GetQueueCheckLabXray(queue_id int, data *structs.GetQueueCheckLabXray) error {
	query := configs.DB1.Table("queues")
	query = query.Select("checks.id")
	query = query.Joins("INNER JOIN checks ON queues.id = checks.queue_id")
	query = query.Where("queues.id = ?", queue_id)
	query = query.Where("checks.chk_type_id != ? ", 1)
	query = query.Where("checks.chk_is_active = ?", 1)
	query = query.Group("queues.id")
	query = query.Find(&data)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetMedicalCertDetailById(mdcId int, data *structs.MedicalCertDetail) error {
	query := configs.DB1.Table("medical_certs")
	query = query.Select("medical_certs.*,medical_cert_types.mdct_th, users.user_fullname, users.user_fullname_en, users.user_license, medical_cert_types.mdct_en, medical_cert_types.mdct_group_id, opds.user_id, opds.queue_id, opds.customer_id, opds.opd_code, opds.opd_date, opds.opd_bw, opds.opd_ht, opds.opd_bmi, opds.opd_t, opds.opd_bsa, opds.opd_vas, opds.opd_pr, opds.opd_bp, opds.opd_rr, opds.opd_sys, opds.opd_dia, opds.opd_o2, opds.opd_fag, opds.opd_alcohol, opds.opd_cc, opds.opd_hpi, opds.opd_pmh, opds.opd_dx, opds.opd_iopat_le, opds.opd_vasc_re, opds.opd_vasc_le, opds.opd_vacc_re, opds.opd_vacc_le, opds.opd_iopat_re, opds.opd_ga, opds.opd_pe, opds.opd_note, opds.opd_sick_startdate, opds.opd_sick_enddate, opds.opd_sick_notrest, opds.opd_sick_air, opds.opd_is_data, opds.opd_is_del, opds.opd_update, opds.opd_create, opd_diagnostics.diagnostic_code, opd_diagnostics.diagnostic_th, opd_diagnostics.diagnostic_en, opd_diagnostics.diagnostic_detail, opd_customs.opdc_name, opd_customs.opdc_value, queues.shop_id")
	query = query.Joins("INNER JOIN medical_cert_types ON medical_certs.medical_cert_type_id = medical_cert_types.id")
	query = query.Joins("INNER JOIN opds ON medical_certs.opd_id = opds.id")
	query = query.Joins("LEFT JOIN opd_diagnostics ON opd_diagnostics.opd_id = opds.id")
	query = query.Joins("LEFT JOIN opd_customs ON opd_customs.opd_id = opds.id")
	query = query.Joins("INNER JOIN queues ON opds.queue_id = queues.id")
	query = query.Joins("LEFT JOIN users ON users.id = medical_certs.user_id")
	query = query.Where("medical_certs.id = ?", mdcId)
	query = query.Where("medical_certs.mdc_is_del = 0")
	query = query.Group("medical_certs.id")
	query = query.Find(&data)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func UpdateDirectionInQue(QueId int, data *structs.PayloadQueueDirection) (err error) {
	query := configs.DB1.Table("queues")
	query = query.Where("queues.id = ?", QueId)
	query = query.Model(&data)
	query = query.Updates(&data)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateQueueStatusID(QueId int) (err error) {
	query := configs.DB1.Table("queues")
	query = query.Where("queues.id = ?", QueId)
	query = query.Updates(map[string]interface{}{"queues.que_status_id": 2, "queues.que_update": time.Now().Format("2006-01-02 15:04:05")})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetOpdDatetime(opdsId int, objQuery *Opd) error {
	query := configs.DB1.Table("opds")
	query = query.Select("opds.id,opds.opd_date")
	query = query.Where("opds.id = ?", opdsId)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

// func UpdateLabplus(id int, objQuery *structs.UpdateLabplus) (err error) {
// 	query := configs.DB1.Table("labplus_apis")
// 	query = query.Where("id = ?", id)
// 	query = query.Updates(&objQuery)
// 	if err = query.Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func GetQueuelabplus(queue_id int, objQuery *structs.QueuesLabplus) (err error) {
// 	query := configs.DB1.Table("queues_labplus")
// 	query = query.Select("queues_labplus.*")
// 	query = query.Where("queues_labplus.queue_id = ?", queue_id)
// 	query = query.Find(&objQuery)
// 	if err = query.Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func UpdateQueuesLabplus(id int, data *structs.UpdateQueuesLabplus) (err error) {
// 	query := configs.DB1.Table("queues_labplus")
// 	query = query.Where("id = ?", id)
// 	var inInterface map[string]interface{}
// 	in, _ := json.Marshal(&data)
// 	json.Unmarshal(in, &inInterface)
// 	query = query.Updates(&inInterface)
// 	if err = query.Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

func GetCheckCheckingLabplus(queue_id int, data *[]structs.CheckCheckingLabplus) (err error) {
	query := configs.DB1.Table("checks")
	query = query.Select("checks.id,checks.queue_id, checks.checking_id, checkings.checking_is_labplus, checks.chk_code,checks.chk_name, checks.chk_is_active")
	query = query.Joins("INNER JOIN checkings ON checks.checking_id = checkings.id")
	query = query.Where("checks.queue_id = ?", queue_id)
	query = query.Where("checks.chk_is_active = ?", 1)
	query = query.Where("checkings.checking_is_labplus = ?", 1)
	query = query.Find(&data)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func UpdateCheckingLabplus(id int, data *structs.UpdateCheckingLabplus) (err error) {
	query := configs.DB1.Table("checks")
	query = query.Where("id = ?", id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&data)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func AddCheckLabplus(objQuery *structs.AddCheckLabplus) (err error) {
	query := configs.DB1.Table("check_labplus").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetCheckLabplus(queue_id int, data *[]structs.CheckLabplus) (err error) {
	query := configs.DB1.Table("check_labplus")
	query = query.Select("check_labplus.*")
	query = query.Joins("INNER JOIN checks ON check_labplus.check_id = checks.id")
	query = query.Where("check_labplus.queue_id = ?", queue_id)
	query = query.Where("checks.chk_is_active = ?", 1)
	query = query.Find(&data)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetCheckLabplusId(check_id int, data *structs.CheckLabplus) (err error) {
	query := configs.DB1.Table("check_labplus")
	query = query.Select("check_labplus.*")
	query = query.Joins("INNER JOIN checks ON check_labplus.check_id = checks.id")
	query = query.Where("check_labplus.check_id = ?", check_id)
	query = query.Where("checks.chk_is_active = ?", 1)
	query = query.Find(&data)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetCheckCheckingLabplusId(queue_id int, chk_code string, data *structs.CheckCheckingLabplus) (err error) {
	query := configs.DB1.Table("checks")
	query = query.Select("checks.id,checks.queue_id, checks.checking_id, checkings.checking_is_labplus, checks.chk_code,checks.chk_name,checks.chk_value,checks.chk_direction_detail,checks.chk_flag, checks.chk_is_active")
	query = query.Joins("INNER JOIN checkings ON checks.checking_id = checkings.id")
	query = query.Where("checks.queue_id = ?", queue_id)
	query = query.Where("checks.chk_code = ?", chk_code)
	query = query.Where("checks.chk_is_active = ?", 1)
	query = query.Where("checkings.checking_is_labplus = ?", 1)
	query = query.Find(&data)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetCheckXray(queueId int, objQuery *[]structs.ObjQueryCheck) error {
	query := configs.DB1.Table("checks")
	query = query.Select("checks.*,receipts.rec_user_fullname, receipts.rec_code, shops.shop_name, shops.shop_phone, doc_settings.sticker_font_size, doc_settings.sticker_width, doc_settings.sticker_height, customers.ctm_id, customers.ctm_prefix, customers.ctm_fname, customers.ctm_lname, customers.ctm_gender, customers.ctm_birthdate, queues.que_code, directions.direction_name, directions.direction_detail, categorys.category_name")
	query = query.Joins("JOIN shops ON checks.shop_id = shops.id")
	query = query.Joins("JOIN customers ON checks.customer_id = customers.id")
	query = query.Joins("JOIN queues ON checks.queue_id = queues.id")
	query = query.Joins("JOIN doc_settings ON doc_settings.shop_id = shops.id")
	query = query.Joins("LEFT JOIN directions ON checks.direction_id = directions.id")
	query = query.Joins("INNER JOIN checkings ON checks.checking_id = checkings.id")
	query = query.Joins("INNER JOIN categorys ON checkings.category_id = categorys.id")
	query = query.Joins("LEFT JOIN receipts ON checks.receipt_id = receipts.id")
	query = query.Where("checks.chk_type_id = ?", 3)
	query = query.Where("checks.chk_is_active = ?", 1)
	query = query.Where("checks.queue_id = ?", queueId)
	query = query.Order("checks.id ASC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetCheckLab(queueId int, objQuery *[]structs.ObjQueryCheck) error {
	query := configs.DB1.Table("checks")
	query = query.Select("checks.*, receipts.rec_code, shops.shop_name, shops.shop_phone, doc_settings.sticker_font_size, doc_settings.sticker_width, doc_settings.sticker_height, customers.ctm_id, customers.ctm_prefix, customers.ctm_fname, customers.ctm_lname, customers.ctm_gender, customers.ctm_birthdate, queues.que_code, directions.direction_name, directions.direction_detail, categorys.category_name")
	query = query.Joins("JOIN shops ON checks.shop_id = shops.id")
	query = query.Joins("JOIN customers ON checks.customer_id = customers.id")
	query = query.Joins("JOIN queues ON checks.queue_id = queues.id")
	query = query.Joins("JOIN doc_settings ON doc_settings.shop_id = shops.id")
	query = query.Joins("LEFT JOIN directions ON checks.direction_id = directions.id")
	query = query.Joins("INNER JOIN checkings ON checks.checking_id = checkings.id")
	query = query.Joins("INNER JOIN categorys ON checkings.category_id = categorys.id")
	query = query.Joins("LEFT JOIN receipts ON checks.receipt_id = receipts.id")
	query = query.Where("checks.chk_type_id = ?", 2)
	query = query.Where("checks.chk_is_active = ?", 1)
	query = query.Where("checks.queue_id = ?", queueId)
	query = query.Order("checks.id ASC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetShopLabplus(shop_id int, data *structs.Labplus) (err error) {
	query := configs.DB1.Table("labplus_apis")
	query = query.Select("labplus_apis.*")
	query = query.Where("labplus_apis.shop_id = ?", shop_id)
	err = query.Scan(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateLabplus(id int, objQuery *structs.UpdateLabplus) (err error) {
	query := configs.DB1.Table("labplus_apis")
	query = query.Where("id = ?", id)
	query = query.Updates(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetMaxQueueNumber(objPayload *structs.ObjPayloadCreateQueueByOrder) (num int, err error) {
	var data int
	checkDate := time.Now().Format("2006-01-02")
	query := configs.DB1.Table("queues")
	query = query.Select("MAX(que_number) AS maxQueue")
	query = query.Where("DATE(queues.que_datetime) = ? AND shop_id = ? AND que_type_id = ?", checkDate, objPayload.ShopId, objPayload.QueTypeId)
	if objPayload.QueTypeId == 1 {
		query = query.Where("que_admis_id = ?", objPayload.QueAdmisId)
	}
	if err = query.Find(&data).Error; err != nil {
		return 0, err
	}
	return data, nil
}
