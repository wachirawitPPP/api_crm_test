package structs

type ObjUpdateReceive struct {
	Id                 int    `json:"id" binding:"required"`
	Shop_store_id      int    `json:"shop_store_id" binding:"required"`
	Password           string `json:"password" binding:"required"`
	Po_comment_receive string `json:"po_comment_receive"`
}

type PurchaseOrderReceive struct {
	Id                    int    `json:"id"`
	Po_code               string `json:"po_code"`
	Po_status_id          int    `json:"po_status_id"`
	User_id_receive       int    `json:"user_id_receive"`
	User_fullname_receive string `json:"user_fullname_receive"`
	Po_comment_receive    string `json:"po_comment_receive"`
	Po_date_import        string `json:"po_date_import"`
	Shop_store_id         int    `json:"shop_store_id"`
	Po_update             string `json:"po_update"`
}

type PurchaseOrderDetailReceive struct {
	Id                int     `json:"id"`
	Purchase_order_id int     `json:"purchase_order_id"`
	Product_unit_id   int     `json:"product_unit_id"`
	Pd_id             int     `json:"pd_id"`
	Pd_expire         string  `json:"pd_expire"`
	Pod_qty           float64 `json:"pod_qty"`
	Pod_cost          float64 `json:"pod_cost"`
	Pod_price         float64 `json:"pod_price"`
	Pod_discount      float64 `json:"pod_discount"`
	Pod_vat           float64 `json:"pod_vat"`
	Pod_total         float64 `json:"pod_total"`
	Pod_create        string  `json:"pod_create"`
	Pod_update        string  `json:"pod_update"`
	Pd_code           string  `json:"pd_code"`
	Pd_name           string  `json:"pd_name"`
	U_name            string  `json:"u_name"`
}

type PurchaseOrderDetailReceiveExpire struct {
	Pd_expire string `json:"pd_expire"`
}

type ProductStoresReceive struct {
	Id            int     `json:"id"`
	Shop_store_id int     `json:"shop_store_id"`
	Product_id    int     `json:"product_id"`
	Pds_barcode   string  `json:"pds_barcode"`
	Pds_cost      float64 `json:"pds_cost"`
	Pds_in        float64 `json:"pds_in"`
	Pds_out       float64 `json:"pds_out"`
	Pds_total     float64 `json:"pds_total"`
	Pds_date      string  `json:"pds_date"`
	Pds_comment   string  `json:"pds_comment"`
	Pds_is_active int     `json:"pds_is_active"`
	Pds_is_del    int     `json:"pds_is_del"`
	Pds_create    string  `json:"pds_create"`
	Pds_update    string  `json:"pds_update"`
}

type ProductStoresReceiveUpdate struct {
	Id         int     `json:"id"`
	Pds_cost   float64 `json:"pds_cost"`
	Pds_in     float64 `json:"pds_in"`
	Pds_out    float64 `json:"pds_out"`
	Pds_total  float64 `json:"pds_total"`
	Pds_update string  `json:"pds_update"`
}

type ProductStoresOrderReceive struct {
	Id                       int     `json:"id"`
	Product_store_id         int     `json:"product_store_id"`
	Purchase_order_id        int     `json:"purchase_order_id"`
	Purchase_order_detail_id int     `json:"purchase_order_detail_id"`
	Pdso_code                string  `json:"pdso_code"`
	Pdso_cost                float64 `json:"pdso_cost"`
	Pdso_date                string  `json:"pdso_date"`
	Pdso_expire              string  `json:"pdso_expire"`
	Pdso_in                  float64 `json:"pdso_in"`
	Pdso_out                 float64 `json:"pdso_out"`
	Pdso_use                 float64 `json:"pdso_use"`
	Pdso_move                float64 `json:"pdso_move"`
	Pdso_total               float64 `json:"pdso_total"`
	Pdso_is_active           int     `json:"pdso_is_active"`
	Pdso_update              string  `json:"pdso_update"`
	Pdso_create              string  `json:"pdso_create"`
}

