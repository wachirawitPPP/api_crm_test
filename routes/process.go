package routes

import (
	"linecrmapi/controllers"
	"linecrmapi/middlewares"

	"github.com/gin-gonic/gin"
)

func SetRouterProcesss(router *gin.Engine) {
	r := router.Group("/process")
	{
		r.POST("", middlewares.CheckAccessToken, controllers.AddProcess)
		r.DELETE("/:id", middlewares.CheckAccessToken, controllers.CancelProcess)
		r.POST("product", middlewares.CheckAccessToken, controllers.ProcessProduct)
		r.GET("product/expires", middlewares.CheckPublicKey, controllers.ProcessProductExpire)
		r.GET("product/expire/history", middlewares.CheckPublicKey, controllers.ProcessProductExpireHistory)
	}
	return
}

/**
* @apiGroup Online Order
* @api {post} process 4.Add Process
* @apiDescription ประมวลผลชำระเงิน
* @apiVersion 2.0.0
* @apiHeader   {String}    Authorization   Bearer Token "Access Tokens"
* @apiBody {Number}  receipt_id               รหัสใบเสร็จ
* @apiSuccess {Boolean}    response สถานะ : true,false
* @apiSuccess {String} message ข้อความตอบกลับ
* @apiSuccess {String}     data รหัสใบเสร็จ
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*   {
    "data": 131239, รหัสใบเสร็จ
    "message": "Process Receipt Success.",
    "status": true
}
*/
