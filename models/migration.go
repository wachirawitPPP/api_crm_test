package models

type Shop struct {
	ID                   int                 `json:"id"`
	PackageOrderDetailId int                 `json:"package_order_detail_id"`
	NatureTypeID         int                 `json:"nature_type_id"`
	CurrencyID           int                 `json:"currency_id"`
	ShopCode             string              `json:"shop_code"`
	ShopName             string              `json:"shop_name"`
	ShopPackage          string              `json:"shop_package"`
	ShopExpire           string              `json:"shop_expire"`
	ShopMotherId         int                 `json:"shop_mother_id"`
	ShopPointGiveRate    int                 `json:"shop_point_give_rate"`
	ShopPointUseRate     int                 `json:"shop_point_use_rate"`
	ShopPointCourse      int                 `json:"shop_point_course"`
	ShopPointChecking    int                 `json:"shop_point_checking"`
	ShopPointProduct     int                 `json:"shop_point_product"`
	ShopPointCoin        int                 `json:"shop_point_coin"`
	PackageOrderDetail   *PackageOrderDetail `json:"package_order_detail" gorm:"foreignKey:package_order_detail_id"`
	NatureType           *NatureType         `json:"nature_type" gorm:"foreignKey:nature_type_id"`
	Currency             *Currency           `json:"currency" gorm:"foreignKey:currency_id"`
	Users                []User              `json:"users" gorm:"many2many:user_shops;"`
}

type User struct {
	ID                int    `json:"id"`
	Username          string `json:"username"`
	UserEmail         string `json:"user_email"`
	UserFullname      string `json:"user_fullname"`
	UserFullnameEn    string `json:"user_fullname_en"`
	UserAddress       string `json:"user_address"`
	UserAddressEn     string `json:"user_address_en"`
	UserDistrict      string `json:"user_district"`
	UserDistrictEn    string `json:"user_district_en"`
	UserAmphoe        string `json:"user_amphoe"`
	UserAmphoeEn      string `json:"user_amphoe_en"`
	UserProvince      string `json:"user_province"`
	UserProvinceEn    string `json:"user_province_en"`
	UserZipcode       string `json:"user_zipcode"`
	UserZipcodeEn     string `json:"user_zipcode_en"`
	UserImage         string `json:"user_image"`
	UserType          int    `json:"user_type"`
	UserTypeName      string `json:"user_type_name"`
	UserGoogleId      string `json:"user_google_id"`
	UserGoogleEmail   string `json:"user_google_email"`
	UserTel           string `json:"user_tel"`
	UserLicense       string `json:"user_license"`
	NotDisplayQueue   int    `json:"not_display_queue"`
	NotDisplayAppoint int    `json:"not_display_appoint"`
	UserIsActivate    int    `json:"user_is_activate"`
	RoleID            int    `json:"role_id"`
	RoleNameTh        string `json:"role_name_th"`
	RoleNameEN        string `json:"role_name_en"`
	ShopRoleId        int    `json:"shop_role_id"`
	ShopID            string `json:"shop_id"`
	ShopName          string `json:"shop_name"`
	ShopNameEN        string `json:"shop_name_en"`
	TimesetOpen       string `json:"timeset_open"`
	TimesetClose      string `json:"timeset_close"`
	Shops             []Shop `json:"shops" gorm:"many2many:user_shops;"`
}

type Role struct {
	ID          int    `json:"id"`
	RoleNameTh  string `json:"role_name_th"`
	RoleNameEn  string `json:"role_name_en"`
	RoleNameLo  string `json:"role_name_lo"`
	RoleNameMm  string `json:"role_name_mm"`
	RoleSort    int    `json:"role_sort"`
	RoleIsOwner int    `json:"role_is_owner"`
	RoleCreate  string `json:"role_create"`
	ROleUpdate  string `json:"role_update"`
	Menus       []Menu `json:"menus" gorm:"many2many:role_menus;"`
}

type ShopRole struct {
	ID              int    `json:"id"`
	ShopId          int    `json:"shop_id"`
	RoleNameTh      string `json:"role_name_th"`
	RoleNameEn      string `json:"role_name_en"`
	RoleNameLo      string `json:"role_name_lo"`
	RoleNameMm      string `json:"role_name_mm"`
	RoleSort        int    `json:"role_sort"`
	RoleIsOwner     int    `json:"role_is_owner"`
	RoleDefault     int    `json:"role_default"`
	SrShowCost      int    `json:"sr_show_cost"`
	SrShowFeeUser   int    `json:"sr_show_fee_user"`
	SrShowFeeDoctor int    `json:"sr_show_fee_doctor"`
	SrShowCom       int    `json:"sr_show_com"`
	SrShowDashboard int    `json:"sr_show_dashboard"`
	SrShowCitizen   int    `json:"sr_show_citizen"`
	RoleCreate      string `json:"role_create"`
	RoleUpdate      string `json:"role_update"`
	Menus           []Menu `json:"menus" gorm:"many2many:shop_role_menus;"`
}

type ShopRoleMenu struct {
	ID         int `json:"id"`
	ShopRoleId int `json:"shop_role_id"`
	MenuId     int `json:"menu_id"`
}

type Menu struct {
	ID           int    `json:"id"`
	MenuGroupId  int    `json:"menu_group_id"`
	MenuTh       string `json:"menu_th"`
	MenuEn       string `json:"menu_en"`
	MenuLo       string `json:"menu_lo"`
	MenuMm       string `json:"menu_mm"`
	MenuLink     string `json:"menu_link"`
	MenuStatusId int    `json:"menu_status_id"`
	MenuOpenLink string `json:"menu_openlink"`
	MenuSort     int    `json:"menu_sort"`
	MenuUpdate   string `json:"menu_update"`
}

type Package struct {
	ID            int          `json:"id"`
	PackageTypeID int          `json:"package_type_id"`
	PackageName   string       `json:"package_name"`
	PackageCost   int          `json:"package_cost"`
	PackageUser   int          `json:"package_user"`
	PackageShop   int          `json:"package_shop"`
	PackageDate   int          `json:"package_date"`
	PackageSms    int          `json:"package_sms"`
	PackageCheck  int          `json:"package_check"`
	PackageDetail string       `json:"package_detail"`
	PackageSort   int          `json:"package_sort"`
	PackageCreate string       `json:"package_create"`
	PackageUpdate string       `json:"package_update"`
	PackageType   *PackageType `json:"package_type" gorm:"foreignKey:package_type_id"`
}

type PackageType struct {
	ID              int    `json:"id"`
	PackageTypeName string `json:"package_type_name"`
}

type PackageOrder struct {
	ID                    int                  `json:"id"`
	UserId                int                  `json:"user_id"`
	PackageOrderCode      string               `json:"package_order_code"`
	PackageOrderPaymentId int                  `json:"package_order_payment_id"`
	PackageOrderStatusId  int                  `json:"package_order_status_id"`
	PackageOrderComment   string               `json:"package_order_comment"`
	PackageOrderCreate    string               `json:"package_order_create"`
	PackageOrderUpdate    string               `json:"package_order_update"`
	PackageOrderDetails   []PackageOrderDetail `json:"package_order_details"`
	User                  *User                `json:"user"`
}

type PackageOrderDetail struct {
	ID                       int           `json:"id"`
	PackageOrderId           int           `json:"package_order_id"`
	PackageId                int           `json:"package_id"`
	PodShopReady             int           `json:"pod_shop_ready"`
	PackageOrderDetailCreate string        `json:"package_order_detail_create"`
	PackageOrderDetailUpdate string        `json:"package_order_detail_update"`
	PackageOrder             *PackageOrder `json:"package_order" gorm:"foreignKey:package_order_id"`
	Package                  *Package      `json:"package" gorm:"foreignKey:package_id"`
}

