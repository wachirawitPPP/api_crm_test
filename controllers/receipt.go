package controllers

import (
	"fmt"
	"linecrmapi/libs"
	"linecrmapi/middlewares"
	"linecrmapi/models"
	"linecrmapi/structs"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func AddReceipt(c *gin.Context) {

	var objPayloadAddReceipt structs.ObjPayloadAddReceiptOrder

	if errSBJ := c.ShouldBindJSON(&objPayloadAddReceipt); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errSBJ.Error(),
		})
		return
	}

	var objInvoice models.Invoice
	errQueryInvoice := models.GetInvoiceById(objPayloadAddReceipt.InvoiceId, &objInvoice)
	if errQueryInvoice != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Get invoice error.",
			"data":    "",
		})
		return
	}

	Shop_id := objInvoice.ShopId

	var AccountListId int = 0
	var AccountList models.AccountList
	if errAccountCode := models.GetAccountList(Shop_id, &AccountList); errAccountCode != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Account List Invalid!",
			"data":    errAccountCode.Error(),
		})
		return
	}
	if AccountList.ID == 0 {
		var AccountCode structs.AccountCode
		if errAccountCode := models.GetAccountCode(Shop_id, &AccountCode); errAccountCode != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Error Get Account Code Invalid!",
				"data":    errAccountCode.Error(),
			})
			return
		}

		if AccountCode.Id == 0 {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Error Get Account Code Invalid!",
				"data":    "",
			})
			return
		} else {
			objCreate := models.AccountList{
				ShopId:        Shop_id,
				AccountCodeId: AccountCode.Id,
				AclCode:       "P001",
				AclName:       "Paysolutions",
				AclTypeId:     4,
				AclIsDel:      0,
				AclCreate:     time.Now().Format("2006-01-02 15:04:05"),
				AclUpdate:     time.Now().Format("2006-01-02 15:04:05"),
			}
			errCreateAccountList := models.CreateAccountList(&objCreate)
			if errCreateAccountList != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Error Create AccountList!",
					"data":    errCreateAccountList.Error(),
				})
				return
			}
			AccountListId = AccountList.ID
		}
	} else {
		AccountListId = AccountList.ID
	}

	var UserType1 structs.UserType1
	if errUser := models.GetUserType1(Shop_id, &UserType1); errUser != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get User Invalid!",
			"data":    errUser.Error(),
		})
		return
	}
	userId := UserType1.Id
	User_email := UserType1.User_email
	User_fullname := UserType1.User_fullname

	objPayload := structs.ObjPayloadAddReceipt{
		InvoiceId:       objPayloadAddReceipt.InvoiceId,
		ShopId:          Shop_id,
		RecTypeId:       1,
		AccountListId:   &AccountListId,
		RecPaymentType:  4,
		RecPay:          objInvoice.InvTotal,
		RecDescription:  objInvoice.InvComment,
		RecAccount:      123456789,
		RecUserId:       userId,
		RecUserFullname: User_fullname,
		RecUserEmail:    User_email,
		RecPayDatetime:  objInvoice.InvDatetime,
		DpmId:           nil,
	}

	ok, _ := middlewares.CompareDateTime(objInvoice.InvDatetime, objPayload.RecPayDatetime, 2, false)

	// เปรียบเทียบวัน เวลา
	if ok {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": fmt.Sprintf("Invoice date %s  is more then current date.", objInvoice.InvDatetime),
			"data":    "",
		})
		return
	}

	if objInvoice.InvIsActive != 1 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invoice is active invalid. Please refresh.",
			"data":    "Please refresh.",
		})
		return
	}

	var objInvoiceDetail []models.InvoiceDetail
	errQueryInvoiceDetail := models.GetInvoiceDetailById(objInvoice.ID, &objInvoiceDetail)
	if errQueryInvoiceDetail != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Get invoice detail error.",
			"data":    "",
		})
		return
	}

	var objReceiptDocSetting structs.ObjQueryReceiptDocSetting
	errQueryReceiptDocSetting := models.GetReceiptDocSetting(objPayload.ShopId, &objReceiptDocSetting)
	if errQueryReceiptDocSetting != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Get doc setting error.",
			"data":    "",
		})
		return
	}

	var objReceiptHistory []models.Receipt
	errQueryReceiptHistory := models.GetReceiptHistory(objInvoice.ID, &objReceiptHistory)
	if errQueryReceiptHistory != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Get receipt history error.",
			"data":    "",
		})
		return
	}

	var objShop models.Shop
	errQueryShop := models.GetReceiptShopById(objInvoice.ShopId, &objShop)
	if errQueryShop != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Get shop error.",
			"data":    "",
		})
		return
	}

	var objCustomer models.Customer
	errQueryCustomer := models.GetReceiptCustomerById(objInvoice.CustomerId, &objCustomer)
	if errQueryCustomer != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Get customer error.",
			"data":    "",
		})
		return
	}

	CtmCoinAmount := objCustomer.CtmCoin
	CtmPointAmount := 0.00
	if objCustomer.CtmPoint > 0 && objShop.ShopPointUseRate > 0 {
		CtmPointAmount = float64(objCustomer.CtmPoint / objShop.ShopPointUseRate)
	}

	if objPayload.RecPaymentType == 5 {
		if CtmCoinAmount < objPayload.RecPay {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Coin amount error.",
				"data":    "",
			})
			return
		}

	}

	if objPayload.RecPaymentType == 6 {
		if CtmPointAmount < objPayload.RecPay {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Point amount error.",
				"data":    "",
			})
			return
		}
	}

	invIsGetPoint := 1
	invGetPointAmount := 0
	if objPayload.RecPaymentType == 6 {
		invIsGetPoint = 0
	}

	invPayTotal := objPayload.RecPay
	for _, item := range objReceiptHistory {
		invPayTotal = invPayTotal + item.RecPay
		if invIsGetPoint == 1 && item.RecPaymentType == 6 {
			invIsGetPoint = 0
		}
	}

	invIsActive := 1
	if invPayTotal >= objInvoice.InvTotal {
		invIsActive = 2
		if invIsGetPoint == 1 {
			sumPriceTotal := 0.0
			for _, item := range objInvoiceDetail {
				if objShop.ShopPointChecking == 1 && item.InvdTypeId == 1 {
					sumPriceTotal = sumPriceTotal + item.InvdTotal
				} else if objShop.ShopPointCourse == 1 && item.InvdTypeId == 2 {
					sumPriceTotal = sumPriceTotal + item.InvdTotal
				} else if objShop.ShopPointProduct == 1 && item.InvdTypeId == 3 {
					sumPriceTotal = sumPriceTotal + item.InvdTotal
				}
			}
			if sumPriceTotal > 0 && objShop.ShopPointGiveRate > 0 {
				invGetPointAmount = int(sumPriceTotal / float64(objShop.ShopPointGiveRate))
			}
		}
	}

	recCode := libs.SetDocSettingCode(objReceiptDocSetting.ReceiptIdDefault, objReceiptDocSetting.ReceiptNumberDigit, objReceiptDocSetting.ReceiptNumberDefault, objReceiptDocSetting.ReceiptType)
	recIsProcess := 1
	recPeriod := 1
	if objPayload.RecTypeId == 2 {
		recPeriod = 0
	}
	if len(objReceiptHistory) > 0 {
		objReceiptLast := objReceiptHistory[len(objReceiptHistory)-1]
		recIsProcess = 0
		recPeriod = objReceiptLast.RecPeriod + 1
	}
	recCodeNext := objReceiptDocSetting.ReceiptNumberDefault + 1

	//เช็คยาก่อน
	// if recIsProcess == 1 {
	// 	var Invd []structs.InvoiceDetailCheck
	// 	if errInvd := models.GetInvoiceDetailProductCheck(objPayload.InvoiceId, &Invd); errInvd != nil {
	// 		c.AbortWithStatusJSON(200, gin.H{
	// 			"status":  false,
	// 			"message": "Invoice detail product invalid.",
	// 			"data":    errInvd.Error(),
	// 		})
	// 		return
	// 	}
	// 	for _, Invdproduct := range Invd {
	// 		var productStoreOrder structs.ProcessProductStoreOrderCheck
	// 		if errInvdp := models.GetProcessProductStoreOrderCheck(*Invdproduct.Product_id, *Invdproduct.Product_store_id, &productStoreOrder); errInvdp != nil {
	// 			c.AbortWithStatusJSON(200, gin.H{
	// 				"status":  false,
	// 				"message": "Invoice detail product invalid.",
	// 				"data":    errInvdp.Error(),
	// 			})
	// 			return
	// 		}
	// 		if productStoreOrder.Id == 0 {
	// 			c.AbortWithStatusJSON(200, gin.H{
	// 				"status":  false,
	// 				"message": "Invoice detail product empty.",
	// 				"data":    Invdproduct.Invd_name + " ( " + Invdproduct.Invd_code + " ) ",
	// 			})
	// 			return
	// 		}
	// 		var InvoiceDetailCheckCourse structs.InvoiceDetailCheckCourse
	// 		if errInvdpC := models.GetInvoiceDetailProductCheckCourse(objPayload.InvoiceId, *Invdproduct.Product_store_id, &InvoiceDetailCheckCourse); errInvdpC != nil {
	// 			c.AbortWithStatusJSON(200, gin.H{
	// 				"status":  false,
	// 				"message": "Invoice detail product invalid.",
	// 				"data":    errInvdpC.Error(),
	// 			})
	// 			return
	// 		}
	// 		var Invd_qty float64 = Invdproduct.Invd_qty
	// 		if InvoiceDetailCheckCourse.Invd_qty != 0 {
	// 			Invd_qty = Invdproduct.Invd_qty - InvoiceDetailCheckCourse.Invd_qty
	// 		}
	// 		if productStoreOrder.Pdso_total < Invd_qty {
	// 			c.AbortWithStatusJSON(200, gin.H{
	// 				"status":  false,
	// 				"message": "Invoice detail product are not enough to pay.",
	// 				"data":    Invdproduct.Invd_name + " ( " + Invdproduct.Invd_code + " ) ",
	// 			})
	// 			return
	// 		}
	// 	}
	// }

	recId, errAddReceiptTx := models.CreateReceiptTx(&objInvoice, &objInvoiceDetail, &objPayload, &objShop, &objCustomer, recCode, recIsProcess, recPeriod, recCodeNext, invIsActive, invPayTotal, invIsGetPoint, invGetPointAmount)
	if errAddReceiptTx != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Create receipt error.",
			"data":    errAddReceiptTx.Error(),
		})
		return
	}

	if recId == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Create receipt error. (recId = 0)",
			"data":    recId,
		})
		return
	} else {
		models.AddLogReceipt(&structs.LogReceipt{
			Username:   c.Params.ByName("userEmail"),
			Log_type:   "Add Receipt",
			Log_text:   "Add Receipt Id = " + recCode + " (" + strconv.Itoa(recId) + ")",
			Log_create: time.Now().Format("2006-01-02 15:04:05"),
			Shop_id:    libs.StrToInt(c.Params.ByName("shopId")),
		})
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Create receipt success.",
		"data": map[string]interface{}{
			"rec_id":         recId,
			"rec_is_process": recIsProcess,
		},
	})

}

