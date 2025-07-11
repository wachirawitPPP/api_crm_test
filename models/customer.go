package models

import (
	"encoding/json"
	"linecrmapi/configs"
	"linecrmapi/structs"
	"time"
	// "time"
)

func CountCustomerPagination(objPayload *structs.ObjPayloadGetCustomerPagination) (int64, error) {
	var count int64
	query := configs.DB1.Table("customers")
	// query = query.Joins("LEFT JOIN shop_roles on shop_roles.id = user_shops.shop_role_id")
	query = query.Where("customers.shop_id = ?", objPayload.ShopId)
	if *objPayload.CtmIsActive != -1 {
		query = query.Where("customers.ctm_is_active = ?", objPayload.CtmIsActive)
	}
	if *objPayload.SearchText != "" {
		query = query.Where("customers.ctm_id LIKE ? OR customers.ctm_fname LIKE ? OR customers.ctm_lname LIKE ? OR customers.ctm_nname LIKE ? OR customers.ctm_nname LIKE ? OR customers.ctm_nname LIKE ? OR customers.ctm_fname_en LIKE ? OR customers.ctm_lname_en LIKE ? OR customers.ctm_citizen_id LIKE ? OR customers.ctm_passport_id LIKE ? OR customers.ctm_tel LIKE ? OR CONCAT(customers.ctm_fname,' ',customers.ctm_lname) LIKE ?", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%")
	}
	query = query.Count(&count)
	if query.Error != nil {
		return 0, query.Error
	}
	return count, nil
}

func GetCustomerPagination(objPayload *structs.ObjPayloadGetCustomerPagination, objQuery *[]structs.ObjQueryGetCustomerPagination) error {
	query := configs.DB1.Table("customers")
	query = query.Select("customers.*, customer_groups.cg_name, ref_right_treatments.rt_code, ref_right_treatments.rt_name")
	query = query.Joins("LEFT JOIN customer_groups on customer_groups.id = customers.customer_group_id")
	query = query.Joins("LEFT JOIN ref_right_treatments on ref_right_treatments.id = customers.right_treatment_id")
	query = query.Where("customers.shop_id = ?", objPayload.ShopId)
	if *objPayload.CtmIsActive != -1 {
		query = query.Where("customers.ctm_is_active = ?", objPayload.CtmIsActive)
	}
	if *objPayload.SearchText != "" {
		query = query.Where("customers.ctm_id LIKE ? OR customers.ctm_fname LIKE ? OR customers.ctm_lname LIKE ? OR customers.ctm_nname LIKE ? OR customers.ctm_nname LIKE ? OR customers.ctm_nname LIKE ? OR customers.ctm_fname_en LIKE ? OR customers.ctm_lname_en LIKE ? OR customers.ctm_citizen_id LIKE ? OR customers.ctm_passport_id LIKE ? OR customers.ctm_tel LIKE ? OR CONCAT(customers.ctm_fname,' ',customers.ctm_lname) LIKE ?", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%")
	}
	// <pagination
	offset := objPayload.CurrentPage * objPayload.PerPage
	query = query.Limit(objPayload.PerPage)
	query = query.Offset(offset)
	// pagination>
	query = query.Order("LENGTH(customers.ctm_id), customers.ctm_id")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetCustomerById(customerId int, objQuery *structs.ObjQueryCustomer) error {
	query := configs.DB1.Table("customers")
	// query = query.Select("customers.*, customer_groups.cg_name, ref_right_treatments.rt_code, ref_right_treatments.rt_name")
	query = query.Select("customers.*, customer_groups.cg_name,customer_groups.cg_save_type,customer_groups.cg_save, ref_right_treatments.rt_code, ref_right_treatments.rt_name, ref_right_treatments.rt_name_en")
	query = query.Joins("LEFT JOIN customer_groups on customer_groups.id = customers.customer_group_id")
	query = query.Joins("LEFT JOIN ref_right_treatments on ref_right_treatments.id = customers.right_treatment_id")
	query = query.Where("customers.id = ?", customerId)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
func GetCustomerOnlinesById(customerId int, objQuery *structs.ObjGetCustomerOnline) error {
	query := configs.DB1.Table("customer_onlines")
	// query = query.Select("customers.*, customer_groups.cg_name, ref_right_treatments.rt_code, ref_right_treatments.rt_name")
	query = query.Select("customer_onlines.co_fname, customer_onlines.co_lname,customer_onlines.id,customer_onlines.co_email,customer_onlines.co_line_id,customer_onlines.co_citizen_id")
	query = query.Where("customer_onlines.id = ?", customerId)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetCustomerTagById(customerId int, objQuery *[]structs.ObjQueryCustomerTag) error {
	query := configs.DB1.Table("customer_tags")
	query = query.Select("customer_tags.*, tags.tag_name, tag_types.tag_type_th, tag_types.tag_type_en")
	query = query.Joins("LEFT JOIN tags on tags.id = customer_tags.tag_id")
	query = query.Joins("LEFT JOIN tag_types on tag_types.id = tags.tag_type_id")
	query = query.Where("customer_tags.customer_id = ?", customerId)
	query = query.Where("tags.tag_type_id = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetCustomerFamilyById(customerId int, objQuery *[]structs.ObjQueryCustomerFamily) error {
	query := configs.DB1.Table("customer_familys")
	query = query.Select("customer_familys.*, customers.ctm_fname, customers.ctm_lname")
	query = query.Joins("LEFT JOIN customers on customers.id = customer_familys.cf_customer_id")
	query = query.Where("customer_familys.customer_id = ?", customerId)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetCustomerContactById(customerId int, objQuery *[]structs.ObjQueryCustomerContact) error {
	query := configs.DB1.Table("customer_contacts")
	query = query.Select("customer_contacts.*")
	query = query.Where("customer_contacts.customer_id = ?", customerId)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetCustomerCheckImport(shopMotherId int, CtmId string, CtmCitizenId string, objQuery *structs.ObjQueryCustomer) error {
	query := configs.DB1.Table("customers")
	query = query.Select("customers.id, customers.ctm_id, customers.ctm_citizen_id, customers.ctm_passport_id")
	query = query.Where("customers.shop_mother_id = ?", shopMotherId)
	query = query.Where("customers.ctm_id = '" + CtmId + "' OR customers.ctm_citizen_id = '" + CtmCitizenId + "'")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetCheckImportCode(shopId int, ctmIds []string, objQuery *[]structs.ObjQueryCheckImport) error {
	query := configs.DB1.Table("customers")
	query = query.Select("customers.ctm_id, customers.ctm_citizen_id")
	query = query.Where("customers.shop_id = ?", shopId)
	query = query.Where("customers.ctm_id IN ?", ctmIds)
	query = query.Find(&objQuery)
	return query.Error
}

func GetCheckImportCitizen(shopMotherId int, ctmCitizenIds []string, objQuery *[]structs.ObjQueryCheckImport) error {
	query := configs.DB1.Table("customers")
	query = query.Select("customers.ctm_id, customers.ctm_citizen_id")
	query = query.Where("customers.shop_mother_id = ?", shopMotherId)
	query = query.Where("customers.ctm_citizen_id IN ?", ctmCitizenIds)
	query = query.Find(&objQuery)
	return query.Error
}

func SearchFamily(objPayload *structs.ObjPayloadSearchFamily, objQuery *[]structs.ObjResponseSearchFamily) error {
	query := configs.DB1.Table("customers")
	query = query.Select("customers.id, customers.ctm_fname, customers.ctm_lname")
	query = query.Where("customers.shop_id = ?", objPayload.ShopId)
	query = query.Where("customers.ctm_is_active = ?", 1)
	if len(objPayload.NotInIds) > 0 {
		query = query.Where("customers.id NOT IN ?", objPayload.NotInIds)
	}
	query = query.Where("customers.ctm_id LIKE ? OR customers.ctm_fname LIKE ? OR customers.ctm_lname LIKE ? OR customers.ctm_nname LIKE ? OR customers.ctm_nname LIKE ? OR customers.ctm_nname LIKE ? OR customers.ctm_fname_en LIKE ? OR customers.ctm_lname_en LIKE ? OR customers.ctm_citizen_id LIKE ? OR customers.ctm_passport_id LIKE ?", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%")
	query = query.Order("customers.ctm_fname ASC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetCustomerDocSetting(shopId int, data *structs.ObjQueryCustomerDocSetting) (err error) {
	query := configs.DB1.Table("doc_settings")
	query = query.Select("customer_id_default, customer_number_default, customer_number_digit, customer_type")
	query = query.Where("shop_id = ?", shopId)
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCustomerDocSetting(shopId int, nextNumberDefault int, data *DocSetting) (err error) {
	query := configs.DB1.Table("doc_settings")
	query = query.Where("shop_id = ?", shopId)
	query = query.Model(&data)
	query = query.Updates(map[string]interface{}{"customer_number_default": nextNumberDefault})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CheckCustomerCode(shopId int, ctmId string) (int64, error) {
	var count int64
	query := configs.DB1.Table("customers")
	query = query.Select("customers.id")
	query = query.Where("customers.shop_id = ?", shopId)
	query = query.Where("customers.ctm_id = ?", ctmId)
	query = query.Count(&count)
	if query.Error != nil {
		return 0, query.Error
	}
	return count, nil
}

func CreateCustomer(objQuery *Customer) (err error) {
	query := configs.DB1.Table("customers").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CreateCustomerBatch(objCreate *[]Customer) error {
	query := configs.DB1.CreateInBatches(&objCreate, 100)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CreateCustomerTagBatch(objCreate *[]CustomerTag) error {
	query := configs.DB1.CreateInBatches(&objCreate, 100)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func DeleteCustomerTagBatch(customerId int, deleteIds []int, objDelete *CustomerTag) error {
	query := configs.DB1.Table("customer_tags")
	query = query.Where("customer_tags.customer_id = ?", customerId)
	query = query.Where("customer_tags.tag_id IN ?", deleteIds)
	query = query.Delete(&objDelete)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func ClearCustomerTag(customerId int, objDelete *CustomerTag) error {
	query := configs.DB1.Table("customer_tags")
	query = query.Where("customer_tags.customer_id = ?", customerId)
	query = query.Delete(&objDelete)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CreateCustomerFamilyBatch(objCreate *[]CustomerFamilys) error {
	query := configs.DB1.CreateInBatches(&objCreate, 100)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func DeleteCustomerFamilyBatch(customerId int, deleteIds []int, objDelete *CustomerFamilys) error {
	query := configs.DB1.Table("customer_familys")
	query = query.Where("customer_familys.customer_id = ?", customerId)
	query = query.Where("customer_familys.cf_customer_id IN ?", deleteIds)
	query = query.Delete(&objDelete)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func ClearCustomerFamily(customerId int, objDelete *CustomerFamilys) error {
	query := configs.DB1.Table("customer_familys")
	query = query.Where("customer_familys.customer_id = ?", customerId)
	query = query.Delete(&objDelete)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CreateCustomerContactBatch(objCreate *[]CustomerContact) error {
	query := configs.DB1.CreateInBatches(&objCreate, 100)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func DeleteCustomerContactBatch(customerId int, deleteIds []string, objDelete *CustomerContact) error {
	query := configs.DB1.Table("customer_contacts")
	query = query.Where("customer_contacts.customer_id = ?", customerId)
	query = query.Where("customer_contacts.cc_name IN ?", deleteIds)
	query = query.Delete(&objDelete)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func ClearCustomerContact(customerId int, objDelete *CustomerContact) error {
	query := configs.DB1.Table("customer_contacts")
	query = query.Where("customer_contacts.customer_id = ?", customerId)
	query = query.Delete(&objDelete)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func UpdateCustomer(customerId int, objUpdate *structs.ObjQueryUpdateCustomer) (err error) {

	var inInterface map[string]interface{}
	in, _ := json.Marshal(&objUpdate)
	json.Unmarshal(in, &inInterface)

	query := configs.DB1.Table("customers")
	query = query.Where("id = ?", customerId)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil

}

func ActiveCustomer(customerId int, objUpdate *structs.ObjQueryActiveCustomer) error {
	query := configs.DB1.Table("customers")
	query = query.Where("id = ?", customerId)
	query = query.Updates(&objUpdate)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CountCustomerOpdPagination(objPayload *structs.ObjPayloadGetCustomerOpdPagination, CustomerId int, shopId int) (int64, error) {
	var count int64
	query := configs.DB1.Table("opds")
	query = query.Joins("LEFT JOIN queues ON queues.id = opds.queue_id")
	query = query.Joins("LEFT JOIN users ON users.id = opds.user_id")
	query = query.Joins("LEFT JOIN customers ON customers.id = opds.customer_id")
	query = query.Joins("LEFT JOIN shops ON customers.shop_id = shops.id")
	query = query.Where("shops.shop_mother_id = ?", shopId)
	query = query.Where("queues.customer_id = ?", CustomerId)
	query = query.Where("queues.que_type_id = ?", objPayload.QueueTypeId)
	if *objPayload.SearchText != "" {
		query = query.Where("opds.opd_code LIKE ? OR queues.que_code LIKE ? OR users.user_fullname LIKE ?", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%")
	}
	query = query.Count(&count)
	if query.Error != nil {
		return 0, query.Error
	}
	return count, nil
}

func GetCustomerOpdPagination(objPayload *structs.ObjPayloadGetCustomerOpdPagination, objQuery *[]structs.ObjQueryGetCustomerOpdPagination, CustomerId int, shopId int) error {
	query := configs.DB1.Table("opds")
	query = query.Select("opds.id AS opd_id, opds.opd_code, opds.opd_date, opds.opd_create, queues.id AS que_id,	queues.shop_id AS shop_id, queues.que_code, queues.que_datetime, queues.que_admis_id, users.user_fullname")
	query = query.Joins("LEFT JOIN queues ON queues.id = opds.queue_id")
	query = query.Joins("LEFT JOIN users ON users.id = opds.user_id")
	query = query.Joins("LEFT JOIN customers ON customers.id = opds.customer_id")
	query = query.Joins("LEFT JOIN shops ON customers.shop_id = shops.id")
	query = query.Where("shops.shop_mother_id = ?", shopId)
	query = query.Where("queues.customer_id = ?", CustomerId)
	query = query.Where("queues.que_type_id = ?", objPayload.QueueTypeId)
	if *objPayload.SearchText != "" {
		query = query.Where("opds.opd_code LIKE ? OR queues.que_code LIKE ? OR users.user_fullname LIKE ?", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%")
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

func GetCustomerOpdShopByIds(shopId int, objQuery *structs.Shop) error {
	query := configs.DB1.Table("shops")
	query = query.Select("shops.*")
	query = query.Where("shops.id = ?", shopId)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetCustomerOpdDiagnosticByIds(opdIds []int, objQuery *[]OpdDiagnostic) error {
	query := configs.DB1.Table("opd_diagnostics")
	query = query.Select("opd_diagnostics.*")
	query = query.Where("opd_diagnostics.opd_id IN ?", opdIds)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CountCustomerCheckPagination(objPayload *structs.ObjPayloadGetCustomerCheckPagination) (int64, error) {
	var count int64
	query := configs.DB1.Table("checks")
	query = query.Joins("LEFT JOIN directions ON directions.id = checks.direction_id")
	query = query.Joins("LEFT JOIN customers ON customers.id = checks.customer_id")
	query = query.Joins("LEFT JOIN shops ON customers.shop_id = shops.id")
	query = query.Joins("LEFT JOIN queues ON checks.queue_id = queues.id")
	query = query.Where("shops.shop_mother_id = ?", objPayload.ShopId)
	// query = query.Where("checks.shop_id = ?", objPayload.ShopId)
	query = query.Where("checks.customer_id = ?", objPayload.CustomerId)
	if *objPayload.SearchText != "" {
		query = query.Where("checks.chk_code LIKE ? OR checks.chk_name LIKE ?", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%")
	}
	query = query.Count(&count)
	if query.Error != nil {
		return 0, query.Error
	}
	return count, nil
}

func GetCustomerCheckPagination(objPayload *structs.ObjPayloadGetCustomerCheckPagination, objQuery *[]structs.ObjQueryGetCustomerCheckPagination) error {
	query := configs.DB1.Table("checks")
	query = query.Select("checks.*, directions.direction_name, queues.que_code, queues.que_datetime , queues.shop_id AS que_shop_id")
	query = query.Joins("LEFT JOIN directions ON directions.id = checks.direction_id")
	query = query.Joins("LEFT JOIN customers ON customers.id = checks.customer_id")
	query = query.Joins("LEFT JOIN shops ON customers.shop_id = shops.id")
	query = query.Joins("LEFT JOIN queues ON checks.queue_id = queues.id")
	query = query.Where("shops.shop_mother_id = ?", objPayload.ShopId)
	// query = query.Where("checks.shop_id = ?", objPayload.ShopId)
	query = query.Where("checks.customer_id = ?", objPayload.CustomerId)
	if *objPayload.SearchText != "" {
		query = query.Where("checks.chk_code LIKE ? OR checks.chk_name LIKE ?", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%")
	}
	// <pagination
	offset := objPayload.CurrentPage * objPayload.PerPage
	query = query.Limit(objPayload.PerPage)
	query = query.Offset(offset)
	// pagination>
	query = query.Order("checks.chk_create DESC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CountCustomerServicePagination(objPayload *structs.ObjPayloadGetCustomerServicePagination) (int64, error) {
	var count int64
	query := configs.DB1.Table("services")
	query = query.Joins("LEFT JOIN users ON users.id = services.user_id")

	query = query.Joins("LEFT JOIN customers ON customers.id = services.customer_id")
	query = query.Joins("LEFT JOIN shops ON customers.shop_id = shops.id")
	query = query.Where("shops.shop_mother_id = ?", objPayload.ShopId)

	query = query.Where("services.ser_is_active = ? ", 1)
	// query = query.Where("services.shop_id = ?", objPayload.ShopId)
	query = query.Where("services.customer_id = ?", objPayload.CustomerId)
	if *objPayload.SearchText != "" {
		query = query.Where("services.ser_code LIKE ? OR services.ser_name LIKE ?", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%")
	}
	query = query.Count(&count)
	if query.Error != nil {
		return 0, query.Error
	}
	return count, nil
}

func GetCustomerServicePagination(objPayload *structs.ObjPayloadGetCustomerServicePagination, objQuery *[]structs.ObjQueryGetCustomerServicePagination) error {
	query := configs.DB1.Table("services")
	query = query.Select("services.*, users.user_fullname")
	query = query.Joins("LEFT JOIN users ON users.id = services.user_id")

	query = query.Joins("LEFT JOIN customers ON customers.id = services.customer_id")
	query = query.Joins("LEFT JOIN shops ON customers.shop_id = shops.id")
	query = query.Where("shops.shop_mother_id = ?", objPayload.ShopId)

	query = query.Where("services.ser_is_active = ? ", 1)
	// query = query.Where("services.shop_id = ?", objPayload.ShopId)
	query = query.Where("services.customer_id = ?", objPayload.CustomerId)
	if *objPayload.SearchText != "" {
		query = query.Where("services.ser_code LIKE ? OR services.ser_name LIKE ?", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%")
	}
	// <pagination
	offset := objPayload.CurrentPage * objPayload.PerPage
	query = query.Limit(objPayload.PerPage)
	query = query.Offset(offset)
	// pagination>
	query = query.Order("services.ser_create DESC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetReceiptCustomerList(filter structs.ObjPayloadSearchReceiptCustomer, isCount bool, dt *[]structs.ReceiptListCustomer, customerId int, shopId int) (err error) {
	query := configs.DB1.Table("receipts")
	if isCount == true {
		query = query.Select("receipts.*,receipts.shop_id AS shop_id,customers.shop_id AS customer_shop_id,receipts.rec_create AS rec_datetime, shops.shop_name, users.user_fullname, customers.ctm_id, customers.ctm_fname, customers.ctm_lname, customers.ctm_fname_en, customers.ctm_lname_en, queues.que_code, invoices.inv_code")
	} else {
		query = query.Select("receipts.id")
	}

	query = query.Joins("INNER JOIN customers ON receipts.customer_id = customers.id")
	query = query.Joins("INNER JOIN users ON receipts.user_id = users.id")
	query = query.Joins("INNER JOIN invoices ON receipts.invoice_id = invoices.id")
	query = query.Joins("LEFT JOIN queues ON receipts.queue_id = queues.id")
	query = query.Joins("INNER JOIN shops ON receipts.shop_id = shops.id")
	query = query.Where("shops.shop_mother_id = ?", shopId)
	query = query.Where("receipts.rec_is_active != 0")
	query = query.Where("receipts.customer_id = ?", customerId)

	if *filter.Rec_is_active != "" {
		if *filter.Rec_is_active == "1" {
			query = query.Where("receipts.rec_is_active = 1")
		} else if *filter.Rec_is_active == "2" {
			query = query.Where("receipts.rec_is_active = 2")
		}
	}
	if *filter.Search != "" {
		query = query.Where("queues.que_code LIKE '%" + *filter.Search + "%' OR invoices.inv_code LIKE '%" + *filter.Search + "%' OR receipts.rec_code LIKE '%" + *filter.Search + "%'")
	}
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

func GetShopMother_(shopId int, objQuery *Shop) error {
	query := configs.DB1.Table("shops")
	query = query.Select("shops.id, shops.shop_mother_id")
	query = query.Where("shops.id = ?", shopId)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetShopCode(shopId int, objQuery *Shop) error {
	query := configs.DB1.Table("shops")
	query = query.Select("shops.id, shops.shop_code")
	query = query.Where("shops.id = ?", shopId)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetCustomer_(shopMotherId int, CustomerId int, citizenId string, objQuery *structs.CheckCustomer) error {
	query := configs.DB1.Table("customers")
	query = query.Select("customers.id, shops.shop_code, shops.shop_name")
	query = query.Joins("LEFT JOIN shops ON shops.id = customers.shop_id")
	query = query.Where("customers.shop_mother_id = ?", shopMotherId)
	if CustomerId != 0 {
		query = query.Where("customers.id != ?", CustomerId)
	}

	query = query.Where("customers.ctm_citizen_id = ?", citizenId)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func AcceptPDPA(customerId int, customerImage string) error {
	query := configs.DB1.Table("customers")
	query = query.Where("id = ?", customerId)
	query = query.Updates(map[string]interface{}{"customers.ctm_subscribe_pdpa": 1, "customers.ctm_subscribe_pdpa_image": customerImage, "customers.ctm_update": time.Now().Format("2006-01-02 15:04:05")})
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func AddLogCustomer(log *structs.LogCustomer) (err error) {
	query := configs.DBL1.Table("log_customers").Create(&log)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func AppointmentSearchByCustomer(filter structs.PayloadSearchAppointmentByCustomer, isCount bool, product *[]structs.AppointmentListHistory, customerId int) (err error) {
	query := configs.DB1.Table("appointments")
	query = query.Select("appointments.*,customers.ctm_fname,customers.ctm_lname,customers.ctm_fname_en,customers.ctm_lname_en,users.user_fullname,users.user_fullname_en,shops.shop_name")
	query = query.Joins("LEFT JOIN customers ON customers.id = appointments.customer_id")
	query = query.Joins("LEFT JOIN users ON users.id = appointments.user_id")
	query = query.Joins("LEFT JOIN shops ON appointments.shop_id = shops.id")
	// query = query.Where("appointments.shop_id = ?", filter.Shop_id)
	query = query.Where("appointments.customer_id = ?", customerId)
	query = query.Where("appointments.ap_is_del = 0")

	if *filter.Search != "" {
		query = query.Where("appointments.ap_topic LIKE '%" + *filter.Search + "%' OR appointments.ap_note LIKE '%" + *filter.Search + "%' ")
	}
	if *filter.Date != "" {
		query = query.Where("appointments.ap_datetime LIKE '%" + *filter.Date + "%'")
	}
	if *filter.Type != "" {
		query = query.Where("appointments.ap_type = ?", *filter.Type)
	}
	if *filter.Is_active != "" {
		query = query.Where("appointments.ap_status_id = ?", *filter.Is_active)
	}
	if *filter.Date_from != ""{
		query = query.Where("appointments.ap_datetime >=?", *filter.Date_from)
	}

	if isCount == true {
		offset := filter.ActivePage * filter.PerPage
		query = query.Limit(filter.PerPage)
		query = query.Offset(offset)
	}
	query = query.Order("appointments.ap_datetime DESC")
	query = query.Order("appointments.id DESC")
	if err = query.Scan(&product).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTokenPDPA(customerId int, token string) error {
	query := configs.DB1.Table("customers")
	query = query.Where("id = ?", customerId)
	query = query.Updates(map[string]interface{}{"customers.ctm_subscribe_pdpa_token": token, "customers.ctm_update": time.Now().Format("2006-01-02 15:04:05")})
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetCustomerByToken(token string, objQuery *structs.ObjQueryCustomer) error {
	query := configs.DB1.Table("customers")
	// query = query.Select("customers.*, customer_groups.cg_name, ref_right_treatments.rt_code, ref_right_treatments.rt_name")
	query = query.Select("customers.*, customer_groups.cg_name,customer_groups.cg_save_type,customer_groups.cg_save, ref_right_treatments.rt_code, ref_right_treatments.rt_name, ref_right_treatments.rt_name_en")
	query = query.Joins("LEFT JOIN customer_groups on customer_groups.id = customers.customer_group_id")
	query = query.Joins("LEFT JOIN ref_right_treatments on ref_right_treatments.id = customers.right_treatment_id")
	query = query.Where("customers.ctm_subscribe_pdpa_token = ?", token)
	query = query.Where("customers.ctm_subscribe_pdpa = ?", 0)
	query = query.Where("customers.ctm_is_del = ?", 0)
	query = query.Where("customers.ctm_is_active = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetCustomerPaymentBalance(Shop_id int, customerId int, objQuery *structs.DashboardPayment) error {
	query := configs.DB1.Table("invoices")
	query = query.Select(`SUM(invoices.inv_total) - SUM(invoices.inv_pay_total) AS balance_total`)
	query = query.Joins("INNER JOIN customers ON invoices.customer_id = customers.id")
	query = query.Joins("INNER JOIN users ON invoices.user_id = users.id")
	query = query.Joins("INNER JOIN customer_groups ON customers.customer_group_id = customer_groups.id")
	query = query.Where("invoices.shop_id = ?", Shop_id)
	query = query.Where("customers.id = ?", customerId)
	query = query.Where("invoices.inv_is_active != 0")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

// Add on
func GetShopPackage(ShopId int, data *ShopPackage) (err error) {
	query := configs.DB1.Table("shops")
	query = query.Select("shops.id,shops.shop_type_id,package_order_details.package_order_id, package_order_details.package_id")
	query = query.Joins("JOIN package_order_details ON shops.package_order_detail_id = package_order_details.id")
	query = query.Where("shops.id = ?", ShopId)
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetShopPackageOrderIds(package_order_id int, package_ids []int, data *[]ShopPackageOrderId) (err error) {
	query := configs.DB1.Table("package_order_details")
	query = query.Select("package_order_details.id,package_order_details.package_order_id,package_order_details.package_id")
	query = query.Where("package_order_details.package_order_id = ?", package_order_id)
	query = query.Where("package_order_details.package_id IN ?", package_ids)
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetShopPackageOrderId(package_order_id int, package_id int, data *ShopPackageOrderId) (err error) {
	query := configs.DB1.Table("package_order_details")
	query = query.Select("package_order_details.id,package_order_details.package_order_id,package_order_details.package_id")
	query = query.Where("package_order_details.package_order_id = ?", package_order_id)
	query = query.Where("package_order_details.package_id = ?", package_id)
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetShopAddonIds(ShopId int, package_ids []int, data *[]ShopAddon) (err error) {
	query := configs.DB1.Table("shop_addons")
	query = query.Select("shop_addons.id, shop_addons.shop_id, shop_addons.package_id")
	query = query.Where("shop_addons.shop_id = ?", ShopId)
	query = query.Where("shop_addons.package_id IN ?", package_ids)
	query = query.Where("shop_addons.shop_addon_is_del = ?", 0)
	query = query.Where("shop_addons.shop_addon_expire > DATE(NOW())")
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetShopAddonId(ShopId int, package_id int, data *ShopAddon) (err error) {
	query := configs.DB1.Table("shop_addons")
	query = query.Select("shop_addons.id, shop_addons.shop_id, shop_addons.package_id")
	query = query.Where("shop_addons.shop_id = ?", ShopId)
	query = query.Where("shop_addons.package_id = ?", package_id)
	query = query.Where("shop_addons.shop_addon_is_del = ?", 0)
	query = query.Where("shop_addons.shop_addon_expire > DATE(NOW())")
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func CountHistoryCustomerCheckPagination(objPayload *structs.ObjPayloadGetCustomerCheckLabXaryPagination, customerId int, ShopMotherId int) (int64, error) {
	var count int64
	query := configs.DB1.Table("checks")
	query = query.Joins("LEFT JOIN directions ON directions.id = checks.direction_id")
	query = query.Joins("LEFT JOIN customers ON customers.id = checks.customer_id")
	query = query.Joins("LEFT JOIN shops ON customers.shop_id = shops.id")
	query = query.Joins("LEFT JOIN queues ON checks.queue_id = queues.id")
	query = query.Where("shops.shop_mother_id = ?", ShopMotherId)
	query = query.Where("checks.chk_type_id = ?", 1)
	query = query.Where("checks.customer_id = ?", customerId)
	if *objPayload.SearchText != "" {
		query = query.Where("checks.chk_code LIKE ? OR checks.chk_name LIKE ?", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%")
	}
	if *objPayload.SearchDate != "" {
		query = query.Where("checks.chk_datetime LIKE '%" + *objPayload.SearchDate + "%'")
	}
	query = query.Count(&count)
	if query.Error != nil {
		return 0, query.Error
	}
	return count, nil
}

func GetHistoryCustomerCheckPagination(objPayload *structs.ObjPayloadGetCustomerCheckLabXaryPagination, objQuery *[]structs.ObjQueryGetCustomerCheckPagination, customerId int, ShopMotherId int) error {
	query := configs.DB1.Table("checks")
	query = query.Select("checks.*, directions.direction_name, queues.que_code, queues.que_datetime, queues.shop_id AS que_shop_id, receipts.rec_code, users.user_fullname, users.user_fullname_en")
	query = query.Joins("LEFT JOIN directions ON directions.id = checks.direction_id")
	query = query.Joins("LEFT JOIN customers ON customers.id = checks.customer_id")
	query = query.Joins("LEFT JOIN shops ON customers.shop_id = shops.id")
	query = query.Joins("LEFT JOIN queues ON checks.queue_id = queues.id")
	query = query.Joins("LEFT JOIN receipts ON receipts.id = checks.receipt_id")
	query = query.Joins("LEFT JOIN users ON users.id = checks.user_id")
	query = query.Where("shops.shop_mother_id = ?", ShopMotherId)
	query = query.Where("checks.chk_type_id = ?", 1)
	query = query.Where("checks.customer_id = ?", customerId)
	if *objPayload.SearchText != "" {
		query = query.Where("checks.chk_code LIKE ? OR checks.chk_name LIKE ?", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%")
	}
	if *objPayload.SearchDate != "" {
		query = query.Where("checks.chk_datetime LIKE '%" + *objPayload.SearchDate + "%'")
	}
	// <pagination
	offset := objPayload.CurrentPage * objPayload.PerPage
	query = query.Limit(objPayload.PerPage)
	query = query.Offset(offset)
	// pagination>
	query = query.Order("checks.chk_datetime DESC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CountHistoryCustomerLabPagination(objPayload *structs.ObjPayloadGetCustomerCheckLabXaryPagination, customerId int, ShopMotherId int) (int64, error) {
	var count int64
	query := configs.DB1.Table("checks")
	query = query.Joins("LEFT JOIN directions ON directions.id = checks.direction_id")
	query = query.Joins("LEFT JOIN customers ON customers.id = checks.customer_id")
	query = query.Joins("LEFT JOIN shops ON customers.shop_id = shops.id")
	query = query.Joins("LEFT JOIN queues ON checks.queue_id = queues.id")
	query = query.Where("shops.shop_mother_id = ?", ShopMotherId)
	query = query.Where("checks.chk_type_id = ?", 2)
	query = query.Where("checks.customer_id = ?", customerId)
	query = query.Where("queues.que_status_id =?", 4)
	if *objPayload.SearchText != "" {
		query = query.Where("checks.chk_code LIKE ? OR checks.chk_name LIKE ?", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%")
	}
	if *objPayload.SearchDate != "" {
		query = query.Where("checks.chk_datetime LIKE '%" + *objPayload.SearchDate + "%'")
	}
	query = query.Count(&count)
	if query.Error != nil {
		return 0, query.Error
	}
	return count, nil
}

func GetHistoryCustomerLabPagination(objPayload *structs.ObjPayloadGetCustomerCheckLabXaryPagination, objQuery *[]structs.ObjQueryGetCustomerCheckPagination, customerId int, ShopMotherId int) error {
	query := configs.DB1.Table("checks")
	query = query.Select("checks.*, directions.direction_name, queues.que_code, queues.que_datetime, queues.shop_id AS que_shop_id, receipts.rec_code, users.user_fullname, users.user_fullname_en")
	query = query.Joins("LEFT JOIN directions ON directions.id = checks.direction_id")
	query = query.Joins("LEFT JOIN customers ON customers.id = checks.customer_id")
	query = query.Joins("LEFT JOIN shops ON customers.shop_id = shops.id")
	query = query.Joins("LEFT JOIN queues ON checks.queue_id = queues.id")
	query = query.Joins("LEFT JOIN receipts ON receipts.id = checks.receipt_id")
	query = query.Joins("LEFT JOIN users ON users.id = checks.user_id")
	query = query.Where("shops.shop_mother_id = ?", objPayload.ShopId)
	query = query.Where("checks.chk_type_id = ?", 2)
	query = query.Where("checks.customer_id = ?", customerId)
	query = query.Where("queues.que_status_id =?", 4)
	if *objPayload.SearchText != "" {
		query = query.Where("checks.chk_code LIKE ? OR checks.chk_name LIKE ?", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%")
	}
	if *objPayload.SearchDate != "" {
		query = query.Where("checks.chk_datetime LIKE '%" + *objPayload.SearchDate + "%'")
	}
	if objPayload.ShopId != nil {
		// *objPayload.ShopId is a valid int
		query = query.Where("shops.id = ?", *objPayload.ShopId)
	}
	// <pagination
	offset := objPayload.CurrentPage * objPayload.PerPage
	query = query.Limit(objPayload.PerPage)
	query = query.Offset(offset)
	// pagination>
	query = query.Order("checks.chk_datetime DESC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CountHistoryCustomerXrayPagination(objPayload *structs.ObjPayloadGetCustomerCheckLabXaryPagination, customerId int, ShopMotherId int) (int64, error) {
	var count int64
	query := configs.DB1.Table("checks")
	query = query.Joins("LEFT JOIN directions ON directions.id = checks.direction_id")
	query = query.Joins("LEFT JOIN customers ON customers.id = checks.customer_id")
	query = query.Joins("LEFT JOIN shops ON customers.shop_id = shops.id")
	query = query.Joins("LEFT JOIN queues ON checks.queue_id = queues.id")
	query = query.Where("shops.shop_mother_id = ?", ShopMotherId)
	query = query.Where("checks.chk_type_id = ?", 3)
	query = query.Where("checks.customer_id = ?", customerId)
	if *objPayload.SearchText != "" {
		query = query.Where("checks.chk_code LIKE ? OR checks.chk_name LIKE ?", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%")
	}
	if *objPayload.SearchDate != "" {
		query = query.Where("checks.chk_datetime LIKE '%" + *objPayload.SearchDate + "%'")
	}
	query = query.Count(&count)
	if query.Error != nil {
		return 0, query.Error
	}
	return count, nil
}

func GetHistoryCustomerXrayPagination(objPayload *structs.ObjPayloadGetCustomerCheckLabXaryPagination, objQuery *[]structs.ObjQueryGetCustomerCheckPagination, customerId int, ShopMotherId int) error {
	query := configs.DB1.Table("checks")
	query = query.Select("checks.*, directions.direction_name, queues.que_code, queues.que_datetime, queues.shop_id AS que_shop_id, receipts.rec_code, users.user_fullname, users.user_fullname_en")
	query = query.Joins("LEFT JOIN directions ON directions.id = checks.direction_id")
	query = query.Joins("LEFT JOIN customers ON customers.id = checks.customer_id")
	query = query.Joins("LEFT JOIN shops ON customers.shop_id = shops.id")
	query = query.Joins("LEFT JOIN queues ON checks.queue_id = queues.id")
	query = query.Joins("LEFT JOIN receipts ON receipts.id = checks.receipt_id")
	query = query.Joins("LEFT JOIN users ON users.id = checks.user_id")
	query = query.Where("shops.shop_mother_id = ?", ShopMotherId)
	query = query.Where("checks.chk_type_id = ?", 3)
	query = query.Where("checks.customer_id = ?", customerId)
	if *objPayload.SearchText != "" {
		query = query.Where("checks.chk_code LIKE ? OR checks.chk_name LIKE ?", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%")
	}
	if *objPayload.SearchDate != "" {
		query = query.Where("checks.chk_datetime LIKE '%" + *objPayload.SearchDate + "%'")
	}
	// <pagination
	offset := objPayload.CurrentPage * objPayload.PerPage
	query = query.Limit(objPayload.PerPage)
	query = query.Offset(offset)
	// pagination>
	query = query.Order("checks.chk_datetime DESC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func AppointmentSearchByCustomerHistory(filter structs.PayloadSearchAppointmentByCustomerHistory, isCount bool, product *[]structs.AppointmentListHistory) (err error) {
	query := configs.DB1.Table("appointments")
	query = query.Select("appointments.*,customers.ctm_fname,customers.ctm_lname,customers.ctm_fname_en,customers.ctm_lname_en,users.user_fullname,users.user_fullname_en")
	query = query.Joins("LEFT JOIN customers ON customers.id = appointments.customer_id")
	query = query.Joins("LEFT JOIN users ON users.id = appointments.user_id")
	query = query.Joins("LEFT JOIN shops ON shops.id = appointments.shop_id")
	query = query.Where("appointments.customer_id = ?", filter.Customer_id)
	query = query.Where("shops.shop_mother_id = ?", filter.ShopMotherId)
	query = query.Where("appointments.ap_is_del = 0")

	if *filter.Search != "" {
		query = query.Where("appointments.ap_topic LIKE '%" + *filter.Search + "%' OR appointments.ap_note LIKE '%" + *filter.Search + "%' ")
	}
	if *filter.Date != "" {
		query = query.Where("appointments.ap_datetime LIKE '%" + *filter.Date + "%'")
	}
	if *filter.Type != "" {
		query = query.Where("appointments.ap_type = ?", *filter.Type)
	}
	if isCount == true {
		offset := filter.ActivePage * filter.PerPage
		query = query.Limit(filter.PerPage)
		query = query.Offset(offset)
	}
	query = query.Order("appointments.ap_datetime DESC")
	query = query.Order("appointments.id DESC")
	if err = query.Scan(&product).Error; err != nil {
		return err
	}
	return nil
}

func CountHistoryDocument(objPayload *structs.ObjPayloadPaginationHistory, customerId int, ShopMotherId int) (int64, error) {
	var count int64
	query := configs.DB1.Table("medical_certs")
	query = query.Joins("LEFT JOIN medical_cert_types on medical_cert_types.id = medical_certs.medical_cert_type_id")
	query = query.Joins("LEFT JOIN users ON users.id = medical_certs.user_id")
	query = query.Joins("LEFT JOIN opds on opds.id = medical_certs.opd_id")
	query = query.Joins("LEFT JOIN queues on queues.id = opds.queue_id")
	query = query.Joins("LEFT JOIN shops ON shops.id = queues.shop_id")
	query = query.Joins("LEFT JOIN customers on customers.id = opds.customer_id")
	query = query.Where("medical_certs.mdc_is_del = ?", 0)
	query = query.Where("shops.shop_mother_id = ?", ShopMotherId)
	query = query.Where("customers.id = ?", customerId)
	if *objPayload.SearchText != "" {
		query = query.Where("customers.ctm_id LIKE ? OR medical_certs.mdc_code LIKE ? OR customers.ctm_fname LIKE ? OR customers.ctm_lname LIKE ? OR CONCAT(customers.ctm_fname,' ',customers.ctm_lname) LIKE ?", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%")
	}
	if *objPayload.SearchDate != "" {
		query = query.Where("medical_certs.mdc_create LIKE '%" + *objPayload.SearchDate + "%'")
	}
	query = query.Count(&count)
	if query.Error != nil {
		return 0, query.Error
	}
	return count, nil
}

func CountCourseHistory(objPayload *structs.ObjPayloadPaginationHistory, customerId int, ShopMotherId int) (int64, error) {
	var count int64
	query := configs.DB1.Table("service_useds")
	query = query.Joins("INNER JOIN queues ON service_useds.queue_id = queues.id")
	query = query.Joins("INNER JOIN receipts ON service_useds.receipt_id = receipts.id")
	query = query.Joins("INNER JOIN services ON service_useds.service_id = services.id")
	query = query.Joins("INNER JOIN users ON service_useds.user_id = users.id")
	query = query.Joins("INNER JOIN shops ON service_useds.shop_id = shops.id")
	query = query.Where("shops.shop_mother_id = ?", ShopMotherId)
	query = query.Where("service_useds.seru_is_active != ? ", 0)
	query = query.Where("services.ser_is_active != ? ", 0)
	query = query.Where("service_useds.customer_id = ?", customerId)
	if *objPayload.SearchText != "" {
		query = query.Where("queues.que_code LIKE ? OR receipts.rec_code LIKE ? OR services.ser_code LIKE ? OR services.ser_name LIKE ?", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%")
	}
	if *objPayload.SearchDate != "" {
		query = query.Where("service_useds.seru_date LIKE '%" + *objPayload.SearchDate + "%'")
	}
	query = query.Count(&count)
	if query.Error != nil {
		return 0, query.Error
	}
	return count, nil
}

func GetCourseHistory(objPayload *structs.ObjPayloadPaginationHistory, objQuery *[]structs.ObjQueryGetCustomerCourseHistory, customerId int, ShopMotherId int) error {
	query := configs.DB1.Table("service_useds")
	query = query.Select("queues.que_code,receipts.rec_code, service_useds.seru_date,service_useds.seru_code,service_useds.seru_name,service_useds.seru_qty,( services.ser_price_total / services.ser_qty ) * service_useds.seru_qty AS ser_price,users.user_fullname,users.user_fullname_en")
	query = query.Joins("INNER JOIN queues ON service_useds.queue_id = queues.id")
	query = query.Joins("INNER JOIN receipts ON service_useds.receipt_id = receipts.id")
	query = query.Joins("INNER JOIN services ON service_useds.service_id = services.id")
	query = query.Joins("INNER JOIN users ON service_useds.user_id = users.id")
	query = query.Joins("INNER JOIN shops ON service_useds.shop_id = shops.id")
	query = query.Where("shops.shop_mother_id = ?", ShopMotherId)
	query = query.Where("service_useds.seru_is_active != ? ", 0)
	query = query.Where("services.ser_is_active != ? ", 0)
	query = query.Where("service_useds.customer_id = ?", customerId)
	if *objPayload.SearchText != "" {
		query = query.Where("queues.que_code LIKE ? OR receipts.rec_code LIKE ? OR services.ser_code LIKE ? OR services.ser_name LIKE ?", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%")
	}
	if *objPayload.SearchDate != "" {
		query = query.Where("service_useds.seru_date LIKE '%" + *objPayload.SearchDate + "%'")
	}
	// <pagination
	offset := objPayload.CurrentPage * objPayload.PerPage
	query = query.Limit(objPayload.PerPage)
	query = query.Offset(offset)
	// pagination>
	query = query.Order("service_useds.seru_date DESC")
	query = query.Order("service_useds.seru_datetime DESC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CountMedicineHistory(objPayload *structs.ObjPayloadPaginationHistory, customerId int, ShopMotherId int) (int64, error) {
	var count int64
	query := configs.DB1.Table("receipt_details")
	query = query.Joins("INNER JOIN receipts ON receipt_details.receipt_id = receipts.id")
	query = query.Joins("INNER JOIN users ON receipts.user_id = users.id")
	query = query.Joins("INNER JOIN shops ON receipts.shop_id = shops.id")
	query = query.Joins("LEFT JOIN queues ON receipt_details.queue_id = queues.id")
	query = query.Where("shops.shop_mother_id = ?", ShopMotherId)
	query = query.Where("receipt_details.recd_is_active = ? ", 1)
	query = query.Where("receipt_details.recd_type_id = ? ", 3)
	query = query.Where("receipts.rec_is_active != ? ", 0)
	query = query.Where("receipts.rec_is_process = ? ", 1)
	query = query.Where("receipts.customer_id = ?", customerId)
	if *objPayload.SearchText != "" {
		query = query.Where("queues.que_code LIKE ? OR receipts.rec_code LIKE ? OR receipt_details.recd_code LIKE ? OR receipt_details.recd_name LIKE ?", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%")
	}
	if *objPayload.SearchDate != "" {
		query = query.Where("receipts.rec_pay_datetime LIKE '%" + *objPayload.SearchDate + "%'")
	}
	query = query.Count(&count)
	if query.Error != nil {
		return 0, query.Error
	}
	return count, nil
}

func GetMedicineHistory(objPayload *structs.ObjPayloadPaginationHistory, objQuery *[]structs.ObjQueryGetCustomerMedicineHistory, customerId int, ShopMotherId int) error {
	query := configs.DB1.Table("receipt_details")
	query = query.Select("IFNULL(queues.que_code,'-') AS que_code,receipts.rec_code,receipt_details.recd_code AS quep_code,receipt_details.recd_name AS quep_name,receipt_details.recd_qty AS quep_qty,receipt_details.recd_total AS quep_total,receipt_details.recd_topical AS quep_topical,receipts.rec_pay_datetime AS que_datetime,users.user_fullname,users.user_fullname_en")
	query = query.Joins("INNER JOIN receipts ON receipt_details.receipt_id = receipts.id")
	query = query.Joins("INNER JOIN users ON receipts.user_id = users.id")
	query = query.Joins("INNER JOIN shops ON receipts.shop_id = shops.id")
	query = query.Joins("LEFT JOIN queues ON receipt_details.queue_id = queues.id")
	query = query.Where("shops.shop_mother_id = ?", ShopMotherId)
	query = query.Where("receipt_details.recd_is_active = ? ", 1)
	query = query.Where("receipt_details.recd_type_id = ? ", 3)
	query = query.Where("receipts.rec_is_active != ? ", 0)
	query = query.Where("receipts.rec_is_process = ? ", 1)
	query = query.Where("receipts.customer_id = ?", customerId)
	if *objPayload.SearchText != "" {
		query = query.Where("queues.que_code LIKE ? OR receipts.rec_code LIKE ? OR receipt_details.recd_code LIKE ? OR receipt_details.recd_name LIKE ?", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%")
	}
	if *objPayload.SearchDate != "" {
		query = query.Where("receipts.rec_pay_datetime LIKE '%" + *objPayload.SearchDate + "%'")
	}
	// <pagination
	offset := objPayload.CurrentPage * objPayload.PerPage
	query = query.Limit(objPayload.PerPage)
	query = query.Offset(offset)
	// pagination>
	query = query.Order("receipts.rec_pay_datetime DESC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetHistoryDocument(objPayload *structs.ObjPayloadPaginationHistory, objQuery *[]structs.ObjQueryGetMedicalCertPagination, customerId int, ShopMotherId int) error {
	query := configs.DB1.Table("medical_certs")
	query = query.Select("medical_certs.*, medical_cert_types.mdct_th, medical_cert_types.mdct_en, users.user_fullname,users.user_fullname_en, customers.ctm_id, customers.ctm_prefix, customers.ctm_fname, customers.ctm_lname, opds.opd_code")
	query = query.Joins("LEFT JOIN medical_cert_types on medical_cert_types.id = medical_certs.medical_cert_type_id")
	query = query.Joins("LEFT JOIN users ON users.id = medical_certs.user_id")
	query = query.Joins("LEFT JOIN opds on opds.id = medical_certs.opd_id")
	query = query.Joins("LEFT JOIN queues on queues.id = opds.queue_id")
	query = query.Joins("LEFT JOIN shops ON shops.id = queues.shop_id")
	query = query.Joins("LEFT JOIN customers on customers.id = opds.customer_id")
	query = query.Where("medical_certs.mdc_is_del = ?", 0)
	query = query.Where("shops.shop_mother_id = ?", ShopMotherId)
	query = query.Where("customers.id = ?", customerId)
	if *objPayload.SearchText != "" {
		query = query.Where("medical_certs.mdc_code LIKE ? OR customers.ctm_id LIKE ? OR customers.ctm_fname LIKE ? OR customers.ctm_lname LIKE ? OR CONCAT(customers.ctm_fname,' ',customers.ctm_lname) LIKE ?", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%", "%"+*objPayload.SearchText+"%")
	}
	if *objPayload.SearchDate != "" {
		query = query.Where("medical_certs.mdc_create LIKE '%" + *objPayload.SearchDate + "%'")
	}
	// <pagination
	offset := objPayload.CurrentPage * objPayload.PerPage
	query = query.Limit(objPayload.PerPage)
	query = query.Offset(offset)
	// pagination>
	query = query.Order("medical_certs.mdc_create DESC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