type Currency struct {
	ID           int    `json:"id"`
	CurrencyName string `json:"currency_name"`
	CurrencyTh   string `json:"currency_th"`
	CurrencyEn   string `json:"currency_en"`
	CurrencyLo   string `json:"currency_lo"`
	CurrencyMm   string `json:"currency_mm"`
}

type NatureType struct {
	ID               int    `json:"id"`
	NatureTypeName   string `json:"nature_type_name"`
	NatureTypeNameEn string `json:"nature_type_name_en"`
}

type Nature struct {
	ID                   int    `json:"id"`
	NatureTypeId         string `json:"nature_type_id"`
	NatureCode           string `json:"nature_code"`
	NatureName           string `json:"nature_name"`
	NatureNameEn         string `json:"nature_name_en"`
	NatureDiseaseChronic string `json:"nature_disease_chronic"`
	NatureDiseaseType    string `json:"nature_disease_type"`
}

type DocSetting struct {
	ID                    int    `json:"id"`
	CustomerIdDefault     string `json:"customer_id_default"`
	CustomerNumberDefault string `json:"customer_number_default"`
	CustomerNumberDigit   int    `json:"customer_number_digit"`
	CustomerType          int    `json:"customer_type"`
	OpdIdDefault          string `json:"opd_id_default"`
	OpdNumberDefault      string `json:"opd_number_default"`
	OpdNumberDigit        int    `json:"opd_number_digit"`
	OpdType               int    `json:"opd_type"`
	IpdIdDefault          string `json:"ipd_id_default"`
	IpdNumberDefault      string `json:"ipd_number_default"`
	IpdNumberDigit        int    `json:"ipd_number_digit"`
	IpdType               int    `json:"ipd_type"`
	CertIdDefault         string `json:"cert_id_default"`
	CertNumberDefault     string `json:"cert_number_default"`
	CertNumberDigit       int    `json:"cert_number_digit"`
	CertType              int    `json:"cert_type"`
	SickIdDefault         string `json:"sick_id_default"`
	SickNumberDefault     string `json:"sick_number_default"`
	SickNumberDigit       int    `json:"sick_number_digit"`
	SickType              int    `json:"sick_type"`
	PhrfIdDefault         string `json:"phrf_id_default"`
	PhrfNumberDefault     string `json:"phrf_number_default"`
	PhrfNumberDigit       int    `json:"phrf_number_digit"`
	PhrfType              int    `json:"phrf_type"`
	ServeIdDefault        string `json:"serve_id_default"`
	ServeNumberDefault    string `json:"serve_number_default"`
	ServeNumberDigit      int    `json:"serve_number_digit"`
	ServeType             int    `json:"serve_type"`
	StickerIdDefault      string `json:"sticker_id_default"`
	StickerNumberDefault  string `json:"sticker_number_default"`
	StickerNumberDigit    int    `json:"sticker_number_digit"`
	StickerType           int    `json:"sticker_type"`
	TransferIdDefault     string `json:"transfer_id_default"`
	TransferNumberDefault string `json:"transfer_number_default"`
	TransferNumberDigit   int    `json:"transfer_number_digit"`
	TransferType          int    `json:"transfer_type"`
	PurchaseIdDefault     string `json:"purchase_id_default"`
	PurchaseNumberDefault string `json:"purchase_number_default"`
	PurchaseNumberDigit   int    `json:"purchase_number_digit"`
	PurchaseType          int    `json:"purchase_type"`
	DruglotIdDefault      string `json:"druglot_id_default"`
	DruglotNumberDefault  string `json:"druglot_number_default"`
	DruglotNumberDigit    int    `json:"druglot_number_digit"`
	DruglotType           int    `json:"druglot_type"`
	DrugIdDefault         string `json:"drug_id_default"`
	DrugNumberDefault     string `json:"drug_number_default"`
	DrugNumberDigit       int    `json:"drug_number_digit"`
	DrugType              int    `json:"drug_type"`
	ToolIdDefault         string `json:"tool_id_default"`
	ToolNumberDefault     string `json:"tool_number_default"`
	ToolNumberDigit       int    `json:"tool_number_digit"`
	ToolType              int    `json:"tool_type"`
	CourseIdDefault       string `json:"course_id_default"`
	CourseNumberDefault   string `json:"course_number_default"`
	CourseNumberDigit     int    `json:"course_number_digit"`
	CourseType            int    `json:"course_type"`
	CheckIdDefault        string `json:"check_id_default"`
	CheckNumberDefault    string `json:"check_number_default"`
	CheckNumberDigit      int    `json:"check_number_digit"`
	CheckType             int    `json:"check_type"`
	LabIdDefault          string `json:"lab_id_default"`
	LabNumberDefault      string `json:"lab_number_default"`
	LabNumberDigit        int    `json:"lab_number_digit"`
	LabType               int    `json:"lab_type"`
	XrayIdDefault         string `json:"xray_id_default"`
	XrayNumberDefault     string `json:"xray_number_default"`
	XrayNumberDigit       int    `json:"xray_number_digit"`
	XrayType              int    `json:"xray_type"`
	InvoiceIdDefault      string `json:"invoice_id_default"`
	InvoiceNumberDefault  string `json:"invoice_number_default"`
	InvoiceNumberDigit    int    `json:"invoice_number_digit"`
	InvoiceType           int    `json:"invoice_type"`
	ReceiptIdDefault      string `json:"receipt_id_default"`
	ReceiptNumberDefault  string `json:"receipt_number_default"`
	ReceiptNumberDigit    int    `json:"receipt_number_digit"`
	ReceiptType           int    `json:"receipt_type"`
	TaxIdDefault          string `json:"tax_id_default"`
	TaxNumberDefault      string `json:"tax_number_default"`
	TaxNumberDigit        int    `json:"tax_number_digit"`
	TaxType               int    `json:"tax_type"`
	StickerFontSize       int    `json:"sticker_font_size"`
	StickerWidth          int    `json:"sticker_width"`
	StickerHeight         int    `json:"sticker_height"`
	StickerShowName       int    `json:"sticker_show_name"`
	StickerShowAddress    int    `json:"sticker_show_address"`
	StickerShowTel        int    `json:"sticker_show_tel"`
	StickerShowDate       int    `json:"sticker_show_date"`
	StickerShowExpdate    int    `json:"sticker_show_expdate"`
	StickerShowDetail     int    `json:"sticker_show_detail"`
	ShowDrugToolId        int    `json:"show_drug_tool_id"`
	ShowDrugToolTh        string `json:"show_drug_tool_th"`
	ShowDrugToolEn        string `json:"show_drug_tool_en"`
	ShowCourseCheckId     int    `json:"show_course_check_id"`
	ShowCourseCheckTh     string `json:"show_course_check_th"`
	ShowCourseCheckEn     string `json:"show_course_check_en"`
	PrintSetId            int    `json:"print_set_id"`
	ShowDateId            int    `json:"show_date_id"`
	ShowPageId            int    `json:"show_page_id"`
	ShowPointId           int    `json:"show_point_id"`
	PrintTh               int    `json:"print_th"`
	PrintEn               int    `json:"print_en"`
	PrintLa               int    `json:"print_la"`
	PrintCa               int    `json:"print_ca"`
	PrintA4               int    `json:"print_a4"`
	PrintA5               int    `json:"print_a5"`
	Print_80              int    `json:"print_80"`
	TaxId                 int    `json:"tax_id"`
	TaxSumId              int    `json:"tax_sum_id"`
	TaxSplitId            int    `json:"tax_split_id"`
	TaxRate               int    `json:"tax_rate"`
	TaxWithRate           int    `json:"tax_with_rate"`
	InvoiceCopy           int    `json:"invoice_copy"`
	ReceiptCopy           int    `json:"receipt_copy"`
	TaxCopy               int    `json:"tax_copy"`
	PurchaseCopy          int    `json:"purchase_copy"`
	TransferCopy          int    `json:"transfer_copy"`
	InvoiceCommentId      int    `json:"invoice_comment_id"`
	ReceiptCommentId      int    `json:"receipt_comment_id"`
	TaxCommentId          int    `json:"tax_comment_id"`
	PurchaseCommentId     int    `json:"purchase_comment_id"`
	TransferCommentId     int    `json:"transfer_comment_id"`
	InvoiceComment        string `json:"invoice_comment"`
	ReceiptComment        string `json:"receipt_comment"`
	TaxComment            string `json:"tax_comment"`
	PurchaseComment       string `json:"purchase_comment"`
	TransferComment       string `json:"transfer_comment"`
}

