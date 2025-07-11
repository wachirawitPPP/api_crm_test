package structs

type ObjPayloadSearchReceipt struct {
	Search           *string `json:"search" binding:"required,omitempty"`
	Rec_is_active    *string `json:"rec_is_active" binding:"required,omitempty"`
	Rec_pay_datetime *string `json:"rec_pay_datetime"`
	Customer_id      *string `json:"customer_id" binding:"required,omitempty"`
	Shop_id          int     `json:"shop_id"`
	ActivePage       int     `json:"active_page" binding:"required"`
	PerPage          int     `json:"per_page" binding:"required"`
}

type ResponsePaginationReceipt struct {
	Result_data   []ReceiptList `json:"result_data"`
	Count_of_page int           `json:"count_of_page"`
	Count_all     int           `json:"count_all"`
}

type ObjPayloadReceipt struct {
	Id                int     `json:"id" binding:"required"`
	Rec_payment_type  int     `json:"rec_payment_type"`
	Rec_type_id       int     `json:"rec_type_id"`
	Rec_pay           float64 `json:"rec_pay"`
	Rec_discription   string  `json:"rec_discription"`
	Rec_account       int     `json:"rec_account"`
	Rec_user_id       int     `json:"rec_user_id"`
	Rec_user_fullname string  `json:"rec_user_fullname"`
}

type ReceiptList struct {
	Id                      int      `json:"id"`
	Shop_id                 int      `json:"shop_id"`
	Shop_name               string   `json:"shop_name"`
	User_id                 int      `json:"user_id"`
	User_fullname           string   `json:"user_fullname"`
	User_fullname_en        string   `json:"user_fullname_en"`
	Customer_id             int      `json:"customer_id"`
	Ctm_id                  string   `json:"ctm_id"`
	CtmShopId               int      `json:"ctm_shop_id"`
	CtmShopName             string   `json:"ctm_shop_name"`
	Queue_id                *int     `json:"queue_id"`
	Invoice_id              int      `json:"invoice_id"`
	Inv_code                string   `json:"inv_code"`
	Rec_code                string   `json:"rec_code"`
	Rec_fullname            string   `json:"rec_fullname"`
	Rec_tel                 string   `json:"rec_tel"`
	Rec_email               string   `json:"rec_email"`
	Rec_address             string   `json:"rec_address"`
	Rec_district            string   `json:"rec_district"`
	Rec_amphoe              string   `json:"rec_amphoe"`
	Rec_province            string   `json:"rec_province"`
	Rec_zipcode             string   `json:"rec_zipcode"`
	Rec_comment             string   `json:"rec_comment"`
	Rec_total_price         float64  `json:"rec_total_price"`
	Rec_discount            float64  `json:"rec_discount"`
	Rec_befor_vat           float64  `json:"rec_befor_vat"`
	Tax_type_id             int      `json:"tax_type_id"`
	Tax_rate                int      `json:"tax_rate"`
	Rec_vat                 float64  `json:"rec_vat"`
	Rec_total               float64  `json:"rec_total"`
	Rec_payment_type        int      `json:"rec_payment_type"`
	Rec_type_id             int      `json:"rec_type_id"`
	Rec_period              int      `json:"rec_period"`
	Rec_pay                 float64  `json:"rec_pay"`
	Rec_balance             float64  `json:"rec_balance"`
	Rec_pay_total           float64  `json:"rec_pay_total"`
	Rec_discription         string   `json:"rec_discription"`
	Rec_account             int      `json:"rec_account"`
	Rec_is_process          int      `json:"rec_is_process"`
	Rec_is_active           int      `json:"rec_is_active"`
	User_id_cancel          int      `json:"user_id_cancel"`
	User_fullname_cancel    string   `json:"user_fullname_cancel"`
	User_fullname_en_cancel string   `json:"user_fullname_en_cancel"`
	Rec_user_id             int      `json:"rec_user_id"`
	Rec_user_fullname       string   `json:"rec_user_fullname"`
	Rec_file                string   `json:"rec_file"`
	Rec_datetime            string   `json:"rec_datetime"`
	Rec_create              string   `json:"rec_create"`
	Rec_update              string   `json:"rec_update"`
	Rec_is_cancel           int      `json:"rec_is_cancel"`
	RecPayDatetime          string   `json:"rec_pay_datetime"`
	RecDiscountTypeId       int      `json:"rec_discount_type_id"`
	Rec_discount_item       *float64 `json:"rec_discount_item"`
	DpmId                   *int     `json:"dpm_id"`
	Rec_eclaim_id           *int     `json:"rec_eclaim_id"`
	Rec_eclaim_rate         float64  `json:"rec_eclaim_rate"`
	Rec_eclaim_over         float64  `json:"rec_eclaim_over"`
	Rec_eclaim_total        float64  `json:"rec_eclaim_total"`
}