type ProductStoresHistoryReceive struct {
	Id                       int     `json:"id"`
	Shop_id                  int     `json:"shop_id"`
	Shop_store_id            int     `json:"shop_store_id"`
	Product_store_id         int     `json:"product_store_id"`
	Pdsh_in                  float64 `json:"pdsh_in"`
	Product_store_order_id   int     `json:"product_store_order_id"`
	Purchase_order_detail_id int     `json:"purchase_order_detail_id"`
	Pdsh_out                 float64 `json:"pdsh_out"`
	Pdsh_inout               int     `json:"pdsh_inout"`
	Pdsh_out_id              int     `json:"pdsh_out_id"`
	Pdsh_order_forward       float64 `json:"pdsh_order_forward"`
	Pdsh_total_forward       float64 `json:"pdsh_total_forward"`
	Pdsh_amount              float64 `json:"pdsh_amount"`
	Pdsh_order_total         float64 `json:"pdsh_order_total"`
	Pdsh_total               float64 `json:"pdsh_total"`
	Pdsh_ref_doc_no          string  `json:"pdsh_ref_doc_no"`
	Pdsh_comment             string  `json:"pdsh_comment"`
	Pdsh_date                string  `json:"pdsh_date"`
	Pdsh_modify              string  `json:"pdsh_modify"`
	User_id                  int     `json:"user_id"`
	Product_id               int     `json:"product_id"`
	Pd_code                  string  `json:"pd_code"`
	Pd_name                  string  `json:"pd_name"`
	Pdsh_type_id             int     `json:"pdsh_type_id"`
}

type ShopStoreReceive struct {
	Id         int    `json:"id"`
	Shop_id    int    `json:"shop_id"`
	Ss_name    string `json:"ss_name"`
	Ss_is_over int    `json:"ss_is_over"`
}

type ProductIdReceive struct {
	Id      int    `json:"id"`
	Pd_code string `json:"pd_code"`
	Pd_name string `json:"pd_name"`
}

type LogPOReceive struct {
	Username   string `json:"username"`
	Log_type   string `json:"log_type"`
	Log_text   string `json:"log_text"`
	Log_create string `json:"log_create"`
}

// Transfers
type ObjCheckTran struct {
	Id int `json:"id" binding:"required"`
}

type ObjUpdateReceiveTran struct {
	Id                 int    `json:"id" binding:"required"`
	Password           string `json:"password" binding:"required"`
	Tf_comment_receive string `json:"tf_comment_receive"`
}

type TransferReceive struct {
	Id                    int    `json:"id"`
	Tf_code               string `json:"tf_code"`
	Shop_id               int    `json:"shop_id"`
	Shop_store_id         int    `json:"shop_store_id"`
	Shop_id_to            int    `json:"shop_id_to"`
	Shop_store_id_to      int    `json:"shop_store_id_to"`
	Tf_status_out_id      int    `json:"tf_status_out_id"`
	Tf_status_in_id       int    `json:"tf_status_in_id"`
	User_id               int    `json:"user_id"`
	User_id_receive       int    `json:"user_id_receive"`
	User_fullname_receive string `json:"user_fullname_receive"`
	Tf_comment_receive    string `json:"tf_comment_receive"`
	Tf_date_receive       string `json:"tf_date_receive"`
	Tf_update             string `json:"tf_update"`
}

type TransferReceiveUpdate struct {
	Id                    int    `json:"id"`
	Tf_status_out_id      int    `json:"tf_status_out_id"`
	Tf_status_in_id       int    `json:"tf_status_in_id"`
	User_id               int    `json:"user_id"`
	User_id_receive       int    `json:"user_id_receive"`
	User_fullname_receive string `json:"user_fullname_receive"`
	Tf_comment_receive    string `json:"tf_comment_receive"`
	Tf_date_receive       string `json:"tf_date_receive"`
	Tf_update             string `json:"tf_update"`
}

type TransferDetailReceive struct {
	Id                     int     `json:"id"`
	Transfer_id            int     `json:"transfer_id"`
	Product_store_id       int     `json:"product_store_id"`
	Product_unit_id        int     `json:"product_unit_id"`
	Product_store_order_id int     `json:"product_store_order_id"`
	Pd_id                  int     `json:"pd_id"`
	Pd_code                string  `json:"pd_code"`
	Pd_name                string  `json:"pd_name"`
	Pd_expire              string  `json:"pd_expire"`
	Tfd_qty                float64 `json:"tfd_qty"`
	Tfd_cost               float64 `json:"tfd_cost"`
	Tfd_total              float64 `json:"tfd_total"`
	Tfd_create             string  `json:"tfd_create"`
	Tfd_update             string  `json:"tfd_update"`
}

type ProductStoresOrderCheck struct {
	Id               int     `json:"id"`
	Product_store_id int     `json:"product_store_id"`
	Pdso_total       float64 `json:"pdso_total"`
}

