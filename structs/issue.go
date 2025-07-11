package structs

type ObjPayloadSearchISS struct {
	Search        *string `json:"search" binding:"required,omitempty"`
	Isu_status_id *string `json:"isu_status_id" binding:"required,omitempty"`
	Shop_id       int     `json:"shop_id"`
	ActivePage    int     `json:"active_page" binding:"required"`
	PerPage       int     `json:"per_page" binding:"required"`
}

type ResponsePaginationISS struct {
	Result_data   []IssueList `json:"result_data"`
	Count_of_page int         `json:"count_of_page"`
	Count_all     int         `json:"count_all"`
}

type IssueList struct {
	Id                    int    `json:"id"`
	Shop_id               int    `json:"shop_id"`
	Shop_name             string `json:"shop_name"`
	Shop_store_id         int    `json:"shop_store_id"`
	Ss_name               string `json:"ss_name"`
	Isu_code              string `json:"isu_code"`
	Isu_cost_id           int    `json:"isu_cost_id"`
	Isu_status_id         int    `json:"isu_status_id"`
	Isu_date              string `json:"isu_date"`
	User_id               int    `json:"user_id"`
	User_fullname         string `json:"user_fullname"`
	User_fullname_en      string `json:"user_fullname_en"`
	Isu_comment           string `json:"isu_comment"`
	User_id_confirm       int    `json:"user_id_confirm"`
	User_fullname_confirm string `json:"user_fullname_confirm"`
	Isu_date_confirm      string `json:"isu_date_confirm"`
	User_id_cancel        int    `json:"user_id_cancel"`
	User_fullname_cancel  string `json:"user_fullname_cancel"`
	Isu_comment_cancel    string `json:"isu_comment_cancel"`
	Isu_create            string `json:"isu_create"`
	Isu_update            string `json:"isu_update"`
}

type Issue struct {
	Id                    int    `json:"id"`
	Shop_id               int    `json:"shop_id"`
	Shop_name             string `json:"shop_name"`
	Shop_store_id         int    `json:"shop_store_id"`
	Ss_name               string `json:"ss_name"`
	Isu_code              string `json:"isu_code"`
	Isu_cost_id           int    `json:"isu_cost_id"`
	Isu_status_id         int    `json:"isu_status_id"`
	Isu_date              string `json:"isu_date"`
	User_id               int    `json:"user_id"`
	User_fullname         string `json:"user_fullname"`
	Isu_comment           string `json:"isu_comment"`
	User_id_confirm       int    `json:"user_id_confirm"`
	User_fullname_confirm string `json:"user_fullname_confirm"`
	Isu_date_confirm      string `json:"isu_date_confirm"`
	User_id_cancel        int    `json:"user_id_cancel"`
	User_fullname_cancel  string `json:"user_fullname_cancel"`
	Isu_comment_cancel    string `json:"isu_comment_cancel"`
	Isu_create            string `json:"isu_create"`
	Isu_update            string `json:"isu_update"`
}

type IssueDetail struct {
	Id                    int          `json:"id"`
	Shop_id               int          `json:"shop_id"`
	Shop_name             string       `json:"shop_name"`
	Shop_code             string       `json:"shop_code"`
	Shop_tax              string       `json:"shop_tax"`
	Shop_phone            string       `json:"shop_phone"`
	Shop_email            string       `json:"shop_email"`
	Shop_address          string       `json:"shop_address"`
	Shop_district         string       `json:"shop_district"`
	Shop_amphoe           string       `json:"shop_amphoe"`
	Shop_province         string       `json:"shop_province"`
	Shop_zipcode          string       `json:"shop_zipcode"`
	Shop_store_id         int          `json:"shop_store_id"`
	Ss_name               string       `json:"ss_name"`
	Isu_code              string       `json:"isu_code"`
	Isu_cost_id           int          `json:"isu_cost_id"`
	Isu_status_id         int          `json:"isu_status_id"`
	Isu_date              string       `json:"isu_date"`
	User_id               int          `json:"user_id"`
	User_fullname         string       `json:"user_fullname"`
	User_fullname_en      string       `json:"user_fullname_en"`
	Isu_comment           string       `json:"isu_comment"`
	User_id_confirm       int          `json:"user_id_confirm"`
	User_fullname_confirm string       `json:"user_fullname_confirm"`
	Isu_date_confirm      string       `json:"isu_date_confirm"`
	User_id_cancel        int          `json:"user_id_cancel"`
	User_fullname_cancel  string       `json:"user_fullname_cancel"`
	Isu_comment_cancel    string       `json:"isu_comment_cancel"`
	Isu_create            string       `json:"isu_create"`
	Isu_update            string       `json:"isu_update"`
	Subs                  *[]IssueSubs `json:"subs" gorm:"-"`
	Shop                  ReceiptShop  `json:"shop" gorm:"-"`
}

