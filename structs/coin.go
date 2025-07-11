package structs

type ObjPayloadSearchCoin struct {
	Search     *string `json:"search" binding:"required,omitempty"`
	Is_active  *string `json:"is_active" binding:"required,omitempty"`
	Shop_id    int     `json:"shop_id"`
	ActivePage int     `json:"active_page" binding:"required"`
	PerPage    int     `json:"per_page" binding:"required"`
}

type ResponsePaginationCoin struct {
	Result_data   []Coin `json:"result_data"`
	Count_of_page int    `json:"count_of_page"`
	Count_all     int    `json:"count_all"`
}

type Coin struct {
	Id                int     `json:"id"`
	Shop_id           int     `json:"shop_id"`
	Coin_code         string  `json:"coin_code"`
	Coin_name         string  `json:"coin_name"`
	Coin_price        float64 `json:"coin_price"`
	Coin_amount       float64 `json:"coin_amount"`
	Coin_is_active    int     `json:"coin_is_active"`
	Coin_create       string  `json:"coin_create"`
	Customer_group_id int     `json:"customer_group_id"`
	CoinCustomerGroups
}

type ObjAddCoin struct {
	Id                int     `json:"id"`
	Shop_id           int     `json:"shop_id"`
	Coin_code         *string `json:"coin_code" binding:"required,omitempty"`
	Coin_name         string  `json:"coin_name" binding:"required"`
	Coin_price        float64 `json:"coin_price" binding:"required,omitempty"`
	Coin_amount       float64 `json:"coin_amount" binding:"required,omitempty"`
	Coin_is_active    int     `json:"coin_is_active"`
	Coin_create       string  `json:"coin_create"`
	Customer_group_id int     `json:"customer_group_id" binding:"required"`
}

type ObjUpdateCoin struct {
	Id                int     `json:"id" binding:"required"`
	Shop_id           int     `json:"shop_id"`
	Coin_code         *string `json:"coin_code" binding:"required,omitempty"`
	Coin_name         string  `json:"coin_name" binding:"required"`
	Coin_price        float64 `json:"coin_price" binding:"required,omitempty"`
	Coin_amount       float64 `json:"coin_amount" binding:"required,omitempty"`
	Coin_is_active    *int    `json:"coin_is_active" binding:"required,omitempty"`
	Customer_group_id int     `json:"customer_group_id" binding:"required"`
}

type LogCoin struct {
	Username   string `json:"username"`
	Log_type   string `json:"log_type"`
	Log_text   string `json:"log_text"`
	Log_create string `json:"log_create"`
}

type CoinCustomerGroups struct {
	Id      int    `json:"id"`
	Shop_id int    `json:"shop_id"`
	Cg_name string `json:"cg_name"`
}
