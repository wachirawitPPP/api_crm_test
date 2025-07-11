package structs

type ObjPayloadSearchChecking struct {
	Search         *string `json:"search" binding:"required,omitempty"`
	Is_active      *string `json:"is_active" binding:"required,omitempty"`
	Shop_id        int     `json:"shop_id"`
	CategoryId     int     `json:"category_id" binding:"required"`
	CheckingTypeId int     `json:"checking_type_id" binding:"required"`
	ActivePage     int     `json:"active_page" binding:"required"`
	PerPage        int     `json:"per_page" binding:"required"`
}

type ObjPayloadSearchCheckingList struct {
	Search      *string `json:"search" binding:"required,omitempty"`
	Is_active   *string `json:"is_active" binding:"required,omitempty"`
	Shop_id     int     `json:"shop_id"`
	Checking_id int     `json:"checking_id" binding:"required"`
	ActivePage  int     `json:"active_page" binding:"required"`
	PerPage     int     `json:"per_page" binding:"required"`
}

type ObjPayloadSearchProduct struct {
	Search     *string `json:"search" binding:"required,omitempty"`
	Shop_id    int     `json:"shop_id"`
	ActivePage int     `json:"active_page" binding:"required"`
	PerPage    int     `json:"per_page" binding:"required"`
}

type ResponsePaginationEmpty struct {
	Result_data   []string `json:"result_data"`
	Count_of_page int      `json:"count_of_page"`
	Count_all     int      `json:"count_all"`
}

type ResponsePaginationChecking struct {
	Result_data   []CheckingList `json:"result_data"`
	Count_of_page int            `json:"count_of_page"`
	Count_all     int            `json:"count_all"`
}

type ResponsePaginationCheckingNotSet struct {
	Result_data   []CheckingNotSetList `json:"result_data"`
	Count_of_page int                  `json:"count_of_page"`
	Count_all     int                  `json:"count_all"`
}

type ResponsePaginationCheckingProduct struct {
	Result_data   []Product `json:"result_data"`
	Count_of_page int       `json:"count_of_page"`
	Count_all     int       `json:"count_all"`
}

type CheckingList struct {
	Id                  int     `json:"id"`
	Shop_id             int     `json:"shop_id"`
	Category_id         int     `json:"category_id"`
	Checking_type_id    int     `json:"checking_type_id"`
	Checking_code       string  `json:"checking_code"`
	Checking_code_acc   string  `json:"checking_code_acc"`
	Checking_name_acc   string  `json:"checking_name_acc"`
	Checking_name       string  `json:"checking_name"`
	Checking_unit       string  `json:"checking_unit"`
	Checking_opd        float64 `json:"checking_opd"`
	Checking_ipd        float64 `json:"checking_ipd"`
	Checking_cost       float64 `json:"checking_cost"`
	Checking_fee_df     float64 `json:"checking_fee_df"`
	Checking_fee_nr     float64 `json:"checking_fee_nr"`
	Checking_fee_tr     float64 `json:"checking_fee_tr"`
	Checking_fee        float64 `json:"checking_fee"`
	Checking_is_active  int     `json:"checking_is_active"`
	Checking_is_del     int     `json:"checking_is_del"`
	Checking_is_labplus int     `json:"checking_is_labplus"`
	Checking_create     string  `json:"checking_create"`
	Checking_update     string  `json:"checking_update"`
	Category
}

