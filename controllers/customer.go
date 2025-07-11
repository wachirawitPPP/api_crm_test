package controllers

import (
	"encoding/json"
	"linecrmapi/libs"
	"linecrmapi/models"
	"linecrmapi/structs"
	"os"
	"strconv"

	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetCustomerById(c *gin.Context) { //use

	customerId := libs.StrToInt(c.Params.ByName("customerId"))
	shopId := libs.StrToInt(c.Params.ByName("shopId"))

	var objQueryCustomer structs.ObjQueryCustomer
	var objQueryTag []structs.ObjQueryCustomerTag
	var objQueryFamily []structs.ObjQueryCustomerFamily
	var objQueryContact []structs.ObjQueryCustomerContact
	var paymentBalance structs.DashboardPayment
	errCustomer := models.GetCustomerById(customerId, &objQueryCustomer)
	errTag := models.GetCustomerTagById(customerId, &objQueryTag)
	errFamily := models.GetCustomerFamilyById(customerId, &objQueryFamily)
	errContact := models.GetCustomerContactById(customerId, &objQueryContact)
	errBalance := models.GetCustomerPaymentBalance(shopId, customerId, &paymentBalance)
	if errCustomer != nil && errTag != nil && errFamily != nil && errContact != nil && errBalance != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Get data error.",
			"data":    "",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Get data successful.",
			"data": structs.ObjResponseGetCustomerById{
				Customer: &objQueryCustomer,
				Tag:      objQueryTag,
				Family:   objQueryFamily,
				Contact:  objQueryContact,
				Balance:  &paymentBalance,
			},
		})
	}

}
func GetCustomerOnlineById(c *gin.Context) {
	customerId := libs.StrToInt(c.Params.ByName("customerOnlineId"))
	var objQueryCustomerOnline structs.ObjGetCustomerOnline

	errCustomer := models.GetCustomerOnlinesById(customerId, &objQueryCustomerOnline)
	if errCustomer != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Get data error.",
			"data":    "",
		})
		return
	}

	// Use citizen_id for internal check
	isOnlineSync := objQueryCustomerOnline.CoCitizenId != ""

	// Marshal to JSON and unmarshal to map
	customerJSON, _ := json.Marshal(objQueryCustomerOnline)

	// Convert to map[string]interface{} for filtering
	var data map[string]interface{}
	json.Unmarshal(customerJSON, &data)

	// Remove citizen_id before returning
	delete(data, "citizen_id")

	// Add custom field
	data["is_online_data_sync"] = isOnlineSync

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Get data successful.",
		"data":    data,
	})
}

func CheckCitizenId(c *gin.Context) {

	var objPayload structs.ObjPayloadCheckCitizenId

	if errSBJ := c.ShouldBindJSON(&objPayload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	var objQueryShopMother models.Shop
	errShopMother := models.GetShopMother_(objPayload.ShopId, &objQueryShopMother)
	if errShopMother != nil || objQueryShopMother.ShopMotherId == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Get data error.",
			"data":    "",
		})
		return
	}

	var objQueryCustomer structs.CheckCustomer
	errCustomer := models.GetCustomer_(objQueryShopMother.ShopMotherId, objPayload.CustomerId, objPayload.CitizenId, &objQueryCustomer)
	if errCustomer != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Get data error.",
			"data":    "",
		})
		return
	}

	if objQueryCustomer.ID != 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": fmt.Sprintf("Citizen ID \"%s\" duplicate in shop \"%s\"", objPayload.CitizenId, objQueryCustomer.ShopName),
			"data":    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Get data successful.",
		"data":    "",
	})

}

