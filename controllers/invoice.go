package controllers

import (
	"fmt"
	"linecrmapi/libs"
	"linecrmapi/models"
	"linecrmapi/structs"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func InvoicesSearch(c *gin.Context) {
	var filter structs.ObjPayloadSearchInvoice
	if errSBJ := c.ShouldBindJSON(&filter); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errSBJ.Error(),
		})
		return
	}

	filter.Shop_id = libs.StrToInt(c.Params.ByName("shopId"))
	var countList []structs.InvoiceList
	if filter.ActivePage < 1 {
		filter.ActivePage = 0
	} else {
		filter.ActivePage -= 1
	}

	err := models.GetInvoiceList(filter, false, &countList)
	emptySlice := []string{}
	if err != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": err.Error(),
			"data":    emptySlice,
		})
		c.Abort()
	} else {
		var INVList []structs.InvoiceList
		if errMD := models.GetInvoiceList(filter, true, &INVList); errMD != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Something went wrong.",
				"data":    "",
			})
			return
		}

		if len(INVList) == 0 {
			emptyINVList := []structs.InvoiceList{}
			c.JSON(200, gin.H{
				"status":  true,
				"message": "",
				"data": structs.ResponsePaginationInvoice{
					Result_data:   emptyINVList,
					Count_of_page: 0,
					Count_all:     0,
				},
			})
			return
		}

		for i := range INVList {
			if INVList[i].User_id_cancel != 0 {
				var Users structs.UserCancel
				if errUsers := models.GetUserCancel(INVList[i].User_id_cancel, &Users); errUsers != nil {
					INVList[i].User_fullname_cancel = ""
					INVList[i].User_fullname_en_cancel = ""
				} else {
					INVList[i].User_fullname_cancel = Users.User_fullname
					INVList[i].User_fullname_en_cancel = Users.User_fullname_en
				}
			}
		}

		res := structs.ResponsePaginationInvoice{
			Result_data:   INVList,
			Count_of_page: len(INVList),
			Count_all:     len(countList),
		}

		c.JSON(200, gin.H{
			"status":  true,
			"message": "",
			"data":    res,
		})
	}
}