type CheckingNotSetList struct {
	Id                  int     `json:"checking_id"`
	Shop_id             int     `json:"shop_id"`
	Category_id         int     `json:"category_id"`
	Checking_type_id    int     `json:"checking_type_id"`
	Checking_code       string  `json:"checking_code"`
	Checking_code_acc   string  `json:"checking_code_acc"`
	Checking_name_acc   string  `json:"checking_name_acc"`
	Checking_name       string  `json:"checking_name"`
	Checking_unit       string  `json:"checking_unit"`
	Checking_opd        float64 `json:"checking_opd"`
	Checking_ipd        float64 `json:"checking_ipd"`
	Checking_cost       float64 `json:"checking_cost"`
	Checking_fee_df     float64 `json:"checking_fee_df"`
	Checking_fee_nr     float64 `json:"checking_fee_nr"`
	Checking_fee_tr     float64 `json:"checking_fee_tr"`
	Checking_fee        float64 `json:"checking_fee"`
	Checking_is_active  int     `json:"checking_is_active"`
	Checking_is_del     int     `json:"checking_is_del"`
	Checking_is_labplus int     `json:"checking_is_labplus"`
	Checking_create     string  `json:"checking_create"`
	Checking_update     string  `json:"checking_update"`
	Category
}

type CheckingDetail struct {
	Id                  int     `json:"id"`
	Shop_id             int     `json:"shop_id"`
	Category_id         int     `json:"category_id"`
	Checking_type_id    int     `json:"checking_type_id"`
	Checking_code       string  `json:"checking_code"`
	Checking_code_acc   string  `json:"checking_code_acc"`
	Checking_name_acc   string  `json:"checking_name_acc"`
	Checking_name       string  `json:"checking_name"`
	Checking_unit       string  `json:"checking_unit"`
	Checking_opd        float64 `json:"checking_opd"`
	Checking_ipd        float64 `json:"checking_ipd"`
	Category_eclaim_id  int     `json:"category_eclaim_id"`
	Checking_ofc        float64 `json:"checking_ofc"`
	Checking_lgo        float64 `json:"checking_lgo"`
	Checking_ucs        float64 `json:"checking_ucs"`
	Checking_sss        float64 `json:"checking_sss"`
	Checking_nhs        float64 `json:"checking_nhs"`
	Checking_ssi        float64 `json:"checking_ssi"`
	Checking_cost       float64 `json:"checking_cost"`
	Checking_fee_df     float64 `json:"checking_fee_df"`
	Checking_fee_nr     float64 `json:"checking_fee_nr"`
	Checking_fee_tr     float64 `json:"checking_fee_tr"`
	Checking_fee        float64 `json:"checking_fee"`
	Acc_code_id_cost    int     `json:"acc_code_id_cost"`
	Acc_code_id_fee_df  int     `json:"acc_code_id_fee_df" `
	Acc_code_id_fee     int     `json:"acc_code_id_fee"`
	Acc_code_id_com     int     `json:"acc_code_id_com"`
	Checking_image      string  `json:"checking_image"`
	Checking_is_active  int     `json:"checking_is_active"`
	Checking_is_labplus int     `json:"checking_is_labplus"`
	Checking_create     string  `json:"checking_create"`
	Checking_update     string  `json:"checking_update"`
	Checking_set_id     int     `json:"checking_set_id"`
	Labplus             int     `json:"labplus"`
	Category
	CheckingListProduct *[]Checking_product `json:"checking_product" gorm:"foreignKey:Id;references:Category_id"`
	CheckingList        *[]CheckingSubList  `json:"checking_list" gorm:"foreignKey:Id;references:Category_id"`
}

