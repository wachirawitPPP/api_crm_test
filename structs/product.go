package structs

type ObjPayloadSearch struct {
	Search  *string `json:"search" binding:"required,omitempty"`
	Shop_id int     `json:"shop_id"`
}

type ResponsePaginationProduct struct {
	Result_data   []ProductList `json:"result_data"`
	Count_of_page int           `json:"count_of_page"`
	Count_all     int           `json:"count_all"`
}

type PayloadSearchProduct struct {
	Search     *string `json:"search" binding:"required,omitempty"`
	Is_active  *string `json:"is_active" binding:"required,omitempty"`
	CategoryId int     `json:"category_id"`
	PdTypeId   int     `json:"pd_type_id"`
	Shop_id    int     `json:"shop_id"`
	Shop_m_id  int     `json:"shop_m_id"`
	ActivePage int     `json:"active_page" binding:"required"`
	PerPage    int     `json:"per_page" binding:"required"`
}

type ProductList struct {
	Id               int    `json:"id"`
	ShopId           int    `json:"shop_id"`
	CategoryId       int    `json:"category_id"`
	PdTypeId         int    `json:"pd_type_id"`
	PdCode           string `json:"pd_code"`
	PdName           string `json:"pd_name"`
	PdCodeAcc        string `json:"pd_code_acc"`
	PdNameAcc        string `json:"pd_name_acc"`
	PdImage          string `json:"pd_image"`
	PdDescription    string `json:"pd_description"`
	PdBarcode        string `json:"pd_barcode"`
	PdIsSerial       int    `json:"pd_is_serial"`
	PdNarcotic2      int    `json:"pd_narcotic2"`
	PdNarcotic3      int    `json:"pd_narcotic3"`
	PdNarcotic4      int    `json:"pd_narcotic4"`
	PdNarcotic5      int    `json:"pd_narcotic5"`
	CategoryEclaimId *int   `json:"category_eclaim_id"`
	DrugDirection    string `json:"drug_direction"`
	TopicalId        *int   `json:"topical_id"`
	UserId           *int   `json:"user_id"`
	PdAmountNoti     int    `json:"pd_amount_noti"`
	PdExpireNoti     string `json:"pd_expire_noti"`
	PdIsOver         int    `json:"pd_is_over"`
	PdIsActive       int    `json:"pd_is_active"`
	PdIsDel          int    `json:"pd_is_del"`
	PdCreate         string `json:"pd_create"`
	PdUpdate         string `json:"pd_update"`
	Category
}

type GetProductAll struct {
	Id int `json:"id"`
}

type ProductByID struct {
	Id               int    `json:"id" binding:"required"`
	ShopId           int    `json:"shop_id"`
	CategoryId       int    `json:"category_id"`
	PdTypeId         int    `json:"pd_type_id"`
	PdCode           string `json:"pd_code"`
	PdName           string `json:"pd_name"`
	PdCodeAcc        string `json:"pd_code_acc"`
	PdNameAcc        string `json:"pd_name_acc"`
	PdImage          string `json:"pd_image"`
	PdDescription    string `json:"pd_description"`
	PdBarcode        string `json:"pd_barcode"`
	PdIsSerial       int    `json:"pd_is_serial"`
	PdNarcotic2      int    `json:"pd_narcotic2"`
	PdNarcotic3      int    `json:"pd_narcotic3"`
	PdNarcotic4      int    `json:"pd_narcotic4"`
	PdNarcotic5      int    `json:"pd_narcotic5"`
	CategoryEclaimId int    `json:"category_eclaim_id"`
	DrugDirection    string `json:"drug_direction"`
	PdAmountNoti     int    `json:"pd_amount_noti"`
	PdExpireNoti     string `json:"pd_expire_noti"`
	PdIsOver         int    `json:"pd_is_over"`
	PdIsActive       int    `json:"pd_is_active"`
	PdIsDel          int    `json:"pd_is_del"`
	PdCreate         string `json:"pd_create"`
	PdUpdate         string `json:"pd_update"`
}

type OdjProductStore struct {
	ID          int     `json:"id"`
	ShopId      int     `json:"shop_id"`
	ShopStoreId int     `json:"shop_store_id"`
	ProductId   int     `json:"product_id"`
	PdsBarcode  string  `json:"pds_barcode"`
	PdsCost     float64 `json:"pds_cost" gorm:"type:decimal(10,2)"`
	PdsIn       float64 `json:"pds_in" gorm:"type:decimal(10,2)"`
	PdsOut      float64 `json:"pds_out" gorm:"type:decimal(10,2)"`
	PdsExp      float64 `json:"pds_exp" gorm:"type:decimal(10,2)"`
	PdsTotal    float64 `json:"pds_total" gorm:"type:decimal(10,2)"`
	PdsDate     string  `json:"pds_date"`
	PdsComment  string  `json:"pds_comment"`
	PdsIsActive int     `json:"pds_is_active"`
	PdsIsDel    int     `json:"pds_is_del"`
	PdsCreate   string  `json:"pds_create"`
	PdsUpdate   string  `json:"pds_update"`
}

