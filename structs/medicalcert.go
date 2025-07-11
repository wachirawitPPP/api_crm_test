package structs

type ObjPayloadGetMedicalCertPagination struct {
	ShopId      int     `json:"shop_id" binding:"required"`
	UserId      *int    `json:"user_id" binding:"required"`
	MdctId      *int    `json:"mdct_id" binding:"required"`
	SearchDate  *string `json:"search_date" binding:"required"`
	SearchText  *string `json:"search_text" binding:"required"`
	CurrentPage int     `json:"current_page" binding:"required"`
	PerPage     int     `json:"per_page" binding:"required"`
}

type ObjQueryGetMedicalCertPagination struct {
	ID                int    `json:"id"`
	MedicalCertTypeId int    `json:"medical_cert_type_id"`
	UserId            int    `json:"user_id"`
	OpdId             int    `json:"opd_id"`
	MdcCode           string `json:"mdc_code"`
	MdcIsPrint        int    `json:"mdc_is_print"`
	MdcIsDel          int    `json:"mdc_is_del"`
	MdcCreate         string `json:"mdc_create"`
	MdcUpdate         string `json:"mdc_update"`
	// join
	CtmId          string `json:"ctm_id"`
	CtmPrefix      string `json:"ctm_prefix"`
	CtmFname       string `json:"ctm_fname"`
	CtmLname       string `json:"ctm_lname"`
	CtmNname       string `json:"ctm_nname"`
	CtmFnameEn     string `json:"ctm_fname_en"`
	CtmLnameEn     string `json:"ctm_lname_en"`
	UserFullname   string `json:"user_fullname"`
	UserFullnameEn string `json:"user_fullname_en"`
	MdctTh         string `json:"mdct_th"`
	MdctEn         string `json:"mdct_en"`
}

type ObjResponseGetMedicalCertPagination struct {
	Items     []ObjQueryGetMedicalCertPagination `json:"items"`
	CountPage int                                `json:"count_page"`
	CountAll  int64                              `json:"count_all"`
}

type MedicalCertDetail struct {
	Shop_id              int              `json:"shop_id"`
	Customer_id          int              `json:"customer_id"`
	Mdct_th              string           `json:"mdct_th"`
	Mdct_en              string           `json:"mdct_en"`
	Mdct_group_id        int              `json:"mdct_group_id"`
	Id                   int              `json:"id"`
	Medical_cert_type_id int              `json:"medical_cert_type_id"`
	User_id              int              `json:"user_id"`
	User_fullname        string           `json:"user_fullname"`
	User_fullname_en     string           `json:"user_fullname_en"`
	User_license         *string          `json:"user_license"`
	Opd_id               int              `json:"opd_id"`
	Mdc_code             string           `json:"mdc_code"`
	Mdc_is_print         int              `json:"mdc_is_print"`
	Mdc_create           string           `json:"mdc_create"`
	Mdc_update           string           `json:"mdc_update"`
	Queue_id             int              `json:"queue_id"`
	Opd_code             string           `json:"opd_code"`
	Opd_date             string           `json:"opd_date"`
	Opd_bw               *float64         `json:"opd_bw"`
	Opd_ht               *float64         `json:"opd_ht"`
	Opd_bmi              *float64         `json:"opd_bmi"`
	Opd_t                string           `json:"opd_t"`
	Opd_bsa              string           `json:"opd_bsa"`
	Opd_vas              string           `json:"opd_vas"`
	Opd_pr               string           `json:"opd_pr"`
	Opd_bp               string           `json:"opd_bp"`
	Opd_rr               string           `json:"opd_rr"`
	Opd_sys              string           `json:"opd_sys"`
	Opd_dia              string           `json:"opd_dia"`
	Opd_o2               string           `json:"opd_o2"`
	Opd_fag              *int             `json:"opd_fag"`
	Opd_alcohol          *int             `json:"opd_alcohol"`
	Opd_cc               string           `json:"opd_cc"`
	Opd_hpi              string           `json:"opd_hpi"`
	Opd_pmh              string           `json:"opd_pmh"`
	Opd_dx               string           `json:"opd_dx"`
	Opd_iopat_le         string           `json:"opd_iopat_le"`
	Opd_vasc_re          string           `json:"opd_vasc_re"`
	Opd_vasc_le          string           `json:"opd_vasc_le"`
	Opd_vacc_re          string           `json:"opd_vacc_re"`
	Opd_vacc_le          string           `json:"opd_vacc_le"`
	Opd_iopat_re         string           `json:"opd_iopat_re"`
	Opd_ga               string           `json:"opd_ga"`
	Opd_pe               string           `json:"opd_pe"`
	Opd_note             string           `json:"opd_note"`
	Opd_sick_startdate   string           `json:"opd_sick_startdate"`
	Opd_sick_enddate     string           `json:"opd_sick_enddate"`
	Opd_sick_notrest     *int             `json:"opd_sick_notrest"`
	Opd_sick_air         *int             `json:"opd_sick_air"`
	Opd_is_data          string           `json:"opd_is_data"`
	Opd_update           string           `json:"opd_update"`
	Opd_create           string           `json:"opd_create"`
	Diagnostic_code      *string          `json:"diagnostic_code"`
	Diagnostic_th        *string          `json:"diagnostic_th"`
	Diagnostic_en        *string          `json:"diagnostic_en"`
	Diagnostic_detail    *string          `json:"diagnostic_detail"`
	Opdc_name            *string          `json:"opdc_name"`
	Opdc_value           *int             `json:"opdc_value"`
	Shop                 ReceiptShop      `json:"shop" gorm:"-"`
	Customer             ObjQueryCustomer `json:"customer" gorm:"-"`
}
