package routes

import (
	"linecrmapi/controllers"
	"linecrmapi/middlewares"

	"github.com/gin-gonic/gin"
)

func SetRouterAppointment(router *gin.Engine) {
	r := router.Group("/appointment")
	{
		r.GET("/topic", middlewares.CheckAccessToken, controllers.GetAppointmentTopic)
		// r.GET("/tag", middlewares.CheckAccessToken, controllers.GetAppointmentTag)
		r.GET("/time", middlewares.CheckAccessToken, controllers.ShopTimeOpen)
		r.GET("/doctor", middlewares.CheckAccessToken, controllers.GetAppointmentDoctor)

		r.GET("/list", middlewares.CheckAccessToken, controllers.GetAppointmentList)
		r.POST("/calendar", middlewares.CheckAccessToken, controllers.AppointmentCalendar)
		r.POST("/add", middlewares.CheckAccessToken, controllers.AddAppointment)
		r.DELETE("/:apId", middlewares.CheckAccessToken, controllers.DeleteAppointment)
	}
}

/**
 * @apiVersion 1.0.0
 * @apiGroup Appointment
 * @api {GET} appointment/topic 1.Get Appointment Topic
 * @apiDescription Get Appointment Topic
 * @apiHeader   {String}    Authorization   Bearer Token "Access Tokens"
 * @apiSuccess {Boolean}    status        	สถานะ : true,false
 * @apiSuccess {String}     message       	ข้อความตอบกลับ
 * @apiSuccess {Object}     data          	{}
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 * {
 *       "data": [
			{
            "id": 37,
            "topic_th": "กายภาพบำบัด",
            "topic_en": "Physical Therapy"
        	}
 *       ],
 *       "status": true,
 *       "message": "Success"
 *     }
*/

/**
 * @apiVersion 1.0.0
 * @apiGroup Appointment
 * @api {GET} appointment/time 2.Get Appointment Time
 * @apiDescription Get Appointment Time
 * @apiHeader   {String}    Authorization   Bearer Token "Access Tokens"
 * @apiSuccess {Boolean}    status        	สถานะ : true,false
 * @apiSuccess {String}     message       	ข้อความตอบกลับ
 * @apiSuccess {Object}     data          	{}
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 * {
 *       "data": [
			{
            "TimeData": "08:00:00",
            "day_datas": null
			},
			{
				"TimeData": "08:30:00",
				"day_datas": null
			}
 *       ],
 *       "status": true,
 *       "message": "Success"
 *     }
*/

/**
 * @apiVersion 1.0.0
 * @apiGroup Appointment
 * @api {GET} appointment/doctor 3.Get Appointment Doctor
 * @apiDescription Get Appointment Doctor
 * @apiHeader   {String}    Authorization   Bearer Token "Access Tokens"
 * @apiSuccess {Boolean}    status        	สถานะ : true,false
 * @apiSuccess {String}     message       	ข้อความตอบกลับ
 * @apiSuccess {Object}     data          	{}
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 * {
 *       "data": [
			{
            "id": 11,
            "user_email": "teendoistudio@gmail.com",
            "user_fullname": "ตีนดอยแน๊ะแจ๊ะ",
            "user_fullname_en": "Teendoi Studio",
            "role_name_th": "ผู้ดูแลระบบ",
            "role_name_en": "Admin"
        	}
 *       ],
 *       "status": true,
 *       "message": "Success"
 *     }
*/

/**
 * @apiVersion 1.0.0
 * @apiGroup Appointment
 * @api {GET} appointment/list 4.Get Appointment List
 * @apiDescription Get Appointment List
 * @apiHeader   {String}    Authorization   Bearer Token "Access Tokens"
 * @apiSuccess {Boolean}    status        	สถานะ : true,false
 * @apiSuccess {String}     message       	ข้อความตอบกลับ
 * @apiSuccess {Object}     data          	{}
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 * {
 *       "data": [
			{
            "id": 157,
            "shop_id": 619,
            "user_id": 11,
            "user_fullname": "ตีนดอยแน๊ะแจ๊ะ",
            "user_fullname_en": "Teendoi Studio",
            "role_name_th": "",
            "role_name_en": "",
            "customer_id": 37,
            "ctm_id": "HN00037",
            "customer_fullname": "ทดสอบ 1",
            "ctm_fname_en": "test",
            "ctm_lname_en": "1",
            "ctm_birthdate": "2023-08-21T00:00:00+07:00",
            "ctm_tel": "0835652123",
            "ctm_tel_2": "00",
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
            "ap_opd_type": 0,
            "ap_is_tele": 1,
            "ap_tele_code": "",
            "ap_tele_url": "",
            "ap_create": "2024-10-31T10:35:57+07:00",
            "ap_update": "2024-10-31T10:35:57+07:00"
        	}
 *       ],
 *       "status": true,
 *       "message": "Success"
 *     }
*/

/**
 * @apiVersion 1.0.0
 * @apiGroup Appointment
 * @api {POST} appointment/add 5.Add Appointment
 * @apiDescription Add Appointment
 * @apiHeader   {String}    Authorization   Bearer Token "Access Tokens"
 * @apiBody {Int}   user_id         รหัสหมอ user_id
 * @apiBody {String}   ap_topic     หัวข้อนัดหมาย
 * @apiBody {String}   ap_datetime  วันที่ + เวลา "2024-11-04 14:30:00"
 * @apiSuccess {Boolean}    status        	สถานะ : true,false
 * @apiSuccess {String}     message       	ข้อความตอบกลับ
 * @apiSuccess {Object}     data          	{}
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 * {
        "data": "",
        "message": "Add Appointment Success.",
        "status": true
*  }
*/

/**
 * @apiVersion 1.0.0
 * @apiGroup Appointment
 * @api {DELETE} appointment/:apId 6.Delete Appointment
 * @apiDescription Delete Appointment
 * @apiHeader   {String}    Authorization   Bearer Token "Access Tokens"
 * @apiParam {Number}       apId            รหัสนัดหมาย
 * @apiSuccess {Boolean}    status        	สถานะ : true,false
 * @apiSuccess {String}     message       	ข้อความตอบกลับ
 * @apiSuccess {Object}     data          	{}
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 * {
        "data": "",
        "message": "Delete Appointment Success.",
        "status": true
*  }
*/
