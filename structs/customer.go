package structs

type ObjPayloadGetCustomerPagination struct {
	ShopId      int     `json:"shop_id" binding:"required"`
	CtmIsActive *int    `json:"ctm_is_active" binding:"required"`
	SearchText  *string `json:"search_text" binding:"required"`
	CurrentPage int     `json:"current_page" binding:"required"`
	PerPage     int     `json:"per_page" binding:"required"`
}

type ObjQueryGetCustomerPagination struct {
	ID                  int     `json:"id"`
	ShopId              int     `json:"shop_id"`
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
	CtmComment          string  `json:"ctm_comment"`
	CtmWeight           float64 `json:"ctm_weight" gorm:"type:decimal(10,2)"`
	CtmHeight           float64 `json:"ctm_height" gorm:"type:decimal(10,2)"`
	CtmWaistline        float64 `json:"ctm_waistline" gorm:"type:decimal(10,2)"`
	CtmChest            float64 `json:"ctm_chest" gorm:"type:decimal(10,2)"`
	CtmTreatmentType    int     `json:"ctm_treatment_type"`
	RightTreatmentId    int     `json:"right_treatment_id"`
	CtmAllergic         string  `json:"ctm_allergic"`
	CtmMentalHealth     string  `json:"ctm_mental_health"`
	CtmDisease          string  `json:"ctm_disease"`
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
	Ctm_subscribe_pdpa  int     `json:"ctm_subscribe_pdpa"`
	CtmIsActive         int     `json:"ctm_is_active"`
	CtmIsDel            int     `json:"ctm_is_del"`
	CtmCreate           string  `json:"ctm_create"`
	CtmUpdate           string  `json:"ctm_update"`
	// join
	CgName  string `json:"cg_name"`
	Rt_code string `json:"rt_code"`
	Rt_name string `json:"rt_name"`
}

type ObjResponseGetCustomerPagination struct {
	Items     []ObjQueryGetCustomerPagination `json:"items"`
	CountPage int                             `json:"count_page"`
	CountAll  int64                           `json:"count_all"`
}

type ObjQueryCustomer struct {
	ID                    int     `json:"id"`
	ShopId                int     `json:"shop_id"`
	CustomerGroupId       int     `json:"customer_group_id"`
	UserId                int     `json:"user_id"`
	CtmId                 string  `json:"ctm_id"`
	CtmCitizenId          string  `json:"ctm_citizen_id"`
	CtmPassportId         string  `json:"ctm_passport_id"`
	CtmPrefix             string  `json:"ctm_prefix"`
	CtmFname              string  `json:"ctm_fname"`
	CtmLname              string  `json:"ctm_lname"`
	CtmNname              string  `json:"ctm_nname"`
	CtmFnameEn            string  `json:"ctm_fname_en"`
	CtmLnameEn            string  `json:"ctm_lname_en"`
	CtmGender             string  `json:"ctm_gender"`
	CtmNation             string  `json:"ctm_nation"`
	CtmReligion           string  `json:"ctm_religion"`
	CtmEduLevel           string  `json:"ctm_edu_level"`
	CtmMaritalStatus      string  `json:"ctm_marital_status"`
	CtmBlood              string  `json:"ctm_blood"`
	CtmEmail              string  `json:"ctm_email"`
	CtmTel                string  `json:"ctm_tel"`
	CtmTel_2              string  `json:"ctm_tel_2"`
	CtmBirthdate          string  `json:"ctm_birthdate"`
	CtmAddress            string  `json:"ctm_address"`
	CtmDistrict           string  `json:"ctm_district"`
	CtmAmphoe             string  `json:"ctm_amphoe"`
	CtmProvince           string  `json:"ctm_province"`
	CtmZipcode            string  `json:"ctm_zipcode"`
	CtmComment            string  `json:"ctm_comment"`
	CtmWeight             float64 `json:"ctm_weight" gorm:"type:decimal(10,2)"`
	CtmHeight             float64 `json:"ctm_height" gorm:"type:decimal(10,2)"`
	CtmWaistline          float64 `json:"ctm_waistline" gorm:"type:decimal(10,2)"`
	CtmChest              float64 `json:"ctm_chest" gorm:"type:decimal(10,2)"`
	CtmTreatmentType      int     `json:"ctm_treatment_type"`
	RightTreatmentId      int     `json:"right_treatment_id"`
	CtmAllergic           string  `json:"ctm_allergic"`
	CtmMentalHealth       string  `json:"ctm_mental_health"`
	CtmDisease            string  `json:"ctm_disease"`
	CtmHealthComment      string  `json:"ctm_health_comment"`
	CtmImage              string  `json:"ctm_image"`
	CtmImageSize          int     `json:"ctm_image_size"`
	CtmPoint              int     `json:"ctm_point"`
	CtmCoin               float64 `json:"ctm_coin" gorm:"type:decimal(10,2)"`
	LineToken             string  `json:"line_token"`
	LineSend              int     `json:"line_send"`
	LineSendDate          string  `json:"line_send_date"`
	FacebookId            string  `json:"facebook_id"`
	CompanyName           string  `json:"company_name"`
	CompanyTax            string  `json:"company_tax"`
	CompanyTel            string  `json:"company_tel"`
	CompanyEmail          string  `json:"company_email"`
	CompanyAddress        string  `json:"company_address"`
	CompanyDistrict       string  `json:"company_district"`
	CompanyAmphoe         string  `json:"company_amphoe"`
	CompanyProvince       string  `json:"company_province"`
	CompanyZipcode        string  `json:"company_zipcode"`
	CtmSubscribeOpd       int     `json:"ctm_subscribe_opd"`
	CtmSubscribeLab       int     `json:"ctm_subscribe_lab"`
	CtmSubscribeCert      int     `json:"ctm_subscribe_cert"`
	CtmSubscribeReceipt   int     `json:"ctm_subscribe_receipt"`
	CtmSubscribeAppoint   int     `json:"ctm_subscribe_appoint"`
	Ctm_subscribe_pdpa    int     `json:"ctm_subscribe_pdpa"`
	CtmIsActive           int     `json:"ctm_is_active"`
	CtmIsDel              int     `json:"ctm_is_del"`
	CtmCreate             string  `json:"ctm_create"`
	CtmUpdate             string  `json:"ctm_update"`
	CtmSubscribePdpa      int     `json:"ctm_subscribe_pdpa"`
	CtmSubscribePdpaToken string  `json:"ctm_subscribe_pdpa_token"`
	CtmSubscribePdpaImage string  `json:"ctm_subscribe_pdpa_image"`
	// join
	CgName     string `json:"cg_name"`
	CgSaveType int    `json:"cg_save_type"` //ประเภทส่วนลด : 1 สกุลเงิน, 2 %
	CgSave     int    `json:"cg_save"`
	RtCode     string `json:"rt_code"`
	RtName     string `json:"rt_name"`
	RtNameEn   string `json:"rt_name_en"`
}
type ObjGetCustomerOnline struct {
	ID          int    `json:"id"`
	CoCitizenId string `json:"co_citizen_id"`
	CoEmail     string `json:"co_email"`
	CoFname     string `json:"co_fname"`
	CoLname     string `json:"co_lname"`
	Co_Line_id  string `json:"co_line_id"`
}

