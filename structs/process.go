package structs

type ObjPayloadProcess struct {
	Id int `json:"id" binding:"required"`
}

type ObjPayloadProcessOrder struct {
	Receipt_id int `json:"receipt_id" binding:"required"`
}

type ProcessReceipt struct {
	Id                int     `json:"id"`
	Shop_id           int     `json:"shop_id"`
	User_id           int     `json:"user_id"`
	Customer_id       int     `json:"customer_id"`
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
	Rec_period        float64 `json:"rec_period"`
	Rec_pay           float64 `json:"rec_pay"`
	Rec_balance       float64 `json:"rec_balance"`
	Rec_pay_total     float64 `json:"rec_pay_total"`
	Rec_discription   string  `json:"rec_discription"`
	Rec_account       int     `json:"rec_account"`
	Rec_is_process    int     `json:"rec_is_process"`
	Rec_is_active     int     `json:"rec_is_active"`
	Rec_user_id       int     `json:"rec_user_id"`
	Rec_user_fullname string  `json:"rec_user_fullname"`
	Rec_pay_datetime  string  `json:"rec_pay_datetime"`
	Rec_create        string  `json:"rec_create"`
	Rec_update        string  `json:"rec_update"`
}

type ProcessReceiptDetail struct {
	Id                int     `json:"id"`
	Receipt_id        int     `json:"receipt_id"`
	Course_id         *int    `json:"course_id" `
	Checking_id       *int    `json:"checking_id"`
	Product_id        *int    `json:"product_id"`
	Product_store_id  *int    `json:"product_store_id"`
	Product_unit_id   *int    `json:"product_unit_id"`
	Coin_id           *int    `json:"coin_id"`
	Room_id           *int    `json:"room_id"`
	Queue_id          *int    `json:"queue_id"`
	Invoice_detail_id int     `json:"invoice_detail_id"`
	Recd_type_id      int     `json:"recd_type_id"`
	Recd_code         string  `json:"recd_code"`
	Recd_name         string  `json:"recd_name"`
	Recd_qty          float64 `json:"recd_qty"`
	Recd_rate         float64 `json:"recd_rate"`
	Recd_set_qty      float64 `json:"recd_set_qty"`
	Recd_limit_qty    float64 `json:"recd_limit_qty"`
	Recd_unit         string  `json:"recd_unit"`
	Recd_cost         float64 `json:"recd_cost"`
	Recd_price        float64 `json:"recd_price"`
	Recd_discount     float64 `json:"recd_discount"`
	Recd_amount       float64 `json:"recd_amount"`
	Tax_type_id       int     `json:"tax_type_id"`
	Tax_rate          int     `json:"tax_rate"`
	Recd_vat          float64 `json:"recd_vat"`
	Topical_id        string  `json:"topical_id"`
	Recd_topical      string  `json:"recd_topical"`
	Recd_direction    string  `json:"recd_direction"`
	Recd_total        float64 `json:"recd_total"`
	Recd_is_set       int     `json:"recd_is_set"`
	Recd_is_active    int     `json:"recd_is_active"`
	Recd_modify       string  `json:"recd_modify"`
}

