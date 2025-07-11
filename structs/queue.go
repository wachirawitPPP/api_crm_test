package structs

type ObjPayloadGetQueuePagination struct {
	ShopId      int     `json:"shop_id" binding:"required"`
	QueAdmisId  *int    `json:"que_admis_id" binding:"required"`
	DoctorId    *int    `json:"doctor_id" binding:"required"`
	QueTypeId   *int    `json:"que_type_id" binding:"required"`
	QueStatusId *int    `json:"que_status_id" binding:"required"`
	QueDate     *string `json:"que_date" binding:"required"`
	SearchText  *string `json:"search_text" binding:"required"`
	CurrentPage int     `json:"current_page" binding:"required"`
	PerPage     int     `json:"per_page" binding:"required"`
}

type ObjQueryGetQueuePagination struct {
	ID              int    `json:"id"`
	ShopId          int    `json:"shop_id"`
	CustomerId      int    `json:"customer_id"`
	RoomId          *int   `json:"room_id"`
	BedId           *int   `json:"bed_id"`
	UserId          int    `json:"user_id"`
	QueUserId       int    `json:"que_user_id"`
	QueUserFullname string `json:"que_user_fullname"`
	QueCode         string `json:"que_code"`
	QueTypeId       int    `json:"que_type_id"`
	QueAdmisId      int    `json:"que_admis_id"`
	QuePriorityId   int    `json:"que_priority_id"`
	QueStatusId     int    `json:"que_status_id"`
	QueDatetime     string `json:"que_datetime"`
	QueDatetimeOut  string `json:"que_datetime_out"`
	QueCreate       string `json:"que_create"`
	QueUpdate       string `json:"que_update"`
	CtmId           string `json:"ctm_id"`
	CtmImage        string `json:"ctm_image"`
	CtmPrefix       string `json:"ctm_prefix"`
	CtmGender       string `json:"ctm_gender"`
	CtmFname        string `json:"ctm_fname"`
	CtmLname        string `json:"ctm_lname"`
	CtmFnameEn      string `json:"ctm_fname_en"`
	CtmLnameEn      string `json:"ctm_lname_en"`
	CtmShopId       int    `json:"ctm_shop_id"`
	CtmShopName     string `json:"ctm_shop_name"`
	RoomCode        string `json:"room_code"`
	RoomTH          string `json:"room_th"`
	RoomEn          string `json:"room_en"`
	BedCode         string `json:"bed_code"`
	InvoiceId       int    `json:"invoice_id"`
	QueTimeEnd      int    `json:"que_time_end"`
	QueHours        int    `json:"que_hours"`
	QueMin          int    `json:"que_min"`
	LabXray         int    `json:"lab_xray"`
}

type ObjPayloadGetShopAddon struct {
	ShopId    int `json:"shop_id" binding:"required"`
	PackageId int `json:"package_id" binding:"required"`
}

type GetQueueCheckLabXray struct {
	ID int `json:"id"`
}

