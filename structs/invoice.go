package structs

type ObjPayloadSearchInvoice struct {
	Search        *string `json:"search" binding:"required,omitempty"`
	Inv_is_active *string `json:"inv_is_active" binding:"required,omitempty"`
	Inv_datetime  *string `json:"inv_datetime"`
	Customer_id   *string `json:"customer_id" binding:"required,omitempty"`
	Shop_id       int     `json:"shop_id"`
	ActivePage    int     `json:"active_page" binding:"required"`
	PerPage       int     `json:"per_page" binding:"required"`
}

type ObjPayloadInvoice struct {
	Id int `json:"id" binding:"required"`
}

type ObjPayloadInvoiceOrder struct {
	Order_id int `json:"order_id" binding:"required"`
}

type ResponsePaginationInvoice struct {
	Result_data   []InvoiceList `json:"result_data"`
	Count_of_page int           `json:"count_of_page"`
	Count_all     int           `json:"count_all"`
}

type InvoiceList struct {
	Id                      int      `json:"id"`
	Shop_id                 int      `json:"shop_id"`
	User_id                 int      `json:"user_id"`
	User_fullname           string   `json:"user_fullname"`
	User_fullname_en        string   `json:"user_fullname_en"`
	Customer_id             int      `json:"customer_id"`
	Ctm_id                  string   `json:"ctm_id"`
	CtmShopId               int      `json:"ctm_shop_id"`
	CtmShopName             string   `json:"ctm_shop_name"`
	Queue_id                *int     `json:"queue_id"`
	Que_code                *string  `json:"que_code"`
	Order_id                int      `json:"order_id"`
	Inv_code                string   `json:"inv_code"`
	Inv_fullname            string   `json:"inv_fullname"`
	Inv_tel                 string   `json:"inv_tel"`
	Inv_email               string   `json:"inv_email"`
	Inv_address             string   `json:"inv_address"`
	Inv_district            string   `json:"inv_district"`
	Inv_amphoe              string   `json:"inv_amphoe"`
	Inv_province            string   `json:"inv_province"`
	Inv_zipcode             string   `json:"inv_zipcode"`
	Inv_comment             string   `json:"inv_comment"`
	Inv_total_price         float64  `json:"inv_total_price"`
	Inv_discount            float64  `json:"inv_discount"`
	Inv_befor_vat           float64  `json:"inv_befor_vat"`
	Tax_type_id             int      `json:"tax_type_id"`
	Tax_rate                int      `json:"tax_rate"`
	Inv_vat                 float64  `json:"inv_vat"`
	Inv_total               float64  `json:"inv_total"`
	Inv_deposit             float64  `json:"inv_deposit"`
	Inv_pay_total           float64  `json:"inv_pay_total"`
	Inv_is_active           int      `json:"inv_is_active"`
	User_id_cancel          int      `json:"user_id_cancel"`
	User_fullname_cancel    string   `json:"user_fullname_cancel"`
	User_fullname_en_cancel string   `json:"user_fullname_en_cancel"`
	Inv_datetime            string   `json:"inv_datetime"`
	Inv_create              string   `json:"inv_create"`
	Inv_update              string   `json:"inv_update"`
	Receipt_id              *int     `json:"receipt_id"`
	DpmId                   int      `json:"dpm_id"`
	Inv_discount_type_id    int      `json:"inv_discount_type_id"`
	Inv_discount_item       float64  `json:"inv_discount_item"`
	Inv_discount_value      float64  `json:"inv_discount_value"`
	Inv_eclaim_id           *int     `json:"inv_eclaim_id"`
	Inv_eclaim_rate         *float64 `json:"inv_eclaim_rate"`
	Inv_eclaim_over         *float64 `json:"inv_eclaim_over"`
	Inv_eclaim_total        *float64 `json:"inv_eclaim_total"`
}