type TimeSet struct {
	ID                  int    `json:"id"`
	Timeset_open        string `json:"timeset_open"`
	Timeset_close       string `json:"timeset_close"`
	Timeset_range       int    `json:"timeset_range"`
	Timeset_day_id      int    `json:"timeset_day_id"`
	Timeset_day_amount  int    `json:"timeset_day_amount"`
	Timeset_room_id     int    `json:"timeset_room_id"`
	Timeset_room_amount int    `json:"timeset_room_amount"`
	Timeset_comment_id  int    `json:"timeset_comment_id"`
	Timeset_sunday      int    `json:"timeset_sunday"`
	Timeset_monday      int    `json:"timeset_monday"`
	Timeset_tuesday     int    `json:"timeset_tuesday"`
	Timeset_wednesday   int    `json:"timeset_wednesday"`
	Timeset_thursday    int    `json:"timeset_thursday"`
	Timeset_friday      int    `json:"timeset_friday"`
	Timeset_saturday    int    `json:"timeset_saturday"`
}

type ResponsePaginationEmpty struct {
	Result_data   []string `json:"result_data"`
	Count_of_page int      `json:"count_of_page"`
	Count_all     int      `json:"count_all"`
}

type Tag struct {
	ID        int    `json:"id"`
	ShopId    int    `json:"shop_id"`
	TagTypeId int    `json:"tag_type_id"`
	TagName   string `json:"tag_name"`
	TagIsDel  int    `json:"tag_is_del"`
}

type TagType struct {
	ID        int    `json:"id"`
	TagTypeTh string `json:"tag_type_th"`
	TagTypeEn string `json:"tag_type_en"`
}

type Holiday struct {
	Id              int    `json:"id"`
	Shop_id         int    `json:"shop_id"`
	Holiday_type_id int    `json:"holiday_type_id"`
	Holiday_name    string `json:"holiday_name"`
	Holiday_date    string `json:"holiday_date"`
	Holiday_is_del  int    `json:"holiday_is_del"`
}

type HolidayType struct {
	Id            int    `json:"id"`
	HolidayTypeTh string `json:"holiday_type_th"`
	HolidayTypeEn string `json:"holiday_type_en"`
}

type Topic struct {
	Id           int    `json:"id"`
	Shop_id      int    `json:"shop_id"`
	Topic_th     string `json:"topic_th"`
	Topic_en     string `json:"topic_en"`
	Topic_is_del int    `json:"topic_is_del"`
}

type Commission struct {
	ID               int               `json:"id"`
	ShopId           int               `json:"shop_id"`
	AccountCodeId    int               `json:"account_code_id"`
	CommissionName   string            `json:"commission_name"`
	CommissionType   int               `json:"commission_type"`
	CommissionRate   float64           `json:"commission_rate" gorm:"type:decimal(10,2)"`
	CommissionIsDel  int               `json:"commission_is_del"`
	CommissionUpdate string            `json:"commission_update"`
	CommissionCreate string            `json:"commission_create"`
	CommissionLevels []CommissionLevel `json:"commission_levels"`
}

type CommissionLevel struct {
	ID                    int     `json:"id"`
	CommissionId          int     `json:"commission_id"`
	CommissionLevelAmount float64 `json:"commission_level_amount" gorm:"type:decimal(10,2)"`
	CommissionLevelRate   float64 `json:"commission_level_rate" gorm:"type:decimal(10,2)"`
	CommissionLevelIsDel  int     `json:"commission_level_is_del"`
	CommissionLevelCreate string  `json:"commission_level_create"`
	CommissionLevelUpdate string  `json:"commission_level_update"`
}

type Fee struct {
	ID            int     `json:"id"`
	ShopId        int     `json:"shop_id"`
	FeeName       string  `json:"fee_name"`
	FeeType       int     `json:"fee_type"`
	AccountCodeId int     `json:"account_code_id"`
	FeeRate       float64 `json:"fee_rate" gorm:"type:decimal(10,2)"`
	FeeIsDel      int     `json:"fee_is_del"`
	FeeCreate     string  `json:"fee_create"`
	FeeUpdate     string  `json:"fee_update"`
}

type LogPaySolutions struct {
	ID         int    `json:"id"`
	Log_text   string `json:"log_text"`
	Log_create string `json:"log_create"`
}

type PackagePayment struct {
	ID             int     `json:"id"`
	UserId         int     `json:"user_id"`
	PackageOrderId int     `json:"package_order_id"`
	PpNumber       string  `json:"pp_number"`
	PpBy           string  `json:"pp_by"`
	PpCost         float64 `json:"pp_cost" gorm:"type:decimal(10,2)"`
	PpDate         string  `json:"pp_date"`
	PpTime         string  `json:"pp_time"`
	PpEvidence     string  `json:"pp_evidence"`
	PpCheck        int     `json:"pp_check"`
	PpCreate       string  `json:"pp_create"`
	PpUpdate       string  `json:"pp_update"`
}

type UserRefreshToken struct {
	UserRftkId        int    `json:"user_rftk_id"`
	UserId            int    `json:"user_id"`
	UserRftkToken     string `json:"user_rftk_token"`
	UserRftkIpAddress string `json:"user_rftk_ip_address"`
	UserRftkBrowser   string `json:"user_rftk_browser"`
	UserRftkSite      int    `json:"user_rftk_site"`
	UserRftkStatus    int    `json:"user_rftk_status"`
	UserRftkExpire    string `json:"user_rftk_expire"`
	UserRftkCreate    string `json:"user_rftk_create"`
}

type UserAccessToken struct {
	UserActkId        int    `json:"user_actk_id"`
	UserId            int    `json:"user_id"`
	UserActkToken     string `json:"user_actk_token"`
	UserActkIpAddress string `json:"user_actk_ip_address"`
	UserActkBrowser   string `json:"user_actk_browser"`
	UserActkExpire    string `json:"user_actk_expire"`
	UserActkCreate    string `json:"user_actk_create"`
}

type Notification struct {
	ID            int    `json:"id"`
	UserId        int    `json:"user_id"`
	NtTitle       string `json:"nt_title"`
	NtDescription string `json:"nt_description"`
	NtImage       string `json:"nt_image"`
	NtLink        string `json:"nt_link"`
	NtType        string `json:"nt_type"`
	NtIsRead      int    `json:"nt_is_read"`
	NtCreate      string `json:"nt_create"`
	NtUpdate      string `json:"nt_update"`
}

type Course struct {
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
	CourseFee      float64 `json:"course_fee"`
	CourseIsActive int     `json:"course_is_active"`
	CourseIsDel    int     `json:"course_is_del"`
	CourseCreate   string  `json:"course_create"`
	CourseUpdate   string  `json:"course_update"`
}

