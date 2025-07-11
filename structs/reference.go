package structs

type ObjShopStore struct {
	ID            int    `json:"id"`
	ShopId        int    `json:"shop_id"`
	ShopCode      string `json:"shop_code"`
	ShopName      string `json:"shop_name"`
	AccountCodeId int    `json:"account_code_id"`
	SsTypeId      int    `json:"ss_type_id"`
	SsName        string `json:"ss_name"`
	SsIsOver      int    `json:"ss_is_over"`
	SsIsActive    int    `json:"ss_is_active"`
	SsCreate      string `json:"ss_create"`
	SsUpdate      string `json:"ss_update"`
}

// docter picker
type ObjPayloadSearchDoctorPicker struct {
	ShopId     int     `json:"shop_id" binding:"required"`
	SearchText *string `json:"search_text" binding:"required"`
}

type ObjResponseSearchDoctorPicker struct {
	ID           int    `json:"id"`
	UserEmail    string `json:"user_email"`
	UserFullname string `json:"user_fullname"`
}

// user picker
type ObjPayloadSearchUserPicker struct {
	ShopId     int     `json:"shop_id" binding:"required"`
	SearchText *string `json:"search_text" binding:"required"`
}

type ObjResponseSearchUserPicker struct {
	ID             int    `json:"id"`
	UserEmail      string `json:"user_email"`
	UserFullname   string `json:"user_fullname"`
	UserFullnameEn string `json:"user_fullname_en"`
}