type ObjQueryGetQueue struct {
	ID                  int     `json:"id"`
	ShopId              int     `json:"shop_id"`
	CustomerId          int     `json:"customer_id"`
	RoomId              *int    `json:"room_id"`
	BedId               *int    `json:"bed_id"`
	UserId              int     `json:"user_id"`
	QueUserId           int     `json:"que_user_id"`
	DpmId               int     `json:"dpm_id"`
	QueUserFullname     string  `json:"que_user_fullname"`
	CtmFullname         string  `json:"ctm_fullname"`
	CtmFullnameEn       string  `json:"ctm_fullname_en"`
	QueCode             string  `json:"que_code"`
	QueTypeId           int     `json:"que_type_id"`
	QueAdmisId          int     `json:"que_admis_id"`
	QuePriorityId       int     `json:"que_priority_id"`
	QueStatusId         int     `json:"que_status_id"`
	QueComment          string  `json:"que_comment"`
	QueNote             string  `json:"que_note"`
	QueDatetimeOut      string  `json:"que_datetime_out"`
	QueRoomTotal        string  `json:"que_room_total"`
	QueDatetime         string  `json:"que_datetime"`
	QueTeleCode         string  `json:"que_tele_code"`
	QueTeleUrl          string  `json:"que_tele_url"`
	QueCreate           string  `json:"que_create"`
	QueUpdate           string  `json:"que_update"`
	CtmId               string  `json:"ctm_id"`
	CtmImage            string  `json:"ctm_image"`
	CtmGender           string  `json:"ctm_gender"`
	CtmPrefix           string  `json:"ctm_prefix"`
	CtmFname            string  `json:"ctm_fname"`
	CtmLname            string  `json:"ctm_lname"`
	CtmFnameEn          string  `json:"ctm_fname_en"`
	CtmLnameEn          string  `json:"ctm_lname_en"`
	CtmBirthdate        string  `json:"ctm_birthdate"`
	CtmBlood            string  `json:"ctm_blood"`
	CtmWeight           float64 `json:"ctm_weight" gorm:"type:decimal(10,2)"`
	CtmHeight           float64 `json:"ctm_height" gorm:"type:decimal(10,2)"`
	CtmWaistline        float64 `json:"ctm_waistline" gorm:"type:decimal(10,2)"`
	CtmChest            float64 `json:"ctm_chest" gorm:"type:decimal(10,2)"`
	CtmTreatmentType    int     `json:"ctm_treatment_type"`
	CtmHealthComment    string  `json:"ctm_health_comment"`
	CtmDisease          string  `json:"ctm_disease"`
	CtmMentalHealth     string  `json:"ctm_mental_health"`
	CtmAllergic         string  `json:"ctm_allergic"`
	UserFullname        string  `json:"user_fullname"`
	RtCode              string  `json:"rt_code"`
	RtName              string  `json:"rt_name"`
	RoomCode            string  `json:"room_code"`
	RoomTH              string  `json:"room_th"`
	RoomEn              string  `json:"room_en"`
	BedCode             string  `json:"bed_code"`
	Directions_id       int     `json:"directions_id"`
	Que_directions      string  `json:"que_directions"`
	Que_directions_name string  `json:"que_directions_name"`
	Qlp_id              int     `json:"qlp_id"`
	Qlp_no              string  `json:"qlp_no"`
	Qlp_message_th      string  `json:"qlp_message_th"`
	Qlp_process_code    string  `json:"qlp_process_code"`
	Qlp_process_name    string  `json:"qlp_process_name"`
	Qlp_datetime        string  `json:"qlp_datetime"`
	Qlp_update          string  `json:"qlp_update"`
}

type PayloadQueueDirection struct {
	Id                  int     `json:"id"`
	Directions_id       int     `json:"directions_id" binding:"required"`
	Que_directions_name *string `json:"que_directions_name" binding:"required,omitempty"`
	Que_directions      *string `json:"que_directions" binding:"required,omitempty"`
}

type ObjResponseGetQueuePagination struct {
	Items     []ObjQueryGetQueuePagination `json:"items"`
	CountPage int                          `json:"count_page"`
	CountAll  int64                        `json:"count_all"`
}

type ObjQueryQueueLabel struct {
	ID         int `json:"id"`
	ShopId     int `json:"shop_id"`
	CustomerId int `json:"customer_id"`
	// RoomId           *int    `json:"room_id"`
	// BedId            *int    `json:"bed_id"`
	UserId    int `json:"user_id"`
	QueUserId int `json:"que_user_id"`
	// QueUserFullname  string  `json:"que_user_fullname"`
	// CtmFullname      string  `json:"ctm_fullname"`
	// CtmFullnameEn    string  `json:"ctm_fullname_en"`
	// QueCode          string  `json:"que_code"`
	QueTypeId     int `json:"que_type_id"`
	QueAdmisId    int `json:"que_admis_id"`
	QuePriorityId int `json:"que_priority_id"`
	QueStatusId   int `json:"que_status_id"`
	// QueComment       string  `json:"que_comment"`
	// QueNote          string  `json:"que_note"`
	// QueDatetimeOut   string  `json:"que_datetime_out"`
	// QueRoomTotal     string  `json:"que_room_total"`
	QueDatetime string `json:"que_datetime"`
	QueCreate   string `json:"que_create"`
	// QueUpdate        string  `json:"que_update"`
	// CtmId            string  `json:"ctm_id"`
	// CtmImage         string  `json:"ctm_image"`
	// CtmFname         string  `json:"ctm_fname"`
	// CtmLname         string  `json:"ctm_lname"`
	// CtmBirthdate     string  `json:"ctm_birthdate"`
	// CtmBlood         string  `json:"ctm_blood"`
	// CtmWeight        float64 `json:"ctm_weight" gorm:"type:decimal(10,2)"`
	// CtmHeight        float64 `json:"ctm_height" gorm:"type:decimal(10,2)"`
	// CtmWaistline     float64 `json:"ctm_waistline" gorm:"type:decimal(10,2)"`
	// CtmChest         float64 `json:"ctm_chest" gorm:"type:decimal(10,2)"`
	// CtmTreatmentType int     `json:"ctm_treatment_type"`
	// CtmHealthComment string  `json:"ctm_health_comment"`
	// CtmDisease       string  `json:"ctm_disease"`
	// CtmMentalHealth  string  `json:"ctm_mental_health"`
	// CtmAllergic      string  `json:"ctm_allergic"`
	// UserFullname     string  `json:"user_fullname"`
	// RtCode           string  `json:"rt_code"`
	// RtName           string  `json:"rt_name"`
	// RoomCode         string  `json:"room_code"`
	// RoomTH           string  `json:"room_th"`
	// BedCode          string  `json:"bed_code"`
}

