package controllers

import (
	"linecrmapi/libs"
	"linecrmapi/middlewares"
	"linecrmapi/models"
	"linecrmapi/structs"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddOrderOnline(c *gin.Context) {
	var payload structs.ObjPayloadCreateOrderOnline
	var createObj structs.OrderOnline = structs.ConvertPayloadToOrderOnline(payload)
	if errPL := c.ShouldBindJSON(&payload); errPL != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errPL.Error(),
		})
		return
	}

	if payload.Or_website_id != 0 {
		if errShop := models.CreateOrderOnline(&createObj); errShop != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Something went wrong.",
				"data":    errShop.Error(),
			})
			return
		}
		if createObj.Id > 0 {
			var detailList []structs.OrderOnlineDetail
			for _, item := range payload.Or_detail {
				var createObjDetail structs.OrderOnlineDetail = structs.ConvertPayloadDetailToOrderDetail(item, createObj.Id)
				detailList = append(detailList, createObjDetail)
			}

			if len(detailList) > 0 {
				if errShop := models.CreateOrderOnlineDetail(&detailList); errShop != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Something went wrong.",
						"data":    errShop.Error(),
					})
					return
				}
			}
		}
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Add order success",
		"data":    createObj.Id,
	})
}

func UpdateOrderOnline(c *gin.Context) {
	var payload structs.ObjGetOrderOnline
	if errPL := c.ShouldBindJSON(&payload); errPL != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errPL.Error(),
		})
		return
	}

	if payload.Id > 0 {
		var order structs.OrderOnline
		if err1 := models.GetOrderOnline(payload.Id, &order); err1 != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Something went wrong.",
				"data":    err1.Error(),
			})
			return
		}
		var orderDetail []structs.OrderOnlineDetail
		if err2 := models.GetOrderOnlineDetailById(payload.Id, &orderDetail); err2 != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Something went wrong.",
				"data":    err2.Error(),
			})
			return
		}
		//check change value
		var initOrder structs.UpdateOrderOnline = structs.ConvertOrderToUpdateOrder(order)
		var updateObj structs.UpdateOrderOnline = structs.ConvertObjGetToUpdateOrder(payload)
		if !reflect.DeepEqual(initOrder, updateObj) {
			if errShop := models.UpdateOrderOnline(&updateObj); errShop != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Something went wrong.",
					"data":    errShop.Error(),
				})
				return
			}
		}

		if len(payload.Or_detail) > 0 {
			updateList, createList := structs.ConvertObjGetDetailsToUpdateDetails(payload.Or_detail)
			if len(updateList) > 0 {
				for _, v := range updateList {
					var change bool = false
					if len(orderDetail) > 0 {
						for _, d := range orderDetail {
							if v.Id == d.Id {
								v.Ord_update = ""
								initOrderDetail := structs.ConvertToUpdateOrderOnlineDetail(d)
								change = !reflect.DeepEqual(initOrderDetail, v)
							}
						}
					}
					if change {
						if errShop := models.UpdateOrderOnlineDetail(&v); errShop != nil {
							c.AbortWithStatusJSON(200, gin.H{
								"status":  false,
								"message": "Something went wrong.",
								"data":    errShop.Error(),
							})
							return
						}
					}
				}
			}

			if len(createList) > 0 {
				if errShop := models.CreateOrderOnlineDetail(&createList); errShop != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Something went wrong.",
						"data":    errShop.Error(),
					})
					return
				}
			}
		}
	} else {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Order Invalid!",
			"data":    "",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Update order success",
		"data":    "",
	})
}

func SearchOrderOnline(c *gin.Context) {
	var payload structs.ObjPayloadSearchOrderOnline
	var resOrder []structs.ObjGetOrderOnline
	if errPL := c.ShouldBindJSON(&payload); errPL != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errPL.Error(),
		})
		return
	}

	var counterAll []structs.OrderOnline
	if err5 := models.SearchOrderOnline(structs.ObjPayloadSearchOrderOnline{
		Order_online_id: -1,
		Search_text:     "",
		Or_website_id:   -1,
		Payment_status:  -1,
		Or_date:         "",
		Or_is_active:    -1,
		Active_page:     0,
		Per_page:        0,
	}, &counterAll); err5 != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    err5.Error(),
		})
		return
	}
	var orderlist []structs.OrderOnline
	if err1 := models.SearchOrderOnline(payload, &orderlist); err1 != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    err1.Error(),
		})
		return
	}

	if len(orderlist) > 0 {
		var orderIds []int
		for _, item := range orderlist {
			orderIds = append(orderIds, item.Id)
		}

		var orderDetail []structs.OrderOnlineDetail
		if err2 := models.GetOrderOnlineDetailByIds(orderIds, payload.Or_is_active, &orderDetail); err2 != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Something went wrong.",
				"data":    err2.Error(),
			})
			return
		}

		if len(orderDetail) > 0 {
			for _, item := range orderlist {
				var details []structs.OrderOnlineDetail
				for _, itemDetail := range orderDetail {
					if item.Id == itemDetail.Order_online_id {
						details = append(details, itemDetail)
					}
				}
				temp := structs.ConvertToObjGetOrderOnline(item, details)
				resOrder = append(resOrder, temp)
			}
		}
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "get order success",
		"data": structs.ObjPaginationOrderOnline{
			Result_data:   resOrder,
			Count_of_page: len(resOrder),
			Count_all:     len(counterAll),
		},
	})
}

