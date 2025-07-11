package structs

type ObjPayloadQueueProduct struct {
	Queue_id              int     `json:"queue_id" binding:"required"`
	Que_admis_id          int     `json:"que_admis_id" binding:"required"` // 1 = IPD, 2 = OPD
	Product_id            int     `json:"product_id" binding:"required"`
	Product_store_id      int     `json:"product_store_id"`
	Product_unit_id       int     `json:"product_unit_id"`
	Quep_qty              float64 `json:"quep_qty"`
	Quep_unit             string  `json:"quep_unit"`
	Quep_price            float64 `json:"quep_price"`
	Quep_discount_type_id int     `json:"quep_discount_type_id"`
	Quep_discount_item    float64 `json:"quep_discount_item"`
	Quep_discount         float64 `json:"quep_discount"`
	DpmId                 int     `json:"dpm_id"`
}

type ObjPayloadQueueCourse struct {
	Queue_id              int     `json:"queue_id" binding:"required"`
	Que_admis_id          int     `json:"que_admis_id" binding:"required"`
	Course_id             int     `json:"course_id" binding:"required"`
	Quec_qty              float64 `json:"quec_qty"`
	Quec_unit             string  `json:"quec_unit"`
	Quec_price            float64 `json:"quec_price"`
	Quec_discount_type_id int     `json:"queci_discount_type_id"`
	Quec_discount_item    float64 `json:"queci_discount_item"`
	Quec_discount         float64 `json:"quec_discount"`
	DpmId                 int     `json:"dpm_id"`
}

type ObjPayloadQueueChecking struct {
	Queue_id               int     `json:"queue_id" binding:"required"`
	Que_admis_id           int     `json:"que_admis_id" binding:"required"`
	Checking_id            int     `json:"checking_id" binding:"required"`
	Queci_qty              float64 `json:"queci_qty"`
	Queci_unit             string  `json:"queci_unit"`
	Queci_price            float64 `json:"queci_price"`
	Queci_discount_type_id int     `json:"queci_discount_type_id"`
	Queci_discount_item    float64 `json:"queci_discount_item"`
	Queci_discount         float64 `json:"queci_discount"`
	DpmId                  int     `json:"dpm_id"`
}

type ObjPayloadAddQueueItem struct {
	User_id   int                       `json:"user_id"`
	Shop_id   int                       `json:"shop_id"`
	Products  []ObjPayloadQueueProduct  `json:"products"`
	Courses   []ObjPayloadQueueCourse   `json:"courses"`
	Checkings []ObjPayloadQueueChecking `json:"checkings"`
}

type ObjPayloadAddItemSet struct {
	User_id  int              `json:"user_id"`
	Shop_id  int              `json:"shop_id"`
	Admis_id int              `json:"admis_id"`
	Queue_id int              `json:"queue_id"`
	DpmId    int              `json:"dpm_id"`
	ItemSet  []QueueItemtList `json:"item_set"`
}

type QueueItemProduct struct {
	Id             int    `json:"id"`
	Pd_type_id     int    `json:"pd_type_id"`
	Pd_code        string `json:"pd_code"`
	Pd_name        string `json:"pd_name"`
	Topical_id     *int   `json:"topical_id"`
	Drug_direction string `json:"drug_direction"`
}

type QueueItemTopical struct {
	Id             int    `json:"id"`
	Topical_name   string `json:"topical_name"`
	Topical_detail string `json:"topical_detail"`
}

