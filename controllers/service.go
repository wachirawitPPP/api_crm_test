package controllers

import (
	"fmt"
	"linecrmapi/libs"
	"linecrmapi/models"
	"linecrmapi/structs"
	"time"

	"github.com/gin-gonic/gin"
)

func ItemService(c *gin.Context) { //use
	Service_id := libs.StrToInt(c.Params.ByName("serviceId"))
	// customerId := libs.StrToInt(c.Params.ByName("customerId"))
	var itemservice structs.ServiceList
	if errSCO := models.GetServiceId(Service_id, &itemservice); errSCO != nil || itemservice.Id == 0 {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Service Invalid!",
			"data":    "",
		})
		return
	}
	var Tranfer_is_qty_course int = 1
	var Tranfer_is_qty_product int = 1
	var Adjust_is_qty_course int = 1
	var Adjust_is_day_course int = 1
	var Adjust_is_qty_product int = 1
	var Ser_exp_day int = 0
	if itemservice.Ser_lock_drug == 1 {
		Tranfer_is_qty_course = 0
		Tranfer_is_qty_product = 1
		Adjust_is_qty_course = 0
		Adjust_is_qty_product = 1
	} else {
		Tranfer_is_qty_course = 1
		Tranfer_is_qty_product = 0
		Adjust_is_qty_course = 1
		Adjust_is_qty_product = 0
	}
	if itemservice.Ser_exp == 0 {
		Adjust_is_day_course = 0
	} else {
		if itemservice.Ser_exp_date != nil {
			format := "2006-01-02T15:04:05-07:00"
			now := time.Now()
			thenStr := *itemservice.Ser_exp_date
			then, err := time.Parse(format, thenStr)
			if err != nil {
				fmt.Println("Error parsing date:", err)
				return
			}
			duration := then.Sub(now)
			days := int(duration.Hours() / 24)
			Ser_exp_day = (days + 1)
		}
	}
	var Ser_amount = itemservice.Ser_qty - (itemservice.Ser_use + itemservice.Ser_tranfer)
	if itemservice.Ser_lock_drug == 1 {
		Ser_amount = 1
	}
	if itemservice.Ser_is_active == 2 {
		Ser_amount = 0
	}
	serviceList := structs.ServiceList{
		Id:                     itemservice.Id,
		Course_id:              itemservice.Course_id,
		Ser_code:               itemservice.Ser_code,
		Ser_name:               itemservice.Ser_name,
		Ser_lock_drug:          itemservice.Ser_lock_drug,
		Ser_qty:                itemservice.Ser_qty,
		Ser_unit:               itemservice.Ser_unit,
		Ser_use_date:           itemservice.Ser_use_date,
		Ser_exp:                itemservice.Ser_exp,
		Ser_exp_day:            Ser_exp_day,
		Ser_exp_date:           itemservice.Ser_exp_date,
		Ser_use:                itemservice.Ser_use,
		Ser_tranfer:            itemservice.Ser_tranfer,
		Ser_amount:             Ser_amount,
		Ser_price_total:        itemservice.Ser_price_total,
		Tranfer_is_qty_course:  Tranfer_is_qty_course,
		Tranfer_is_qty_product: Tranfer_is_qty_product,
		Adjust_is_qty_course:   Adjust_is_qty_course,
		Adjust_is_day_course:   Adjust_is_day_course,
		Adjust_is_qty_product:  Adjust_is_qty_product,
	}
	//product
	if itemservice.Ser_lock_drug == 1 {
		var productSet []structs.ServiceProductList
		if errc := models.GetQueueItemCourseServiceProductsList(Service_id, &productSet); errc != nil || len(productSet) == 0 {
			serviceList.Products = &[]structs.ServiceProductList{}
		} else {
			serviceList.Products = &productSet
		}
	} else {
		serviceList.Products = &[]structs.ServiceProductList{}
	}
	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    serviceList,
	})
}

