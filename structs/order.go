package structs

type ObjPayloadSearchCart struct {
	Search      *string `json:"search" binding:"required,omitempty"`
	Customer_id int     `json:"customer_id" binding:"required"`
	Shop_id     int     `json:"shop_id"`
	ActivePage  int     `json:"active_page" binding:"required"`
	PerPage     int     `json:"per_page" binding:"required"`
}

type ResponsePaginationCart struct {
	Result_data   []Carts `json:"result_data"`
	Count_of_page int     `json:"count_of_page"`
	Count_all     int     `json:"count_all"`
}

type ObjPayloadSearchOrder struct {
	Search       *string `json:"search" binding:"required,omitempty"`
	Or_is_active *string `json:"or_is_active" binding:"required,omitempty"`
	Or_datetime  *string `json:"or_datetime"`
	Customer_id  *string `json:"customer_id" binding:"required,omitempty"`
	Shop_id      int     `json:"shop_id"`
	ActivePage   int     `json:"active_page" binding:"required"`
	PerPage      int     `json:"per_page" binding:"required"`
}

type ResponsePaginationOrder struct {
	Result_data   []Order `json:"result_data"`
	Count_of_page int     `json:"count_of_page"`
	Count_all     int     `json:"count_all"`
}

type ResponsePaginationOrderLists struct {
	Result_data   []OrderLists `json:"result_data"`
	Count_of_page int          `json:"count_of_page"`
	Count_all     int          `json:"count_all"`
}

type ObjPayloadSearchOrderCustomer struct {
	// ShopIds    []int   `json:"shop_ids" binding:"required"`
	ShopId     int     `json:"shop_id"`
	SearchText *string `json:"search_text" binding:"required"`
}

type ObjResponseSearchOrderCustomer struct {
	// ID       int    `json:"id"`
	// CtmId    string `json:"ctm_id"`
	// CtmFname string `json:"ctm_fname"`
	// CtmLname string `json:"ctm_lname"`
	ID                  int     `json:"id"`
	ShopId              int     `json:"shop_id"`
	CustomerGroupId     int     `json:"customer_group_id"`
	UserId              int     `json:"user_id"`
	CtmId               string  `json:"ctm_id"`
	CtmCitizenId        string  `json:"ctm_citizen_id"`
	CtmPassportId       string  `json:"ctm_passport_id"`
	CtmPrefix           string  `json:"ctm_prefix"`
	CtmFname            string  `json:"ctm_fname"`
	CtmLname            string  `json:"ctm_lname"`
	CtmNname            string  `json:"ctm_nname"`
	CtmFnameEn          string  `json:"ctm_fname_en"`
	CtmLnameEn          string  `json:"ctm_lname_en"`
	CtmGender           string  `json:"ctm_gender"`
	CtmNation           string  `json:"ctm_nation"`
	CtmReligion         string  `json:"ctm_religion"`
	CtmEduLevel         string  `json:"ctm_edu_level"`
	CtmMaritalStatus    string  `json:"ctm_marital_status"`
	CtmBlood            string  `json:"ctm_blood"`
	CtmEmail            string  `json:"ctm_email"`
	CtmTel              string  `json:"ctm_tel"`
	CtmTel_2            string  `json:"ctm_tel_2"`
	CtmBirthdate        string  `json:"ctm_birthdate"`
	CtmAddress          string  `json:"ctm_address"`
	CtmDistrict         string  `json:"ctm_district"`
	CtmAmphoe           string  `json:"ctm_amphoe"`
	CtmProvince         string  `json:"ctm_province"`
	CtmZipcode          string  `json:"ctm_zipcode"`
	CtmComment          string  `json:"ctm_comment"`
	CtmWeight           float64 `json:"ctm_weight" gorm:"type:decimal(10,2)"`
	CtmHeight           float64 `json:"ctm_height" gorm:"type:decimal(10,2)"`
	CtmWaistline        float64 `json:"ctm_waistline" gorm:"type:decimal(10,2)"`
	CtmChest            float64 `json:"ctm_chest" gorm:"type:decimal(10,2)"`
	CtmTreatmentType    int     `json:"ctm_treatment_type"`
	RightTreatmentId    int     `json:"right_treatment_id"`
	CtmAllergic         string  `json:"ctm_allergic"`
	CtmMentalHealth     string  `json:"ctm_mental_health"`
	CtmDisease          string  `json:"ctm_disease"`
	CtmHealthComment    string  `json:"ctm_health_comment"`
	CtmImage            string  `json:"ctm_image"`
	CtmImageSize        int     `json:"ctm_image_size"`
	CtmPoint            int     `json:"ctm_point"`
	CtmCoin             float64 `json:"ctm_coin" gorm:"type:decimal(10,2)"`
	LineToken           string  `json:"line_token"`
	LineSend            int     `json:"line_send"`
	LineSendDate        string  `json:"line_send_date"`
	FacebookId          string  `json:"facebook_id"`
	CompanyName         string  `json:"company_name"`
	CompanyTax          string  `json:"company_tax"`
	CompanyTel          string  `json:"company_tel"`
	CompanyEmail        string  `json:"company_email"`
	CompanyAddress      string  `json:"company_address"`
	CompanyDistrict     string  `json:"company_district"`
	CompanyAmphoe       string  `json:"company_amphoe"`
	CompanyProvince     string  `json:"company_province"`
	CompanyZipcode      string  `json:"company_zipcode"`
	CtmSubscribeOpd     int     `json:"ctm_subscribe_opd"`
	CtmSubscribeLab     int     `json:"ctm_subscribe_lab"`
	CtmSubscribeCert    int     `json:"ctm_subscribe_cert"`
	CtmSubscribeReceipt int     `json:"ctm_subscribe_receipt"`
	CtmSubscribeAppoint int     `json:"ctm_subscribe_appoint"`
	CtmIsActive         int     `json:"ctm_is_active"`
	CtmIsDel            int     `json:"ctm_is_del"`
	CtmCreate           string  `json:"ctm_create"`
	CtmUpdate           string  `json:"ctm_update"`
	// join
	CgName     string `json:"cg_name"`
	CgSaveType int    `json:"cg_save_type"` //ประเภทส่วนลด : 1 สกุลเงิน, 2 %
	CgSave     int    `json:"cg_save"`
	Rt_code    string `json:"rt_code"`
	Rt_name    string `json:"rt_name"`
}