func UpdateCustomer(c *gin.Context) {

	customerId := libs.StrToInt(c.Params.ByName("customerId"))
	var objPayload structs.ObjPayloadUpdateCustomer

	if errSBJ := c.ShouldBindJSON(&objPayload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errSBJ,
		})
		return
	}

	var objQueryCustomer structs.ObjQueryCustomer
	var objQueryTag []structs.ObjQueryCustomerTag
	var objQueryFamily []structs.ObjQueryCustomerFamily
	var objQueryContact []structs.ObjQueryCustomerContact

	errCustomer := models.GetCustomerById(customerId, &objQueryCustomer)
	errTag := models.GetCustomerTagById(customerId, &objQueryTag)
	errFamily := models.GetCustomerFamilyById(customerId, &objQueryFamily)
	errContact := models.GetCustomerContactById(customerId, &objQueryContact)
	if errCustomer != nil && errTag != nil && errFamily != nil && errContact != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Update customer error.",
			"data":    "",
		})
		return
	}

	ctmImage := objQueryCustomer.CtmImage
	ctmImageSize := objQueryCustomer.CtmImageSize

	if objPayload.CtmImage != "" {
		var objQueryShop models.Shop
		errShopMother := models.GetShopCode(objPayload.ShopId, &objQueryShop)
		if errShopMother != nil || objQueryShop.ID == 0 {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Get data error.",
				"data":    "",
			})
			return
		}
		filename := objQueryShop.ShopCode + "_" + objQueryCustomer.CtmId + "_" + time.Now().Format("20060102150405")
		ctmImage, ctmImageSize = libs.UploadFileFilenameS3(objPayload.CtmImage, "customer", filename)
	}

	objQueryUpdateCustomer := structs.ObjQueryUpdateCustomer{
		CustomerGroupId:     objPayload.CustomerGroupId,
		CtmCitizenId:        objPayload.CtmCitizenId,
		CtmPassportId:       objPayload.CtmPassportId,
		CtmPrefix:           objPayload.CtmPrefix,
		CtmFname:            objPayload.CtmFname,
		CtmLname:            objPayload.CtmLname,
		CtmNname:            objPayload.CtmNname,
		CtmFnameEn:          objPayload.CtmFnameEn,
		CtmLnameEn:          objPayload.CtmLnameEn,
		CtmGender:           objPayload.CtmGender,
		CtmNation:           objPayload.CtmNation,
		CtmReligion:         objPayload.CtmReligion,
		CtmEduLevel:         objPayload.CtmEduLevel,
		CtmMaritalStatus:    objPayload.CtmMaritalStatus,
		CtmBlood:            objPayload.CtmBlood,
		CtmEmail:            objPayload.CtmEmail,
		CtmTel:              objPayload.CtmTel,
		CtmTel_2:            objPayload.CtmTel_2,
		CtmBirthdate:        objPayload.CtmBirthdate,
		CtmAddress:          objPayload.CtmAddress,
		CtmDistrict:         objPayload.CtmDistrict,
		CtmAmphoe:           objPayload.CtmAmphoe,
		CtmProvince:         objPayload.CtmProvince,
		CtmZipcode:          objPayload.CtmZipcode,
		CtmWeight:           objPayload.CtmWeight,
		CtmHeight:           objPayload.CtmHeight,
		CtmWaistline:        objPayload.CtmWaistline,
		CtmChest:            objPayload.CtmChest,
		CtmTreatmentType:    objPayload.CtmTreatmentType,
		RightTreatmentId:    objPayload.RightTreatmentId,
		CtmAllergic:         objPayload.CtmAllergic,
		CtmMentalHealth:     objPayload.CtmMentalHealth,
		CtmDisease:          objPayload.CtmDisease,
		CtmHealthComment:    objPayload.CtmHealthComment,
		CtmComment:          objPayload.CtmComment,
		CtmImage:            ctmImage,
		CtmImageSize:        ctmImageSize,
		CtmPoint:            objPayload.CtmPoint,
		CompanyName:         objPayload.CompanyName,
		CompanyTax:          objPayload.CompanyTax,
		CompanyTel:          objPayload.CompanyTel,
		CompanyEmail:        objPayload.CompanyEmail,
		CompanyAddress:      objPayload.CompanyAddress,
		CompanyDistrict:     objPayload.CompanyDistrict,
		CompanyAmphoe:       objPayload.CompanyAmphoe,
		CompanyProvince:     objPayload.CompanyProvince,
		CompanyZipcode:      objPayload.CompanyZipcode,
		CtmSubscribeOpd:     objPayload.CtmSubscribeOpd,
		CtmSubscribeLab:     objPayload.CtmSubscribeLab,
		CtmSubscribeCert:    objPayload.CtmSubscribeCert,
		CtmSubscribeReceipt: objPayload.CtmSubscribeReceipt,
		CtmSubscribeAppoint: objPayload.CtmSubscribeAppoint,
		CtmIsActive:         objPayload.CtmIsActive,
		CtmUpdate:           time.Now().Format("2006-01-02 15:04:05"),
	}

	errCreate := models.UpdateCustomer(customerId, &objQueryUpdateCustomer)
	if errCreate != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Create customer error.",
			"data":    "",
		})
		return
	}

	// tag
	if len(objPayload.TagSelected) > 0 {
		// delete
		var newIds = objPayload.TagSelected
		var oldIds []int
		for _, item := range objQueryTag {
			oldIds = append(oldIds, item.TagId)
		}
		deleteIds := libs.GetArrayIntDiff(newIds, oldIds)
		if len(deleteIds) > 0 {
			objQueryDeleteCustomerTag := models.CustomerTag{}
			errDeleteCustomerTagBatch := models.DeleteCustomerTagBatch(customerId, deleteIds, &objQueryDeleteCustomerTag)
			if errDeleteCustomerTagBatch != nil {
				fmt.Println(errDeleteCustomerTagBatch)
			}
		}
		// create
		objQueryCreateCustomerTag := []models.CustomerTag{}
		for _, item_1 := range objPayload.TagSelected {
			var isExit = 0
			for _, item_2 := range objQueryTag {
				if item_1 == item_2.TagId {
					isExit = 1
					break
				}
			}
			if isExit == 0 {
				objQueryCreateCustomerTag = append(objQueryCreateCustomerTag, models.CustomerTag{
					CustomerId: customerId,
					TagId:      item_1,
				})
			}
		}
		if len(objQueryCreateCustomerTag) > 0 {
			// create
			errCreateCustomerTag := models.CreateCustomerTagBatch(&objQueryCreateCustomerTag)
			if errCreateCustomerTag != nil {
				fmt.Println(errCreateCustomerTag)
			}
		}
	} else {
		if len(objQueryTag) > 0 {
			// clear
			objQueryDeleteCustomerTag := models.CustomerTag{}
			errClearCustomerTag := models.ClearCustomerTag(customerId, &objQueryDeleteCustomerTag)
			if errClearCustomerTag != nil {
				fmt.Println(errClearCustomerTag)
			}

		}
	}

	// family
	if len(objPayload.FamilySelected) > 0 {
		// delete
		var newIds []int
		for _, item := range objPayload.FamilySelected {
			newIds = append(newIds, item.ID)
		}
		var oldIds []int
		for _, item := range objQueryFamily {
			oldIds = append(oldIds, item.CfCustomerId)
		}
		deleteIds := libs.GetArrayIntDiff(newIds, oldIds)
		if len(deleteIds) > 0 {
			// delete
			objQueryDeleteCustomerFamily := models.CustomerFamilys{}
			errDeleteCustomerFamilyBatch := models.DeleteCustomerFamilyBatch(customerId, deleteIds, &objQueryDeleteCustomerFamily)
			if errDeleteCustomerFamilyBatch != nil {
				fmt.Println(errDeleteCustomerFamilyBatch)
			}
		}
		objQueryCreateCustomerFamily := []models.CustomerFamilys{}
		for _, item_1 := range objPayload.FamilySelected {
			var isExit = 0
			for _, item_2 := range objQueryFamily {
				if item_1.ID == item_2.CfCustomerId {
					isExit = 1
					break
				}
			}
			if isExit == 0 {
				objQueryCreateCustomerFamily = append(objQueryCreateCustomerFamily, models.CustomerFamilys{
					CustomerId:   customerId,
					CfCustomerId: item_1.ID,
					CfRelation:   item_1.Relation,
					CfCreate:     time.Now().Format("2006-01-02 15:04:05"),
					CfUpdate:     time.Now().Format("2006-01-02 15:04:05"),
				})
			}
		}
		if len(objQueryCreateCustomerFamily) > 0 {
			// create
			errCreateCustomerFamily := models.CreateCustomerFamilyBatch(&objQueryCreateCustomerFamily)
			if errCreateCustomerFamily != nil {
				fmt.Println(errCreateCustomerFamily)
			}
		}
	} else {
		if len(objQueryFamily) > 0 {
			// clear
			objQueryDeleteCustomerFamily := models.CustomerFamilys{}
			errClearCustomerFamily := models.ClearCustomerFamily(customerId, &objQueryDeleteCustomerFamily)
			if errClearCustomerFamily != nil {
				fmt.Println(errClearCustomerFamily)
			}
		}
	}

	// contact
	if len(objPayload.ContactSelected) > 0 {
		// delete
		var newIds []string
		for _, item := range objPayload.ContactSelected {
			newIds = append(newIds, item.Name)
		}
		var oldIds []string
		for _, item := range objQueryContact {
			oldIds = append(oldIds, item.CcName)
		}
		deleteIds := libs.GetArrayStrDiff(newIds, oldIds)
		if len(deleteIds) > 0 {
			// delete
			objQueryDeleteCustomerContact := models.CustomerContact{}
			errDeleteCustomerContactBatch := models.DeleteCustomerContactBatch(customerId, deleteIds, &objQueryDeleteCustomerContact)
			if errDeleteCustomerContactBatch != nil {
				fmt.Println(errDeleteCustomerContactBatch)
			}
		}
		objQueryCreateCustomerContact := []models.CustomerContact{}
		for _, item_1 := range objPayload.ContactSelected {
			var isExit = 0
			for _, item_2 := range objQueryContact {
				if item_1.Name == item_2.CcName {
					isExit = 1
					break
				}
			}
			if isExit == 0 {
				objQueryCreateCustomerContact = append(objQueryCreateCustomerContact, models.CustomerContact{
					CustomerId: customerId,
					CcName:     item_1.Name,
					CcTel:      item_1.Tel,
					CcRelation: item_1.Relation,
					CcCreate:   time.Now().Format("2006-01-02 15:04:05"),
					CcUpdate:   time.Now().Format("2006-01-02 15:04:05"),
				})
			}
		}
		if len(objQueryCreateCustomerContact) > 0 {
			// create
			errCreateCustomerContact := models.CreateCustomerContactBatch(&objQueryCreateCustomerContact)
			if errCreateCustomerContact != nil {
				fmt.Println(errCreateCustomerContact)
			}
		}
	} else {
		if len(objQueryContact) > 0 {
			// clear
			objQueryDeleteCustomerContact := models.CustomerContact{}
			errClearCustomerContact := models.ClearCustomerContact(customerId, &objQueryDeleteCustomerContact)
			if errClearCustomerContact != nil {
				fmt.Println(errClearCustomerContact)
			}
		}
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Update customer success.",
		"data":    customerId,
	})

}