// func ReceiptsSearch(c *gin.Context) {
// 	var filter structs.ObjPayloadSearchReceipt
// 	if errSBJ := c.ShouldBindJSON(&filter); errSBJ != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Invalid request data.",
// 			"data":    errSBJ.Error(),
// 		})
// 		return
// 	}

// 	filter.Shop_id = libs.StrToInt(c.Params.ByName("shopId"))
// 	var countList []structs.ReceiptList
// 	if errMD := models.GetReceiptList(filter, false, &countList); errMD != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Something went wrong.",
// 			"data":    "",
// 		})
// 		return
// 	}

// 	if filter.ActivePage < 1 {
// 		filter.ActivePage = 0
// 	} else {
// 		filter.ActivePage -= 1
// 	}

// 	var RCList []structs.ReceiptList
// 	if errMD := models.GetReceiptList(filter, true, &RCList); errMD != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Something went wrong.",
// 			"data":    "",
// 		})
// 		return
// 	}

// 	for i, _ := range RCList {
// 		var Rec_is_cancel int = 1
// 		if RCList[i].Queue_id != nil {
// 			var serviceUsed []structs.ServiceUsed
// 			if errCKS := models.CheckQueueReceiptServiceUsed(RCList[i].Id, *RCList[i].Queue_id, &serviceUsed); errCKS != nil {
// 				c.AbortWithStatusJSON(200, gin.H{
// 					"status":  false,
// 					"message": "Receipt Check invalid.",
// 					"data":    errCKS.Error(),
// 				})
// 				return
// 			}
// 			if len(serviceUsed) > 0 && RCList[i].Rec_is_process == 1 {
// 				Rec_is_cancel = 0
// 			}
// 		} else {
// 			var serviceUsed []structs.ServiceUsed
// 			if errCKS := models.CheckReceiptServiceUsed(RCList[i].Id, &serviceUsed); errCKS != nil {
// 				c.AbortWithStatusJSON(200, gin.H{
// 					"status":  false,
// 					"message": "Receipt Check invalid.",
// 					"data":    errCKS.Error(),
// 				})
// 				return
// 			}
// 			if len(serviceUsed) > 0 && RCList[i].Rec_is_process == 1 {
// 				Rec_is_cancel = 0
// 			}
// 		}
// 		var maxReceipt structs.MaxReceipt
// 		if errCMR := models.CheckMaxReceipt(RCList[i].Invoice_id, &maxReceipt); errCMR != nil {
// 			c.AbortWithStatusJSON(200, gin.H{
// 				"status":  false,
// 				"message": "Check receipt max invalid.",
// 				"data":    errCMR.Error(),
// 			})
// 			return
// 		}
// 		if maxReceipt.Rec_period_max != RCList[i].Rec_period {
// 			Rec_is_cancel = 0
// 		}
// 		RCList[i].Rec_is_cancel = Rec_is_cancel
// 	}