type Receipt struct {
	Id                 int      `json:"id"`
	Shop_id            int      `json:"shop_id"`
	User_id            int      `json:"user_id"`
	Customer_id        int      `json:"customer_id"`
	Queue_id           *int     `json:"queue_id"`
	Invoice_id         int      `json:"invoice_id"`
	Rec_code           string   `json:"rec_code"`
	Rec_fullname       string   `json:"rec_fullname"`
	Rec_tel            string   `json:"rec_tel"`
	Rec_email          string   `json:"rec_email"`
	Rec_address        string   `json:"rec_address"`
	Rec_district       string   `json:"rec_district"`
	Rec_amphoe         string   `json:"rec_amphoe"`
	Rec_province       string   `json:"rec_province"`
	Rec_zipcode        string   `json:"rec_zipcode"`
	Rec_comment        string   `json:"rec_comment"`
	Rec_total_price    float64  `json:"rec_total_price"`
	Rec_discount       float64  `json:"rec_discount"`
	Rec_befor_vat      float64  `json:"rec_befor_vat"`
	Tax_type_id        int      `json:"tax_type_id"`
	Tax_rate           int      `json:"tax_rate"`
	Rec_vat            float64  `json:"rec_vat"`
	Rec_total          float64  `json:"rec_total"`
	Rec_payment_type   int      `json:"rec_payment_type"`
	Rec_type_id        int      `json:"rec_type_id"`
	Rec_period         int      `json:"rec_period"`
	Rec_pay            float64  `json:"rec_pay"`
	Rec_balance        float64  `json:"rec_balance"`
	Rec_pay_total      float64  `json:"rec_pay_total"`
	Rec_discription    string   `json:"rec_discription"`
	Rec_account        int      `json:"rec_account"`
	Rec_is_process     int      `json:"rec_is_process"`
	Rec_point_give     int      `json:"rec_point_give"`
	Rec_point_used     int      `json:"rec_point_used"`
	Rec_is_active      int      `json:"rec_is_active"`
	Rec_user_id        int      `json:"rec_user_id"`
	Rec_user_fullname  string   `json:"rec_user_fullname"`
	Rec_file           string   `json:"rec_file"`
	Rec_datetime       string   `json:"rec_datetime"`
	Rec_create         string   `json:"rec_create"`
	Rec_update         string   `json:"rec_update"`
	RecDiscountTypeId  int      `json:"rec_discount_type_id"`
	Rec_discount_item  *float64 `json:"rec_discount_item"`
	Rec_discount_value *float64 `json:"rec_discount_value"`
	DpmId              *int     `json:"dpm_id"`
	Rec_eclaim_id      *int     `json:"rec_eclaim_id"`
	Rec_eclaim_rate    float64  `json:"rec_eclaim_rate"`
	Rec_eclaim_over    float64  `json:"rec_eclaim_over"`
	Rec_eclaim_total   float64  `json:"rec_eclaim_total"`
}

