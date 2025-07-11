package structs

type ObjPayloadGetOPDPagination struct {
	Search_text *string `json:"search_text"`
	Search_date *string `json:"search_date"`
	User_id     *int    `json:"user_id"`
	Shop_id     int     `json:"shop_id"`
	ActivePage  int     `json:"active_page" binding:"required"`
	PerPage     int     `json:"per_page" binding:"required"`
}

type ObjQueryGetOPDPagination struct {
	ID               int     `json:"id"`
	UserId           int     `json:"user_id"`
	QueueId          int     `json:"queue_id"`
	QueCode          string  `json:"que_code"`
	CustomerId       int     `json:"customer_id"`
	CtmId            string  `json:"ctm_id"`
	CtmFname         string  `json:"ctm_fname"`
	CtmLname         string  `json:"ctm_lname"`
	OpdCode          string  `json:"opd_code"`
	OpdDate          string  `json:"opd_date"`
	OpdBw            float64 `json:"opd_bw"`
	OpdHt            float64 `json:"opd_ht"`
	OpdBmi           float64 `json:"opd_bmi"`
	OpdT             string  `json:"opd_t"`
	OpdBsa           string  `json:"opd_bsa"`
	OpdVas           string  `json:"opd_vas"`
	OpdPr            string  `json:"opd_pr"`
	OpdBp            string  `json:"opd_bp"`
	OpdRr            string  `json:"opd_rr"`
	OpdSys           string  `json:"opd_sys"`
	OpdDia           string  `json:"opd_dia"`
	OpdO2            string  `json:"opd_o2"`
	OpdFag           int     `json:"opd_fag"`
	OpdAlcohol       int     `json:"opd_alcohol"`
	OpdCc            string  `json:"opd_cc"`
	OpdHpi           string  `json:"opd_hpi"`
	OpdPmh           string  `json:"opd_pmh"`
	OpdDx            string  `json:"opd_dx"`
	OpdIopatLe       string  `json:"opd_iopat_le"`
	OpdVascRe        string  `json:"opd_vasc_re"`
	OpdVascLe        string  `json:"opd_vasc_le"`
	OpdVaccRe        string  `json:"opd_vacc_re"`
	OpdVaccLe        string  `json:"opd_vacc_le"`
	OpdIopatRe       string  `json:"opd_iopat_re"`
	OpdGa            string  `json:"opd_ga"`
	OpdPe            string  `json:"opd_pe"`
	OpdNote          string  `json:"opd_note"`
	OpdIsData        *int    `json:"opd_is_data"`
	OpdSickStartdate *string `json:"opd_sick_startdate"`
	OpdSickEnddate   *string `json:"opd_sick_enddate"`
	OpdSickNotrest   int     `json:"opd_sick_notrest"`
	OpdSickAir       int     `json:"opd_sick_air"`
	UserFullname     string  `json:"user_fullname"`
	UserFullnameEn   string  `json:"user_fullname_en"`
	// OpdCustom        []ObjPayloadOpdCustom     `json:"opd_custom"`
	// OpdDiagnostic    []ObjPayloadOpdDiagnostic `json:"opd_diagnostic"`
}

type OpdCustom struct {
	OpdcName  string `json:"opdc_name"`
	OpdcValue string `json:"opdc_value"`
}

type OpdDiagnostic struct {
	DiagnosticId   int    `json:"diagnostic_id"`
	DiagnosticCode string `json:"diagnostic_code"`
	DiagnosticTh   string `json:"diagnostic_th"`
}

type ObjResponseGetOPDPagination struct {
	Result_data   []ObjQueryGetOPDPagination `json:"result_data"`
	Count_of_page int                        `json:"count_of_page"`
	Count_all     int                        `json:"count_all"`
}