type ObjResponseQueueLabel struct {
	QueAll       int `json:"que_all"`
	QueMe        int `json:"que_me"`
	QueType1     int `json:"que_type_1"`
	QueType2     int `json:"que_type_2"`
	QuePriority1 int `json:"que_priority_1"`
	QuePriority2 int `json:"que_priority_2"`
	QuePriority3 int `json:"que_priority_3"`
	QuePriority4 int `json:"que_priority_4"`
}

type ObjPayloadSearchQueueCustomer struct {
	// ShopIds    []int   `json:"shop_ids" binding:"required"`
	ShopId     int     `json:"shop_ids"`
	SearchText *string `json:"search_text" binding:"required"`
}

type ObjResponseSearchQueueCustomer struct {
	ID         int    `json:"id"`
	CtmId      string `json:"ctm_id"`
	CtmFname   string `json:"ctm_fname"`
	CtmLname   string `json:"ctm_lname"`
	CtmFnameEn string `json:"ctm_fname_en"`
	CtmLnameEn string `json:"ctm_lname_en"`
}

type ObjPayloadSearchQueueDoctor struct {
	ShopId     int     `json:"shop_id" binding:"required"`
	SearchText *string `json:"search_text" binding:"required"`
}

type ObjResponseSearchQueueDoctor struct {
	ID             int    `json:"id"`
	UserEmail      string `json:"user_email"`
	UserFullname   string `json:"user_fullname"`
	UserFullnameEn string `json:"user_fullname_en"`
}

type ObjPayloadSearchQueueDiagnostic struct {
	ShopId     int     `json:"shop_id" binding:"required"`
	SearchText *string `json:"search_text" binding:"required"`
}

type ObjResponseSearchQueueDiagnostic struct {
	ID                 int    `json:"id"`
	ShopId             int    `json:"shop_id"`
	CategoryId         int    `json:"category_id"`
	DiagnosticCode     string `json:"diagnostic_code"`
	DiagnosticTh       string `json:"diagnostic_th"`
	DiagnosticEn       string `json:"diagnostic_en"`
	DiagnosticDetail   string `json:"diagnostic_detail"`
	DiagnosticIsActive int    `json:"diagnostic_is_active"`
	DiagnosticIsDel    int    `json:"diagnostic_is_del"`
}