type ObjQueryCheckImport struct {
	CtmId        string `json:"ctm_id"`
	CtmCitizenId string `json:"ctm_citizen_id"`
}

type ObjQueryCustomerTag struct {
	ID         int `json:"id"`
	CustomerId int `json:"customer_id"`
	TagId      int `json:"tag_id"`
	// join
	TagName   string `json:"tag_name"`
	TagTypeTh string `json:"tag_type_th"`
	TagTypeEn string `json:"tag_type_en"`
}

type ObjQueryCustomerFamily struct {
	TableName    string `gorm:"-"`
	ID           int    `json:"id"`
	CustomerId   int    `json:"customer_id"`
	CfCustomerId int    `json:"cf_customer_id"`
	CtmFname     string `json:"ctm_fname"`
	CtmLname     string `json:"ctm_lname"`
	CfRelation   string `json:"cf_relation"`
	CfCreate     string `json:"cf_create"`
	CfUpdate     string `json:"cf_update"`
}

type ObjQueryCustomerContact struct {
	ID         int    `json:"id"`
	CustomerId int    `json:"customer_id"`
	CcName     string `json:"cc_name"`
	CcTel      string `json:"cc_tel"`
	CcRelation string `json:"cc_relation"`
	CcCreate   string `json:"cc_create"`
	CcUpdate   string `json:"cc_update"`
}

type ObjResponseGetCustomerById struct {
	Customer *ObjQueryCustomer         `json:"customer"`
	Tag      []ObjQueryCustomerTag     `json:"tag"`
	Family   []ObjQueryCustomerFamily  `json:"family"`
	Contact  []ObjQueryCustomerContact `json:"contact"`
	Balance  *DashboardPayment         `json:"balance"`
}
type ObjResponseGetCustomerOnlineById struct {
	ID      int    `json:"id"`
	CoEmail string `json:"co_email"`
	CoFname string `json:"co_fname"`
	CoLname string `json:"co_lname"`
}

type ObjPayloadSearchFamily struct {
	ShopId     int     `json:"shop_id" binding:"required"`
	NotInIds   []int   `json:"not_in_ids" binding:"required"`
	SearchText *string `json:"search_text" binding:"required"`
}

type ObjResponseSearchFamily struct {
	ID       int    `json:"id"`
	CtmFname string `json:"ctm_fname"`
	CtmLname string `json:"ctm_lname"`
}

type ObjPayloadCheckCitizenId struct {
	CustomerId int    `json:"customer_id"`
	ShopId     int    `json:"shop_id" binding:"required"`
	CitizenId  string `json:"citizen_id" binding:"required"`
}