// 	if len(RCList) == 0 {
// 		c.JSON(200, gin.H{
// 			"status":  true,
// 			"message": "",
// 			"data": structs.ResponsePaginationReceipt{
// 				Result_data:   []structs.ReceiptList{},
// 				Count_of_page: 0,
// 				Count_all:     0,
// 			},
// 		})
// 		return
// 	}

// 	for i := range RCList {
// 		if RCList[i].User_id_cancel != 0 {
// 			var Users structs.UserCancel
// 			if errUsers := models.GetUserCancel(RCList[i].User_id_cancel, &Users); errUsers != nil {
// 				RCList[i].User_fullname_cancel = ""
// 				RCList[i].User_fullname_en_cancel = ""
// 			} else {
// 				RCList[i].User_fullname_cancel = Users.User_fullname
// 				RCList[i].User_fullname_en_cancel = Users.User_fullname_en
// 			}
// 		}
// 	}

// 	res := structs.ResponsePaginationReceipt{
// 		Result_data:   RCList,
// 		Count_of_page: len(RCList),
// 		Count_all:     len(countList),
// 	}

// 	c.JSON(200, gin.H{
// 		"status":  true,
// 		"message": "",
// 		"data":    res,
// 	})
// }

func ReceiptDetail(c *gin.Context) {
	RCId := libs.StrToInt(c.Params.ByName("id"))
	shopId := libs.StrToInt(c.Params.ByName("shopId"))
	customerId := libs.StrToInt(c.Params.ByName("customerId"))
	var RCData structs.ReceiptPrint
	err := models.GetReceipById(RCId, shopId, &RCData, customerId)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Invoice Invalid!",
			"data":    err.Error(),
		})
		return
	} else if RCData.Id == 0 {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Invoice Invalid!",
			"data":    "",
		})
		return
	}

	// 1 cash, 2 bank, 3 card, 4 pay solutions, 5 coin, 6 point
	if RCData.Rec_payment_type == 1 {
		RCData.Rec_payment_type_th = "เงินสด"
		RCData.Rec_payment_type_en = "Cash"
	}
	if RCData.Rec_payment_type == 2 {
		RCData.Rec_payment_type_th = "ธนาคาร " + RCData.Acl_name
		RCData.Rec_payment_type_en = "Bank " + RCData.Acl_name
	}
	if RCData.Rec_payment_type == 3 {
		RCData.Rec_payment_type_th = "บัตรเครดิต/เดบิต " + RCData.Acl_name
		RCData.Rec_payment_type_en = "Card " + RCData.Acl_name
	}
	if RCData.Rec_payment_type == 4 {
		RCData.Rec_payment_type_th = "Pay Solutions " + RCData.Acl_name
		RCData.Rec_payment_type_en = "Pay Solutions " + RCData.Acl_name
	}
	if RCData.Rec_payment_type == 5 {
		RCData.Rec_payment_type_th = "วงเงิน " + RCData.Acl_name
		RCData.Rec_payment_type_en = "Coin " + RCData.Acl_name
	}
	if RCData.Rec_payment_type == 6 {
		RCData.Rec_payment_type_th = "แต้ม " + RCData.Acl_name
		RCData.Rec_payment_type_en = "Point " + RCData.Acl_name
	}

	var RCShop structs.ReceiptShop
	if errShop := models.GetShopReceiptById(RCData.Shop_id, &RCShop); errShop != nil || RCShop.Id == 0 {
		RCData.Shop = structs.ReceiptShop{}
	} else {
		RCData.Shop = RCShop
	}

	var RCCus structs.ObjQueryCustomer
	if errCus := models.GetCustomerById(RCData.Customer_id, &RCCus); errCus != nil || RCCus.ID == 0 {
		RCData.Customer = structs.ObjQueryCustomer{}
	} else {
		RCData.Customer = RCCus
	}

	var RCSubs []structs.ReceiptDetail
	if errINVS := models.GetReceiptDetail(RCId, &RCSubs); errINVS != nil || len(RCSubs) == 0 {
		RCData.Subs = &[]structs.ReceiptDetail{}
	} else {
		RCData.Subs = &RCSubs
	}

	if RCData.User_id_cancel != 0 {
		var Users structs.UserCancel
		if errUsers := models.GetUserCancel(RCData.User_id_cancel, &Users); errUsers != nil {
			RCData.User_fullname_cancel = ""
			RCData.User_fullname_en_cancel = ""
		} else {
			RCData.User_fullname_cancel = Users.User_fullname
			RCData.User_fullname_en_cancel = Users.User_fullname_en
		}
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    RCData,
	})

}

