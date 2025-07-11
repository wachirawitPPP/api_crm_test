package structs

type ShopTax struct {
	Id            int `json:"id"`
	Shop_id       int `json:"shop_id"`
	Tax_id        int `json:"tax_id"`
	Tax_sum_id    int `json:"tax_sum_id"`
	Tax_split_id  int `json:"tax_split_id"`
	Tax_rate      int `json:"tax_rate"`
	Tax_with_rate int `json:"tax_with_rate"`
}

type Tax struct {
	Id          int    `json:"id"`
	Shop_id     int    `json:"shop_id"`
	Tax_name    string `json:"tax_name"`
	Tax_name_en string `json:"tax_name_en"`
	Tax_type_id int    `json:"tax_type_id"`
	Tax_rate    int    `json:"tax_rate"`
}
