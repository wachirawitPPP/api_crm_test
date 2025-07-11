package routes

import (
	"linecrmapi/controllers"
	"linecrmapi/middlewares"

	"github.com/gin-gonic/gin"
)

func SetRouterItem(router *gin.Engine) {
	r := router.Group("/item")
	{
		r.POST("/product", middlewares.CheckPublicKey, controllers.ItemProductList)
		r.POST("/course", middlewares.CheckPublicKey, controllers.ItemCourseList)
		r.POST("/checking", middlewares.CheckPublicKey, controllers.ItemCheckingList)
	}
	return
}

/**
 * @apiVersion 2.0.0
 * @apiGroup Online Item
 * @api {POST} item/product 1.ItemProductList
 * @apiDescription ItemProductList
 * @apiHeader   {String}    Authorization   Bearer Token "Access Tokens"
 * @apiBody    {String}    Search        	คำค้นหา
 * @apiBody    {Number}    ActivePage    	active_page
 * @apiBody    {Number}    PerPage       	per_page
 * @apiSuccess {Boolean}    status        	สถานะ : true,false
 * @apiSuccess {String}     message       	ข้อความตอบกลับ
 * @apiSuccess {Object}     data          	{}
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
    "data": [
        {
            "id": 182,
            "shop_id": 619,
            "shop_name": "Teendoi Studio",
            "shop_name_en": "Teendoi Studio",
            "shop_code": "S00619",
            "product_id": 182,
            "pd_code": "P0001",
            "pd_barcode": "P0001",
            "pd_name": "พาราเซตามอล online",
            "pd_type_id": 1,
            "pd_image_1": "https://apsth-x.s3.ap-southeast-1.amazonaws.com/product/a13a80c2-28de-4cce-a893-2e361b2be368.png",
            "pd_image_2": "https://apsth-x.s3.ap-southeast-1.amazonaws.com/product/58edaf73-fb95-46ed-9196-4f56826ac790.png",
            "pd_image_3": "https://apsth-x.s3.ap-southeast-1.amazonaws.com/product/a9b88dd7-bcd5-4f2f-8d04-1076b7f15004.jpeg",
            "pd_image_4": "",
            "pd_detail": "",
            "subs": [
                {
                    "id": 182,
                    "shop_id": 0,
                    "product_id": 182,
                    "product_store_id": 258,
                    "product_units_id": 319,
                    "pd_code": "P0001",
                    "pd_name": "พาราเซตามอล online",
                    "u_name": "เม็ด",
                    "pu_amount": 1,
                    "pu_rate": 1,
                    "balance": 1467,
                    "psp_price_ipd": 2,
                    "psp_price_opd": 5,
                    "topical_id": 0,
                    "topical_detail": "",
                    "drug_direction": "",
										"pd_image_1": "https://apsth-x.s3.ap-southeast-1.amazonaws.com/product/a13a80c2-28de-4cce-a893-2e361b2be368.png",
										"pd_image_2": "https://apsth-x.s3.ap-southeast-1.amazonaws.com/product/58edaf73-fb95-46ed-9196-4f56826ac790.png",
										"pd_image_3": "https://apsth-x.s3.ap-southeast-1.amazonaws.com/product/a9b88dd7-bcd5-4f2f-8d04-1076b7f15004.jpeg",
										"pd_image_4": "",
										"pd_detail": "",
                    "label": "",
                    "is_set": 0,
                    "id_set": null,
                    "units": [
                        {
                            "id": 182,
                            "product_id": 182,
                            "product_units_id": 319,
                            "pd_code": "P0001",
                            "pd_name": "พาราเซตามอล online",
                            "pu_amount": 1,
                            "pu_rate": 1,
                            "u_name": "เม็ด"
                        },
                        {
                            "id": 182,
                            "product_id": 182,
                            "product_units_id": 320,
                            "pd_code": "P0001",
                            "pd_name": "พาราเซตามอล online",
                            "pu_amount": 1,
                            "pu_rate": 12,
                            "u_name": "กล่อง"
                        }
                    ]
                }
            ]
        }
    ],
    "message": "",
    "status": true
}
*/