// func DelReceipt(c *gin.Context) {
// 	RECId := libs.StrToInt(c.Params.ByName("id"))
// 	var REC structs.Receipt
// 	if errCK := models.GetReceiptId(RECId, &REC); errCK != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Receipt invalid.",
// 			"data":    errCK.Error(),
// 		})
// 		return
// 	}

// 	if REC.Id < 1 {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Receipt invalid.",
// 			"data":    REC,
// 		})
// 		return
// 	}

// 	if REC.Rec_is_active != 1 {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Receipt Status invalid.",
// 			"data":    REC,
// 		})
// 		return
// 	}

// 	var objInvoice models.Invoice
// 	errQueryInvoice := models.GetInvoiceById(REC.Invoice_id, &objInvoice)
// 	if errQueryInvoice != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Get invoice error.",
// 			"data":    "",
// 		})
// 		return
// 	}

// 	if REC.Queue_id != nil {
// 		var serviceUsed []structs.ServiceUsed
// 		if errCKS := models.CheckQueueReceiptServiceUsed(RECId, *REC.Queue_id, &serviceUsed); errCKS != nil {
// 			c.AbortWithStatusJSON(200, gin.H{
// 				"status":  false,
// 				"message": "Receipt Check invalid.",
// 				"data":    errCKS.Error(),
// 			})
// 			return
// 		}

