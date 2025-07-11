package structs

type ObjPayloadServiceQueueProduct struct {
	Queue_id         int     `json:"queue_id" binding:"required"`
	Product_id       int     `json:"product_id" binding:"required"`
	Product_store_id int     `json:"product_store_id"`
	Product_unit_id  int     `json:"product_unit_id"`
	Quep_qty         float64 `json:"quep_qty"`
	Quep_unit        string  `json:"quep_unit"`
	// Que_admis_id     int     `json:"que_admis_id" binding:"required"` // 1 = IPD, 2 = OPD
	// Quep_price       float64 `json:"quep_price"`
	// Quep_discount    float64 `json:"quep_discount"`
}

type ServiceQueueItemProduct struct {
	Id         int    `json:"id"`
	Pd_type_id int    `json:"pd_type_id"`
	Pd_code    string `json:"pd_code"`
	Pd_name    string `json:"pd_name"`
}

type ServiceQueueProductSet struct {
	Id               int     `json:"id"`
	Product_id       int     `json:"product_id"`
	Product_store_id int     `json:"product_store_id"`
	Product_units_id int     `json:"product_units_id"`
	Pd_code          string  `json:"pd_code"`
	Pd_name          string  `json:"pd_name"`
	Topical_id       *int    `json:"topical_id"`
	Drug_direction   string  `json:"drug_direction"`
	U_name           string  `json:"u_name"`
	Pu_amount        int     `json:"pu_amount"`
	Psp_price_ipd    float64 `json:"psp_price_ipd"`
	Psp_price_opd    float64 `json:"psp_price_opd"`
	Pds_cost         float64 `json:"pds_cost"`
}

type ObjPayloadServiceQueueCourse struct {
	Queue_id   int     `json:"queue_id" binding:"required"`
	Course_id  int     `json:"course_id" binding:"required"`
	Service_id int     `json:"service_id" binding:"required"`
	Receipt_id int     `json:"receipt_id"`
	Quec_qty   float64 `json:"quec_qty"`
	// Que_admis_id  int     `json:"que_admis_id" binding:"required"`
	// Quec_unit     string  `json:"quec_unit"`
	// Quec_price    float64 `json:"quec_price"`
	// Quec_discount float64 `json:"quec_discount"`
}

type ServiceQueueProduct struct {
	Id                 *int    `json:"id"`
	Queue_id           int     `json:"queue_id"`
	User_id            int     `json:"user_id"`
	Product_id         int     `json:"product_id"`
	Product_store_id   int     `json:"product_store_id"`
	Product_unit_id    int     `json:"product_unit_id"`
	Quep_type_id       int     `json:"quep_type_id"`
	Queue_checking_id  *int    `json:"queue_checking_id"`
	Checking_id        *int    `json:"checking_id"`
	Queue_course_id    *int    `json:"queue_course_id"`
	Course_id          *int    `json:"course_id"`
	Service_id         *int    `json:"service_id"`
	Service_product_id *int    `json:"service_product_id"`
	Quep_code          string  `json:"quep_code"`
	Quep_name          string  `json:"quep_name"`
	Quep_qty           float64 `json:"quep_qty"`
	Quep_set_qty       float64 `json:"quep_set_qty"`
	Quep_limit_qty     float64 `json:"quep_limit_qty"`
	Quep_unit          string  `json:"quep_unit"`
	Quep_cost          float64 `json:"quep_cost"`
	Quep_price         float64 `json:"quep_price"`
	Quep_discount      float64 `json:"quep_discount"`
	Quep_total         float64 `json:"quep_total"`
	Quep_is_set        int     `json:"quep_is_set"`
	Quep_is_active     int     `json:"quep_is_active"`
	Quep_modify        string  `json:"quep_modify"`
}

type ServiceQueueItemCourse struct {
	Id               *int    `json:"id"`
	Course_type_id   int     `json:"course_type_id"`
	Course_code      string  `json:"course_code"`
	Course_name      string  `json:"course_name"`
	Course_lock_drug int     `json:"course_lock_drug"`
	Course_cost      float64 `json:"course_cost"`
}