type ObjPayloadSearchProductUnit struct {
	Search     *string `json:"search" binding:"required,omitempty"`
	ActivePage int     `json:"active_page" binding:"required"`
	PerPage    int     `json:"per_page" binding:"required"`
}

type ResponsePaginationProductUnit struct {
	Result_data   []Product `json:"result_data"`
	Count_of_page int       `json:"count_of_page"`
	Count_all     int       `json:"count_all"`
}

type ProductDetail struct {
	Id              int    `json:"id"`
	ShopId          int    `json:"shop_id"`
	CategoryId      int    `json:"category_id"`
	CategoryName    string `json:"category_name"`
	PdTypeId        int    `json:"pd_type_id"`
	PdCode          string `json:"pd_code"`
	PdName          string `json:"pd_name"`
	PdCodeAcc       string `json:"pd_code_acc"`
	PdNameAcc       string `json:"pd_name_acc"`
	PdImage         string `json:"pd_image"`
	PdDescription   string `json:"pd_description"`
	PdBarcode       string `json:"pd_barcode"`
	PdIsSerial      int    `json:"pd_is_serial"`
	PdNarcotic2     int    `json:"pd_narcotic2"`
	PdNarcotic3     int    `json:"pd_narcotic3"`
	PdNarcotic4     int    `json:"pd_narcotic4"`
	PdNarcotic5     int    `json:"pd_narcotic5"`
	DrugDirection   string `json:"drug_direction"`
	TopicalId       *int   `json:"topical_id"`
	TopicalName     string `json:"topical_name"`
	UserId          *int   `json:"user_id"`
	UserFullname    string `json:"user_fullname"`
	UserFullname_en string `json:"user_fullname_en"`
	PdAmountNoti    int    `json:"pd_amount_noti"`
	PdExpireNoti    int    `json:"pd_expire_noti"`
	PdIsOver        int    `json:"pd_is_over"`
	PdIsActive      int    `json:"pd_is_active"`
	PdIsDel         int    `json:"pd_is_del"`
	PdCreate        string `json:"pd_create"`
	PdUpdate        string `json:"pd_update"`
	Category
	ProductUnitList *[]ProductUnitList `json:"product_unit" gorm:"-" binding:"omitempty"`
}

type ObjPayloadCreateProduct struct {
	CategoryId    int               `json:"category_id" binding:"required"`
	PdTypeId      int               `json:"pd_type_id" binding:"required"`
	PdCode        string            `json:"pd_code" binding:"required"`
	PdName        string            `json:"pd_name" binding:"required"`
	PdCodeAcc     string            `json:"pd_code_acc"`
	PdNameAcc     string            `json:"pd_name_acc"`
	PdImage       string            `json:"pd_image"`
	PdDescription string            `json:"pd_description"`
	PdBarcode     string            `json:"pd_barcode"`
	PdIsSerial    int               `json:"pd_is_serial"`
	PdNarcotic2   int               `json:"pd_narcotic2"`
	PdNarcotic3   int               `json:"pd_narcotic3"`
	PdNarcotic4   int               `json:"pd_narcotic4"`
	PdNarcotic5   int               `json:"pd_narcotic5"`
	DrugDirection string            `json:"drug_direction"`
	TopicalId     *int              `json:"topical_id"`
	UserId        *int              `json:"user_id"`
	PdAmountNoti  int               `json:"pd_amount_noti"`
	PdExpireNoti  int               `json:"pd_expire_noti"`
	PdIsOver      int               `json:"pd_is_over"`
	PdIsActive    int               `json:"pd_is_active"`
	PdIsDel       int               `json:"pd_is_del"`
	ProductUnit   *[]ObjProductUnit `json:"product_unit" gorm:"-" binding:"required,omitempty"`
}

