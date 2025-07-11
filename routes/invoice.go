package routes

import (
	"linecrmapi/controllers"
	"linecrmapi/middlewares"

	"github.com/gin-gonic/gin"
)

func SetRouterInvoices(router *gin.Engine) {
	r := router.Group("/invoice")
	{
		r.POST("/search", middlewares.CheckAccessToken, controllers.InvoicesSearch)
		r.GET("/:id", middlewares.CheckAccessToken, controllers.InvoicesDetail)
		r.POST("", middlewares.CheckAccessToken, controllers.AddInvoice)
		r.DELETE("/:id", middlewares.CheckAccessToken, controllers.DelInvoice)
	}
	return
}

/**
* @apiGroup Online Order
* @api {post} invoice 2.Add Invoice
* @apiDescription เพิ่มใบแจ้งหนี้
* @apiVersion 2.0.0
* @apiHeader   {String}    Authorization   Bearer Token "Access Tokens"
* @apiBody {Number}  order_id               รหัส order
* @apiSuccess {Boolean}    response สถานะ : true,false
* @apiSuccess {String} message ข้อความตอบกลับ
* @apiSuccess {String} data รหัสใบแจ้งหนี้
* @apiSuccess {String} code เลขที่ใบแจ้งหนี้
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*   {
   "code": "INV0353", รหัสใบแจ้งหนี้
   "data": 112059, เลขที่ใบแจ้งหนี้
   "message": "Add Invoice Success.",
   "status": true
}
*/