type Carts struct {
	Id               int     `json:"id"`
	Shop_id          int     `json:"shop_id"`
	User_id          int     `json:"user_id"`
	Customer_id      int     `json:"customer_id"`
	Queue_id         int     `json:"queue_id"`
	Course_id        int     `json:"course_id"`
	Checking_id      int     `json:"checking_id"`
	Product_id       int     `json:"product_id"`
	Product_store_id int     `json:"product_store_id"`
	Product_unit_id  int     `json:"product_unit_id"`
	Room_id          int     `json:"room_id"`
	Ordt_type_id     int     `json:"ordt_type_id"`
	Ordt_code        string  `json:"ordt_code"`
	Ordt_name        string  `json:"ordt_name"`
	Ordt_qty         float64 `json:"ordt_qty"`
	Ordt_unit        string  `json:"ordt_unit"`
	Ordt_cost        float64 `json:"ordt_cost"`
	Ordt_price       float64 `json:"ordt_price"`
	Ordt_discount    float64 `json:"ordt_discount"`
	Tax_type_id      int     `json:"tax_type_id"`
	Tax_rate         int     `json:"tax_rate"`
	Ordt_vat         float64 `json:"ordt_vat"`
	Ordt_total       float64 `json:"ordt_total"`
	Ordt_is_active   int     `json:"ordt_is_active"`
	Ordt_modify      string  `json:"ordt_modify"`
	Room_code        string  `json:"room_code"`
	Room_th          string  `json:"room_th"`
	Room_en          string  `json:"room_en"`
	Room_type_th     string  `json:"room_type_th"`
	Room_type_en     string  `json:"room_type_en"`
	Course_code      string  `json:"course_code"`
	Course_name      string  `json:"course_name"`
	Course_unit      string  `json:"course_unit"`
	Checking_code    string  `json:"checking_code"`
	Checking_name    string  `json:"checking_name"`
	Checking_unit    string  `json:"checking_unit"`
	Ctm_fname        string  `json:"ctm_fname"`
	Ctm_lname        string  `json:"ctm_lname"`
	Pd_code          string  `json:"pd_code"`
	Pd_name          string  `json:"pd_name"`
	U_name           string  `json:"u_name"`
}

type Order struct {
	Id                  int     `json:"id"`
	Shop_id             int     `json:"shop_id"`
	User_id             int     `json:"user_id"`
	User_fullname       string  `json:"user_fullname"`
	User_fullname_en    string  `json:"user_fullname_en"`
	Customer_id         int     `json:"customer_id"`
	Ctm_id              string  `json:"ctm_id"`
	Queue_id            int     `json:"queue_id"`
	Que_code            string  `json:"que_code"`
	Or_fullname         string  `json:"or_fullname"`
	Or_tel              string  `json:"or_tel"`
	Or_email            string  `json:"or_email"`
	Or_address          string  `json:"or_address"`
	Or_district         string  `json:"or_district"`
	Or_amphoe           string  `json:"or_amphoe"`
	Or_province         string  `json:"or_province"`
	Or_zipcode          string  `json:"or_zipcode"`
	Or_comment          string  `json:"or_comment"`
	Or_total_price      float64 `json:"or_total_price"`
	Or_discount_type_id int     `json:"or_discount_type_id"`
	Or_discount_item    float64 `json:"or_discount_item"`
	Or_discount_value   float64 `json:"or_discount_value"`
	Or_discount         float64 `json:"or_discount"`
	Or_befor_vat        float64 `json:"or_befor_vat"`
	Tax_type_id         int     `json:"tax_type_id"`
	Tax_rate            int     `json:"tax_rate"`
	Or_vat              float64 `json:"or_vat"`
	Or_total            float64 `json:"or_total"`
	Or_is_active        int     `json:"or_is_active"`
	Or_datetime         string  `json:"or_datetime"`
	Or_create           string  `json:"or_create"`
	Or_update           string  `json:"or_update"`
}