func GetCustomerOpdPagination(c *gin.Context) { //use
	customerId := libs.StrToInt(c.Params.ByName("customerId"))
	shopId := libs.StrToInt(c.Params.ByName("shopId"))
	var objPayload structs.ObjPayloadGetCustomerOpdPagination

	if errSBJ := c.ShouldBindJSON(&objPayload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	shopMId := shopId
	if libs.StrToInt(c.Params.ByName("shopMotherId")) > 0 {
		shopMId = libs.StrToInt(c.Params.ByName("shopMotherId"))
	}
	shopId = shopMId

	if objPayload.CurrentPage < 1 {
		objPayload.CurrentPage = 0
	} else {
		objPayload.CurrentPage -= 1
	}

	countAll, errCountAll := models.CountCustomerOpdPagination(&objPayload, customerId, shopId)
	if errCountAll != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Get data error.",
			"data":    0,
		})
		return
	}

	var objQueryCustomerOpd []structs.ObjQueryGetCustomerOpdPagination
	errCustomerOpd := models.GetCustomerOpdPagination(&objPayload, &objQueryCustomerOpd, customerId, shopId)
	if errCustomerOpd != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Get data error.",
			"data": structs.ObjResponseGetCustomerOpdPagination{
				Items:     []structs.ObjQueryGetCustomerOpdPagination{},
				CountPage: 0,
				CountAll:  0,
			},
		})
		return
	} else {
		if len(objQueryCustomerOpd) > 0 {
			var opdIds []int
			for i, opd := range objQueryCustomerOpd {
				var objQueryCustomerOpdShop structs.Shop
				errCustomerOpdShop := models.GetCustomerOpdShopByIds(opd.ShopId, &objQueryCustomerOpdShop)
				if errCustomerOpdShop != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Get data error.",
						"data":    "",
					})
					return
				}
				objQueryCustomerOpd[i].ShopName = objQueryCustomerOpdShop.ShopName
				opdIds = append(opdIds, opd.OpdId)
			}
			var objQueryCustomerOpdDiagnostic []models.OpdDiagnostic
			errQueryCustomerOpdDiagnostic := models.GetCustomerOpdDiagnosticByIds(opdIds, &objQueryCustomerOpdDiagnostic)
			if errQueryCustomerOpdDiagnostic != nil {
				fmt.Println(errQueryCustomerOpdDiagnostic)
			}
			c.JSON(http.StatusOK, gin.H{
				"status":  true,
				"message": "Get data successful.",
				"data": structs.ObjResponseGetCustomerOpdPagination{
					Items:     objQueryCustomerOpd,
					Subs:      objQueryCustomerOpdDiagnostic,
					CountPage: len(objQueryCustomerOpd),
					CountAll:  countAll,
				},
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  true,
				"message": "Get data successful.",
				"data": structs.ObjResponseGetCustomerOpdPagination{
					Items:     objQueryCustomerOpd,
					Subs:      []models.OpdDiagnostic{},
					CountPage: len(objQueryCustomerOpd),
					CountAll:  countAll,
				},
			})
		}
	}

}