type QueueProduct struct {
	Id                    *int    `json:"id"`
	Queue_id              int     `json:"queue_id"`
	User_id               int     `json:"user_id"`
	Product_id            int     `json:"product_id"`
	Product_store_id      int     `json:"product_store_id"`
	Product_unit_id       int     `json:"product_unit_id"`
	Quep_type_id          int     `json:"quep_type_id"`
	Queue_checking_id     *int    `json:"queue_checking_id"`
	Checking_id           *int    `json:"checking_id"`
	Queue_course_id       *int    `json:"queue_course_id"`
	Course_id             *int    `json:"course_id"`
	Quep_code             string  `json:"quep_code"`
	Quep_name             string  `json:"quep_name"`
	Quep_qty              float64 `json:"quep_qty"`
	Quep_set_qty          float64 `json:"quep_set_qty"`
	Quep_limit_qty        float64 `json:"quep_limit_qty"`
	Quep_unit             string  `json:"quep_unit"`
	Quep_cost             float64 `json:"quep_cost"`
	Quep_price            float64 `json:"quep_price"`
	Quep_discount_type_id int     `json:"quep_discount_type_id"`
	Quep_discount_item    float64 `json:"quep_discount_item"`
	Quep_discount         float64 `json:"quep_discount"`
	Topical_id            *int    `json:"topical_id"`
	Quep_topical          string  `json:"quep_topical"`
	Quep_direction        string  `json:"quep_direction"`
	Quep_total            float64 `json:"quep_total"`
	Quep_is_set           int     `json:"quep_is_set"`
	Quep_id_set           *int    `json:"quep_id_set"`
	Quep_id_ref           int     `json:"quep_id_ref"`
	Quep_is_active        int     `json:"quep_is_active"`
	Quep_modify           string  `json:"quep_modify"`
	DpmId                 int     `json:"dpm_id"`
}

type QueueProduct2 struct {
	Id                    *int    `json:"id"`
	Queue_id              int     `json:"queue_id"`
	User_id               int     `json:"user_id"`
	Product_id            int     `json:"product_id"`
	Product_store_id      int     `json:"product_store_id"`
	Product_unit_id       int     `json:"product_unit_id"`
	Quep_type_id          int     `json:"quep_type_id"`
	Queue_checking_id     *int    `json:"queue_checking_id"`
	Checking_id           *int    `json:"checking_id"`
	Queue_course_id       *int    `json:"queue_course_id"`
	Course_id             *int    `json:"course_id"`
	Quep_code             string  `json:"quep_code"`
	Quep_name             string  `json:"quep_name"`
	Quep_qty              float64 `json:"quep_qty"`
	Quep_set_qty          float64 `json:"quep_set_qty"`
	Quep_limit_qty        float64 `json:"quep_limit_qty"`
	Quep_unit             string  `json:"quep_unit"`
	Quep_cost             float64 `json:"quep_cost"`
	Quep_price            float64 `json:"quep_price"`
	Quep_discount_type_id int     `json:"quep_discount_type_id"`
	Quep_discount_item    float64 `json:"quep_discount_item"`
	Quep_discount         float64 `json:"quep_discount"`
	Topical_id            *int    `json:"topical_id"`
	Quep_topical          string  `json:"quep_topical"`
	Quep_direction        string  `json:"quep_direction"`
	Quep_total            float64 `json:"quep_total"`
	Quep_is_set           int     `json:"quep_is_set"`
	Quep_id_set           *int    `json:"quep_id_set"`
	Quep_is_active        int     `json:"quep_is_active"`
	Quep_modify           string  `json:"quep_modify"`
	DpmId                 int     `json:"dpm_id"`
	Dpm_name              string  `json:"dpm_name"`
	Dpm_name_en           string  `json:"dpm_name_en"`
	Inv_id                int     `json:"inv_id"`
	Inv_code              string  `json:"inv_code"`
	Inv_datetime          string  `json:"inv_datetime"`
}

type QueueItemProductSet struct {
	Product_id     int `json:"product_id"`
	Product_amount int `json:"product_amount"`
	// Product_code   string  `json:"product_code"`
	// Product_name   string  `json:"product_name"`
	// Product_unit   string  `json:"product_unit"`
	// Product_opd    float64 `json:"product_opd"`
	// Product_ipd    float64 `json:"product_ipd"`
}