type OrderLists struct {
	Id                      int     `json:"id"`
	Shop_id                 int     `json:"shop_id"`
	User_id                 int     `json:"user_id"`
	User_fullname           string  `json:"user_fullname"`
	User_fullname_en        string  `json:"user_fullname_en"`
	Customer_id             int     `json:"customer_id"`
	Ctm_id                  string  `json:"ctm_id"`
	CtmShopId               int     `json:"ctm_shop_id"`
	CtmShopName             string  `json:"ctm_shop_name"`
	Queue_id                int     `json:"queue_id"`
	Que_code                string  `json:"que_code"`
	Or_fullname             string  `json:"or_fullname"`
	Or_tel                  string  `json:"or_tel"`
	Or_email                string  `json:"or_email"`
	Or_address              string  `json:"or_address"`
	Or_district             string  `json:"or_district"`
	Or_amphoe               string  `json:"or_amphoe"`
	Or_province             string  `json:"or_province"`
	Or_zipcode              string  `json:"or_zipcode"`
	Or_comment              string  `json:"or_comment"`
	Or_total_price          float64 `json:"or_total_price"`
	Or_discount_type_id     int     `json:"or_discount_type_id"`
	Or_discount_item        float64 `json:"or_discount_item"`
	Or_discount_value       float64 `json:"or_discount_value"`
	Or_discount             float64 `json:"or_discount"`
	Or_befor_vat            float64 `json:"or_befor_vat"`
	Tax_type_id             int     `json:"tax_type_id"`
	Tax_rate                int     `json:"tax_rate"`
	Or_vat                  float64 `json:"or_vat"`
	Or_total                float64 `json:"or_total"`
	Or_is_active            int     `json:"or_is_active"`
	User_id_cancel          int     `json:"user_id_cancel"`
	User_fullname_cancel    string  `json:"user_fullname_cancel"`
	User_fullname_en_cancel string  `json:"user_fullname_en_cancel"`
	Or_datetime             string  `json:"or_datetime"`
	Or_create               string  `json:"or_create"`
	Or_update               string  `json:"or_update"`
}

type OrderSub struct {
	Id                   int                `json:"id"`
	Order_id             int                `json:"order_id"`
	Course_id            *int               `json:"course_id" gorm:"default:0"`
	Checking_id          *int               `json:"checking_id" gorm:"default:0"`
	Product_id           *int               `json:"product_id" gorm:"default:0"`
	Product_store_id     *int               `json:"product_store_id" gorm:"default:0"`
	Product_unit_id      *int               `json:"product_unit_id" gorm:"default:0"`
	Coin_id              *int               `json:"coin_id" gorm:"default:0"`
	Room_id              *int               `json:"room_id" gorm:"default:0"`
	Queue_id             *int               `json:"queue_id" gorm:"default:0"`
	Queue_ord_id         *int               `json:"queue_ord_id" gorm:"default:0"`
	Queue_checking_id    *int               `json:"queue_checking_id" gorm:"default:0"`
	Queue_course_id      *int               `json:"queue_course_id" gorm:"default:0"`
	Queue_product_id     *int               `json:"queue_product_id" gorm:"default:0"`
	Ord_type_id          int                `json:"ord_type_id"`
	Ord_code             string             `json:"ord_code"`
	Ord_name             string             `json:"ord_name"`
	Ord_qty              float64            `json:"ord_qty"`
	Ord_rate             float64            `json:"ord_rate"`
	Ord_set_qty          float64            `json:"ord_set_qty"`
	Ord_limit_qty        float64            `json:"ord_limit_qty"`
	Ord_unit             string             `json:"ord_unit"`
	Ord_cost             float64            `json:"ord_cost"`
	Ord_price            float64            `json:"ord_price"`
	Ord_amount           float64            `json:"ord_amount"`
	Ord_discount_type_id int                `json:"ord_discount_type_id"`
	Ord_discount_item    float64            `json:"ord_discount_item"`
	Ord_discount         float64            `json:"ord_discount"`
	Tax_type_id          int                `json:"tax_type_id"`
	Tax_rate             int                `json:"tax_rate"`
	Ord_vat              float64            `json:"ord_vat"`
	Ord_total            float64            `json:"ord_total"`
	Topical_id           *int               `json:"topical_id" gorm:"default:0"`
	Ord_topical          string             `json:"ord_topical"`
	Ord_direction        string             `json:"ord_direction"`
	Ord_is_set           int                `json:"ord_is_set"`
	Ord_id_set           *int               `json:"ord_id_set"`
	Ord_is_use           int                `json:"ord_is_use"`
	Ord_is_active        int                `json:"ord_is_active"`
	Ord_modify           string             `json:"ord_modify"`
	Room_code            string             `json:"room_code"`
	Room_th              string             `json:"room_th"`
	Room_en              string             `json:"room_en"`
	Room_type_th         string             `json:"room_type_th"`
	Room_type_en         string             `json:"room_type_en"`
	Units                *[]ProductUnitList `json:"units" gorm:"-"`
	Label                string             `json:"label"`
	U_name               string             `json:"u_name"`
	U_name_en            string             `json:"u_name_en"`
	Balance              float64            `json:"balance"`
	Ord_eclaim           *float64           `json:"ord_eclaim"`
	Claim_price_ofc      float64            `json:"claim_price_ofc"`
	Claim_price_lgo      float64            `json:"claim_price_lgo"`
	Claim_price_ucs      float64            `json:"claim_price_ucs"`
	Claim_price_sss      float64            `json:"claim_price_sss"`
	Claim_price_nhs      float64            `json:"claim_price_nhs"`
	Claim_price_ssi      float64            `json:"claim_price_ssi"`
}

type OrderSubUpdate struct {
	Tax_type_id          int     `json:"tax_type_id"`
	Tax_rate             int     `json:"tax_rate"`
	Ord_qty              float64 `json:"ord_qty"`
	Ord_amount           float64 `json:"ord_amount"`
	Ord_vat              float64 `json:"ord_vat"`
	Ord_total            float64 `json:"ord_total"`
	Ord_discount_type_id int     `json:"ord_discount_type_id"`
	Ord_discount_item    float64 `json:"ord_discount_item"`
	Ord_discount         float64 `json:"ord_discount"`
}

type OrderTags struct {
	Id       int     `json:"id"`
	Order_id int     `json:"order_id"`
	Tags_id  int     `json:"tags_id"`
	Tag_name *string `json:"tag_name"`
}