/**
 * @apiVersion 2.0.0
 * @apiGroup Online Item
 * @api {POST} item/course 2.ItemCourseList
 * @apiDescription ItemCourseList
 * @apiHeader   {String}    Authorization   Bearer Token "Access Tokens"
 * @apiBody    {String}    Search        	คำค้นหา
 * @apiBody    {Number}    ActivePage    	active_page
 * @apiBody    {Number}    PerPage       	per_page
 * @apiSuccess {Boolean}    status        	สถานะ : true,false
 * @apiSuccess {String}     message       	ข้อความตอบกลับ
 * @apiSuccess {Object}     data          	{}
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
    "data": [
        {
            "id": 5478,
            "shop_id": 619,
            "shop_name": "Teendoi Studio",
            "shop_name_en": "Teendoi Studio",
            "shop_code": "S00619",
            "course_type_id": 1,
            "course_code": "CO0054",
            "course_name": "Course_is_line",
            "course_amount": 1,
            "course_unit": "line",
            "course_qtyset": 1,
            "course_cost": 0,
            "course_opd": 0,
            "course_ipd": 0,
            "course_lock_drug": 1,
            "course_image_1": "https://apsth-x.s3.ap-southeast-1.amazonaws.com/course/7a2fff0f-8b8b-4047-a8aa-648661ad60f3.png",
            "course_image_2": "https://apsth-x.s3.ap-southeast-1.amazonaws.com/course/2c91b97d-f25b-4287-9353-b2fda7cea7f3.jpeg",
            "course_image_3": "https://apsth-x.s3.ap-southeast-1.amazonaws.com/course/5e31da81-a653-4087-bf8b-2f88952d6783.png",
            "course_image_4": "https://apsth-x.s3.ap-southeast-1.amazonaws.com/course/ed10e372-f8a4-40f9-b684-4720f06b29a8.jpeg",
						"course_detail": "",
            "subs": [
                {
                    "id": 5478,
                    "course_type_id": 1,
                    "course_code": "CO0054",
                    "course_name": "Course_is_line",
                    "course_amount": 1,
                    "course_unit": "line",
                    "course_qtyset": 0,
                    "course_cost": 0,
                    "course_opd": 0,
                    "course_ipd": 0,
                    "course_list_qtyset": 0,
                    "course_list_opd": 0,
                    "course_list_ipd": 0,
                    "course_lock_drug": 0,
										"course_image_1": "https://apsth-x.s3.ap-southeast-1.amazonaws.com/course/7a2fff0f-8b8b-4047-a8aa-648661ad60f3.png",
										"course_image_2": "https://apsth-x.s3.ap-southeast-1.amazonaws.com/course/2c91b97d-f25b-4287-9353-b2fda7cea7f3.jpeg",
										"course_image_3": "https://apsth-x.s3.ap-southeast-1.amazonaws.com/course/5e31da81-a653-4087-bf8b-2f88952d6783.png",
										"course_image_4": "https://apsth-x.s3.ap-southeast-1.amazonaws.com/course/ed10e372-f8a4-40f9-b684-4720f06b29a8.jpeg",
										"course_detail": "",
                    "label": "",
                    "is_set": 0,
                    "id_set": null,
                    "products": [
                        {
                            "id": 182,
                            "product_id": 182,
                            "product_store_id": 258,
                            "product_units_id": 319,
                            "pd_code": "P0001",
                            "pd_name": "พาราเซตามอล online",
                            "u_name": "เม็ด",
                            "pu_amount": 10,
                            "Pu_rate": 1,
                            "balance": 1467,
                            "label": "CO0054:Course_is_line",
                            "is_set": 0,
                            "id_set": null,
                            "qty_set": 0
                        },
                        {
                            "id": 183,
                            "product_id": 183,
                            "product_store_id": 259,
                            "product_units_id": 321,
                            "pd_code": "P0002",
                            "pd_name": "แก้แพ้ online",
                            "u_name": "เม็ด",
                            "pu_amount": 1,
                            "Pu_rate": 1,
                            "balance": 1614.45,
                            "label": "CO0054:Course_is_line",
                            "is_set": 0,
                            "id_set": null,
                            "qty_set": 0
                        }
                    ]
                }
            ]
        }
    ],
    "message": "",
    "status": true
}
*/

/**
 * @apiVersion 2.0.0
 * @apiGroup Online Item
 * @api {POST} item/checking 3.ItemCheckingList
 * @apiDescription ItemCheckingList
 * @apiHeader   {String}    Authorization   Bearer Token "Access Tokens"
 * @apiBody    {String}    Search        	คำค้นหา
 * @apiBody    {Number}    ActivePage    	active_page
 * @apiBody    {Number}    PerPage       	per_page
 * @apiSuccess {Boolean}    status        	สถานะ : true,false
 * @apiSuccess {String}     message       	ข้อความตอบกลับ
 * @apiSuccess {Object}     data          	{}
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
    "data": [
        {
            "id": 93141,
            "shop_id": 619,
            "shop_name": "Teendoi Studio",
            "shop_name_en": "Teendoi Studio",
            "shop_code": "S00619",
            "checking_type_id": 1,
            "checking_code": "C0020",
            "checking_amount": 1,
            "checking_name": "TestLine",
            "checking_unit": "Line",
            "checking_cost": 0,
            "checking_opd": 0,
            "checking_ipd": 0,
            "checking_image_1": "https://apsth-x.s3.ap-southeast-1.amazonaws.com/check/3c8256d5-5039-4ed4-969f-f570c0b23c1c.jpeg",
            "checking_image_2": "https://apsth-x.s3.ap-southeast-1.amazonaws.com/check/4fbe0f59-9771-4d0b-b319-c98408fa524b.jpeg",
            "checking_image_3": "https://apsth-x.s3.ap-southeast-1.amazonaws.com/check/093c478d-f1c7-4fa0-85bf-1d1311795deb.jpeg",
            "checking_image_4": "https://apsth-x.s3.ap-southeast-1.amazonaws.com/check/28fb8f62-e5f7-4f7e-9ec8-5a5fd2f29fae.png",
						"checking_detail": "",
            "subs": [
                {
                    "id": 93141,
                    "checking_type_id": 1,
                    "checking_code": "C0020",
                    "checking_amount": 1,
                    "checking_name": "TestLine",
                    "checking_unit": "Line",
                    "checking_opd": 0,
                    "checking_cost": 0,
                    "checking_ipd": 0,
                    "checking_list_opd": 0,
                    "checking_list_ipd": 0,
										"checking_image_1": "https://apsth-x.s3.ap-southeast-1.amazonaws.com/course/7a2fff0f-8b8b-4047-a8aa-648661ad60f3.png",
										"checking_image_2": "https://apsth-x.s3.ap-southeast-1.amazonaws.com/course/2c91b97d-f25b-4287-9353-b2fda7cea7f3.jpeg",
										"checking_image_3": "https://apsth-x.s3.ap-southeast-1.amazonaws.com/course/5e31da81-a653-4087-bf8b-2f88952d6783.png",
										"checking_image_4": "https://apsth-x.s3.ap-southeast-1.amazonaws.com/course/ed10e372-f8a4-40f9-b684-4720f06b29a8.jpeg",
										"checking_detail": "",
                    "label": "",
                    "is_set": 0,
                    "id_set": null,
                    "products": null
                }
            ]
        }
    ],
    "message": "",
    "status": true
}
*/