type QueueProductSet struct {
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

type QueueItemCourse struct {
	Id               *int    `json:"id"`
	Course_type_id   int     `json:"course_type_id"`
	Course_code      string  `json:"course_code"`
	Course_name      string  `json:"course_name"`
	Course_lock_drug int     `json:"course_lock_drug"`
	Course_cost      float64 `json:"course_cost"`
	Course_opd       float64 `json:"course_opd"`
	Course_ipd       float64 `json:"course_ipd"`
	Course_unit      string  `json:"course_unit"`
	Course_qtyset    int     `json:"course_qtyset"`
}

type QueueCourse struct {
	Id                    *int    `json:"id"`
	Queue_id              int     `json:"queue_id"`
	Course_id             int     `json:"course_id"`
	User_id               int     `json:"user_id"`
	Quec_code             string  `json:"quec_code"`
	Quec_name             string  `json:"quec_name"`
	Quec_qty              float64 `json:"quec_qty"`
	Quec_unit             string  `json:"quec_unit"`
	Quec_cost             float64 `json:"quec_cost"`
	Quec_price            float64 `json:"quec_price"`
	Quec_discount_type_id int     `json:"queci_discount_type_id"`
	Quec_discount_item    float64 `json:"queci_discount_item"`
	Quec_discount         float64 `json:"quec_discount"`
	Quec_total            float64 `json:"quec_total"`
	Quec_is_set           int     `json:"quec_is_set"`
	Quec_id_ref           int     `json:"quec_id_ref"`
	Quec_id_set           *int    `json:"quec_id_set"`
	Quec_is_active        int     `json:"quec_is_active"`
	Quec_modify           string  `json:"quec_modify"`
	DpmId                 int     `json:"dpm_id"`
}

type QueueCourse2 struct {
	Id                    *int    `json:"id"`
	Queue_id              int     `json:"queue_id"`
	Course_id             int     `json:"course_id"`
	User_id               int     `json:"user_id"`
	Quec_code             string  `json:"quec_code"`
	Quec_name             string  `json:"quec_name"`
	Quec_qty              float64 `json:"quec_qty"`
	Quec_unit             string  `json:"quec_unit"`
	Quec_cost             float64 `json:"quec_cost"`
	Quec_price            float64 `json:"quec_price"`
	Quec_discount_type_id int     `json:"queci_discount_type_id"`
	Quec_discount_item    float64 `json:"queci_discount_item"`
	Quec_discount         float64 `json:"quec_discount"`
	Quec_total            float64 `json:"quec_total"`
	Quec_is_set           int     `json:"quec_is_set"`
	Quec_id_set           *int    `json:"quec_id_set"`
	Quec_is_active        int     `json:"quec_is_active"`
	Quec_modify           string  `json:"quec_modify"`
	DpmId                 int     `json:"dpm_id"`
	Dpm_name              string  `json:"dpm_name"`
	Dpm_name_en           string  `json:"dpm_name_en"`
	Inv_id                int     `json:"inv_id"`
	Inv_code              string  `json:"inv_code"`
	Inv_datetime          string  `json:"inv_datetime"`
}

type QueueItemCourseSet struct {
	Id                 int     `json:"id"`
	Course_code        string  `json:"course_code"`
	Course_name        string  `json:"course_name"`
	Course_amount      int     `json:"course_amount"`
	Course_unit        string  `json:"course_unit"`
	Course_opd         float64 `json:"course_opd"`
	Course_ipd         float64 `json:"course_ipd"`
	Course_list_qtyset int     `json:"course_list_qtyset"`
	Course_list_opd    float64 `json:"course_list_opd"`
	Course_list_ipd    float64 `json:"course_list_ipd"`
	Course_lock_drug   int     `json:"course_lock_drug"`
	Course_cost        float64 `json:"course_cost"`
}

type QueueCourseProductSet struct {
	Id         int     `json:"id"`
	Product_id int     `json:"product_id"`
	Cp_amount  float64 `json:"cp_amount"`
}

type QueueCourseProductSetByCourseList struct {
	Id         int     `json:"id"`
	CourseId   int     `json:"course_id"`
	Product_id int     `json:"product_id"`
	Cp_amount  float64 `json:"cp_amount"`
	CourseCode string  `json:"course_code"`
	CourseName string  `json:"course_name"`
	PdCode     string  `json:"pd_code"`
	PdName     string  `json:"pd_name"`
}

type QueueCourseProductSetByCourse struct {
	CourseCode string `json:"course_code"`
	CourseName string `json:"course_name"`
	PdCode     string `json:"pd_code"`
	PdName     string `json:"pd_name"`
}
type QueueCourseProduct struct {
	CourseCode string   `json:"course_code"`
	PdCode     []string `json:"pd_code"`
}

type QueueItemChecking struct {
	Id               *int    `json:"id"`
	Checking_type_id int     `json:"checking_type_id"`
	Checking_code    string  `json:"checking_code"`
	Checking_name    string  `json:"checking_name"`
	Checking_cost    float64 `json:"checking_cost"`
	Checking_opd     float64 `json:"checking_opd"`
	Checking_ipd     float64 `json:"checking_ipd"`
	Checking_unit    string  `json:"checking_unit"`
}

type QueueChecking struct {
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
	DpmId                  int     `json:"dpm_id"`
}

type QueueChecking2 struct {
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
	Queci_id_set           *int    `json:"queci_id_set"`
	Queci_is_active        int     `json:"queci_is_active"`
	Queci_modify           string  `json:"queci_modify"`
	DpmId                  int     `json:"dpm_id"`
	Dpm_name               string  `json:"dpm_name"`
	Dpm_name_en            string  `json:"dpm_name_en"`
	Inv_id                 int     `json:"inv_id"`
	Inv_code               string  `json:"inv_code"`
	Inv_datetime           string  `json:"inv_datetime"`
}

type QueueInvoiceList struct {
	Inv_id       int    `json:"inv_id"`
	Queue_id     int    `json:"queue_id"`
	Inv_code     string `json:"inv_code"`
	Inv_datetime string `json:"inv_datetime"`
	Dpm_id       int    `json:"dpm_id"`
	Dpm_name     string `json:"dpm_name"`
	Dpm_name_en  string `json:"dpm_name_en"`
}

type QueueInvoiceListResponse struct {
	Inv_id       int              `json:"inv_id"`
	Queue_id     int              `json:"queue_id"`
	Inv_code     string           `json:"inv_code"`
	Inv_datetime string           `json:"inv_datetime"`
	Products     []QueueProduct2  `json:"products"`
	Course       []QueueCourse2   `json:"courses"`
	Checkings    []QueueChecking2 `json:"checkings"`
}

type QueueItemCheckingSet struct {
	Id                int     `json:"id"`
	Checking_code     string  `json:"checking_code"`
	Checking_name     string  `json:"checking_name"`
	Checking_amount   int     `json:"checking_amount"`
	Checking_unit     string  `json:"checking_unit"`
	Checking_opd      float64 `json:"checking_opd"`
	Checking_ipd      float64 `json:"checking_ipd"`
	Checking_cost     float64 `json:"checking_cost"`
	Checking_list_opd float64 `json:"checking_list_opd"`
	Checking_list_ipd float64 `json:"checking_list_ipd"`
}

type QueueCheckingProductSet struct {
	Id         int     `json:"id"`
	Product_id int     `json:"product_id"`
	Cip_amount float64 `json:"cip_amount"`
}

// type QueueItemtCourseList struct {
// 	Id              int     `json:"id"`
// 	Queue_id        int     `json:"queue_id"`
// 	User_id         int     `json:"user_id"`
// 	Queue_type_id   int     `json:"queue_type_id"`
// 	Course_id       int     `json:"course_id"`
// 	Queue_code      string  `json:"queue_code"`
// 	Queue_name      string  `json:"queue_name"`
// 	Queue_label     string  `json:"queue_label"`
// 	Queue_set_qty   float64 `json:"queue_set_qty"`
// 	Queue_limit_qty float64 `json:"queue_limit_qty"`
// 	Queue_qty       float64 `json:"queue_qty"`
// 	Queue_unit      string  `json:"queue_unit"`
// 	Queue_cost      float64 `json:"queue_cost"`
// 	Queue_price     float64 `json:"queue_price"`
// 	Queue_discount  float64 `json:"queue_discount"`
// 	Queue_total     float64 `json:"queue_total"`
// 	Queue_is_set    int     `json:"queue_is_set"`
// 	Queue_tap       int     `json:"queue_tap"`
// }

type QueueItemtList struct {
	Id                     int                `json:"id"`
	Queue_id               int                `json:"queue_id"`
	User_id                int                `json:"user_id"`
	Queue_type_id          int                `json:"queue_type_id"`
	Product_id             int                `json:"product_id"`
	Product_store_id       int                `json:"product_store_id"`
	Product_unit_id        int                `json:"product_unit_id"`
	Queue_checking_id      *int               `json:"queue_checking_id"`
	Queue_course_id        *int               `json:"queue_course_id"`
	Checking_id            *int               `json:"checking_id"`
	Course_id              *int               `json:"course_id"`
	Queue_code             string             `json:"queue_code"`
	Queue_name             string             `json:"queue_name"`
	Queue_label            string             `json:"queue_label"`
	Queue_set_qty          float64            `json:"queue_set_qty"`
	Queue_limit_qty        float64            `json:"queue_limit_qty"`
	Queue_qty              float64            `json:"queue_qty"`
	Queue_unit             string             `json:"queue_unit"`
	Queue_cost             float64            `json:"queue_cost"`
	Queue_price            float64            `json:"queue_price"`
	Queue_discount         float64            `json:"queue_discount"`
	Queue_total            float64            `json:"queue_total"`
	Topical_id             *int               `json:"topical_id"`
	Queue_topical          string             `json:"queue_topical"`
	Queue_discount_type_id int                `json:"queue_discount_type_id"`
	Queue_discount_item    float64            `json:"queue_discount_item"`
	Queue_direction        string             `json:"queue_direction"`
	Queue_is_set           int                `json:"queue_is_set"`
	Queue_id_set           *int               `json:"queue_id_set"`
	Queue_tap              int                `json:"queue_tap"`
	Balance                float64            `json:"balance"`
	Cost_total             float64            `json:"cost_total"`
	Units                  *[]ProductUnitList `json:"units" gorm:"-"`
	DpmId                  int                `json:"dpm_id"`
	Dpm_name               string             `json:"dpm_name"`
	Dpm_name_en            string             `json:"dpm_name_en"`
	Inv_id                 int                `json:"inv_id"`
	Inv_code               string             `json:"inv_code"`
	Inv_datetime           string             `json:"inv_datetime"`
}

type QueueItemUpdateCancel struct {
	Id        int `json:"id"`
	Queue_tap int `json:"queue_tap"`
}

type QueueItemUpdate struct {
	Id               *int    `json:"id"`
	Product_unit_id  *int    `json:"product_unit_id"`
	Queue_tap        int     `json:"queue_tap"`
	Qty              float64 `json:"qty"`
	Unit             string  `json:"unit"`
	Price            float64 `json:"price"`
	Discount_type_id int     `json:"discount_type_id"`
	Discount_item    float64 `json:"discount_item"`
	Discount         float64 `json:"discount"`
}

type QueueCheckingUpdate struct {
	Id                     *int    `json:"id"`
	Queci_qty              float64 `json:"queci_qty"`
	Queci_unit             string  `json:"queci_unit"`
	Queci_price            float64 `json:"queci_price"`
	Queci_discount_type_id int     `json:"queci_discount_type_id"`
	Queci_discount_item    float64 `json:"queci_discount_item"`
	Queci_discount         float64 `json:"queci_discount"`
	Queci_total            float64 `json:"queci_total"`
	Queci_modify           string  `json:"queci_modify"`
}

type QueueCourseUpdate struct {
	Id                    *int    `json:"id"`
	Quec_qty              float64 `json:"quec_qty"`
	Quec_unit             string  `json:"quec_unit"`
	Quec_price            float64 `json:"quec_price"`
	Quec_discount_type_id int     `json:"quec_discount_type_id"`
	Quec_discount_item    float64 `json:"quec_discount_item"`
	Quec_discount         float64 `json:"quec_discount"`
	Quec_total            float64 `json:"quec_total"`
	Quec_modify           string  `json:"quec_modify"`
}

type QueueProductUpdate struct {
	Id                    *int    `json:"id"`
	Product_unit_id       int     `json:"product_unit_id"`
	Quep_qty              float64 `json:"quep_qty"`
	Quep_unit             string  `json:"quep_unit"`
	Quep_price            float64 `json:"quep_price"`
	Quep_discount_type_id int     `json:"quep_discount_type_id"`
	Quep_discount_item    float64 `json:"quep_discount_item"`
	Quep_discount         float64 `json:"quep_discount"`
	Quep_total            float64 `json:"quep_total"`
	Quep_modify           string  `json:"quep_modify"`
}

type QueueProductStoresCost struct {
	Id       int     `json:"id"`
	Pds_cost float64 `json:"pds_cost"`
}

type QueueDirection struct {
	Id             *int   `json:"id"`
	Topical_id     *int   `json:"topical_id"`
	Quep_topical   string `json:"quep_topical"`
	Quep_direction string `json:"quep_direction"`
	Quep_modify    string `json:"quep_modify"`
}

type ObjPayloadQueueDirection struct {
	Id             *int   `json:"id"`
	Topical_id     *int   `json:"topical_id"`
	Quep_topical   string `json:"quep_topical"`
	Quep_direction string `json:"quep_direction"`
}

type QueueTopicalNoti struct {
	Id             int     `json:"id"`
	Product_id     int     `json:"product_id"`
	Tpd_amount     float64 `json:"tpd_amount"`
	Topical_name   string  `json:"topical_name"`
	Topical_detail string  `json:"topical_detail"`
}

type ObjPayloadQueueMax struct {
	Id           int `json:"id"`
	Customer_id  int `json:"customer_id"`
	InvoiceLimit int `json:"invoice_limit"`
	DpmId        int `json:"dpm_id"`
}

type ObjPayloadQueueInvoices struct {
	Customer_id  int `json:"customer_id"`
	InvoiceLimit int `json:"invoice_limit"`
	Shop_id      int `json:"shop_id"`
}

type QueueMax struct {
	Id int `json:"id"`
}

type GetQueue struct {
	ID            int    `json:"id"`
	Shop_id       int    `json:"shop_id"`
	Customer_id   int    `json:"customer_id"`
	User_Id       int    `json:"user_id"`
	Que_admis_id  int    `json:"que_admis_id"` // 1 IPD 2 OPD
	Que_code      string `json:"que_code"`
	Que_ref_ipd   int    `json:"que_ref_ipd"`
	Que_status_id int    `json:"que_status_id"`
}

type CostTotal struct {
	Cost_total float64 `json:"cost_total"`
}

type QueueProductCheck struct {
	Id               *int    `json:"id"`
	Product_id       int     `json:"product_id"`
	Product_store_id int     `json:"product_store_id"`
	Quep_code        string  `json:"quep_code"`
	Quep_name        string  `json:"quep_name"`
	Quep_qty         float64 `json:"quep_qty"`
}
