package structs

type PayloadShopOauth struct {
	Shop_id int `json:"shop_id" binding:"required"`
}

type InShopList struct {
	Id           int    `json:"id"`
	Code         string `json:"code"`
	Name         string `json:"name"`
	NatureType   string `json:"nature_type"`
	Nature       string `json:"nature"`
	Province     string `json:"province"`
	Latlong      string `json:"latlong"`
	Image        string `json:"image"`
	ShopMotherId int    `json:"shop_mother_id"`
}

// get
type ShopReadResponse struct {
	// primary key
	Id int `json:"shop_id"`
	// foreign key
	ShopNatureId  int `json:"shop_nature_id"`
	ShopTypeId    int `json:"shop_type_id"`
	ShopPackageId int `json:"shop_package_id"`
	CurrencyId    int `json:"currency_id"`
	// fields
	ShopCode     string `json:"shop_code" `
	ShopName     string `json:"shop_name"`
	ShopLicense  string `json:"shop_license"`
	ShopNature   string `json:"shop_nature"`
	ShopTax      string `json:"shop_tax"`
	ShopTel      string `json:"shop_tel"`
	ShopPhone    string `json:"shop_phone"`
	ShopFax      string `json:"shop_fax"`
	ShopEmail    string `json:"shop_email"`
	ShopAddress  string `json:"shop_address"`
	ShopDistrict string `json:"shop_district"`
	ShopAmphoe   string `json:"shop_amphoe"`
	ShopProvince string `json:"shop_province"`
	ShopZipcode  string `json:"shop_zipcode"`

	ShopNameEn     string `json:"shop_name_en"`
	ShopAddressEn  string `json:"shop_address_en"`
	ShopDistrictEn string `json:"shop_district_en"`
	ShopAmphoeEn   string `json:"shop_amphoe_en"`
	ShopProvinceEn string `json:"shop_province_en"`

	ShopLatlong           string `json:"shop_latlong"`
	ShopPromptpayIdd      string `json:"shop_promptpay_id"`
	ShopPromptpayName     string `json:"shop_promptpay_name"`
	ShopMotherId          int    `json:"shop_mother_id"`
	ShopImage             string `json:"shop_image"`
	ShopDetail            string `json:"shop_detail"`
	ShopLang              string `json:"shop_lang"`
	ShopStatusId          int    `json:"shop_status_id"`
	ShopSmsSum            string `json:"shop_sms_sum"`
	ShopSmsAll            string `json:"shop_sms_all"`
	ShopPointExchange     string `json:"shop_point_exchange"`
	ShopDrugLowest        string `json:"shop_drug_lowest"`
	ShopDrugExpire        string `json:"shop_drug_expire"`
	ShopPackageType       string `json:"shop_package_type"`
	ShopPackageUsemanager string `json:"shop_package_usemanager"`
	ShopPackageUseuser    string `json:"shop_package_useuser"`
	ShopPackageUsedoctor  string `json:"shop_package_usedoctor"`
	ShopPackageUseshop    string `json:"shop_package_useshop"`
	ShopPackage           string `json:"shop_package"`
	ShopExpire            string `json:"shop_expire"`
	ShopFacebook          string `json:"shop_facebook"`
	ShopLine              string `json:"shop_line"`
	ShopInstagram         string `json:"shop_instagram"`
	ShopPortIdcard        string `json:"shop_port_idcard"`
	ShopPrintType         string `json:"shop_print_type"`
	ShopComment           string `json:"shop_comment"`
	ShopConnectLineStatus string `json:"shop_connect_line_status"`
	ShopPaysolutionsRefno string `json:"shop_paysolutions_refno"`
	ShopCmline            string `json:"shop_cmline"`
	ShopS3DeleteId        int    `json:"shop_s3_delete_id"`
	ShopShippingCost      string `json:"shop_shipping_cost"`
	PaysolutionsMerchant  string `json:"paysolutions_merchant"`
	PaysolutionsSecretkey string `json:"paysolutions_secretkey"`
	PaysolutionsApikey    string `json:"paysolutions_apikey"`
	PaysolutionsActive    string `json:"paysolutions_active"`
	ShopPointGiveRate     int    `json:"shop_point_give_rate"`
	ShopPointUseRate      int    `json:"shop_point_use_rate"`
	ShopPointCourse       int    `json:"shop_point_course"`
	ShopPointChecking     int    `json:"shop_point_checking"`
	ShopPointProduct      int    `json:"shop_point_product"`
	ShopPointCoin         int    `json:"shop_point_coin"`
	Invoice_copy          int    `json:"invoice_copy"`
	Receipt_copy          int    `json:"receipt_copy"`
	Tax_copy              int    `json:"tax_copy"`
	Purchase_copy         int    `json:"purchase_copy"`
	Transfer_copy         int    `json:"transfer_copy"`
	ShopCreate            string `json:"shop_create"`
	ShopUpdate            string `json:"shop_update"`
	UserFullname          string `json:"user_fullname"`
	UserFullnameEn        string `json:"user_fullname_en"`
}

type ShopReadPayload struct {
	SearchText *string `json:"search_text" binding:"required,omitempty"`
}

type ObjCreateUserShop struct {
	ID          int
	ShopRoleId  int
	ShopId      int
	UserId      int
	MusIsActive int
	MusIsOwner  int
}

type ObjCreateDocSetting struct {
	ID     int
	ShopId int
}

type ObjCreateShopRole struct {
	ID          int
	ShopId      int
	RoleNameTh  string
	RoleNameEn  string
	RoleNameLo  string
	RoleNameMm  string
	RoleSort    int
	RoleIsOwner int
	RoleDefault int
	RoleCreate  string
	RoleUpdate  string
}

type ObjCreateShopRoleMenu struct {
	ID         int
	ShopRoleId int
	MenuId     int
}

type ObjCheckCreateShop struct {
	ShopId               int `json:"shop_id"`
	PackageId            int `json:"package_id"`
	PackageOrderStatusId int `json:"package_order_status_id"`
}

type ObjInitCreateShop struct {
	ShopId int
}

type ObjSwitchShop struct {
	ShopId   int    `json:"shop_id" binding:"required"`
	ShopName string `json:"shop_name" binding:"required"`
}

type AddShopAuth struct {
	Id         int    `json:"ID"`
	UserId     int    `json:"user_id"`
	ShopId     int    `json:"shop_id"`
	SaToken    string `json:"sa_token"`
	IsExpired  int    `json:"is_expired"`
	CreateDate string `json:"create_date"`
}

type ShopCalendar struct {
	Id            int    `json:"shop_id"`
	Shop_name     string `json:"shop_name"`
	Shop_gc_token string `json:"shop_gc_token"`
	Shop_gid      string `json:"shop_gid"`
}

type ObjCheckShopAuth struct {
	ShopId       int    `json:"shop_id"`
	UserId       int    `json:"user_id"`
	ShopStatusId int    `json:"shop_status_id"`
	ShopExpire   string `json:"shop_expire"`
}

type PayloadShopListOnline struct {
	Co_citizen_id int `json:"co_citizen_id" binding:"required"`
	Co_tel        int `json:"co_tel" binding:"required"`
	Co_email      int `json:"co_email" binding:"required"`
}
