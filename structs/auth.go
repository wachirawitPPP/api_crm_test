package structs

type PayloadLogin struct {
	Line_id    string `json:"line_id" binding:"required"`
	Line_name  string `json:"line_name" binding:"required"`
	Line_email string `json:"line_email" binding:"required"`
	Citizen_id string `json:"citizen_id" binding:"required"`
	Tel        string `json:"tel" binding:"required"`
}
type PayloadLoginTest struct {
	Line_id    string `json:"line_id" binding:"required"`
	Line_name  string `json:"line_name" binding:"required"`
	Line_email string `json:"line_email" binding:"required"`
	Citizen_id string `json:"citizen_id" binding:"required"`
	Tel        string `json:"tel" binding:"required"`
}

type PayloadLoginOnline struct {
	Co_email    string `json:"co_email" binding:"required"`
	Co_password string `json:"co_password" binding:"required"`
}

type PayloadLoginTel struct {
	Co_tel string `json:"co_tel" binding:"required"`
}

type PayloadOtp struct {
	Line_id string `json:"line_id" binding:"required"`
	Tel     string `json:"tel" binding:"required"`
	Otp     string `json:"otp" binding:"required"`
}

type PayloadChecklineid struct {
	Line_id string `json:"line_id" binding:"required"`
}

type CustomerOnline struct {
	ID            int    `json:"id"`
	Co_line_name  string `json:"co_line_name"`
	Co_line_email string `json:"co_line_email"`
	Co_line_id    string `json:"co_line_id"`
	Co_citizen_id string `json:"co_citizen_id"`
	Co_tel        string `json:"co_tel"`
	Co_otp        string `json:"co_otp"`
	Co_otp_key    string `json:"co_otp_key"`
	Co_otp_expire string `json:"co_otp_expire"`
	Co_is_active  int    `json:"co_is_active"`
	Co_is_del     int    `json:"co_is_del"`
	Co_create     string `json:"co_create"`
	Co_update     string `json:"co_update"`
	Co_email      string `json:"co_email"`
	Co_password   string `json:"co_password"`
	Co_prefix     string `json:"co_prefix"`
	Co_fname      string `json:"co_fname"`
	Co_lname      string `json:"co_lname"`
	Co_gender     string `json:"co_gender"`
	Co_birthdate  string `json:"co_birthdate"`
	Co_address    string `json:"co_address"`
	Co_district   string `json:"co_district"`
	Co_amphoe     string `json:"co_amphoe"`
	Co_province   string `json:"co_province"`
	Co_zipcode    string `json:"co_zipcode"`
}

type ResponseOauth struct {
	AccessToken string `json:"access_token"`
}

type ObjLogCustomerLogin struct {
	Line_id        string `json:"line_id"`
	Log_ip_address string `json:"log_ip_address"`
	Log_browser    string `json:"log_browser"`
	Log_text       string `json:"log_text"`
	Log_create     string `json:"log_create"`
}

type ShopAuth struct {
	Id           int    `json:"id"`
	UserId       int    `json:"user_id"`
	ShopId       int    `json:"shop_id"`
	ShopMotherId int    `json:"shop_mother_id"`
	SaToken      string `json:"sa_token"`
	IsExpired    int    `json:"is_expired"`
	CreateDate   string `json:"create_date"`
	Shop
}

type Shop struct {
	ID              int    `json:"id"`
	ShopCode        string `json:"shop_code"`
	ShopName        string `json:"shop_name"`
	ShopNatureID    int    `json:"shop_nature_id"`
	ShopMotherId    int    `json:"shop_mother_id"`
	ShopLang        int    `json:"shop_lang"`
	Currency_symbol string `json:"currency_symbol"`
}

type PayloadRegis struct {
	// Co_citizen_id string `json:"co_citizen_id" binding:"required"`
	Co_tel      string `json:"co_tel" binding:"required"`
	Co_password string `json:"co_password" binding:"required"`
	Co_email    string `json:"co_email" binding:"required"`
	Co_prefix   string `json:"co_prefix" binding:"required"`
	Co_fname    string `json:"co_fname" binding:"required"`
	Co_lname    string `json:"co_lname" binding:"required"`
}