type ObjPayloadCreateCustomer struct {
	ShopId           int     `json:"shop_id"`
	CustomerGroupId  int     `json:"customer_group_id"`
	UserId           int     `json:"user_id"`
	CtmId            string  `json:"ctm_id"`
	CtmCitizenId     string  `json:"ctm_citizen_id"`
	CtmPassportId    string  `json:"ctm_passport_id"`
	CtmPrefix        string  `json:"ctm_prefix"`
	CtmFname         string  `json:"ctm_fname"`
	CtmLname         string  `json:"ctm_lname"`
	CtmNname         string  `json:"ctm_nname"`
	CtmFnameEn       string  `json:"ctm_fname_en"`
	CtmLnameEn       string  `json:"ctm_lname_en"`
	CtmGender        string  `json:"ctm_gender"`
	CtmNation        string  `json:"ctm_nation"`
	CtmReligion      string  `json:"ctm_religion"`
	CtmEduLevel      string  `json:"ctm_edu_level"`
	CtmMaritalStatus string  `json:"ctm_marital_status"`
	CtmBlood         string  `json:"ctm_blood"`
	CtmEmail         string  `json:"ctm_email"`
	CtmTel           string  `json:"ctm_tel"`
	CtmTel_2         string  `json:"ctm_tel_2"`
	CtmBirthdate     string  `json:"ctm_birthdate"`
	CtmAddress       string  `json:"ctm_address"`
	CtmDistrict      string  `json:"ctm_district"`
	CtmAmphoe        string  `json:"ctm_amphoe"`
	CtmProvince      string  `json:"ctm_province"`
	CtmZipcode       string  `json:"ctm_zipcode"`
	CtmWeight        float64 `json:"ctm_weight" gorm:"type:decimal(10,2)"`
	CtmHeight        float64 `json:"ctm_height" gorm:"type:decimal(10,2)"`
	CtmWaistline     float64 `json:"ctm_waistline" gorm:"type:decimal(10,2)"`
	CtmChest         float64 `json:"ctm_chest" gorm:"type:decimal(10,2)"`
	CtmTreatmentType int     `json:"ctm_treatment_type"`
	RightTreatmentId int     `json:"right_treatment_id"`
	CtmAllergic      string  `json:"ctm_allergic"`
	CtmMentalHealth  string  `json:"ctm_mental_health"`
	CtmDisease       string  `json:"ctm_disease"`
	CtmComment       string  `json:"ctm_comment"`
	CtmHealthComment string  `json:"ctm_health_comment"`
	CtmImage         string  `json:"ctm_image"`
	CtmImageSize     int     `json:"ctm_image_size"`
	CtmPoint         int     `json:"ctm_point"`
	CtmCoin          float64 `json:"ctm_coin" gorm:"type:decimal(10,2)"`
	// LineToken           string `json:"line_token"`
	// LineSend            int    `json:"line_send"`
	// LineSendDate        string `json:"line_send_date"`
	// FacebookId          string `json:"facebook_id"`
	CompanyName         string `json:"company_name"`
	CompanyTax          string `json:"company_tax"`
	CompanyTel          string `json:"company_tel"`
	CompanyEmail        string `json:"company_email"`
	CompanyAddress      string `json:"company_address"`
	CompanyDistrict     string `json:"company_district"`
	CompanyAmphoe       string `json:"company_amphoe"`
	CompanyProvince     string `json:"company_province"`
	CompanyZipcode      string `json:"company_zipcode"`
	CtmSubscribeOpd     int    `json:"ctm_subscribe_opd"`
	CtmSubscribeLab     int    `json:"ctm_subscribe_lab"`
	CtmSubscribeCert    int    `json:"ctm_subscribe_cert"`
	CtmSubscribeReceipt int    `json:"ctm_subscribe_receipt"`
	CtmSubscribeAppoint int    `json:"ctm_subscribe_appoint"`
	CtmIsActive         int    `json:"ctm_is_active"`
	// CtmIsDel            string `json:"ctm_is_del"`
	// CtmCreate           string `json:"ctm_create"`
	// CtmUpdate           string `json:"ctm_update"`
	TagSelected     []int             `json:"tag_selected"`
	FamilySelected  []FamilySelected  `json:"family_selected"`
	ContactSelected []ContactSelected `json:"contact_selected"`
}