type ProcessCheck struct {
	Id                int     `json:"id"`
	Shop_id           int     `json:"shop_id"`
	Receipt_id        int     `json:"receipt_id" `
	Receipt_detail_id int     `json:"receipt_detail_id"`
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

type ProcessCheckProduct struct {
	Id               int     `json:"id"`
	Shop_id          int     `json:"shop_id"`
	Check_id         int     `json:"check_id" `
	Checking_id      int     `json:"checking_id" `
	Queue_id         *int    `json:"queue_id"`
	Receipt_id       int     `json:"receipt_id" `
	Customer_id      int     `json:"customer_id"`
	Product_id       int     `json:"product_id"`
	Product_store_id int     `json:"product_store_id"`
	Product_unit_id  int     `json:"product_unit_id"`
	Chkp_code        string  `json:"chkp_code"`
	Chkp_name        string  `json:"chkp_name"`
	Chkp_qty         float64 `json:"chkp_qty"`
	Chkp_unit        string  `json:"chkp_unit"`
	Chkp_is_active   int     `json:"chkp_is_active"`
	Chkp_datetime    string  `json:"chkp_datetime"`
	Chkp_create      string  `json:"chkp_create"`
	Chkp_modify      string  `json:"chkp_modify"`
}

type ProcessService struct {
	Id                int     `json:"id"`
	Shop_id           int     `json:"shop_id"`
	Shop_mother_id    int     `json:"shop_mother_id"`
	Receipt_id        int     `json:"receipt_id" `
	Receipt_detail_id int     `json:"receipt_detail_id"`
	User_id           int     `json:"user_id"`
	Ser_customer_id   int     `json:"ser_customer_id"`
	Customer_id       int     `json:"customer_id"`
	Course_id         int     `json:"course_id"`
	Ser_code          string  `json:"ser_code"`
	Ser_name          string  `json:"ser_name"`
	Ser_lock_drug     int     `json:"ser_lock_drug"`
	Ser_qty           int     `json:"ser_qty"`
	Ser_unit          string  `json:"ser_unit"`
	Ser_use_date      int     `json:"ser_use_date"`
	Ser_exp           int     `json:"ser_exp"`
	Ser_exp_date      *string `json:"ser_exp_date"`
	Ser_use           int     `json:"ser_use"`
	Ser_price_total   float64 `json:"ser_price_total"`
	Ser_is_active     int     `json:"ser_is_active"`
	Ser_datetime      string  `json:"ser_datetime"`
	Ser_create        string  `json:"ser_create"`
	Ser_update        string  `json:"ser_update"`
}

type ProcessServiceUpdate struct {
	Ser_is_active int    `json:"ser_is_active"`
	Ser_update    string `json:"ser_update"`
}

type ProcessServiceProduct struct {
	Id                int     `json:"id"`
	Shop_id           int     `json:"shop_id"`
	Service_id        int     `json:"service_id" `
	Course_id         int     `json:"course_id"`
	Receipt_id        int     `json:"receipt_id"`
	Receipt_detail_id int     `json:"receipt_detail_id"`
	Product_id        int     `json:"product_id"`
	Product_store_id  int     `json:"product_store_id"`
	Product_unit_id   int     `json:"product_unit_id"`
	Serp_code         string  `json:"serp_code"`
	Serp_name         string  `json:"serp_name"`
	Serp_qty          float64 `json:"serp_qty"`
	Serp_use          float64 `json:"serp_use"`
	Serp_tranfer      float64 `json:"serp_tranfer"`
	Serp_balance      float64 `json:"serp_balance"`
	Serp_unit         string  `json:"serp_unit"`
	Serp_lock_drug    int     `json:"serp_lock_drug"`
	Serp_use_set_qty  float64 `json:"serp_use_set_qty"`
	Serp_is_active    int     `json:"serp_is_active"`
	Serp_datetime     string  `json:"serp_datetime"`
	Serp_create       string  `json:"serp_create"`
	Serp_modify       string  `json:"serp_modify"`
}

type ProcessServiceUsed struct {
	Id             int     `json:"id"`
	Shop_id        int     `json:"shop_id"`
	Shop_mother_id int     `json:"shop_mother_id"`
	Shop_used_id   int     `json:"shop_used_id"`
	Service_id     int     `json:"service_id"`
	Queue_id       int     `json:"queue_id"`
	Receipt_id     int     `json:"receipt_id"`
	Course_id      int     `json:"course_id"`
	Customer_id    int     `json:"customer_id"`
	User_id        int     `json:"user_id"`
	Seru_code      string  `json:"seru_code"`
	Seru_name      string  `json:"seru_name"`
	Seru_qty       int     `json:"seru_qty"`
	Seru_unit      string  `json:"seru_unit"`
	Seru_cost      float64 `json:"seru_cost"`
	Seru_date      string  `json:"seru_date"`
	Seru_is_active int     `json:"seru_is_active"`
	Seru_datetime  string  `json:"seru_datetime"`
	Seru_create    string  `json:"seru_create"`
	Seru_update    string  `json:"seru_update"`
}

type ProcessServiceProductUsed struct {
	Id                 int     `json:"id"`
	Shop_id            int     `json:"shop_id"`
	Service_id         int     `json:"service_id" `
	Service_used_id    int     `json:"service_used_id"`
	Course_id          int     `json:"course_id"`
	Queue_id           int     `json:"queue_id"`
	Receipt_id         int     `json:"receipt_id"`
	Customer_id        int     `json:"customer_id"`
	Service_product_id int     `json:"service_product_id"`
	Product_id         int     `json:"product_id"`
	Product_store_id   int     `json:"product_store_id"`
	Product_unit_id    int     `json:"product_unit_id"`
	Serpu_code         string  `json:"serpu_code"`
	Serpu_name         string  `json:"serpu_name"`
	Serpu_qty          float64 `json:"serpu_qty"`
	Serpu_unit         string  `json:"serpu_unit"`
	Serpu_is_active    int     `json:"serpu_is_active"`
	Serpu_datetime     string  `json:"serpu_datetime"`
	Serpu_create       string  `json:"serpu_create"`
	Serpu_modify       string  `json:"serpu_modify"`
}

type ProcessCheckChecking struct {
	Id               int `json:"id"`
	Checking_type_id int `json:"checking_type_id"`
}

type ProcessCourseService struct {
	Id               int     `json:"id"`
	Course_lock_drug int     `json:"course_lock_drug"`
	Course_cost      float64 `json:"course_cost"`
	Course_amount    int     `json:"course_amount"`
	Course_use_date  int     `json:"course_use_date"`
	Course_exp_date  int     `json:"course_exp_date"`
}

type ProcessCourseProductSet struct {
	Id         int     `json:"id"`
	Product_id int     `json:"product_id"`
	Cp_amount  float64 `json:"cp_amount"`
}

type ProcessProductSet struct {
	Id               int     `json:"id"`
	Product_id       int     `json:"product_id"`
	Product_store_id int     `json:"product_store_id"`
	Product_units_id int     `json:"product_units_id"`
	Pd_code          string  `json:"pd_code"`
	Pd_name          string  `json:"pd_name"`
	U_name           string  `json:"u_name"`
	Pu_amount        int     `json:"pu_amount"`
	Psp_price_ipd    float64 `json:"psp_price_ipd"`
	Psp_price_opd    float64 `json:"psp_price_opd"`
	Pds_cost         float64 `json:"pds_cost"`
}

type ProcessProductStoreOrder struct {
	Id               int     `json:"id"`
	Product_store_id int     `json:"product_store_id"`
	Pdso_expire      string  `json:"pdso_expire"`
	Pdso_out         float64 `json:"pdso_out"`
	Pdso_use         float64 `json:"pdso_use"`
	Pdso_total       float64 `json:"pdso_total"`
	Pdso_update      string  `json:"pdso_update"`
}

type ProcessProductStore struct {
	Id            int     `json:"id"`
	Shop_store_id int     `json:"shop_store_id"`
	Pds_out       float64 `json:"pds_out"`
	Pds_total     float64 `json:"pds_total"`
	Pds_update    string  `json:"pds_update"`
}

type ProcessProductStoreHistory struct {
	Id                      int     `json:"id"`
	Shop_id                 int     `json:"shop_id"`
	Shop_store_id           int     `json:"shop_store_id"`
	Product_store_id        int     `json:"product_store_id"`
	Product_store_order_id  int     `json:"product_store_order_id"`
	Receipt_id              *int    `json:"receipt_id"`
	Receipt_detail_id       *int    `json:"receipt_detail_id"`
	Check_product_id        *int    `json:"check_product_id"`
	Service_product_used_id *int    `json:"service_product_used_id"`
	Queue_id                *int    `json:"queue_id"`
	Pdsh_in                 float64 `json:"pdsh_in"`
	Pdsh_out                float64 `json:"pdsh_out"`
	Pdsh_order_forward      float64 `json:"pdsh_order_forward"`
	Pdsh_total_forward      float64 `json:"pdsh_total_forward"`
	Pdsh_amount             float64 `json:"pdsh_amount"`
	Pdsh_order_total        float64 `json:"pdsh_order_total"`
	Pdsh_total              float64 `json:"pdsh_total"`
	Pdsh_inout              int     `json:"pdsh_inout"`
	Pdsh_out_id             int     `json:"pdsh_out_id"`
	Pdsh_ref_doc_no         string  `json:"pdsh_ref_doc_no"`
	Pdsh_comment            string  `json:"pdsh_comment"`
	Pdsh_date               string  `json:"pdsh_date"`
	Pdsh_modify             string  `json:"pdsh_modify"`
	User_id                 int     `json:"user_id"`
	Product_id              int     `json:"product_id"`
	Pd_code                 string  `json:"pd_code"`
	Pd_name                 string  `json:"pd_name"`
	Pdsh_type_id            int     `json:"pdsh_type_id"`
	Customer_id             int     `json:"customer_id"`
	Ctm_prefix              string  `json:"ctm_prefix"`
	Ctm_fname               string  `json:"ctm_fname"`
	Ctm_lname               string  `json:"ctm_lname"`
	Ctm_fname_en            string  `json:"ctm_fname_en"`
	Ctm_lname_en            string  `json:"ctm_lname_en"`
	Ctm_tel                 string  `json:"ctm_tel"`
	Ctm_gender              string  `json:"ctm_gender"`
}

type ProcessProductStoreHistoryExp struct {
	Id                      int     `json:"id"`
	Shop_id                 int     `json:"shop_id"`
	Shop_store_id           int     `json:"shop_store_id"`
	Product_store_id        int     `json:"product_store_id"`
	Product_store_order_id  int     `json:"product_store_order_id"`
	Receipt_id              *int    `json:"receipt_id"`
	Receipt_detail_id       *int    `json:"receipt_detail_id"`
	Check_product_id        *int    `json:"check_product_id"`
	Service_product_used_id *int    `json:"service_product_used_id"`
	Queue_id                *int    `json:"queue_id"`
	Pdsh_in                 float64 `json:"pdsh_in"`
	Pdsh_out                float64 `json:"pdsh_out"`
	Pdsh_order_forward      float64 `json:"pdsh_order_forward"`
	Pdsh_total_forward      float64 `json:"pdsh_total_forward"`
	Pdsh_amount             float64 `json:"pdsh_amount"`
	Pdsh_order_total        float64 `json:"pdsh_order_total"`
	Pdsh_total              float64 `json:"pdsh_total"`
	Pdsh_inout              int     `json:"pdsh_inout"`
	Pdsh_out_id             int     `json:"pdsh_out_id"`
	Pdsh_ref_doc_no         string  `json:"pdsh_ref_doc_no"`
	Product_id              int     `json:"product_id"`
	Pd_code                 string  `json:"pd_code"`
	Pd_name                 string  `json:"pd_name"`
	Pdsh_type_id            int     `json:"pdsh_type_id"`
	Pdsh_comment            string  `json:"pdsh_comment"`
	Pdsh_date               string  `json:"pdsh_date"`
	Pdsh_modify             string  `json:"pdsh_modify"`
}

type ObjPayloadProcessProduct struct {
	Customer_id             int     `json:"customer_id"`
	User_id                 int     `json:"user_id"`
	Product_id              int     `json:"product_id"`
	Product_store_id        int     `json:"product_store_id"`
	Receipt_id              *int    `json:"receipt_id"`
	Receipt_detail_id       *int    `json:"receipt_detail_id"`
	Check_product_id        *int    `json:"check_product_id"`
	Service_product_used_id *int    `json:"service_product_used_id"`
	Queue_id                *int    `json:"queue_id"`
	Pdsh_out_id             int     `json:"pdsh_out_id"`
	Pdsh_ref_doc_no         string  `json:"pdsh_ref_doc_no"`
	Pdso_qty                float64 `json:"pdso_qty"`
	Recd_code               string  `json:"recd_code"`
	Recd_name               string  `json:"recd_name"`
	Recd_unit               string  `json:"recd_unit"`
	Recd_price              float64 `json:"recd_price"`
	Recd_topical            string  `json:"recd_topical"`
	Recd_direction          string  `json:"recd_direction"`
}

type ProcessProduct struct {
	Shop_id                 int     `json:"shop_id"`
	Customer_id             int     `json:"customer_id"`
	User_id                 int     `json:"user_id"`
	Product_id              int     `json:"product_id"`
	Product_store_id        int     `json:"product_store_id"`
	Receipt_id              *int    `json:"receipt_id"`
	Receipt_detail_id       *int    `json:"receipt_detail_id"`
	Check_product_id        *int    `json:"check_product_id"`
	Service_product_used_id *int    `json:"service_product_used_id"`
	Queue_id                *int    `json:"queue_id"`
	Pdsh_out_id             int     `json:"pdsh_out_id"`
	Pdsh_ref_doc_no         string  `json:"pdsh_ref_doc_no"`
	Pdso_qty                float64 `json:"pdso_qty"`
	Recd_code               string  `json:"recd_code"`
	Recd_name               string  `json:"recd_name"`
	Recd_unit               string  `json:"recd_unit"`
	Recd_price              float64 `json:"recd_price"`
	Recd_topical            string  `json:"recd_topical"`
	Recd_direction          string  `json:"recd_direction"`
	Pdsh_type_id            int     `json:"pdsh_type_id"`
	Ctm_prefix              string  `json:"ctm_prefix"`
	Ctm_fname               string  `json:"ctm_fname"`
	Ctm_lname               string  `json:"ctm_lname"`
	Ctm_fname_en            string  `json:"ctm_fname_en"`
	Ctm_lname_en            string  `json:"ctm_lname_en"`
	Ctm_tel                 string  `json:"ctm_tel"`
	Ctm_gender              string  `json:"ctm_gender"`
}

type ProcessProductSticker struct {
	Id                  int     `json:"id"`
	Shop_id             int     `json:"shop_id"`
	Product_id          int     `json:"product_id"`
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

type CancelProductReceipt struct {
	Shop_id     int  `json:"shop_id"`
	Customer_id int  `json:"customer_id"`
	User_id     int  `json:"user_id"`
	Receipt_id  int  `json:"receipt_id"`
	Queue_id    *int `json:"queue_id"`
}

type CancelProduct struct {
	Shop_id                 int  `json:"shop_id"`
	Customer_id             int  `json:"customer_id"`
	User_id                 int  `json:"user_id"`
	Receipt_id              *int `json:"receipt_id"`
	Receipt_detail_id       *int `json:"receipt_detail_id"`
	Check_product_id        *int `json:"check_product_id"`
	Service_product_used_id *int `json:"service_product_used_id"`
	Queue_id                *int `json:"queue_id"`
}

type ProcessCoin struct {
	Id                int     `json:"id"`
	Ctm_coin          float64 `json:"ctm_coin"`
	Customer_group_id int     `json:"customer_group_id"`
	Ctm_update        string  `json:"ctm_update"`
}

type ProcessCoinCancelUpdate struct {
	Id         int     `json:"id"`
	Ctm_coin   float64 `json:"ctm_coin"`
	Ctm_update string  `json:"ctm_update"`
}

type CoinCustomerGroupId struct {
	Id                int `json:"id"`
	Customer_group_id int `json:"customer_group_id"`
}

type ProcessProductType struct {
	Id          int    `json:"id"`
	Pd_type_id  int    `json:"pd_type_id"`
	Pd_name_acc string `json:"pd_name_acc"`
}

type ProcessQueue struct {
	Id            int    `json:"id"`
	Customer_id   int    `json:"customer_id"`
	Que_code      string `json:"que_code"`
	Que_status_id int    `json:"que_status_id"`
	Que_datetime  string `json:"que_datetime"`
	Que_update    string `json:"que_update"`
}

type QueueProcessUpdate struct {
	Que_status_id    int    `json:"que_status_id"`
	Que_datetime_out string `json:"que_datetime_out"`
	Que_time_end     int    `json:"que_time_end"`
	Que_update       string `json:"que_update"`
}

type ProcessInvoice struct {
	Id           int    `json:"id"`
	Inv_datetime string `json:"inv_datetime"`
}

type InvoiceProcessUpdate struct {
	Inv_time_end int    `json:"inv_time_end"`
	Inv_update   string `json:"inv_update"`
}

type CheckProductUpdate struct {
	Chkp_is_active int    `json:"chkp_is_active"`
	Chkp_modify    string `json:"chkp_modify"`
}

type CheckUpdate struct {
	Chk_is_active int    `json:"chk_is_active"`
	Chk_update    string `json:"chk_update"`
}

type ServiceProductUsedUpdate struct {
	Serpu_is_active int    `json:"serpu_is_active"`
	Serpu_modify    string `json:"serpu_modify"`
}

type ServiceProductUpdate struct {
	Serp_is_active int    `json:"serp_is_active"`
	Serp_modify    string `json:"serp_modify"`
}

type ServiceUsedUpdate struct {
	Seru_is_active int    `json:"seru_is_active"`
	Seru_update    string `json:"seru_update"`
}

type StickerUpdate struct {
	Sticker_is_del int `json:"sticker_is_del"`
}

type ProcessProductStoreOrderExpire struct {
	Id               int     `json:"id"`
	Product_store_id int     `json:"product_store_id"`
	Pdso_expire      string  `json:"pdso_expire"`
	Pdso_exp         float64 `json:"Pdso_exp"`
	Pdso_total       float64 `json:"pdso_total"`
	Pdso_update      string  `json:"pdso_update"`
	// Pdso_cost        float64 `json:"pdso_cost"`
}

type ProcessProductStoreExpire struct {
	Id            int     `json:"id"`
	Shop_id       int     `json:"shop_id"`
	Shop_store_id int     `json:"shop_store_id"`
	Product_id    int     `json:"product_id"`
	Pds_exp       float64 `json:"pds_exp"`
	Pds_total     float64 `json:"pds_total"`
	Pds_update    string  `json:"pds_update"`
}

type UpdateProcessProductStoreExpire struct {
	Id            int     `json:"id"`
	Shop_store_id int     `json:"shop_store_id"`
	Pds_exp       float64 `json:"pds_exp"`
	Pds_total     float64 `json:"pds_total"`
	Pds_update    string  `json:"pds_update"`
}

type ProductIdProcess struct {
	Id      int    `json:"id"`
	Pd_code string `json:"pd_code"`
	Pd_name string `json:"pd_name"`
}

type ProductRate struct {
	Id          int    `json:"id"`
	Pd_type_id  int    `json:"pd_type_id"`
	Pd_name_acc string `json:"pd_name_acc"`
	Pu_rate     int    `json:"pu_rate"`
	U_name      string `json:"u_name"`
	U_name_en   string `json:"u_name_en"`
}

type CustomerIdProduct struct {
	Id           int    `json:"id"`
	Ctm_prefix   string `json:"ctm_prefix"`
	Ctm_fname    string `json:"ctm_fname"`
	Ctm_lname    string `json:"ctm_lname"`
	Ctm_fname_en string `json:"ctm_fname_en"`
	Ctm_lname_en string `json:"ctm_lname_en"`
	Ctm_tel      string `json:"ctm_tel"`
	Ctm_gender   string `json:"ctm_gender"`
}

type ProcessServiceQueueCourse struct {
	Queue_course_id int `json:"queue_course_id"`
}

type ProcessServiceQueueCourseUpdate struct {
	Service_id int `json:"service_id"`
	Receipt_id int `json:"receipt_id"`
}

type ProcessInvoiceDetailById struct {
	Invoice_id        int `json:"invoice_id"`
	Invoice_detail_id int `json:"invoice_detail_id"`
}