type ObjPayloadAddProduct struct {
	Shop_id             int                     `json:"shop_id"`
	Category_id         int                     `json:"category_id" binding:"required"`
	Checking_type_id    int                     `json:"checking_type_id" binding:"required"`
	Checking_code       string                  `json:"checking_code" binding:"required"`
	Checking_code_acc   *string                 `json:"checking_code_acc" binding:"required,omitempty"`
	Checking_name_acc   *string                 `json:"checking_name_acc" binding:"required,omitempty"`
	Checking_name       string                  `json:"checking_name" binding:"required"`
	Checking_unit       string                  `json:"checking_unit" binding:"required"`
	Checking_opd        *float64                `json:"checking_opd" binding:"required,omitempty"`
	Checking_ipd        *float64                `json:"checking_ipd" binding:"required,omitempty"`
	Category_eclaim_id  *int                    `json:"category_eclaim_id" binding:"required,omitempty"`
	Checking_ofc        *float64                `json:"checking_ofc" binding:"required,omitempty"`
	Checking_lgo        *float64                `json:"checking_lgo" binding:"required,omitempty"`
	Checking_ucs        *float64                `json:"checking_ucs" binding:"required,omitempty"`
	Checking_sss        *float64                `json:"checking_sss" binding:"required,omitempty"`
	Checking_nhs        *float64                `json:"checking_nhs" binding:"required,omitempty"`
	Checking_ssi        *float64                `json:"checking_ssi" binding:"required,omitempty"`
	Checking_cost       *float64                `json:"checking_cost" binding:"required,omitempty"`
	Checking_fee_df     *float64                `json:"checking_fee_df" binding:"required,omitempty"`
	Checking_fee_nr     *float64                `json:"checking_fee_nr" binding:"required,omitempty"`
	Checking_fee_tr     *float64                `json:"checking_fee_tr" binding:"required,omitempty"`
	Checking_fee        *float64                `json:"checking_fee" binding:"required,omitempty"`
	Acc_code_id_cost    *int                    `json:"acc_code_id_cost" binding:"required,omitempty"`
	Acc_code_id_fee_df  *int                    `json:"acc_code_id_fee_df" binding:"required,omitempty"`
	Acc_code_id_fee     *int                    `json:"acc_code_id_fee" binding:"required,omitempty"`
	Acc_code_id_com     *int                    `json:"acc_code_id_com" binding:"required,omitempty"`
	Checking_image      *string                 `json:"checking_image" binding:"required,omitempty"`
	Checking_is_active  int                     `json:"checking_is_active"`
	Checking_is_del     int                     `json:"checking_is_del"`
	Checking_is_labplus int                     `json:"checking_is_labplus"`
	Checking_create     string                  `json:"checking_create"`
	Checking_update     string                  `json:"checking_update"`
	Checking_product    *[]Obj_Checking_product `json:"checking_product" binding:"required,omitempty"`
}

type ObjPayloadAddList struct {
	Shop_id             int                  `json:"shop_id"`
	Category_id         int                  `json:"category_id" binding:"required"`
	Checking_type_id    int                  `json:"checking_type_id" binding:"required"`
	Checking_code       string               `json:"checking_code" binding:"required"`
	Checking_code_acc   *string              `json:"checking_code_acc" binding:"required,omitempty"`
	Checking_name_acc   *string              `json:"checking_name_acc" binding:"required,omitempty"`
	Checking_name       string               `json:"checking_name" binding:"required"`
	Checking_unit       string               `json:"checking_unit" binding:"required"`
	Checking_opd        *float64             `json:"checking_opd" binding:"required,omitempty"`
	Checking_ipd        *float64             `json:"checking_ipd" binding:"required,omitempty"`
	Category_eclaim_id  *int                 `json:"category_eclaim_id" binding:"required,omitempty"`
	Checking_ofc        *float64             `json:"checking_ofc" binding:"required,omitempty"`
	Checking_lgo        *float64             `json:"checking_lgo" binding:"required,omitempty"`
	Checking_ucs        *float64             `json:"checking_ucs" binding:"required,omitempty"`
	Checking_sss        *float64             `json:"checking_sss" binding:"required,omitempty"`
	Checking_nhs        *float64             `json:"checking_nhs" binding:"required,omitempty"`
	Checking_ssi        *float64             `json:"checking_ssi" binding:"required,omitempty"`
	Checking_cost       *float64             `json:"checking_cost" binding:"required,omitempty"`
	Checking_fee_df     *float64             `json:"checking_fee_df" binding:"required,omitempty"`
	Checking_fee_nr     *float64             `json:"checking_fee_nr" binding:"required,omitempty"`
	Checking_fee_tr     *float64             `json:"checking_fee_tr" binding:"required,omitempty"`
	Checking_fee        *float64             `json:"checking_fee" binding:"required,omitempty"`
	Acc_code_id_cost    *int                 `json:"acc_code_id_cost" binding:"required,omitempty"`
	Acc_code_id_fee_df  *int                 `json:"acc_code_id_fee_df" binding:"required,omitempty"`
	Acc_code_id_fee     *int                 `json:"acc_code_id_fee" binding:"required,omitempty"`
	Acc_code_id_com     *int                 `json:"acc_code_id_com" binding:"required,omitempty"`
	Checking_image      *string              `json:"checking_image" binding:"required,omitempty"`
	Checking_is_active  int                  `json:"checking_is_active"`
	Checking_is_del     int                  `json:"checking_is_del"`
	Checking_is_labplus int                  `json:"checking_is_labplus"`
	Checking_create     string               `json:"checking_create"`
	Checking_update     string               `json:"checking_update"`
	Checking_List       *[]Obj_Checking_list `json:"checking_list" binding:"required,omitempty"`
}