// 		if len(serviceUsed) > 0 && REC.Rec_is_process == 1 {
// 			c.AbortWithStatusJSON(200, gin.H{
// 				"status":  false,
// 				"message": "Check Queue Service Used.",
// 				"data":    "",
// 			})
// 			return
// 		}
// 		if REC.Rec_is_process == 1 {
// 			if errAD := models.UpdateReceiptQueueCancal(*REC.Queue_id); errAD != nil {
// 				c.AbortWithStatusJSON(200, gin.H{
// 					"status":  false,
// 					"message": "Cannot Update Queue",
// 					"data":    errAD.Error(),
// 				})
// 				return
// 			}
// 		}

// 	} else {
// 		var serviceUsed []structs.ServiceUsed
// 		if errCKS := models.CheckReceiptServiceUsed(RECId, &serviceUsed); errCKS != nil {
// 			c.AbortWithStatusJSON(200, gin.H{
// 				"status":  false,
// 				"message": "Receipt Check invalid.",
// 				"data":    errCKS.Error(),
// 			})
// 			return
// 		}

// 		if len(serviceUsed) > 0 && REC.Rec_is_process == 1 {
// 			c.AbortWithStatusJSON(200, gin.H{
// 				"status":  false,
// 				"message": "Check Queue Service Used.",
// 				"data":    "",
// 			})
// 			return
// 		}
// 	}
// 	var objCustomer models.Customer
// 	errCustomer := models.GetReceiptCustomerById(REC.Customer_id, &objCustomer)
// 	if errCustomer != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Get customer data error.",
// 			"data":    "",
// 		})
// 		return
// 	}

// 	if REC.Rec_point_give > 0 {
// 		//ลบแต้มออกจากลูกค้า
// 		pointCustomerUpdate := structs.PointCustomerUpdate{
// 			Ctm_point:  objCustomer.CtmPoint - REC.Rec_point_give,
// 			Ctm_update: time.Now().Format("2006-01-02 15:04:05"),
// 		}
// 		if errpcu := models.UpdatePointCustomerId(REC.Customer_id, &pointCustomerUpdate); errpcu != nil {
// 			c.AbortWithStatusJSON(200, gin.H{
// 				"status":  false,
// 				"message": "Update Services error.",
// 				"data":    errpcu.Error(),
// 			})
// 			return
// 		}
// 		// Add point history
// 		AddPointHistory := structs.AddPointHistory{
// 			Shop_id:     REC.Shop_id,
// 			Customer_id: REC.Customer_id,
// 			Receipt_id:  RECId,
// 			Rec_code:    REC.Rec_code,
// 			Ph_forward:  float64(objCustomer.CtmPoint),
// 			Ph_amount:   math.Ceil(float64(REC.Rec_point_give) * (-1)),
// 			Ph_total:    float64(objCustomer.CtmPoint) - math.Ceil(float64(REC.Rec_point_give)),
// 			Ph_comment:  "Cancel Receipt: Minus Point History",
// 			Ph_create:   time.Now().Format("2006-01-02 15:04:05"),
// 		}
// 		if errPointHistory := models.CreatePointHistory(&AddPointHistory); errPointHistory != nil {
// 			c.AbortWithStatusJSON(200, gin.H{
// 				"status":  false,
// 				"message": "Cancel point history error.",
// 				"data":    errPointHistory.Error(),
// 			})
// 			return
// 		}
// 	}

// 	if REC.Rec_point_used > 0 {
// 		//บวกแต้มคืนลูกค้า
// 		pointCustomerUpdate := structs.PointCustomerUpdate{
// 			Ctm_point:  objCustomer.CtmPoint + REC.Rec_point_used,
// 			Ctm_update: time.Now().Format("2006-01-02 15:04:05"),
// 		}
// 		if errpcu := models.UpdatePointCustomerId(REC.Customer_id, &pointCustomerUpdate); errpcu != nil {
// 			c.AbortWithStatusJSON(200, gin.H{
// 				"status":  false,
// 				"message": "Update Services error.",
// 				"data":    errpcu.Error(),
// 			})
// 			return
// 		}
// 		// Add point history
// 		AddPointHistory := structs.AddPointHistory{
// 			Shop_id:     REC.Shop_id,
// 			Customer_id: REC.Customer_id,
// 			Receipt_id:  RECId,
// 			Rec_code:    REC.Rec_code,
// 			Ph_forward:  float64(objCustomer.CtmPoint),
// 			Ph_amount:   math.Ceil(float64(REC.Rec_point_used)),
// 			Ph_total:    float64(objCustomer.CtmPoint) + math.Ceil(float64(REC.Rec_point_used)),
// 			Ph_comment:  "Cancel Receipt: Add Point History",
// 			Ph_create:   time.Now().Format("2006-01-02 15:04:05"),
// 		}
// 		if errPointHistory := models.CreatePointHistory(&AddPointHistory); errPointHistory != nil {
// 			c.AbortWithStatusJSON(200, gin.H{
// 				"status":  false,
// 				"message": "Cancel point history error.",
// 				"data":    errPointHistory.Error(),
// 			})
// 			return
// 		}
// 	}