type ServiceQueueCourse struct {
	Id             *int    `json:"id"`
	Queue_id       int     `json:"queue_id"`
	Course_id      int     `json:"course_id"`
	User_id        int     `json:"user_id"`
	Service_id     int     `json:"service_id"`
	Receipt_id     int     `json:"receipt_id"`
	Quec_code      string  `json:"quec_code"`
	Quec_name      string  `json:"quec_name"`
	Quec_qty       float64 `json:"quec_qty"`
	Quec_unit      string  `json:"quec_unit"`
	Quec_cost      float64 `json:"quec_cost"`
	Quec_price     float64 `json:"quec_price"`
	Quec_discount  float64 `json:"quec_discount"`
	Quec_total     float64 `json:"quec_total"`
	Quec_is_set    int     `json:"quec_is_set"`
	Quec_is_active int     `json:"quec_is_active"`
	Quec_modify    string  `json:"quec_modify"`
}

type ServiceQueueCourseReceipt struct {
	Id             *int    `json:"id"`
	Queue_id       int     `json:"queue_id"`
	Course_id      int     `json:"course_id"`
	User_id        int     `json:"user_id"`
	Service_id     int     `json:"service_id"`
	Receipt_id     int     `json:"receipt_id"`
	Rec_code       string  `json:"rec_code"`
	Ser_lock_drug  int     `json:"ser_lock_drug"`
	Quec_code      string  `json:"quec_code"`
	Quec_name      string  `json:"quec_name"`
	Quec_qty       float64 `json:"quec_qty"`
	Quec_unit      string  `json:"quec_unit"`
	Quec_cost      float64 `json:"quec_cost"`
	Quec_price     float64 `json:"quec_price"`
	Quec_discount  float64 `json:"quec_discount"`
	Quec_total     float64 `json:"quec_total"`
	Quec_is_set    int     `json:"quec_is_set"`
	Quec_is_active int     `json:"quec_is_active"`
	Quec_modify    string  `json:"quec_modify"`
}

type ServiceQueueItemCourseSet struct {
	Id               int     `json:"id"`
	Course_code      string  `json:"course_code"`
	Course_name      string  `json:"course_name"`
	Course_amount    int     `json:"course_amount"`
	Course_unit      string  `json:"course_unit"`
	Course_opd       float64 `json:"course_opd"`
	Course_ipd       float64 `json:"course_ipd"`
	Course_lock_drug int     `json:"course_lock_drug"`
	Course_cost      float64 `json:"course_cost"`
}

type ServiceQueueCourseProductSet struct {
	Id         int     `json:"id"`
	Product_id int     `json:"product_id"`
	Pd_name    string  `json:"pd_name"`
	Cp_amount  float64 `json:"cp_amount"`
}

type ServiceQueueItemtList struct {
	Id                 int     `json:"id"`
	Queue_id           int     `json:"queue_id"`
	User_id            int     `json:"user_id"`
	Queue_type_id      int     `json:"queue_type_id"`
	Product_id         int     `json:"product_id"`
	Product_store_id   int     `json:"product_store_id"`
	Product_unit_id    int     `json:"product_unit_id"`
	Queue_checking_id  *int    `json:"queue_checking_id"`
	Queue_course_id    *int    `json:"queue_course_id"`
	Checking_id        *int    `json:"checking_id"`
	Course_id          *int    `json:"course_id"`
	Service_id         *int    `json:"service_id"`
	Service_product_id *int    `json:"service_product_id"`
	Rec_code           string  `json:"rec_code"`
	Queue_limit        string  `json:"queue_limit"`
	Queue_code         string  `json:"queue_code"`
	Queue_name         string  `json:"queue_name"`
	Queue_label        string  `json:"queue_label"`
	Queue_set_qty      float64 `json:"queue_set_qty"`
	Queue_limit_qty    float64 `json:"queue_limit_qty"`
	Queue_qty          float64 `json:"queue_qty"`
	Queue_unit         string  `json:"queue_unit"`
	Queue_cost         float64 `json:"queue_cost"`
	Queue_price        float64 `json:"queue_price"`
	Queue_discount     float64 `json:"queue_discount"`
	Queue_total        float64 `json:"queue_total"`
	Queue_is_set       int     `json:"queue_is_set"`
	Balance            float64 `json:"balance"`
	Cost_total         float64 `json:"cost_total"`
	Queue_tap          int     `json:"queue_tap"`
}

type ServiceQueueItemUpdateCancel struct {
	Id        int `json:"id"`
	Queue_tap int `json:"queue_tap"`
}

type ServiceQueueCourseUpdate struct {
	Id          *int    `json:"id"`
	Quec_qty    float64 `json:"quec_qty"`
	Quec_modify string  `json:"quec_modify"`
}

