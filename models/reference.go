package models

import (
	"linecrmapi/configs"
	"linecrmapi/structs"
	// "gorm.io/gorm"
)

// global2

// Generic function for simple table queries
func getSimpleTableData[T any](tableName string, result *[]T) error {
	return configs.DB1.Table(tableName).Find(result).Error
}

// GetRightTreatment uses the generic function
func GetRightTreatment(objResponse *[]RefRightTreatment) error {
	return getSimpleTableData("ref_right_treatments", objResponse)
}

// GetTagType uses the generic function
func GetTagType(objResponse *[]TagType) error {
	return getSimpleTableData("tag_types", objResponse)
}

func GetMedicalCertType(objResponse *[]MedicalCertType) error {
	query := configs.DB1.Table("medical_cert_types")
	query = query.Find(&objResponse)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetRole(objResponse *[]Role) error {
	query := configs.DB1.Table("roles")
	query = query.Find(&objResponse)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

// by shopId

func GetShopTagById(shopId int, tagTypeIds []int, objResponse *[]Tag) error {
	return configs.DB1.Table("tags").
		Where(`
			tag_is_del = ? AND
			shop_id = ? AND
			tags.tag_type_id IN ?
		`,
			0,
			shopId,
			tagTypeIds,
		).
		Find(objResponse).Error
}

func GetCustomerGroupById(shopId int, objResponse *[]CustomerGroup) error {
	query := configs.DB1.Table("customer_groups")
	query = query.Where("cg_is_active = ?", 1)
	query = query.Where("cg_is_del = ?", 0)
	query = query.Where("shop_id = ?", shopId)
	query = query.Find(&objResponse)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetShopStoreById(shopId int, objResponse *[]ShopStore) error {
	query := configs.DB1.Table("shop_stores")
	query = query.Where("ss_is_active = ?", 1)
	query = query.Where("shop_id = ?", shopId)
	query = query.Find(&objResponse)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetShopStoreSelfById(shopId int, objResponse *[]structs.ObjShopStore) error {
	query := configs.DB1.Table("shop_stores")
	query = query.Select("shop_stores.*, shops.shop_code, shops.shop_name")
	query = query.Joins("LEFT JOIN shops ON shops.id = shop_stores.shop_id")
	query = query.Where("ss_is_active = ?", 1)
	query = query.Where("shop_id = ?", shopId)
	query = query.Find(&objResponse)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetShopStoreBranchById(shopMotherId int, objResponse *[]structs.ObjShopStore) error {
	query := configs.DB1.Table("shop_stores")
	query = query.Select("shop_stores.*, shops.shop_code, shops.shop_name")
	query = query.Joins("LEFT JOIN shops ON shops.id = shop_stores.shop_id")
	query = query.Where("shop_stores.ss_is_active = ?", 1)
	query = query.Where("shops.shop_mother_id = ?", shopMotherId)
	query = query.Find(&objResponse)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetCategoryById(shopId int, objResponse *[]Categorys) error {
	query := configs.DB1.Table("categorys")
	query = query.Where("category_is_del = ?", 0)
	query = query.Where("shop_id = ?", shopId)
	query = query.Find(&objResponse)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetCategoryTypeById(shopId int, category_type_id int, objResponse *[]Categorys) error {
	query := configs.DB1.Table("categorys")
	query = query.Where("category_is_del = ?", 0)
	query = query.Where("shop_id = ?", shopId)
	query = query.Where("category_type_id = ?", category_type_id)
	query = query.Find(&objResponse)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetShopBranchById(shopId int, objResponse *[]Shop) error {
	query := configs.DB1.Table("shops")
	query = query.Where("shop_mother_id = ?", shopId)
	query = query.Find(&objResponse)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetRoomById(shopId int, roomTypeId int, objResponse *[]Room) error {
	query := configs.DB1.Table("rooms")
	query = query.Select("rooms.*,COUNT(beds.room_id) AS count_bed ")
	query = query.Joins("JOIN beds ON beds.room_id = rooms.id ")
	query = query.Where("rooms.room_is_del = ?", 0)
	query = query.Where("rooms.shop_id = ? AND beds.bed_is_del = 0", shopId)
	if roomTypeId != 0 {
		query = query.Where("room_type_id = ?", roomTypeId)
	}
	query = query.Group("rooms.id")
	query = query.Find(&objResponse)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetBedById(roomId int, objResponse *[]Bed, bedLock int) error {
	query := configs.DB1.Table("beds")
	query = query.Select("beds.*, queues.id AS queue_id,queues.que_status_id")
	query = query.Joins("LEFT JOIN queues ON queues.bed_id = beds.id AND queues.que_status_id < 4")
	query = query.Where("beds.room_id = ?", roomId)
	query = query.Where("beds.bed_is_del = ? AND beds.bed_is_active = ?", 0, 1)
	if bedLock >= 0 {
		query = query.Where("beds.bed_lock = ?", bedLock)
	}
	query = query.Group("beds.id")
	query = query.Order("beds.bed_code ASC")
	query = query.Find(&objResponse)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetDirectionById_(shopId int, directionTypeIds []int, objResponse *[]Direction) error {
	query := configs.DB1.Table("directions")
	query = query.Where("direction_is_active = ?", 1)
	query = query.Where("direction_is_del = ?", 0)
	query = query.Where("shop_id = ?", shopId)
	query = query.Where("direction_type_id IN ?", directionTypeIds)
	query = query.Find(&objResponse)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetAccountListById(shopId int, objResponse *[]AccountList) error {
	query := configs.DB1.Table("account_lists")
	query = query.Where("acl_is_del = ?", 0)
	query = query.Where("shop_id = ?", shopId)
	query = query.Find(&objResponse)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetShopAccountCodeById(shopId int, objResponse *[]AccountCode) error {
	query := configs.DB1.Table("account_codes")
	query = query.Where("account_codes.acc_is_del = ?", 0)
	query = query.Where("account_codes.shop_id = ?", shopId)
	// query = query.Where("account_codes.account_type_id != 1 && account_codes.account_type_id != 4")
	query = query.Find(&objResponse)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetDepartmentById(shopId int, objResponse *[]RefDepartment) error {
	query := configs.DB1.Table("ref_departments")
	query = query.Where("ref_departments.dpm_is_del = ?", 0)
	query = query.Where("ref_departments.shop_id = ?", shopId)
	query = query.Find(&objResponse)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

// docter picker
func SearchDoctorPicker(objPayload *structs.ObjPayloadSearchDoctorPicker, objQuery *[]structs.ObjResponseSearchDoctorPicker) error {
	query := configs.DB1.Table("user_shops").
		Select(`
			users.id,
			users.user_email,
			users.user_fullname
		`).
		Joins("LEFT JOIN users ON users.id = user_shops.user_id").
		Where(`
			user_shops.shop_id = ? AND
			users.user_is_active = ?
		`,
			objPayload.ShopId,
			1,
		)

	// Apply search filter if provided
	if objPayload.SearchText != nil && *objPayload.SearchText != "" {
		searchTerm := "%" + *objPayload.SearchText + "%"
		query = query.Where(`
			users.user_email LIKE ? OR 
			users.user_fullname LIKE ? OR 
			users.user_tel LIKE ?
		`,
			searchTerm,
			searchTerm,
			searchTerm,
		)
	}

	return query.Order("users.user_fullname ASC").
		Find(objQuery).Error
}

// user picker
func SearchUserPicker(objPayload *structs.ObjPayloadSearchUserPicker, objQuery *[]structs.ObjResponseSearchUserPicker) error {
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

func GetAllUsers(ShopId int, objQuery *[]structs.ObjResponseSearchUserPicker) error {
	query := configs.DB1.Table("user_shops")
	query = query.Select("users.id, users.user_email, users.user_fullname, users.user_fullname_en")
	query = query.Joins("LEFT JOIN users ON users.id = user_shops.user_id")
	query = query.Where("user_shops.shop_id = ?", ShopId)
	query = query.Where("users.user_is_active = ?", 1)
	query = query.Order("users.user_fullname ASC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetDoctorUsers(ShopId int, objQuery *[]structs.ObjResponseSearchUserPicker) error {
	query := configs.DB1.Table("user_shops")
	query = query.Select("users.id, users.user_email, users.user_fullname, users.user_fullname_en")
	query = query.Joins("LEFT JOIN users ON users.id = user_shops.user_id")
	query = query.Where("user_shops.shop_id = ?", ShopId)
	query = query.Where("users.user_is_active = ?", 1)
	query = query.Order("users.user_fullname ASC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetDoctorsUsers(ShopId int, objQuery *[]structs.ObjResponseSearchAppointmentDoctor) error {
	return configs.DB1.Table("user_shops").
		Select(`
			users.id,
			users.user_image,
			users.user_email,
			users.user_fullname,
			users.user_fullname_en,
			shop_roles.role_name_th,
			shop_roles.role_name_en
		`).
		Joins(`
			JOIN users ON users.id = user_shops.user_id
			JOIN shop_roles ON shop_roles.id = user_shops.shop_role_id
		`).
		Where(`
			user_shops.shop_id = ? AND
			shop_roles.role_id IN ? AND
			users.user_is_active = ? 
		`,
			ShopId,
			[]int{1, 2, 6, 8},
			1,
		).Where("user_shops.us_is_active != 0 AND user_shops.us_invite != 3").
		Order("users.user_fullname ASC").
		Find(objQuery).Error
}

func GetDoctorsUsersSearch(userId int, ShopId int, objQuery *[]structs.ObjResponseSearchAppointmentDoctor) error {
	query := configs.DB1.Table("user_shops")
	query = query.Select(`
		users.id,
		users.user_image,
		users.user_email,
		users.user_fullname,
		users.user_fullname_en,
		shop_roles.role_name_th,
		shop_roles.role_name_en
	`)
	query = query.Joins(`
		JOIN users ON users.id = user_shops.user_id
		JOIN shop_roles ON shop_roles.id = user_shops.shop_role_id
	`)
	query = query.Where(`
		user_shops.shop_id = ? AND
		shop_roles.role_id IN ? AND
		users.user_is_active = ? AND
		user_shops.us_is_active = 1`,
		ShopId,
		[]int{1, 2, 6, 8},
		1,
	)
	query = query.Where("user_shops.us_is_active != 0 AND user_shops.us_invite != 3")
	if userId > 0 {
		query = query.Where("users.id = ?", userId)
	}
	query = query.Order("users.user_fullname ASC")
	return query.Find(&objQuery).Error
}

func GetFeeComUsers(ShopId int, objQuery *[]structs.ObjResponseSearchAppointmentDoctor) error {
	query := configs.DB1.Table("user_shops")
	query = query.Select("users.id, users.user_email, users.user_fullname, users.user_fullname_en,shop_roles.role_name_th,shop_roles.role_name_en")
	query = query.Joins("JOIN users ON users.id = user_shops.user_id")
	query = query.Joins("JOIN shop_roles ON shop_roles.id = user_shops.shop_role_id")
	query = query.Where("user_shops.shop_id = ?", ShopId)
	// query = query.Where("shop_roles.role_id IN ?", []int{1, 2, 6, 8})
	query = query.Where("users.user_is_active = ?", 1)
	query = query.Order("users.user_fullname ASC")
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
