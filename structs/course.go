package structs

type ResponsePaginationCourse struct {
	Result_data   []CourseList `json:"result_data"`
	Count_of_page int          `json:"count_of_page"`
	Count_all     int          `json:"count_all"`
}

type ObjPayloadSearchCourse struct {
	Search       *string `json:"search" binding:"required,omitempty"`
	Is_active    *string `json:"is_active" binding:"required,omitempty"`
	Shop_id      int     `json:"shop_id"`
	ActivePage   int     `json:"active_page" binding:"required"`
	CategoryId   int     `json:"category_id" binding:"required"`
	CourseTypeId int     `json:"course_type_id" binding:"required"`
	PerPage      int     `json:"per_page" binding:"required"`
}

type ObjPayloadSearchCourseList struct {
	Search     *string `json:"search" binding:"required,omitempty"`
	Is_active  *string `json:"is_active" binding:"required,omitempty"`
	Shop_id    int     `json:"shop_id"`
	Course_id  int     `json:"course_id" binding:"required"`
	ActivePage int     `json:"active_page" binding:"required"`
	PerPage    int     `json:"per_page" binding:"required"`
}

type ObjPayloadCourse struct {
	Id             int    `json:"id" binding:"required"`
	Shop_id        int    `json:"shop_id"`
	Course_type_id int    `json:"course_type_id" binding:"required"`
	Course_name    string `json:"course_name" binding:"required"`
	Course_is_del  int    `json:"course_is_del"`
}

type ResponsePaginationCourseNotSet struct {
	Result_data   []CourseNotSetList `json:"result_data"`
	Count_of_page int                `json:"count_of_page"`
	Count_all     int                `json:"count_all"`
}

type CourseList struct {
	Id                 int     `json:"id"`
	Shop_id            int     `json:"shop_id"`
	Category_id        int     `json:"category_id"`
	Course_type_id     int     `json:"course_type_id"`
	Course_code        string  `json:"course_code"`
	Course_name        string  `json:"course_name"`
	Course_amount      int     `json:"course_amount"`
	Course_unit        string  `json:"course_unit"`
	Course_use_date    int     `json:"course_use_date"`
	Course_exp_date    int     `json:"course_exp_date"`
	Course_lock_drug   int     `json:"course_lock_drug"`
	Course_opd         float64 `json:"course_opd" gorm:"type:decimal(10,2)"`
	Course_ipd         float64 `json:"course_ipd" gorm:"type:decimal(10,2)"`
	Course_cost        float64 `json:"course_cost" gorm:"type:decimal(10,2)"`
	Course_fee_df      float64 `json:"course_fee_df" gorm:"type:decimal(10,2)"`
	Course_fee_nr      float64 `json:"course_fee_nr" gorm:"type:decimal(10,2)"`
	Course_fee_tr      float64 `json:"course_fee_tr" gorm:"type:decimal(10,2)"`
	Course_fee         float64 `json:"course_fee" gorm:"type:decimal(10,2)"`
	Acc_code_id_cost   *int    `json:"acc_code_id_cost" binding:"required,omitempty"`
	Acc_code_id_fee_df *int    `json:"acc_code_id_fee_df" binding:"required,omitempty"`
	Acc_code_id_fee    *int    `json:"acc_code_id_fee" binding:"required,omitempty"`
	Acc_code_id_com    *int    `json:"acc_code_id_com" binding:"required,omitempty"`
	Course_image       *string `json:"course_image" binding:"required,omitempty"`
	Course_is_active   int     `json:"course_is_active"`
	Course_is_del      int     `json:"course_is_del"`
	Course_create      string  `json:"course_create"`
	Course_update      string  `json:"course_update"`
	Category
}