func GetCustomerCheckPagination(c *gin.Context) { //use
	// customerId := libs.StrToInt(c.Params.ByName("customerId"))
	var objPayload structs.ObjPayloadGetCustomerCheckPagination

	if errSBJ := c.ShouldBindJSON(&objPayload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	if objPayload.CurrentPage < 1 {
		objPayload.CurrentPage = 0
	} else {
		objPayload.CurrentPage -= 1
	}

	countAll, errCountAll := models.CountCustomerCheckPagination(&objPayload)
	if errCountAll != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Get data error.",
			"data":    0,
		})
		return
	}

	var objQueryCheck []structs.ObjQueryGetCustomerCheckPagination
	errQueryCheck := models.GetCustomerCheckPagination(&objPayload, &objQueryCheck)
	if errQueryCheck != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Get data error.",
			"data": structs.ObjResponseGetCustomerCheckPagination{
				Items:     []structs.ObjQueryGetCustomerCheckPagination{},
				CountPage: 0,
				CountAll:  0,
			},
		})
		return
	}
	if len(objQueryCheck) > 0 {
		for i, check := range objQueryCheck {
			var objQueryCustomerCheckShop structs.Shop
			errCustomerCheckShop := models.GetCustomerOpdShopByIds(check.ShopId, &objQueryCustomerCheckShop)
			if errCustomerCheckShop != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Get data error.",
					"data":    "",
				})
				return
			}
			objQueryCheck[i].ShopName = objQueryCustomerCheckShop.ShopName
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Get data successful.",
		"data": structs.ObjResponseGetCustomerCheckPagination{
			Items:     objQueryCheck,
			CountPage: len(objQueryCheck),
			CountAll:  countAll,
		},
	})

}