type Invoice struct {
	Id                   int      `json:"id"`
	Shop_id              int      `json:"shop_id"`
	User_id              int      `json:"user_id"`
	Customer_id          int      `json:"customer_id"`
	Customer_online_id   int      `json:"customer_online_id"`
	Queue_id             *int     `json:"queue_id"`
	Order_id             int      `json:"order_id"`
	Inv_code             string   `json:"inv_code"`
	Inv_fullname         string   `json:"inv_fullname"`
	Inv_tel              string   `json:"inv_tel"`
	Inv_email            string   `json:"inv_email"`
	Inv_address          string   `json:"inv_address"`
	Inv_district         string   `json:"inv_district"`
	Inv_amphoe           string   `json:"inv_amphoe"`
	Inv_province         string   `json:"inv_province"`
	Inv_zipcode          string   `json:"inv_zipcode"`
	Inv_comment          string   `json:"inv_comment"`
	Inv_total_price      float64  `json:"inv_total_price"`
	Inv_discount         float64  `json:"inv_discount"`
	Inv_befor_vat        float64  `json:"inv_befor_vat"`
	Tax_type_id          int      `json:"tax_type_id"`
	Tax_rate             int      `json:"tax_rate"`
	Inv_vat              float64  `json:"inv_vat"`
	Inv_total            float64  `json:"inv_total"`
	Inv_pay              float64  `json:"inv_pay"`
	Inv_pay_total        float64  `json:"inv_pay_total"`
	Inv_is_active        int      `json:"inv_is_active"`
	Inv_datetime         string   `json:"inv_datetime"`
	Inv_tele_code        string   `json:"inv_tele_code"`
	Inv_create           string   `json:"inv_create"`
	Inv_update           string   `json:"inv_update"`
	DpmId                int      `json:"dpm_id" gorm:"default:null"`
	Inv_discount_type_id int      `json:"inv_discount_type_id"`
	Inv_discount_item    float64  `json:"inv_discount_item"`
	Inv_discount_value   float64  `json:"inv_discount_value"`
	Inv_eclaim_id        *int     `json:"inv_eclaim_id"`
	Inv_eclaim_rate      *float64 `json:"inv_eclaim_rate"`
	Inv_eclaim_over      *float64 `json:"inv_eclaim_over"`
	Inv_eclaim_total     *float64 `json:"inv_eclaim_total"`
}

type InvoiceDetail struct {
	Id                    int      `json:"id"`
	Invoice_id            int      `json:"invoice_id"`
	Course_id             *int     `json:"course_id" `
	Checking_id           *int     `json:"checking_id"`
	Product_id            *int     `json:"product_id"`
	Product_store_id      *int     `json:"product_store_id"`
	Product_unit_id       *int     `json:"product_unit_id"`
	Coin_id               *int     `json:"coin_id"`
	Room_id               *int     `json:"room_id"`
	Queue_id              *int     `json:"queue_id"`
	Order_detail_id       int      `json:"order_detail_id"`
	Invd_type_id          int      `json:"invd_type_id"`
	Invd_code             string   `json:"invd_code"`
	Invd_name             string   `json:"invd_name"`
	Invd_qty              float64  `json:"invd_qty"`
	Invd_set_qty          float64  `json:"invd_set_qty"`
	Invd_limit_qty        float64  `json:"invd_limit_qty"`
	Invd_rate             float64  `json:"invd_rate"`
	Topical_id            *int     `json:"topical_id"`
	Invd_topical          string   `json:"invd_topical"`
	Invd_direction        string   `json:"invd_direction"`
	Invd_unit             string   `json:"invd_unit"`
	Invd_cost             float64  `json:"invd_cost"`
	Invd_price            float64  `json:"invd_price"`
	Invd_discount         float64  `json:"invd_discount"`
	Invd_amount           float64  `json:"invd_amount"`
	Tax_type_id           int      `json:"tax_type_id"`
	Tax_rate              int      `json:"tax_rate"`
	Invd_vat              float64  `json:"invd_vat"`
	Invd_total            float64  `json:"invd_total"`
	Invd_is_set           int      `json:"invd_is_set"`
	Invd_id_set           *int     `json:"invd_id_set"`
	Category_eclaim_id    *int     `json:"category_eclaim_id"`
	Invd_is_active        int      `json:"invd_is_active"`
	Invd_modify           string   `json:"invd_modify"`
	Invd_eclaim           *float64 `json:"invd_eclaim"`
	Invd_discount_type_id int      `json:"invd_discount_type_id"`
	Invd_discount_item    float64  `json:"invd_discount_item"`
}