type ObjQueryQueueDocSetting struct {
	ShopId             int    `json:"shop_id"`
	OpdIdDefault       string `json:"opd_id_default"`
	OpdNumberDefault   int    `json:"opd_number_default"`
	OpdNumberDigit     int    `json:"opd_number_digit"`
	OpdType            int    `json:"opd_type"`
	IpdIdDefault       string `json:"ipd_id_default"`
	IpdNumberDefault   int    `json:"ipd_number_default"`
	IpdNumberDigit     int    `json:"ipd_number_digit"`
	IpdType            int    `json:"ipd_type"`
	ServeIdDefault     string `json:"serve_id_default"`
	ServeNumberDefault int    `json:"serve_number_default"`
	ServeNumberDigit   int    `json:"serve_number_digit"`
	ServeType          int    `json:"serve_type"`
	CertIdDefault      string `json:"cert_id_default"`
	CertNumberDefault  int    `json:"cert_number_default"`
	CertNumberDigit    int    `json:"cert_number_digit"`
	CertType           int    `json:"cert_type"`
	PhrfIdDefault      string `json:"phrf_id_default"`
	PhrfNumberDefault  int    `json:"phrf_number_default"`
	PhrfNumberDigit    int    `json:"phrf_number_digit"`
	PhrfType           int    `json:"phrf_type"`
	SickIdDefault      string `json:"sick_id_default"`
	SickNumberDefault  int    `json:"sick_number_default"`
	SickNumberDigit    int    `json:"sick_number_digit"`
	SickType           int    `json:"sick_type"`
}

type ObjPayloadCreateQueue struct {
	ShopId          int    `json:"shop_id"`
	CustomerId      int    `json:"customer_id"`
	DoctorId        int    `json:"doctor_id"`
	DoctorFullname  string `json:"doctor_fullname"`
	RoomId          *int   `json:"room_id"`
	BedId           *int   `json:"bed_id"`
	QueUserId       int    `json:"que_user_id"`
	QueUserFullname string `json:"que_user_fullname"`
	QueTypeId       int    `json:"que_type_id"`
	QueAdmisId      int    `json:"que_admis_id"`
	QuePriorityId   int    `json:"que_priority_id"`
	QueDatetime     string `json:"que_datetime"`
	QueTeleCode     string `json:"que_tele_code"`
	QueTeleUrl      string `json:"que_tele_url"`
	DpmId           int    `json:"dpm_id"`
}

type ObjQueryQueueHistory struct {
	QueId          int    `json:"que_id"`
	QueCode        string `json:"que_code"`
	QueDatetime    string `json:"que_datetime"`
	QueCreate      string `json:"que_create"`
	QueAdmisId     int    `json:"que_admis_id"`
	UserFullname   string `json:"user_fullname"`
	UserFullnameEn string `json:"user_fullname_en"`
}

type ObjQueryOpdHistory struct {
	OpdId          int    `json:"opd_id"`
	OpdCode        string `json:"opd_code"`
	OpdDate        string `json:"opd_date"`
	OpdIsData      int    `json:"opd_is_data"`
	OpdCreate      string `json:"opd_create"`
	QueId          int    `json:"que_id"`
	QueAdmisId     int    `json:"que_admis_id"`
	QueCode        string `json:"que_code"`
	QueDatetime    string `json:"que_datetime"`
	QueCreate      string `json:"que_create"`
	UserId         int    `json:"user_id"`
	UserFullname   string `json:"user_fullname"`
	UserFullnameEn string `json:"user_fullname_en"`
	Detail         string `json:"detail"`
}

type ObjPayloadCreateOpd struct {
	UserId           int                             `json:"user_id" binding:"required"`
	QueueId          int                             `json:"queue_id" binding:"required"`
	CustomerId       int                             `json:"customer_id" binding:"required"`
	DpmId            *int                            `json:"dpm_id"`
	OpdCode          string                          `json:"opd_code"`
	OpdDate          string                          `json:"opd_date"`
	OpdBw            float64                         `json:"opd_bw"`
	OpdHt            float64                         `json:"opd_ht"`
	OpdBmi           float64                         `json:"opd_bmi"`
	OpdT             string                          `json:"opd_t"`
	OpdBsa           string                          `json:"opd_bsa"`
	OpdVas           string                          `json:"opd_vas"`
	OpdPr            string                          `json:"opd_pr"`
	OpdBp            string                          `json:"opd_bp"`
	OpdRr            string                          `json:"opd_rr"`
	OpdSys           string                          `json:"opd_sys"`
	OpdDia           string                          `json:"opd_dia"`
	OpdO2            string                          `json:"opd_o2"`
	OpdFag           int                             `json:"opd_fag"`
	OpdAlcohol       int                             `json:"opd_alcohol"`
	OpdCc            string                          `json:"opd_cc"`
	OpdHpi           string                          `json:"opd_hpi"`
	OpdPmh           string                          `json:"opd_pmh"`
	OpdDx            string                          `json:"opd_dx"`
	OpdIopatLe       string                          `json:"opd_iopat_le"`
	OpdVascRe        string                          `json:"opd_vasc_re"`
	OpdVascLe        string                          `json:"opd_vasc_le"`
	OpdVaccRe        string                          `json:"opd_vacc_re"`
	OpdVaccLe        string                          `json:"opd_vacc_le"`
	OpdIopatRe       string                          `json:"opd_iopat_re"`
	OpdGa            string                          `json:"opd_ga"`
	OpdPe            string                          `json:"opd_pe"`
	OpdNote          string                          `json:"opd_note"`
	OpdIsData        *int                            `json:"opd_is_data"`
	OpdSickStartdate *string                         `json:"opd_sick_startdate"`
	OpdSickEnddate   *string                         `json:"opd_sick_enddate"`
	OpdSickNotrest   int                             `json:"opd_sick_notrest"`
	OpdSickAir       int                             `json:"opd_sick_air"`
	OpdCustom        []ObjPayloadCreateOpdCustom     `json:"opd_custom"`
	OpdDiagnostic    []ObjPayloadCreateOpdDiagnostic `json:"opd_diagnostic"`
}

