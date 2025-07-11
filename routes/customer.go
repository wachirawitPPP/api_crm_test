package routes

import (
	"linecrmapi/controllers"
	"linecrmapi/middlewares"

	"github.com/gin-gonic/gin"
)

func SetRouterCustomer(router *gin.Engine) {
	r := router.Group("/customer")
	{
		r.GET("", middlewares.CheckAccessToken, controllers.GetCustomerById)                                  //use
		r.POST("/opd/pagination", middlewares.CheckAccessToken, controllers.GetCustomerOpdPagination)         //use
		r.POST("/receipt/pagination", middlewares.CheckAccessToken, controllers.GetCustomerReceiptPagination) //use
		r.POST("/appointment", middlewares.CheckAccessToken, controllers.SearchAppointmentByCustomer)         //use
		r.POST("/history/lab", middlewares.CheckAccessToken, controllers.HistoryLab)                          //use
		r.POST("/history/xray", middlewares.CheckAccessToken, controllers.HistoryXray)                        //use
		r.POST("/history/document", middlewares.CheckAccessToken, controllers.HistoryDocument)                //use
        r.GET("/co_profile", middlewares.CheckAccessToken, controllers.GetCustomerOnlineById)                                    //use
	}
	return
}

/**
 * @apiVersion 1.0.0
 * @apiGroup Customer
 * @api {GET} customer 1.Get CustomerById
 * @apiDescription Get Customer By Id
 * @apiHeader   {String}    Authorization   Bearer Token "Access Tokens"
 * @apiSuccess {Boolean}    status        	สถานะ : true,false
 * @apiSuccess {String}     message       	ข้อความตอบกลับ
 * @apiSuccess {Object}     data          	{}
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 * {
 *       "data": {
			"customer": {
				"id": 42901,
				"shop_id": 619,
				"customer_group_id": 1,
				"user_id": 0,
				"ctm_id": "AS001",
				"ctm_citizen_id": "AS001",
				"ctm_passport_id": "",
				"ctm_prefix": "นางสาว",
				"ctm_fname": "Chiew xx",
				"ctm_lname": "szelynn",
				"ctm_nname": "",
				"ctm_fname_en": "",
				"ctm_lname_en": "",
				"ctm_gender": "หญิง",
				"ctm_nation": "ไม่ระบุ",
				"ctm_religion": "ไม่ระบุ",
				"ctm_edu_level": "ไม่ระบุ",
				"ctm_marital_status": "ไม่ระบุ",
				"ctm_blood": "ไม่ระบุ",
				"ctm_email": "",
				"ctm_tel": "",
				"ctm_tel_2": "",
				"ctm_birthdate": "1991-01-01T00:00:00+07:00",
				"ctm_address": "ไม่ระบุ",
				"ctm_district": "เกาะลันตาใหญ่",
				"ctm_amphoe": "เกาะลันตา",
				"ctm_province": "กระบี่",
				"ctm_zipcode": "81150",
				"ctm_comment": "",
				"ctm_weight": 0,
				"ctm_height": 0,
				"ctm_waistline": 0,
				"ctm_chest": 0,
				"ctm_treatment_type": 2,
				"right_treatment_id": 23,
				"ctm_allergic": "",
				"ctm_mental_health": "",
				"ctm_disease": "",
				"ctm_health_comment": "",
				"ctm_image": "https://aps-x.s3.ap-southeast-1.amazonaws.com/customer/image.jpeg",
				"ctm_image_size": 99,
				"ctm_point": 18,
				"ctm_coin": 0,
				"line_token": "",
				"line_send": 0,
				"line_send_date": "0001-01-01T00:00:00Z",
				"facebook_id": "",
				"company_name": "",
				"company_tax": "",
				"company_tel": "",
				"company_email": "",
				"company_address": "",
				"company_district": "",
				"company_amphoe": "",
				"company_province": "",
				"company_zipcode": "",
				"ctm_subscribe_opd": 0,
				"ctm_subscribe_lab": 0,
				"ctm_subscribe_cert": 0,
				"ctm_subscribe_receipt": 0,
				"ctm_subscribe_appoint": 0,
				"ctm_is_active": 1,
				"ctm_is_del": 0,
				"ctm_create": "2024-05-31T14:49:33+07:00",
				"ctm_update": "2024-09-06T19:32:12+07:00",
				"ctm_subscribe_pdpa_token": "",
				"ctm_subscribe_pdpa_image": "",
				"cg_name": "group - 619 - 1",
				"cg_save_type": 2,
				"cg_save": 1,
				"rt_code": "PIC",
				"rt_name": "เจนเนอราลี่ ประกันภัย (ไทยแลนด์)",
				"rt_name_en": "Generali Insurance (Thailand) Public Company Limited"
			},
			"tag": [],
			"family": [],
			"contact": [],
			"balance": {
				"pay_total": 0,
				"balance_total": 200
			}
 *       },
 *       "status": true,
 *       "message": "CustomerById Success"
 *     }
*/