type ObjPayloadUpdateCustomer struct {
	ShopId           int     `json:"shop_id"`
	CustomerGroupId  int     `json:"customer_group_id"`
	UserId           int     `json:"user_id"`
	CtmId            string  `json:"ctm_id"`
	CtmCitizenId     string  `json:"ctm_citizen_id"`
	CtmPassportId    string  `json:"ctm_passport_id"`
	CtmPrefix        string  `json:"ctm_prefix"`
	CtmFname         string  `json:"ctm_fname"`
	CtmLname         string  `json:"ctm_lname"`
	CtmNname         string  `json:"ctm_nname"`
	CtmFnameEn       string  `json:"ctm_fname_en"`
	CtmLnameEn       string  `json:"ctm_lname_en"`
	CtmGender        string  `json:"ctm_gender"`
	CtmNation        string  `json:"ctm_nation"`
	CtmReligion      string  `json:"ctm_religion"`
	CtmEduLevel      string  `json:"ctm_edu_level"`
	CtmMaritalStatus string  `json:"ctm_marital_status"`
	CtmBlood         string  `json:"ctm_blood"`
	CtmEmail         string  `json:"ctm_email"`
	CtmTel           string  `json:"ctm_tel"`
	CtmTel_2         string  `json:"ctm_tel_2"`
	CtmBirthdate     string  `json:"ctm_birthdate"`
	CtmAddress       string  `json:"ctm_address"`
	CtmDistrict      string  `json:"ctm_district"`
	CtmAmphoe        string  `json:"ctm_amphoe"`
	CtmProvince      string  `json:"ctm_province"`
	CtmZipcode       string  `json:"ctm_zipcode"`
	CtmWeight        float64 `json:"ctm_weight" gorm:"type:decimal(10,2)"`
	CtmHeight        float64 `json:"ctm_height" gorm:"type:decimal(10,2)"`
	CtmWaistline     float64 `json:"ctm_waistline" gorm:"type:decimal(10,2)"`
	CtmChest         float64 `json:"ctm_chest" gorm:"type:decimal(10,2)"`
	CtmTreatmentType int     `json:"ctm_treatment_type"`
	RightTreatmentId int     `json:"right_treatment_id"`
	CtmAllergic      string  `json:"ctm_allergic"`
	CtmMentalHealth  string  `json:"ctm_mental_health"`
	CtmDisease       string  `json:"ctm_disease"`
	CtmComment       string  `json:"ctm_comment"`
	CtmHealthComment string  `json:"ctm_health_comment"`
	CtmImage         string  `json:"ctm_image"`
	CtmImageSize     int     `json:"ctm_image_size"`
	CtmPoint         int     `json:"ctm_point"`
	// CtmCoin             float64           `json:"ctm_coin" gorm:"type:decimal(10,2)"`
	CompanyName         string            `json:"company_name"`
	CompanyTax          string            `json:"company_tax"`
	CompanyTel          string            `json:"company_tel"`
	CompanyEmail        string            `json:"company_email"`
	CompanyAddress      string            `json:"company_address"`
	CompanyDistrict     string            `json:"company_district"`
	CompanyAmphoe       string            `json:"company_amphoe"`
	CompanyProvince     string            `json:"company_province"`
	CompanyZipcode      string            `json:"company_zipcode"`
	CtmSubscribeOpd     int               `json:"ctm_subscribe_opd"`
	CtmSubscribeLab     int               `json:"ctm_subscribe_lab"`
	CtmSubscribeCert    int               `json:"ctm_subscribe_cert"`
	CtmSubscribeReceipt int               `json:"ctm_subscribe_receipt"`
	CtmSubscribeAppoint int               `json:"ctm_subscribe_appoint"`
	CtmIsActive         int               `json:"ctm_is_active"`
	TagSelected         []int             `json:"tag_selected"`
	FamilySelected      []FamilySelected  `json:"family_selected"`
	ContactSelected     []ContactSelected `json:"contact_selected"`
}

type ObjQueryUpdateCustomer struct {
	CustomerGroupId  int     `json:"customer_group_id"`
	CtmCitizenId     string  `json:"ctm_citizen_id"`
	CtmPassportId    string  `json:"ctm_passport_id"`
	CtmPrefix        string  `json:"ctm_prefix"`
	CtmFname         string  `json:"ctm_fname"`
	CtmLname         string  `json:"ctm_lname"`
	CtmNname         string  `json:"ctm_nname"`
	CtmFnameEn       string  `json:"ctm_fname_en"`
	CtmLnameEn       string  `json:"ctm_lname_en"`
	CtmGender        string  `json:"ctm_gender"`
	CtmNation        string  `json:"ctm_nation"`
	CtmReligion      string  `json:"ctm_religion"`
	CtmEduLevel      string  `json:"ctm_edu_level"`
	CtmMaritalStatus string  `json:"ctm_marital_status"`
	CtmBlood         string  `json:"ctm_blood"`
	CtmEmail         string  `json:"ctm_email"`
	CtmTel           string  `json:"ctm_tel"`
	CtmTel_2         string  `json:"ctm_tel_2"`
	CtmBirthdate     string  `json:"ctm_birthdate"`
	CtmAddress       string  `json:"ctm_address"`
	CtmDistrict      string  `json:"ctm_district"`
	CtmAmphoe        string  `json:"ctm_amphoe"`
	CtmProvince      string  `json:"ctm_province"`
	CtmZipcode       string  `json:"ctm_zipcode"`
	CtmWeight        float64 `json:"ctm_weight" gorm:"type:decimal(10,2)"`
	CtmHeight        float64 `json:"ctm_height" gorm:"type:decimal(10,2)"`
	CtmWaistline     float64 `json:"ctm_waistline" gorm:"type:decimal(10,2)"`
	CtmChest         float64 `json:"ctm_chest" gorm:"type:decimal(10,2)"`
	CtmTreatmentType int     `json:"ctm_treatment_type"`
	RightTreatmentId int     `json:"right_treatment_id"`
	CtmAllergic      string  `json:"ctm_allergic"`
	CtmMentalHealth  string  `json:"ctm_mental_health"`
	CtmDisease       string  `json:"ctm_disease"`
	CtmComment       string  `json:"ctm_comment"`
	CtmHealthComment string  `json:"ctm_health_comment"`
	CtmImage         string  `json:"ctm_image"`
	CtmImageSize     int     `json:"ctm_image_size"`
	CtmPoint         int     `json:"ctm_point"`
	// CtmCoin             float64 `json:"ctm_coin" gorm:"type:decimal(10,2)"`
	CompanyName         string `json:"company_name"`
	CompanyTax          string `json:"company_tax"`
	CompanyTel          string `json:"company_tel"`
	CompanyEmail        string `json:"company_email"`
	CompanyAddress      string `json:"company_address"`
	CompanyDistrict     string `json:"company_district"`
	CompanyAmphoe       string `json:"company_amphoe"`
	CompanyProvince     string `json:"company_province"`
	CompanyZipcode      string `json:"company_zipcode"`
	CtmSubscribeOpd     int    `json:"ctm_subscribe_opd"`
	CtmSubscribeLab     int    `json:"ctm_subscribe_lab"`
	CtmSubscribeCert    int    `json:"ctm_subscribe_cert"`
	CtmSubscribeReceipt int    `json:"ctm_subscribe_receipt"`
	CtmSubscribeAppoint int    `json:"ctm_subscribe_appoint"`
	CtmIsActive         int    `json:"ctm_is_active"`
	CtmUpdate           string `json:"ctm_update"`
}