func GetCustomerReceiptPagination(c *gin.Context) { //use
	customerId := libs.StrToInt(c.Params.ByName("customerId"))
	shopId := libs.StrToInt(c.Params.ByName("shopId"))
	var filter structs.ObjPayloadSearchReceiptCustomer
	if errSBJ := c.ShouldBindJSON(&filter); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errSBJ.Error(),
		})
		return
	}

	shopMId := libs.StrToInt(c.Params.ByName("shopId"))
	if libs.StrToInt(c.Params.ByName("shopMotherId")) > 0 {
		shopMId = libs.StrToInt(c.Params.ByName("shopMotherId"))
	}
	shopId = shopMId

	var countList []structs.ReceiptListCustomer
	if errMD := models.GetReceiptCustomerList(filter, false, &countList, customerId, shopId); errMD != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    "",
		})
		return
	}

	if filter.ActivePage < 1 {
		filter.ActivePage = 0
	} else {
		filter.ActivePage -= 1
	}

	var RCList []structs.ReceiptListCustomer
	if errMD := models.GetReceiptCustomerList(filter, true, &RCList, customerId, shopId); errMD != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    "",
		})
		return
	}

	for i, _ := range RCList {
		var Rec_is_cancel int = 1
		if RCList[i].Queue_id != nil {
			var serviceUsed []structs.ServiceUsed
			if errCKS := models.CheckQueueReceiptServiceUsed(RCList[i].Id, *RCList[i].Queue_id, &serviceUsed); errCKS != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Receipt Check invalid.",
					"data":    errCKS.Error(),
				})
				return
			}
			if len(serviceUsed) > 0 && RCList[i].Rec_is_process == 1 {
				Rec_is_cancel = 0
			}
		} else {
			var serviceUsed []structs.ServiceUsed
			if errCKS := models.CheckReceiptServiceUsed(RCList[i].Id, &serviceUsed); errCKS != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Receipt Check invalid.",
					"data":    errCKS.Error(),
				})
				return
			}
			if len(serviceUsed) > 0 && RCList[i].Rec_is_process == 1 {
				Rec_is_cancel = 0
			}
		}
		var maxReceipt structs.MaxReceipt
		if errCMR := models.CheckMaxReceipt(RCList[i].Invoice_id, &maxReceipt); errCMR != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Check receipt max invalid.",
				"data":    errCMR.Error(),
			})
			return
		}
		if maxReceipt.Rec_period_max != RCList[i].Rec_period {
			Rec_is_cancel = 0
		}
		RCList[i].Rec_is_cancel = Rec_is_cancel
	}

	if len(RCList) == 0 {
		c.JSON(200, gin.H{
			"status":  true,
			"message": "",
			"data": structs.ResponsePaginationReceiptCustomer{
				Result_data:   []structs.ReceiptListCustomer{},
				Count_of_page: 0,
				Count_all:     0,
			},
		})
		return
	}

	res := structs.ResponsePaginationReceiptCustomer{
		Result_data:   RCList,
		Count_of_page: len(RCList),
		Count_all:     len(countList),
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    res,
	})
}