type ObjPayloadUpdateProduct struct {
	Id            int               `json:"id" binding:"required"`
	Shop_id       int               `json:"shop_id" binding:"required"`
	CategoryId    int               `json:"category_id" binding:"required"`
	PdTypeId      int               `json:"pd_type_id" binding:"required"`
	PdCode        string            `json:"pd_code"`
	PdName        string            `json:"pd_name"`
	PdCodeAcc     string            `json:"pd_code_acc"`
	PdNameAcc     string            `json:"pd_name_acc"`
	PdImage       string            `json:"pd_image"`
	PdDescription string            `json:"pd_description"`
	PdBarcode     string            `json:"pd_barcode"`
	PdIsSerial    int               `json:"pd_is_serial"`
	PdNarcotic2   int               `json:"pd_narcotic2"`
	PdNarcotic3   int               `json:"pd_narcotic3"`
	PdNarcotic4   int               `json:"pd_narcotic4"`
	PdNarcotic5   int               `json:"pd_narcotic5"`
	TopicalId     *int              `json:"topical_id"`
	UserId        *int              `json:"user_id"`
	DrugDirection string            `json:"drug_direction"`
	PdAmountNoti  int               `json:"pd_amount_noti"`
	PdExpireNoti  int               `json:"pd_expire_noti"`
	PdIsOver      int               `json:"pd_is_over"`
	PdIsActive    int               `json:"pd_is_active"`
	PdIsDel       int               `json:"pd_is_del"`
	ProductUnit   *[]ObjProductUnit `json:"product_unit" gorm:"-" binding:"required,omitempty"`
}

type ProductAction struct {
	Id            int    `json:"id" binding:"required"`
	ShopId        int    `json:"shop_id" binding:"required"`
	CategoryId    int    `json:"category_id" binding:"required"`
	PdTypeId      int    `json:"pd_type_id" binding:"required"`
	PdCode        string `json:"pd_code" binding:"required"`
	PdName        string `json:"pd_name" binding:"required"`
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
	DrugDirection string `json:"drug_direction"`
	PdAmountNoti  int    `json:"pd_amount_noti"`
	PdExpireNoti  int    `json:"pd_expire_noti"`
	PdIsOver      int    `json:"pd_is_over"`
	PdIsActive    int    `json:"pd_is_active"`
	PdIsDel       int    `json:"pd_is_del"`
	PdCreate      string `json:"pd_create"`
	PdUpdate      string `json:"pd_update"`
}

type ProductShopAction struct {
	Id        int    `json:"id"`
	ProductId int    `json:"product_id" binding:"required"`
	ShopId    int    `json:"shop_id" binding:"required"`
	UserId    *int   `json:"user_id" gorm:"default:null"`
	TopicalId *int   `json:"topical_id" gorm:"default:null"`
	PsCreate  string `json:"ps_create"`
	PsUpdate  string `json:"ps_update"`
}

type ProductUnitAction struct {
	Id        int    `json:"id"`
	ProductId int    `json:"product_id" binding:"required"`
	UnitId    int    `json:"unit_id" binding:"required"`
	PuRate    int    `json:"pu_rate"`
	PuIsDel   int    `json:"pu_is_del"`
	PuCreate  string `json:"pu_create"`
	PuUpdate  string `json:"pu_update"`
}

type ProductShopPriceAction struct {
	Id               int     `json:"id"`
	ShopId           int     `json:"shop_id" binding:"required"`
	ProductUnitId    int     `json:"product_unit_id" binding:"required"`
	CategoryEclaimId *int    `json:"category_eclaim_id"`
	PspPriceOpd      float64 `json:"psp_price_opd" gorm:"type:decimal(10,2)"`
	PspPriceIpd      float64 `json:"psp_price_ipd" gorm:"type:decimal(10,2)"`
	PspPriceOfc      float64 `json:"psp_price_ofc" gorm:"type:decimal(10,2)"`
	PspPriceLgo      float64 `json:"psp_price_lgo" gorm:"type:decimal(10,2)"`
	PspPriceUcs      float64 `json:"psp_price_ucs" gorm:"type:decimal(10,2)"`
	PspPriceSss      float64 `json:"psp_price_sss" gorm:"type:decimal(10,2)"`
	PspPriceNhs      float64 `json:"psp_price_nhs" gorm:"type:decimal(10,2)"`
	PspPriceSsi      float64 `json:"psp_price_ssi" gorm:"type:decimal(10,2)"`
	PspIsDefault     int     `json:"psp_is_default"`
	PspIsDel         int     `json:"psp_is_del"`
	PspCreate        string  `json:"psp_create"`
	PspUpdate        string  `json:"psp_update"`
}