type CourseNotSetList struct {
	Id                 int     `json:"course_id"`
	Shop_id            int     `json:"shop_id"`
	Category_id        int     `json:"category_id"`
	Course_type_id     int     `json:"course_type_id"`
	Course_code        string  `json:"course_code"`
	Course_name        string  `json:"course_name"`
	Course_amount      int     `json:"course_amount"`
	Course_unit        string  `json:"course_unit"`
	Course_use_date    int     `json:"course_use_date"`
	Course_exp_date    int     `json:"course_exp_date"`
	Course_lock_drug   int     `json:"course_lock_drug"`
	Course_opd         float64 `json:"course_opd"`
	Course_ipd         float64 `json:"course_ipd"`
	Course_cost        float64 `json:"course_cost"`
	Course_fee_df      float64 `json:"course_fee_df"`
	Course_fee_nr      float64 `json:"course_fee_nr"`
	Course_fee_tr      float64 `json:"course_fee_tr"`
	Course_fee         float64 `json:"course_fee"`
	Acc_code_id_cost   *int    `json:"acc_code_id_cost" binding:"required,omitempty"`
	Acc_code_id_fee_df *int    `json:"acc_code_id_fee_df" binding:"required,omitempty"`
	Acc_code_id_fee    *int    `json:"acc_code_id_fee" binding:"required,omitempty"`
	Acc_code_id_com    *int    `json:"acc_code_id_com" binding:"required,omitempty"`
	Course_image       *string `json:"course_image" binding:"required,omitempty"`
	Course_is_active   int     `json:"course_is_active"`
	Course_is_del      int     `json:"course_is_del"`
	Course_create      string  `json:"course_create"`
	Course_update      string  `json:"course_update"`
	Category
}

type CourseByID struct {
	Id                 int     `json:"id" binding:"required"`
	Shop_id            int     `json:"shop_id"`
	Category_id        int     `json:"category_id"`
	Course_type_id     int     `json:"course_type_id"`
	Course_code        string  `json:"course_code"`
	Course_name        string  `json:"course_name"`
	Course_amount      int     `json:"course_amount"`
	Course_unit        string  `json:"course_unit"`
	Course_use_date    int     `json:"course_use_date"`
	Course_exp_date    int     `json:"course_exp_date"`
	Course_lock_drug   int     `json:"course_lock_drug"`
	Course_opd         float64 `json:"course_opd" gorm:"type:decimal(10,2)"`
	Course_ipd         float64 `json:"course_ipd" gorm:"type:decimal(10,2)"`
	Course_cost        float64 `json:"course_cost" gorm:"type:decimal(10,2)"`
	Course_fee_df      float64 `json:"course_fee_df" gorm:"type:decimal(10,2)"`
	Course_fee_nr      float64 `json:"course_fee_nr" gorm:"type:decimal(10,2)"`
	Course_fee_tr      float64 `json:"course_fee_tr" gorm:"type:decimal(10,2)"`
	Course_fee         float64 `json:"course_fee" gorm:"type:decimal(10,2)"`
	Acc_code_id_cost   *int    `json:"acc_code_id_cost" binding:"required,omitempty"`
	Acc_code_id_fee_df *int    `json:"acc_code_id_fee_df" binding:"required,omitempty"`
	Acc_code_id_fee    *int    `json:"acc_code_id_fee" binding:"required,omitempty"`
	Acc_code_id_com    *int    `json:"acc_code_id_com" binding:"required,omitempty"`
	Course_image       *string `json:"course_image" binding:"required,omitempty"`
	Course_is_active   int     `json:"course_is_active"`
	Course_is_del      int     `json:"course_is_del"`
	Course_create      string  `json:"course_create"`
	Course_update      string  `json:"course_update"`
	Category
}

// ////////////////////////////////////////////////////////////////////////////

type ObjPayloadSearchProductCourse struct {
	Search     *string `json:"search" binding:"required,omitempty"`
	Shop_id    int     `json:"shop_id"`
	ActivePage int     `json:"active_page" binding:"required"`
	PerPage    int     `json:"per_page" binding:"required"`
}

type ResponsePaginationCourseProduct struct {
	Result_data   []Product `json:"result_data"`
	Count_of_page int       `json:"count_of_page"`
	Count_all     int       `json:"count_all"`
}