type InvoiceOrder struct {
	Id                  int      `json:"id"`
	Shop_id             int      `json:"shop_id"`
	User_id             int      `json:"user_id"`
	Customer_id         int      `json:"customer_id"`
	Customer_online_id  int      `json:"customer_online_id"`
	Queue_id            *int     `json:"queue_id"`
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
	Or_tele_code        string   `json:"or_tele_code"`
	Or_create           string   `json:"or_create"`
	Or_update           string   `json:"or_update"`
	DpmId               int      `json:"dpm_id"`
	Or_eclaim_id        *int     `json:"or_eclaim_id"`
	Or_eclaim_rate      *float64 `json:"or_eclaim_rate"`
	Or_eclaim_over      *float64 `json:"or_eclaim_over"`
	Or_eclaim_total     *float64 `json:"or_eclaim_total"`
}

type InvoiceOrderDetail struct {
	Id                   int      `json:"id"`
	Order_id             int      `json:"order_id"`
	Course_id            *int     `json:"invoice_id" `
	Checking_id          *int     `json:"checking_id"`
	Product_id           *int     `json:"product_id"`
	Product_store_id     *int     `json:"product_store_id"`
	Product_unit_id      *int     `json:"product_unit_id"`
	Coin_id              *int     `json:"coin_id"`
	Room_id              *int     `json:"room_id"`
	Queue_id             *int     `json:"queue_id"`
	Ord_type_id          int      `json:"ord_type_id"`
	Ord_code             string   `json:"ord_code"`
	Ord_name             string   `json:"ord_name"`
	Ord_qty              float64  `json:"ord_qty"`
	Ord_set_qty          float64  `json:"ord_set_qty"`
	Ord_limit_qty        float64  `json:"ord_limit_qty"`
	Ord_rate             float64  `json:"ord_rate"`
	Topical_id           *int     `json:"topical_id"`
	Ord_topical          string   `json:"ord_topical"`
	Ord_direction        string   `json:"ord_direction"`
	Ord_unit             string   `json:"ord_unit"`
	Ord_cost             float64  `json:"ord_cost"`
	Ord_price            float64  `json:"ord_price"`
	Ord_discount         float64  `json:"ord_discount"`
	Ord_amount           float64  `json:"ord_amount"`
	Tax_type_id          int      `json:"tax_type_id"`
	Tax_rate             int      `json:"tax_rate"`
	Ord_vat              float64  `json:"ord_vat"`
	Ord_total            float64  `json:"ord_total"`
	Category_eclaim_id   *int     `json:"category_eclaim_id"`
	Ord_is_active        int      `json:"ord_is_active"`
	Ord_is_set           int      `json:"ord_is_set"`
	Ord_id_set           *int     `json:"ord_id_set"`
	Ord_modify           string   `json:"ord_modify"`
	Ord_eclaim           *float64 `json:"ord_eclaim"`
	Ord_discount_type_id int      `json:"ord_discount_type_id"`
	Ord_discount_item    float64  `json:"ord_discount_item"`
}

type InvoiceId struct {
	Id                      int              `json:"id"`
	Shop_id                 int              `json:"shop_id"`
	User_id                 int              `json:"user_id"`
	User_fullname           string           `json:"user_fullname"`
	User_fullname_en        string           `json:"user_fullname_en"`
	Customer_id             int              `json:"customer_id"`
	Queue_id                int              `json:"queue_id"`
	Order_id                int              `json:"order_id"`
	Inv_code                string           `json:"inv_code"`
	Inv_fullname            string           `json:"inv_fullname"`
	Inv_tel                 string           `json:"inv_tel"`
	Inv_email               string           `json:"inv_email"`
	Inv_address             string           `json:"inv_address"`
	Inv_district            string           `json:"inv_district"`
	Inv_amphoe              string           `json:"inv_amphoe"`
	Inv_province            string           `json:"inv_province"`
	Inv_zipcode             string           `json:"inv_zipcode"`
	Inv_comment             string           `json:"inv_comment"`
	Inv_total_price         float64          `json:"inv_total_price"`
	Inv_discount            float64          `json:"inv_discount"`
	Inv_befor_vat           float64          `json:"inv_befor_vat"`
	Tax_type_id             int              `json:"tax_type_id"`
	Tax_rate                int              `json:"tax_rate"`
	Inv_vat                 float64          `json:"inv_vat"`
	Inv_total               float64          `json:"inv_total"`
	Inv_is_active           int              `json:"inv_is_active"`
	User_id_cancel          int              `json:"user_id_cancel"`
	User_fullname_cancel    string           `json:"user_fullname_cancel"`
	User_fullname_en_cancel string           `json:"user_fullname_en_cancel"`
	Inv_datetime            string           `json:"inv_datetime"`
	Inv_create              string           `json:"inv_create"`
	Inv_update              string           `json:"inv_update"`
	Shop                    ReceiptShop      `json:"shop" gorm:"-"`
	Customer                ObjQueryCustomer `json:"customer" gorm:"-"`
	Tags                    *[]OrderTags     `json:"tags" gorm:"-"`
	Subs                    *[]InvoiceSub    `json:"subs" gorm:"-"`
	DpmId                   int              `json:"dpm_id"`
	Inv_discount_type_id    int              `json:"inv_discount_type_id"`
	Inv_discount_item       float64          `json:"inv_discount_item"`
	Inv_discount_value      float64          `json:"inv_discount_value"`
	Inv_eclaim_id           *int             `json:"inv_eclaim_id"`
	Inv_eclaim_rate         *float64         `json:"inv_eclaim_rate"`
	Inv_eclaim_over         *float64         `json:"inv_eclaim_over"`
	Inv_eclaim_total        *float64         `json:"inv_eclaim_total"`
}