type ObjProductUnit struct {
	Id               int      `json:"id"`
	ProductId        int      `json:"product_id"`
	UnitId           int      `json:"unit_id"`
	UnitName         string   `json:"u_name"`
	PuRate           int      `json:"pu_rate"`
	PuIsDel          int      `json:"pu_is_del"`
	PuCreate         string   `json:"pu_create"`
	PuUpdate         string   `json:"pu_update"`
	CategoryEclaimId *int     `json:"category_eclaim_id"`
	PspPriceOpd      *float64 `json:"psp_price_opd" gorm:"type:decimal(10,2)"`
	PspPriceIpd      *float64 `json:"psp_price_ipd" gorm:"type:decimal(10,2)"`
	PspPriceOfc      *float64 `json:"psp_price_ofc" gorm:"type:decimal(10,2)"`
	PspPriceLgo      *float64 `json:"psp_price_lgo" gorm:"type:decimal(10,2)"`
	PspPriceUcs      *float64 `json:"psp_price_ucs" gorm:"type:decimal(10,2)"`
	PspPriceSss      *float64 `json:"psp_price_sss" gorm:"type:decimal(10,2)"`
	PspPriceNhs      *float64 `json:"psp_price_nhs" gorm:"type:decimal(10,2)"`
	PspPriceSsi      *float64 `json:"psp_price_ssi" gorm:"type:decimal(10,2)"`
}

type GetShopByMotherId struct {
	ShopId int `json:"shop_id"`
}

type ProductUnitList struct {
	Id               int     `json:"id"`
	Product_units_id int     `json:"product_units_id"`
	ProductId        int     `json:"product_id"`
	UnitId           int     `json:"unit_id"`
	UName            string  `json:"u_name"`
	UNameEn          string  `json:"u_name_en"`
	PuRate           int     `json:"pu_rate"`
	PuIsDel          int     `json:"pu_is_del"`
	PuCreate         string  `json:"pu_create"`
	PuUpdate         string  `json:"pu_update"`
	CategoryEclaimId int     `json:"category_eclaim_id"`
	PspPriceOpd      float64 `json:"psp_price_opd" gorm:"type:decimal(10,2)"`
	PspPriceIpd      float64 `json:"psp_price_ipd" gorm:"type:decimal(10,2)"`
	PspPriceOfc      float64 `json:"psp_price_ofc" gorm:"type:decimal(10,2)"`
	PspPriceLgo      float64 `json:"psp_price_lgo" gorm:"type:decimal(10,2)"`
	PspPriceUcs      float64 `json:"psp_price_ucs" gorm:"type:decimal(10,2)"`
	PspPriceSss      float64 `json:"psp_price_sss" gorm:"type:decimal(10,2)"`
	PspPriceNhs      float64 `json:"psp_price_nhs" gorm:"type:decimal(10,2)"`
	PspPriceSsi      float64 `json:"psp_price_ssi" gorm:"type:decimal(10,2)"`
}

type ProductShop struct {
	Id              int    `json:"id"`
	ProductId       int    `json:"product_id"`
	ShopId          int    `json:"shop_id"`
	TopicalId       int    `json:"topical_id"`
	TopicalName     string `json:"topical_name"`
	UserId          int    `json:"user_id"`
	UserFullname    string `json:"user_fullname"`
	UserFullname_en string `json:"user_fullname_en"`
	PsCreate        string `json:"ps_create"`
	PsUpdate        string `json:"ps_update"`
}

type ProductShopPrice struct {
	Id               int     `json:"id"`
	ShopId           int     `json:"shop_id"`
	Product_units_id int     `json:"product_units_id"`
	PspPriceOpd      float64 `json:"psp_price_opd" gorm:"type:decimal(10,2)"`
	PspPriceIpd      float64 `json:"psp_price_ipd" gorm:"type:decimal(10,2)"`
	PspPriceOfc      float64 `json:"psp_price_ofc" gorm:"type:decimal(10,2)"`
	PspPriceLgo      float64 `json:"psp_price_lgo" gorm:"type:decimal(10,2)"`
	PspPriceUcs      float64 `json:"psp_price_ucs" gorm:"type:decimal(10,2)"`
	PspPriceSss      float64 `json:"psp_price_sss" gorm:"type:decimal(10,2)"`
	PspPriceNhs      float64 `json:"psp_price_nhs" gorm:"type:decimal(10,2)"`
	PspPriceSsi      float64 `json:"psp_price_ssi" gorm:"type:decimal(10,2)"`
	PspIsDefault     int     `json:"psp_is_default"`
	PspIsDel         int     `json:"psp_is_del"`
	PspCreate        string  `json:"psp_create"`
	PspUpdate        string  `json:"psp_update"`
}

