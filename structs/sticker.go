package structs

type ObjPayloadSearchSticker struct {
	Search_text *string `json:"search_text"`
	Search_date *string `json:"search_date"`
	User_id     *int    `json:"user_id"`
	Shop_id     int     `json:"shop_id"`
	ActivePage  int     `json:"active_page" binding:"required"`
	PerPage     int     `json:"per_page" binding:"required"`
	StkActiveId *int    `json:"sticker_active_id"`
}

type ObjPayloadList struct {
	Id *[]int `json:"id" binding:"required,omitempty"`
}

type ObjPayloadLists struct {
	Invoice_id *[]int `json:"invoice_id" binding:"required,omitempty"`
}

type ResponsePaginationSticker struct {
	Result_data   []StickersList `json:"result_data"`
	Count_of_page int            `json:"count_of_page"`
	Count_all     int            `json:"count_all"`
}

type Stickers struct {
	Id                  int     `json:"id"`
	Product_id          int     `json:"product_id"`
	Customer_id         int     `json:"customer_id"`
	Invoice_id          int     `json:"invoice_id"`
	Inv_code            string  `json:"inv_code"`
	Receipt_id          int     `json:"receipt_id"`
	Rec_code            string  `json:"rec_code"`
	Ctm_id              string  `json:"ctm_id"`
	Ctm_prefix          string  `json:"ctm_prefix"`
	Ctm_fname           string  `json:"ctm_fname"`
	Ctm_lname           string  `json:"ctm_lname"`
	Ctm_fname_en        string  `json:"ctm_fname_en"`
	Ctm_lname_en        string  `json:"ctm_lname_en"`
	User_id             int     `json:"user_id"`
	User_fullname       string  `json:"user_fullname"`
	User_fullname_en    string  `json:"user_fullname_en"`
	Sticker_products    string  `json:"sticker_products"`
	Sticker_direction   string  `json:"sticker_direction"`
	Sticker_active_id   int     `json:"sticker_active_id"`
	Sticker_print_label int     `json:"sticker_print_label"`
	Sticker_print_order int     `json:"sticker_print_order"`
	Sticker_is_del      int     `json:"sticker_is_del"`
	Sticker_modify      string  `json:"sticker_modify"`
	ShopId              int     `json:"shop_id"`
	OrderId             int     `json:"order_id"`
	StickerCode         string  `json:"sticker_code"`
	StickerName         string  `json:"sticker_name"`
	StickerAmount       float64 `json:"sticker_amount"`
	StickerUnit         string  `json:"sticker_unit"`
	StickerExpdate      string  `json:"sticker_expdate"`
	StickerTopical      string  `json:"sticker_topical"`
}

type PrintStickerDrugLabel struct {
	Id                  int     `json:"id"`
	Product_id          int     `json:"product_id"`
	Customer_id         int     `json:"customer_id"`
	User_id             int     `json:"user_id"`
	ShopId              int     `json:"shop_id"`
	Sticker_code        string  `json:"sticker_code"`
	Sticker_name        string  `json:"sticker_name"`
	Sticker_name_acc    string  `json:"sticker_name_acc"`
	Sticker_amount      float64 `json:"sticker_amount"`
	Sticker_unit        string  `json:"sticker_unit"`
	Sticker_unit_en     string  `json:"sticker_unit_en"`
	Sticker_price       float64 `json:"sticker_price"`
	Sticker_expdate     string  `json:"sticker_expdate"`
	Sticker_active_id   int     `json:"sticker_active_id"`
	Sticker_print_label int     `json:"sticker_print_label"`
	Sticker_print_order int     `json:"sticker_print_order"`
	Sticker_topical     string  `json:"sticker_topical"`
	Sticker_direction   string  `json:"sticker_direction"`
	Sticker_is_del      int     `json:"sticker_is_del"`
	Sticker_modify      string  `json:"sticker_modify"`
}

