package structs

type ObjPayloaItem struct {
	Search     *string `json:"search" binding:"required,omitempty"`
	ActivePage int     `json:"active_page" binding:"required"`
	PerPage    int     `json:"per_page" binding:"required"`
}

type ItemProduct struct {
	Id           int              `json:"id"`
	Shop_id      int              `json:"shop_id"`
	Shop_name    string           `json:"shop_name"`
	Shop_name_en string           `json:"shop_name_en"`
	Shop_code    string           `json:"shop_code"`
	Product_id   int              `json:"product_id"`
	Pd_code      string           `json:"pd_code"`
	Pd_barcode   string           `json:"pd_barcode"`
	Pd_name      string           `json:"pd_name"`
	Pd_type_id   int              `json:"pd_type_id"`
	Pd_image_1   string           `json:"pd_image_1"`
	Pd_image_2   string           `json:"pd_image_2"`
	Pd_image_3   string           `json:"pd_image_3"`
	Pd_image_4   string           `json:"pd_image_4"`
	Pd_detail    string           `json:"pd_detail"`
	Subs         []ItemProductSub `json:"subs" gorm:"-"`
}

type ItemProductSub struct {
	Id               int               `json:"id"`
	Shop_id          int               `json:"shop_id"`
	Product_id       int               `json:"product_id"`
	Product_store_id int               `json:"product_store_id"`
	Product_units_id int               `json:"product_units_id"`
	Pd_code          string            `json:"pd_code"`
	Pd_name          string            `json:"pd_name"`
	Pd_image_1       string            `json:"pd_image_1"`
	Pd_image_2       string            `json:"pd_image_2"`
	Pd_image_3       string            `json:"pd_image_3"`
	Pd_image_4       string            `json:"pd_image_4"`
	Pd_detail        string            `json:"pd_detail"`
	U_name           string            `json:"u_name"`
	Pu_amount        float64           `json:"pu_amount"`
	Pu_rate          float64           `json:"pu_rate"`
	Balance          float64           `json:"balance"`
	Psp_price_ipd    float64           `json:"psp_price_ipd"`
	Psp_price_opd    float64           `json:"psp_price_opd"`
	Topical_id       int               `json:"topical_id"`
	Topical_detail   string            `json:"topical_detail"`
	Drug_direction   string            `json:"drug_direction"`
	Label            string            `json:"label"`
	Is_set           int               `json:"is_set"`
	Id_set           *int              `json:"id_set"`
	Units            []ItemProductUnit `json:"units" gorm:"-"`
}

type ItemProductUnit struct {
	Id               int     `json:"id"`
	Product_id       int     `json:"product_id"`
	Product_units_id int     `json:"product_units_id"`
	Pd_code          string  `json:"pd_code"`
	Pd_name          string  `json:"pd_name"`
	Pu_amount        float64 `json:"pu_amount"`
	Pu_rate          float64 `json:"pu_rate"`
	U_name           string  `json:"u_name"`
}

type ObjQueryProductStoreBalance struct {
	Pds_balance float64 `json:"pds_balance" gorm:"type:decimal(10,2)"`
}

type ItemTopical struct {
	Id             int    `json:"id"`
	Topical_name   string `json:"topical_name"`
	Topical_detail string `json:"topical_detail"`
}

type ItemProductSet struct {
	Id               int     `json:"id"`
	Product_amount   float64 `json:"product_amount"`
	Product_list_opd float64 `json:"product_list_opd"`
	Product_list_ipd float64 `json:"product_list_ipd"`
	Topical_id       int     `json:"topical_id"`
	Topical_detail   string  `json:"topical_detail"`
	Drug_direction   string  `json:"drug_direction"`
}

type ItemCourse struct {
	Id               int             `json:"id"`
	Shop_id          int             `json:"shop_id"`
	Shop_name        string          `json:"shop_name"`
	Shop_name_en     string          `json:"shop_name_en"`
	Shop_code        string          `json:"shop_code"`
	Course_type_id   int             `json:"course_type_id"`
	Course_code      string          `json:"course_code"`
	Course_name      string          `json:"course_name"`
	Course_amount    int             `json:"course_amount"`
	Course_unit      string          `json:"course_unit"`
	Course_qtyset    int             `json:"course_qtyset"`
	Course_cost      float64         `json:"course_cost"`
	Course_opd       float64         `json:"course_opd"`
	Course_ipd       float64         `json:"course_ipd"`
	Course_lock_drug int             `json:"course_lock_drug"`
	Course_image_1   string          `json:"course_image_1"`
	Course_image_2   string          `json:"course_image_2"`
	Course_image_3   string          `json:"course_image_3"`
	Course_image_4   string          `json:"course_image_4"`
	Course_detail    string          `json:"course_detail"`
	Subs             []ItemCourseSub `json:"subs" gorm:"-"`
}