type ObjPayloadEditProduct struct {
	Id                  int                     `json:"id" binding:"required"`
	Shop_id             int                     `json:"shop_id"`
	Category_id         int                     `json:"category_id" binding:"required"`
	Checking_type_id    int                     `json:"checking_type_id" binding:"required"`
	Checking_code       string                  `json:"checking_code" binding:"required"`
	Checking_code_acc   *string                 `json:"checking_code_acc" binding:"required,omitempty"`
	Checking_name_acc   *string                 `json:"checking_name_acc" binding:"required,omitempty"`
	Checking_name       string                  `json:"checking_name" binding:"required"`
	Checking_unit       string                  `json:"checking_unit" binding:"required"`
	Checking_opd        *float64                `json:"checking_opd" binding:"required,omitempty"`
	Checking_ipd        *float64                `json:"checking_ipd" binding:"required,omitempty"`
	Category_eclaim_id  *int                    `json:"category_eclaim_id" binding:"required,omitempty"`
	Checking_ofc        *float64                `json:"checking_ofc" binding:"required,omitempty"`
	Checking_lgo        *float64                `json:"checking_lgo" binding:"required,omitempty"`
	Checking_ucs        *float64                `json:"checking_ucs" binding:"required,omitempty"`
	Checking_sss        *float64                `json:"checking_sss" binding:"required,omitempty"`
	Checking_nhs        *float64                `json:"checking_nhs" binding:"required,omitempty"`
	Checking_ssi        *float64                `json:"checking_ssi" binding:"required,omitempty"`
	Checking_cost       *float64                `json:"checking_cost" binding:"required,omitempty"`
	Checking_fee_df     *float64                `json:"checking_fee_df" binding:"required,omitempty"`
	Checking_fee_nr     *float64                `json:"checking_fee_nr" binding:"required,omitempty"`
	Checking_fee_tr     *float64                `json:"checking_fee_tr" binding:"required,omitempty"`
	Checking_fee        *float64                `json:"checking_fee" binding:"required,omitempty"`
	Acc_code_id_cost    *int                    `json:"acc_code_id_cost" binding:"required,omitempty"`
	Acc_code_id_fee_df  *int                    `json:"acc_code_id_fee_df" binding:"required,omitempty"`
	Acc_code_id_fee     *int                    `json:"acc_code_id_fee" binding:"required,omitempty"`
	Acc_code_id_com     *int                    `json:"acc_code_id_com" binding:"required,omitempty"`
	Checking_image      *string                 `json:"checking_image" binding:"required,omitempty"`
	Checking_is_active  *int                    `json:"checking_is_active" binding:"omitempty"`
	Checking_is_del     int                     `json:"checking_is_del"`
	Checking_is_labplus int                     `json:"checking_is_labplus"`
	Checking_update     string                  `json:"checking_update"`
	Checking_product    *[]Obj_Checking_product `json:"checking_product" binding:"required,omitempty"`
}