type OrderDetail struct {
	Id                  int              `json:"id"`
	Shop_id             int              `json:"shop_id"`
	User_id             int              `json:"user_id"`
	User_fullname       string           `json:"user_fullname"`
	User_fullname_en    string           `json:"user_fullname_en"`
	Customer_id         int              `json:"customer_id"`
	Customer_online_id  int              `json:"customer_online_id"`
	Queue_id            int              `json:"queue_id"`
	Que_code            string           `json:"que_code"`
	Or_fullname         string           `json:"or_fullname"`
	Or_tel              string           `json:"or_tel"`
	Or_email            string           `json:"or_email"`
	Or_address          string           `json:"or_address"`
	Or_district         string           `json:"or_district"`
	Or_amphoe           string           `json:"or_amphoe"`
	Or_province         string           `json:"or_province"`
	Or_zipcode          string           `json:"or_zipcode"`
	Or_comment          string           `json:"or_comment"`
	Or_total_price      float64          `json:"or_total_price"`
	Or_discount_type_id int              `json:"or_discount_type_id"`
	Or_discount_item    float64          `json:"or_discount_item"`
	Or_discount_value   float64          `json:"or_discount_value"`
	Or_discount         float64          `json:"or_discount"`
	Or_befor_vat        float64          `json:"or_befor_vat"`
	Tax_type_id         int              `json:"tax_type_id"`
	Tax_rate            int              `json:"tax_rate"`
	Or_vat              float64          `json:"or_vat"`
	Or_total            float64          `json:"or_total"`
	Or_is_active        int              `json:"or_is_active"`
	Or_datetime         string           `json:"or_datetime"`
	Qr_tele_code        string           `json:"or_tele_code"`
	Or_create           string           `json:"or_create"`
	Or_update           string           `json:"or_update"`
	Shop                ReceiptShop      `json:"shop" gorm:"-"`
	Customer            ObjQueryCustomer `json:"customer" gorm:"-"`
	Tags                *[]OrderTags     `json:"tags" gorm:"-"`
	Subs                *[]OrderSub      `json:"subs" gorm:"-"`
	DpmId               *int             `json:"dpm_id"`
	Or_eclaim_id        *int             `json:"or_eclaim_id"`
	Or_eclaim_rate      *float64         `json:"or_eclaim_rate"`
	Or_eclaim_over      *float64         `json:"or_eclaim_over"`
	Or_eclaim_total     *float64         `json:"or_eclaim_total"`
	Or_tele_code        string           `json:"or_tele_code"`
}

type OrderDetailData struct {
	Id                      int              `json:"id"`
	Shop_id                 int              `json:"shop_id"`
	User_id                 int              `json:"user_id"`
	User_fullname           string           `json:"user_fullname"`
	User_fullname_en        string           `json:"user_fullname_en"`
	Customer_id             int              `json:"customer_id"`
	Queue_id                int              `json:"queue_id"`
	Que_code                string           `json:"que_code"`
	Or_fullname             string           `json:"or_fullname"`
	Or_tel                  string           `json:"or_tel"`
	Or_email                string           `json:"or_email"`
	Or_address              string           `json:"or_address"`
	Or_district             string           `json:"or_district"`
	Or_amphoe               string           `json:"or_amphoe"`
	Or_province             string           `json:"or_province"`
	Or_zipcode              string           `json:"or_zipcode"`
	Or_comment              string           `json:"or_comment"`
	Or_total_price          float64          `json:"or_total_price"`
	Or_discount_type_id     int              `json:"or_discount_type_id"`
	Or_discount_item        float64          `json:"or_discount_item"`
	Or_discount_value       float64          `json:"or_discount_value"`
	Or_discount             float64          `json:"or_discount"`
	Or_befor_vat            float64          `json:"or_befor_vat"`
	Tax_type_id             int              `json:"tax_type_id"`
	Tax_rate                int              `json:"tax_rate"`
	Or_vat                  float64          `json:"or_vat"`
	Or_total                float64          `json:"or_total"`
	Or_is_active            int              `json:"or_is_active"`
	User_id_cancel          int              `json:"user_id_cancel"`
	User_fullname_cancel    string           `json:"user_fullname_cancel"`
	User_fullname_en_cancel string           `json:"user_fullname_en_cancel"`
	Or_datetime             string           `json:"or_datetime"`
	Qr_tele_code            string           `json:"or_tele_code"`
	Or_create               string           `json:"or_create"`
	Or_update               string           `json:"or_update"`
	Shop                    ReceiptShop      `json:"shop" gorm:"-"`
	Customer                ObjQueryCustomer `json:"customer" gorm:"-"`
	Tags                    *[]OrderTags     `json:"tags" gorm:"-"`
	Subs                    *[]OrderSub      `json:"subs" gorm:"-"`
	DpmId                   int              `json:"dpm_id"`
	Or_eclaim_id            int              `json:"or_eclaim_id"`
	Or_eclaim_rate          float64          `json:"or_eclaim_rate"`
	Or_eclaim_over          float64          `json:"or_eclaim_over"`
	Or_eclaim_total         float64          `json:"or_eclaim_total"`
}

type UserCancel struct {
	User_fullname    string `json:"user_fullname"`
	User_fullname_en string `json:"user_fullname_en"`
}