type InvoiceSub struct {
	Id                    int                `json:"id"`
	Invoice_id            int                `json:"invoice_id"`
	Course_id             *int               `json:"course_id" `
	Checking_id           *int               `json:"checking_id"`
	Product_id            *int               `json:"product_id"`
	Product_unit_id       *int               `json:"product_unit_id"`
	Coin_id               *int               `json:"coin_id"`
	Room_id               *int               `json:"room_id"`
	Order_detail_id       int                `json:"order_detail_id"`
	Invd_type_id          int                `json:"invd_type_id"`
	Invd_code             string             `json:"invd_code"`
	Invd_name             string             `json:"invd_name"`
	Invd_qty              float64            `json:"invd_qty"`
	Invd_set_qty          float64            `json:"invd_set_qty"`
	Invd_limit_qty        float64            `json:"invd_limit_qty"`
	Invd_unit             string             `json:"invd_unit"`
	U_name_en             string             `json:"u_name_en"`
	Invd_cost             float64            `json:"invd_cost"`
	Invd_price            float64            `json:"invd_price"`
	Invd_discount         float64            `json:"invd_discount"`
	Invd_amount           float64            `json:"invd_amount"`
	Tax_type_id           int                `json:"tax_type_id"`
	Tax_rate              int                `json:"tax_rate"`
	Invd_vat              float64            `json:"invd_vat"`
	Invd_is_set           int                `json:"invd_is_set"`
	Invd_id_set           *int               `json:"invd_id_set"`
	Invd_total            float64            `json:"invd_total"`
	Invd_is_active        int                `json:"invd_is_active"`
	Invd_modify           string             `json:"invd_modify"`
	Room_code             string             `json:"room_code"`
	Room_th               string             `json:"room_th"`
	Room_en               string             `json:"room_en"`
	Room_type_th          string             `json:"room_type_th"`
	Room_type_en          string             `json:"room_type_en"`
	Units                 *[]ProductUnitList `json:"units" gorm:"-"`
	Invd_eclaim           *float64           `json:"invd_eclaim"`
	Invd_discount_type_id int                `json:"invd_discount_type_id"`
	Invd_discount_item    float64            `json:"invd_discount_item"`
	Course_code           string             `json:"course_code"`
	Course_name           string             `json:"course_name"`
	Checking_code         string             `json:"checking_code"`
	Checking_name         string             `json:"checking_name"`
}

type DocInvoice struct {
	ShopId                 int    `json:"shop_id"`
	Invoice_id_default     string `json:"invoice_id_default"`
	Invoice_number_default string `json:"invoice_number_default"`
	Invoice_number_digit   int    `json:"invoice_number_digit"`
	Invoice_type           int    `json:"invoice_type"`
}

type LogInvoice struct {
	Username   string `json:"username"`
	Log_type   string `json:"log_type"`
	Log_text   string `json:"log_text"`
	Log_create string `json:"log_create"`
	Shop_id    int    `json:"shop_id"`
}

type OrderInvoiceUpdate struct {
	Id           int    `json:"id"`
	Or_is_active int    `json:"or_is_active"`
	Or_update    string `json:"or_update"`
}

type CategoryEclaimId struct {
	Category_eclaim_id *int `json:"category_eclaim_id"`
}