type IssueSubs struct {
	Id                     int     `json:"id"`
	Issue_id               int     `json:"issue_id"`
	Product_store_id       int     `json:"product_store_id"`
	Product_unit_id        int     `json:"product_unit_id"`
	Product_store_order_id int     `json:"product_store_order_id"`
	Pd_id                  int     `json:"pd_id"`
	Pd_code                string  `json:"pd_code"`
	Pd_name                string  `json:"pd_name"`
	Pd_expire              string  `json:"pd_expire"`
	Isud_qty               float64 `json:"isud_qty"`
	Isud_cost              float64 `json:"isud_cost"`
	Isud_total             float64 `json:"isud_total"`
	Isud_create            string  `json:"isud_create"`
	Isud_update            string  `json:"isud_update"`
	U_name                 string  `json:"u_name"`
	Pu_rate                int     `json:"pu_rate"`
	Pod_cost               float64 `json:"pod_cost"`
}

type DocIssue struct {
	Shop_id              int    `json:"shop_id"`
	Issue_id_default     string `json:"issue_id_default"`
	Issue_number_default string `json:"issue_number_default"`
	Issue_number_digit   int    `json:"issue_number_digit"`
	Issue_type           int    `json:"issue_type"`
}

type LogIssue struct {
	Username   string `json:"username"`
	Log_type   string `json:"log_type"`
	Log_text   string `json:"log_text"`
	Log_create string `json:"log_create"`
}

type ObjPayloadIssueAdd struct {
	Id            int          `json:"id"`
	Shop_id       int          `json:"shop_id" binding:"required"`
	Shop_store_id int          `json:"shop_store_id" binding:"required"`
	Isu_code      string       `json:"isu_code" binding:"required"`
	Isu_cost_id   int          `json:"isu_cost_id" binding:"required"`
	Isu_status_id int          `json:"isu_status_id"`
	Isu_date      string       `json:"isu_date" binding:"required"`
	User_id       int          `json:"user_id"`
	User_fullname string       `json:"user_fullname"`
	Isu_comment   *string      `json:"isu_comment" binding:"required,omitempty"`
	Isu_create    string       `json:"isu_create"`
	Isu_update    string       `json:"isu_update"`
	Subs          *[]IssueSubs `json:"subs" gorm:"-" binding:"required,omitempty"`
}

type ObjPayloadIssueEdit struct {
	Id            int          `json:"id" binding:"required"`
	Shop_id       int          `json:"shop_id" binding:"required"`
	Shop_store_id int          `json:"shop_store_id" binding:"required"`
	Isu_code      string       `json:"isu_code" binding:"required"`
	Isu_cost_id   int          `json:"isu_cost_id" binding:"required"`
	Isu_status_id int          `json:"isu_status_id"`
	Isu_date      string       `json:"isu_date" binding:"required"`
	User_id       int          `json:"user_id" binding:"required"`
	User_fullname string       `json:"user_fullname" binding:"required"`
	Isu_comment   *string      `json:"isu_comment" binding:"required,omitempty"`
	Subs          *[]IssueSubs `json:"subs" gorm:"-" binding:"required,omitempty"`
}

type ProductStoreOrderIssue struct {
	Id         int     `json:"id"`
	Pdso_in    float64 `json:"pdso_in"`
	Pdso_out   float64 `json:"pdso_out"`
	Pdso_use   float64 `json:"pdso_use"`
	Pdso_move  float64 `json:"pdso_move"`
	Pdso_total float64 `json:"pdso_total"`
}

type CancelIssue struct {
	Id                   int     `json:"id" binding:"required"`
	User_id_cancel       int     `json:"user_id_cancel"`
	User_fullname_cancel string  `json:"user_fullname_cancel"`
	Isu_status_id        int     `json:"isu_status_id"`
	Isu_password_confirm string  `json:"isu_password_confirm" binding:"required"`
	Isu_comment_cancel   *string `json:"isu_comment_cancel" binding:"required,omitempty"`
	Isu_update           string  `json:"isu_update"`
}