func SendSMSPDPA(c *gin.Context) {
	var objPayload structs.ObjPayloadSmsPDPA
	if errSBJ := c.ShouldBindJSON(&objPayload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errSBJ.Error(),
		})
		return
	}

	var objQueryCustomer structs.ObjQueryCustomer
	errCustomer := models.GetCustomerById(objPayload.CustomerId, &objQueryCustomer)
	if errCustomer != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Customer Invalid",
			"data":    "",
		})
		return
	}
	currentTime := time.Now()
	futureTime := currentTime.Add(24 * time.Hour)
	expiredTime := futureTime.Format("2006-01-02 15:04:05")

	// Encrypt the data
	Token, err := libs.Encrypt(expiredTime)
	if err != nil {
		fmt.Println("Error encrypting:", err)
		return
	}

	s := libs.Sms{
		Msisdn:  objPayload.CtmTel,
		Message: "Plese Accept PDPA by link " + os.Getenv("BASE_URL") + "/customer/pdpa/" + Token,
	}
	res, _ := s.Send()
	if res.StatusCode == 201 {
		err := models.UpdateTokenPDPA(objPayload.CustomerId, Token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "Error Create Token",
				"code":    err.Error(),
			})
		}

		models.AddLogCustomer(&structs.LogCustomer{
			Username:   c.Params.ByName("userEmail"),
			Log_type:   "Send PDPA Customer",
			Log_text:   "Send PDPA Customer Id = " + strconv.Itoa(objPayload.CustomerId) + " in token = " + Token,
			Log_create: time.Now().Format("2006-01-02 15:04:05"),
		})

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Send Link Accept PDPA to Number Phone " + objPayload.CtmTel,
			"data":    res.StatusCode,
			// "link":    "Plese Accept PDPA by link " + os.Getenv("BASE_URL") + "/customer/pdpa/" + Token,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "SMS is error.",
			"code":    res,
		})
	}

}

func UpdateSmsCustomerPDPA(c *gin.Context) {
	token := c.Params.ByName("token")

	// Decrypt the ciphertext

	decrypted, err := libs.Decrypt(token)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}
	layout := "2006-01-02 15:04:05" // Assuming the date string follows this layout
	datetimeStr := decrypted        // Example date string to calculate the difference from
	datetime, err := time.Parse(layout, datetimeStr)
	if err != nil {
		fmt.Println("Error parsing datetime:", err)
		return
	}
	now := time.Now()
	duration := datetime.Sub(now)
	minutes := int(duration.Minutes())

	if minutes < 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"expired": true,
			"message": "Token Expired. Please contact Administrator.",
			"data":    minutes,
		})
	}

	var objQueryCustomer structs.ObjQueryCustomer
	errCustomer := models.GetCustomerByToken(token, &objQueryCustomer)
	if errCustomer != nil || objQueryCustomer.ID == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"expired": true,
			"message": "Data Invalid",
			"data":    "",
		})
		return
	}

	if err := models.AcceptPDPA(objQueryCustomer.ID, objQueryCustomer.CtmSubscribePdpaImage); err != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"expired": false,
			"message": "Cannot Accept PDPA",
			"data":    err.Error(),
		})
		return
	}

	models.AddLogCustomer(&structs.LogCustomer{
		Username:   c.Params.ByName("userEmail"),
		Log_type:   "Accept PDPA Customer",
		Log_text:   "Accept PDPA Customer Id = " + strconv.Itoa(objQueryCustomer.ID),
		Log_create: time.Now().Format("2006-01-02 15:04:05"),
	})

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"expired": false,
		"message": "Accept PDPA Successful.",
		"data":    "",
	})
}