type ObjPayloadCreateOpdCustom struct {
	OpdcName  string `json:"opdc_name"`
	OpdcValue string `json:"opdc_value"`
}

type ObjPayloadCreateOpdDiagnostic struct {
	DiagnosticId     int    `json:"diagnostic_id"`
	DiagnosticCode   string `json:"diagnostic_code"`
	DiagnosticTh     string `json:"diagnostic_th"`
	DiagnosticEn     string `json:"diagnostic_en"`
	DiagnosticDetail string `json:"diagnostic_detail"`
}

type ObjPayloadUpdateOpd struct {
	UserId           int                             `json:"user_id" binding:"required"`
	DpmId            *int                            `json:"dpm_id"`
	OpdDate          string                          `json:"opd_date"`
	OpdBw            float64                         `json:"opd_bw"`
	OpdHt            float64                         `json:"opd_ht"`
	OpdBmi           float64                         `json:"opd_bmi"`
	OpdT             string                          `json:"opd_t"`
	OpdBsa           string                          `json:"opd_bsa"`
	OpdVas           string                          `json:"opd_vas"`
	OpdPr            string                          `json:"opd_pr"`
	OpdBp            string                          `json:"opd_bp"`
	OpdRr            string                          `json:"opd_rr"`
	OpdSys           string                          `json:"opd_sys"`
	OpdDia           string                          `json:"opd_dia"`
	OpdO2            string                          `json:"opd_o2"`
	OpdFag           int                             `json:"opd_fag"`
	OpdAlcohol       int                             `json:"opd_alcohol"`
	OpdCc            *string                         `json:"opd_cc"`
	OpdHpi           *string                         `json:"opd_hpi"`
	OpdPmh           *string                         `json:"opd_pmh"`
	OpdDx            *string                         `json:"opd_dx"`
	OpdGa            *string                         `json:"opd_ga"`
	OpdPe            *string                         `json:"opd_pe"`
	OpdNote          *string                         `json:"opd_note"`
	OpdIopatLe       string                          `json:"opd_iopat_le"`
	OpdVascRe        string                          `json:"opd_vasc_re"`
	OpdVascLe        string                          `json:"opd_vasc_le"`
	OpdVaccRe        string                          `json:"opd_vacc_re"`
	OpdVaccLe        string                          `json:"opd_vacc_le"`
	OpdIopatRe       string                          `json:"opd_iopat_re"`
	OpdIsData        *int                            `json:"opd_is_data"`
	OpdSickStartdate *string                         `json:"opd_sick_startdate"`
	OpdSickEnddate   *string                         `json:"opd_sick_enddate"`
	OpdSickNotrest   int                             `json:"opd_sick_notrest"`
	OpdSickAir       int                             `json:"opd_sick_air"`
	OpdCustom        []ObjPayloadUpdateOpdCustom     `json:"opd_custom"`
	OpdDiagnostic    []ObjPayloadUpdateOpdDiagnostic `json:"opd_diagnostic"`
}

type ObjPayloadUpdateOpdCustom struct {
	ID        int    `json:"id"`
	OpdcName  string `json:"opdc_name"`
	OpdcValue string `json:"opdc_value"`
}