type Checking struct {
	ID                int     `json:"id"`
	ShopId            int     `json:"shop_id"`
	CategoryId        int     `json:"category_id"`
	CheckingTypeId    int     `json:"checking_type_id"`
	CheckingCode      string  `json:"checking_code"`
	CheckingCodeAcc   string  `json:"checking_code_acc"`
	CheckingNameAcc   string  `json:"checking_name_acc"`
	CheckingName      string  `json:"checking_name"`
	CheckingUnit      string  `json:"checking_unit"`
	CheckingOpd       float64 `json:"checking_opd"`
	CheckingIpd       float64 `json:"checking_ipd"`
	CheckingCost      float64 `json:"checking_cost"`
	CheckingFeeDf     float64 `json:"checking_fee_df"`
	CheckingFee       float64 `json:"checking_fee"`
	CheckingIsLabplus int     `json:"checking_is_labplus"`
	AccCodeIdCost     float64 `json:"acc_code_id_cost"`
	AccCodeIdFeeDf    float64 `json:"acc_code_id_fee_df"`
	AccCodeIdFee      int     `json:"acc_code_id_fee"`
	AccCodeIdCom      int     `json:"acc_code_id_com"`
	CheckingImage     string  `json:"checking_image"`
	CheckingIsActive  int     `json:"checking_is_active"`
	CheckingIsDel     int     `json:"checking_is_del"`
	CheckingCreate    string  `json:"checking_create"`
	CheckingUpdate    string  `json:"checking_update"`
}

type Product struct {
	ID            int    `json:"id"`
	ShopId        int    `json:"shop_id"`
	CategoryId    int    `json:"category_id"`
	PdTypeId      int    `json:"pd_type_id"`
	PdCode        string `json:"pd_code"`
	PdName        string `json:"pd_name"`
	PdCodeAcc     string `json:"pd_code_acc"`
	PdNameAcc     string `json:"pd_name_acc"`
	PdImage       string `json:"pd_image"`
	PdDescription string `json:"pd_description"`
	PdBarcode     string `json:"pd_barcode"`
	PdIsSerial    int    `json:"pd_is_serial"`
	PdNarcotic2   int    `json:"pd_narcotic2"`
	PdNarcotic3   int    `json:"pd_narcotic3"`
	PdNarcotic4   int    `json:"pd_narcotic4"`
	PdNarcotic5   int    `json:"pd_narcotic5"`
	DrugDirection string `json:"checking_code"`
	PdAmountNoti  int    `json:"pd_amount_noti"`
	PdExpireNoti  int    `json:"pd_expire_noti"`
	PdIsOver      int    `json:"pd_is_over"`
	PdIsActive    int    `json:"pd_is_active"`
	PdIsDel       int    `json:"pd_is_del"`
	PdCreate      string `json:"pd_create"`
	PdUpdate      string `json:"pd_update"`
}

type Customer struct {
	ID                  int     `json:"id"`
	ShopId              int     `json:"shop_id"`
	ShopMotherId        int     `json:"shop_mother_id"`
	CustomerGroupId     int     `json:"customer_group_id"`
	UserId              int     `json:"user_id"`
	CtmId               string  `json:"ctm_id"`
	CtmCitizenId        string  `json:"ctm_citizen_id"`
	CtmPassportId       string  `json:"ctm_passport_id"`
	CtmPrefix           string  `json:"ctm_prefix"`
	CtmFname            string  `json:"ctm_fname"`
	CtmLname            string  `json:"ctm_lname"`
	CtmNname            string  `json:"ctm_nname"`
	CtmFnameEn          string  `json:"ctm_fname_en"`
	CtmLnameEn          string  `json:"ctm_lname_en"`
	CtmGender           string  `json:"ctm_gender"`
	CtmNation           string  `json:"ctm_nation"`
	CtmReligion         string  `json:"ctm_religion"`
	CtmEduLevel         string  `json:"ctm_edu_level"`
	CtmMaritalStatus    string  `json:"ctm_marital_status"`
	CtmBlood            string  `json:"ctm_blood"`
	CtmEmail            string  `json:"ctm_email"`
	CtmTel              string  `json:"ctm_tel"`
	CtmTel_2            string  `json:"ctm_tel_2"`
	CtmBirthdate        string  `json:"ctm_birthdate"`
	CtmAddress          string  `json:"ctm_address"`
	CtmDistrict         string  `json:"ctm_district"`
	CtmAmphoe           string  `json:"ctm_amphoe"`
	CtmProvince         string  `json:"ctm_province"`
	CtmZipcode          string  `json:"ctm_zipcode"`
	CtmWeight           float64 `json:"ctm_weight" gorm:"type:decimal(10,2)"`
	CtmHeight           float64 `json:"ctm_height" gorm:"type:decimal(10,2)"`
	CtmWaistline        float64 `json:"ctm_waistline" gorm:"type:decimal(10,2)"`
	CtmChest            float64 `json:"ctm_chest" gorm:"type:decimal(10,2)"`
	CtmTreatmentType    int     `json:"ctm_treatment_type"`
	RightTreatmentId    int     `json:"right_treatment_id"`
	CtmAllergic         string  `json:"ctm_allergic"`
	CtmMentalHealth     string  `json:"ctm_mental_health"`
	CtmDisease          string  `json:"ctm_disease"`
	CtmComment          string  `json:"ctm_comment"`
	CtmHealthComment    string  `json:"ctm_health_comment"`
	CtmImage            string  `json:"ctm_image"`
	CtmImageSize        int     `json:"ctm_image_size"`
	CtmPoint            int     `json:"ctm_point"`
	CtmCoin             float64 `json:"ctm_coin" gorm:"type:decimal(10,2)"`
	LineToken           string  `json:"line_token"`
	LineSend            int     `json:"line_send"`
	LineSendDate        string  `json:"line_send_date"`
	FacebookId          string  `json:"facebook_id"`
	CompanyName         string  `json:"company_name"`
	CompanyTax          string  `json:"company_tax"`
	CompanyTel          string  `json:"company_tel"`
	CompanyEmail        string  `json:"company_email"`
	CompanyAddress      string  `json:"company_address"`
	CompanyDistrict     string  `json:"company_district"`
	CompanyAmphoe       string  `json:"company_amphoe"`
	CompanyProvince     string  `json:"company_province"`
	CompanyZipcode      string  `json:"company_zipcode"`
	CtmSubscribeOpd     int     `json:"ctm_subscribe_opd"`
	CtmSubscribeLab     int     `json:"ctm_subscribe_lab"`
	CtmSubscribeCert    int     `json:"ctm_subscribe_cert"`
	CtmSubscribeReceipt int     `json:"ctm_subscribe_receipt"`
	CtmSubscribeAppoint int     `json:"ctm_subscribe_appoint"`
	CtmIsActive         int     `json:"ctm_is_active"`
	CtmIsDel            int     `json:"ctm_is_del"`
	CtmCreate           string  `json:"ctm_create"`
	CtmUpdate           string  `json:"ctm_update"`
	// CustomerGroup       *CustomerGroup `json:"customer_group" gorm:"foreignKey:customer_group_id"`
}

type RefRightTreatment struct {
	ID       int    `json:"id"`
	RtCode   string `json:"rt_code"`
	RtName   string `json:"rt_name"`
	RtNameEn string `json:"rt_name_en"`
}

type RefDepartment struct {
	ID        int    `json:"id"`
	ShopId    int    `json:"shop_id"`
	DpmName   string `json:"dpm_name"`
	DpmNameEn string `json:"dpm_name_en"`
	DpmIsDel  int    `json:"dpm_is_del"`
	DpmCreate string `json:"dpm_create"`
	DpmUpdate string `json:"dpm_update"`
}

