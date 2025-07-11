package routes

import (
	"linecrmapi/controllers"
	"linecrmapi/middlewares"

	"github.com/gin-gonic/gin"
)

func SetRouterOrderonline(router *gin.Engine) {
	r := router.Group("/orderonline")
	{
		r.GET("/:id", middlewares.CheckAccessToken, controllers.GetOrderOnlineItem)
		r.GET("/history/:id", middlewares.CheckAccessToken, controllers.GetCustomerOrderOnlineHistory)
		r.POST("", middlewares.CheckAccessToken, controllers.AddOrderOnline)
		r.POST("/search", middlewares.CheckAccessToken, controllers.SearchOrderOnline)
		r.POST("/payment", middlewares.CheckAccessToken, controllers.PaymentOrderOnline)
		r.PUT("/:id", middlewares.CheckAccessToken, controllers.UpdateOrderOnline)
		r.DELETE("/cancel/:id", middlewares.CheckAccessToken, controllers.CancelOrderOnline)
		r.DELETE("/detail/:id", middlewares.CheckAccessToken, controllers.DeleteOrderOnlineDetail)
	}
	return
}