type PayloadExamedRegis struct {
	Co_citizen_id string `json:"co_citizen_id" binding:"required"`
	Co_tel        string `json:"co_tel" binding:"required"`
	Co_email      string `json:"co_email" binding:"required"`
	Co_prefix     string `json:"co_prefix" binding:"required"`
	Co_fname      string `json:"co_fname" binding:"required"`
	Co_lname      string `json:"co_lname" binding:"required"`
}

type PayloadRegisTelOnly struct {
	Co_tel string `json:"co_tel" binding:"required"`
}

type PayloadOtpRegis struct {
	// Co_citizen_id string `json:"co_citizen_id" binding:"required"`
	Co_email string `json:"co_email" binding:"required"`
	Co_tel   string `json:"co_tel" binding:"required"`
	Otp      string `json:"otp" binding:"required"`
}

type PayloadOtpTelOnlyRegis struct {
	Co_tel string `json:"co_tel" binding:"required"`
	Otp    string `json:"otp" binding:"required"`
}

type CustomerOnlineUpdate struct {
	// Co_citizen_id string `json:"co_citizen_id"`
	Co_otp        string `json:"co_otp"`
	Co_otp_key    string `json:"co_otp_key"`
	Co_otp_expire string `json:"co_otp_expire"`
	Co_update     string `json:"co_update"`
}
type CustomerOnlineOtpUpdate struct {
	// Co_citizen_id string `json:"co_citizen_id"`
	Co_otp        string `json:"co_otp"`
	Co_otp_key    string `json:"co_otp_key"`
	Co_otp_expire string `json:"co_otp_expire"`
	Co_update     string `json:"co_update"`
}

type CustomerOnlineUpdateTel struct {
	Co_tel        string `json:"co_tel"`
	Co_otp        string `json:"co_otp"`
	Co_otp_expire string `json:"co_otp_expire"`
	Co_update     string `json:"co_update"`
}

type CustomerOnlineUpdateSync struct {
	Co_citizen_id string `json:"co_citizen_id"`
	Co_otp        string `json:"co_otp"`
	Co_otp_key    string `json:"co_otp_key"`
	Co_otp_expire string `json:"co_otp_expire"`
	Co_update     string `json:"co_update"`
}

type PayloadSync struct {
	Co_citizen_id string `json:"co_citizen_id" binding:"required"`
}

type PayloadOtpSync struct {
	Cid string `json:"cid" binding:"required"`
	Otp string `json:"otp" binding:"required"`
}
type PayloadOtpSyncLockShop struct {
	Cid string `json:"cid" binding:"required"`
	Otp string `json:"otp" binding:"required"`
	Shop_Id int `json:"shop_id" binding:"required"`
}

type PayloadSyncAccountWithEmail struct {
	Line_id    string `json:"line_id" binding:"required"`
	Co_email   string `json:"co_email" binding:"required"`
	Line_email string `json:"line_email" binding:"required"`
	Line_name  string `json:"line_name" binding:"required"`
}

type SyncAccountWithEmail struct {
	Co_Line_id    string `json:"co_line_id" binding:"required"`
	Co_update     string `json:"co_update"`
	Co_Line_email string `json:"co_line_email" binding:"required"`
	Co_Line_name  string `json:"co_line_name" binding:"required"`
}
type PayloadEmailVerify struct {
	Co_email   string `json:"co_email" binding:"required"`
	Co_otp     string `json:"co_otp" binding:"required"`
	Co_otp_key string `json:"co_otp_key" binding:"required"`
}
type PayloadResendEmailOtp struct {
	Co_email string `json:"co_email" binding:"required"`
}

type PayloadLoginHealthSurvey struct {
	Co_citizen_id string `json:"co_citizen_id" binding:"required"`
}
type PayloadVerifyHealthSurvey struct {
	Co_citizen_id string `json:"co_citizen_id" binding:"required"`
	Co_otp         string `json:"co_otp" binding:"required"`
	Shop_id 	 int    `json:"shop_id" binding:"required"`
}
type TestRemoveAccountWithLineId struct {
 // Co_line_id    string json:"line_id" binding:"required"
 Co_line_name  string `json:"co_line_name"`
 Co_line_email string `json:"co_line_email"`
 Co_line_id    string `json:"co_line_id"`
 Co_citizen_id string `json:"co_citizen_id"`
}
