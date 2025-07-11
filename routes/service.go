package routes

import (
	"linecrmapi/controllers"
	"linecrmapi/middlewares"

	"github.com/gin-gonic/gin"
)

func SetRouterService(router *gin.Engine) {
	r := router.Group("/service")
	{
		r.GET("id/:serviceId", middlewares.CheckAccessToken, controllers.ItemService)       //use
		r.POST("search", middlewares.CheckAccessToken, controllers.ServiceSearch)           //use
		r.GET("used/:serviceId", middlewares.CheckAccessToken, controllers.ItemServiceUsed) //use
	}
	return
}

/**
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @api {POST} service/search 1.ServiceSearch
 * @apiDescription ServiceSearch
 * @apiHeader   {String}    Authorization   Bearer Token "Access Tokens"
 * @apiBody {Number}      Search            search
 * @apiBody {Number}      ActivePage        active_page
 * @apiBody {Number}      PerPage           per_page
 * @apiSuccess {Boolean}    status        	สถานะ : true,false
 * @apiSuccess {String}     message       	ข้อความตอบกลับ
 * @apiSuccess {Object}     data          	{}
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 {
    "data": {
        "result_data": [
            {
                "id": 419,
                "receipt_id": 645,
                "queue_id": 0,
                "rec_code": "RE0217",
                "receipt_detail_id": 1712,
                "shop_id": 619,
                "customer_shop_id": 619,
                "user_id": 4,
                "ser_customer_id": 37,
                "customer_id": 37,
                "course_id": 118,
                "ser_tranfer_id": 0,
                "ser_code": "CO0029",
                "ser_name": "Lock",
                "ser_lock_drug": 1,
                "ser_qty": 2,
                "ser_unit": "Lock",
                "ser_use_date": 1,
                "ser_exp": 30,
                "ser_exp_date": "2024-02-14T00:00:00+07:00",
                "ser_use": 2,
                "ser_tranfer": 0,
                "ser_is_active": 1,
                "ser_datetime": "2024-01-15T10:00:43+07:00",
                "ser_create": "2024-01-15T10:00:43+07:00",
                "ser_update": "2024-01-17T09:05:46+07:00",
                "course_amount": 1,
                "course_ipd": 10,
                "course_opd": 20,
                "course_cost": 10,
                "ser_amount": 1
            }
        ],
        "count_of_page": 2,
        "count_all": 2
    },
    "message": "",
    "status": true
}
*/

/**
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @api {GET} service/id/:serviceId 2.ItemService
 * @apiParam {Number} serviceId
 * @apiDescription ItemService
 * @apiHeader   {String}    Authorization   Bearer Token "Access Tokens"
 * @apiSuccess {Boolean}    status        	สถานะ : true,false
 * @apiSuccess {String}     message       	ข้อความตอบกลับ
 * @apiSuccess {Object}     data          	{}
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
{
    "data": {
        "id": 419,
        "course_id": 118,
        "ser_code": "CO0029",
        "ser_name": "Lock",
        "ser_lock_drug": 1,
        "ser_qty": 2,
        "ser_unit": "Lock",
        "ser_use_date": 1,
        "ser_exp": 30,
        "ser_exp_day": -260,
        "ser_exp_date": "2024-02-14T00:00:00+07:00",
        "ser_use": 2,
        "ser_tranfer": 0,
        "ser_amount": 1,
        "ser_price_total": 2000,
        "ser_is_active": 0,
        "tranfer_is_qty_course": 0,
        "tranfer_is_qty_product": 1,
        "adjust_is_qty_course": 0,
        "adjust_is_day_course": 1,
        "adjust_is_qty_product": 1,
        "products": [
            {
                "id": 326,
                "shop_id": 619,
                "service_id": 419,
                "course_id": 118,
                "receipt_id": 645,
                "receipt_detail_id": 1712,
                "product_id": 292,
                "product_store_id": 261,
                "product_unit_id": 401,
                "serp_code": "PD0029",
                "serp_name": "ยาแก้ไอ",
                "serp_qty": 200,
                "serp_use": 2,
                "serp_tranfer": 10,
                "serp_adjust": 0,
                "serp_balance": 188,
                "serp_unit": "เม็ด",
                "serp_lock_drug": 1,
                "serp_use_set_qty": 1,
                "serp_is_active": 1,
                "serp_datetime": "2024-01-15T10:00:43+07:00",
                "serp_create": "2024-01-15T10:00:43+07:00",
                "serp_modify": "2024-01-17T09:05:46+07:00"
            },
            {
                "id": 327,
                "shop_id": 619,
                "service_id": 419,
                "course_id": 118,
                "receipt_id": 645,
                "receipt_detail_id": 1712,
                "product_id": 300,
                "product_store_id": 254,
                "product_unit_id": 413,
                "serp_code": "PD0035",
                "serp_name": "เข็มฉีดยา",
                "serp_qty": 200,
                "serp_use": 2,
                "serp_tranfer": 10,
                "serp_adjust": 0,
                "serp_balance": 188,
                "serp_unit": "เข็ม",
                "serp_lock_drug": 1,
                "serp_use_set_qty": 1,
                "serp_is_active": 1,
                "serp_datetime": "2024-01-15T10:00:43+07:00",
                "serp_create": "2024-01-15T10:00:43+07:00",
                "serp_modify": "2024-01-17T09:05:46+07:00"
            }
        ]
    },
    "message": "",
    "status": true
}
*/

/**
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @api {GET} service/used/:serviceId 3.ItemServiceUsed
 * @apiParam {Number} serviceId
 * @apiDescription ItemServiceUsed
 * @apiHeader   {String}    Authorization   Bearer Token "Access Tokens"
 * @apiSuccess {Boolean}    status        	สถานะ : true,false
 * @apiSuccess {String}     message       	ข้อความตอบกลับ
 * @apiSuccess {Object}     data          	{}
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
    "data": [
        {
            "id": 387,
            "shop_id": 619,
            "shop_name": "Teendoi Studio",
            "service_id": 419,
            "queue_id": 550,
            "que_code": "SEV2024010051",
            "receipt_id": 645,
            "course_id": 118,
            "customer_id": 37,
            "user_id": 4,
            "user_fullname": "Nut",
            "seru_code": "CO0029",
            "seru_name": "Lock",
            "seru_qty": 1,
            "seru_unit": "Lock",
            "seru_cost": 10,
            "seru_date": "2024-01-17T00:00:00+07:00",
            "seru_is_active": 1,
            "seru_datetime": "2024-01-17T09:05:46+07:00",
            "seru_create": "2024-01-17T09:05:46+07:00",
            "seru_update": "2024-01-17T09:05:46+07:00"
        }
    ],
    "message": "",
    "status": true
}
*/