type ServiceQueueProductUpdate struct {
	Id              *int    `json:"id"`
	Product_unit_id *int    `json:"product_unit_id"`
	Quep_qty        float64 `json:"quep_qty"`
	Quep_unit       string  `json:"quep_unit"`
	Quep_price      float64 `json:"quep_price"`
	Quep_discount   float64 `json:"quep_discount"`
	Quep_total      float64 `json:"quep_total"`
	Quep_modify     string  `json:"quep_modify"`
}

type ServiceQueueProductStoresCost struct {
	Id       int     `json:"id"`
	Pds_cost float64 `json:"pds_cost"`
}

type ServiceQueueItemProductSet struct {
	Product_id     int `json:"product_id"`
	Product_amount int `json:"product_amount"`
	// Product_code   string  `json:"product_code"`
	// Product_name   string  `json:"product_name"`
	// Product_unit   string  `json:"product_unit"`
	// Product_opd    float64 `json:"product_opd"`
	// Product_ipd    float64 `json:"product_ipd"`
}

type ServiceQueueTopicalNoti struct {
	Id             int     `json:"id"`
	Product_id     int     `json:"product_id"`
	Tpd_amount     float64 `json:"tpd_amount"`
	Topical_name   string  `json:"topical_name"`
	Topical_detail string  `json:"topical_detail"`
}