type ReceiptDetail struct {
	Id                    int      `json:"id"`
	Receipt_id            int      `json:"receipt_id"`
	Course_id             *int     `json:"course_id"`
	Ser_exp_date          string   `json:"ser_exp_date"`
	Course_type_id        *int     `json:"course_type_id"`
	Checking_id           *int     `json:"checking_id"`
	Product_id            *int     `json:"product_id"`
	Product_store_id      *int     `json:"product_store_id"`
	Product_unit_id       *int     `json:"product_unit_id"`
	Coin_id               *int     `json:"coin_id"`
	Room_id               *int     `json:"room_id"`
	Queue_id              *int     `json:"queue_id"`
	Invoice_detail_id     int      `json:"invoice_detail_id"`
	Recd_type_id          int      `json:"recd_type_id"`
	Recd_code             string   `json:"recd_code"`
	Recd_name             string   `json:"recd_name"`
	Recd_qty              float64  `json:"recd_qty"`
	Recd_set_qty          float64  `json:"recd_set_qty"`
	Recd_limit_qty        float64  `json:"recd_limit_qty"`
	Recd_unit             string   `json:"recd_unit"`
	U_name_en             string   `json:"u_name_en"`
	Recd_cost             float64  `json:"recd_cost"`
	Recd_price            float64  `json:"recd_price"`
	Recd_discount         float64  `json:"recd_discount"`
	Recd_amount           float64  `json:"recd_amount"`
	Tax_type_id           int      `json:"tax_type_id"`
	Tax_rate              int      `json:"tax_rate"`
	Recd_vat              float64  `json:"recd_vat"`
	Recd_total            float64  `json:"recd_total"`
	Recd_is_set           int      `json:"recd_is_set"`
	Recd_is_active        int      `json:"recd_is_active"`
	Recd_modify           string   `json:"recd_modify"`
	Recd_eclaim           *float64 `json:"recd_eclaim"`
	Recd_discount_type_id int      `json:"recd_discount_type_id"`
	Recd_discount_item    *float64 `json:"recd_discount_item"`
}

type ReceiptInvoice struct {
	Id              int     `json:"id"`
	Shop_id         int     `json:"shop_id"`
	User_id         int     `json:"user_id"`
	Customer_id     int     `json:"customer_id"`
	Queue_id        *int    `json:"queue_id"`
	Order_id        int     `json:"order_id"`
	Inv_code        string  `json:"inv_code"`
	Inv_fullname    string  `json:"inv_fullname"`
	Inv_tel         string  `json:"inv_tel"`
	Inv_email       string  `json:"inv_email"`
	Inv_address     string  `json:"inv_address"`
	Inv_district    string  `json:"inv_district"`
	Inv_amphoe      string  `json:"inv_amphoe"`
	Inv_province    string  `json:"inv_province"`
	Inv_zipcode     string  `json:"inv_zipcode"`
	Inv_comment     string  `json:"inv_comment"`
	Inv_total_price float64 `json:"inv_total_price"`
	Inv_discount    float64 `json:"inv_discount"`
	Inv_befor_vat   float64 `json:"inv_befor_vat"`
	Tax_type_id     int     `json:"tax_type_id"`
	Tax_rate        int     `json:"tax_rate"`
	Inv_vat         float64 `json:"inv_vat"`
	Inv_total       float64 `json:"inv_total"`
	Inv_pay         float64 `json:"inv_pay"`
	Inv_is_active   int     `json:"inv_is_active"`
	Inv_datetime    string  `json:"inv_datetime"`
	Inv_create      string  `json:"inv_create"`
	Inv_update      string  `json:"inv_update"`
}

type ReceiptInvoiceDetail struct {
	Id               int     `json:"id"`
	Invoice_id       int     `json:"invoice_id"`
	Course_id        *int    `json:"course_id" `
	Checking_id      *int    `json:"checking_id"`
	Product_id       *int    `json:"product_id"`
	Product_store_id *int    `json:"product_store_id"`
	Product_unit_id  *int    `json:"product_unit_id"`
	Coin_id          *int    `json:"coin_id"`
	Room_id          *int    `json:"room_id"`
	Queue_id         *int    `json:"queue_id"`
	Order_detail_id  int     `json:"order_detail_id"`
	Invd_type_id     int     `json:"invd_type_id"`
	Invd_code        string  `json:"invd_code"`
	Invd_name        string  `json:"invd_name"`
	Invd_qty         float64 `json:"invd_qty"`
	Invd_set_qty     float64 `json:"invd_set_qty"`
	Invd_limit_qty   float64 `json:"invd_limit_qty"`
	Invd_unit        string  `json:"invd_unit"`
	Invd_cost        float64 `json:"invd_cost"`
	Invd_price       float64 `json:"invd_price"`
	Invd_discount    float64 `json:"invd_discount"`
	Invd_amount      float64 `json:"invd_amount"`
	Tax_type_id      int     `json:"tax_type_id"`
	Tax_rate         int     `json:"tax_rate"`
	Invd_vat         float64 `json:"invd_vat"`
	Invd_total       float64 `json:"invd_total"`
	Invd_is_set      int     `json:"invd_is_set"`
	Invd_is_active   int     `json:"invd_is_active"`
	Invd_modify      string  `json:"invd_modify"`
}