type ItemCourseSub struct {
	Id                 int                 `json:"id"`
	Course_type_id     int                 `json:"course_type_id"`
	Course_code        string              `json:"course_code"`
	Course_name        string              `json:"course_name"`
	Course_amount      int                 `json:"course_amount"`
	Course_unit        string              `json:"course_unit"`
	Course_qtyset      int                 `json:"course_qtyset"`
	Course_cost        float64             `json:"course_cost"`
	Course_opd         float64             `json:"course_opd"`
	Course_ipd         float64             `json:"course_ipd"`
	Course_image_1     string              `json:"course_image_1"`
	Course_image_2     string              `json:"course_image_2"`
	Course_image_3     string              `json:"course_image_3"`
	Course_image_4     string              `json:"course_image_4"`
	Course_detail      string              `json:"course_detail"`
	Course_list_qtyset int                 `json:"course_list_qtyset"`
	Course_list_opd    float64             `json:"course_list_opd"`
	Course_list_ipd    float64             `json:"course_list_ipd"`
	Course_lock_drug   int                 `json:"course_lock_drug"`
	Label              string              `json:"label"`
	Is_set             int                 `json:"is_set"`
	Id_set             *int                `json:"id_set"`
	Products           []ItemProductSubSet `json:"products" gorm:"-"`
}

type ItemProductSubSet struct {
	Id               int     `json:"id"`
	Product_id       int     `json:"product_id"`
	Product_store_id int     `json:"product_store_id"`
	Product_units_id int     `json:"product_units_id"`
	Pd_code          string  `json:"pd_code"`
	Pd_name          string  `json:"pd_name"`
	U_name           string  `json:"u_name"`
	Pu_amount        float64 `json:"pu_amount"`
	Pu_rate          float64 `json:"Pu_rate"`
	Balance          float64 `json:"balance"`
	Label            string  `json:"label"`
	Is_set           int     `json:"is_set"`
	Id_set           *int    `json:"id_set"`
	Qty_set          float64 `json:"qty_set"`
}

type ItemChecking struct {
	Id               int               `json:"id"`
	Shop_id          int               `json:"shop_id"`
	Shop_name        string            `json:"shop_name"`
	Shop_name_en     string            `json:"shop_name_en"`
	Shop_code        string            `json:"shop_code"`
	Checking_type_id int               `json:"checking_type_id"`
	Checking_code    string            `json:"checking_code"`
	Checking_amount  int               `json:"checking_amount"`
	Checking_name    string            `json:"checking_name"`
	Checking_unit    string            `json:"checking_unit"`
	Checking_cost    float64           `json:"checking_cost"`
	Checking_opd     float64           `json:"checking_opd"`
	Checking_ipd     float64           `json:"checking_ipd"`
	Checking_image_1 string            `json:"checking_image_1"`
	Checking_image_2 string            `json:"checking_image_2"`
	Checking_image_3 string            `json:"checking_image_3"`
	Checking_image_4 string            `json:"checking_image_4"`
	Checking_detail  string            `json:"checking_detail"`
	Subs             []ItemCheckingSub `json:"subs" gorm:"-"`
}

type ItemCheckingSub struct {
	Id                int                 `json:"id"`
	Checking_type_id  int                 `json:"checking_type_id"`
	Checking_code     string              `json:"checking_code"`
	Checking_amount   int                 `json:"checking_amount"`
	Checking_name     string              `json:"checking_name"`
	Checking_unit     string              `json:"checking_unit"`
	Checking_opd      float64             `json:"checking_opd"`
	Checking_cost     float64             `json:"checking_cost"`
	Checking_ipd      float64             `json:"checking_ipd"`
	Checking_image_1  string              `json:"checking_image_1"`
	Checking_image_2  string              `json:"checking_image_2"`
	Checking_image_3  string              `json:"checking_image_3"`
	Checking_image_4  string              `json:"checking_image_4"`
	Checking_detail   string              `json:"checking_detail"`
	Checking_list_opd float64             `json:"checking_list_opd"`
	Checking_list_ipd float64             `json:"checking_list_ipd"`
	Label             string              `json:"label"`
	Is_set            int                 `json:"is_set"`
	Id_set            *int                `json:"id_set"`
	Products          []ItemProductSubSet `json:"products" gorm:"-"`
}