// 	if REC.Rec_payment_type == 5 {
// 		//คืนวงเงิน
// 		coinCustomerUpdate := structs.CoinCustomerUpdate{
// 			Ctm_coin:   objCustomer.CtmCoin + REC.Rec_pay,
// 			Ctm_update: time.Now().Format("2006-01-02 15:04:05"),
// 		}
// 		if errpcu := models.UpdateCoinCustomerId(REC.Customer_id, &coinCustomerUpdate); errpcu != nil {
// 			c.AbortWithStatusJSON(200, gin.H{
// 				"status":  false,
// 				"message": "Update Customer Coin error.",
// 				"data":    errpcu.Error(),
// 			})
// 			return
// 		}
// 		// Add coin history
// 		AddCoinHistory := structs.AddCoinHistory{
// 			Shop_id:     REC.Shop_id,
// 			Customer_id: REC.Customer_id,
// 			Receipt_id:  RECId,
// 			Rec_code:    REC.Rec_code,
// 			Ch_forward:  float64(objCustomer.CtmCoin),
// 			Ch_amount:   float64(REC.Rec_pay),
// 			Ch_total:    float64(objCustomer.CtmCoin) + float64(REC.Rec_pay),
// 			Ch_comment:  "Cancel Receipt: Add Coin History",
// 			Ch_create:   time.Now().Format("2006-01-02 15:04:05"),
// 		}
// 		if errCoinHistory := models.CreateCoinHistory(&AddCoinHistory); errCoinHistory != nil {
// 			c.AbortWithStatusJSON(200, gin.H{
// 				"status":  false,
// 				"message": "Cancel coin history error.",
// 				"data":    errCoinHistory.Error(),
// 			})
// 			return
// 		}
// 	}
// 	userId := libs.StrToInt(c.Params.ByName("userID"))
// 	if errAD := models.CancelReceipt(RECId, userId, &REC); errAD != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Cannot Remove Receipt",
// 			"data":    errAD.Error(),
// 		})
// 		return
// 	}
// 	var InvDeposit float64 = objInvoice.InvDeposit
// 	if REC.Rec_type_id == 2 {
// 		InvDeposit = 0
// 	}
// 	cancelInvoiceUpdate := structs.CancelInvoiceUpdate{
// 		Inv_is_active: 1,
// 		Inv_time_end:  0,
// 		Inv_pay_total: objInvoice.InvPayTotal - REC.Rec_pay,
// 		Inv_deposit:   InvDeposit,
// 		Inv_update:    time.Now().Format("2006-01-02 15:04:05"),
// 	}
// 	if errpcoinupdate := models.UpdateReceiptInvoiceCancal(REC.Invoice_id, &cancelInvoiceUpdate); errpcoinupdate != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Update Services error.",
// 			"data":    errpcoinupdate.Error(),
// 		})
// 		return
// 	}

// 	models.AddLogReceipt(&structs.LogReceipt{
// 		Username:   c.Params.ByName("userEmail"),
// 		Log_type:   "Cancel Receipt",
// 		Log_text:   "Cancel Receipt Id = " + REC.Rec_code + " (" + strconv.Itoa(RECId) + ")",
// 		Log_create: time.Now().Format("2006-01-02 15:04:05"),
// 	})

// 	c.JSON(200, gin.H{
// 		"status":  true,
// 		"message": "Cancel Receipt ID " + REC.Rec_code + " success",
// 		"data": map[string]interface{}{
// 			"rec_id":         REC.Id,
// 			"rec_is_process": REC.Rec_is_process,
// 		},
// 	})
// }

// // tom code
// func GetReceiptShop(c *gin.Context) {

// 	shopId := libs.StrToInt(c.Params.ByName("shopId"))

// 	var objResponse models.Shop

// 	err := models.GetReceiptShopById(shopId, &objResponse)
// 	if err != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Get data error.",
// 			"data":    "",
// 		})
// 		return
// 	} else {
// 		c.JSON(200, gin.H{
// 			"status":  true,
// 			"message": "Get data successful.",
// 			"data":    objResponse,
// 		})
// 	}
// }

// func GetReceiptCustomer(c *gin.Context) {

// 	ctmId := libs.StrToInt(c.Params.ByName("ctmId"))

// 	var objResponse models.Customer

// 	err := models.GetReceiptCustomerById(ctmId, &objResponse)
// 	if err != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Get data error.",
// 			"data":    "",
// 		})
// 		return
// 	} else {
// 		c.JSON(200, gin.H{
// 			"status":  true,
// 			"message": "Get data successful.",
// 			"data":    objResponse,
// 		})
// 	}
// }

// func GetReceiptHistory(c *gin.Context) {