type StickersList struct {
	Id                  int              `json:"id"`
	Product_id          int              `json:"product_id"`
	Customer_id         int              `json:"customer_id"`
	Invoice_id          int              `json:"invoice_id"`
	Inv_code            string           `json:"inv_code"`
	Receipt_id          int              `json:"receipt_id"`
	Rec_code            string           `json:"rec_code"`
	Ctm_id              string           `json:"ctm_id"`
	Ctm_prefix          string           `json:"ctm_prefix"`
	Ctm_fname           string           `json:"ctm_fname"`
	Ctm_lname           string           `json:"ctm_lname"`
	Ctm_fname_en        string           `json:"ctm_fname_en"`
	Ctm_lname_en        string           `json:"ctm_lname_en"`
	User_id             int              `json:"user_id"`
	User_fullname       string           `json:"user_fullname"`
	User_fullname_en    string           `json:"user_fullname_en"`
	Sticker_products    string           `json:"sticker_products"`
	Sticker_direction   string           `json:"sticker_direction"`
	Sticker_active_id   int              `json:"sticker_active_id"`
	Sticker_print_label int              `json:"sticker_print_label"`
	Sticker_print_order int              `json:"sticker_print_order"`
	Sticker_is_del      int              `json:"sticker_is_del"`
	Sticker_modify      string           `json:"sticker_modify"`
	StickerProduct      []StickerProduct `json:"sticker_product"`
	ShopId              int              `json:"shop_id"`
	OrderId             int              `json:"order_id"`
	StickerCode         string           `json:"sticker_code"`
	StickerName         string           `json:"sticker_name"`
	StickerAmount       float64          `json:"sticker_amount"`
	StickerUnit         string           `json:"sticker_unit"`
	StickerExpdate      string           `json:"sticker_expdate"`
	StickerTopical      string           `json:"sticker_topical"`
}

type StickerProduct struct {
	Id                int     `json:"id"`
	Invoice_id        int     `json:"invoice_id"`
	Sticker_code      string  `json:"sticker_code"`
	Sticker_name      string  `json:"sticker_name"`
	Sticker_name_acc  string  `json:"sticker_name_acc"`
	Sticker_amount    float64 `json:"sticker_amount"`
	Sticker_price     float64 `json:"sticker_price"`
	Sticker_topical   string  `json:"sticker_topical"`
	Sticker_direction string  `json:"sticker_direction"`
	Sticker_unit      string  `json:"sticker_unit"`
	Sticker_expdate   string  `json:"sticker_expdate"`
	Sticker_modify    string  `json:"sticker_modify"`
	Invd_discount     float64 `json:"invd_discount"`
	Invd_amount       float64 `json:"invd_amount"`
	Invd_total        float64 `json:"invd_total"`
	Invd_vat          float64 `json:"invd_vat"`
	Pd_code_acc       string  `json:"pd_code_acc"`
	Pd_name_acc       string  `json:"pd_name_acc"`
}

type UpdateSticker struct {
	Id                int    `json:"id"`
	Sticker_topical   string `json:"sticker_topical"`
	Sticker_direction string `json:"sticker_direction"`
}

type StickerDetail struct {
	Id               int    `json:"id"`
	Product_id       int    `json:"product_id"`
	Customer_id      int    `json:"customer_id"`
	Invoice_id       int    `json:"invoice_id"`
	Inv_code         string `json:"inv_code"`
	Receipt_id       int    `json:"receipt_id"`
	Rec_code         string `json:"rec_code"`
	Ctm_id           string `json:"ctm_id"`
	Ctm_prefix       string `json:"ctm_prefix"`
	Ctm_fname        string `json:"ctm_fname"`
	Ctm_lname        string `json:"ctm_lname"`
	Ctm_fname_en     string `json:"ctm_fname_en"`
	Ctm_lname_en     string `json:"ctm_lname_en"`
	Ctm_birthdate    string `json:"ctm_birthdate"`
	User_id          int    `json:"user_id"`
	User_fullname    string `json:"user_fullname"`
	Shop_name        string `json:"shop_name"`
	Shop_name_en     string `json:"shop_name_en"`
	Shop_license     string `json:"shop_license"`
	Shop_tax         string `json:"shop_tax"`
	Shop_phone       string `json:"shop_phone"`
	Shop_address     string `json:"shop_address"`
	Shop_address_en  string `json:"shop_address_en"`
	Shop_district    string `json:"shop_district"`
	Shop_district_en string `json:"shop_district_en"`
	Shop_amphoe      string `json:"shop_amphoe"`
	Shop_amphoe_en   string `json:"shop_amphoe_en"`
	Shop_province    string `json:"shop_province"`
	Shop_province_en string `json:"shop_province_en"`
	Shop_zipcode     string `json:"shop_zipcode"`
	Shop_zipcode_en  string `json:"shop_zipcode_en"`
	Shop_image       string `json:"shop_image"`
	StickerDocSetting
	StickerProduct
}

