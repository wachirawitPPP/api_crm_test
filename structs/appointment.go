package structs

import "time"

// Doctor Table
type PayloadDoctorTable struct {
	Shop_id   int    `json:"shop_id"`
	DateStart string `json:"date_start" binding:"required"`
	DateEnd   string `json:"date_end" binding:"required"`
}

type ObjAPIdList struct {
	Ap_id *[]int `json:"id" binding:"required,omitempty"`
}

type ShopTimeset struct {
	ID                int    `json:"id"`
	ShopID            int    `json:"shop_id"`
	TimesetOpen       string `json:"timeset_open"`
	TimesetClose      string `json:"timeset_close"`
	TimesetRange      int    `json:"timeset_range"`
	TimesetDayId      int    `json:"timeset_day_id"`
	TimesetDayAmount  int    `json:"timeset_day_amount"`
	TimesetRoomId     int    `json:"timeset_room_id"`
	TimesetRoomAmount int    `json:"timeset_room_amount"`
	TimesetCommentId  int    `json:"timeset_comment_id"`
	TimesetSunday     int    `json:"timeset_sunday"`
	TimesetMonday     int    `json:"timeset_monday"`
	TimesetTuesday    int    `json:"timeset_tuesday"`
	TimesetWednesday  int    `json:"timeset_wednesday"`
	TimesetThursday   int    `json:"timeset_thursday"`
	TimesetFriday     int    `json:"timeset_friday"`
	TimesetSaturday   int    `json:"timeset_saturday"`
}

type UserSettimeList struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	DayID     int    `json:"day_id"`
	TimeStart string `json:"time_start"`
	TimeEnd   string `json:"time_end"`
}

type TableTime struct {
	TimeData string      `json:"TimeData"`
	DayDatas []TableDays `json:"day_datas"`
}

type TableDays struct {
	Date    string             `json:"date"`
	Day     string             `json:"day"`
	Day_th  string             `json:"day_th"`
	Members []UserSettimeToDay `json:"members"`
}

type DoctorMember struct {
	ID             int    `json:"id"`
	User_email     string `json:"user_email"`
	UserFullname   string `json:"user_fullname"`
	UserFullnameEn string `json:"user_fullname_en"`
	RoleNameTh     string `json:"role_name_th"`
	RoleNameEn     string `json:"role_name_en"`
	User_image     string `json:"user_image"`
	User_tel       string `json:"user_tel"`
}

type PayloadSearchAppointment struct {
	Shop_id    int     `json:"shop_id"`
	Search     *string `json:"search" binding:"required,omitempty"`
	Date       *string `json:"date" binding:"required,omitempty"`
	Type       *string `json:"type" binding:"required,omitempty"`
	OpdType    *string `json:"opd_type" binding:"required,omitempty"`
	Is_active  *string `json:"is_active" binding:"required,omitempty"`
	ActivePage int     `json:"active_page" binding:"required"`
	PerPage    int     `json:"per_page" binding:"required"`
}

type PayloadCalendarAppointment struct {
	Date_start *string `json:"date_start" binding:"required,omitempty"`
	Date_end   *string `json:"date_end" binding:"required,omitempty"`
}

type ResponseAppointment struct {
	Result_data   []AppointmentList `json:"result_data"`
	Count_of_page int               `json:"count_of_page"`
	Count_all     int               `json:"count_all"`
}