type CustomerGroup struct {
	ID         int    `json:"id"`
	ShopId     int    `json:"shop_id"`
	CgName     string `json:"cg_name"`
	CgSaveType int    `json:"cg_save_type"`
	CgSave     int    `json:"cg_save"`
	CgIsActive int    `json:"cg_is_active"`
	CgIsDel    int    `json:"cg_is_del"`
	CgCreate   string `json:"cg_create"`
	CgUpdate   string `json:"cg_update"`
}

type CustomerTag struct {
	ID         int `json:"id"`
	CustomerId int `json:"customer_id"`
	TagId      int `json:"tag_id"`
}

type CustomerFamilys struct {
	TableName    string `gorm:"-"`
	ID           int    `json:"id"`
	CustomerId   int    `json:"customer_id"`
	CfCustomerId int    `json:"cf_customer_id"`
	CfRelation   string `json:"cf_relation"`
	CfCreate     string `json:"cf_create"`
	CfUpdate     string `json:"cf_update"`
}

type CustomerContact struct {
	ID         int    `json:"id"`
	CustomerId int    `json:"customer_id"`
	CcName     string `json:"cc_name"`
	CcTel      string `json:"cc_tel"`
	CcRelation string `json:"cc_relation"`
	CcCreate   string `json:"cc_create"`
	CcUpdate   string `json:"cc_update"`
}

type Promotion struct {
	ID                 int    `json:"id"`
	ShopId             int    `json:"shop_id"`
	UserId             int    `json:"user_id"`
	PmtMessageTypeId   int    `json:"pmt_message_type_id"`
	PmtSendTypeId      int    `json:"pmt_send_type_id"`
	CustomerId         int    `json:"customer_id"`
	CustomerGroupId    int    `json:"customer_group_id"`
	CustomerBirthMonth int    `json:"customer_birth_month"`
	PmtTitle           string `json:"pmt_title"`
	PmtContent         string `json:"pmt_content"`
	PmtIsDel           int    `json:"pmt_is_del"`
	PmtUpdate          string `json:"pmt_update"`
	PmtCreate          string `json:"pmt_create"`
}

type SentMessage struct {
	ID           int    `json:"id"`
	PromotionId  int    `json:"promotion_id"`
	CustomerId   int    `json:"customer_id"`
	SmCreditUsed int    `json:"sm_credit_used"`
	SmTo         string `json:"sm_to"`
	SmStatus     int    `json:"sm_status"`
	SmIsDel      int    `json:"sm_is_del"`
	SmCreate     string `json:"sm_create"`
}

type Article struct {
	ID            int          `json:"id"`
	ShopId        int          `json:"shop_id"`
	UserId        int          `json:"user_id"`
	ArTitle       string       `json:"ar_title"`
	ArSlug        string       `json:"ar_slug"`
	ArExcerpt     string       `json:"ar_excerpt"`
	ArText        string       `json:"ar_text"`
	ArDescription string       `json:"ar_description"`
	ArKeyword     string       `json:"ar_keyword"`
	ArThumbnail   string       `json:"ar_thumbnail"`
	ArStart       string       `json:"ar_start"`
	ArEnd         string       `json:"ar_end"`
	ArStatusId    int          `json:"ar_status_id"`
	ArTypeId      int          `json:"ar_type_id"`
	ArView        int          `json:"ar_view"`
	ArCreate      string       `json:"ar_create"`
	ArUpdate      string       `json:"ar_update"`
	User          *User        `json:"user"`
	ArticleTags   []ArticleTag `json:"article_tags" gorm:"many2many:article_tag_maps;"`
}

type ArticleTag struct {
	ID      int    `json:"id"`
	ShopId  int    `json:"shop_id"`
	ArtText string `json:"art_text"`
}

type ShopStore struct {
	ID            int    `json:"id"`
	ShopId        int    `json:"shop_id"`
	AccountCodeId int    `json:"account_code_id"`
	SsTypeId      int    `json:"ss_type_id"`
	SsName        string `json:"ss_name"`
	SsIsOver      int    `json:"ss_is_over"`
	SsIsActive    int    `json:"ss_is_active"`
	SsCreate      string `json:"ss_create"`
	SsUpdate      string `json:"ss_update"`
}

type Topical struct {
	ID              int              `json:"id"`
	ShopId          int              `json:"shop_id"`
	TopicalTypeId   int              `json:"topical_type_id"`
	TopicalName     string           `json:"topical_name"`
	TopicalDetail   string           `json:"topical_detail"`
	TopicalIsActive int              `json:"topical_is_active"`
	TopicalIsDel    int              `json:"topical_is_del"`
	TopicalCreate   string           `json:"topical_create"`
	TopicalUpdate   string           `json:"topical_update"`
	TopicalProduct  []TopicalProduct `json:"topical_product"`
}

type TopicalProduct struct {
	ID        int     `json:"id"`
	TopicalId int     `json:"topical_id"`
	ProductId int     `json:"product_id"`
	TpdAmount float64 `json:"tpd_amount" gorm:"type:decimal(10,2)"`
}

type Expense struct {
	ID            int     `json:"id"`
	ShopId        int     `json:"shop_id"`
	CategoryId    int     `json:"category_id"`
	AccountCodeId *int    `json:"account_code_id"`
	UserId        int     `json:"user_id"`
	EpsCode       string  `json:"eps_code"`
	EpsName       string  `json:"eps_name"`
	EpsWallet     float64 `json:"eps_wallet" gorm:"type:decimal(20,2)"`
	Eps_file      string  `json:"eps_file"`
	EpsDate       string  `json:"eps_date"`
	EpsIsDel      int     `json:"eps_is_del"`
	EpsModify     string  `json:"eps_modify"`
}

type Categorys struct {
	TableName      string `gorm:"-"`
	ID             int    `json:"id"`
	ShopId         int    `json:"shop_id"`
	CategoryTypeId int    `json:"category_type_id"`
	CategoryName   string `json:"category_name"`
	CategoryIsDel  int    `json:"category_is_del"`
}

type Queue struct {
	ID              int    `json:"id"`
	ShopId          int    `json:"shop_id"`
	CustomerId      int    `json:"customer_id"`
	RoomId          *int   `json:"room_id"`
	BedId           *int   `json:"bed_id"`
	UserId          int    `json:"user_id"`
	QueUserId       int    `json:"que_user_id"`
	QueUserFullname string `json:"que_user_fullname"`
	CtmFullname     string `json:"ctm_fullname"`
	CtmFullnameEn   string `json:"ctm_fullname_en"`
	QueCode         string `json:"que_code"`
	QueTypeId       int    `json:"que_type_id"`
	QueAdmisId      int    `json:"que_admis_id"`
	QuePriorityId   int    `json:"que_priority_id"`
	QueStatusId     int    `json:"que_status_id"`
	QueComment      string `json:"que_comment"`
	QueNote         string `json:"que_note"`
	QueDatetimeOut  string `json:"que_datetime_out"`
	QueRoomTotal    string `json:"que_room_total"`
	QueDatetime     string `json:"que_datetime"`
	QueTimeEnd      int    `json:"que_time_end"`
	QueTeleCode     string `json:"que_tele_code"`
	QueTeleUrl      string `json:"que_tele_url"`
	QueCreate       string `json:"que_create"`
	QueUpdate       string `json:"que_update"`
	DpmId           int    `json:"dpm_id"`
	QueNumber       int    `json:"que_number"`
	QueDpmCode      string `json:"que_dpm_code"`
}