func GetOrderOnlineItem(c *gin.Context) {
	order_online_id := libs.StrToInt(c.Params.ByName("id"))
	var resOrder []structs.ObjGetOrderOnline
	var orderlist []structs.OrderOnline
	if err1 := models.SearchOrderOnline(structs.ObjPayloadSearchOrderOnline{
		Order_online_id: order_online_id,
		Search_text:     "",
		Or_website_id:   -1,
		Payment_status:  -1,
		Or_date:         "",
		Or_is_active:    -1,
		Active_page:     1,
		Per_page:        1,
	}, &orderlist); err1 != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    err1.Error(),
		})
		return
	}

	if len(orderlist) > 0 {
		var orderIds []int
		for _, item := range orderlist {
			orderIds = append(orderIds, item.Id)
		}

		var orderDetail []structs.OrderOnlineDetail
		if err2 := models.GetOrderOnlineDetailByIds(orderIds, -1, &orderDetail); err2 != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Something went wrong.",
				"data":    err2.Error(),
			})
			return
		}

		if len(orderDetail) > 0 {
			for _, item := range orderlist {
				var details []structs.OrderOnlineDetail
				for _, itemDetail := range orderDetail {
					if item.Id == itemDetail.Order_online_id {
						details = append(details, itemDetail)
					}
				}
				temp := structs.ConvertToObjGetOrderOnline(item, details)
				resOrder = append(resOrder, temp)
			}
		}
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "get order success",
		"data":    resOrder,
	})
}

func CancelOrderOnline(c *gin.Context) {
	order_online_id := libs.StrToInt(c.Params.ByName("id"))
	bearer := c.Request.Header.Get("Authorization")
	actkToken := strings.TrimPrefix(bearer, "Bearer ")
	token, err6 := middlewares.DecodeToken(actkToken)
	if err6 != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    err6.Error(),
		})
		return
	}
	payload := structs.ObjCancelOrderOnline{Order_online_id: order_online_id, User_id_cancel: token.CustomerOnlineID}

	if err1 := models.CancelOrderOnline(payload); err1 != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    err1.Error(),
		})
		return
	}

	if err2 := models.CancelOrderOnlineDetail(payload.Order_online_id); err2 != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    err2.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Cancel order success",
		"data":    "",
	})
}

func DeleteOrderOnlineDetail(c *gin.Context) {
	order_online_id := libs.StrToInt(c.Params.ByName("id"))
	if err1 := models.DeleteOrderOnlineDetail(order_online_id); err1 != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    err1.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Delete order success",
		"data":    "",
	})
}

func PaymentOrderOnline(c *gin.Context) {
	var payload structs.ObjPaymentOrderOnline
	if errPL := c.ShouldBindJSON(&payload); errPL != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errPL.Error(),
		})
		return
	}

	if payload.Order_online_id != 0 {
		if err1 := models.PaymentOrderOnline(&payload); err1 != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Something went wrong.",
				"data":    err1.Error(),
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Update order success",
		"data":    "",
	})
}

func GetCustomerOrderOnlineHistory(c *gin.Context) {
	customer_online_id := libs.StrToInt(c.Params.ByName("id"))
	var resOrder []structs.ObjGetOrderOnline
	var orderlist []structs.OrderOnline
	if err1 := models.GetCustomerOrderOnlineHistory(customer_online_id, &orderlist); err1 != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    err1.Error(),
		})
		return
	}

	if len(orderlist) > 0 {
		var orderIds []int
		for _, item := range orderlist {
			orderIds = append(orderIds, item.Id)
		}

		var orderDetail []structs.OrderOnlineDetail
		if err2 := models.GetOrderOnlineDetailByIds(orderIds, -1, &orderDetail); err2 != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Something went wrong.",
				"data":    err2.Error(),
			})
			return
		}

		if len(orderDetail) > 0 {
			for _, item := range orderlist {
				var details []structs.OrderOnlineDetail
				for _, itemDetail := range orderDetail {
					if item.Id == itemDetail.Order_online_id {
						details = append(details, itemDetail)
					}
				}
				temp := structs.ConvertToObjGetOrderOnline(item, details)
				resOrder = append(resOrder, temp)
			}
		}
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "get order success",
		"data":    resOrder,
	})
}