// Add Sticker
type StickerInvoiceDetail struct {
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
	Invd_rate        float64 `json:"invd_rate"`
	Invd_set_qty     float64 `json:"invd_set_qty"`
	Invd_limit_qty   float64 `json:"invd_limit_qty"`
	Invd_unit        string  `json:"invd_unit"`
	Invd_cost        float64 `json:"invd_cost"`
	Invd_price       float64 `json:"invd_price"`
	Invd_discount    float64 `json:"invd_discount"`
	Invd_amount      float64 `json:"invd_amount"`
	Invd_vat         float64 `json:"invd_vat"`
	Topical_id       string  `json:"topical_id"`
	Invd_topical     string  `json:"invd_topical"`
	Invd_direction   string  `json:"invd_direction"`
	Invd_total       float64 `json:"invd_total"`
	Invd_is_set      int     `json:"invd_is_set"`
	Invd_is_active   int     `json:"invd_is_active"`
	Invd_modify      string  `json:"invd_modify"`
}

type StickerProcessProduct struct {
	Shop_id           int     `json:"shop_id"`
	Customer_id       int     `json:"customer_id"`
	User_id           int     `json:"user_id"`
	Product_id        int     `json:"product_id"`
	Product_store_id  int     `json:"product_store_id"`
	Invoice_id        *int    `json:"invoice_id"`
	Invoice_detail_id *int    `json:"invoice_detail_id"`
	Pdso_qty          float64 `json:"pdso_qty"`
	Invd_code         string  `json:"invd_code"`
	Invd_name         string  `json:"invd_name"`
	Invd_unit         string  `json:"invd_unit"`
	Invd_price        float64 `json:"invd_price"`
	Invd_topical      string  `json:"invd_topical"`
	Invd_direction    string  `json:"invd_direction"`
}

type ProcessProductStickerInvoice struct {
	Id                  int     `json:"id"`
	Product_id          int     `json:"product_id"`
	Shop_id             int     `json:"shop_id"`
	Customer_id         int     `json:"customer_id"`
	User_id             int     `json:"user_id"`
	Invoice_id          *int    `json:"invoice_id"`
	Invoice_detail_id   *int    `json:"invoice_detail_id"`
	Receipt_id          *int    `json:"receipt_id"`
	Receipt_detail_id   *int    `json:"receipt_detail_id"`
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

// check
type CheckChecking struct {
	Id               int `json:"id"`
	Checking_type_id int `json:"checking_type_id"`
}

type InvoiceCheck struct {
	Id                int     `json:"id"`
	Shop_id           int     `json:"shop_id"`
	Invoice_id        int     `json:"invoice_id" `
	Invoice_detail_id int     `json:"invoice_detail_id"`
	User_id           int     `json:"user_id"`
	Customer_id       int     `json:"customer_id"`
	Queue_id          *int    `json:"queue_id"`
	Checking_id       int     `json:"checking_id"`
	Chk_type_id       int     `json:"chk_type_id"`
	Chk_code          string  `json:"chk_code"`
	Chk_name          string  `json:"chk_name"`
	Chk_unit          string  `json:"chk_unit"`
	Chk_value         *string `json:"chk_value"`
	Chk_upload        *string `json:"chk_upload"`
	Chk_upload_size   *int    `json:"chk_upload_size"`
	Direction_id      *int    `json:"direction_id"`
	Chk_flag          *string `json:"chk_flag"`
	Chk_date          *string `json:"chk_date"`
	Chk_is_print      *int    `json:"chk_is_print"`
	Chk_is_report     *int    `json:"chk_is_report"`
	Chk_is_active     int     `json:"chk_is_active"`
	Chk_datetime      string  `json:"chk_datetime"`
	Chk_create        string  `json:"chk_create"`
	Chk_update        string  `json:"chk_update"`
}

type GetQueueDateTime struct {
	Que_datetime string `json:"que_datetime"`
}

type InvoiceDetailCheck struct {
	Id               int     `json:"id"`
	Product_id       *int    `json:"product_id"`
	Product_store_id *int    `json:"product_store_id"`
	Invd_code        string  `json:"invd_code"`
	Invd_name        string  `json:"invd_name"`
	Invd_qty         float64 `json:"invd_qty"`
}

type ProcessProductStoreOrderCheck struct {
	Id         int     `json:"id"`
	Pdso_total float64 `json:"pdso_total"`
}

type InvoiceDetailCheckCourse struct {
	Id       int     `json:"id"`
	Invd_qty float64 `json:"invd_qty"`
}