type StickerDetail2 struct {
	Id               int    `json:"id"`
	Product_id       int    `json:"product_id"`
	Customer_id      int    `json:"customer_id"`
	Invoice_id       int    `json:"invoice_id"`
	Inv_code         string `json:"inv_code"`
	Receipt_id       int    `json:"receipt_id"`
	Rec_code         string `json:"rec_code"`
	Ctm_id           string `json:"ctm_id"`
	Ctm_prefix       string `json:"ctm_prefix"`
	Ctm_fname        string `json:"ctm_fname"`
	Ctm_lname        string `json:"ctm_lname"`
	Ctm_fname_en     string `json:"ctm_fname_en"`
	Ctm_lname_en     string `json:"ctm_lname_en"`
	Ctm_birthdate    string `json:"ctm_birthdate"`
	User_id          int    `json:"user_id"`
	User_fullname    string `json:"user_fullname"`
	Shop_name        string `json:"shop_name"`
	Shop_image       string `json:"shop_image"`
	Shop_name_en     string `json:"shop_name_en"`
	Shop_license     string `json:"shop_license"`
	Shop_tax         string `json:"shop_tax"`
	Shop_phone       string `json:"shop_phone"`
	Shop_address     string `json:"shop_address"`
	Shop_address_en  string `json:"shop_address_en"`
	Shop_district    string `json:"shop_district"`
	Shop_district_en string `json:"shop_district_en"`
	Shop_amphoe      string `json:"shop_amphoe"`
	Shop_amphoe_en   string `json:"shop_amphoe_en"`
	Shop_province    string `json:"shop_province"`
	Shop_province_en string `json:"shop_province_en"`
	Shop_zipcode     string `json:"shop_zipcode"`
	Shop_zipcode_en  string `json:"shop_zipcode_en"`
	StickerDocSetting
	StickerProduct []StickerProduct `json:"sticker_product" gorm:"-"`
}