func UpdateCustomerPDPA(c *gin.Context) {

	var objPayload structs.ObjPayloadAcceptPDPA
	if errSBJ := c.ShouldBindJSON(&objPayload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errSBJ.Error(),
		})
		return
	}
	var objQueryCustomer structs.ObjQueryCustomer
	errCustomer := models.GetCustomerById(objPayload.CustomerId, &objQueryCustomer)
	if errCustomer != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Customer Invalid",
			"data":    "",
		})
		return
	}

	CtmImage := libs.UploadImageS3(objPayload.CustomerImage, "", "pdpa")

	if err := models.AcceptPDPA(objPayload.CustomerId, CtmImage); err != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Cannot Accept PDPA",
			"data":    err.Error(),
		})
		return
	}

	models.AddLogCustomer(&structs.LogCustomer{
		Username:   c.Params.ByName("userEmail"),
		Log_type:   "Accept PDPA Customer",
		Log_text:   "Accept PDPA Customer Id = " + strconv.Itoa(objPayload.CustomerId),
		Log_create: time.Now().Format("2006-01-02 15:04:05"),
	})

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Accept PDPA Successful.",
		"data":    "",
	})
}

func SearchAppointmentByCustomer(c *gin.Context) { //use
	customerId := libs.StrToInt(c.Params.ByName("customerId"))
	var filter structs.PayloadSearchAppointmentByCustomer
	if errSBJ := c.ShouldBindJSON(&filter); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errSBJ.Error(),
		})
		return
	}

	// shopId = libs.StrToInt(c.Params.ByName("shopId"))

	var countApList []structs.AppointmentListHistory

	if filter.ActivePage < 1 {
		filter.ActivePage = 0
	} else {
		filter.ActivePage -= 1
	}

	err := models.AppointmentSearchByCustomer(filter, false, &countApList, customerId)
	emptySlice := []string{}
	if err != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": err.Error(),
			"data":    emptySlice,
		})
		return
	} else {
		var apList []structs.AppointmentListHistory
		if errMD := models.AppointmentSearchByCustomer(filter, true, &apList, customerId); errMD != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Something went wrong.",
				"data":    "",
			})
			return
		}

		if len(apList) == 0 {
			c.JSON(200, gin.H{
				"status":  true,
				"message": "",
				"data": models.ResponsePaginationEmpty{
					Result_data:   emptySlice,
					Count_of_page: 0,
					Count_all:     0,
				},
			})
			return
		}

		res := structs.ResponseAppointmentHistory{
			Result_data:   apList,
			Count_of_page: len(apList),
			Count_all:     len(countApList),
		}

		c.JSON(200, gin.H{
			"status":  true,
			"message": "",
			"data":    res,
		})
	}
}

func HistoryAppointment(c *gin.Context) {

	var filter structs.PayloadSearchAppointmentByCustomerHistory
	if errSBJ := c.ShouldBindJSON(&filter); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errSBJ.Error(),
		})
		return
	}

	// filter.Shop_id = libs.StrToInt(c.Params.ByName("shopId"))

	var countApList []structs.AppointmentListHistory

	if filter.ActivePage < 1 {
		filter.ActivePage = 0
	} else {
		filter.ActivePage -= 1
	}

	err := models.AppointmentSearchByCustomerHistory(filter, false, &countApList)
	emptySlice := []string{}
	if err != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": err.Error(),
			"data":    emptySlice,
		})
		return
	} else {
		var apList []structs.AppointmentListHistory
		if errMD := models.AppointmentSearchByCustomerHistory(filter, true, &apList); errMD != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Something went wrong.",
				"data":    "",
			})
			return
		}

		if len(apList) == 0 {
			c.JSON(200, gin.H{
				"status":  true,
				"message": "",
				"data": models.ResponsePaginationEmpty{
					Result_data:   emptySlice,
					Count_of_page: 0,
					Count_all:     0,
				},
			})
			return
		}

		res := structs.ResponseAppointmentHistory{
			Result_data:   apList,
			Count_of_page: len(apList),
			Count_all:     len(countApList),
		}

		c.JSON(200, gin.H{
			"status":  true,
			"message": "",
			"data":    res,
		})
	}
}

