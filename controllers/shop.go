package controllers

import (
	"linecrmapi/libs"
	"linecrmapi/middlewares"
	"linecrmapi/models"
	"linecrmapi/structs"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

func ShopOauth(c *gin.Context) {

	var payload structs.PayloadShopOauth

	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	lineId := c.Params.ByName("lineId")

	// get customer online by line id
	var customerOnline structs.CustomerOnline
	if errCO := models.GetCustomerOnlineByLineId(lineId, &customerOnline); errCO != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Customer Online Invalid!",
			"data":    errCO.Error(),
		})
		return
	}
	if customerOnline.ID == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Customer Online Invalid!",
			"data":    "",
		})
		return
	}

	// get shop
	var shop structs.ShopReadResponse
	if errS := models.GetShopById(payload.Shop_id, &shop); errS != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Shop Invalid!",
			"data":    errS.Error(),
		})
		return
	}

	if shop.Id == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Shop Invalid!",
			"data":    "",
		})
		return
	}

	// get customer first in shop by citizen id
	var customer models.Customer
	if errC := models.GetShopCustomerById(shop.ShopMotherId, customerOnline.Co_citizen_id, &customer); errC != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Customer Invalid!",
			"data":    errC.Error(),
		})
		return
	}
	if customer.ID == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Customer Shop Invalid!",
			"data":    "",
		})
		return
	}

	var coActk = middlewares.CreateAccessToken(customerOnline.ID, customer.ID, shop.Id, shop.ShopMotherId, customerOnline.Co_line_id, customerOnline.Co_citizen_id)

	//add log
	if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
		Line_id:        customerOnline.Co_line_id,
		Log_ip_address: c.ClientIP(),
		Log_browser:    c.GetHeader("User-Agent"),
		Log_text:       "Login Shop by Shop ID",
		Log_create:     time.Now().Format("2006-01-02 15:04:05"),
	}); errL != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Create Log Customer Error.",
			"data":    "",
		})
		return
	}

	responseData := structs.ResponseOauth{
		AccessToken: coActk,
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    responseData,
	})
}

func GetInShopList(c *gin.Context) {
	citizenId := c.Params.ByName("citizenId")
	// mother_id := []int64{}

	mother_id, errm := models.GetInShopList(citizenId)
	if errm != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Get shop mother error!",
			"data":    errm,
		})
	} else {
		var objShopList []structs.InShopList
		err := models.GetInShopListByMother(mother_id, &objShopList)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "Get data error!",
				"data":    err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  true,
				"message": "Get Data Success",
				"data":    objShopList,
			})
		}
	}
}

func GetInShopListOnline(c *gin.Context) {
	// var payload structs.PayloadShopListOnline

	// if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
	// 	c.AbortWithStatusJSON(200, gin.H{
	// 		"status":  false,
	// 		"message": "Invalid request data.",
	// 		"data":    "",
	// 	})
	// 	return
	// }
	citizenId := c.Params.ByName("citizenId")

	mother_id, errm := models.GetInShopList(citizenId)
	if errm != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Get shop mother error!",
			"data":    errm,
		})
	} else {
		var objShopList []structs.InShopList
		err := models.GetInShopListByMother(mother_id, &objShopList)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "Get data error!",
				"data":    err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  true,
				"message": "Get Data Success",
				"data":    objShopList,
			})
		}
	}
}

func ShopOauthOnline(c *gin.Context) {

	var payload structs.PayloadShopOauth

	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	customerOnlineId := libs.StrToInt(c.Params.ByName("customerOnlineId"))

	// get customer online by line id
	var customerOnline structs.CustomerOnline
	if errCO := models.GetCustomerOnlineById(customerOnlineId, &customerOnline); errCO != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Customer Online Invalid!",
			"data":    errCO.Error(),
		})
		return
	}
	if customerOnline.ID == 0 || customerOnline.Co_citizen_id == "" {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Customer Online Invalid!",
			"data":    "",
		})
		return
	}

	// get shop
	var shop structs.ShopReadResponse
	if errS := models.GetShopById(payload.Shop_id, &shop); errS != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Shop Invalid!",
			"data":    errS.Error(),
		})
		return
	}

	if shop.Id == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Shop Invalid!",
			"data":    "",
		})
		return
	}

	// get customer first in shop by citizen id
	var customer models.Customer
	if errC := models.GetShopCustomerById(shop.ShopMotherId, customerOnline.Co_citizen_id, &customer); errC != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Customer Invalid!",
			"data":    errC.Error(),
		})
		return
	}
	if customer.ID == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Customer Shop Invalid!",
			"data":    "",
		})
		return
	}

	var coActk = middlewares.CreateAccessToken(customerOnline.ID, customer.ID, shop.Id, shop.ShopMotherId, customerOnline.Co_line_id, customerOnline.Co_citizen_id)

	//add log
	if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
		Line_id:        customerOnline.Co_line_id,
		Log_ip_address: c.ClientIP(),
		Log_browser:    c.GetHeader("User-Agent"),
		Log_text:       "Login Shop by Shop ID",
		Log_create:     time.Now().Format("2006-01-02 15:04:05"),
	}); errL != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Create Log Customer Error.",
			"data":    "",
		})
		return
	}

	responseData := structs.ResponseOauth{
		AccessToken: coActk,
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    responseData,
	})
}