// type Prescription struct {
// 	Id                  int              `json:"id"`
// 	Shop_id             int              `json:"shop_id"`
// 	Shop_lang           string           `json:"shop_lang"`
// 	User_id             int              `json:"user_id"`
// 	Customer_id         int              `json:"customer_id"`
// 	Queue_id            int              `json:"queue_id"`
// 	Que_code            string           `json:"que_code"`
// 	Que_user_fullname   string           `json:"que_user_fullname"`
// 	User_fullname       string           `json:"user_fullname"`
// 	Invoice_id          int              `json:"invoice_id"`
// 	Inv_code            string           `json:"inv_code"`
// 	Rec_code            string           `json:"rec_code"`
// 	Acl_code            string           `json:"acl_code"`
// 	Acl_name            string           `json:"acl_name"`
// 	Rec_fullname        string           `json:"rec_fullname"`
// 	Rec_tel             string           `json:"rec_tel"`
// 	Rec_email           string           `json:"rec_email"`
// 	Rec_address         string           `json:"rec_address"`
// 	Rec_district        string           `json:"rec_district"`
// 	Rec_amphoe          string           `json:"rec_amphoe"`
// 	Rec_province        string           `json:"rec_province"`
// 	Rec_zipcode         string           `json:"rec_zipcode"`
// 	Rec_comment         string           `json:"rec_comment"`
// 	Rec_total_price     float64          `json:"rec_total_price"`
// 	Rec_discount        float64          `json:"rec_discount"`
// 	Rec_befor_vat       float64          `json:"rec_befor_vat"`
// 	Tax_type_id         int              `json:"tax_type_id"`
// 	Tax_rate            int              `json:"tax_rate"`
// 	Rec_vat             float64          `json:"rec_vat"`
// 	Rec_total           float64          `json:"rec_total"`
// 	Rec_payment_type    int              `json:"rec_payment_type"`
// 	Rec_payment_type_th string           `json:"rec_payment_type_th" gorm:"-"`
// 	Rec_payment_type_en string           `json:"rec_payment_type_en" gorm:"-"`
// 	Rec_type_id         int              `json:"rec_type_id"`
// 	Rec_period          int              `json:"rec_period"`
// 	Rec_pay             float64          `json:"rec_pay"`
// 	Rec_balance         float64          `json:"rec_balance"`
// 	Rec_pay_total       float64          `json:"rec_pay_total"`
// 	Rec_description     string           `json:"rec_description"`
// 	Rec_account         int              `json:"rec_account"`
// 	Rec_is_process      int              `json:"rec_is_process"`
// 	Rec_is_active       int              `json:"rec_is_active"`
// 	Rec_user_id         int              `json:"rec_user_id"`
// 	Rec_user_fullname   string           `json:"rec_user_fullname"`
// 	Rec_datetime        string           `json:"rec_datetime"`
// 	Shop                ReceiptShop      `json:"shop" gorm:"-"`
// 	Customer            ObjQueryCustomer `json:"customer" gorm:"-"`
// 	Subs                []StickerDetail  `json:"subs" gorm:"-"`
// }

type StickerDocSetting struct {
	Sticker_font_size    *int `json:"sticker_font_size" binding:"required,omitempty"`
	Sticker_width        *int `json:"sticker_width" binding:"required,omitempty"`
	Sticker_height       *int `json:"sticker_height" binding:"required,omitempty"`
	Sticker_show_name    *int `json:"sticker_show_name" binding:"required,omitempty"`
	Sticker_show_tel     *int `json:"sticker_show_tel" binding:"required,omitempty"`
	Sticker_show_detail  *int `json:"sticker_show_detail" binding:"required,omitempty"`
	Sticker_show_date    *int `json:"sticker_show_date" binding:"required,omitempty"`
	Sticker_show_expdate *int `json:"sticker_show_expdate" binding:"required,omitempty"`
	Show_product_acc     *int `json:"show_product_acc" binding:"required,omitempty"`
	Show_shop_image      *int `json:"show_shop_image" binding:"required,omitempty"`
}

type LogStk struct {
	Username   string `json:"username"`
	Log_type   string `json:"log_type"`
	Log_text   string `json:"log_text"`
	Log_create string `json:"log_create"`
	Shop_id    int    `json:"shop_id"`
}

type Prescription struct {
	Id                int              `json:"id"`
	Shop_id           int              `json:"shop_id"`
	Shop_lang         string           `json:"shop_lang"`
	User_id           int              `json:"user_id"`
	Customer_id       int              `json:"customer_id"`
	Queue_id          int              `json:"queue_id"`
	Que_code          string           `json:"que_code"`
	Que_user_fullname string           `json:"que_user_fullname"`
	User_fullname     string           `json:"user_fullname"`
	Invoice_id        int              `json:"invoice_id"`
	Inv_code          string           `json:"inv_code"`
	Inv_fullname      string           `json:"inv_fullname"`
	Inv_tel           string           `json:"inv_tel"`
	Inv_email         string           `json:"inv_email"`
	Inv_address       string           `json:"inv_address"`
	Inv_district      string           `json:"inv_district"`
	Inv_amphoe        string           `json:"inv_amphoe"`
	Inv_province      string           `json:"inv_province"`
	Inv_zipcode       string           `json:"inv_zipcode"`
	Inv_comment       string           `json:"inv_comment"`
	Inv_total_price   float64          `json:"inv_total_price"`
	Inv_discount      float64          `json:"inv_discount"`
	Inv_befor_vat     float64          `json:"inv_befor_vat"`
	Tax_type_id       int              `json:"tax_type_id"`
	Tax_rate          int              `json:"tax_rate"`
	Inv_vat           float64          `json:"inv_vat"`
	Inv_total         float64          `json:"inv_total"`
	Inv_type_id       int              `json:"inv_type_id"`
	Inv_is_active     int              `json:"inv_is_active"`
	Inv_user_id       int              `json:"inv_user_id"`
	Inv_user_fullname string           `json:"inv_user_fullname"`
	Inv_datetime      string           `json:"inv_datetime"`
	Shop              StickerShop      `json:"shop" gorm:"-"`
	Customer          ObjQueryCustomer `json:"customer" gorm:"-"`
	Subs              []StickerDetail  `json:"subs" gorm:"-"`
}

