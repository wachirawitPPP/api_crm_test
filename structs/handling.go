package structs

type ObjPayloadUpdateHandling struct {
	InvoiceId *int       `json:"invoice_id"`
	QueueId   *int       `json:"queue_id"`
	Handlings []Handling `json:"handlings"`
}

type Handling struct {
	ID             int     `json:"id"`
	ShopId         int     `json:"shop_id"`
	HandFromId     int     `json:"hand_from_id"`
	InvoiceId      *int    `json:"invoice_id"`
	QueueId        *int    `json:"queue_id"`
	UserId         int     `json:"user_id"`
	RoleId         int     `json:"role_id"`
	HandTypeId     int     `json:"hand_type_id"`
	HandPrice      float64 `json:"hand_price" gorm:"type:decimal(10,2)"`
	HandIsDeposit  int     `json:"hand_is_deposit"`
	Uuid           string  `json:"uuid"`
	UserFullname   string  `json:"user_fullname"`
	UserFullnameEn string  `json:"user_fullname_en"`
}

type ObjPayloadGetHandlingUser struct {
	ShopId     int     `json:"shop_id" binding:"required"`
	SearchText *string `json:"search_text" binding:"required"`
}

type ObjResponseGetHandlingUser struct {
	UserShopId     int    `json:"user_shop_id"`
	RoleId         int    `json:"role_id"`
	UserId         int    `json:"user_id"`
	UserEmail      string `json:"user_email"`
	UserFullname   string `json:"user_fullname"`
	UserFullnameEn string `json:"user_fullname_en"`
}

type ObjPayloadCalculateHandling struct {
	InvoiceId     int `json:"invoice_id"`
	QueueId       int `json:"queue_id"`
	UserShopId    int `json:"user_shop_id"`
	HandTypeId    int `json:"hand_type_id"`
	HandIsDeposit int `json:"hand_is_deposit"`
}

type ObjQueryUserHandling struct {
	ID              int `json:"id"`
	ShopRoleId      int `json:"shop_role_id"`
	RoleId          int `json:"role_id"`
	ShopId          int `json:"shop_id"`
	UserId          int `json:"user_id"`
	Commission_1_id int `json:"commission_1_id"`
	Commission_2_id int `json:"commission_2_id"`
	Commission_3_id int `json:"commission_3_id"`
	Commission_4_id int `json:"commission_4_id"`
	Commission_5_id int `json:"commission_5_id"`
	FeeId           int `json:"fee_id"`
}

type ObjQueryQueueTimeEnd struct {
	QueueId    int `json:"queue_id"`
	QueTimeEnd int `json:"que_time_end"`
}

type ObjQueryInvoiceChecking struct {
	InvoiceId       int     `json:"invoice_id"`
	InvoiceDetailId int     `json:"invoice_detail_id"`
	CheckingId      int     `json:"checking_id"`
	InvdTotal       float64 `json:"invd_total" gorm:"type:decimal(10,2)"`
	CheckingFeeDf   float64 `json:"checking_fee_df" gorm:"type:decimal(10,2)"`
	CheckingFeeNr   float64 `json:"checking_fee_nr" gorm:"type:decimal(10,2)"`
	CheckingFeeTr   float64 `json:"checking_fee_tr" gorm:"type:decimal(10,2)"`
	CheckingFee     float64 `json:"checking_fee" gorm:"type:decimal(10,2)"`
	InvdQty         float64 `json:"invd_qty" gorm:"type:decimal(10,2)"`
}

type ObjQueryInvoiceCourse struct {
	InvoiceId       int     `json:"invoice_id"`
	InvoiceDetailId int     `json:"invoice_detail_id"`
	CourseId        int     `json:"course_id"`
	InvdTotal       float64 `json:"invd_total" gorm:"type:decimal(10,2)"`
	CourseFeeDf     float64 `json:"course_fee_df" gorm:"type:decimal(10,2)"`
	CourseFeeNr     float64 `json:"course_fee_nr" gorm:"type:decimal(10,2)"`
	CourseFeeTr     float64 `json:"course_fee_tr" gorm:"type:decimal(10,2)"`
	CourseFee       float64 `json:"course_fee" gorm:"type:decimal(10,2)"`
}

type ObjQueryQueueService struct {
	QueueId       int     `json:"queue_id"`
	QueueCourseId int     `json:"queue_course_id"`
	ServiceId     int     `json:"service_id"`
	SerPriceTotal float64 `json:"ser_price_total" gorm:"type:decimal(10,2)"`
	CourseFeeDf   float64 `json:"course_fee_df" gorm:"type:decimal(10,2)"`
	CourseFeeNr   float64 `json:"course_fee_nr" gorm:"type:decimal(10,2)"`
	CourseFeeTr   float64 `json:"course_fee_tr" gorm:"type:decimal(10,2)"`
	CourseFee     float64 `json:"course_fee" gorm:"type:decimal(10,2)"`
	QuecQty       float64 `json:"quec_qty" gorm:"type:decimal(10,2)"`
}