type ObjPayloadAddOrder struct {
	Id                  int          `json:"id"`
	Shop_id             int          `json:"shop_id"`
	Customer_online_id  int          `json:"customer_online_id" binding:"required"`
	Co_citizen_id       string       `json:"co_citizen_id"`
	Co_prefix           string       `json:"co_prefix"`
	Co_fname            string       `json:"co_fname"`
	Co_lname            string       `json:"co_lname"`
	Co_gender           string       `json:"co_gender"`
	Co_email            string       `json:"co_email"`
	Co_tel              string       `json:"co_tel"`
	Co_birthdate        string       `json:"co_birthdate"`
	Co_address          string       `json:"co_address"`
	Co_district         string       `json:"co_district"`
	Co_amphoe           string       `json:"co_amphoe"`
	Co_province         string       `json:"co_province"`
	Co_zipcode          string       `json:"co_zipcode"`
	Queue_id            *int         `json:"queue_id" binding:"required,omitempty"`
	Or_fullname         string       `json:"or_fullname" binding:"required"`
	Or_tel              string       `json:"or_tel"`
	Or_email            string       `json:"or_email"`
	Or_address          *string      `json:"or_address" binding:"required,omitempty"`
	Or_district         string       `json:"or_district" binding:"required"`
	Or_amphoe           string       `json:"or_amphoe" binding:"required"`
	Or_province         string       `json:"or_province" binding:"required"`
	Or_zipcode          string       `json:"or_zipcode" binding:"required"`
	Or_comment          string       `json:"or_comment"`
	Or_total_price      *float64     `json:"or_total_price" binding:"required,omitempty"`
	Or_discount_type_id int          `json:"or_discount_type_id" binding:"required"`
	Or_discount_item    *float64     `json:"or_discount_item" binding:"required,omitempty"`
	Or_discount_value   *float64     `json:"or_discount_value" binding:"required,omitempty"`
	Or_discount         *float64     `json:"or_discount" binding:"required,omitempty"`
	Or_befor_vat        float64      `json:"or_befor_vat"`
	Tax_type_id         int          `json:"tax_type_id" binding:"required"`
	Tax_rate            *int         `json:"tax_rate" binding:"omitempty"`
	Or_vat              *float64     `json:"or_vat" binding:"omitempty"`
	Or_total            *float64     `json:"or_total" binding:"omitempty"`
	Or_is_active        int          `json:"or_is_active"`
	Or_datetime         string       `json:"or_datetime"`
	Or_create           string       `json:"or_create"`
	Or_update           string       `json:"or_update"`
	Tags                *[]OrderTags `json:"tags" binding:"required,omitempty"`
	Subs                []OrderSub   `json:"subs" gorm:"-" binding:"required"`
	DpmId               *int         `json:"dpm_id" binding:"omitempty"`
	Or_eclaim_id        *int         `json:"or_eclaim_id" binding:"omitempty"`
	Or_eclaim_rate      *float64     `json:"or_eclaim_rate" binding:"omitempty"`
	Or_eclaim_over      *float64     `json:"or_eclaim_over" binding:"omitempty"`
	Or_eclaim_total     *float64     `json:"or_eclaim_total" binding:"omitempty"`
}

type ObjPayloadEditOrder struct {
	Id                  int          `json:"id" binding:"required"`
	Shop_id             int          `json:"shop_id"`
	User_id             int          `json:"user_id"`
	Customer_id         int          `json:"customer_id" binding:"required"`
	Queue_id            *int         `json:"queue_id" binding:"required,omitempty"`
	Or_fullname         string       `json:"or_fullname" binding:"required"`
	Or_tel              string       `json:"or_tel"`
	Or_email            string       `json:"or_email"`
	Or_address          *string      `json:"or_address" binding:"required,omitempty"`
	Or_district         string       `json:"or_district" binding:"required"`
	Or_amphoe           string       `json:"or_amphoe" binding:"required"`
	Or_province         string       `json:"or_province" binding:"required"`
	Or_zipcode          string       `json:"or_zipcode" binding:"required"`
	Or_comment          string       `json:"or_comment"`
	Or_total_price      *float64     `json:"or_total_price" binding:"required,omitempty"`
	Or_discount_type_id int          `json:"or_discount_type_id" binding:"required"`
	Or_discount_item    *float64     `json:"or_discount_item" binding:"required,omitempty"`
	Or_discount_value   *float64     `json:"or_discount_value" binding:"required,omitempty"`
	Or_discount         *float64     `json:"or_discount" binding:"required,omitempty"`
	Or_befor_vat        float64      `json:"or_befor_vat"`
	Tax_type_id         int          `json:"tax_type_id" binding:"required"`
	Tax_rate            *int         `json:"tax_rate" binding:"required,omitempty"`
	Or_vat              *float64     `json:"or_vat" binding:"required,omitempty"`
	Or_total            *float64     `json:"or_total" binding:"required,omitempty"`
	Or_is_active        int          `json:"or_is_active"`
	Or_datetime         string       `json:"or_datetime"`
	Or_create           string       `json:"or_create"`
	Or_update           string       `json:"or_update"`
	Tags                *[]OrderTags `json:"tags" binding:"required,omitempty"`
	Subs                []OrderSub   `json:"subs" gorm:"-" binding:"required"`
	DpmId               *int         `json:"dpm_id" binding:"omitempty"`
	Or_eclaim_id        *int         `json:"or_eclaim_id" binding:"omitempty"`
	Or_eclaim_rate      *float64     `json:"or_eclaim_rate" binding:"omitempty"`
	Or_eclaim_over      *float64     `json:"or_eclaim_over" binding:"omitempty"`
	Or_eclaim_total     *float64     `json:"or_eclaim_total" binding:"omitempty"`
}