func ItemServiceUsed(c *gin.Context) { //use
	Service_id := libs.StrToInt(c.Params.ByName("serviceId"))
	// customerId := libs.StrToInt(c.Params.ByName("customerId"))
	var itemserviceused []structs.ServiceUsedList
	if errSCO := models.GetServiceUsedList(Service_id, &itemserviceused); errSCO != nil || len(itemserviceused) == 0 {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Service Invalid!",
			"data":    "",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    itemserviceused,
	})
}

func ServiceSearch(c *gin.Context) { //use
	customerId := libs.StrToInt(c.Params.ByName("customerId"))
	var filter structs.ObjPayloadSearchService
	if errSBJ := c.ShouldBindJSON(&filter); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errSBJ.Error(),
		})
		return
	}

	var countList []structs.ServiceSearch
	if filter.ActivePage < 1 {
		filter.ActivePage = 0
	} else {
		filter.ActivePage -= 1
	}

	err := models.GetServiceList(filter, false, &countList, customerId)
	emptySlice := []string{}
	if err != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": err.Error(),
			"data":    emptySlice,
		})
		c.Abort()
	} else {
		var SsList []structs.ServiceSearch
		if errMD := models.GetServiceList(filter, true, &SsList, customerId); errMD != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Something went wrong.",
				"data":    "",
			})
			return
		}

		if len(SsList) == 0 {
			emptyServiceSearch := []structs.ServiceSearch{}
			c.JSON(200, gin.H{
				"status":  true,
				"message": "",
				"data": structs.ResponsePaginationService{
					Result_data:   emptyServiceSearch,
					Count_of_page: 0,
					Count_all:     0,
				},
			})
			return
		}

		ItemList := []structs.ServiceSearch{}
		for _, queryService := range SsList {
			var SerAmount = queryService.SerQty - (queryService.SerUse + queryService.SerTranfer)
			if queryService.SerLockDrug == 1 {
				SerAmount = 1
			}
			if queryService.SerIsActive == 2 {
				SerAmount = 0
			}
			obj := structs.ServiceSearch{
				ID:              queryService.ID,
				ReceiptId:       queryService.ReceiptId,
				RecCode:         queryService.RecCode,
				ReceiptDetailId: queryService.ReceiptDetailId,
				ShopId:          queryService.ShopId,
				CustomerShopId:  queryService.CustomerShopId,
				UserId:          queryService.UserId,
				SerCustomerId:   queryService.SerCustomerId,
				CustomerId:      queryService.CustomerId,
				CourseId:        queryService.CourseId,
				SerTranferId:    queryService.SerTranferId,
				SerCode:         queryService.SerCode,
				SerName:         queryService.SerName,
				SerLockDrug:     queryService.SerLockDrug,
				SerQty:          queryService.SerQty,
				SerUnit:         queryService.SerUnit,
				SerUseDate:      queryService.SerUseDate,
				SerExp:          queryService.SerExp,
				SerExpDate:      queryService.SerExpDate,
				SerUse:          queryService.SerUse,
				SerTranfer:      queryService.SerTranfer,
				SerIsActive:     queryService.SerIsActive,
				SerDatetime:     queryService.SerDatetime,
				SerCreate:       queryService.SerCreate,
				SerUpdate:       queryService.SerUpdate,
				CourseAmount:    queryService.CourseAmount,
				CourseIpd:       queryService.CourseIpd,
				CourseOpd:       queryService.CourseOpd,
				CourseCost:      queryService.CourseCost,
				SerAmount:       SerAmount,
			}
			ItemList = append(ItemList, obj)
		}

		res := structs.ResponsePaginationService{
			Result_data:   ItemList,
			Count_of_page: len(ItemList),
			Count_all:     len(countList),
		}

		c.JSON(200, gin.H{
			"status":  true,
			"message": "",
			"data":    res,
		})
	}

}
