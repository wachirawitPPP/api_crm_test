package structs

type ObjPayloadDashboard struct {
	Datestart *string `json:"date_start"`
	Dateend   *string `json:"date_end"`
}

type ObjPayloadDashboards struct {
	Datestartnow    *string `json:"date_start_now"`
	Dateendnow      *string `json:"date_end_now"`
	Datestartbefore *string `json:"date_start_before"`
	Dateendbefore   *string `json:"date_end_before"`
	Search          *string `json:"search"`
}

type DashboardSalesSummaryTotal struct {
	Id               int     `json:"id"`
	User_fullname    string  `json:"user_fullname"`
	User_fullname_en string  `json:"user_fullname_en"`
	User_image       string  `json:"user_image"`
	Hand_price       float64 `json:"hand_price"`
	Role_name_en     string  `json:"role_name_en"`
	Role_name_th     string  `json:"role_name_th"`
	Role_name_lo     string  `json:"role_name_lo"`
	Role_name_vn     string  `json:"role_name_vn"`
	Role_name_kh     string  `json:"role_name_kh"`
}

type DashboardLastReceipt struct {
	Id               int     `json:"id"`
	Rec_fullname     string  `json:"rec_fullname"`
	Rec_code         string  `json:"rec_code"`
	Rec_total        float64 `json:"rec_total"`
	Rec_pay          float64 `json:"rec_pay"`
	Rec_payment_type int     `json:"rec_payment_type"`
	Rec_pay_datetime string  `json:"rec_pay_datetime"`
	Ctm_image        string  `json:"ctm_image"`
	Ctm_id           string  `json:"ctm_id"`
}

type DashboardPayment struct {
	Pay_total     float64 `json:"pay_total"`
	Balance_total float64 `json:"balance_total"`
}

type DashboardOrderTotal struct {
	Total        float64 `json:"total"`
	Total_before float64 `json:"total_before"`
}

type DashboardOrderList struct {
	Checking float64 `json:"checking"`
	Courses  float64 `json:"courses"`
	Products float64 `json:"products"`
	Coin     float64 `json:"coin"`
}

type DashboardFeeComTotal struct {
	Fee_total        float64 `json:"fee_total"`
	Fee_total_before float64 `json:"fee_total_before"`
	Com_total        float64 `json:"com_total"`
	Com_total_before float64 `json:"com_total_before"`
}

type DashboardCustomer struct {
	M_old     float64 `json:"m_old"`
	F_old     float64 `json:"f_old"`
	Other_old float64 `json:"other_old"`
	M_new     float64 `json:"m_new"`
	F_new     float64 `json:"f_new"`
	Other_new float64 `json:"other_new"`
}

type DashboardOrderYear struct {
	Years int                   `json:"years"`
	Data  []DashboardOrderMouth `json:"data"`
}

type DashboardOrderMouth struct {
	Months int     `json:"months"`
	Total  float64 `json:"total"`
}

type DashboardOrderMouthYear struct {
	Years  int     `json:"years"`
	Months int     `json:"months"`
	Total  float64 `json:"total"`
}

type ShopAddon struct {
	Id         int `json:"in"`
	Shop_id    int `json:"shop_id"`
	Package_id int `json:"package_id"`
}

type DashboardTotal struct {
	Id                   int     `json:"id"`
	Shop_id              int     `json:"shop_id"`
	User_id              int     `json:"user_id"`
	Dt_payment_balance   float64 `json:"dt_payment_balance"`
	Dt_payment_pay       float64 `json:"dt_payment_pay"`
	Dt_ordertotal_before float64 `json:"dt_ordertotal_before"`
	Dt_ordertotal_now    float64 `json:"dt_ordertotal_now"`
	Dt_profit_before     float64 `json:"dt_profit_before"`
	Dt_profit_now        float64 `json:"dt_profit_now"`
	Dt_fee_before        float64 `json:"dt_fee_before"`
	Dt_fee_now           float64 `json:"dt_fee_now"`
	Dt_com_before        float64 `json:"dt_com_before"`
	Dt_com_now           float64 `json:"dt_com_now"`
	Dt_order_checking    float64 `json:"dt_order_checking"`
	Dt_order_courses     float64 `json:"dt_order_courses"`
	Dt_order_coin        float64 `json:"dt_order_coin"`
	Dt_order_products    float64 `json:"dt_order_products"`
	Dt_f_old             float64 `json:"dt_f_old"`
	Dt_f_new             float64 `json:"dt_f_new"`
	Dt_m_old             float64 `json:"dt_m_old"`
	Dt_m_new             float64 `json:"dt_m_new"`
	Dt_other_old         float64 `json:"dt_other_old"`
	Dt_other_new         float64 `json:"dt_other_new"`
	Dt_create            string  `json:"dt_create"`
}

type DashboardOrdertotalmouth struct {
	Id        int     `json:"id"`
	Shop_id   int     `json:"shop_id"`
	Dom_year  int     `json:"dom_year"`
	Dom_mouth int     `json:"dom_mouth"`
	Dom_total float64 `json:"dom_total"`
}