type ObjQueryCustomerDocSetting struct {
	CustomerIdDefault     string `json:"customer_id_default"`
	CustomerNumberDefault int    `json:"customer_number_default"`
	CustomerNumberDigit   int    `json:"customer_number_digit"`
	CustomerType          int    `json:"customer_type"`
}

type FamilySelected struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Relation string `json:"relation"`
}

type ContactSelected struct {
	Name     string `json:"name"`
	Tel      string `json:"tel"`
	Relation string `json:"relation"`
}

type ObjResponseCreateCustomer struct {
	ID int `json:"id"`
}

type ObjPayloadActiveCustomer struct {
	CtmIsActive *int `json:"ctm_is_active"`
}

type ObjQueryActiveCustomer struct {
	CtmIsActive *int
}

type ObjPayloadGetCustomerOpdPagination struct {
	// ShopId int `json:"shop_id" binding:"required"`
	// CustomerId  int     `json:"customer_id" binding:"required"`
	QueueTypeId int     `json:"queue_type_id" binding:"required"`
	SearchText  *string `json:"search_text" binding:"required"`
	CurrentPage int     `json:"current_page" binding:"required"`
	PerPage     int     `json:"per_page" binding:"required"`
}

type ObjQueryGetCustomerOpdPagination struct {
	OpdId        int    `json:"opd_id"`
	OpdCode      string `json:"opd_code"`
	OpdDate      string `json:"opd_date"`
	OpdCreate    string `json:"opd_create"`
	QueId        int    `json:"que_id"`
	QueAdmisId   int    `json:"que_admis_id"`
	QueCode      string `json:"que_code"`
	QueDatetime  string `json:"que_datetime"`
	QueCreate    string `json:"que_create"`
	UserFullname string `json:"user_fullname"`
	ShopId       int    `json:"shop_id"`
	ShopName     string `json:"shop_name"`
}

type ObjResponseGetCustomerOpdPagination struct {
	Items     any   `json:"items"`
	Subs      any   `json:"subs"`
	CountPage int   `json:"count_page"`
	CountAll  int64 `json:"count_all"`
}

type ObjPayloadGetCustomerCheckPagination struct {
	ShopId      int     `json:"shop_id" binding:"required"`
	CustomerId  int     `json:"customer_id" binding:"required"`
	SearchText  *string `json:"search_text" binding:"required"`
	CurrentPage int     `json:"current_page" binding:"required"`
	PerPage     int     `json:"per_page" binding:"required"`
}