type ObjPayloadEditList struct {
	Id                  int                  `json:"id" binding:"required"`
	Shop_id             int                  `json:"shop_id"`
	Category_id         int                  `json:"category_id" binding:"required"`
	Checking_type_id    int                  `json:"checking_type_id" binding:"required"`
	Checking_code       string               `json:"checking_code" binding:"required"`
	Checking_code_acc   *string              `json:"checking_code_acc" binding:"required,omitempty"`
	Checking_name_acc   *string              `json:"checking_name_acc" binding:"required,omitempty"`
	Checking_name       string               `json:"checking_name" binding:"required"`
	Checking_unit       string               `json:"checking_unit" binding:"required"`
	Checking_opd        *float64             `json:"checking_opd" binding:"required,omitempty"`
	Checking_ipd        *float64             `json:"checking_ipd" binding:"required,omitempty"`
	Category_eclaim_id  *int                 `json:"category_eclaim_id" binding:"required,omitempty"`
	Checking_ofc        *float64             `json:"checking_ofc" binding:"required,omitempty"`
	Checking_lgo        *float64             `json:"checking_lgo" binding:"required,omitempty"`
	Checking_ucs        *float64             `json:"checking_ucs" binding:"required,omitempty"`
	Checking_sss        *float64             `json:"checking_sss" binding:"required,omitempty"`
	Checking_nhs        *float64             `json:"checking_nhs" binding:"required,omitempty"`
	Checking_ssi        *float64             `json:"checking_ssi" binding:"required,omitempty"`
	Checking_cost       *float64             `json:"checking_cost" binding:"required,omitempty"`
	Checking_fee_df     *float64             `json:"checking_fee_df" binding:"required,omitempty"`
	Checking_fee_nr     *float64             `json:"checking_fee_nr" binding:"required,omitempty"`
	Checking_fee_tr     *float64             `json:"checking_fee_tr" binding:"required,omitempty"`
	Checking_fee        *float64             `json:"checking_fee" binding:"required,omitempty"`
	Acc_code_id_cost    *int                 `json:"acc_code_id_cost" binding:"required,omitempty"`
	Acc_code_id_fee_df  *int                 `json:"acc_code_id_fee_df" binding:"required,omitempty"`
	Acc_code_id_fee     *int                 `json:"acc_code_id_fee" binding:"required,omitempty"`
	Acc_code_id_com     *int                 `json:"acc_code_id_com" binding:"required,omitempty"`
	Checking_image      *string              `json:"checking_image" binding:"required,omitempty"`
	Checking_is_active  *int                 `json:"checking_is_active" binding:"omitempty"`
	Checking_is_del     int                  `json:"checking_is_del"`
	Checking_create     string               `json:"checking_create"`
	Checking_update     string               `json:"checking_update"`
	Checking_is_labplus int                  `json:"checking_is_labplus"`
	Checking_set_id     int                  `json:"checking_set_id"`
	Checking_List       *[]Obj_Checking_list `json:"checking_list" binding:"required,omitempty"`
}