type PayloadSearchUnit struct {
	Search *string `json:"search"`
}

type Unit struct {
	Id      int    `json:"id"`
	UName   string `json:"u_name"`
	UNameEn string `json:"u_name_en"`
	USort   int    `json:"u_sort"`
	UIsDel  int    `json:"u_is_del"`
}

type UserProduct struct {
	Id             int    `json:"id"`
	UserEmail      string `json:"user_email"`
	UserFullname   string `json:"user_fullname"`
	UserFullnameEn string `json:"user_fullname_en"`
	UserTel        string `json:"user_tel"`
}

type LogProduct struct {
	Username   string `json:"username"`
	Log_type   string `json:"log_type"`
	Log_text   string `json:"log_text"`
	Log_create string `json:"log_create"`
}

type DocNoProduct struct {
	ShopId               int    `json:"shop_id"`
	ProductIdDefault     string `json:"product_id_default"`
	ProductNumberDefault string `json:"product_number_default"`
	ProductNumberDigit   int    `json:"product_number_digit"`
	ProductType          int    `json:"product_type"`
}

type ProductTopical struct {
	Id            int    `json:"id"`
	TopicalName   string `json:"topical_name"`
	TopicalDetail string `json:"topical_detail"`
}

type ObjQueryProduct struct {
	ID               int    `json:"id"`
	ShopId           int    `json:"shop_id"`
	CategoryId       int    `json:"category_id"`
	PdTypeId         int    `json:"pd_type_id"`
	PdCode           string `json:"pd_code"`
	PdName           string `json:"pd_name"`
	PdCodeAcc        string `json:"pd_code_acc"`
	PdNameAcc        string `json:"pd_name_acc"`
	PdImage          string `json:"pd_image"`
	PdDescription    string `json:"pd_description"`
	UserId           int    `json:"user_id"`
	PdBarcode        string `json:"pd_barcode"`
	PdIsSerial       int    `json:"pd_is_serial"`
	PdNarcotic2      int    `json:"pd_narcotic2"`
	PdNarcotic3      int    `json:"pd_narcotic3"`
	PdNarcotic4      int    `json:"pd_narcotic4"`
	PdNarcotic5      int    `json:"pd_narcotic5"`
	CategoryEclaimId int    `json:"category_eclaim_id"`
	DrugDirection    string `json:"checking_code"`
	PdAmountNoti     int    `json:"pd_amount_noti"`
	PdExpireNoti     int    `json:"pd_expire_noti"`
	PdIsOver         int    `json:"pd_is_over"`
	TopicalId        int    `json:"topical_id"`
	PdIsActive       int    `json:"pd_is_active"`
	PdIsDel          int    `json:"pd_is_del"`
	PdCreate         string `json:"pd_create"`
	PdUpdate         string `json:"pd_update"`
}

type ObjCheckExcelProduct struct {
	ShopId        int      `json:"shop_id"`
	CategoryId    int      `json:"category_id"`
	PdCode        string   `json:"pd_code"`
	PdBarcode     string   `json:"pd_barcode"`
	PdName        string   `json:"pd_name"`
	PdCodeAcc     string   `json:"pd_code_acc"`
	PdNameAcc     string   `json:"pd_name_acc"`
	PdAmountNoti  int      `json:"pd_amount_noti"`
	PdExpireNoti  int      `json:"pd_expire_noti"`
	DrugDirection string   `json:"checking_code"`
	PdTypeId      int      `json:"pd_type_id"`
	UnitId        int      `json:"unit_id"`
	UnitName      string   `json:"unit_name"`
	OpdPrice      float64  `json:"opd_price"`
	IpdPrice      float64  `json:"ipd_price"`
	Message       []string `json:"message"`
}

type ObjPayloadImportExcelProduct struct {
	ImportData []ObjCheckExcelProduct `json:"import_data"`
}

type ObjQueryCheckImportProduct struct {
	PdCode    string `json:"pd_code"`
	PdBarcode string `json:"pd_barcode"`
}

type ObjQueryCheckUnit struct {
	ID       int    `json:"id"`
	U_Name   string `json:"u_name"`
	U_NameEn string `json:"u_name_en"`
}

type ObjCreateProductUnit struct {
	ID        int    `json:"id"`
	ProductId int    `json:"product_id"`
	UnitId    int    `json:"unit_id"`
	PuRate    int    `json:"pu_rate"`
	PuIsDel   int    `json:"pu_is_del"`
	PuCreate  string `json:"pu_create"`
	PuUpdate  string `json:"pu_update"`
}