type StickerShop struct {
	// primary key
	Id int `json:"shop_id"`
	// foreign key
	ShopNatureId   int    `json:"shop_nature_id"`
	ShopTypeId     int    `json:"shop_type_id"`
	ShopPackageId  int    `json:"shop_package_id"`
	CurrencyId     int    `json:"currency_id"`
	CurrencySymbol string `json:"currency_symbol"`
	// fields
	ShopCode                 string `json:"shop_code" `
	ShopName                 string `json:"shop_name"`
	ShopLicense              string `json:"shop_license"`
	ShopNature               string `json:"shop_nature"`
	ShopTax                  string `json:"shop_tax"`
	ShopPhone                string `json:"shop_phone"`
	ShopFax                  string `json:"shop_fax"`
	ShopEmail                string `json:"shop_email"`
	ShopAddress              string `json:"shop_address"`
	ShopDistrict             string `json:"shop_district"`
	ShopAmphoe               string `json:"shop_amphoe"`
	ShopProvince             string `json:"shop_province"`
	ShopZipcode              string `json:"shop_zipcode"`
	Shop_name_en             string `json:"shop_name_en"`
	Shop_nature_en           string `json:"shop_nature_en"`
	Shop_address_en          string `json:"shop_address_en"`
	Shop_district_en         string `json:"shop_district_en"`
	Shop_amphoe_en           string `json:"shop_amphoe_en"`
	Shop_province_en         string `json:"shop_province_en"`
	Shop_zipcode_en          string `json:"shop_zipcode_en"`
	Shop_country_en          string `json:"shop_country_en"`
	Shop_company_name        string `json:"shop_company_name"`
	Shop_company_name_en     string `json:"shop_company_name_en"`
	Shop_company_address     string `json:"shop_company_address"`
	Shop_company_address_en  string `json:"shop_company_address_en"`
	Shop_company_district_en string `json:"shop_company_district_en"`
	Shop_company_district    string `json:"shop_company_district"`
	Shop_company_amphoe      string `json:"shop_company_amphoe"`
	Shop_company_amphoe_en   string `json:"shop_company_amphoe_en"`
	Shop_company_province    string `json:"shop_company_province"`
	Shop_company_province_en string `json:"shop_company_province_en"`
	Shop_company_zipcode     string `json:"shop_company_zipcode"`
	Shop_company_zipcode_en  string `json:"shop_company_zipcode_en"`
	Shop_company_country     string `json:"shop_company_country"`
	Shop_company_country_en  string `json:"shop_company_country_en"`
	Shop_company_tax         string `json:"shop_company_tax"`
	Shop_company_phone       string `json:"shop_company_phone"`
	Shop_company_email       string `json:"shop_company_email"`
	ShopLatlong              string `json:"shop_latlong"`
	ShopPromptpayIdd         string `json:"shop_promptpay_id"`
	ShopPromptpayName        string `json:"shop_promptpay_name"`
	ShopMotherId             int    `json:"shop_mother_id"`
	ShopImage                string `json:"shop_image"`
	ShopDetail               string `json:"shop_detail"`
	ShopLang                 string `json:"shop_lang"`
	ShopStatusId             int    `json:"shop_status_id"`
	ShopSmsSum               string `json:"shop_sms_sum"`
	ShopSmsAll               string `json:"shop_sms_all"`
	ShopPointExchange        string `json:"shop_point_exchange"`
	ShopDrugLowest           string `json:"shop_drug_lowest"`
	ShopDrugExpire           string `json:"shop_drug_expire"`
	ShopPackageType          string `json:"shop_package_type"`
	ShopPackageUsemanager    string `json:"shop_package_usemanager"`
	ShopPackageUseuser       string `json:"shop_package_useuser"`
	ShopPackageUsedoctor     string `json:"shop_package_usedoctor"`
	ShopPackageUseshop       string `json:"shop_package_useshop"`
	ShopPackage              string `json:"shop_package"`
	ShopExpire               string `json:"shop_expire"`
	ShopFacebook             string `json:"shop_facebook"`
	ShopLine                 string `json:"shop_line"`
	ShopInstagram            string `json:"shop_instagram"`
	ShopPortIdcard           string `json:"shop_port_idcard"`
	ShopPrintType            string `json:"shop_print_type"`
	ShopComment              string `json:"shop_comment"`
	ShopConnectLineStatus    string `json:"shop_connect_line_status"`
	ShopPaysolutionsRefno    string `json:"shop_paysolutions_refno"`
	ShopCmline               string `json:"shop_cmline"`
	ShopS3DeleteId           int    `json:"shop_s3_delete_id"`
	ShopShippingCost         string `json:"shop_shipping_cost"`
	PaysolutionsMerchant     string `json:"paysolutions_merchant"`
	PaysolutionsSecretkey    string `json:"paysolutions_secretkey"`
	PaysolutionsApikey       string `json:"paysolutions_apikey"`
	PaysolutionsActive       string `json:"paysolutions_active"`
	Show_product_th          string `json:"show_product_th"`
	Show_product_en          string `json:"show_product_en"`
	Show_course_check_th     string `json:"show_course_check_th"`
	Show_course_check_en     string `json:"show_course_check_en"`
	Show_date_id             int    `json:"show_date_id"`
	Show_page_id             int    `json:"show_page_id"`
	Print_th                 int    `json:"print_th"`
	Print_en                 int    `json:"print_en"`
	Print_la                 int    `json:"print_la"`
	Print_a4                 int    `json:"print_a4"`
	Print_ca                 int    `json:"print_ca"`
	Print_a5                 int    `json:"print_a5"`
	Print_80                 int    `json:"print_80"`
	Invoice_comment_id       int    `json:"invoice_comment_id"`
	Invoice_comment          string `json:"invoice_comment"`
	Receipt_comment_id       int    `json:"receipt_comment_id"`
	Receipt_comment          string `json:"receipt_comment"`
	Tax_comment_id           int    `json:"tax_comment_id"`
	Tax_comment              string `json:"tax_comment"`
	Purchase_comment_id      int    `json:"purchase_comment_id"`
	Purchase_comment         string `json:"purchase_comment"`
	Transfer_comment_id      int    `json:"transfer_comment_id"`
	Transfer_comment         string `json:"transfer_comment"`
	Invoice_copy             int    `json:"invoice_copy"`
	Receipt_copy             int    `json:"receipt_copy"`
	Tax_copy                 int    `json:"tax_copy"`
	Purchase_copy            int    `json:"purchase_copy"`
	Transfer_copy            int    `json:"transfer_copy"`
	Sticker_font_size        int    `json:"sticker_font_size"`
	Sticker_width            int    `json:"sticker_width"`
	Sticker_height           int    `json:"sticker_height"`
	Sticker_show_name        int    `json:"sticker_show_name"`
	Sticker_show_address     int    `json:"sticker_show_address"`
	Sticker_show_tel         int    `json:"sticker_show_tel"`
	Sticker_show_date        int    `json:"sticker_show_date"`
	Sticker_show_expdate     int    `json:"sticker_show_expdate"`
	Sticker_show_detail      int    `json:"sticker_show_detail"`
}
