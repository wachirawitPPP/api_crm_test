package structs

type ObjPayloadSearchCoupon struct {
	Search     *string `json:"search" binding:"required,omitempty"`
	Is_active  *string `json:"is_active" binding:"required,omitempty"`
	Shop_id    int     `json:"shop_id"`
	ActivePage int     `json:"active_page" binding:"required"`
	PerPage    int     `json:"per_page" binding:"required"`
}

type ResponsePaginationCoupon struct {
	Result_data   []Coupon `json:"result_data"`
	Count_of_page int      `json:"count_of_page"`
	Count_all     int      `json:"count_all"`
}

type Coupon struct {
	Id               int     `json:"id"`
	Shop_id          int     `json:"shop_id"`
	Coupon_code      string  `json:"coupon_code"`
	Coupon_name      string  `json:"coupon_name"`
	Coupon_limit     int     `json:"coupon_limit"`
	Coupon_use       int     `json:"coupon_use"`
	Coupon_amount    float64 `json:"coupon_amount"`
	Coupon_expdate   string  `json:"coupon_expdate"`
	Coupon_create    string  `json:"coupon_create"`
	Coupon_is_active int     `json:"coupon_is_active"`
}

type ObjAddCoupon struct {
	Id               int     `json:"id"`
	Shop_id          int     `json:"shop_id"`
	Coupon_code      string  `json:"coupon_code" binding:"required"`
	Coupon_name      string  `json:"coupon_name" binding:"required"`
	Coupon_limit     int     `json:"coupon_limit" binding:"required"`
	Coupon_use       int     `json:"coupon_use"`
	Coupon_amount    float64 `json:"coupon_amount" binding:"required"`
	Coupon_expdate   string  `json:"coupon_expdate" binding:"required"`
	Coupon_create    string  `json:"coupon_create"`
	Coupon_is_active int     `json:"coupon_is_active"`
}

type ObjUpdateCoupon struct {
	Id               int     `json:"id" binding:"required"`
	Shop_id          int     `json:"shop_id"`
	Coupon_code      string  `json:"coupon_code" binding:"omitempty"`
	Coupon_name      string  `json:"coupon_name" binding:"required"`
	Coupon_limit     int     `json:"coupon_limit" binding:"required"`
	Coupon_use       *int    `json:"coupon_use" binding:"omitempty"`
	Coupon_amount    float64 `json:"coupon_amount" binding:"required"`
	Coupon_expdate   string  `json:"coupon_expdate" binding:"required"`
	Coupon_is_active *int    `json:"coupon_is_active" binding:"required,omitempty"`
}

type LogCoupon struct {
	Username   string `json:"username"`
	Log_type   string `json:"log_type"`
	Log_text   string `json:"log_text"`
	Log_create string `json:"log_create"`
}

// use coupon
type ObjPayloadUseSearch struct {
	ShopId     int    `json:"shop_id" binding:"required"`
	SearchText string `json:"search_text" binding:"required"`
}

type ObjPayloadUseProcess struct {
	CouponId   int `json:"coupon_id" binding:"required"`
	CustomerId int `json:"customer_id" binding:"required"`
}

// use coupon history
type CouponHistory struct {
	Id          int     `json:"id"`
	Coupon_id   int     `json:"coupon_id"`
	Shop_id     int     `json:"shop_id"`
	User_id     int     `json:"user_id"`
	Customer_id *int    `json:"customer_id"`
	Cph_code    string  `json:"cph_code"`
	Cph_name    string  `json:"cph_name"`
	Cph_type_id int     `json:"cph_type_id"`
	Cph_limit   int     `json:"cph_limit"`
	Cph_use     int     `json:"cph_use"`
	Cph_total   int     `json:"cph_total"`
	Cph_amount  float64 `json:"cph_amount"`
	Cph_date    string  `json:"cph_date"`
	Cph_expdate string  `json:"cph_expdate"`
	Cph_create  string  `json:"cph_create"`
}