type ProductStoresTotal struct {
	Pd_id     int     `json:"pd_id"`
	Pd_total  float64 `json:"pd_total"`
	Pd_cost   float64 `json:"pd_cost"`
	Pd_expire string  `json:"pd_expire"`
}

type ProductStoresOrderTransfer struct {
	Id          int     `json:"id"`
	Pdso_move   float64 `json:"pdso_move"`
	Pdso_total  float64 `json:"pdso_total"`
	Pdso_update string  `json:"pdso_update"`
}

type ProductStoresOrderReceiveTransfer struct {
	Id               int     `json:"id"`
	Product_store_id int     `json:"product_store_id"`
	Transfer_id      int     `json:"transfer_id"`
	Pdso_code        string  `json:"pdso_code"`
	Pdso_cost        float64 `json:"pdso_cost"`
	Pdso_date        string  `json:"pdso_date"`
	Pdso_expire      string  `json:"pdso_expire"`
	Pdso_in          float64 `json:"pdso_in"`
	Pdso_out         float64 `json:"pdso_out"`
	Pdso_use         float64 `json:"pdso_use"`
	Pdso_move        float64 `json:"pdso_move"`
	Pdso_total       float64 `json:"pdso_total"`
	Pdso_is_active   int     `json:"pdso_is_active"`
	Pdso_update      string  `json:"pdso_update"`
	Pdso_create      string  `json:"pdso_create"`
}

type ProductStoresHistoryTransfer struct {
	Id                     int     `json:"id"`
	Shop_id                int     `json:"shop_id"`
	Shop_store_id          int     `json:"shop_store_id"`
	Product_store_id       int     `json:"product_store_id"`
	Product_store_order_id int     `json:"product_store_order_id"`
	Pdsh_in                float64 `json:"pdsh_in"`
	Pdsh_out               float64 `json:"pdsh_out"`
	Transfer_id            int     `json:"transfer_id"`
	Transfer_detail_id     int     `json:"transfer_detail_id"`
	Pdsh_inout             int     `json:"pdsh_inout"`
	Pdsh_out_id            int     `json:"pdsh_out_id"`
	Pdsh_order_forward     float64 `json:"pdsh_order_forward"`
	Pdsh_total_forward     float64 `json:"pdsh_total_forward"`
	Pdsh_amount            float64 `json:"pdsh_amount"`
	Pdsh_order_total       float64 `json:"pdsh_order_total"`
	Pdsh_total             float64 `json:"pdsh_total"`
	Pdsh_ref_doc_no        string  `json:"pdsh_ref_doc_no"`
	Pdsh_comment           string  `json:"pdsh_comment"`
	Pdsh_date              string  `json:"pdsh_date"`
	Pdsh_modify            string  `json:"pdsh_modify"`
	User_id                int     `json:"user_id"`
	Product_id             int     `json:"product_id"`
	Pd_code                string  `json:"pd_code"`
	Pd_name                string  `json:"pd_name"`
	Pdsh_type_id           int     `json:"pdsh_type_id"`
}

type ProductStoresHistoryReceiveTransfer struct {
	Id                     int     `json:"id"`
	Shop_id                int     `json:"shop_id"`
	Shop_store_id          int     `json:"shop_store_id"`
	Product_store_id       int     `json:"product_store_id"`
	Product_store_order_id int     `json:"product_store_order_id"`
	Pdsh_in                float64 `json:"pdsh_in"`
	Pdsh_out               float64 `json:"pdsh_out"`
	Transfer_id            int     `json:"transfer_id"`
	Pdsh_inout             int     `json:"pdsh_inout"`
	Pdsh_out_id            int     `json:"pdsh_out_id"`
	Pdsh_order_forward     float64 `json:"pdsh_order_forward"`
	Pdsh_total_forward     float64 `json:"pdsh_total_forward"`
	Pdsh_amount            float64 `json:"pdsh_amount"`
	Pdsh_order_total       float64 `json:"pdsh_order_total"`
	Pdsh_total             float64 `json:"pdsh_total"`
	Pdsh_ref_doc_no        string  `json:"pdsh_ref_doc_no"`
	Pdsh_comment           string  `json:"pdsh_comment"`
	Pdsh_date              string  `json:"pdsh_date"`
	Pdsh_modify            string  `json:"pdsh_modify"`
	User_id                int     `json:"user_id"`
	Product_id             int     `json:"product_id"`
	Pd_code                string  `json:"pd_code"`
	Pd_name                string  `json:"pd_name"`
	Pdsh_type_id           int     `json:"pdsh_type_id"`
}