func AddInvoice(c *gin.Context) {
	var payload structs.ObjPayloadInvoiceOrder
	if errPL := c.ShouldBindJSON(&payload); errPL != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errPL.Error(),
		})
		return
	}

	ORDId := payload.Order_id
	var ORD structs.InvoiceOrder
	if errCK := models.GetOrderId(ORDId, &ORD); errCK != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invoice invalid.",
			"data":    errCK.Error(),
		})
		return
	}

	if ORD.Or_total < 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "The net total cannot be negative.",
			"data":    "",
		})
		return
	}

	shopId := ORD.Shop_id
	userId := ORD.User_id
	dpmId := ORD.DpmId

	if ORD.Or_is_active != 1 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Order is active invalid. Please refresh.",
			"data":    "",
		})
		return
	}

	var docno structs.DocInvoice
	errDN := models.GetInvoiceDocNoData(shopId, &docno)
	if errDN != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Doc Data Invalid. Please Setting Doc no in shop setting.",
			"data":    errDN.Error(),
		})
		return
	}

	dateTime, _ := time.Parse(time.RFC3339, ORD.Or_datetime)
	ORD.Or_datetime = dateTime.Format("2006-01-02 15:04:05")

	createDatetime, _ := time.Parse(time.RFC3339, ORD.Or_create)
	ORD.Or_create = createDatetime.Format("2006-01-02 15:04:05")

	updateDateTime, _ := time.Parse(time.RFC3339, ORD.Or_update)
	ORD.Or_update = updateDateTime.Format("2006-01-02 15:04:05")

	var invoice_id = 0
	if ORD.Id > 0 {
		var ORDD []structs.InvoiceOrderDetail
		if errOR := models.GetOrderDetailId(ORDId, &ORDD); errOR != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Invoice Detail invalid.",
				"data":    "",
			})
			return
		}

		if len(ORDD) == 0 {
			c.JSON(200, gin.H{
				"status":  false,
				"message": "Please checked to add list order.",
				"data":    "",
			})
			return
		}

		for i, OrderDetail := range ORDD {
			var category_eclaim_id *int = nil
			var CEId structs.CategoryEclaimId
			if OrderDetail.Ord_type_id == 1 {
				if errCEId := models.GetCategoryEclaimCheck(*OrderDetail.Checking_id, &CEId); errCEId != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Invoice Detail invalid.",
						"data":    "",
					})
					return
				}
				category_eclaim_id = CEId.Category_eclaim_id
			} else if OrderDetail.Ord_type_id == 2 {
				if errCEId := models.GetCategoryEclaimCourse(*OrderDetail.Course_id, &CEId); errCEId != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Invoice Detail invalid.",
						"data":    "",
					})
					return
				}
				category_eclaim_id = CEId.Category_eclaim_id
			} else if OrderDetail.Ord_type_id == 3 {
				if errCEId := models.GetCategoryEclaimProduct(*OrderDetail.Product_id, *OrderDetail.Product_unit_id, ORD.Shop_id, &CEId); errCEId != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Invoice Detail invalid.",
						"data":    "",
					})
					return
				}
				category_eclaim_id = CEId.Category_eclaim_id
			}
			ORDD[i].Category_eclaim_id = category_eclaim_id
		}
		if len(ORDD) != 0 {
			invoiceId, err := models.AddInvoice(&ORD, &ORDD, convertDocInvoice(&docno), userId, dpmId)
			if err != nil {
				c.JSON(200, gin.H{
					"status":  false,
					"message": "Cannot Add Invoice",
					"data":    err.Error(),
				})
				return
			}
			invoice_id = invoiceId
		}
		//Add Sticker
		var Invd []structs.StickerInvoiceDetail
		if errInvdc := models.GetStickerInvoiceDetailProduct(invoice_id, &Invd); errInvdc != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Invoice detail product invalid.",
				"data":    errInvdc.Error(),
			})
			return
		}
		for _, Invdproduct := range Invd {
			if Invdproduct.Course_id == nil {
				var Invd_qty float64 = 0
				if Invdproduct.Invd_type_id == 3 {
					Invd_qty = Invdproduct.Invd_qty * Invdproduct.Invd_rate
				} else {
					Invd_qty = Invdproduct.Invd_qty
				}
				var Invoice_id = Invdproduct.Invoice_id
				var Invoice_detail_id = Invdproduct.Id
				productSticker := structs.StickerProcessProduct{
					Shop_id:           shopId,
					Customer_id:       ORD.Customer_id,
					User_id:           userId,
					Product_id:        *Invdproduct.Product_id,
					Product_store_id:  *Invdproduct.Product_store_id,
					Invoice_id:        &Invoice_id,
					Invoice_detail_id: &Invoice_detail_id,
					Pdso_qty:          Invd_qty,
					Invd_code:         Invdproduct.Invd_code,
					Invd_name:         Invdproduct.Invd_name,
					Invd_unit:         Invdproduct.Invd_unit,
					Invd_price:        Invdproduct.Invd_price,
					Invd_topical:      Invdproduct.Invd_topical,
					Invd_direction:    Invdproduct.Invd_direction,
				}
				ProductSticker(&productSticker)
			}
		}

		//สร้างรายการตรวจ
		if ORD.Queue_id != nil {
			var queueDateTime structs.GetQueueDateTime
			if errqueueDateTime := models.GetQueueDateTime(*ORD.Queue_id, &queueDateTime); errqueueDateTime != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Queue DateTime invalid.",
					"data":    errqueueDateTime.Error(),
				})
				return
			}
			Chk_datetime, _ := time.Parse(time.RFC3339, queueDateTime.Que_datetime)
			var InvdCheck []structs.InvoiceDetail
			if errInvd := models.GetInvoiceDetailCheck(invoice_id, &InvdCheck); errInvd != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Invoice detail checking invalid.",
					"data":    errInvd.Error(),
				})
				return
			}
			for _, Invdcheck := range InvdCheck {
				var checking structs.CheckChecking
				if errChecking := models.GetCheckingChecksId(*Invdcheck.Checking_id, &checking); errChecking != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Checking invalid.",
						"data":    errChecking.Error(),
					})
					return
				}
				if checking.Checking_type_id != 4 {
					var i float64
					for i = 0; i < Invdcheck.Invd_qty; i++ {
						var Check_Id = 0

						invoiceCheck := structs.InvoiceCheck{
							Id:                0,
							Shop_id:           ORD.Shop_id,
							Invoice_id:        invoice_id,
							Invoice_detail_id: Invdcheck.Id,
							User_id:           userId,
							Customer_id:       ORD.Customer_id,
							Queue_id:          ORD.Queue_id,
							Checking_id:       *Invdcheck.Checking_id,
							Chk_type_id:       checking.Checking_type_id, //1 ทั่วไป, 2 lab, 3 X-Lay
							Chk_code:          Invdcheck.Invd_code,
							Chk_name:          Invdcheck.Invd_name,
							Chk_unit:          Invdcheck.Invd_unit,
							Chk_is_active:     1,
							Chk_datetime:      Chk_datetime.Format("2006-01-02 15:04:05"),
							Chk_create:        time.Now().Format("2006-01-02 15:04:05"),
							Chk_update:        time.Now().Format("2006-01-02 15:04:05"),
						}
						check_Id, errCCheck := models.CreateCheck(&invoiceCheck)
						if errCCheck != nil {
							c.AbortWithStatusJSON(200, gin.H{
								"status":  false,
								"message": "Create Check error.",
								"data":    "",
							})
							return
						}
						Check_Id = check_Id
						if errUID := models.UpdateInvoiceDetails(Invdcheck.Id, Check_Id); errUID != nil {
							c.AbortWithStatusJSON(200, gin.H{
								"status":  false,
								"message": "Cannot Update Check",
								"data":    errUID.Error(),
							})
							return
						}
					}
				}
			}
		}
		//จบสร้างรายการตรวจ
	}

	if invoice_id != 0 {
		errNewDocno := models.UpdateInvoiceDocno(shopId, &structs.DocInvoice{
			ShopId:                 shopId,
			Invoice_number_default: strconv.Itoa(libs.StrToInt(docno.Invoice_number_default) + 1),
		})
		if errNewDocno != nil {
			c.JSON(200, gin.H{
				"status":  false,
				"message": "Cannot Update Doc No Invoice",
				"data":    errNewDocno.Error(),
			})
			return
		}
	}

	models.AddLogInvoice(&structs.LogInvoice{
		Username:   c.Params.ByName("userEmail"),
		Log_type:   "Add Invoice",
		Log_text:   "Add Invoice Id = " + convertDocInvoice(&docno) + " (" + strconv.Itoa(invoice_id) + ")",
		Log_create: time.Now().Format("2006-01-02 15:04:05"),
		Shop_id:    libs.StrToInt(c.Params.ByName("shopId")),
	})

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Add Invoice Success.",
		"data":    invoice_id,
		"code":    convertDocInvoice(&docno),
	})
}