type OrderAction struct {
	Id                  int      `json:"id"`
	Shop_id             int      `json:"shop_id"`
	User_id             int      `json:"user_id"`
	Customer_id         int      `json:"customer_id"`
	Customer_online_id  int      `json:"customer_online_id"`
	Queue_id            int      `json:"queue_id" gorm:"-"`
	Or_fullname         string   `json:"or_fullname"`
	Or_tel              string   `json:"or_tel"`
	Or_email            string   `json:"or_email"`
	Or_address          string   `json:"or_address"`
	Or_district         string   `json:"or_district"`
	Or_amphoe           string   `json:"or_amphoe"`
	Or_province         string   `json:"or_province"`
	Or_zipcode          string   `json:"or_zipcode"`
	Or_comment          string   `json:"or_comment"`
	Or_total_price      float64  `json:"or_total_price"`
	Or_discount_type_id int      `json:"or_discount_type_id"`
	Or_discount_item    float64  `json:"or_discount_item"`
	Or_discount_value   float64  `json:"or_discount_value"`
	Or_discount         float64  `json:"or_discount"`
	Or_befor_vat        float64  `json:"or_befor_vat"`
	Tax_type_id         int      `json:"tax_type_id"`
	Tax_rate            int      `json:"tax_rate"`
	Or_vat              float64  `json:"or_vat"`
	Or_total            float64  `json:"or_total"`
	Or_is_active        int      `json:"or_is_active"`
	Or_datetime         string   `json:"or_datetime"`
	Or_create           string   `json:"or_create"`
	Or_update           string   `json:"or_update"`
	DpmId               *int     `json:"dpm_id"`
	Or_eclaim_id        *int     `json:"or_eclaim_id"`
	Or_eclaim_rate      *float64 `json:"or_eclaim_rate"`
	Or_eclaim_over      *float64 `json:"or_eclaim_over"`
	Or_eclaim_total     *float64 `json:"or_eclaim_total"`
	Or_tele_code        string   `json:"or_tele_code"`
}

type LogOrders struct {
	Username   string `json:"username"`
	Log_type   string `json:"log_type"`
	Log_text   string `json:"log_text"`
	Log_create string `json:"log_create"`
	Shop_id    int    `json:"shop_id"`
}

type ItemCourseOrder struct {
	Id               *int    `json:"id"`
	Course_type_id   int     `json:"course_type_id"`
	Course_code      string  `json:"course_code"`
	Course_name      string  `json:"course_name"`
	Course_lock_drug int     `json:"course_lock_drug"`
	Course_cost      float64 `json:"course_cost"`
	Course_ofc       float64 `json:"course_ofc"`
	Course_lgo       float64 `json:"course_lgo"`
	Course_ucs       float64 `json:"course_ucs"`
	Course_sss       float64 `json:"course_sss"`
	Course_ssi       float64 `json:"course_ssi"`
	Course_nhs       float64 `json:"course_nhs"`
}

type ItemCheckingOrder struct {
	Id               *int    `json:"id"`
	Checking_type_id int     `json:"checking_type_id"`
	Checking_code    string  `json:"checking_code"`
	Checking_name    string  `json:"checking_name"`
	Checking_cost    float64 `json:"checking_cost"`
	Checking_ofc     float64 `json:"checking_ofc"`
	Checking_lgo     float64 `json:"checking_lgo"`
	Checking_ucs     float64 `json:"checking_ucs"`
	Checking_sss     float64 `json:"checking_sss"`
	Checking_nhs     float64 `json:"checking_nhs"`
	Checking_ssi     float64 `json:"checking_ssi"`
}

type OrderTopical struct {
	Id             int    `json:"id"`
	Topical_name   string `json:"topical_name"`
	Topical_detail string `json:"topical_detail"`
}

type OrderTopicalNotiProduct struct {
	Id                  int    `json:"id"`
	Noti_product_detail string `json:"noti_product_detail"`
	Noti_product_name   string `json:"noti_product_name"`
}

type CheckQueueStatusID struct {
	Id                int    `json:"id"`
	UserId            int    `json:"user_id"`
	Que_status_id     int    `json:"que_status_id"`
	Que_user_fullname string `json:"que_user_fullname"`
	Que_tele_code     string `json:"que_tele_code"`
	Shop_id           int    `json:"shop_id"`
}

type ObjPayloadAddOrderQueueIPD struct {
	Id             int                    `json:"id"`
	Shop_id        int                    `json:"shop_id"`
	User_id        int                    `json:"user_id"`
	Customer_id    int                    `json:"customer_id" binding:"required"`
	Queue_id       *int                   `json:"queue_id" binding:"required,omitempty"`
	Or_fullname    string                 `json:"or_fullname" binding:"required"`
	Or_tel         string                 `json:"or_tel"`
	Or_email       string                 `json:"or_email"`
	Or_address     *string                `json:"or_address" binding:"required,omitempty"`
	Or_district    string                 `json:"or_district" binding:"required"`
	Or_amphoe      string                 `json:"or_amphoe" binding:"required"`
	Or_province    string                 `json:"or_province" binding:"required"`
	Or_zipcode     string                 `json:"or_zipcode" binding:"required"`
	Or_comment     string                 `json:"or_comment"`
	Or_total_price *float64               `json:"or_total_price" binding:"required,omitempty"`
	Or_discount    *float64               `json:"or_discount" binding:"required,omitempty"`
	Or_befor_vat   float64                `json:"or_befor_vat"`
	Tax_type_id    int                    `json:"tax_type_id" binding:"required"`
	Tax_rate       *int                   `json:"tax_rate" binding:"required,omitempty"`
	Or_vat         *float64               `json:"or_vat" binding:"required,omitempty"`
	Or_total       *float64               `json:"or_total" binding:"required,omitempty"`
	Or_is_active   int                    `json:"or_is_active"`
	Or_datetime    string                 `json:"or_datetime"`
	Or_create      string                 `json:"or_create"`
	Or_update      string                 `json:"or_update"`
	Tags           *[]OrderTags           `json:"tags" binding:"required,omitempty"`
	Subs           []OrderSub             `json:"subs" gorm:"-" binding:"required"`
	Queue_ipd      ObjPayloadMoveQueueIPD `json:"queue_ipd" gorm:"-" binding:"required"`
}