type ObjPayloadUpdateOpdDiagnostic struct {
	ID               int    `json:"id"`
	DiagnosticId     int    `json:"diagnostic_id"`
	DiagnosticCode   string `json:"diagnostic_code"`
	DiagnosticTh     string `json:"diagnostic_th"`
	DiagnosticEn     string `json:"diagnostic_en"`
	DiagnosticDetail string `json:"diagnostic_detail"`
}

type ObjQueryUpdateOpd struct {
	UserId           int     `json:"user_id" binding:"required"`
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
}

type ObjPayloadCreateFile struct {
	QueueId    int    `json:"queue_id" binding:"required"`
	FileBase64 string `json:"file_base64" binding:"required"`
	IsUse      int    `json:"is_use" binding:"required"`
}

type ObjPayloadUpdateDoctorNote struct {
	QueNote string `json:"que_note"`
}

type ObjQueryMedicalCert struct {
	ID                int    `json:"id"`
	MedicalCertTypeId int    `json:"medical_cert_type_id"`
	UserId            int    `json:"user_id"`
	OpdId             int    `json:"opd_id"`
	OpdCode           string `json:"opd_code"`
	MdcCode           string `json:"mdc_code"`
	MdcIsPrint        int    `json:"mdc_is_print"`
	MdcIsDel          int    `json:"mdc_is_del"`
	MdcCreate         string `json:"mdc_create"`
	MdcUpdate         string `json:"mdc_update"`
	CtmPrefix         string `json:"ctm_prefix"`
	CtmFname          string `json:"ctm_fname"`
	CtmLname          string `json:"ctm_lname"`
	CtmNname          string `json:"ctm_nname"`
	CtmFnameEn        string `json:"ctm_fname_en"`
	CtmLnameEn        string `json:"ctm_lname_en"`
	UserFullname      string `json:"user_fullname"`
	MdctTh            string `json:"mdct_th"`
	MdctEn            string `json:"mdct_en"`
}

type ObjPayloadCreateMedicalCert struct {
	ShopId      int `json:"shop_id"`
	UserId      int `json:"user_id"`
	OpdId       int `json:"opd_id"`
	MdctId      int `json:"mdct_id"`
	MdctGroupId int `json:"mdct_group_id"`
}

type ObjPayloadCreateQueueTag struct {
	QueueId int   `json:"queue_id"`
	TagIds  []int `json:"tag_ids"`
}

type ObjQueryQueueTag struct {
	ID      int    `json:"id"`
	TagId   int    `json:"tag_id"`
	QueueId int    `json:"queue_id"`
	TagName string `json:"tag_name"`
}

// checking
type ObjQueryCheck struct {
	ID                 int    `json:"id"`
	ShopId             int    `json:"shop_id"`
	ShopName           string `json:"shop_name"`
	ShopPhone          string `json:"shop_phone"`
	ReceiptId          int    `json:"receipt_id"`
	ReceiptDetailId    int    `json:"receipt_detail_id"`
	Rec_code           string `json:"rec_code"`
	Rec_user_fullname  string `json:"rec_user_fullname"`
	UserId             int    `json:"user_id"`
	CustomerId         int    `json:"customer_id"`
	Ctm_id             string `json:"ctm_id"`
	Ctm_prefix         string `json:"ctm_prefix"`
	Ctm_fname          string `json:"ctm_fname"`
	Ctm_lname          string `json:"ctm_lname"`
	Ctm_fname_en       string `json:"ctm_fname_en"`
	Ctm_lname_en       string `json:"ctm_lname_en"`
	Ctm_gender         string `json:"ctm_gender"`
	Ctm_birthdate      string `json:"ctm_birthdate"`
	QueueId            int    `json:"queue_id"`
	Que_code           string `json:"que_code"`
	CheckingId         int    `json:"checking_id"`
	Category_name      string `json:"category_name"`
	ChkTypeId          int    `json:"chk_type_id"`
	ChkCode            string `json:"chk_code"`
	ChkName            string `json:"chk_name"`
	ChkUnit            string `json:"chk_unit"`
	ChkValue           string `json:"chk_value"`
	ChkUpload          string `json:"chk_upload"`
	ChkUploadSize      int    `json:"chk_upload_size"`
	ChkOld             string `json:"chk_old"`
	DirectionId        int    `json:"direction_id"`
	DirectionName      string `json:"direction_name"`
	DirectionDetail    string `json:"direction_detail"`
	ChkDirectionDetail string `json:"chk_direction_detail"`
	ChkFlag            string `json:"chk_flag"`
	Chk_date           string `json:"chk_date"`
	Chk_datetime       string `json:"chk_datetime"`
	ChkIsPrint         int    `json:"chk_is_print"`
	ChkIsReport        int    `json:"chk_is_report"`
	ChkIsActive        int    `json:"chk_is_active"`
	ChkCreate          string `json:"chk_create"`
	ChkUpdate          string `json:"chk_update"`
	Sticker_font_size  int    `json:"sticker_font_size"`
	Sticker_width      int    `json:"sticker_width"`
	Sticker_height     int    `json:"sticker_height"`
}