// 	invId := libs.StrToInt(c.Params.ByName("invId"))

// 	var objResponse []models.Receipt

// 	err := models.GetReceiptHistory(invId, &objResponse)
// 	if err != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Get data error.",
// 			"data":    "",
// 		})
// 		return
// 	} else {
// 		c.JSON(200, gin.H{
// 			"status":  true,
// 			"message": "Get data successful.",
// 			"data":    objResponse,
// 		})
// 	}
// }

// func AddReceipt(c *gin.Context) {

// 	var objPayload structs.ObjPayloadAddReceipt

// 	if errSBJ := c.ShouldBindJSON(&objPayload); errSBJ != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Invalid request data.",
// 			"data":    errSBJ.Error(),
// 		})
// 		return
// 	}

// 	var objInvoice models.Invoice
// 	errQueryInvoice := models.GetInvoiceById(objPayload.InvoiceId, &objInvoice)
// 	if errQueryInvoice != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Get invoice error.",
// 			"data":    "",
// 		})
// 		return
// 	}

// 	if objInvoice.InvIsActive != 1 {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Invoice is active invalid. Please refresh.",
// 			"data":    "Please refresh.",
// 		})
// 		return
// 	}

// 	var objInvoiceDetail []models.InvoiceDetail
// 	errQueryInvoiceDetail := models.GetInvoiceDetailById(objInvoice.ID, &objInvoiceDetail)
// 	if errQueryInvoiceDetail != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Get invoice detail error.",
// 			"data":    "",
// 		})
// 		return
// 	}

// 	var objReceiptDocSetting structs.ObjQueryReceiptDocSetting
// 	errQueryReceiptDocSetting := models.GetReceiptDocSetting(objPayload.ShopId, &objReceiptDocSetting)
// 	if errQueryReceiptDocSetting != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Get doc setting error.",
// 			"data":    "",
// 		})
// 		return
// 	}

// 	var objReceiptHistory []models.Receipt
// 	errQueryReceiptHistory := models.GetReceiptHistory(objInvoice.ID, &objReceiptHistory)
// 	if errQueryReceiptHistory != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Get receipt history error.",
// 			"data":    "",
// 		})
// 		return
// 	}

// 	var objShop models.Shop
// 	errQueryShop := models.GetReceiptShopById(objInvoice.ShopId, &objShop)
// 	if errQueryShop != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Get shop error.",
// 			"data":    "",
// 		})
// 		return
// 	}

// 	var objCustomer models.Customer
// 	errQueryCustomer := models.GetReceiptCustomerById(objInvoice.CustomerId, &objCustomer)
// 	if errQueryCustomer != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Get customer error.",
// 			"data":    "",
// 		})
// 		return
// 	}

// 	CtmCoinAmount := objCustomer.CtmCoin
// 	CtmPointAmount := 0.00
// 	if objCustomer.CtmPoint > 0 && objShop.ShopPointUseRate > 0 {
// 		CtmPointAmount = float64(objCustomer.CtmPoint / objShop.ShopPointUseRate)
// 	}

// 	if objPayload.RecPaymentType == 5 {
// 		if CtmCoinAmount < objPayload.RecPay {
// 			c.AbortWithStatusJSON(200, gin.H{
// 				"status":  false,
// 				"message": "Coin amount error.",
// 				"data":    "",
// 			})
// 			return
// 		}

// 	}

// 	if objPayload.RecPaymentType == 6 {
// 		if CtmPointAmount < objPayload.RecPay {
// 			c.AbortWithStatusJSON(200, gin.H{
// 				"status":  false,
// 				"message": "Point amount error.",
// 				"data":    "",
// 			})
// 			return
// 		}
// 	}

// 	invIsGetPoint := 1
// 	invGetPointAmount := 0
// 	if objPayload.RecPaymentType == 6 {
// 		invIsGetPoint = 0
// 	}

// 	invPayTotal := objPayload.RecPay
// 	for _, item := range objReceiptHistory {
// 		invPayTotal = invPayTotal + item.RecPay
// 		if invIsGetPoint == 1 && item.RecPaymentType == 6 {
// 			invIsGetPoint = 0
// 		}
// 	}

// 	invIsActive := 1
// 	if invPayTotal >= objInvoice.InvTotal {
// 		invIsActive = 2
// 		if invIsGetPoint == 1 {
// 			sumPriceTotal := 0.0
// 			for _, item := range objInvoiceDetail {
// 				if objShop.ShopPointChecking == 1 && item.InvdTypeId == 1 {
// 					sumPriceTotal = sumPriceTotal + item.InvdTotal
// 				} else if objShop.ShopPointCourse == 1 && item.InvdTypeId == 2 {
// 					sumPriceTotal = sumPriceTotal + item.InvdTotal
// 				} else if objShop.ShopPointProduct == 1 && item.InvdTypeId == 3 {
// 					sumPriceTotal = sumPriceTotal + item.InvdTotal
// 				}
// 			}
// 			if sumPriceTotal > 0 && objShop.ShopPointGiveRate > 0 {
// 				invGetPointAmount = int(sumPriceTotal / float64(objShop.ShopPointGiveRate))
// 			}
// 		}
// 	}