type ObjPayloadMoveQueueIPD struct {
	Queue_id       int    `json:"queue_id"`
	CustomerId     int    `json:"customer_id"`
	DoctorId       int    `json:"doctor_id"`
	DoctorFullname string `json:"doctor_fullname"`
	RoomId         *int   `json:"room_id"`
	BedId          *int   `json:"bed_id"`
	QuePriorityId  int    `json:"que_priority_id"`
	QueDatetime    string `json:"que_datetime"`
}

type GetQueueId struct {
	ID            int    `json:"id"`
	Shop_id       int    `json:"shop_id"`
	Customer_id   int    `json:"customer_id"`
	User_Id       int    `json:"user_id"`
	Que_admis_id  int    `json:"que_admis_id"` // 1 IPD 2 OPD
	Que_code      string `json:"que_code"`
	Que_ref_ipd   int    `json:"que_ref_ipd"`
	Que_status_id int    `json:"que_status_id"`
}

type AddMoveQueueIPD struct {
	ID              int    `json:"id"`
	ShopId          int    `json:"shop_id"`
	CustomerId      int    `json:"customer_id"`
	RoomId          *int   `json:"room_id"`
	BedId           *int   `json:"bed_id"`
	UserId          int    `json:"user_id"`
	QueUserId       int    `json:"que_user_id"`
	QueUserFullname string `json:"que_user_fullname"`
	CtmFullname     string `json:"ctm_fullname"`
	CtmFullnameEn   string `json:"ctm_fullname_en"`
	QueCode         string `json:"que_code"`
	QueTypeId       int    `json:"que_type_id"`
	QueAdmisId      int    `json:"que_admis_id"`
	QueRefIpd       int    `json:"que_ref_ipd"`
	QueRefId        int    `json:"que_ref_id"`
	QuePriorityId   int    `json:"que_priority_id"`
	QueStatusId     int    `json:"que_status_id"`
	QueComment      string `json:"que_comment"`
	QueNote         string `json:"que_note"`
	QueDatetimeOut  string `json:"que_datetime_out"`
	QueRoomTotal    string `json:"que_room_total"`
	QueDatetime     string `json:"que_datetime"`
	QueTimeEnd      int    `json:"que_time_end"`
	QueCreate       string `json:"que_create"`
	QueUpdate       string `json:"que_update"`
}

type QueueProductIPD struct {
	Id                *int    `json:"id"`
	Queue_id          int     `json:"queue_id"`
	User_id           int     `json:"user_id"`
	Product_id        int     `json:"product_id"`
	Product_store_id  int     `json:"product_store_id"`
	Product_unit_id   int     `json:"product_unit_id"`
	Quep_type_id      int     `json:"quep_type_id"`
	Queue_checking_id *int    `json:"queue_checking_id"`
	Checking_id       *int    `json:"checking_id"`
	Queue_course_id   *int    `json:"queue_course_id"`
	Course_id         *int    `json:"course_id"`
	Quep_code         string  `json:"quep_code"`
	Quep_name         string  `json:"quep_name"`
	Quep_qty          float64 `json:"quep_qty"`
	Quep_set_qty      float64 `json:"quep_set_qty"`
	Quep_limit_qty    float64 `json:"quep_limit_qty"`
	Quep_unit         string  `json:"quep_unit"`
	Quep_cost         float64 `json:"quep_cost"`
	Quep_price        float64 `json:"quep_price"`
	Quep_discount     float64 `json:"quep_discount"`
	Topical_id        *int    `json:"topical_id"`
	Quep_topical      string  `json:"quep_topical"`
	Quep_direction    string  `json:"quep_direction"`
	Quep_total        float64 `json:"quep_total"`
	Quep_is_set       int     `json:"quep_is_set"`
	Quep_id_ref       int     `json:"quep_id_ref"`
	Quep_ipd_order    int     `json:"quep_ipd_order"`
	Quep_is_active    int     `json:"quep_is_active"`
	Quep_modify       string  `json:"quep_modify"`
}

type QueueCourseIPD struct {
	Id             *int    `json:"id"`
	Queue_id       int     `json:"queue_id"`
	Course_id      int     `json:"course_id"`
	User_id        int     `json:"user_id"`
	Quec_code      string  `json:"quec_code"`
	Quec_name      string  `json:"quec_name"`
	Quec_qty       float64 `json:"quec_qty"`
	Quec_unit      string  `json:"quec_unit"`
	Quec_cost      float64 `json:"quec_cost"`
	Quec_price     float64 `json:"quec_price"`
	Quec_discount  float64 `json:"quec_discount"`
	Quec_total     float64 `json:"quec_total"`
	Quec_is_set    int     `json:"quec_is_set"`
	Quec_id_ref    int     `json:"quec_id_ref"`
	Quec_ipd_order int     `json:"quec_ipd_order"`
	Quec_is_active int     `json:"quec_is_active"`
	Quec_modify    string  `json:"quec_modify"`
}

type QueueCheckingIPD struct {
	Id              *int    `json:"id"`
	Queue_id        int     `json:"queue_id"`
	Checking_id     int     `json:"checking_id"`
	User_id         int     `json:"user_id"`
	Queci_code      string  `json:"queci_code"`
	Queci_name      string  `json:"queci_name"`
	Queci_qty       float64 `json:"queci_qty"`
	Queci_unit      string  `json:"queci_unit"`
	Queci_cost      float64 `json:"queci_cost"`
	Queci_price     float64 `json:"queci_price"`
	Queci_discount  float64 `json:"queci_discount"`
	Queci_total     float64 `json:"queci_total"`
	Queci_is_set    int     `json:"queci_is_set"`
	Queci_id_ref    int     `json:"queci_id_ref"`
	Queci_ipd_order int     `json:"queci_ipd_order"`
	Queci_is_active int     `json:"queci_is_active"`
	Queci_modify    string  `json:"queci_modify"`
}