/**
* @apiVersion 1.0.0
* @apiGroup Customer
* @api {POST} customer/opd/pagination 2.CustomerOpdPagination
* @apiDescription CustomerOpdPagination
* @apiHeader   {String}    Authorization   Bearer Token "Access Tokens"
* @apiBody {Number}        QueueTypeId             queue_type_id
* @apiBody {String}        SearchText              search_text
* @apiBody {Number}        CurrentPage             current_page
* @apiBody {Number}        PerPage                 per_page
* @apiSuccess {Boolean}    status        	สถานะ : true,false
* @apiSuccess {String}     message       	ข้อความตอบกลับ
* @apiSuccess {Object}     data          	{}
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
{
    "data": {
        "items": [
            {
                "opd_id": 193408,
                "opd_code": "OPD20240166-2",
                "opd_date": "2024-10-18T00:00:00+07:00",
                "opd_create": "2024-10-18T09:18:08+07:00",
                "que_id": 96787,
                "que_admis_id": 2,
                "que_code": "OPD20240166",
                "que_datetime": "2024-10-18T08:57:44+07:00",
                "que_create": "",
                "user_fullname": "ทดสอบ เดฟ",
                "shop_id": 619,
                "shop_name": "Teendoi Studio"
            }
        ],
        "subs": [
            {
                "id": 95677,
                "opd_id": 193408,
                "diagnostic_id": 1,
                "diagnostic_code": "0001",
                "diagnostic_th": "ทดสอบแก้วินิจฉัย",
                "diagnostic_en": "Test12",
                "diagnostic_detail": ""
            }
        ],
        "count_page": 10,
        "count_all": 39
    },
    "message": "Get data successful.",
    "status": true
}
*/

/**
 * @apiVersion 1.0.0
 * @apiGroup Customer
 * @api {POST} customer/receipt/pagination 3.GetCustomerReceiptPagination
 * @apiDescription GetCustomerReceiptPagination
 * @apiHeader   {String}    Authorization   Bearer Token "Access Tokens"
 * @apiBody {String}        Search                  search
 * @apiBody {String}        Rec_is_active           rec_is_active
 * @apiBody {Number}        ActivePage              active_page
 * @apiBody {Number}        PerPage                 per_page
 * @apiSuccess {Boolean}    status        	สถานะ : true,false
 * @apiSuccess {String}     message       	ข้อความตอบกลับ
 * @apiSuccess {Object}     data          	{}
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
{
    "data": {
        "result_data": [
            {
                "id": 157,
                "shop_id": 619,
                "user_id": 11,
                "user_image": "",
                "user_fullname": "ทดสอบ",
                "user_fullname_en": "ทดสอบ",
                "role_name_th": "",
                "role_name_en": "",
                "customer_id": 37,
                "customer_fullname": "ทดสอบ 1",
                "ctm_fname": "ทดสอบ",
                "ctm_lname": "1",
                "ctm_fname_en": "test",
                "ctm_lname_en": "1",
                "ap_type": 2,
                "ap_topic": "ทดสอบนัดหมาย postman",
                "ap_tel": "081-111-2222",
                "ap_datetime": "2024-10-31T14:30:00+07:00",
                "ap_note": "-",
                "ap_comment": "-",
                "ap_color": "#fffaeb",
                "ap_confirm": 0,
                "ap_status_id": 1,
                "ap_status_sms": 0,
                "ap_status_line": 0,
                "ap_sms": "",
                "ap_is_gcalendar": 0,
                "ap_gid": "",
                "ap_user_id": 11,
                "ap_is_del": 0,
                "ap_create": "2024-10-31T10:35:57+07:00",
                "ap_update": "2024-10-31T10:35:57+07:00",
                "shop_name": "Teendoi Studio"
            }
        ],
        "count_of_page": 10,
        "count_all": 37
    },
    "message": "",
    "status": true
}
*/