type CheckingAction struct {
	Id                  int     `json:"id"`
	Shop_id             int     `json:"shop_id"`
	Category_id         int     `json:"category_id"`
	Checking_type_id    int     `json:"checking_type_id"`
	Checking_code       string  `json:"checking_code"`
	Checking_code_acc   string  `json:"checking_code_acc"`
	Checking_name_acc   string  `json:"checking_name_acc"`
	Checking_name       string  `json:"checking_name"`
	Checking_unit       string  `json:"checking_unit"`
	Checking_opd        float64 `json:"checking_opd"`
	Checking_ipd        float64 `json:"checking_ipd"`
	Category_eclaim_id  *int    `json:"category_eclaim_id"`
	Checking_ofc        float64 `json:"checking_ofc"`
	Checking_lgo        float64 `json:"checking_lgo"`
	Checking_ucs        float64 `json:"checking_ucs"`
	Checking_sss        float64 `json:"checking_sss"`
	Checking_nhs        float64 `json:"checking_nhs"`
	Checking_ssi        float64 `json:"checking_ssi"`
	Checking_cost       float64 `json:"checking_cost"`
	Checking_fee_df     float64 `json:"checking_fee_df"`
	Checking_fee_nr     float64 `json:"checking_fee_nr"`
	Checking_fee_tr     float64 `json:"checking_fee_tr"`
	Checking_fee        float64 `json:"checking_fee"`
	Acc_code_id_cost    int     `json:"acc_code_id_cost" binding:"required,omitempty"`
	Acc_code_id_fee_df  int     `json:"acc_code_id_fee_df" binding:"required,omitempty"`
	Acc_code_id_fee     int     `json:"acc_code_id_fee" binding:"required,omitempty"`
	Acc_code_id_com     int     `json:"acc_code_id_com" binding:"required,omitempty"`
	Checking_image      string  `json:"checking_image" binding:"required,omitempty"`
	Checking_is_active  int     `json:"checking_is_active"`
	Checking_is_del     int     `json:"checking_is_del"`
	Checking_is_labplus int     `json:"checking_is_labplus"`
	Checking_create     string  `json:"checking_create"`
	Checking_update     string  `json:"checking_update"`
}

type Obj_Checking_product struct {
	Id            int     `json:"id"`
	Checking_id   int     `json:"checking_id"`
	Product_id    int     `json:"product_id" binding:"required"`
	Cip_amount    float64 `json:"cip_amount"`
	Cip_is_active int     `json:"cip_is_active"`
	Cip_is_del    int     `json:"cip_is_del"`
}

type Obj_Checking_list struct {
	Id                int     `json:"id"`
	Checking_set_id   int     `json:"checking_set_id"`
	Checking_id       int     `json:"checking_id"`
	Checking_list_opd float64 `json:"checking_list_opd"`
	Checking_list_ipd float64 `json:"checking_list_ipd"`
}

type Obj_Checking_set struct {
	Id          int `json:"id"`
	Checking_id int `json:"checking_id"`
}

type Checking_product struct {
	Id            int     `json:"id"`
	Checking_id   int     `json:"checking_id"`
	Product_id    int     `json:"product_id"`
	Cip_amount    float64 `json:"cip_amount"`
	Cip_is_active int     `json:"cip_is_active"`
	Cip_is_del    int     `json:"cip_is_del"`
	Product
}

type CheckingSubList struct {
	Id                int     `json:"id"`
	Checking_set_id   int     `json:"checking_set_id"`
	Checking_id       int     `json:"checking_id"`
	Checking_list_opd float64 `json:"checking_list_opd"`
	Checking_list_ipd float64 `json:"checking_list_ipd"`
	CheckingList      `gorm:"foreignKey:checking_id;references:Id"`
}

type CheckingSubListId struct {
	Checking_id int `json:"checking_id"`
}

type Product struct {
	Id          int    `json:"product_id"`
	Shop_id     int    `json:"shop_id"`
	Category_id int    `json:"category_id"`
	Pd_type_id  int    `json:"pd_type_id"`
	Pd_code     string `json:"pd_code"`
	Pd_name     string `json:"pd_name"`
	Pd_code_acc string `json:"pd_code_acc"`
	Pd_name_acc string `json:"pd_name_acc"`
	Product_unit
}

type Product_unit struct {
	PuId    int `json:"pu_id" db:"id"`
	Unit_id int `json:"unit_id"`
	Ref_unit
	Pu_rate  int     `json:"pu_rate"`
	Pu_price float64 `json:"pu_price"`
}