func InvoicesDetail(c *gin.Context) {
	INVId := libs.StrToInt(c.Params.ByName("id"))
	shopId := libs.StrToInt(c.Params.ByName("shopId"))
	var INVdetail structs.InvoiceId
	err := models.GetInvoiceDetail(INVId, shopId, &INVdetail)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Invoice Invalid!",
			"data":    err.Error(),
		})
		return
	} else if INVdetail.Id == 0 {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Invoice Invalid!",
			"data":    "",
		})
		return
	}

	var RCShop structs.ReceiptShop
	if errShop := models.GetShopReceiptById(INVdetail.Shop_id, &RCShop); errShop != nil || RCShop.Id == 0 {
		INVdetail.Shop = structs.ReceiptShop{}
	} else {
		INVdetail.Shop = RCShop
	}

	var RCCus structs.ObjQueryCustomer
	if errCus := models.GetCustomerById(INVdetail.Customer_id, &RCCus); errCus != nil || RCCus.ID == 0 {
		INVdetail.Customer = structs.ObjQueryCustomer{}
	} else {
		INVdetail.Customer = RCCus
	}

	var INVSubs []structs.InvoiceSub
	if errINVS := models.GetInvoiceSub(INVId, &INVSubs); errINVS != nil || len(INVSubs) == 0 {
		INVdetail.Subs = &[]structs.InvoiceSub{}
	} else {
		for i := range INVSubs {
			INVSubs[i].Invd_name = strings.TrimSpace(INVSubs[i].Invd_name)
			INVSubs[i].Invd_code = strings.TrimSpace(INVSubs[i].Invd_code)
		}
		INVdetail.Subs = &INVSubs
	}

	var wg sync.WaitGroup
	INVSub := *INVdetail.Subs
	for i, _ := range INVSub {
		wg.Add(1)
		var units []structs.ProductUnitList
		go func(st string, i int) {
			defer wg.Done()
			if INVSub[i].Product_id != nil {
				if errU := models.GetProductUnit(*INVSub[i].Product_id, shopId, &units); errU != nil || len(units) == 0 {
					INVSub[i].Units = &[]structs.ProductUnitList{}
				} else {
					INVSub[i].Units = &units
				}
			} else {
				INVSub[i].Units = &[]structs.ProductUnitList{}
			}
		}(INVSub[i].Invd_code, i)
	}
	wg.Wait()

	var INVT []structs.OrderTags
	if errINVT := models.GetOrderTags(INVdetail.Order_id, &INVT); errINVT != nil || len(INVT) == 0 {
		INVdetail.Tags = &[]structs.OrderTags{}
	} else {
		INVdetail.Tags = &INVT
	}

	if INVdetail.User_id_cancel != 0 {
		var Users structs.UserCancel
		if errUsers := models.GetUserCancel(INVdetail.User_id_cancel, &Users); errUsers != nil {
			INVdetail.User_fullname_cancel = ""
			INVdetail.User_fullname_en_cancel = ""
		} else {
			INVdetail.User_fullname_cancel = Users.User_fullname
			INVdetail.User_fullname_en_cancel = Users.User_fullname_en
		}
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    INVdetail,
	})
}