func HistoryLab(c *gin.Context) { //use
	customerId := libs.StrToInt(c.Params.ByName("customerId"))
	var objPayload structs.ObjPayloadGetCustomerCheckLabXaryPagination
	shopMId := libs.StrToInt(c.Params.ByName("shopId"))
	if libs.StrToInt(c.Params.ByName("shopMotherId")) > 0 {
		shopMId = libs.StrToInt(c.Params.ByName("shopMotherId"))
	}

	if errSBJ := c.ShouldBindJSON(&objPayload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	if objPayload.CurrentPage < 1 {
		objPayload.CurrentPage = 0
	} else {
		objPayload.CurrentPage -= 1
	}

	countAll, errCountAll := models.CountHistoryCustomerLabPagination(&objPayload, customerId, shopMId)
	if errCountAll != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Get data error.",
			"data":    0,
		})
		return
	}

	var objQueryCheck []structs.ObjQueryGetCustomerCheckPagination
	errQueryCheck := models.GetHistoryCustomerLabPagination(&objPayload, &objQueryCheck, customerId, shopMId)
	if errQueryCheck != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Get data error.",
			"data": structs.ObjResponseGetCustomerCheckPagination{
				Items:     []structs.ObjQueryGetCustomerCheckPagination{},
				CountPage: 0,
				CountAll:  0,
			},
		})
		return
	}
	if len(objQueryCheck) > 0 {
		for i, check := range objQueryCheck {
			var objQueryCustomerCheckShop structs.Shop
			errCustomerCheckShop := models.GetCustomerOpdShopByIds(check.ShopId, &objQueryCustomerCheckShop)
			if errCustomerCheckShop != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Get data error.",
					"data":    "",
				})
				return
			}
			objQueryCheck[i].ShopName = objQueryCustomerCheckShop.ShopName
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Get data successful.",
		"data": structs.ObjResponseGetCustomerCheckPagination{
			Items:     objQueryCheck,
			CountPage: len(objQueryCheck),
			CountAll:  countAll,
		},
	})

}

func HistoryXray(c *gin.Context) { //use
	customerId := libs.StrToInt(c.Params.ByName("customerId"))
	var objPayload structs.ObjPayloadGetCustomerCheckLabXaryPagination

	if errSBJ := c.ShouldBindJSON(&objPayload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}
	shopMId := libs.StrToInt(c.Params.ByName("shopId"))
	if libs.StrToInt(c.Params.ByName("shopMotherId")) > 0 {
		shopMId = libs.StrToInt(c.Params.ByName("shopMotherId"))
	}

	if objPayload.CurrentPage < 1 {
		objPayload.CurrentPage = 0
	} else {
		objPayload.CurrentPage -= 1
	}

	countAll, errCountAll := models.CountHistoryCustomerXrayPagination(&objPayload, customerId, shopMId)
	if errCountAll != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Get data error.",
			"data":    0,
		})
		return
	}

	var objQueryCheck []structs.ObjQueryGetCustomerCheckPagination
	errQueryCheck := models.GetHistoryCustomerXrayPagination(&objPayload, &objQueryCheck, customerId, shopMId)
	if errQueryCheck != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Get data error.",
			"data": structs.ObjResponseGetCustomerCheckPagination{
				Items:     []structs.ObjQueryGetCustomerCheckPagination{},
				CountPage: 0,
				CountAll:  0,
			},
		})
		return
	}
	if len(objQueryCheck) > 0 {
		for i, check := range objQueryCheck {
			var objQueryCustomerCheckShop structs.Shop
			errCustomerCheckShop := models.GetCustomerOpdShopByIds(check.ShopId, &objQueryCustomerCheckShop)
			if errCustomerCheckShop != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Get data error.",
					"data":    "",
				})
				return
			}
			objQueryCheck[i].ShopName = objQueryCustomerCheckShop.ShopName
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Get data successful.",
		"data": structs.ObjResponseGetCustomerCheckPagination{
			Items:     objQueryCheck,
			CountPage: len(objQueryCheck),
			CountAll:  countAll,
		},
	})

}

func HistoryDocument(c *gin.Context) { //use
	customerId := libs.StrToInt(c.Params.ByName("customerId"))
	var objPayload structs.ObjPayloadPaginationHistory

	if errSBJ := c.ShouldBindJSON(&objPayload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	shopMId := libs.StrToInt(c.Params.ByName("shopId"))
	if libs.StrToInt(c.Params.ByName("shopMotherId")) > 0 {
		shopMId = libs.StrToInt(c.Params.ByName("shopMotherId"))
	}

	if objPayload.CurrentPage < 1 {
		objPayload.CurrentPage = 0
	} else {
		objPayload.CurrentPage -= 1
	}

	countAll, errCountAll := models.CountHistoryDocument(&objPayload, customerId, shopMId)
	if errCountAll != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Get data error.",
			"data":    0,
		})
	}

	var objQuery []structs.ObjQueryGetMedicalCertPagination

	err := models.GetHistoryDocument(&objPayload, &objQuery, customerId, shopMId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Get data error.",
			"data": structs.ObjResponseGetMedicalCertPagination{
				Items:     []structs.ObjQueryGetMedicalCertPagination{},
				CountPage: 0,
				CountAll:  0,
			},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Get data successful.",
			"data": structs.ObjResponseGetMedicalCertPagination{
				Items:     objQuery,
				CountPage: len(objQuery),
				CountAll:  countAll,
			},
		})
	}

}