type Ref_unit struct {
	Uid    int    `json:"u_id" db:"id"`
	U_name string `json:"u_name"`
}

type DocChecking struct {
	ShopId               int    `json:"shop_id"`
	Check_id_default     string `json:"check_id_default"`
	Check_number_default string `json:"check_number_default"`
	Check_number_digit   int    `json:"check_number_digit"`
	Check_type           int    `json:"check_type"`
	Lab_id_default       string `json:"lab_id_default"`
	Lab_number_default   string `json:"lab_number_default"`
	Lab_number_digit     int    `json:"lab_number_digit"`
	Lab_type             int    `json:"lab_type"`
	Xray_id_default      string `json:"xray_id_default"`
	Xray_number_default  string `json:"xray_number_default"`
	Xray_number_digit    int    `json:"xray_number_digit"`
	Xray_type            int    `json:"xray_type"`
}

type LogChecking struct {
	Username   string `json:"username"`
	Log_type   string `json:"log_type"`
	Log_text   string `json:"log_text"`
	Log_create string `json:"log_create"`
}

// import
type ObjQueryChecking struct {
	ID               int     `json:"id"`
	ShopId           int     `json:"shop_id"`
	CategoryId       int     `json:"category_id"`
	CheckingTypeId   int     `json:"checking_type_id"`
	CheckingCode     string  `json:"checking_code"`
	CheckingCodeAcc  string  `json:"checking_code_acc"`
	CheckingNameAcc  string  `json:"checking_name_acc"`
	CheckingName     string  `json:"checking_name"`
	CheckingUnit     string  `json:"checking_unit"`
	CheckingOpd      float64 `json:"checking_opd"`
	CheckingIpd      float64 `json:"checking_ipd"`
	CheckingCost     float64 `json:"checking_cost"`
	CheckingFeeDf    float64 `json:"checking_fee_df"`
	CheckingFeeNr    float64 `json:"checking_fee_nr"`
	CheckingFeeTr    float64 `json:"checking_fee_tr"`
	CheckingFee      float64 `json:"checking_fee"`
	AccCodeIdCost    float64 `json:"acc_code_id_cost"`
	AccCodeIdFeeDf   float64 `json:"acc_code_id_fee_df"`
	AccCodeIdFee     int     `json:"acc_code_id_fee"`
	AccCodeIdCom     int     `json:"acc_code_id_com"`
	CheckingImage    string  `json:"checking_image"`
	CheckingIsActive int     `json:"checking_is_active"`
	CheckingIsDel    int     `json:"checking_is_del"`
	CheckingCreate   string  `json:"checking_create"`
	CheckingUpdate   string  `json:"checking_update"`
}

type ObjCheckExcelChecking struct {
	ShopId            int      `json:"shop_id"`
	CategoryId        int      `json:"category_id"`
	CheckingTypeId    int      `json:"checking_type_id"`
	CheckingCode      string   `json:"checking_code"`
	CheckingCodeAcc   string   `json:"checking_code_acc"`
	CheckingNameAcc   string   `json:"checking_name_acc"`
	CheckingName      string   `json:"checking_name"`
	CheckingUnit      string   `json:"checking_unit"`
	CheckingOpd       float64  `json:"checking_opd"`
	CheckingIpd       float64  `json:"checking_ipd"`
	CheckingCost      float64  `json:"checking_cost"`
	CheckingFeeDf     float64  `json:"checking_fee_df"`
	CheckingFeeNr     float64  `json:"checking_fee_nr"`
	CheckingFeeTr     float64  `json:"checking_fee_tr"`
	CheckingFee       float64  `json:"checking_fee"`
	CheckingIsLabplus int      `json:"checking_is_labplus"`
	Message           []string `json:"message"`
}

type ObjPayloadImportExcelChecking struct {
	ImportData []ObjCheckExcelChecking `json:"import_data"`
}