type AppointmentList struct {
	ID               int    `json:"id"`
	ShopID           int    `json:"shop_id"`
	UserID           int    `json:"user_id"`
	User_image       string `json:"user_image"`
	UserFullname     string `json:"user_fullname"`
	UserFullnameEn   string `json:"user_fullname_en"`
	RoleNameTh       string `json:"role_name_th"`
	RoleNameEn       string `json:"role_name_en"`
	CustomerID       int    `json:"customer_id"`
	CustomerFullname string `json:"customer_fullname"`
	CtmId            string `json:"ctm_id"`
	CtmFname         string `json:"ctm_fname"`
	CtmLname         string `json:"ctm_lname"`
	CtmFnameEn       string `json:"ctm_fname_en"`
	CtmLnameEn       string `json:"ctm_lname_en"`
	ApType           int    `json:"ap_type"`
	ApTopic          string `json:"ap_topic"`
	ApTel            string `json:"ap_tel"`
	ApDatetime       string `json:"ap_datetime"`
	ApNote           string `json:"ap_note"`
	ApComment        string `json:"ap_comment"`
	ApColor          string `json:"ap_color"`
	ApConfirm        int    `json:"ap_confirm"`
	ApStatusID       int    `json:"ap_status_id"`
	ApStatusSMS      int    `json:"ap_status_sms"`
	ApStatusLine     int    `json:"ap_status_line"`
	ApSms            string `json:"ap_sms"`
	ApIsGcalendar    int    `json:"ap_is_gcalendar"`
	ApGid            string `json:"ap_gid"`
	ApUserID         int    `json:"ap_user_id"`
	ApIsDel          int    `json:"ap_is_del"`
	ApOpdType        int    `json:"ap_opd_type"`
	ApIsTele         int    `json:"ap_is_tele"`
	ApTeleCode       string `json:"ap_tele_code"`
	ApTeleUrl        string `json:"ap_tele_url"`
	ApCreate         string `json:"ap_create"`
	ApUpdate         string `json:"ap_update"`
	ShopName         string `json:"shop_name"`
}

type AppointmentDetail struct {
	ID               int    `json:"id"`
	ShopID           int    `json:"shop_id"`
	UserID           int    `json:"user_id"`
	UserFullname     string `json:"user_fullname"`
	UserFullnameEn   string `json:"user_fullname_en"`
	RoleNameTh       string `json:"role_name_th"`
	RoleNameEn       string `json:"role_name_en"`
	CustomerID       *int   `json:"customer_id"`
	CtmId            string `json:"ctm_id"`
	CustomerFullname string `json:"customer_fullname"`
	CtmFnameEn       string `json:"ctm_fname_en"`
	CtmLnameEn       string `json:"ctm_lname_en"`
	Ctm_birthdate    string `json:"ctm_birthdate"`
	Ctm_tel          string `json:"ctm_tel"`
	Ctm_tel_2        string `json:"ctm_tel_2"`
	ApType           int    `json:"ap_type"`
	ApTopic          string `json:"ap_topic"`
	ApTel            string `json:"ap_tel"`
	ApDatetime       string `json:"ap_datetime"`
	ApNote           string `json:"ap_note"`
	ApComment        string `json:"ap_comment"`
	ApColor          string `json:"ap_color"`
	ApConfirm        int    `json:"ap_confirm"`
	ApStatusID       int    `json:"ap_status_id"`
	ApStatusSMS      int    `json:"ap_status_sms"`
	ApStatusLine     int    `json:"ap_status_line"`
	ApSms            string `json:"ap_sms"`
	ApIsGcalendar    int    `json:"ap_is_gcalendar"`
	ApGid            string `json:"ap_gid"`
	ApUserID         int    `json:"ap_user_id"`
	ApIsDel          int    `json:"ap_is_del"`
	ApOpdType        int    `json:"ap_opd_type"`
	ApIsTele         int    `json:"ap_is_tele"`
	ApTeleCode       string `json:"ap_tele_code"`
	ApTeleUrl        string `json:"ap_tele_url"`
	ApCreate         string `json:"ap_create"`
	ApUpdate         string `json:"ap_update"`
}

type ObjPayloadCreateAppointment struct {
	UserID     int    `json:"user_id" binding:"required"`
	ApTopic    string `json:"ap_topic"`
	ApDatetime string `json:"ap_datetime"`
}

type ApDatetime struct {
	Datetime string `json:"datetime"`
}