type ObjPayloadUpdateCheck struct {
	ChkValue           string  `json:"chk_value"`
	ChkUpload          string  `json:"chk_upload"`
	DirectionId        int     `json:"direction_id"`
	ChkDirectionDetail *string `json:"chk_direction_detail"`
	ChkFlag            string  `json:"chk_flag"`
}

type ObjPayloadUpdateCheckOld struct {
	ChkTypeId  int `json:"chk_type_id"`
	CheckingId int `json:"checking_id"`
	CustomerId int `json:"customer_id"`
}

// service
type ObjQueryService struct {
	ID              int     `json:"id"`
	ReceiptId       int     `json:"receipt_id"`
	RecCode         string  `json:"rec_code"`
	ReceiptDetailId int     `json:"receipt_detail_id"`
	ShopId          int     `json:"shop_id"`
	UserId          int     `json:"user_id"`
	SerCustomerId   int     `json:"ser_customer_id"`
	CustomerId      int     `json:"customer_id"`
	CourseId        int     `json:"course_id"`
	SerTranferId    int     `json:"ser_tranfer_id"`
	SerCode         string  `json:"ser_code"`
	SerName         string  `json:"ser_name"`
	SerLockDrug     int     `json:"ser_lock_drug"`
	SerQty          int     `json:"ser_qty"`
	SerUnit         string  `json:"ser_unit"`
	SerUseDate      int     `json:"ser_use_date"`
	SerExp          int     `json:"ser_exp"`
	SerExpDate      string  `json:"ser_exp_date"`
	SerUse          int     `json:"ser_use"`
	SerTranfer      int     `json:"ser_tranfer"`
	SerIsActive     int     `json:"ser_is_active"`
	SerCreate       string  `json:"ser_create"`
	SerUpdate       string  `json:"ser_update"`
	CourseAmount    float64 `json:"course_amount"`
	CourseIpd       float64 `json:"course_ipd"`
	CourseOpd       float64 `json:"course_opd"`
	CourseCost      float64 `json:"course_cost"`
	SerAmount       int     `json:"ser_amount"`
}

// X-Rey Detail
type XReyDetail struct {
	Queue_id       int    `json:"queue_id"`
	Que_code       string `json:"que_code"`
	Que_note       string `json:"que_note"`
	Que_directions string `json:"que_directions"`
	// Customer_id    int              `json:"customer_id"`
	Shop_id  int              `json:"shop_id"`
	Shop     ReceiptShop      `json:"shop" gorm:"-"`
	Customer ObjQueryCustomer `json:"customer" gorm:"-"`
	Checks   *[]ObjQueryCheck `json:"checks" gorm:"-"`
}

type ObjPayloadGetOpdHistoryPagination struct {
	ShopId       int    `json:"shop_id" binding:"required"`
	ShopMotherId int    `json:"shop_mother_id" binding:"required"`
	CustomerId   int    `json:"customer_id" binding:"required"`
	QueueId      int    `json:"queue_id" binding:"required"`
	SearchDate   string `json:"search_date"`
	SearchText   string `json:"search_text"`
	CurrentPage  int    `json:"current_page" binding:"required"`
	PerPage      int    `json:"per_page" binding:"required"`
}

type ObjResponseGetOpdHistoryPagination struct {
	Items     []ObjQueryOpdHistory `json:"items"`
	CountPage int                  `json:"count_page"`
	CountAll  int64                `json:"count_all"`
}

