package structs

import "time"

type ObjUserVerified struct {
	User_email       string `json:"user_email"`
	User_is_activate int    `json:"user_is_activate"`
	User_update      string `json:"user_update"`
}

type LogUser struct {
	Username   string `json:"username"`
	Log_type   string `json:"log_type"`
	Log_text   string `json:"log_text"`
	Log_create string `json:"log_create"`
}

type VerifiedUser struct {
	ID         int `gorm:"primaryKey"`
	UserId     int
	UvfCode    string
	UvfToken   string
	UvfCreate  string
	UvfExpired string
}

type VerifiedUserFull struct {
	UvfId  int `json:"uvf_id"`
	UserId int `json:"user_id"`
	ProfileUser
	UvfCode    string    `json:"uvf_code"`
	UvfToken   string    `json:"uvf_token"`
	UvfCreate  string    `json:"uvf_create"`
	UvfExpired time.Time `json:"uvf_expired"`
}

type ProfileUser struct {
	ID                int    `json:"id"`
	Username          string `json:"username"`
	User_email        string `json:"user_email"`
	UserFullname      string `json:"user_fullname"`
	UserFullnameEn    string `json:"user_fullname_en"`
	UserAddress       string `json:"user_address"`
	UserAddressEn     string `json:"user_address_en"`
	UserDistrict      string `json:"user_district"`
	UserDistrictEn    string `json:"user_district_en"`
	UserAmphoe        string `json:"user_amphoe"`
	UserAmphoeEn      string `json:"user_amphoe_en"`
	UserProvince      string `json:"user_province"`
	UserProvinceEn    string `json:"user_province_en"`
	UserZipcode       string `json:"user_zipcode"`
	UserZipcodeEn     string `json:"user_zipcode_en"`
	User_type         int    `json:"user_type"`
	User_type_name    string `json:"user_type_name"`
	User_image        string `json:"user_image"`
	User_tel          string `json:"user_tel"`
	UserLicense       string `json:"user_license"`
	NotDisplayQueue   int    `json:"not_display_queue"`
	NotDisplayAppoint int    `json:"not_display_appoint"`
	User_is_activate  int    `json:"user_is_activate"`
	User_update       string `json:"user_update"`
	UserShop
	DefaultShop
	ShopRole
}

type UserShop struct {
	ShopId int `json:"shop_id"`
}

type DefaultShop struct {
	ShopName   string `json:"shop_name"`
	ShopNameEN string `json:"shop_name_en"`
}

type ShopRole struct {
	RoleNameTh string `json:"role_name_th"`
	RoleNameEN string `json:"role_name_en"`
}

type ObjUpdateUser struct {
	UserFullname   string  `json:"user_fullname" binding:"required"`
	UserFullnameEn string  `json:"user_fullname_en"`
	UserAddress    string  `json:"user_address"`
	UserAddressEn  string  `json:"user_address_en"`
	UserDistrict   string  `json:"user_district"`
	UserDistrictEn string  `json:"user_district_en"`
	UserAmphoe     string  `json:"user_amphoe"`
	UserAmphoeEn   string  `json:"user_amphoe_en"`
	UserProvince   string  `json:"user_province"`
	UserProvinceEn string  `json:"user_province_en"`
	UserZipcode    string  `json:"user_zipcode"`
	UserZipcodeEn  string  `json:"user_zipcode_en"`
	UserImage      *string `json:"user_image" binding:"required,omitempty"`
	UserTel        *string `json:"user_tel" binding:"required,omitempty"`
	UserLicense    string  `json:"user_license"`
	UserUpdate     string  `json:"user_update"`
}

type ForgetPwdStamp struct {
	UfpId      int `gorm:"primary"`
	UserId     int
	UfpToken   string
	UfpCreate  string
	UfpExpired time.Time
}

type ForgetPwdUser struct {
	UfpId  int `gorm:"primary"`
	UserId int
	ProfileUser
	UfpToken   string
	UfpCreate  string
	UfpExpired time.Time
}

type ObjChangePwdUser struct {
	Password    string `json:"password"`
	User_update string `json:"user_update"`
}

type ObjUpdateExpireStatus struct {
	UfpIsExpired int
}

type PayloadChangePassword struct {
	OldPassword        string `json:"old_password" binding:"required"`
	NewPassword        string `json:"new_password" binding:"required"`
	ConfirmNewPassword string `json:"confirm_new_password" binding:"required"`
}

type UserSettimeDays struct {
	ID        int    `json:"id"`
	DayNameTh string `json:"day_name_th"`
	DayNameEn string `json:"day_name_en"`
	Check     int    `json:"check"`
	UserId    int    `json:"user_id"`
	DayId     int    `json:"day_id"`
	TimeStart string `json:"time_start"`
	TimeEnd   string `json:"time_end"`
}

type UserSettime struct {
	ID        int    `json:"id"`
	ShopId    int    `json:"shop_id" binding:"required"`
	UserId    int    `json:"user_id" binding:"required"`
	DayId     int    `json:"day_id" binding:"required"`
	TimeStart string `json:"time_start" binding:"omitempty"`
	TimeEnd   string `json:"time_end" binding:"omitempty"`
}

type ObjUpdateUserSettimeList struct {
	TimeStart string `json:"time_start" binding:"omitempty"`
	TimeEnd   string `json:"time_end" binding:"omitempty"`
}

type Days struct {
	ID        int    `json:"id"`
	DayNameTh string `json:"day_name_th"`
	DayNameEn string `json:"day_name_en"`
}

type ObjUpdateUserSettime struct {
	NotDisplayQueue   int    `json:"not_display_queue"`
	NotDisplayAppoint int    `json:"not_display_appoint"`
	UserUpdate        string `json:"user_update"`
}

type ObjPlayloadUserSettime struct {
	NotDisplayQueue   int                `json:"not_display_queue"`
	NotDisplayAppoint int                `json:"not_display_appoint"`
	Settime           *[]UserSettimeDays `json:"settime" binding:"omitempty"`
}

type ObjPayloadSearchSelectUser struct {
	ShopId     int     `json:"shop_id" binding:"required"`
	SearchText *string `json:"search_text" binding:"required"`
}

type ObjResponseSearchSelectUser struct {
	ID             int    `json:"id"`
	UserEmail      string `json:"user_email"`
	UserFullname   string `json:"user_fullname"`
	UserFullnameEn string `json:"user_fullname_en"`
	UserImage      string `json:"user_image"`
}

type ReqCreateUserPopup struct {
	PopupId int `json:"popup_id"`
}

type ObjPayloadUpdateUserId struct {
	Mode   string `json:"mode"`
	UserId int    `json:"user_id"`
}

type ObjUserEmail struct {
	User_email string `json:"user_email"`
}