// Issues
type ObjUpdateStockIssue struct {
	Id       int    `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type IssueStock struct {
	Id                    int    `json:"id"`
	Shop_store_id         int    `json:"shop_store_id"`
	Isu_code              string `json:"isu_code"`
	Isu_cost_id           int    `json:"isu_cost_id"`
	Isu_status_id         int    `json:"isu_status_id"`
	User_id_confirm       int    `json:"user_id_confirm"`
	User_fullname_confirm string `json:"user_fullname_confirm"`
	Isu_date_confirm      string `json:"isu_date_confirm"`
	Isu_update            string `json:"isu_update"`
}

type IssueStockUpdate struct {
	Id                    int    `json:"id"`
	Isu_status_id         int    `json:"isu_status_id"`
	User_id_confirm       int    `json:"user_id_confirm"`
	User_fullname_confirm string `json:"user_fullname_confirm"`
	Isu_date_confirm      string `json:"isu_date_confirm"`
	Isu_update            string `json:"isu_update"`
}

type IssueDetailStock struct {
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
}

type ProductStoresIssues struct {
	Id            int     `json:"id"`
	Shop_store_id int     `json:"shop_store_id"`
	Product_id    int     `json:"product_id"`
	Pds_barcode   string  `json:"pds_barcode"`
	Pds_cost      float64 `json:"pds_cost"`
	Pds_in        float64 `json:"pds_in"`
	Pds_out       float64 `json:"pds_out"`
	Pds_total     float64 `json:"pds_total"`
	Pds_date      string  `json:"pds_date"`
	Pds_comment   string  `json:"pds_comment"`
	Pds_is_active int     `json:"pds_is_active"`
	Pds_is_del    int     `json:"pds_is_del"`
	Pds_create    string  `json:"pds_create"`
	Pds_update    string  `json:"pds_update"`
}

type ProductStoresIssuesUpdate struct {
	Id         int     `json:"id"`
	Pds_in     float64 `json:"pds_in"`
	Pds_out    float64 `json:"pds_out"`
	Pds_total  float64 `json:"pds_total"`
	Pds_update string  `json:"pds_update"`
}

type ProductStoresOrderIssues struct {
	Id          int     `json:"id"`
	Pdso_move   float64 `json:"pdso_move"`
	Pdso_total  float64 `json:"pdso_total"`
	Pdso_update string  `json:"pdso_update"`
}

type ProductStoresHistoryIssues struct {
	Id                     int     `json:"id"`
	Shop_id                int     `json:"shop_id"`
	Shop_store_id          int     `json:"shop_store_id"`
	Product_store_id       int     `json:"product_store_id"`
	Product_store_order_id int     `json:"product_store_order_id"`
	Pdsh_in                float64 `json:"pdsh_in"`
	Pdsh_out               float64 `json:"pdsh_out"`
	Issue_id               int     `json:"issue_id"`
	Issue_detail_id        int     `json:"issue_detail_id"`
	Pdsh_inout             int     `json:"pdsh_inout"`
	Pdsh_out_id            int     `json:"pdsh_out_id"`
	Pdsh_order_forward     float64 `json:"pdsh_order_forward"`
	Pdsh_total_forward     float64 `json:"pdsh_total_forward"`
	Pdsh_amount            float64 `json:"pdsh_amount"`
	Pdsh_order_total       float64 `json:"pdsh_order_total"`
	Pdsh_total             float64 `json:"pdsh_total"`
	Pdsh_ref_doc_no        string  `json:"pdsh_ref_doc_no"`
	Pdsh_comment           string  `json:"pdsh_comment"`
	Pdsh_date              string  `json:"pdsh_date"`
	Pdsh_modify            string  `json:"pdsh_modify"`
	User_id                int     `json:"user_id"`
	Product_id             int     `json:"product_id"`
	Pd_code                string  `json:"pd_code"`
	Pd_name                string  `json:"pd_name"`
	Pdsh_type_id           int     `json:"pdsh_type_id"`
}

type LogISS struct {
	Username   string `json:"username"`
	Log_type   string `json:"log_type"`
	Log_text   string `json:"log_text"`
	Log_create string `json:"log_create"`
}

// Adjusts
type ObjUpdateStockAdjust struct {
	Id       int    `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AdjustStock struct {
	Id            int    `json:"id"`
	Shop_store_id int    `json:"shop_store_id"`
	Adj_code      string `json:"adj_code"`
	Adj_status_id int    `json:"adj_status_id"`
	Adj_update    string `json:"adj_update"`
}

type AdjustDetailStock struct {
	Id                     int     `json:"id"`
	Adjust_id              int     `json:"adjust_id"`
	Product_store_id       int     `json:"product_store_id"`
	Product_unit_id        int     `json:"product_unit_id"`
	Product_store_order_id int     `json:"product_store_order_id"`
	Pd_id                  int     `json:"pd_id"`
	Pd_code                string  `json:"pd_code"`
	Pd_name                string  `json:"pd_name"`
	Pd_expire              string  `json:"pd_expire"`
	Adjd_qty               float64 `json:"adjd_qty"`
	Adjd_balance           float64 `json:"adjd_balance"`
	Adjd_amount            float64 `json:"adjd_amount"`
	Adjd_cost              float64 `json:"adjd_cost"`
	Adjd_total             float64 `json:"adjd_total"`
	Adjd_create            string  `json:"adjd_create"`
	Adjd_update            string  `json:"adjd_update"`
}

type AdjustDetailUpdate struct {
	Id           int     `json:"id"`
	Adjd_balance float64 `json:"adjd_balance"`
	Adjd_qty     float64 `json:"adjd_qty"`
	Adjd_in_out  int     `json:"adjd_in_out"`
	Adjd_total   float64 `json:"adjd_total"`
	Adjd_update  string  `json:"adjd_update"`
}

type ProductStoresAdjust struct {
	Id            int     `json:"id"`
	Shop_store_id int     `json:"shop_store_id"`
	Product_id    int     `json:"product_id"`
	Pds_barcode   string  `json:"pds_barcode"`
	Pds_cost      float64 `json:"pds_cost"`
	Pds_in        float64 `json:"pds_in"`
	Pds_out       float64 `json:"pds_out"`
	Pds_total     float64 `json:"pds_total"`
	Pds_date      string  `json:"pds_date"`
	Pds_comment   string  `json:"pds_comment"`
	Pds_is_active int     `json:"pds_is_active"`
	Pds_is_del    int     `json:"pds_is_del"`
	Pds_create    string  `json:"pds_create"`
	Pds_update    string  `json:"pds_update"`
}

type ProductStoresAdjustUpdate struct {
	Id         int     `json:"id"`
	Pds_in     float64 `json:"pds_in"`
	Pds_out    float64 `json:"pds_out"`
	Pds_total  float64 `json:"pds_total"`
	Pds_update string  `json:"pds_update"`
}

type ProductStoresOrderAdjust struct {
	Id          int     `json:"id"`
	Pdso_in     float64 `json:"pdso_in"`
	Pdso_move   float64 `json:"pdso_move"`
	Pdso_total  float64 `json:"pdso_total"`
	Pdso_update string  `json:"pdso_update"`
}

type ProductStoresHistoryAdjust struct {
	Id                     int     `json:"id"`
	Shop_id                int     `json:"shop_id"`
	Shop_store_id          int     `json:"shop_store_id"`
	Product_store_id       int     `json:"product_store_id"`
	Product_store_order_id int     `json:"product_store_order_id"`
	Pdsh_in                float64 `json:"pdsh_in"`
	Pdsh_out               float64 `json:"pdsh_out"`
	Adjust_id              int     `json:"adjust_id"`
	Adjust_detail_id       int     `json:"adjust_detail_id"`
	Pdsh_inout             int     `json:"pdsh_inout"`
	Pdsh_out_id            int     `json:"pdsh_out_id"`
	Pdsh_order_forward     float64 `json:"pdsh_order_forward"`
	Pdsh_total_forward     float64 `json:"pdsh_total_forward"`
	Pdsh_amount            float64 `json:"pdsh_amount"`
	Pdsh_order_total       float64 `json:"pdsh_order_total"`
	Pdsh_total             float64 `json:"pdsh_total"`
	Pdsh_ref_doc_no        string  `json:"pdsh_ref_doc_no"`
	Pdsh_comment           string  `json:"pdsh_comment"`
	Pdsh_date              string  `json:"pdsh_date"`
	Pdsh_modify            string  `json:"pdsh_modify"`
	User_id                int     `json:"user_id"`
	Product_id             int     `json:"product_id"`
	Pd_code                string  `json:"pd_code"`
	Pd_name                string  `json:"pd_name"`
	Pdsh_type_id           int     `json:"pdsh_type_id"`
}

type LogADJ struct {
	Username   string `json:"username"`
	Log_type   string `json:"log_type"`
	Log_text   string `json:"log_text"`
	Log_create string `json:"log_create"`
}