type ObjUpdateOpd struct {
	UserId           int     `json:"user_id"`
	DpmId            *int    `json:"dpm_id"`
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
	OpdCc            *string `json:"opd_cc"`
	OpdHpi           *string `json:"opd_hpi"`
	OpdPmh           *string `json:"opd_pmh"`
	OpdDx            *string `json:"opd_dx"`
	OpdGa            *string `json:"opd_ga"`
	OpdPe            *string `json:"opd_pe"`
	OpdNote          *string `json:"opd_note"`
	OpdIopatLe       string  `json:"opd_iopat_le"`
	OpdVascRe        string  `json:"opd_vasc_re"`
	OpdVascLe        string  `json:"opd_vasc_le"`
	OpdVaccRe        string  `json:"opd_vacc_re"`
	OpdVaccLe        string  `json:"opd_vacc_le"`
	OpdIopatRe       string  `json:"opd_iopat_re"`
	OpdIsData        *int    `json:"opd_is_data"`
	OpdSickStartdate *string `json:"opd_sick_startdate"`
	OpdSickEnddate   *string `json:"opd_sick_enddate"`
	OpdSickNotrest   int     `json:"opd_sick_notrest"`
	OpdSickAir       int     `json:"opd_sick_air"`
	OpdUpdate        string  `json:"opd_update"`
	OpdIsDel         int     `json:"opd_is_del"`
}

// labplus
type CheckCheckingLabplus struct {
	Id                   int    `json:"id"`
	Queue_id             int    `json:"queue_id"`
	Checking_id          int    `json:"checking_id"`
	Checking_is_labplus  int    `json:"checking_is_labplus"`
	Chk_value            string `json:"chk_value"`
	Chk_direction_detail string `json:"chk_direction_detail"`
	Chk_flag             string `json:"chk_flag"`
	Chk_code             string `json:"chk_code"`
	Chk_name             string `json:"chk_name"`
	Chk_is_active        int    `json:"chk_is_active"`
}

type UpdateCheckingLabplus struct {
	Chk_value            string `json:"chk_value"`
	Chk_direction_detail string `json:"chk_direction_detail"`
	Chk_flag             string `json:"chk_flag"`
	Chk_labplus_id       string `json:"chk_labplus_id"`
	Chk_labplus_name     string `json:"chk_labplus_name"`
	Chk_is_labplus       int    `json:"chk_is_labplus"`
	Chk_update           string `json:"chk_update"`
}

type CheckLabplus struct {
	Id                 int    `json:"id"`
	Shop_id            int    `json:"Shop_id"`
	Queue_id           int    `json:"queue_id"`
	Check_id           int    `json:"check_id"`
	Chkl_item_code     string `json:"chkl_item_code"`
	Chkl_item_name     string `json:"chkl_item_name"`
	Chkl_result        string `json:"chkl_result"`
	Chkl_result_normal string `json:"chkl_result_normal"`
	Chkl_item_unit     string `json:"chkl_item_unit"`
	Chkl_remark        string `json:"chkl_remark"`
	Chkl_flag          string `json:"chkl_flag"`
	Chkl_sort          int    `json:"chkl_sort"`
	Chkl_is_active     int    `json:"chkl_is_active"`
	Chkl_datetime      string `json:"chkl_datetime"`
}

type AddCheckLabplus struct {
	Shop_id            int    `json:"Shop_id"`
	Queue_id           int    `json:"queue_id"`
	Check_id           int    `json:"check_id"`
	Chkl_item_code     string `json:"chkl_item_code"`
	Chkl_item_name     string `json:"chkl_item_name"`
	Chkl_result        string `json:"chkl_result"`
	Chkl_result_normal string `json:"chkl_result_normal"`
	Chkl_item_unit     string `json:"chkl_item_unit"`
	Chkl_remark        string `json:"chkl_remark"`
	Chkl_flag          string `json:"chkl_flag"`
	Chkl_sort          int    `json:"chkl_sort"`
	Chkl_is_active     int    `json:"chkl_is_active"`
	Chkl_datetime      string `json:"chkl_datetime"`
}

type ObjPayloadCreateImageFileForAI struct {
	FileBase64 string `json:"file_base64" binding:"required"`
}