func DelInvoice(c *gin.Context) {
	INVId := libs.StrToInt(c.Params.ByName("id"))
	var INV structs.Invoice
	if errCK := models.GetInvoiceId(INVId, &INV); errCK != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invoice invalid.",
			"data":    errCK.Error(),
		})
		return
	}

	if INV.Id < 1 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invoice invalid.",
			"data":    INV,
		})
		return
	}
	userId := libs.StrToInt(c.Params.ByName("userID"))
	if errAD := models.CancelInvoice(INVId, userId, &INV); errAD != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Cannot Remove Invoice",
			"data":    errAD.Error(),
		})
		return
	}

	if errAD := models.UpdateInvoiceOrderCancal(INV.Order_id); errAD != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Cannot Update Order",
			"data":    errAD.Error(),
		})
		return
	}

	if errCS := models.UpdateCancalStickerInvoice(INVId); errCS != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Cannot Update Sticker",
			"data":    errCS.Error(),
		})
		return
	}

	//ยกเเลิกรายการตรวจ
	if errCS := models.UpdateCancalCheckInvoice(INVId); errCS != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Cannot Update Check",
			"data":    errCS.Error(),
		})
		return
	}
	//จบยกเเลิกรายการตรวจ

	models.AddLogInvoice(&structs.LogInvoice{
		Username:   c.Params.ByName("userEmail"),
		Log_type:   "Cancel Invoice",
		Log_text:   "Cancel Invoice Id = " + INV.Inv_code + " (" + strconv.Itoa(INVId) + ")",
		Log_create: time.Now().Format("2006-01-02 15:04:05"),
		Shop_id:    libs.StrToInt(c.Params.ByName("shopId")),
	})

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Cancel Invoice ID " + INV.Inv_code + " success",
		"data":    "",
	})
}

func convertDocInvoice(docno *structs.DocInvoice) string {

	middleFormat := ""

	switch docno.Invoice_type {
	case 1:
		middleFormat = ""
		break
	case 2: // YYYYx
		middleFormat = time.Now().Format("2006")
		break
	case 3: // YYYYMMx
		middleFormat = time.Now().Format("200601")
		break
	case 4: // YYYYMMDDx
		middleFormat = time.Now().Format("20060102")
		break
	}
	runno := docno.Invoice_number_default

	switch docno.Invoice_number_digit {
	case 1:
		runno = docno.Invoice_number_default
		break
	case 2: // YYYYx
		runno = fmt.Sprintf("%02v", docno.Invoice_number_default)
		break
	case 3: // YYYYMMx
		runno = fmt.Sprintf("%03v", docno.Invoice_number_default)
		break
	case 4: // YYYYMMDDx
		runno = fmt.Sprintf("%04v", docno.Invoice_number_default)
		break
	}

	idDefault := ""
	if docno.Invoice_id_default != "" {
		idDefault = docno.Invoice_id_default
	}

	return idDefault + middleFormat + runno
}
