package routes

import (
	"linecrmapi/controllers"
	"linecrmapi/middlewares"

	"github.com/gin-gonic/gin"
)

func SetRouterSticker(router *gin.Engine) {
	r := router.Group("/sticker")
	{
		r.POST("/search", middlewares.CheckAccessToken, controllers.STKSearch)
		r.POST("/list", middlewares.CheckAccessToken, controllers.GetSTKList)
		r.POST("/detail/list", middlewares.CheckAccessToken, controllers.GetSTKDetailList)
		r.POST("/druglabel", middlewares.CheckAccessToken, controllers.CreatePrintDruglabel)
		r.GET("/:invoice_id", middlewares.CheckAccessToken, controllers.GetSTKDetail)
		r.GET("noinvoice/:sticker_id/:actv_id", middlewares.CheckAccessToken, controllers.GetSTKDetail)
		r.GET("/prescription/:invoice_id", middlewares.CheckAccessToken, controllers.GetPrescription)
		r.GET("/setting", middlewares.CheckAccessToken, controllers.GetSettingSticker)
		r.PUT("/setting", middlewares.CheckAccessToken, controllers.UpdateSettingSticker)
		r.PUT("", middlewares.CheckAccessToken, controllers.UpdateSticker)
	}
	return
}