// type CourseList struct {
// 	Id               int     `json:"id"`
// 	Shop_id          int     `json:"shop_id"`
// 	Category_id      int     `json:"category_id"`
// 	Course_type_id   int     `json:"course_type_id"`
// 	Course_code      string  `json:"course_code"`
// 	Course_code_acc  string  `json:"course_code_acc"`
// 	Course_name      string  `json:"course_name"`
// 	Course_unit      string  `json:"course_unit"`
// 	Course_opd       float64 `json:"course_opd"`
// 	Course_ipd       float64 `json:"course_ipd"`
// 	Course_cost      float64 `json:"course_cost"`
// 	Course_fee_df    float64 `json:"course_fee_df"`
// 	Course_fee_nr    float64 `json:"course_fee_nr"`
// 	Course_fee_tr    float64 `json:"course_fee_tr"`
// 	Course_fee       float64 `json:"course_fee"`
// 	Course_is_active int     `json:"course_is_active"`
// 	Course_is_del    int     `json:"course_is_del"`
// 	Course_create    string  `json:"course_create"`
// 	Course_update    string  `json:"course_update"`
// 	Category
// }

type CourseDetail struct {
	Id                 int     `json:"id"`
	Shop_id            int     `json:"shop_id"`
	Category_id        int     `json:"category_id"`
	Course_type_id     int     `json:"course_type_id"`
	Course_code        string  `json:"course_code"`
	Course_name        string  `json:"course_name"`
	Course_amount      int     `json:"course_amount"`
	Course_unit        string  `json:"course_unit"`
	Course_use_date    int     `json:"course_use_date"`
	Course_exp_date    int     `json:"course_exp_date"`
	Course_lock_drug   int     `json:"course_lock_drug"`
	Course_opd         float64 `json:"course_opd"`
	Course_ipd         float64 `json:"course_ipd"`
	Category_eclaim_id int     `json:"category_eclaim_id"`
	Course_ofc         float64 `json:"course_ofc"`
	Course_lgo         float64 `json:"course_lgo"`
	Course_ucs         float64 `json:"course_ucs"`
	Course_sss         float64 `json:"course_sss"`
	Course_nhs         float64 `json:"course_nhs"`
	Course_ssi         float64 `json:"course_ssi"`
	Course_cost        float64 `json:"course_cost"`
	Course_fee_df      float64 `json:"course_fee_df"`
	Course_fee_nr      float64 `json:"course_fee_nr"`
	Course_fee_tr      float64 `json:"course_fee_tr"`
	Course_fee         float64 `json:"course_fee"`
	Acc_code_id_cost   *int    `json:"acc_code_id_cost" binding:"required,omitempty"`
	Acc_code_id_fee_df *int    `json:"acc_code_id_fee_df" binding:"required,omitempty"`
	Acc_code_id_fee    *int    `json:"acc_code_id_fee" binding:"required,omitempty"`
	Acc_code_id_com    *int    `json:"acc_code_id_com" binding:"required,omitempty"`
	Course_image       *string `json:"course_image" binding:"required,omitempty"`
	Course_is_active   int     `json:"course_is_active"`
	Course_create      string  `json:"course_create"`
	Course_update      string  `json:"course_update"`
	Course_set_id      int     `json:"course_set_id"`
	Category
	CourseListProduct *[]Course_product `json:"course_product" gorm:"foreignKey:Id;references:Category_id"`
	CourseList        *[]CourseSubList  `json:"course_list" gorm:"foreignKey:Id;references:Category_id"`
}