type ObjPayloadUpdateAppointment struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id" binding:"required"`
	CustomerID       *int   `json:"customer_id" binding:"required,omitempty"`
	CustomerFullname string `json:"customer_fullname"`
	ApType           int    `json:"ap_type"`
	ApTopic          string `json:"ap_topic"`
	ApTel            string `json:"ap_tel"`
	ApDatetime       string `json:"ap_datetime"`
	ApNote           string `json:"ap_note"`
	ApComment        string `json:"ap_comment"`
	ApColor          string `json:"ap_color"`
	ApConfirm        int    `json:"ap_confirm"`
	ApStatusID       int    `json:"ap_status_id"`
	ApStatusSMS      int    `json:"ap_status_sms"`
	ApStatusLine     int    `json:"ap_status_line"`
	ApSms            string `json:"ap_sms"`
	ApIsGcalendar    int    `json:"ap_is_gcalendar"`
	ApGid            string `json:"ap_gid"`
	ApOpdType        int    `json:"ap_opd_type"`
	ApIsTele         int    `json:"ap_is_tele"`
	ApUpdate         string `json:"ap_update"`
	TagSelected      []int  `json:"tag_selected"`
}

type ObjAppointmentTag struct {
	AppointmentId int `json:"appointment_id"`
	TagId         int `json:"tag_id"`
	// join
	TagName   string `json:"tag_name"`
	TagTypeTh string `json:"tag_type_th"`
	TagTypeEn string `json:"tag_type_en"`
}

type AppointmentAction struct {
	ID               int    `json:"id"`
	ShopID           int    `json:"shop_id"`
	UserID           int    `json:"user_id"`
	CustomerID       int    `json:"customer_id" gorm:"-"`
	CustomerFullname string `json:"customer_fullname"`
	ApType           int    `json:"ap_type"`
	ApTopic          string `json:"ap_topic"`
	ApTel            string `json:"ap_tel"`
	ApDatetime       string `json:"ap_datetime"`
	ApNote           string `json:"ap_note"`
	ApComment        string `json:"ap_comment"`
	ApColor          string `json:"ap_color"`
	ApConfirm        int    `json:"ap_confirm"`
	ApStatusID       int    `json:"ap_status_id"`
	ApStatusSMS      int    `json:"ap_status_sms"`
	ApStatusLine     int    `json:"ap_status_line"`
	ApSms            string `json:"ap_sms"`
	ApIsGcalendar    int    `json:"ap_is_gcalendar"`
	ApGid            string `json:"ap_gid"`
	ApUserID         int    `json:"ap_user_id"`
	ApIsDel          int    `json:"ap_is_del"`
	ApOpdType        int    `json:"ap_opd_type"`
	ApIsTele         int    `json:"ap_is_tele"`
	ApTeleCode       string `json:"ap_tele_code"`
	ApTeleUrl        string `json:"ap_tele_url"`
	ApCreate         string `json:"ap_create"`
	ApUpdate         string `json:"ap_update"`
}

type ObjPayloadUpdateAppointmentStatus struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id" binding:"required"`
	CustomerID *int   `json:"customer_id" binding:"required,omitempty"`
	ApConfirm  int    `json:"ap_confirm"`
	ApStatusID int    `json:"ap_status_id"`
	ApOpdType  int    `json:"ap_opd_type"`
	ApUpdate   string `json:"ap_update"`
}

type AppointmentStatusAction struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	CustomerID int    `json:"customer_id" gorm:"-"`
	ApConfirm  int    `json:"ap_confirm"`
	ApStatusID int    `json:"ap_status_id"`
	ApOpdType  int    `json:"ap_opd_type"`
	ApUpdate   string `json:"ap_update"`
}

type AppointmentTags struct {
	ID            int `json:"id"`
	AppointmentId int `json:"appointment_id"`
	TagId         int `json:"tag_id"`
}

