package controllers

import (
	"linecrmapi/libs"
	"linecrmapi/middlewares"
	"linecrmapi/models"
	"linecrmapi/structs"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPrintMedicalCert(c *gin.Context) { //use
	customerId := libs.StrToInt(c.Params.ByName("customerId"))
	mdcId := libs.StrToInt(c.Params.ByName("mdcId"))

	var mdcDetail structs.MedicalCertDetail
	err := models.GetMedicalCertDetailById(mdcId, &mdcDetail)
	if err != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Get medical cert error.",
			"data":    err.Error(),
		})
		return
	}

	if mdcDetail.Id < 1 {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Data Not found!",
			"data":    []string{},
		})
		return
	}

	var RCShop structs.ReceiptShop
	if errShop := models.GetShopReceiptById(mdcDetail.Shop_id, &RCShop); errShop != nil || RCShop.Id == 0 {
		mdcDetail.Shop = structs.ReceiptShop{}
	} else {
		mdcDetail.Shop = RCShop
	}

	var RCCus structs.ObjQueryCustomer
	if errCus := models.GetCustomerById(customerId, &RCCus); errCus != nil || RCCus.ID == 0 {
		mdcDetail.Customer = structs.ObjQueryCustomer{}
	} else {
		mdcDetail.Customer = RCCus
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    mdcDetail,
	})
}

func GetXRayByQue(c *gin.Context) { //use
	customerId := libs.StrToInt(c.Params.ByName("customerId"))
	queueId := libs.StrToInt(c.Params.ByName("queueId"))
	var objResponse structs.XReyDetail

	err := models.GetOPDinReceipt(queueId, &objResponse)
	if err != nil || objResponse.Queue_id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Get data error.",
			"data":    structs.XReyDetail{},
		})
	} else {

		var RCShop structs.ReceiptShop
		if errShop := models.GetShopReceiptById(objResponse.Shop_id, &RCShop); errShop != nil || RCShop.Id == 0 {
			objResponse.Shop = structs.ReceiptShop{}
		} else {
			objResponse.Shop = RCShop
		}

		var RCCus structs.ObjQueryCustomer
		if errCus := models.GetCustomerById(customerId, &RCCus); errCus != nil || RCCus.ID == 0 {
			objResponse.Customer = structs.ObjQueryCustomer{}
		} else {
			objResponse.Customer = RCCus
		}

		var checksData []structs.ObjQueryCheck
		if errCk := models.GetCheckXray(queueId, &checksData); errCk != nil || len(checksData) == 0 {
			objResponse.Checks = &[]structs.ObjQueryCheck{}
		} else {
			objResponse.Checks = &checksData
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Get data successful.",
			"data":    objResponse,
		})
	}

}

func GetLabByQue(c *gin.Context) { //use
	customerId := libs.StrToInt(c.Params.ByName("customerId"))
	queueId := libs.StrToInt(c.Params.ByName("queueId"))
	var objResponse structs.XReyDetail

	err := models.GetOPDinReceipt(queueId, &objResponse)
	if err != nil || objResponse.Queue_id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Get data error.",
			"data":    structs.XReyDetail{},
		})
	} else {

		var RCShop structs.ReceiptShop
		if errShop := models.GetShopReceiptById(objResponse.Shop_id, &RCShop); errShop != nil || RCShop.Id == 0 {
			objResponse.Shop = structs.ReceiptShop{}
		} else {
			objResponse.Shop = RCShop
		}

		var RCCus structs.ObjQueryCustomer
		if errCus := models.GetCustomerById(customerId, &RCCus); errCus != nil || RCCus.ID == 0 {
			objResponse.Customer = structs.ObjQueryCustomer{}
		} else {
			objResponse.Customer = RCCus
		}

		var checksData []structs.ObjQueryCheck
		if errCk := models.GetCheckLab(queueId, &checksData); errCk != nil || len(checksData) == 0 {
			objResponse.Checks = &[]structs.ObjQueryCheck{}
		} else {
			objResponse.Checks = &checksData
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Get data successful.",
			"data":    objResponse,
		})
	}

}

func CreateImageFileForAl(c *gin.Context) {
	var objPayload structs.ObjPayloadCreateImageFileForAI
	if errSBJ := c.ShouldBindJSON(&objPayload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errSBJ,
		})
		return
	}
	if objPayload.FileBase64 != "" {
		filename := "AI_" + middlewares.GenerateDateTimeCode()
		path, size := libs.UploadFileFilenameS3(objPayload.FileBase64, "his-ai", filename)
		objQueryCreateFile := models.QueueFile{
			QueueId:    0,
			QuefPath:   path,
			QuefSize:   size,
			QuefIsUse:  1,
			QuefModify: "",
		}
		if size == 0 {
			c.JSON(200, gin.H{
				"status":  false,
				"message": "Upload File to S3 error.",
				"data":    "",
			})
			return
		}

		c.JSON(200, gin.H{
			"status":  true,
			"message": "Create file success.",
			"data":    objQueryCreateFile,
		})
		return
	} else {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Create file error.",
			"data":    "",
		})
		return
	}

}