type ObjPayloadAddProductCourse struct {
	Shop_id            int                   `json:"shop_id"`
	Category_id        int                   `json:"category_id" binding:"required"`
	Course_type_id     int                   `json:"course_type_id" binding:"required"`
	Course_code        string                `json:"course_code" binding:"required"`
	Course_name        string                `json:"course_name" binding:"required"`
	Course_amount      int                   `json:"course_amount"`
	Course_unit        string                `json:"course_unit" binding:"required"`
	Course_use_date    int                   `json:"course_use_date"`
	Course_exp_date    int                   `json:"course_exp_date"`
	Course_lock_drug   int                   `json:"course_lock_drug"`
	Course_opd         *float64              `json:"course_opd" binding:"required,omitempty"`
	Course_ipd         *float64              `json:"course_ipd" binding:"required,omitempty"`
	Category_eclaim_id *int                  `json:"category_eclaim_id" binding:"required,omitempty"`
	Course_ofc         *float64              `json:"course_ofc" binding:"required,omitempty"`
	Course_lgo         *float64              `json:"course_lgo" binding:"required,omitempty"`
	Course_ucs         *float64              `json:"course_ucs" binding:"required,omitempty"`
	Course_sss         *float64              `json:"course_sss" binding:"required,omitempty"`
	Course_nhs         *float64              `json:"course_nhs" binding:"required,omitempty"`
	Course_ssi         *float64              `json:"course_ssi" binding:"required,omitempty"`
	Course_cost        *float64              `json:"course_cost" binding:"required,omitempty"`
	Course_fee_df      *float64              `json:"course_fee_df" binding:"required,omitempty"`
	Course_fee_nr      *float64              `json:"course_fee_nr" binding:"required,omitempty"`
	Course_fee_tr      *float64              `json:"course_fee_tr" binding:"required,omitempty"`
	Course_fee         *float64              `json:"course_fee" binding:"required,omitempty"`
	Acc_code_id_cost   *int                  `json:"acc_code_id_cost" binding:"required,omitempty"`
	Acc_code_id_fee_df *int                  `json:"acc_code_id_fee_df" binding:"required,omitempty"`
	Acc_code_id_fee    *int                  `json:"acc_code_id_fee" binding:"required,omitempty"`
	Acc_code_id_com    *int                  `json:"acc_code_id_com" binding:"required,omitempty"`
	Course_image       *string               `json:"course_image" binding:"required,omitempty"`
	Course_is_active   int                   `json:"course_is_active"`
	Course_is_del      int                   `json:"course_is_del"`
	Course_create      string                `json:"course_create"`
	Course_update      string                `json:"course_update"`
	Course_product     *[]Obj_Course_product `json:"course_product" binding:"required,omitempty"`
}

type ObjPayloadAddListCourse struct {
	Shop_id            int                `json:"shop_id"`
	Category_id        int                `json:"category_id" binding:"required"`
	Course_type_id     int                `json:"course_type_id" binding:"required"`
	Course_code        string             `json:"course_code" binding:"required"`
	Course_name        string             `json:"course_name" binding:"required"`
	Course_amount      int                `json:"course_amount"`
	Course_unit        string             `json:"course_unit" binding:"required"`
	Course_use_date    int                `json:"course_use_date"`
	Course_exp_date    int                `json:"course_exp_date"`
	Course_lock_drug   int                `json:"course_lock_drug"`
	Course_opd         *float64           `json:"course_opd" binding:"required,omitempty"`
	Course_ipd         *float64           `json:"course_ipd" binding:"required,omitempty"`
	Category_eclaim_id *int               `json:"category_eclaim_id" binding:"required,omitempty"`
	Course_ofc         *float64           `json:"course_ofc" binding:"required,omitempty"`
	Course_lgo         *float64           `json:"course_lgo" binding:"required,omitempty"`
	Course_ucs         *float64           `json:"course_ucs" binding:"required,omitempty"`
	Course_sss         *float64           `json:"course_sss" binding:"required,omitempty"`
	Course_nhs         *float64           `json:"course_nhs" binding:"required,omitempty"`
	Course_ssi         *float64           `json:"course_ssi" binding:"required,omitempty"`
	Course_cost        *float64           `json:"course_cost" binding:"required,omitempty"`
	Course_fee_df      *float64           `json:"course_fee_df" binding:"required,omitempty"`
	Course_fee_nr      *float64           `json:"course_fee_nr" binding:"required,omitempty"`
	Course_fee_tr      *float64           `json:"course_fee_tr" binding:"required,omitempty"`
	Course_fee         *float64           `json:"course_fee" binding:"required,omitempty"`
	Acc_code_id_cost   *int               `json:"acc_code_id_cost" binding:"required,omitempty"`
	Acc_code_id_fee_df *int               `json:"acc_code_id_fee_df" binding:"required,omitempty"`
	Acc_code_id_fee    *int               `json:"acc_code_id_fee" binding:"required,omitempty"`
	Acc_code_id_com    *int               `json:"acc_code_id_com" binding:"required,omitempty"`
	Course_image       *string            `json:"course_image" binding:"required,omitempty"`
	Course_is_active   int                `json:"course_is_active"`
	Course_is_del      int                `json:"course_is_del"`
	Course_create      string             `json:"course_create"`
	Course_update      string             `json:"course_update"`
	Course_List        *[]Obj_Course_list `json:"course_list" binding:"required,omitempty"`
}