type ObjResponseSearchAppointmentDoctor struct {
	ID             int    `json:"id"`
	UserEmail      string `json:"user_email"`
	UserFullname   string `json:"user_fullname"`
	UserFullnameEn string `json:"user_fullname_en"`
	RoleNameTh     string `json:"role_name_th"`
	RoleNameEn     string `json:"role_name_en"`
}

type AppointmentUser struct {
	Id             int    `json:"id"`
	UserEmail      string `json:"user_email"`
	UserFullname   string `json:"user_fullname"`
	UserFullnameEn string `json:"user_fullname_en"`
	UserTel        string `json:"user_tel"`
}

type AppointmentCustomer struct {
	Id           int    `json:"id"`
	CtmCitizenId string `json:"ctm_citizen_id"`
	CtmFname     string `json:"ctm_fname"`
	CtmLname     string `json:"ctm_lname"`
	CtmFnameEn   string `json:"ctm_fname_en"`
	CtmLnameEn   string `json:"ctm_lname_en"`
	CtmTel       string `json:"ctm_tel"`
	CtmEmail     string `json:"ctm_email"`
}

type AppointmentTopic struct {
	Id      int    `json:"id"`
	TopicTH string `json:"topic_th"`
	TopicEN string `json:"topic_en"`
}

type AppointmentTag struct {
	Id      int    `json:"id"`
	TagName string `json:"tag_name"`
}

type ShopConfig struct {
	Id      int    `json:"id"`
	TagName string `json:"tag_name"`
}

type CalendarView struct {
	Id              int    `json:"id"`
	Title           string `json:"title"`
	Start           string `json:"start"`
	End             string `json:"end"`
	BackgroundColor string `json:"backgroundColor"`
	BorderColor     string `json:"borderColor"`
	Color           string `json:"color"`
}

type LogAppointment struct {
	Username   string `json:"username"`
	Log_type   string `json:"log_type"`
	Log_text   string `json:"log_text"`
	Log_create string `json:"log_create"`
}

type AppointmentEmail struct {
	ID            int    `json:"id"`
	User_id       int    `json:"user_id"`
	Customer_id   int    `json:"customer_id"`
	Shop_id       int    `json:"shop_id"`
	Shop_email    string `json:"shop_email"`
	Shop_gid      string `json:"shop_gid"`
	Shop_gc_token string `json:"shop_gc_token"`
	User_email    string `json:"user_email"`
	Ctm_email     string `json:"ctm_email"`
}

type UserSettimeToDay struct {
	ID               int    `json:"id"`
	Shop_id          int    `json:"shop_id"`
	User_id          int    `json:"user_id"`
	User_image       string `json:"user_image"`
	User_fullname    string `json:"user_fullname"`
	User_fullname_en string `json:"user_fullname_en"`
	Role_name_th     string `json:"role_name_th"`
	Role_name_en     string `json:"role_name_en"`
	Day_id           int    `json:"day_id"`
	Day_name_th      string `json:"day_name_th"`
	Time_start       string `json:"time_start"`
	Time_end         string `json:"time_end"`
}

type ShopHoliday struct {
	ID              int    `json:"id"`
	Shop_id         int    `json:"shop_id"`
	Holiday_type_id int    `json:"holiday_type_id"`
	Holiday_name    string `json:"holiday_name"`
	Holiday_date    string `json:"holiday_date"`
	Holiday_is_del  int    `json:"holiday_is_del"`
}

type ShopGmail struct {
	ID            int    `json:"id"`
	Shop_gmail    string `json:"shop_gmail"`
	Shop_gtoken   string `json:"shop_gtoken"`
	Shop_gretoken string `json:"shop_gretoken"`
	Shop_gid      string `json:"shop_gid"`
	Shop_gc_token string `json:"shop_gc_token"`
}
type CalendarToken struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	TokenType    string    `json:"token_type"`
	Expiry       time.Time `json:"expiry"`
}