type OpdDetail struct {
	ID               int              `json:"id"`
	UserId           int              `json:"user_id"`
	QueueId          int              `json:"queue_id"`
	QueCode          string           `json:"que_code"`
	CustomerId       int              `json:"customer_id"`
	CtmId            string           `json:"ctm_id"`
	CtmFname         string           `json:"ctm_fname"`
	CtmLname         string           `json:"ctm_lname"`
	OpdCode          string           `json:"opd_code"`
	OpdDate          string           `json:"opd_date"`
	OpdBw            float64          `json:"opd_bw"`
	OpdHt            float64          `json:"opd_ht"`
	OpdBmi           float64          `json:"opd_bmi"`
	OpdT             string           `json:"opd_t"`
	OpdBsa           string           `json:"opd_bsa"`
	OpdVas           string           `json:"opd_vas"`
	OpdPr            string           `json:"opd_pr"`
	OpdBp            string           `json:"opd_bp"`
	OpdRr            string           `json:"opd_rr"`
	OpdSys           string           `json:"opd_sys"`
	OpdDia           string           `json:"opd_dia"`
	OpdO2            string           `json:"opd_o2"`
	OpdFag           int              `json:"opd_fag"`
	OpdAlcohol       int              `json:"opd_alcohol"`
	OpdCc            string           `json:"opd_cc"`
	OpdHpi           string           `json:"opd_hpi"`
	OpdPmh           string           `json:"opd_pmh"`
	OpdDx            string           `json:"opd_dx"`
	OpdIopatLe       string           `json:"opd_iopat_le"`
	OpdVascRe        string           `json:"opd_vasc_re"`
	OpdVascLe        string           `json:"opd_vasc_le"`
	OpdVaccRe        string           `json:"opd_vacc_re"`
	OpdVaccLe        string           `json:"opd_vacc_le"`
	OpdIopatRe       string           `json:"opd_iopat_re"`
	OpdGa            string           `json:"opd_ga"`
	OpdPe            string           `json:"opd_pe"`
	OpdNote          string           `json:"opd_note"`
	OpdIsData        *int             `json:"opd_is_data"`
	OpdSickStartdate *string          `json:"opd_sick_startdate"`
	OpdSickEnddate   *string          `json:"opd_sick_enddate"`
	OpdSickNotrest   int              `json:"opd_sick_notrest"`
	OpdSickAir       int              `json:"opd_sick_air"`
	UserFullname     string           `json:"user_fullname"`
	OpdCustoms       []OpdCustom      `json:"opd_customs" gorm:"-"`
	OpdDiagnostics   []OpdDiagnostic  `json:"opd_diagnostics" gorm:"-"`
	Shop             ReceiptShop      `json:"shop" gorm:"-"`
	Customer         ObjQueryCustomer `json:"customer" gorm:"-"`
	User_fullname_en string           `json:"user_fullname_en"`
	User_license     *string          `json:"user_license"`
	Checkings        []CheckingOPD    `json:"checkings" gorm:"-"`
	Courses          []CourseOPD      `json:"courses" gorm:"-"`
	Products         []ProductOPD     `json:"products" gorm:"-"`
	File_first       string           `json:"file_first" gorm:"-"`
	File_last        string           `json:"file_last" gorm:"-"`
}

type CheckingOPD struct {
	Ord_code string  `json:"ord_code"`
	Ord_name string  `json:"ord_name"`
	Ord_qty  float64 `json:"ord_qty"`
	Ord_unit string  `json:"ord_unit"`
}

type CourseOPD struct {
	Ord_code string  `json:"ord_code"`
	Ord_name string  `json:"ord_name"`
	Ord_qty  float64 `json:"ord_qty"`
	Ord_unit string  `json:"ord_unit"`
}

type ProductOPD struct {
	Ord_code string  `json:"ord_code"`
	Ord_name string  `json:"ord_name"`
	Ord_qty  float64 `json:"ord_qty"`
	Ord_unit string  `json:"ord_unit"`
}

type Queue_files struct {
	Quef_path   string `json:"quef_path"`
	Quef_modify string `json:"quef_modify"`
}