type Opd struct {
	ID               int             `json:"id"`
	UserId           int             `json:"user_id"`
	QueueId          int             `json:"queue_id"`
	CustomerId       int             `json:"customer_id"`
	DpmId            *int            `json:"dpm_id"`
	OpdCode          string          `json:"opd_code"`
	OpdDate          string          `json:"opd_date"`
	OpdBw            float64         `json:"opd_bw"`
	OpdHt            float64         `json:"opd_ht"`
	OpdBmi           float64         `json:"opd_bmi"`
	OpdT             string          `json:"opd_t"`
	OpdBsa           string          `json:"opd_bsa"`
	OpdVas           string          `json:"opd_vas"`
	OpdPr            string          `json:"opd_pr"`
	OpdBp            string          `json:"opd_bp"`
	OpdRr            string          `json:"opd_rr"`
	OpdSys           string          `json:"opd_sys"`
	OpdDia           string          `json:"opd_dia"`
	OpdO2            string          `json:"opd_o2"`
	OpdFag           int             `json:"opd_fag"`
	OpdAlcohol       int             `json:"opd_alcohol"`
	OpdCc            string          `json:"opd_cc"`
	OpdHpi           string          `json:"opd_hpi"`
	OpdPmh           string          `json:"opd_pmh"`
	OpdDx            string          `json:"opd_dx"`
	OpdIopatLe       string          `json:"opd_iopat_le"`
	OpdVascRe        string          `json:"opd_vasc_re"`
	OpdVascLe        string          `json:"opd_vasc_le"`
	OpdVaccRe        string          `json:"opd_vacc_re"`
	OpdVaccLe        string          `json:"opd_vacc_le"`
	OpdIopatRe       string          `json:"opd_iopat_re"`
	OpdGa            string          `json:"opd_ga"`
	OpdPe            string          `json:"opd_pe"`
	OpdNote          string          `json:"opd_note"`
	OpdIsData        int             `json:"opd_is_data"`
	OpdSickStartdate *string         `json:"opd_sick_startdate"`
	OpdSickEnddate   *string         `json:"opd_sick_enddate"`
	OpdSickNotrest   int             `json:"opd_sick_notrest"`
	OpdSickAir       int             `json:"opd_sick_air"`
	OpdIsDel         int             `json:"opd_is_del"`
	OpdUpdate        string          `json:"opd_update"`
	OpdCreate        string          `json:"opd_create"`
	OpdCustoms       []OpdCustom     `json:"opd_custom"`
	OpdDiagnostics   []OpdDiagnostic `json:"opd_diagnostic"`
	User             *User           `json:"user" gorm:"foreignKey:user_id"`
	Queue            *Queue          `json:"queue" gorm:"foreignKey:queue_id"`
}

type OpdCustom struct {
	ID        int    `json:"id"`
	OpdId     int    `json:"opd_id"`
	OpdcName  string `json:"opdc_name"`
	OpdcValue string `json:"opdc_value"`
}

type OpdDiagnostic struct {
	ID               int    `json:"id"`
	OpdId            int    `json:"opd_id"`
	DiagnosticId     int    `json:"diagnostic_id"`
	DiagnosticCode   string `json:"diagnostic_code"`
	DiagnosticTh     string `json:"diagnostic_th"`
	DiagnosticEn     string `json:"diagnostic_en"`
	DiagnosticDetail string `json:"diagnostic_detail"`
}

type Room struct {
	ID         int    `json:"id"`
	ShopId     int    `json:"shop_id"`
	RoomTypeId int    `json:"room_type_id"`
	RoomCode   string `json:"room_code"`
	RoomTh     string `json:"room_th"`
	RoomEn     string `json:"room_en"`
	RoomIsDel  int    `json:"room_is_del"`
}

type Bed struct {
	ID          int    `json:"id"`
	ShopId      int    `json:"shop_id"`
	RoomId      int    `json:"room_id"`
	BedCode     string `json:"bed_code"`
	BedLock     int    `json:"bed_lock"`
	BedIsDel    int    `json:"bed_is_del"`
	QueueId     int    `json:"queue_id"`
	QueStatusId *int   `json:"que_status_id"`
}

type QueueFile struct {
	ID         int    `json:"id"`
	QueueId    int    `json:"queue_id"`
	QuefPath   string `json:"quef_path"`
	QuefSize   int    `json:"quef_size"`
	QuefIsUse  int    `json:"quef_is_use"`
	QuefModify string `json:"quef_modify"`
}

type QueueTag struct {
	ID      int `json:"id"`
	TagId   int `json:"tag_id"`
	QueueId int `json:"queue_id"`
}

type MedicalCertType struct {
	ID           int    `json:"id"`
	MdctTh       string `json:"mdct_th"`
	MdctEn       string `json:"mdct_en"`
	MdctGroup_id int    `json:"mdct_group_id"`
}

type MedicalCert struct {
	ID                int    `json:"id"`
	MedicalCertTypeId int    `json:"medical_cert_type_id"`
	UserId            int    `json:"user_id"`
	OpdId             int    `json:"opd_id"`
	MdcCode           string `json:"mdc_code"`
	MdcIsPrint        int    `json:"mdc_is_print"`
	MdcIsDel          int    `json:"mdc_is_del"`
	MdcCreate         string `json:"mdc_create"`
	MdcUpdate         string `json:"mdc_update"`
}

type Direction struct {
	ID                int    `json:"id"`
	ShopId            int    `json:"shop_id"`
	CategoryId        int    `json:"category_id"`
	DirectionName     string `json:"direction_name"`
	DirectionDetail   string `json:"direction_detail"`
	DirectionTypeId   int    `json:"direction_type_id"`
	DirectionIsActive int    `json:"direction_is_active"`
	DirectionIsDel    int    `json:"direction_is_del"`
}

type Check struct {
	ID                 int    `json:"id"`
	ShopId             int    `json:"shop_id"`
	ReceiptId          int    `json:"receipt_id"`
	ReceiptDetailId    int    `json:"receipt_detail_id"`
	UserId             int    `json:"user_id"`
	CustomerId         int    `json:"customer_id"`
	QueueId            int    `json:"queue_id"`
	CheckingId         int    `json:"checking_id"`
	ChkTypeId          int    `json:"chk_type_id"`
	ChkCode            string `json:"chk_code"`
	ChkName            string `json:"chk_name"`
	ChkUnit            string `json:"chk_unit"`
	ChkValue           string `json:"chk_value"`
	ChkUpload          string `json:"chk_upload"`
	ChkUploadSize      int    `json:"chk_upload_size"`
	ChkOld             string `json:"chk_old"`
	DirectionId        int    `json:"direction_id"`
	ChkDirectionDetail string `json:"chk_direction_detail"`
	ChkFlag            string `json:"chk_flag"`
	ChkDate            string `json:"chk_date"`
	ChkIsPrint         int    `json:"chk_is_print"`
	ChkIsReport        int    `json:"chk_is_report"`
	ChkIsActive        int    `json:"chk_is_active"`
	ChkCreate          string `json:"chk_create"`
	ChkUpdate          string `json:"chk_update"`
}

type Service struct {
	ID              int    `json:"id"`
	ReceiptId       int    `json:"receipt_id"`
	ReceiptDetailId int    `json:"receipt_detail_id"`
	ShopId          int    `json:"shop_id"`
	UserId          int    `json:"user_id"`
	SerCustomerId   int    `json:"ser_customer_id"`
	CustomerId      int    `json:"customer_id"`
	CourseId        int    `json:"course_id"`
	SerTranferId    int    `json:"ser_tranfer_id"`
	SerCode         string `json:"ser_code"`
	SerName         string `json:"ser_name"`
	SerLockDrug     int    `json:"ser_lock_drug"`
	SerQty          int    `json:"ser_qty"`
	SerUnit         string `json:"ser_unit"`
	SerUseDate      int    `json:"ser_use_date"`
	SerExp          int    `json:"ser_exp"`
	SerExpDate      string `json:"ser_exp_date"`
	SerUse          int    `json:"ser_use"`
	SerIsActive     int    `json:"ser_is_active"`
	SerCreate       string `json:"ser_create"`
	SerUpdate       string `json:"ser_update"`
}

