package structs

type CustomerConfigSendmail struct {
	Id                    int    `json:"id"`
	DocNo                 string `json:"doc_no"`
	Customer_id           int    `json:"customer_id"`
	Ctm_subscribe_receipt int    `json:"ctm_subscribe_receipt"`
	Ctm_subscribe_cert    int    `json:"ctm_subscribe_cert"`
	Ctm_subscribe_lab     int    `json:"ctm_subscribe_lab"`
	Ctm_subscribe_opd     int    `json:"ctm_subscribe_opd"`
	Ctm_subscribe_appoint int    `json:"ctm_subscribe_appoint"`
	Ctm_email             string `json:"ctm_email"`
	Shop_id               int    `json:"shop_id"`
	Ctm_id                string `json:"ctm_id"`
	Ctm_fname             string `json:"ctm_fname"`
	Ctm_lname             string `json:"ctm_lname"`
	Ctm_fname_en          string `json:"ctm_fname_en"`
	Ctm_lname_en          string `json:"ctm_lname_en"`
	Ctm_birthdate         string `json:"ctm_birthdate"`
	Shop_name             string `json:"shop_name"`
	Shop_name_en          string `json:"shop_name_en"`
	Shop_email            string `json:"shop_email"`
	Shop_lang             string `json:"shop_lang"` // 1 = TH, 2 = EN
	Email_link            string `json:"link"`
}

type MedicalCertDetailEmail struct {
	Mdct_th              string  `json:"mdct_th"`
	Mdct_en              string  `json:"mdct_en"`
	Mdct_group_id        int     `json:"mdct_group_id"`
	Id                   int     `json:"id"`
	Medical_cert_type_id int     `json:"medical_cert_type_id"`
	User_id              int     `json:"user_id"`
	User_fullname        string  `json:"user_fullname"`
	User_fullname_en     string  `json:"user_fullname_en"`
	User_license         *string `json:"user_license"`
	Opd_id               int     `json:"opd_id"`
	Mdc_code             string  `json:"mdc_code"`
	Mdc_is_print         int     `json:"mdc_is_print"`
	Mdc_create           string  `json:"mdc_create"`
	Mdc_update           string  `json:"mdc_update"`
	CustomerConfigSendmail
}

type OpdEmail struct {
	ID           int    `json:"id"`
	UserId       int    `json:"user_id"`
	QueueId      int    `json:"queue_id"`
	QueCode      string `json:"que_code"`
	OpdCode      string `json:"opd_code"`
	OpdDate      string `json:"opd_date"`
	UserFullname string `json:"user_fullname"`
	CustomerConfigSendmail
}

type AppointEmail struct {
	ID            int    `json:"id"`
	ShopID        int    `json:"shop_id"`
	ApType        int    `json:"ap_type"`
	ApTopic       string `json:"ap_topic"`
	ApTel         string `json:"ap_tel"`
	ApDatetime    string `json:"ap_datetime"`
	ApNote        string `json:"ap_note"`
	ApComment     string `json:"ap_comment"`
	ApColor       string `json:"ap_color"`
	ApConfirm     int    `json:"ap_confirm"`
	ApStatusID    int    `json:"ap_status_id"`
	ApStatusSMS   int    `json:"ap_status_sms"`
	ApStatusLine  int    `json:"ap_status_line"`
	ApSms         string `json:"ap_sms"`
	ApIsGcalendar int    `json:"ap_is_gcalendar"`
	ApGid         string `json:"ap_gid"`
	ApUserID      int    `json:"ap_user_id"`
	ApIsDel       int    `json:"ap_is_del"`
	ApCreate      string `json:"ap_create"`
	ApUpdate      string `json:"ap_update"`
	CustomerConfigSendmail
}

type TaxinvoiceEmail struct {
	Id                int     `json:"id"`
	Receipt_id        int     `json:"receipt_id"`
	Rec_code          string  `json:"rec_code"`
	Tax_code          string  `json:"tax_code"`
	Tax_fullname      string  `json:"tax_fullname"`
	Tax_tel           string  `json:"tax_tel"`
	Tax_email         string  `json:"tax_email"`
	Tax_address       string  `json:"tax_address"`
	Tax_district      string  `json:"tax_district"`
	Tax_amphoe        string  `json:"tax_amphoe"`
	Tax_province      string  `json:"tax_province"`
	Tax_zipcode       string  `json:"tax_zipcode"`
	Tax_customer_tax  string  `json:"tax_customer_tax"`
	Tax_comment       string  `json:"tax_comment"`
	Tax_total_price   float64 `json:"tax_total_price"`
	Tax_discount      float64 `json:"tax_discount"`
	Tax_befor_vat     float64 `json:"tax_befor_vat"`
	Tax_type_id       int     `json:"tax_type_id"`
	Tax_rate          int     `json:"tax_rate"`
	Tax_vat           float64 `json:"tax_vat"`
	Tax_total         float64 `json:"tax_total"`
	Tax_payment_type  int     `json:"tax_payment_type"`
	Tax_rec_type_id   int     `json:"tax_rec_type_id"`
	Tax_period        int     `json:"tax_period"`
	Tax_pay           float64 `json:"tax_pay"`
	Tax_balance       float64 `json:"tax_balance"`
	Tax_pay_total     float64 `json:"tax_pay_total"`
	Tax_discription   string  `json:"tax_discription"`
	Tax_account       int     `json:"tax_account"`
	Tax_is_active     int     `json:"tax_is_active"`
	Tax_user_id       int     `json:"tax_user_id"`
	Tax_user_fullname string  `json:"tax_user_fullname"`
	Tax_datetime      string  `json:"tax_datetime"`
	Tax_create        string  `json:"tax_create"`
	Tax_update        string  `json:"tax_update"`
	CustomerConfigSendmail
}

// X-Rey Detail
type XReyDetailEmail struct {
	Queue_id       int    `json:"queue_id"`
	Que_code       string `json:"que_code"`
	Que_note       string `json:"que_note"`
	Que_directions string `json:"que_directions"`
	CustomerConfigSendmail
}