// 	recCode := libs.SetDocSettingCode(objReceiptDocSetting.ReceiptIdDefault, objReceiptDocSetting.ReceiptNumberDigit, objReceiptDocSetting.ReceiptNumberDefault, objReceiptDocSetting.ReceiptType)
// 	recIsProcess := 1
// 	recPeriod := 1
// 	if objPayload.RecTypeId == 2 {
// 		recPeriod = 0
// 	}
// 	if len(objReceiptHistory) > 0 {
// 		objReceiptLast := objReceiptHistory[len(objReceiptHistory)-1]
// 		recIsProcess = 0
// 		recPeriod = objReceiptLast.RecPeriod + 1
// 	}

// 	recCodeNext := objReceiptDocSetting.ReceiptNumberDefault + 1

// 	//เช็คยาก่อน
// 	if recIsProcess == 1 {
// 		var Invd []structs.InvoiceDetail
// 		if errInvd := models.GetInvoiceDetailProduct(objPayload.InvoiceId, &Invd); errInvd != nil {
// 			c.AbortWithStatusJSON(200, gin.H{
// 				"status":  false,
// 				"message": "Invoice detail product invalid.",
// 				"data":    errInvd.Error(),
// 			})
// 			return
// 		}
// 		for _, Invdproduct := range Invd {
// 			var productStoreOrder structs.ProcessProductStoreOrder
// 			if errInvdp := models.GetProcessProductStoreOrderLast(*Invdproduct.Product_id, *Invdproduct.Product_store_id, &productStoreOrder); errInvdp != nil {
// 				c.AbortWithStatusJSON(200, gin.H{
// 					"status":  false,
// 					"message": "Invoice detail product invalid.",
// 					"data":    errInvdp.Error(),
// 				})
// 				return
// 			}
// 			if productStoreOrder.Id == 0 {
// 				c.AbortWithStatusJSON(200, gin.H{
// 					"status":  false,
// 					"message": "Invoice detail product empty.",
// 					"data":    Invdproduct.Invd_name + " ( " + Invdproduct.Invd_code + " ) ",
// 				})
// 				return
// 			}
// 		}
// 	}

// 	recId, errAddReceiptTx := models.CreateReceiptTx(&objInvoice, &objInvoiceDetail, &objPayload, &objShop, &objCustomer, recCode, recIsProcess, recPeriod, recCodeNext, invIsActive, invPayTotal, invIsGetPoint, invGetPointAmount)
// 	if errAddReceiptTx != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Create receipt error.",
// 			"data":    errAddReceiptTx.Error(),
// 		})
// 		return
// 	}

// 	if recId == 0 {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Create receipt error. (recId = 0)",
// 			"data":    recId,
// 		})
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"status":  true,
// 		"message": "Create receipt success.",
// 		"data": map[string]interface{}{
// 			"rec_id":         recId,
// 			"rec_is_process": recIsProcess,
// 		},
// 	})

// }

// func ReceiptFile(c *gin.Context) {

// 	var objPayload structs.FileReceipt

// 	if errSBJ := c.ShouldBindJSON(&objPayload); errSBJ != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Invalid request data.",
// 			"data":    errSBJ.Error(),
// 		})
// 		return
// 	}

// 	var receipt structs.ReceiptFileCheck
// 	if errCK := models.GetReceiptFileCheckBYId(objPayload.Rec_id, &receipt); errCK != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Receipt invalid.",
// 			"data":    errCK.Error(),
// 		})
// 		return
// 	}

// 	if receipt.Id == 0 {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Receipt invalid.",
// 			"data":    "",
// 		})
// 		return
// 	}

// 	RecFile := ""
// 	rec_file_size := 0
// 	if *objPayload.Rec_file != "" {
// 		RecFile = receipt.Rec_file
// 	}

// 	if *objPayload.Rec_file != "" {
// 		filename := receipt.Shop_code + "_" + receipt.Rec_code + "_" + receipt.Ctm_id + "_" + time.Now().Format("20060102150405")
// 		RecFile, rec_file_size = libs.UploadFileFilenameS3(*objPayload.Rec_file, "receipt", filename)
// 	}

// 	// if RecFile == "" {
// 	// 	c.AbortWithStatusJSON(200, gin.H{
// 	// 		"status":  false,
// 	// 		"message": "Upload file receipt empty.",
// 	// 		"data":    RecFile,
// 	// 	})
// 	// 	return
// 	// }

// 	// UpdateFileReceipt
// 	var UR structs.Receipt
// 	if errUF := models.UpdateFileReceipt(objPayload.Rec_id, rec_file_size, RecFile, &UR); errUF != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Upload file receipt error.",
// 			"data":    errUF.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"status":  true,
// 		"message": "Upload file receipt success.",
// 		"data":    "",
// 	})

// }