type ObjQueryGetCustomerCheckPagination struct {
	ID              int    `json:"id"`
	ShopId          int    `json:"shop_id"`
	ReceiptId       int    `json:"receipt_id"`
	ReceiptDetailId int    `json:"receipt_detail_id"`
	UserId          int    `json:"user_id"`
	CustomerId      int    `json:"customer_id"`
	QueueId         int    `json:"queue_id"`
	CheckingId      int    `json:"checking_id"`
	ChkTypeId       int    `json:"chk_type_id"`
	ChkCode         string `json:"chk_code"`
	ChkName         string `json:"chk_name"`
	ChkUnit         string `json:"chk_unit"`
	ChkValue        string `json:"chk_value"`
	ChkUpload       string `json:"chk_upload"`
	ChkUploadSize   int    `json:"chk_upload_size"`
	ChkOld          string `json:"chk_old"`
	DirectionId     int    `json:"direction_id"`
	ChkFlag         string `json:"chk_flag"`
	ChkDate         string `json:"chk_date"`
	ChkIsPrint      int    `json:"chk_is_print"`
	ChkIsReport     int    `json:"chk_is_report"`
	ChkIsActive     int    `json:"chk_is_active"`
	ChkDatetime     string `json:"chk_datetime"`
	ChkCreate       string `json:"chk_create"`
	ChkUpdate       string `json:"chk_update"`
	DirectionName   string `json:"direction_name"`
	QueCode         string `json:"que_code"`
	QueShopId       int    `json:"que_shop_id"`
	ShopName        string `json:"shop_name"`
	QueDatetime     string `json:"que_datetime"`
	RecCode         string `json:"rec_code"`
	UserFullname    string `json:"user_fullname"`
}

type ObjResponseGetCustomerCheckPagination struct {
	Items     any   `json:"items"`
	CountPage int   `json:"count_page"`
	CountAll  int64 `json:"count_all"`
}

type ObjPayloadGetCustomerServicePagination struct {
	ShopId      int     `json:"shop_id" binding:"required"`
	CustomerId  int     `json:"customer_id" binding:"required"`
	SearchText  *string `json:"search_text" binding:"required"`
	CurrentPage int     `json:"current_page" binding:"required"`
	PerPage     int     `json:"per_page" binding:"required"`
}

type ObjQueryGetCustomerServicePagination struct {
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
	UserFullname    string `json:"user_fullname"`
}

type ObjResponseGetCustomerServicePagination struct {
	Items     any   `json:"items"`
	CountPage int   `json:"count_page"`
	CountAll  int64 `json:"count_all"`
}

type ObjPayloadSearchReceiptCustomer struct {
	Search        *string `json:"search" binding:"required,omitempty"`
	Rec_is_active *string `json:"rec_is_active" binding:"required,omitempty"`
	// Customer_id   *string `json:"customer_id" binding:"required,omitempty"`
	// Shop_id    int `json:"shop_id"`
	ActivePage int `json:"active_page" binding:"required"`
	PerPage    int `json:"per_page" binding:"required"`
}

type ReceiptListCustomer struct {
	Id                int     `json:"id"`
	Shop_id           int     `json:"shop_id"`
	Customer_shop_id  int     `json:"customer_shop_id"`
	Shop_name         string  `json:"shop_name"`
	User_id           int     `json:"user_id"`
	User_fullname     string  `json:"user_fullname"`
	Customer_id       int     `json:"customer_id"`
	Ctm_id            string  `json:"ctm_id"`
	Queue_id          *int    `json:"queue_id"`
	Invoice_id        int     `json:"invoice_id"`
	Inv_code          string  `json:"inv_code"`
	Rec_code          string  `json:"rec_code"`
	Rec_fullname      string  `json:"rec_fullname"`
	Rec_tel           string  `json:"rec_tel"`
	Rec_email         string  `json:"rec_email"`
	Rec_address       string  `json:"rec_address"`
	Rec_district      string  `json:"rec_district"`
	Rec_amphoe        string  `json:"rec_amphoe"`
	Rec_province      string  `json:"rec_province"`
	Rec_zipcode       string  `json:"rec_zipcode"`
	Rec_comment       string  `json:"rec_comment"`
	Rec_total_price   float64 `json:"rec_total_price"`
	Rec_discount      float64 `json:"rec_discount"`
	Rec_befor_vat     float64 `json:"rec_befor_vat"`
	Tax_type_id       int     `json:"tax_type_id"`
	Tax_rate          int     `json:"tax_rate"`
	Rec_vat           float64 `json:"rec_vat"`
	Rec_total         float64 `json:"rec_total"`
	Rec_payment_type  int     `json:"rec_payment_type"`
	Rec_type_id       int     `json:"rec_type_id"`
	Rec_period        int     `json:"rec_period"`
	Rec_pay           float64 `json:"rec_pay"`
	Rec_balance       float64 `json:"rec_balance"`
	Rec_pay_total     float64 `json:"rec_pay_total"`
	Rec_discription   string  `json:"rec_discription"`
	Rec_account       int     `json:"rec_account"`
	Rec_is_process    int     `json:"rec_is_process"`
	Rec_is_active     int     `json:"rec_is_active"`
	Rec_user_id       int     `json:"rec_user_id"`
	Rec_user_fullname string  `json:"rec_user_fullname"`
	Rec_datetime      string  `json:"rec_datetime"`
	Rec_create        string  `json:"rec_create"`
	Rec_update        string  `json:"rec_update"`
	Rec_is_cancel     int     `json:"rec_is_cancel"`
	RecPayDatetime    string  `json:"rec_pay_datetime"`
}

type ResponsePaginationReceiptCustomer struct {
	Result_data   []ReceiptListCustomer `json:"result_data"`
	Count_of_page int                   `json:"count_of_page"`
	Count_all     int                   `json:"count_all"`
}