type Invoice struct {
	ID                   int      `json:"id"`
	ShopId               int      `json:"shop_id"`
	UserId               int      `json:"user_id"`
	CustomerId           int      `json:"customer_id"`
	CustomerOnlineId     int      `json:"customer_online_id"`
	QueueId              *int     `json:"queue_id"`
	OrderId              int      `json:"order_id"`
	InvCode              string   `json:"inv_code"`
	InvFullname          string   `json:"inv_fullname"`
	InvTel               string   `json:"inv_tel"`
	InvEmail             string   `json:"inv_email"`
	InvAddress           string   `json:"inv_address"`
	InvDistrict          string   `json:"inv_district"`
	InvAmphoe            string   `json:"inv_amphoe"`
	InvProvince          string   `json:"inv_province"`
	InvZipcode           string   `json:"inv_zipcode"`
	InvComment           string   `json:"inv_comment"`
	InvTotalPrice        float64  `json:"inv_total_price"`
	InvDiscount          float64  `json:"inv_discount"`
	InvBeforVat          float64  `json:"inv_befor_vat"`
	TaxTypeId            int      `json:"tax_type_id"`
	TaxRate              int      `json:"tax_rate"`
	InvVat               float64  `json:"inv_vat"`
	InvTotal             float64  `json:"inv_total"`
	InvDeposit           float64  `json:"inv_deposit"`
	InvPayTotal          float64  `json:"inv_pay_total"`
	InvIsActive          int      `json:"inv_is_active"`
	InvDatetime          string   `json:"inv_datetime"`
	InvCreate            string   `json:"inv_create"`
	InvUpdate            string   `json:"inv_update"`
	DpmId                int      `json:"dpm_id"`
	Inv_discount_type_id int      `json:"inv_discount_type_id"`
	Inv_discount_item    float64  `json:"inv_discount_item"`
	Inv_discount_value   float64  `json:"inv_discount_value"`
	Inv_eclaim_id        *int     `json:"inv_eclaim_id"`
	Inv_eclaim_rate      *float64 `json:"inv_eclaim_rate"`
	Inv_eclaim_over      *float64 `json:"inv_eclaim_over"`
	Inv_eclaim_total     *float64 `json:"inv_eclaim_total"`
}

type InvoiceDetail struct {
	ID                    int      `json:"id"`
	InvoiceId             int      `json:"invoice_id"`
	CourseId              *int     `json:"course_id" `
	CheckingId            *int     `json:"checking_id"`
	ProductId             *int     `json:"product_id"`
	ProductStoreId        *int     `json:"product_store_id"`
	ProductUnitId         *int     `json:"product_unit_id"`
	CoinId                *int     `json:"coin_id"`
	RoomId                *int     `json:"room_id"`
	QueueId               *int     `json:"queue_id"`
	OrderDetailId         int      `json:"order_detail_id"`
	InvdTypeId            int      `json:"invd_type_id"`
	InvdCode              string   `json:"invd_code"`
	InvdName              string   `json:"invd_name"`
	InvdQty               float64  `json:"invd_qty"`
	InvdSetQty            float64  `json:"invd_set_qty"`
	InvdLimitQty          float64  `json:"invd_limit_qty"`
	InvdUnit              string   `json:"invd_unit"`
	InvdCost              float64  `json:"invd_cost"`
	InvdPrice             float64  `json:"invd_price"`
	InvdDiscount          float64  `json:"invd_discount"`
	InvdAmount            float64  `json:"invd_amount"`
	InvdRate              float64  `json:"invd_rate"`
	TopicalId             *int     `json:"topical_id"`
	InvdTopical           string   `json:"invd_topical"`
	InvdDirection         string   `json:"invd_direction"`
	TaxTypeId             int      `json:"tax_type_id"`
	TaxRate               int      `json:"tax_rate"`
	InvdVat               float64  `json:"invd_vat"`
	InvdTotal             float64  `json:"invd_total"`
	InvdIsSet             int      `json:"invd_is_set"`
	InvdIsActive          int      `json:"invd_is_active"`
	InvdModify            string   `json:"invd_modify"`
	Invd_eclaim           *float64 `json:"invd_eclaim"`
	Invd_discount_type_id int      `json:"invd_discount_type_id"`
	Invd_discount_item    *float64 `json:"invd_discount_item"`
}

type Receipt struct {
	ID                 int      `json:"id"`
	ShopId             int      `json:"shop_id"`
	UserId             int      `json:"user_id"`
	CustomerId         int      `json:"customer_id"`
	CustomerOnlineId   int      `json:"customer_online_id"`
	QueueId            *int     `json:"queue_id"`
	InvoiceId          int      `json:"invoice_id"`
	AccountListId      *int     `json:"account_list_id"`
	RecCode            string   `json:"rec_code"`
	RecFullname        string   `json:"rec_fullname"`
	RecTel             string   `json:"rec_tel"`
	RecEmail           string   `json:"rec_email"`
	RecAddress         string   `json:"rec_address"`
	RecDistrict        string   `json:"rec_district"`
	RecAmphoe          string   `json:"rec_amphoe"`
	RecProvince        string   `json:"rec_province"`
	RecZipcode         string   `json:"rec_zipcode"`
	RecComment         string   `json:"rec_comment"`
	RecTotalPrice      float64  `json:"rec_total_price"`
	RecDiscount        float64  `json:"rec_discount"`
	RecBeforVat        float64  `json:"rec_befor_vat"`
	TaxTypeId          int      `json:"tax_type_id"`
	TaxRate            int      `json:"tax_rate"`
	RecVat             float64  `json:"rec_vat"`
	RecTotal           float64  `json:"rec_total"`
	RecPaymentType     int      `json:"rec_payment_type"`
	RecTypeId          int      `json:"rec_type_id"`
	RecPeriod          int      `json:"rec_period"`
	RecPay             float64  `json:"rec_pay"`
	RecBalance         float64  `json:"rec_balance"`
	RecPayTotal        float64  `json:"rec_pay_total"`
	RecPayDatetime     string   `json:"rec_pay_datetime"`
	RecDescription     string   `json:"rec_discription"`
	RecAccount         int      `json:"rec_account"`
	RecUserId          int      `json:"rec_user_id"`
	RecUserFullname    string   `json:"rec_user_fullname"`
	RecPointGive       int      `json:"rec_point_give"`
	RecIsProcess       int      `json:"rec_point_used"`
	RecPointUsed       int      `json:"rec_is_process"`
	RecIsActive        int      `json:"rec_is_active"`
	RecCreate          string   `json:"rec_create"`
	RecUpdate          string   `json:"rec_update"`
	RecDiscountTypeId  int      `json:"rec_discount_type_id"`
	Rec_discount_item  *float64 `json:"rec_discount_item"`
	Rec_discount_value *float64 `json:"rec_discount_value"`
	DpmId              *int     `json:"dpm_id"`
	Rec_eclaim_id      *int     `json:"rec_eclaim_id"`
	Rec_eclaim_rate    *float64 `json:"rec_eclaim_rate"`
	Rec_eclaim_over    *float64 `json:"rec_eclaim_over"`
	Rec_eclaim_total   *float64 `json:"rec_eclaim_total"`
}