type LogReceipt struct {
	Username   string `json:"username"`
	Log_type   string `json:"log_type"`
	Log_text   string `json:"log_text"`
	Log_create string `json:"log_create"`
	Shop_id    int    `json:"shop_id"`
}

type InvoiceReceiptUpdate struct {
	Id            int     `json:"id"`
	Inv_pay       float64 `json:"inv_pay"`
	Inv_is_active int     `json:"inv_is_active"`
	Inv_update    string  `json:"inv_update"`
}

type ReceiptPrint struct {
	Id                      int              `json:"id"`
	Shop_id                 int              `json:"shop_id"`
	Shop_lang               string           `json:"shop_lang"`
	User_id                 int              `json:"user_id"`
	Customer_id             int              `json:"customer_id"`
	Queue_id                int              `json:"queue_id"`
	Que_code                string           `json:"que_code"`
	Invoice_id              int              `json:"invoice_id"`
	Inv_code                string           `json:"inv_code"`
	Rec_code                string           `json:"rec_code"`
	Rec_pay_datetime        string           `json:"rec_pay_datetime"`
	Acl_code                string           `json:"acl_code"`
	Acl_name                string           `json:"acl_name"`
	Rec_fullname            string           `json:"rec_fullname"`
	Rec_tel                 string           `json:"rec_tel"`
	Rec_email               string           `json:"rec_email"`
	Rec_address             string           `json:"rec_address"`
	Rec_district            string           `json:"rec_district"`
	Rec_amphoe              string           `json:"rec_amphoe"`
	Rec_province            string           `json:"rec_province"`
	Rec_zipcode             string           `json:"rec_zipcode"`
	Rec_comment             string           `json:"rec_comment"`
	Rec_total_price         float64          `json:"rec_total_price"`
	Rec_discount            float64          `json:"rec_discount"`
	Rec_befor_vat           float64          `json:"rec_befor_vat"`
	Tax_type_id             int              `json:"tax_type_id"`
	Tax_rate                int              `json:"tax_rate"`
	Rec_vat                 float64          `json:"rec_vat"`
	Rec_total               float64          `json:"rec_total"`
	Rec_payment_type        int              `json:"rec_payment_type"`
	Rec_payment_type_th     string           `json:"rec_payment_type_th" gorm:"-"`
	Rec_payment_type_en     string           `json:"rec_payment_type_en" gorm:"-"`
	Rec_type_id             int              `json:"rec_type_id"`
	Rec_period              int              `json:"rec_period"`
	Rec_pay                 float64          `json:"rec_pay"`
	Rec_balance             float64          `json:"rec_balance"`
	Rec_pay_total           float64          `json:"rec_pay_total"`
	Rec_description         string           `json:"rec_description"`
	Rec_account             int              `json:"rec_account"`
	Rec_is_process          int              `json:"rec_is_process"`
	Rec_is_active           int              `json:"rec_is_active"`
	Rec_user_id             int              `json:"rec_user_id"`
	Rec_user_fullname       string           `json:"rec_user_fullname"`
	Rec_user_fullname_en    string           `json:"rec_user_fullname_en"`
	User_id_cancel          int              `json:"user_id_cancel"`
	User_fullname_cancel    string           `json:"user_fullname_cancel"`
	User_fullname_en_cancel string           `json:"user_fullname_en_cancel"`
	Rec_update              string           `json:"rec_update"`
	Rec_file                string           `json:"rec_file"`
	Rec_datetime            string           `json:"rec_datetime"`
	Shop                    ReceiptShop      `json:"shop" gorm:"-"`
	Customer                ObjQueryCustomer `json:"customer" gorm:"-"`
	Subs                    *[]ReceiptDetail `json:"subs" gorm:"-"`
	RecDiscountTypeId       int              `json:"rec_discount_type_id"`
	Rec_discount_item       *float64         `json:"rec_discount_item"`
	Rec_discount_value      *float64         `json:"rec_discount_value"`
	DpmId                   *int             `json:"dpm_id"`
	Rec_eclaim_id           *int             `json:"rec_eclaim_id"`
	Rec_eclaim_rate         float64          `json:"rec_eclaim_rate"`
	Rec_eclaim_over         float64          `json:"rec_eclaim_over"`
	Rec_eclaim_total        float64          `json:"rec_eclaim_total"`
}