type ObjCheckExcelCustomer struct {
	ShopId           int     `json:"shop_id"`
	ShopMotherId     int     `json:"shop_mother_id"`
	CustomerGroupId  int     `json:"customer_group_id"`
	CtmId            string  `json:"ctm_id"`
	CtmCitizenId     string  `json:"ctm_citizen_id"`
	CtmPassportId    string  `json:"ctm_passport_id"`
	CtmPrefix        string  `json:"ctm_prefix"`
	CtmFname         string  `json:"ctm_fname"`
	CtmLname         string  `json:"ctm_lname"`
	CtmNname         string  `json:"ctm_nname"`
	CtmFnameEn       string  `json:"ctm_fname_en"`
	CtmLnameEn       string  `json:"ctm_lname_en"`
	CtmGender        string  `json:"ctm_gender"`
	CtmNation        string  `json:"ctm_nation"`
	CtmReligion      string  `json:"ctm_religion"`
	CtmEduLevel      string  `json:"ctm_edu_level"`
	CtmMaritalStatus string  `json:"ctm_marital_status"`
	CtmBlood         string  `json:"ctm_blood"`
	CtmEmail         string  `json:"ctm_email"`
	CtmTel           string  `json:"ctm_tel"`
	CtmTel_2         string  `json:"ctm_tel_2"`
	CtmBirthdate     string  `json:"ctm_birthdate"`
	CtmAddress       string  `json:"ctm_address"`
	CtmDistrict      string  `json:"ctm_district"`
	CtmAmphoe        string  `json:"ctm_amphoe"`
	CtmProvince      string  `json:"ctm_province"`
	CtmZipcode       string  `json:"ctm_zipcode"`
	CtmComment       string  `json:"ctm_comment"`
	CtmWeight        float64 `json:"ctm_weight" gorm:"type:decimal(10,2)"`
	CtmHeight        float64 `json:"ctm_height" gorm:"type:decimal(10,2)"`
	CtmWaistline     float64 `json:"ctm_waistline" gorm:"type:decimal(10,2)"`
	CtmChest         float64 `json:"ctm_chest" gorm:"type:decimal(10,2)"`
	// CtmTreatmentType int     `json:"ctm_treatment_type"`
	// RightTreatmentId int     `json:"right_treatment_id"`
	CtmAllergic      string   `json:"ctm_allergic"`
	CtmMentalHealth  string   `json:"ctm_mental_health"`
	CtmDisease       string   `json:"ctm_disease"`
	CtmHealthComment string   `json:"ctm_health_comment"`
	CompanyName      string   `json:"company_name"`
	CompanyTax       string   `json:"company_tax"`
	CompanyTel       string   `json:"company_tel"`
	CompanyEmail     string   `json:"company_email"`
	CompanyAddress   string   `json:"company_address"`
	CompanyDistrict  string   `json:"company_district"`
	CompanyAmphoe    string   `json:"company_amphoe"`
	CompanyProvince  string   `json:"company_province"`
	CompanyZipcode   string   `json:"company_zipcode"`
	Message          []string `json:"message"`
}

type ObjPayloadImportExcelCustomer struct {
	ImportData []ObjCheckExcelCustomer `json:"import_data"`
}

type ObjPayloadAcceptPDPA struct {
	CustomerId    int    `json:"customer_id" binding:"required"`
	CustomerImage string `json:"customer_image" binding:"required"`
}

type ObjPayloadSmsPDPA struct {
	CustomerId int    `json:"customer_id" binding:"required"`
	CtmTel     string `json:"ctm_tel" binding:"required"`
}

type LogCustomer struct {
	Username   string `json:"username"`
	Log_type   string `json:"log_type"`
	Log_text   string `json:"log_text"`
	Log_create string `json:"log_create"`
}

type PayloadSearchAppointmentByCustomer struct {
	// Shop_id int `json:"shop_id"`
	// Customer_id int     `json:"customer_id" binding:"required"`
	Search     *string `json:"search" binding:"required,omitempty"`
	Date       *string `json:"date" binding:"required,omitempty"`
	Type       *string `json:"type" binding:"required,omitempty"`
	Is_active  *string `json:"is_active" binding:"required,omitempty"`
	ActivePage int     `json:"active_page" binding:"required"`
	PerPage    int     `json:"per_page" binding:"required"`
	Date_from  *string `json:"date_from"`
}

type CheckCustomer struct {
	ID       int    `json:"id"`
	ShopCode string `json:"shop_code"`
	ShopName string `json:"shop_name"`
}

type ObjPayloadGetCustomerCheckLabXaryPagination struct {
	// ShopMotherId int `json:"shop_mother_id" binding:"required"`
	// CustomerId   int     `json:"customer_id" binding:"required"`
	SearchDate  *string `json:"search_date" binding:"required"`
	SearchText  *string `json:"search_text" binding:"required"`
	CurrentPage int     `json:"current_page" binding:"required"`
	PerPage     int     `json:"per_page" binding:"required"`
	ShopId      *int    `json:"shop_id"`
}