/**
 * @apiVersion 1.0.0
 * @apiGroup Customer
 * @api {POST} customer/history/lab 5.HistoryLab
 * @apiDescription HistoryLab
 * @apiHeader   {String}    Authorization   Bearer Token "Access Tokens"
 * @apiBody {String}       search_text		ค้นหา
 * @apiBody {String}       search_date		วันที่
 * @apiBody {Number}       current_page		หน้า
 * @apiBody {Number}       per_page			จำนวนรายการต่อหน้า
 * @apiSuccess {Boolean}    status        	สถานะ : true,false
 * @apiSuccess {String}     message       	ข้อความตอบกลับ
 * @apiSuccess {Object}     data          	{}
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
{
    "data": {
        "items": [
            {
                "id": 163620,
                "shop_id": 619,
                "receipt_id": 127364,
                "receipt_detail_id": 216047,
                "user_id": 4,
                "customer_id": 37,
                "queue_id": 96732,
                "checking_id": 92917,
                "chk_type_id": 2,
                "chk_code": "1071",
                "chk_name": "CBC",
                "chk_unit": "labplus",
                "chk_value": "1",
                "chk_upload": "",
                "chk_upload_size": 0,
                "chk_old": "",
                "direction_id": 0,
                "chk_flag": "",
                "chk_date": "",
                "chk_is_print": 0,
                "chk_is_report": 0,
                "chk_is_active": 1,
                "chk_datetime": "2024-07-25T14:19:32+07:00",
                "chk_create": "2024-07-25T14:19:32+07:00",
                "chk_update": "2024-07-27T21:01:49+07:00",
                "direction_name": "",
                "que_code": "OPD20240144",
                "que_shop_id": 619,
                "shop_name": "Teendoi Studio",
                "que_datetime": "2024-07-16T20:13:44+07:00",
                "rec_code": "RE0309",
                "user_fullname": "Nut"
            }
        ],
        "count_page": 8,
        "count_all": 8
    },
    "message": "Get data successful.",
    "status": true
}
*/

/**
 * @apiVersion 1.0.0
 * @apiGroup Customer
 * @api {POST} customer/history/xray 6.HistoryXray
 * @apiDescription HistoryXray
 * @apiHeader   {String}    Authorization   Bearer Token "Access Tokens"
 * @apiBody {String}       search_text		ค้นหา
 * @apiBody {String}       search_date		วันที่
 * @apiBody {Number}       current_page		หน้า
 * @apiBody {Number}       per_page			จำนวนรายการต่อหน้า
 * @apiSuccess {Boolean}    status        	สถานะ : true,false
 * @apiSuccess {String}     message       	ข้อความตอบกลับ
 * @apiSuccess {Object}     data          	{}
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
{
    "data": {
        "items": [
            {
                "id": 163620,
                "shop_id": 619,
                "receipt_id": 127364,
                "receipt_detail_id": 216047,
                "user_id": 4,
                "customer_id": 37,
                "queue_id": 96732,
                "checking_id": 92917,
                "chk_type_id": 2,
                "chk_code": "1071",
                "chk_name": "CBC",
                "chk_unit": "labplus",
                "chk_value": "1",
                "chk_upload": "",
                "chk_upload_size": 0,
                "chk_old": "",
                "direction_id": 0,
                "chk_flag": "",
                "chk_date": "",
                "chk_is_print": 0,
                "chk_is_report": 0,
                "chk_is_active": 1,
                "chk_datetime": "2024-07-25T14:19:32+07:00",
                "chk_create": "2024-07-25T14:19:32+07:00",
                "chk_update": "2024-07-27T21:01:49+07:00",
                "direction_name": "",
                "que_code": "OPD20240144",
                "que_shop_id": 619,
                "shop_name": "Teendoi Studio",
                "que_datetime": "2024-07-16T20:13:44+07:00",
                "rec_code": "RE0309",
                "user_fullname": "Nut"
            }
        ],
        "count_page": 8,
        "count_all": 8
    },
    "message": "Get data successful.",
    "status": true
}
*/

/**
 * @apiVersion 1.0.0
 * @apiGroup Customer
 * @api {POST} customer/history/document 7.History Medicalcert
 * @apiDescription History Medicalcert
 * @apiHeader   {String}    Authorization   Bearer Token "Access Tokens"
 * @apiBody {String}       search_text		ค้นหา
 * @apiBody {String}       search_date		วันที่
 * @apiBody {Number}       current_page		หน้า
 * @apiBody {Number}       per_page			จำนวนรายการต่อหน้า
 * @apiSuccess {Boolean}    status        	สถานะ : true,false
 * @apiSuccess {String}     message       	ข้อความตอบกลับ
 * @apiSuccess {Object}     data          	{}
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 * {
    "data": {
        "items": [
            {
                "id": 405,
                "medical_cert_type_id": 2,
                "user_id": 2,
                "opd_id": 97712,
                "mdc_code": "SICK0006",
                "mdc_is_print": 1,
                "mdc_is_del": 0,
                "mdc_create": "2024-02-09T13:49:26+07:00",
                "mdc_update": "2024-02-09T13:49:26+07:00",
                "ctm_id": "HN00037",
                "ctm_prefix": "นาง",
                "ctm_fname": "ทดสอบ",
                "ctm_lname": "1",
                "ctm_nname": "",
                "ctm_fname_en": "",
                "ctm_lname_en": "",
                "user_fullname": "ทดสอบ เดฟ",
                "user_fullname_en": "test dev",
                "mdct_th": "ใบรับรองแพทย์เจ็บป่วย",
                "mdct_en": "Medical Certificate (ILLNESS)"
            }
        ],
        "count_page": 10,
        "count_all": 13
    },
    "message": "Get data successful.",
    "status": true
}
*/