type ObjPayloadEditProductCourse struct {
	Id                 int                   `json:"id" binding:"required"`
	Shop_id            int                   `json:"shop_id"`
	Category_id        int                   `json:"category_id" binding:"required"`
	Course_type_id     int                   `json:"course_type_id" binding:"required"`
	Course_code        string                `json:"course_code" binding:"required"`
	Course_name        string                `json:"course_name" binding:"required"`
	Course_amount      int                   `json:"course_amount"`
	Course_unit        string                `json:"course_unit" binding:"required"`
	Course_use_date    *int                  `json:"course_use_date" binding:"required,omitempty"`
	Course_exp_date    *int                  `json:"course_exp_date" binding:"required,omitempty"`
	Course_lock_drug   *int                  `json:"course_lock_drug" binding:"required,omitempty"`
	Course_opd         *float64              `json:"course_opd" binding:"required,omitempty"`
	Course_ipd         *float64              `json:"course_ipd" binding:"required,omitempty"`
	Category_eclaim_id *int                  `json:"category_eclaim_id" binding:"required,omitempty"`
	Course_ofc         *float64              `json:"course_ofc" binding:"required,omitempty"`
	Course_lgo         *float64              `json:"course_lgo" binding:"required,omitempty"`
	Course_ucs         *float64              `json:"course_ucs" binding:"required,omitempty"`
	Course_sss         *float64              `json:"course_sss" binding:"required,omitempty"`
	Course_nhs         *float64              `json:"course_nhs" binding:"required,omitempty"`
	Course_ssi         *float64              `json:"course_ssi" binding:"required,omitempty"`
	Course_cost        *float64              `json:"course_cost" binding:"required,omitempty"`
	Course_fee_df      *float64              `json:"course_fee_df" binding:"required,omitempty"`
	Course_fee_nr      *float64              `json:"course_fee_nr" binding:"required,omitempty"`
	Course_fee_tr      *float64              `json:"course_fee_tr" binding:"required,omitempty"`
	Course_fee         *float64              `json:"course_fee" binding:"required,omitempty"`
	Acc_code_id_cost   *int                  `json:"acc_code_id_cost" binding:"required,omitempty"`
	Acc_code_id_fee_df *int                  `json:"acc_code_id_fee_df" binding:"required,omitempty"`
	Acc_code_id_fee    *int                  `json:"acc_code_id_fee" binding:"required,omitempty"`
	Acc_code_id_com    *int                  `json:"acc_code_id_com" binding:"required,omitempty"`
	Course_image       *string               `json:"course_image" binding:"required,omitempty"`
	Course_is_active   int                   `json:"course_is_active"`
	Course_is_del      int                   `json:"course_is_del"`
	Course_update      string                `json:"course_update"`
	Course_product     *[]Obj_Course_product `json:"course_product" binding:"required,omitempty"`
}