type PayloadSearchAppointmentByCustomerHistory struct {
	ShopMotherId int     `json:"shop_mother_id" binding:"required"`
	Shop_id      int     `json:"shop_id"`
	Customer_id  int     `json:"customer_id" binding:"required"`
	Search       *string `json:"search" binding:"required,omitempty"`
	Date         *string `json:"date" binding:"required,omitempty"`
	Type         *string `json:"type" binding:"required,omitempty"`
	ActivePage   int     `json:"active_page" binding:"required"`
	PerPage      int     `json:"per_page" binding:"required"`
}

type ObjPayloadPaginationHistory struct {
	// CustomerId   int     `json:"customer_id" binding:"required"`
	// ShopMotherId int     `json:"shop_mother_id" binding:"required"`
	SearchText  *string `json:"search_text" binding:"required"`
	SearchDate  *string `json:"search_date" binding:"required"`
	CurrentPage int     `json:"current_page" binding:"required"`
	PerPage     int     `json:"per_page" binding:"required"`
}

type ObjQueryGetCustomerCourseHistory struct {
	Que_code         string  `json:"que_code"`
	Rec_code         string  `json:"rec_code"`
	Seru_code        string  `json:"seru_code"`
	Seru_name        string  `json:"seru_name"`
	Seru_qty         float64 `json:"seru_qty"`
	Ser_price        float64 `json:"ser_price"`
	User_fullname    string  `json:"user_fullname"`
	User_fullname_en string  `json:"user_fullname_en"`
	Seru_date        string  `json:"seru_date"`
}

type ObjQueryGetCustomerMedicineHistory struct {
	Que_code         string  `json:"que_code"`
	Rec_code         string  `json:"rec_code"`
	Quep_code        string  `json:"quep_code"`
	Quep_name        string  `json:"quep_name"`
	Quep_qty         float64 `json:"quep_qty"`
	Quep_total       float64 `json:"quep_total"`
	Quep_topical     string  `json:"quep_topical"`
	User_fullname    string  `json:"user_fullname"`
	User_fullname_en string  `json:"user_fullname_en"`
	Que_datetime     string  `json:"que_datetime"`
}

type ResponseAppointmentHistory struct {
	Result_data   []AppointmentListHistory `json:"result_data"`
	Count_of_page int                      `json:"count_of_page"`
	Count_all     int                      `json:"count_all"`
}

type AppointmentListHistory struct {
	ID               int    `json:"id"`
	ShopID           int    `json:"shop_id"`
	UserID           int    `json:"user_id"`
	User_image       string `json:"user_image"`
	UserFullname     string `json:"user_fullname"`
	UserFullnameEn   string `json:"user_fullname_en"`
	RoleNameTh       string `json:"role_name_th"`
	RoleNameEn       string `json:"role_name_en"`
	CustomerID       int    `json:"customer_id"`
	CustomerFullname string `json:"customer_fullname"`
	CtmFname         string `json:"ctm_fname"`
	CtmLname         string `json:"ctm_lname"`
	CtmFnameEn       string `json:"ctm_fname_en"`
	CtmLnameEn       string `json:"ctm_lname_en"`
	ApType           int    `json:"ap_type"`
	ApTopic          string `json:"ap_topic"`
	ApTel            string `json:"ap_tel"`
	ApDatetime       string `json:"ap_datetime"`
	ApNote           string `json:"ap_note"`
	ApComment        string `json:"ap_comment"`
	ApColor          string `json:"ap_color"`
	ApConfirm        int    `json:"ap_confirm"`
	ApStatusID       int    `json:"ap_status_id"`
	ApStatusSMS      int    `json:"ap_status_sms"`
	ApStatusLine     int    `json:"ap_status_line"`
	ApSms            string `json:"ap_sms"`
	ApIsGcalendar    int    `json:"ap_is_gcalendar"`
	ApGid            string `json:"ap_gid"`
	ApUserID         int    `json:"ap_user_id"`
	ApIsDel          int    `json:"ap_is_del"`
	ApCreate         string `json:"ap_create"`
	ApUpdate         string `json:"ap_update"`
	ShopName         string `json:"shop_name"`
}

// ---------------------- test ObjGetCustomerOnline-------------------------/-
type ObjGetCustomerOnlineTest struct {
	ID          int    `json:"id"`
	CoCitizenId string `json:"co_citizen_id"`
	CoEmail     string `json:"co_email"`
	CoFname     string `json:"co_fname"`
	CoLname     string `json:"co_lname"`
	Co_Line_id  string `json:"co_line_id"`
	CoGender    string `json:"co_gender"`
	CoBirthdate string `json:"co_birthdate"`
	CoAddress   string `json:"co_address"`
	CoTel       string `json:"co_tel"`
	CoDistrict  string `json:"co_district"`
	CoProvince  string `json:"co_province"`
	CoAmphoe    string `json:"co_amphoe"`
	CoZipcode   string `json:"co_zipcode"`
}