type DocReceipt struct {
	ShopId                 int    `json:"shop_id"`
	Receipt_id_default     string `json:"receipt_id_default"`
	Receipt_number_default string `json:"receipt_number_default"`
	Receipt_number_digit   int    `json:"receipt_number_digit"`
	Receipt_type           int    `json:"receipt_type"`
}

type ReceiptQueueUpdate struct {
	Id            int    `json:"id"`
	Que_status_id int    `json:"que_status_id"`
	Que_update    string `json:"que_update"`
}

type ReceiptShop struct {
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

// tom code
type ObjPayloadAddReceipt struct {
	InvoiceId       int     `json:"invoice_id" binding:"required"`
	ShopId          int     `json:"shop_id" binding:"required"`
	RecTypeId       int     `json:"rec_type_id" binding:"required"`
	AccountListId   *int    `json:"account_list_id"`
	RecPaymentType  int     `json:"rec_payment_type" binding:"required"`
	RecPay          float64 `json:"rec_pay"`
	RecDescription  string  `json:"rec_description"`
	RecAccount      int     `json:"rec_account"`
	RecUserId       int     `json:"rec_user_id" `
	RecUserFullname string  `json:"rec_user_fullname"`
	RecUserEmail    string  `json:"rec_user_email"`
	RecPayDatetime  string  `json:"rec_pay_datetime"`
	DpmId           *int    `json:"dpm_id"`
}

type ObjQueryReceiptDocSetting struct {
	ID                   int    `json:"id"`
	ReceiptIdDefault     string `json:"receipt_id_default"`
	ReceiptNumberDefault int    `json:"receipt_number_default"`
	ReceiptNumberDigit   int    `json:"receipt_number_digit"`
	ReceiptType          int    `json:"receipt_type"`
}

type CancelInvoiceUpdate struct {
	Inv_is_active int     `json:"inv_is_active"`
	Inv_pay_total float64 `json:"inv_pay_total"`
	Inv_deposit   float64 `json:"inv_deposit"`
	Inv_time_end  int     `json:"inv_time_end"`
	Inv_update    string  `json:"inv_update"`
}

type PointCustomerUpdate struct {
	Ctm_point  int    `json:"ctm_point"`
	Ctm_update string `json:"ctm_update"`
}

type CoinCustomerUpdate struct {
	Ctm_coin   float64 `json:"ctm_coin"`
	Ctm_update string  `json:"ctm_update"`
}

type MaxReceipt struct {
	Rec_period_max int `json:"rec_period_max"`
}

type FileReceipt struct {
	Rec_id   int     `json:"rec_id" binding:"required"`
	Rec_file *string `json:"rec_file" binding:"required,omitempty"`
}

type ReceiptFileCheck struct {
	Id                int     `json:"id"`
	Shop_id           int     `json:"shop_id"`
	Shop_code         string  `json:"shop_code"`
	User_id           int     `json:"user_id"`
	Customer_id       int     `json:"customer_id"`
	Ctm_id            string  `json:"ctm_id"`
	Queue_id          *int    `json:"queue_id"`
	Invoice_id        int     `json:"invoice_id"`
	Rec_code          string  `json:"rec_code"`
	Rec_fullname      string  `json:"rec_fullname"`
	Rec_tel           string  `json:"rec_tel"`
	Rec_email         string  `json:"rec_email"`
	Rec_address       string  `json:"rec_address"`
	Rec_district      string  `json:"rec_district"`
	Rec_amphoe        string  `json:"rec_amphoe"`
	Rec_province      string  `json:"rec_province"`
	Rec_zipcode       string  `json:"rec_zipcode"`
	Rec_comment       string  `json:"rec_comment"`
	Rec_total_price   float64 `json:"rec_total_price"`
	Rec_discount      float64 `json:"rec_discount"`
	Rec_befor_vat     float64 `json:"rec_befor_vat"`
	Tax_type_id       int     `json:"tax_type_id"`
	Tax_rate          int     `json:"tax_rate"`
	Rec_vat           float64 `json:"rec_vat"`
	Rec_total         float64 `json:"rec_total"`
	Rec_payment_type  int     `json:"rec_payment_type"`
	Rec_type_id       int     `json:"rec_type_id"`
	Rec_period        int     `json:"rec_period"`
	Rec_pay           float64 `json:"rec_pay"`
	Rec_balance       float64 `json:"rec_balance"`
	Rec_pay_total     float64 `json:"rec_pay_total"`
	Rec_discription   string  `json:"rec_discription"`
	Rec_account       int     `json:"rec_account"`
	Rec_is_process    int     `json:"rec_is_process"`
	Rec_point_give    int     `json:"rec_point_give"`
	Rec_point_used    int     `json:"rec_point_used"`
	Rec_is_active     int     `json:"rec_is_active"`
	Rec_user_id       int     `json:"rec_user_id"`
	Rec_user_fullname string  `json:"rec_user_fullname"`
	Rec_file          string  `json:"rec_file"`
	Rec_datetime      string  `json:"rec_datetime"`
	Rec_create        string  `json:"rec_create"`
	Rec_update        string  `json:"rec_update"`
}

// coin historys
type CoinHistory struct {
	Id          int     `json:"id"`
	Shop_id     int     `json:"shop_id"`
	Customer_id int     `json:"customer_id"`
	Receipt_id  int     `json:"receipt_id"`
	Rec_code    string  `json:"rec_code" `
	Ch_forward  float64 `json:"ch_forward"`
	Ch_amount   float64 `json:"ch_amount"`
	Ch_total    float64 `json:"ch_total"`
	Ch_comment  string  `json:"ch_comment"`
	Ch_create   string  `json:"ch_create"`
}

type AddCoinHistory struct {
	Shop_id     int     `json:"shop_id"`
	Customer_id int     `json:"customer_id"`
	Receipt_id  int     `json:"receipt_id"`
	Rec_code    string  `json:"rec_code" `
	Ch_forward  float64 `json:"ch_forward"`
	Ch_amount   float64 `json:"ch_amount"`
	Ch_total    float64 `json:"ch_total"`
	Ch_comment  string  `json:"ch_comment"`
	Ch_create   string  `json:"ch_create"`
}

// point historys
type PointHistory struct {
	Id          int     `json:"id"`
	Shop_id     int     `json:"shop_id"`
	Customer_id int     `json:"customer_id"`
	Receipt_id  int     `json:"receipt_id"`
	Rec_code    string  `json:"rec_code" `
	Ph_forward  float64 `json:"ph_forward"`
	Ph_amount   float64 `json:"ph_amount"`
	Ph_total    float64 `json:"ph_total"`
	Ph_comment  string  `json:"ph_comment"`
	Ph_create   string  `json:"ph_create"`
}

type AddPointHistory struct {
	Shop_id     int     `json:"shop_id"`
	Customer_id int     `json:"customer_id"`
	Receipt_id  int     `json:"receipt_id"`
	Rec_code    string  `json:"rec_code" `
	Ph_forward  float64 `json:"ph_forward"`
	Ph_amount   float64 `json:"ph_amount"`
	Ph_total    float64 `json:"ph_total"`
	Ph_comment  string  `json:"ph_comment"`
	Ph_create   string  `json:"ph_create"`
}

type ObjPayloadAddReceiptOrder struct {
	InvoiceId int `json:"invoice_id" binding:"required"`
	// ShopId          int     `json:"shop_id" binding:"required"`
	// RecTypeId       int     `json:"rec_type_id" binding:"required"`
	// AccountListId   *int    `json:"account_list_id"`
	// RecPaymentType  int     `json:"rec_payment_type" binding:"required"`
	// RecPay          float64 `json:"rec_pay"`
	// RecDescription  string  `json:"rec_description"`
	// RecAccount      int     `json:"rec_account"`
	// RecUserId       int     `json:"rec_user_id" `
	// RecUserFullname string  `json:"rec_user_fullname"`
	// RecUserEmail    string  `json:"rec_user_email"`
	// RecPayDatetime  string  `json:"rec_pay_datetime"`
	// DpmId           *int    `json:"dpm_id"`
}