type ObjPayloadEditListCourse struct {
	Id                 int                `json:"id" binding:"required"`
	Shop_id            int                `json:"shop_id"`
	Category_id        int                `json:"category_id" binding:"required"`
	Course_type_id     int                `json:"course_type_id" binding:"required"`
	Course_code        string             `json:"course_code" binding:"required"`
	Course_name        string             `json:"course_name" binding:"required"`
	Course_amount      int                `json:"course_amount"`
	Course_unit        string             `json:"course_unit" binding:"required"`
	Course_use_date    *int               `json:"course_use_date" binding:"required,omitempty"`
	Course_exp_date    *int               `json:"course_exp_date" binding:"required,omitempty"`
	Course_lock_drug   int                `json:"course_lock_drug"`
	Course_opd         *float64           `json:"course_opd" binding:"required,omitempty"`
	Course_ipd         *float64           `json:"course_ipd" binding:"required,omitempty"`
	Category_eclaim_id *int               `json:"category_eclaim_id" binding:"required,omitempty"`
	Course_ofc         *float64           `json:"course_ofc" binding:"required,omitempty"`
	Course_lgo         *float64           `json:"course_lgo" binding:"required,omitempty"`
	Course_ucs         *float64           `json:"course_ucs" binding:"required,omitempty"`
	Course_sss         *float64           `json:"course_sss" binding:"required,omitempty"`
	Course_nhs         *float64           `json:"course_nhs" binding:"required,omitempty"`
	Course_ssi         *float64           `json:"course_ssi" binding:"required,omitempty"`
	Course_cost        *float64           `json:"course_cost" binding:"required,omitempty"`
	Course_fee_df      *float64           `json:"course_fee_df" binding:"required,omitempty"`
	Course_fee_nr      *float64           `json:"course_fee_nr" binding:"required,omitempty"`
	Course_fee_tr      *float64           `json:"course_fee_tr" binding:"required,omitempty"`
	Course_fee         *float64           `json:"course_fee" binding:"required,omitempty"`
	Acc_code_id_cost   *int               `json:"acc_code_id_cost" binding:"required,omitempty"`
	Acc_code_id_fee_df *int               `json:"acc_code_id_fee_df" binding:"required,omitempty"`
	Acc_code_id_fee    *int               `json:"acc_code_id_fee" binding:"required,omitempty"`
	Acc_code_id_com    *int               `json:"acc_code_id_com" binding:"required,omitempty"`
	Course_image       *string            `json:"course_image" binding:"required,omitempty"`
	Course_is_active   int                `json:"course_is_active"`
	Course_is_del      int                `json:"course_is_del"`
	Course_create      string             `json:"course_create"`
	Course_update      string             `json:"course_update"`
	Course_set_id      int                `json:"course_set_id"`
	Course_List        *[]Obj_Course_list `json:"course_list" binding:"required,omitempty"`
}

type CourseAction struct {
	Id                 int     `json:"id"`
	Shop_id            int     `json:"shop_id"`
	Category_id        int     `json:"category_id"`
	Course_type_id     int     `json:"course_type_id"`
	Course_code        string  `json:"course_code"`
	Course_name        string  `json:"course_name"`
	Course_amount      int     `json:"course_amount"`
	Course_unit        string  `json:"course_unit"`
	Course_use_date    int     `json:"course_use_date"`
	Course_exp_date    int     `json:"course_exp_date"`
	Course_lock_drug   int     `json:"course_lock_drug"`
	Course_opd         float64 `json:"course_opd"`
	Course_ipd         float64 `json:"course_ipd"`
	Category_eclaim_id *int    `json:"category_eclaim_id"`
	Course_ofc         float64 `json:"course_ofc"`
	Course_lgo         float64 `json:"course_lgo"`
	Course_ucs         float64 `json:"course_ucs"`
	Course_sss         float64 `json:"course_sss"`
	Course_nhs         float64 `json:"course_nhs"`
	Course_ssi         float64 `json:"course_ssi"`
	Course_cost        float64 `json:"course_cost"`
	Course_fee_df      float64 `json:"course_fee_df"`
	Course_fee_nr      float64 `json:"course_fee_nr"`
	Course_fee_tr      float64 `json:"course_fee_tr"`
	Course_fee         float64 `json:"course_fee"`
	Acc_code_id_cost   *int    `json:"acc_code_id_cost" binding:"required,omitempty"`
	Acc_code_id_fee_df *int    `json:"acc_code_id_fee_df" binding:"required,omitempty"`
	Acc_code_id_fee    *int    `json:"acc_code_id_fee" binding:"required,omitempty"`
	Acc_code_id_com    *int    `json:"acc_code_id_com" binding:"required,omitempty"`
	Course_image       *string `json:"course_image" binding:"required,omitempty"`
	Course_is_active   int     `json:"course_is_active"`
	Course_is_del      int     `json:"course_is_del"`
	Course_create      string  `json:"course_create"`
	Course_update      string  `json:"course_update"`
}

type Obj_Course_product struct {
	Id           int     `json:"id"`
	Course_id    int     `json:"course_id"`
	Product_id   int     `json:"product_id" binding:"required"`
	Cp_amount    float64 `json:"cp_amount"`
	Cp_is_active int     `json:"cp_is_active"`
	Cp_is_del    int     `json:"cp_is_del"`
}

type Obj_Course_list struct {
	Id              int     `json:"id"`
	Course_set_id   int     `json:"course_set_id"`
	Course_id       int     `json:"course_id"`
	Course_list_opd float64 `json:"course_list_opd"`
	Course_list_ipd float64 `json:"course_list_ipd"`
}

