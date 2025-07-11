package structs

type Category struct {
	Id               int    `json:"id"`
	Shop_id          int    `json:"shop_id"`
	Category_type_id int    `json:"category_type_id"`
	Category_name    string `json:"category_name"`
	CategoryType
}

type CategoryType struct {
	Id             int    `json:"id"`
	CategoryTypeTh string `json:"category_type_th"`
	CategoryTypeEn string `json:"category_type_en"`
}