// labplus
type CheckingLabplus struct {
	Id            int    `json:"id"`
	Checking_code string `json:"checking_code"`
	Checking_name string `json:"checking_name"`
}

type CustomerPatient struct {
	Ctm_id         string `json:"ctm_id"`
	Ctm_prefix     string `json:"ctm_prefix"`
	Ctm_fname      string `json:"ctm_fname"`
	Ctm_lname      string `json:"ctm_lname"`
	Ctm_citizen_id string `json:"ctm_citizen_id"`
	Ctm_birthdate  string `json:"ctm_birthdate"`
	Ctm_gender     string `json:"ctm_gender"`
}

type UserType1 struct {
	Id            int    `json:"id"`
	User_email    string `json:"user_email"`
	User_fullname string `json:"user_fullname"`
}

type CustomerGroupId struct {
	Id int `json:"id"`
}

type AddCustomerGroup struct {
	Id           int    `json:"id"`
	Shop_id      int    `json:"shop_id"`
	Cg_name      string `json:"cg_name"`
	Cg_is_active int    `json:"cg_is_active"`
	Cg_is_online int    `json:"cg_is_online"`
	Cg_create    string `json:"cg_create"`
	Cg_update    string `json:"cg_update"`
}

type AccountCode struct {
	Id int `json:"id"`
}

type RoomBedId struct {
	Room_id int `json:"room_id"`
	Bed_id  int `json:"bed_id"`
}

type ObjPayloadCreateQueueByOrder struct {
	ShopId          int    `json:"shop_id"`
	CustomerId      int    `json:"customer_id"`
	DoctorId        int    `json:"doctor_id"`
	DoctorFullname  string `json:"doctor_fullname"`
	RoomId          int    `json:"room_id"`
	BedId           int    `json:"bed_id"`
	QueUserId       int    `json:"que_user_id"`
	QueUserFullname string `json:"que_user_fullname"`
	QueTypeId       int    `json:"que_type_id"`
	QueAdmisId      int    `json:"que_admis_id"`
	QuePriorityId   int    `json:"que_priority_id"`
	QueDatetime     string `json:"que_datetime"`
	QueTeleCode     string `json:"que_tele_code"`
	QueTeleUrl      string `json:"que_tele_url"`
}

type QueueByOrder struct {
	ID              int    `json:"id"`
	ShopId          int    `json:"shop_id"`
	CustomerId      int    `json:"customer_id"`
	RoomId          int    `json:"room_id"`
	BedId           int    `json:"bed_id"`
	UserId          int    `json:"user_id"`
	QueUserId       int    `json:"que_user_id"`
	QueUserFullname string `json:"que_user_fullname"`
	CtmFullname     string `json:"ctm_fullname"`
	CtmFullnameEn   string `json:"ctm_fullname_en"`
	QueCode         string `json:"que_code"`
	QueTypeId       int    `json:"que_type_id"`
	QueAdmisId      int    `json:"que_admis_id"`
	QuePriorityId   int    `json:"que_priority_id"`
	QueStatusId     int    `json:"que_status_id"`
	QueComment      string `json:"que_comment"`
	QueNote         string `json:"que_note"`
	QueDatetimeOut  string `json:"que_datetime_out"`
	QueRoomTotal    string `json:"que_room_total"`
	QueDatetime     string `json:"que_datetime"`
	QueTimeEnd      int    `json:"que_time_end"`
	QueTeleCode     string `json:"que_tele_code"`
	QueTeleUrl      string `json:"que_tele_url"`
	QueCreate       string `json:"que_create"`
	QueUpdate       string `json:"que_update"`
	DpmId           *int   `json:"dpm_id"`
	QueNumber       int    `json:"que_number"`
	QueDpmCode      string `json:"que_dpm_code"`
}

type QueueCheckingByOrder struct {
	Id                     *int    `json:"id"`
	Queue_id               int     `json:"queue_id"`
	Checking_id            int     `json:"checking_id"`
	User_id                int     `json:"user_id"`
	Queci_code             string  `json:"queci_code"`
	Queci_name             string  `json:"queci_name"`
	Queci_qty              float64 `json:"queci_qty"`
	Queci_unit             string  `json:"queci_unit"`
	Queci_cost             float64 `json:"queci_cost"`
	Queci_price            float64 `json:"queci_price"`
	Queci_discount_type_id int     `json:"queci_discount_type_id"`
	Queci_discount_item    float64 `json:"queci_discount_item"`
	Queci_discount         float64 `json:"queci_discount"`
	Queci_total            float64 `json:"queci_total"`
	Queci_is_set           int     `json:"queci_is_set"`
	Queci_id_ref           int     `json:"queci_id_ref"`
	Queci_id_set           *int    `json:"queci_id_set"`
	Queci_is_active        int     `json:"queci_is_active"`
	Queci_modify           string  `json:"queci_modify"`
	DpmId                  *int    `json:"dpm_id"`
}

type OrderDetailCheck struct {
	Id               int     `json:"id"`
	Product_id       *int    `json:"product_id"`
	Product_store_id *int    `json:"product_store_id"`
	Ord_code         string  `json:"ord_code"`
	Ord_name         string  `json:"ord_name"`
	Ord_qty          float64 `json:"ord_qty"`
}

type OrderDetailCheckCourse struct {
	Id      int     `json:"id"`
	Ord_qty float64 `json:"ord_qty"`
}