type Obj_Course_set struct {
	Id        int `json:"id"`
	Course_id int `json:"course_id"`
	// Course_list_opd float64 `json:"course_opd"`
	// Course_list_ipd float64 `json:"course_ipd"`
}

type Course_product struct {
	Id           int     `json:"id"`
	Course_id    int     `json:"course_id"`
	Product_id   int     `json:"product_id"`
	Cp_amount    float64 `json:"cp_amount"`
	Cp_is_active int     `json:"cp_is_active"`
	Cp_is_del    int     `json:"cp_is_del"`
	Product
}

type CourseSubList struct {
	Id              int     `json:"id"`
	Course_set_id   int     `json:"course_set_id"`
	Course_id       int     `json:"course_id"`
	Course_list_opd float64 `json:"course_list_opd"`
	Course_list_ipd float64 `json:"course_list_ipd"`
	CourseList      `gorm:"foreignKey:course_id;references:Id"`
}

type CourseSubListId struct {
	Course_id int `json:"course_id"`
}

type DocCourse struct {
	ShopId                int    `json:"shop_id"`
	Course_id_default     string `json:"course_id_default"`
	Course_number_default string `json:"course_number_default"`
	Course_number_digit   int    `json:"course_number_digit"`
	Course_type           int    `json:"course_type"`
}

type LogCourse struct {
	Username   string `json:"username"`
	Log_type   string `json:"log_type"`
	Log_text   string `json:"log_text"`
	Log_create string `json:"log_create"`
}

type ObjQueryCourse struct {
	ID             int     `json:"id"`
	ShopId         int     `json:"shop_id"`
	CategoryId     int     `json:"category_id"`
	CourseTypeId   int     `json:"course_type_id"`
	CourseCode     string  `json:"course_code"`
	CourseName     string  `json:"course_name"`
	CourseAmount   int     `json:"course_amount"`
	CourseUnit     string  `json:"course_unit"`
	CourseUseDate  int     `json:"course_use_date"`
	CourseExpDate  int     `json:"course_exp_date"`
	CourseLockDrug int     `json:"course_lock_drug"`
	CourseOpd      float64 `json:"course_opd"`
	CourseIpd      float64 `json:"course_ipd"`
	CourseCost     float64 `json:"course_cost"`
	CourseFeeDf    float64 `json:"course_fee_df"`
	CourseFeeNr    float64 `json:"course_fee_nr"`
	CourseFeeTr    float64 `json:"course_fee_tr"`
	CourseFee      float64 `json:"course_fee"`
	AccCodeIdCost  *int    `json:"acc_code_id_cost"`
	AccCodeIdFeeDf *int    `json:"acc_code_id_fee_df"`
	AccCodeIdFee   *int    `json:"acc_code_id_fee"`
	AccCodeIdCom   *int    `json:"acc_code_id_com"`
	CourseImage    *string `json:"course_image"`
	CourseIsActive int     `json:"course_is_active"`
	CourseCreate   string  `json:"course_create"`
	CourseUpdate   string  `json:"course_update"`
}

type ObjCheckExcelCourse struct {
	ShopId         int      `json:"shop_id"`
	CategoryId     int      `json:"category_id"`
	CourseCode     string   `json:"course_code"`
	CourseName     string   `json:"course_name"`
	CourseOpd      float64  `json:"course_opd"`
	CourseIpd      float64  `json:"course_ipd"`
	CourseUnit     string   `json:"course_unit"`
	CourseCost     float64  `json:"course_cost"`
	CourseFeeDf    float64  `json:"course_fee_df"`
	CourseFeeNr    float64  `json:"course_fee_nr"`
	CourseFeeTr    float64  `json:"course_fee_tr"`
	CourseFee      float64  `json:"course_fee"`
	CourseExpDate  int      `json:"course_exp_date"`
	CourseTypeId   int      `json:"course_type_id"`
	CourseLockDrug int      `json:"course_lock_drug"`
	Message        []string `json:"message"`
}

type ObjPayloadImportExcelCourse struct {
	ImportData []ObjCheckExcelCourse `json:"import_data"`
}