type ReceiptDetail struct {
	ID                    int      `json:"id"`
	ReceiptId             int      `json:"receipt_id"`
	CourseId              *int     `json:"course_id"`
	CheckingId            *int     `json:"checking_id"`
	ProductId             *int     `json:"product_id"`
	ProductStoreId        *int     `json:"product_store_id"`
	ProductUnitId         *int     `json:"product_unit_id"`
	CoinId                *int     `json:"coin_id"`
	RoomId                *int     `json:"room_id"`
	QueueId               *int     `json:"queue_id"`
	InvoiceDetailId       int      `json:"invoice_detail_id"`
	RecdTypeId            int      `json:"recd_type_id"`
	RecdCode              string   `json:"recd_code"`
	RecdName              string   `json:"recd_name"`
	RecdQty               float64  `json:"recd_qty"`
	RecdSetQty            float64  `json:"recd_set_qty"`
	RecdLimitQty          float64  `json:"recd_limit_qty"`
	RecdRate              float64  `json:"recd_rate"`
	TopicalId             *int     `json:"topical_id"`
	RecdTopical           string   `json:"recd_topical"`
	RecdDirection         string   `json:"recd_direction"`
	RecdUnit              string   `json:"recd_unit"`
	RecdCost              float64  `json:"recd_cost"`
	RecdPrice             float64  `json:"recd_price"`
	RecdDiscount          float64  `json:"recd_discount"`
	RecdAmount            float64  `json:"recd_amount"`
	TaxTypeId             int      `json:"tax_type_id"`
	TaxRate               int      `json:"tax_rate"`
	RecdVat               float64  `json:"recd_vat"`
	RecdTotal             float64  `json:"recd_total"`
	RecdIsSet             int      `json:"recd_is_set"`
	RecdIsActive          int      `json:"recd_is_active"`
	RecdModify            string   `json:"recd_modify"`
	Recd_eclaim           *float64 `json:"recd_eclaim"`
	Recd_discount_type_id int      `json:"recd_discount_type_id"`
	Recd_discount_item    *float64 `json:"recd_discount_item"`
}

type Coupon struct {
	ID             int     `json:"id"`
	ShopId         int     `json:"shop_id"`
	CouponCode     string  `json:"coupon_code"`
	CouponName     string  `json:"coupon_name"`
	CouponLimit    int     `json:"coupon_limit"`
	CouponUse      int     `json:"coupon_use"`
	CouponAmount   float64 `json:"coupon_amount"`
	CouponExpdate  string  `json:"coupon_expdate"`
	CouponCreate   string  `json:"coupon_create"`
	CouponIsActive int     `json:"coupon_is_active"`
}

type Handling struct {
	ID            int     `json:"id"`
	ShopId        int     `json:"shop_id"`
	HandFromId    int     `json:"hand_from_id"`
	InvoiceId     *int    `json:"invoice_id"`
	QueueId       *int    `json:"queue_id"`
	UserId        int     `json:"user_id"`
	RoleId        int     `json:"role_id"`
	HandTypeId    int     `json:"hand_type_id"`
	HandIsDeposit int     `json:"hand_is_deposit"`
	HandPrice     float64 `json:"hand_price" gorm:"type:decimal(10,2)"`
	HandIsDel     int     `json:"hand_is_del"`
	HandModify    string  `json:"hand_modify"`
}

type AccountList struct {
	ID            int    `json:"id"`
	ShopId        int    `json:"shop_id"`
	AccountCodeId int    `json:"account_code_id"`
	AclCode       string `json:"acl_code"`
	AclName       string `json:"acl_name"`
	AclTypeId     int    `json:"acl_type_id"`
	AclIsDel      int    `json:"acl_is_del"`
	AclCreate     string `json:"acl_create"`
	AclUpdate     string `json:"acl_update"`
}

type AccountCode struct {
	ID            int    `json:"id"`
	ShopId        int    `json:"shop_id"`
	AccountTypeId int    `json:"account_type_id"`
	AccCode       string `json:"acc_code"`
	AccTh         string `json:"acc_th"`
	AccEn         string `json:"acc_en"`
	AccIsDel      int    `json:"acc_is_del"`
	AccCreate     string `json:"acc_create"`
	AccUpdate     string `json:"acc_update"`
}

type ShopPackage struct {
	Id               int `json:"id"`
	Shop_type_id     int `json:"shop_type_id"`
	Package_order_id int `json:"package_order_id"`
	Package_id       int `json:"package_id"`
}

type ShopPackageOrderId struct {
	Id               int `json:"id"`
	Package_order_id int `json:"package_order_id"`
	Package_id       int `json:"package_id"`
}

type ShopAddon struct {
	Id         int `json:"id"`
	Shop_id    int `json:"shop_id"`
	Package_id int `json:"package_id"`
}

type Diagnostic struct {
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

type LogSms struct {
	LogId     int    `json:"log_id"`
	UserId    int    `json:"user_id"`
	Shopname  string `json:"shopname"`
	Username  string `json:"username"`
	LogType   string `json:"log_type"`
	LogText   string `json:"log_text"`
	LogCredit string `json:"log_credit"`
	LogCreate string `json:"log_create"`
}

type ProductUnit struct {
	ID        int    `json:"id"`
	ProductId int    `json:"product_id"`
	UnitId    int    `json:"unit_id"`
	PuRate    int    `json:"pu_rate"`
	PuIsDel   int    `json:"pu_is_del"`
	PuCreate  string `json:"pu_create"`
	PuUpdate  string `json:"pu_update"`
}

type ProductShopPrice struct {
	ID            int     `json:"id"`
	ShopId        int     `json:"shop_id"`
	ProductUnitId int     `json:"product_unit_id"`
	PspPriceOpd   float64 `json:"psp_price_opd" gorm:"type:decimal(10,2)"`
	PspPriceIpd   float64 `json:"psp_price_ipd" gorm:"type:decimal(10,2)"`
	PspIsDefault  int     `json:"psp_is_default"`
	PspIsDel      int     `json:"psp_is_del"`
	PspCreate     string  `json:"psp_create"`
	PspUpdate     string  `json:"psp_update"`
}

type ProductShop struct {
	ID        int    `json:"id"`
	ShopId    int    `json:"shop_id"`
	ProductId int    `json:"product_id"`
	PsCreate  string `json:"ps_create"`
	PsUpdate  string `json:"ps_update"`
}

type Popup struct {
	PopupId       int    `json:"popup_id"`
	PopupTitle    string `json:"popup_title"`
	PopupText     string `json:"popup_text"`
	PopupImage    string `json:"popup_image"`
	PopupLink     string `json:"popup_link"`
	PopupStatusId int    `json:"popup_status_id"`
	PopupCreate   string `json:"popup_create"`
	PopupUpdate   string `json:"popup_update"`
}

type PopupUser struct {
	ID      int `json:"id"`
	PopupId int `json:"popup_id"`
	UserId  int `json:"user_id"`
}

type Vendor struct {
	ID             int    `json:"id"`
	ShopId         int    `json:"shop_id"`
	VendorCode     string `json:"vendor_code"`
	VendorName     string `json:"vendor_name"`
	VendorTel      string `json:"vendor_tel"`
	VendorCompany  string `json:"vendor_company"`
	VendorBranch   string `json:"vendor_branch"`
	VendorTax      string `json:"vendor_tax"`
	VendorAddress  string `json:"vendor_address"`
	VendorDistrict string `json:"vendor_district"`
	VendorAmphoe   string `json:"vendor_amphoe"`
	VendorProvince string `json:"vendor_province"`
	VendorZipcode  string `json:"vendor_zipcode`
	VendorEmail    string `json:"vendor_email"`
	VendorIsActive int    `json:"vendor_is_active"`
	VendorIdDel    int    `json:"vendor_id_del"`
	VendorCreate   string `json:"vendor_create"`
	VendorUpdate   string `json:"vendor_update"`
}