type ServiceUsed struct {
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

type ServiceProductUsed struct {
	Id                 int     `json:"id"`
	Shop_id            int     `json:"shop_id"`
	Service_id         int     `json:"service_id" `
	Service_used_id    int     `json:"service_used_id"`
	Course_id          int     `json:"course_id"`
	Queue_id           int     `json:"queue_id"`
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

type ServiceProductUsedNotSevice struct {
	Id               int     `json:"id"`
	Shop_id          int     `json:"shop_id"`
	Queue_id         int     `json:"queue_id"`
	Customer_id      int     `json:"customer_id"`
	Product_id       int     `json:"product_id"`
	Product_store_id int     `json:"product_store_id"`
	Product_unit_id  int     `json:"product_unit_id"`
	Serpu_code       string  `json:"serpu_code"`
	Serpu_name       string  `json:"serpu_name"`
	Serpu_qty        float64 `json:"serpu_qty"`
	Serpu_unit       string  `json:"serpu_unit"`
	Serpu_is_active  int     `json:"serpu_is_active"`
	Serpu_datetime   string  `json:"serpu_datetime"`
	Serpu_create     string  `json:"serpu_create"`
	Serpu_modify     string  `json:"serpu_modify"`
}

type ServiceUpdate struct {
	Ser_is_active int    `json:"ser_is_active"`
	Ser_update    string `json:"ser_update"`
}

type ServiceQueueItemUpdate struct {
	Id        *int    `json:"id"`
	Queue_tap int     `json:"queue_tap"`
	Qty       float64 `json:"qty"`
}

type ObjPayloadServiceProcess struct {
	Id int `json:"id" binding:"required"`
}

type ServiceQueue struct {
	Id            int    `json:"id"`
	Customer_id   int    `json:"customer_id"`
	Que_code      string `json:"que_code"`
	Que_status_id int    `json:"que_status_id"`
	Que_datetime  string `json:"que_datetime"`
	Que_update    string `json:"que_update"`
}

type ProcessCourseServiceQueue struct {
	Id               int `json:"id"`
	Course_lock_drug int `json:"course_lock_drug"`
	Course_amount    int `json:"course_amount"`
	Course_use_date  int `json:"course_use_date"`
	Course_exp_date  int `json:"course_exp_date"`
}

type ProcessServiceQueueUpdate struct {
	Ser_use       int     `json:"ser_use"`
	Ser_exp_date  *string `json:"ser_exp_date"`
	Ser_is_active int     `json:"ser_is_active"`
	Ser_update    string  `json:"ser_update"`
}

type ProcessServiceProductQueueUpdate struct {
	Serp_use       float64 `json:"serp_use"`
	Serp_balance   float64 `json:"serp_balance"`
	Serp_is_active int     `json:"serp_is_active"`
	Serp_modify    string  `json:"serp_modify"`
}

type QueueServiceUpdate struct {
	Que_status_id    int    `json:"que_status_id"`
	Que_datetime_out string `json:"que_datetime_out"`
	Que_time_end     int    `json:"que_time_end"`
	Que_update       string `json:"que_update"`
}

type ObjPayloadServiceTranfer struct {
	Id            int                                `json:"id"`
	Course_id     int                                `json:"course_id"`
	Ser_lock_drug int                                `json:"ser_lock_drug"`
	Customer_id   int                                `json:"customer_id"`
	Qty           float64                            `json:"qty"`
	Products      *[]ObjPayloadServiceProductTranfer `json:"products" gorm:"-"`
}

type ObjPayloadServiceProductTranfer struct {
	Id               int     `json:"id"`
	Product_id       int     `json:"product_id"`
	Product_store_id int     `json:"product_store_id"`
	Product_unit_id  int     `json:"product_unit_id"`
	Qty              float64 `json:"qty"`
	Unit             string  `json:"unit"`
}

type ObjPayloadServiceAdjust struct {
	Id            int                               `json:"id"`
	Course_id     int                               `json:"course_id"`
	Ser_lock_drug int                               `json:"ser_lock_drug"`
	Qty           float64                           `json:"qty"`
	Qty_adjust    float64                           `json:"qty_adjust"`
	Day           int                               `json:"day"`
	Day_adjust    int                               `json:"day_adjust"`
	Products      *[]ObjPayloadServiceProductAdjust `json:"products" gorm:"-"`
}

type ObjPayloadServiceProductAdjust struct {
	Id               int     `json:"id"`
	Product_id       int     `json:"product_id"`
	Product_store_id int     `json:"product_store_id"`
	Product_unit_id  int     `json:"product_unit_id"`
	Code             string  `json:"code"`
	Name             string  `json:"name"`
	Qty              float64 `json:"qty"`
	Unit             string  `json:"unit"`
}

type ServiceUpdateTranfer struct {
	Ser_tranfer   int    `json:"ser_tranfer"`
	Ser_is_active int    `json:"ser_is_active"`
	Ser_update    string `json:"ser_update"`
}

type ServiceUpdateAdjust struct {
	Ser_qty       int     `json:"ser_qty"`
	Ser_exp_date  *string `json:"ser_exp_date"`
	Ser_is_active int     `json:"ser_is_active"`
	Ser_update    string  `json:"ser_update"`
}

type ServiceProductUpdateTranfer struct {
	Serp_qty       float64 `json:"serp_qty"`
	Serp_use       float64 `json:"serp_use"`
	Serp_tranfer   float64 `json:"serp_tranfer"`
	Serp_adjust    float64 `json:"serp_adjust"`
	Serp_balance   float64 `json:"serp_balance"`
	Serp_is_active int     `json:"serp_is_active"`
	Serp_modify    string  `json:"serp_modify"`
}

type ServiceProductUpdateAdjust struct {
	Product_id       int     `json:"product_id"`
	Product_store_id int     `json:"product_store_id"`
	Product_unit_id  int     `json:"product_unit_id"`
	Serp_code        string  `json:"serp_code"`
	Serp_name        string  `json:"serp_name"`
	Serp_qty         float64 `json:"serp_qty"`
	Serp_use         float64 `json:"serp_use"`
	Serp_tranfer     float64 `json:"serp_tranfer"`
	Serp_adjust      float64 `json:"serp_adjust"`
	Serp_balance     float64 `json:"serp_balance"`
	Serp_unit        string  `json:"serp_unit"`
	Serp_is_active   int     `json:"serp_is_active"`
	Serp_modify      string  `json:"serp_modify"`
}

type Service struct {
	Id                int     `json:"id"`
	Shop_id           int     `json:"shop_id"`
	Shop_mother_id    int     `json:"shop_mother_id"`
	Receipt_id        int     `json:"receipt_id" `
	Receipt_detail_id int     `json:"receipt_detail_id"`
	User_id           int     `json:"user_id"`
	Ser_customer_id   int     `json:"ser_customer_id"`
	Customer_id       int     `json:"customer_id"`
	Course_id         int     `json:"course_id"`
	Ser_tranfer_id    *int    `json:"ser_tranfer_id"`
	Ser_code          string  `json:"ser_code"`
	Ser_name          string  `json:"ser_name"`
	Ser_lock_drug     int     `json:"ser_lock_drug"`
	Ser_qty           int     `json:"ser_qty"`
	Ser_tranfer       int     `json:"ser_tranfer"`
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

type ServiceProduct struct {
	Id                int     `json:"id"`
	Shop_id           int     `json:"shop_id"`
	Service_id        int     `json:"service_id" `
	Course_id         int     `json:"course_id"`
	Receipt_id        *int    `json:"receipt_id"`
	Receipt_detail_id *int    `json:"receipt_detail_id"`
	Product_id        int     `json:"product_id"`
	Product_store_id  int     `json:"product_store_id"`
	Product_unit_id   int     `json:"product_unit_id"`
	Serp_code         string  `json:"serp_code"`
	Serp_name         string  `json:"serp_name"`
	Serp_qty          float64 `json:"serp_qty"`
	Serp_use          float64 `json:"serp_use"`
	Serp_tranfer      float64 `json:"serp_tranfer"`
	Serp_adjust       float64 `json:"serp_adjust"`
	Serp_balance      float64 `json:"serp_balance"`
	Serp_unit         string  `json:"serp_unit"`
	Serp_lock_drug    int     `json:"serp_lock_drug"`
	Serp_use_set_qty  float64 `json:"serp_use_set_qty"`
	Serp_is_active    int     `json:"serp_is_active"`
	Serp_datetime     string  `json:"serp_datetime"`
	Serp_create       string  `json:"serp_create"`
	Serp_modify       string  `json:"serp_modify"`
}

type ServiceList struct {
	Id                     int                   `json:"id"`
	Course_id              int                   `json:"course_id"`
	Ser_code               string                `json:"ser_code"`
	Ser_name               string                `json:"ser_name"`
	Ser_lock_drug          int                   `json:"ser_lock_drug"`
	Ser_qty                int                   `json:"ser_qty"`
	Ser_unit               string                `json:"ser_unit"`
	Ser_use_date           int                   `json:"ser_use_date"`
	Ser_exp                int                   `json:"ser_exp"`
	Ser_exp_day            int                   `json:"ser_exp_day"`
	Ser_exp_date           *string               `json:"ser_exp_date"`
	Ser_use                int                   `json:"ser_use"`
	Ser_tranfer            int                   `json:"ser_tranfer"`
	Ser_amount             int                   `json:"ser_amount"`
	Ser_price_total        float64               `json:"ser_price_total"`
	Ser_is_active          int                   `json:"ser_is_active"`
	Tranfer_is_qty_course  int                   `json:"tranfer_is_qty_course"`
	Tranfer_is_qty_product int                   `json:"tranfer_is_qty_product"`
	Adjust_is_qty_course   int                   `json:"adjust_is_qty_course"`
	Adjust_is_day_course   int                   `json:"adjust_is_day_course"`
	Adjust_is_qty_product  int                   `json:"adjust_is_qty_product"`
	Products               *[]ServiceProductList `json:"products" gorm:"-"`
}

type ServiceProductList struct {
	Id                int     `json:"id"`
	Shop_id           int     `json:"shop_id"`
	Service_id        int     `json:"service_id" `
	Course_id         int     `json:"course_id"`
	Receipt_id        *int    `json:"receipt_id"`
	Receipt_detail_id *int    `json:"receipt_detail_id"`
	Product_id        int     `json:"product_id"`
	Product_store_id  int     `json:"product_store_id"`
	Product_unit_id   int     `json:"product_unit_id"`
	Serp_code         string  `json:"serp_code"`
	Serp_name         string  `json:"serp_name"`
	Serp_qty          float64 `json:"serp_qty"`
	Serp_use          float64 `json:"serp_use"`
	Serp_tranfer      float64 `json:"serp_tranfer"`
	Serp_adjust       float64 `json:"serp_adjust"`
	Serp_balance      float64 `json:"serp_balance"`
	Serp_unit         string  `json:"serp_unit"`
	Serp_lock_drug    int     `json:"serp_lock_drug"`
	Serp_use_set_qty  float64 `json:"serp_use_set_qty"`
	Serp_is_active    int     `json:"serp_is_active"`
	Serp_datetime     string  `json:"serp_datetime"`
	Serp_create       string  `json:"serp_create"`
	Serp_modify       string  `json:"serp_modify"`
}

type ObjPayloadSearchService struct {
	Search *string `json:"search" binding:"required,omitempty"`
	// Customer_id *string `json:"customer_id" binding:"required,omitempty"`
	ActivePage int `json:"active_page" binding:"required"`
	PerPage    int `json:"per_page" binding:"required"`
}

type ResponsePaginationService struct {
	Result_data   []ServiceSearch `json:"result_data"`
	Count_of_page int             `json:"count_of_page"`
	Count_all     int             `json:"count_all"`
}

// service
type ServiceSearch struct {
	ID              int     `json:"id"`
	ReceiptId       int     `json:"receipt_id"`
	Queue_id        int     `json:"queue_id"`
	RecCode         string  `json:"rec_code"`
	ReceiptDetailId int     `json:"receipt_detail_id"`
	ShopId          int     `json:"shop_id"`
	CustomerShopId  int     `json:"customer_shop_id"`
	UserId          int     `json:"user_id"`
	SerCustomerId   int     `json:"ser_customer_id"`
	CustomerId      int     `json:"customer_id"`
	CourseId        int     `json:"course_id"`
	SerTranferId    int     `json:"ser_tranfer_id"`
	SerCode         string  `json:"ser_code"`
	SerName         string  `json:"ser_name"`
	SerLockDrug     int     `json:"ser_lock_drug"`
	SerQty          int     `json:"ser_qty"`
	SerUnit         string  `json:"ser_unit"`
	SerUseDate      int     `json:"ser_use_date"`
	SerExp          int     `json:"ser_exp"`
	SerExpDate      string  `json:"ser_exp_date"`
	SerUse          int     `json:"ser_use"`
	SerTranfer      int     `json:"ser_tranfer"`
	SerIsActive     int     `json:"ser_is_active"`
	SerDatetime     string  `json:"ser_datetime"`
	SerCreate       string  `json:"ser_create"`
	SerUpdate       string  `json:"ser_update"`
	CourseAmount    float64 `json:"course_amount"`
	CourseIpd       float64 `json:"course_ipd"`
	CourseOpd       float64 `json:"course_opd"`
	CourseCost      float64 `json:"course_cost"`
	SerAmount       int     `json:"ser_amount"`
}

type ServiceUsedList struct {
	Id             int     `json:"id"`
	Shop_id        int     `json:"shop_id"`
	Shop_name      string  `json:"shop_name"`
	Service_id     int     `json:"service_id"`
	Queue_id       int     `json:"queue_id"`
	Que_code       string  `json:"que_code"`
	Receipt_id     int     `json:"receipt_id"`
	Course_id      int     `json:"course_id"`
	Customer_id    int     `json:"customer_id"`
	User_id        int     `json:"user_id"`
	User_fullname  string  `json:"user_fullname"`
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

type LogServices struct {
	Username   string  `json:"username"`
	Log_type   string  `json:"log_type"`
	Log_text   string  `json:"log_text"`
	Is_id      int     `json:"is_id"`
	Is_to_id   *int    `json:"is_to_id"`
	Log_qty    float64 `json:"log_qty"`
	Log_create string  `json:"log_create"`
}

type LogServiceProducts struct {
	Username   string  `json:"username"`
	Log_type   string  `json:"log_type"`
	Log_text   string  `json:"log_text"`
	Is_id      int     `json:"is_id"`
	Is_to_id   *int    `json:"is_to_id"`
	Log_qty    float64 `json:"log_qty"`
	Log_create string  `json:"log_create"`
}

type UpdateServiceProductUsed struct {
	Serpu_is_active int    `json:"serpu_is_active"`
	Serpu_modify    string `json:"serpu_modify"`
}

type UpdateServiceUsed struct {
	Seru_is_active int    `json:"seru_is_active"`
	Seru_update    string `json:"seru_update"`
}

type PrintServiceCourse struct {
	Use_customer_fullname    string  `json:"use_customer_fullname"`
	Use_customer_fullname_en string  `json:"use_customer_fullname_en"`
	Ser_customer_fullname    string  `json:"ser_customer_fullname"`
	Ser_customer_fullname_en string  `json:"ser_customer_fullname_en"`
	Ser_code                 string  `json:"ser_code"`
	Quec_qty                 float64 `json:"quec_qty"`
	Quec_unit                string  `json:"quec_unit"`
	Ser_name                 string  `json:"ser_name"`
	Ser_use_date             string  `json:"ser_use_date"`
	Rec_code                 string  `json:"rec_code"`
	Quec_modify              string  `json:"quec_modify"`
}

type PrintServiceProduct struct {
	Ser_code  string `json:"ser_code"`
	Ser_name  string `json:"ser_name"`
	Quep_code string `json:"quep_code"`
	Quep_name string `json:"quep_name"`
	Quep_qty  string `json:"quep_qty"`
	Quep_unit string `json:"quep_unit"`
}

type PrintService struct {
	Que_code          string `json:"que_code"`
	Que_create        string `json:"que_create"`
	Ctm_fullname      string `json:"ctm_fullname"`
	Ctm_fullname_en   string `json:"ctm_fullname_en"`
	Que_user_fullname string `json:"que_user_fullname"`
	Que_note          string `json:"que_note"`
	ReceiptShop
	Courses  []PrintServiceCourse  `json:"courses" gorm:"-"`
	Products []PrintServiceProduct `json:"products" gorm:"-"`
}
