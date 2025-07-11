package controllers

import (
	"fmt"
	"linecrmapi/libs"
	"linecrmapi/models"
	"linecrmapi/structs"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// func STKSearch(c *gin.Context) {
// 	var filter structs.ObjPayloadSearchSticker
// 	if errSBJ := c.ShouldBindJSON(&filter); errSBJ != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Invalid request data.",
// 			"data":    errSBJ.Error(),
// 		})
// 		return
// 	}

// 	filter.Shop_id = libs.StrToInt(c.Params.ByName("shopId"))
// 	var countList []structs.Stickers
// 	if filter.ActivePage < 1 {
// 		filter.ActivePage = 0
// 	} else {
// 		filter.ActivePage -= 1
// 	}

// 	err := models.GetStickerList(filter, false, &countList)
// 	emptySlice := []string{}
// 	if err != nil {
// 		c.JSON(200, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 			"data":    emptySlice,
// 		})
// 		c.Abort()
// 	} else {
// 		var STKList []structs.Stickers
// 		if errMD := models.GetStickerList(filter, true, &STKList); errMD != nil {
// 			c.AbortWithStatusJSON(200, gin.H{
// 				"status":  false,
// 				"message": "Something went wrong.",
// 				"data":    "",
// 			})
// 			return
// 		}

// 		if len(STKList) == 0 {
// 			emptySTKList := []structs.StickersList{}
// 			c.JSON(200, gin.H{
// 				"status":  true,
// 				"message": "",
// 				"data": structs.ResponsePaginationSticker{
// 					Result_data:   emptySTKList,
// 					Count_of_page: 0,
// 					Count_all:     0,
// 				},
// 			})
// 			return
// 		}

// 		StickersList := []structs.StickersList{}
// 		for i, _ := range STKList {
// 			var STKP []structs.StickerProduct
// 			if errSTKP := models.GetStickerReceipt(STKList[i].Invoice_id, &STKP); errSTKP != nil {
// 				c.AbortWithStatusJSON(200, gin.H{
// 					"status":  false,
// 					"message": "Something went wrong.",
// 					"data":    "",
// 				})
// 				return
// 			}
// 			var num string = ""
// 			for j, stickerProduct := range STKP {
// 				if j == 0 {
// 					num = strconv.Itoa((j + 1)) + "." + stickerProduct.Sticker_name + "*" + strconv.Itoa(int(stickerProduct.Sticker_amount))
// 				} else {
// 					num = num + "<br>" + strconv.Itoa((j + 1)) + "." + stickerProduct.Sticker_name + "*" + strconv.Itoa(int(stickerProduct.Sticker_amount))
// 				}
// 			}
// 			STKList[i].Sticker_products = num
// 			StickersLists := structs.StickersList{
// 				Id:                  STKList[i].Id,
// 				Product_id:          STKList[i].Product_id,
// 				Customer_id:         STKList[i].Customer_id,
// 				Invoice_id:          STKList[i].Invoice_id,
// 				Inv_code:            STKList[i].Inv_code,
// 				Ctm_id:              STKList[i].Ctm_id,
// 				Ctm_prefix:          STKList[i].Ctm_prefix,
// 				Ctm_fname:           STKList[i].Ctm_fname,
// 				Ctm_lname:           STKList[i].Ctm_lname,
// 				Ctm_fname_en:        STKList[i].Ctm_fname_en,
// 				Ctm_lname_en:        STKList[i].Ctm_lname_en,
// 				User_id:             STKList[i].User_id,
// 				User_fullname:       STKList[i].User_fullname,
// 				User_fullname_en:    STKList[i].User_fullname_en,
// 				Sticker_products:    STKList[i].Sticker_products,
// 				Sticker_direction:   STKList[i].Sticker_direction,
// 				Sticker_active_id:   STKList[i].Sticker_active_id,
// 				Sticker_print_label: STKList[i].Sticker_print_label,
// 				Sticker_print_order: STKList[i].Sticker_print_order,
// 				Sticker_is_del:      STKList[i].Sticker_is_del,
// 				Sticker_modify:      STKList[i].Sticker_modify,
// 				StickerProduct:      STKP,
// 				ShopId:              STKList[i].ShopId,
// 				OrderId:             STKList[i].OrderId,
// 				StickerCode:         STKList[i].StickerCode,
// 				StickerName:         STKList[i].StickerName,
// 				StickerAmount:       STKList[i].StickerAmount,
// 				StickerUnit:         STKList[i].StickerUnit,
// 				StickerExpdate:      STKList[i].StickerExpdate,
// 				StickerTopical:      STKList[i].StickerTopical,
// 			}
// 			StickersList = append(StickersList, StickersLists)
// 		}

// 		res := structs.ResponsePaginationSticker{
// 			Result_data:   StickersList,
// 			Count_of_page: len(StickersList),
// 			Count_all:     len(countList),
// 		}

// 		c.JSON(200, gin.H{
// 			"status":  true,
// 			"message": "",
// 			"data":    res,
// 		})
// 	}
// }

func STKSearch(c *gin.Context) {
	var filter structs.ObjPayloadSearchSticker
	if errSBJ := c.ShouldBindJSON(&filter); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errSBJ.Error(),
		})
		return
	}

	filter.Shop_id = libs.StrToInt(c.Params.ByName("shopId"))
	var countList []structs.Stickers
	if filter.ActivePage < 1 {
		filter.ActivePage = 0
	} else {
		filter.ActivePage -= 1
	}

	err := models.GetStickerList(filter, false, &countList)
	emptySlice := []string{}
	if err != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": err.Error(),
			"data":    emptySlice,
		})
		c.Abort()
	} else {
		var STKList []structs.Stickers
		if errMD := models.GetStickerList(filter, true, &STKList); errMD != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Something went wrong.",
				"data":    "",
			})
			return
		}

		if len(STKList) == 0 {
			emptySTKList := []structs.StickersList{}
			c.JSON(200, gin.H{
				"status":  true,
				"message": "",
				"data": structs.ResponsePaginationSticker{
					Result_data:   emptySTKList,
					Count_of_page: 0,
					Count_all:     0,
				},
			})
			return
		}

		StickersList := []structs.StickersList{}
		var ItemListID []int
		for _, data := range STKList {
			ItemListID = append(ItemListID, data.Invoice_id)
		}

		var STKP []structs.StickerProduct
		if errSTKP := models.GetStickerReceiptIdInArray(ItemListID, &STKP); errSTKP != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Something went wrong.",
				"data":    "",
			})
			return
		}

		var wg sync.WaitGroup
		var mu sync.Mutex
		for i, _ := range STKList {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()

				// var num string = ""
				// for j, stickerProduct := range STKP {
				// 	if j == 0 {
				// 		num = strconv.Itoa((j + 1)) + "." + stickerProduct.Sticker_name + "*" + strconv.Itoa(int(stickerProduct.Sticker_amount))
				// 	} else {
				// 		num = num + "<br>" + strconv.Itoa((j + 1)) + "." + stickerProduct.Sticker_name + "*" + strconv.Itoa(int(stickerProduct.Sticker_amount))
				// 	}
				// }
				STKList[i].Sticker_products = ""
				StickersLists := structs.StickersList{
					Id:                  STKList[i].Id,
					Product_id:          STKList[i].Product_id,
					Customer_id:         STKList[i].Customer_id,
					Invoice_id:          STKList[i].Invoice_id,
					Inv_code:            STKList[i].Inv_code,
					Ctm_id:              STKList[i].Ctm_id,
					Ctm_prefix:          STKList[i].Ctm_prefix,
					Ctm_fname:           STKList[i].Ctm_fname,
					Ctm_lname:           STKList[i].Ctm_lname,
					Ctm_fname_en:        STKList[i].Ctm_fname_en,
					Ctm_lname_en:        STKList[i].Ctm_lname_en,
					User_id:             STKList[i].User_id,
					User_fullname:       STKList[i].User_fullname,
					User_fullname_en:    STKList[i].User_fullname_en,
					Sticker_products:    STKList[i].Sticker_products,
					Sticker_direction:   STKList[i].Sticker_direction,
					Sticker_active_id:   STKList[i].Sticker_active_id,
					Sticker_print_label: STKList[i].Sticker_print_label,
					Sticker_print_order: STKList[i].Sticker_print_order,
					Sticker_is_del:      STKList[i].Sticker_is_del,
					Sticker_modify:      STKList[i].Sticker_modify,
					ShopId:              STKList[i].ShopId,
					OrderId:             STKList[i].OrderId,
					StickerCode:         STKList[i].StickerCode,
					StickerName:         STKList[i].StickerName,
					StickerAmount:       STKList[i].StickerAmount,
					StickerUnit:         STKList[i].StickerUnit,
					StickerExpdate:      STKList[i].StickerExpdate,
					StickerTopical:      STKList[i].StickerTopical,
				}

				for _, dataDetail := range STKP {
					if dataDetail.Invoice_id == STKList[i].Invoice_id {
						StickersLists.StickerProduct = append(StickersLists.StickerProduct, structs.StickerProduct{
							Id:                dataDetail.Id,
							Invoice_id:        dataDetail.Invoice_id,
							Sticker_code:      dataDetail.Sticker_code,
							Sticker_name:      dataDetail.Sticker_name,
							Sticker_name_acc:  dataDetail.Sticker_name_acc,
							Sticker_amount:    dataDetail.Sticker_amount,
							Sticker_price:     dataDetail.Sticker_price,
							Sticker_topical:   dataDetail.Sticker_topical,
							Sticker_direction: dataDetail.Sticker_direction,
							Sticker_unit:      dataDetail.Sticker_unit,
							Sticker_expdate:   dataDetail.Sticker_expdate,
							Sticker_modify:    dataDetail.Sticker_modify,
							Invd_discount:     dataDetail.Invd_discount,
							Invd_amount:       dataDetail.Invd_amount,
							Invd_total:        dataDetail.Invd_total,
							Invd_vat:          dataDetail.Invd_vat,
							Pd_code_acc:       dataDetail.Pd_code_acc,
							Pd_name_acc:       dataDetail.Pd_name_acc,
						})
					}
				}
				mu.Lock()
				StickersList = append(StickersList, StickersLists)
				mu.Unlock()
			}(i)
		}
		wg.Wait()
		sort.Slice(StickersList, func(i, j int) bool {
			layout := "2006-01-02 15:04:05" // <-- ปรับตามรูปแบบวันที่จริงใน Sticker_modify
			timeI, err1 := time.Parse(layout, StickersList[i].Sticker_modify)
			timeJ, err2 := time.Parse(layout, StickersList[j].Sticker_modify)

			if err1 != nil || err2 != nil {
				// ถ้าแปลงไม่ได้ ให้เปรียบเทียบเป็น string fallback
				return StickersList[i].Sticker_modify > StickersList[j].Sticker_modify
			}

			return timeI.After(timeJ) // 🔁 เรียงจากใหม่ → เก่า
		})

		res := structs.ResponsePaginationSticker{
			Result_data:   StickersList,
			Count_of_page: len(StickersList),
			Count_all:     len(countList),
		}

		c.JSON(200, gin.H{
			"status":  true,
			"message": "",
			"data":    res,
		})
	}
}

func GetSTKList(c *gin.Context) {
	var IdList structs.ObjPayloadLists
	if errSBJ := c.ShouldBindJSON(&IdList); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errSBJ.Error(),
		})
		return
	}

	if len(*IdList.Invoice_id) == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Receipt Id invalid.",
			"data":    []structs.StickerProduct{},
		})
		return
	}

	var tmp []string

	for _, v := range *IdList.Invoice_id {
		tmp = append(tmp, fmt.Sprint(v))
	}

	var STKDList []structs.StickerDetail
	if errSTKP := models.GetStickerReceiptInArray(strings.Join(tmp, ","), &STKDList); errSTKP != nil || len(STKDList) == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Sticker invalid.",
			"data":    []structs.StickerProduct{},
		})
		return
	}

	// UpdatePrintStickerStatus
	var wg sync.WaitGroup
	for i, _ := range STKDList {
		wg.Add(1)
		go func(st string, i int) {
			defer wg.Done()
			if STKDList[i].Id != 0 {
				if errUp := models.UpdatePrintStickerStatus(*&STKDList[i].Id); errUp != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Sticker invalid.",
						"data":    errUp.Error(),
					})
					return
				}
			}
		}(STKDList[i].Sticker_code, i)
	}
	wg.Wait()

	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    STKDList,
	})
}

func GetSTKDetail(c *gin.Context) {
	InvId := libs.StrToInt(c.Params.ByName("invoice_id"))
	StkId := libs.StrToInt(c.Params.ByName("sticker_id"))
	ActvId := libs.StrToInt(c.Params.ByName("actv_id"))
	// var STKP []structs.StickerDetail
	// if errSTKP := models.GetStickerInvoiceList(InvId, &STKP); errSTKP != nil || len(STKP) == 0 {
	// 	c.AbortWithStatusJSON(200, gin.H{
	// 		"status":  false,
	// 		"message": "Sticker invalid.",
	// 		"data":    []structs.StickerProduct{},
	// 	})
	// 	return
	// }

	// ตรวจสอบว่าไม่มีค่าใดๆ เลย
	if InvId == 0 && StkId == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request. Missing invoice_id or sticker_id.",
			"data":    []structs.StickerProduct{},
		})
		return
	}

	var STKP []structs.StickerDetail
	// ใช้ InvId ก่อน ถ้าไม่มีค่าให้ใช้ OrId
	if InvId != 0 {
		if errSTKP := models.GetStickerInvoiceList(InvId, &STKP); errSTKP != nil || len(STKP) == 0 {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Sticker invalid for invoice.",
				"data":    []structs.StickerProduct{},
			})
			return
		}
	} else if StkId != 0 {
		if errSTKP := models.GetStickerNoInvoiceList(StkId, ActvId, &STKP); errSTKP != nil || len(STKP) == 0 {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Sticker invalid for order.",
				"data":    []structs.StickerProduct{},
			})
			return
		}
	}

	var wg sync.WaitGroup
	for i, _ := range STKP {
		wg.Add(1)
		go func(st string, i int) {
			defer wg.Done()
			if errUp := models.UpdatePrintStickerStatus(*&STKP[i].Id); errUp != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Sticker invalid.",
					"data":    errUp.Error(),
				})
				return
			}
		}(STKP[i].Sticker_code, i)
	}
	wg.Wait()

	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    STKP,
	})
}

func GetSTKDetailList(c *gin.Context) {
	var IdList structs.ObjPayloadLists
	if errSBJ := c.ShouldBindJSON(&IdList); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errSBJ.Error(),
		})
		return
	}

	if len(*IdList.Invoice_id) == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Receipt Id invalid.",
			"data":    []structs.StickerDetail2{},
		})
		return
	}

	var tmp []string

	for _, v := range *IdList.Invoice_id {
		tmp = append(tmp, fmt.Sprint(v))
	}

	var STKList []structs.StickerDetail2
	if errSTKP := models.GetStickerHeadList(strings.Join(tmp, ","), &STKList); errSTKP != nil || len(STKList) == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Sticker invalid.",
			"data":    []structs.StickerDetail2{},
		})
		return
	}
	// UpdatePrintStickerStatus
	var wg sync.WaitGroup
	for i, _ := range STKList {
		wg.Add(1)
		go func(st string, i int) {
			defer wg.Done()
			if STKList[i].Id != 0 {
				var STKP []structs.StickerProduct
				if errSTKP := models.GetStickerReceipt(STKList[i].Invoice_id, &STKP); errSTKP != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Something went wrong.",
						"data":    "",
					})
					return
				}
				STKList[i].StickerProduct = STKP

				if errUp := models.UpdatePrintStickerStatus(*&STKList[i].Id); errUp != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Sticker invalid.",
						"data":    errUp.Error(),
					})
					return
				}
			}
		}(STKList[i].Inv_code, i)
	}
	wg.Wait()

	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    STKList,
	})
}

func GetPrescription(c *gin.Context) {
	InvId := libs.StrToInt(c.Params.ByName("invoice_id"))
	var RCData structs.Prescription
	err := models.GetReceipInPrescriptionById(InvId, &RCData)
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

	var RCShop structs.StickerShop
	if errShop := models.GetShopStickerById(RCData.Shop_id, &RCShop); errShop != nil || RCShop.Id == 0 {
		RCData.Shop = structs.StickerShop{}
	} else {
		RCData.Shop = RCShop
	}

	var RCCus structs.ObjQueryCustomer
	if errCus := models.GetCustomerById(RCData.Customer_id, &RCCus); errCus != nil || RCCus.ID == 0 {
		RCData.Customer = structs.ObjQueryCustomer{}
	} else {
		RCData.Customer = RCCus
	}

	var STKP []structs.StickerDetail
	if errSTKP := models.GetStickerInvoiceList(InvId, &STKP); errSTKP != nil || len(STKP) == 0 {
		RCData.Subs = []structs.StickerDetail{}
	} else {
		RCData.Subs = STKP
	}

	var wg sync.WaitGroup
	for i, _ := range STKP {
		wg.Add(1)
		go func(st string, i int) {
			defer wg.Done()
			if errUp := models.UpdatePrescriptionStatus(*&STKP[i].Id); errUp != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Prescription invalid.",
					"data":    errUp.Error(),
				})
				return
			}
		}(STKP[i].Sticker_code, i)
	}
	wg.Wait()

	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    RCData,
	})

}

func GetSettingSticker(c *gin.Context) {
	Shop_id := libs.StrToInt(c.Params.ByName("shopId"))
	var DS structs.StickerDocSetting
	if errDS := models.GetDocSettingSticker(Shop_id, &DS); errDS != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Doc Setting Invalid",
			"data":    errDS.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    DS,
	})
}
func UpdateSettingSticker(c *gin.Context) {
	var payload structs.StickerDocSetting
	if errPL := c.ShouldBindJSON(&payload); errPL != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errPL.Error(),
		})
		return
	}

	Shop_id := libs.StrToInt(c.Params.ByName("shopId"))
	err := models.UpdateDocSettingSticker(Shop_id, &payload)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Cannot Edit Sticker Setting",
			"data":    err.Error(),
		})
		return
	}

	models.AddLogStk(&structs.LogStk{
		Username:   c.Params.ByName("userEmail"),
		Log_type:   "Update doc_setting sticker",
		Log_text:   "Update doc_setting sticker",
		Log_create: time.Now().Format("2006-01-02 15:04:05"),
		Shop_id:    libs.StrToInt(c.Params.ByName("shopId")),
	})

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Edit Sticker Setting Success.",
		"data":    "",
	})

}

func UpdateSticker(c *gin.Context) {
	var payload structs.UpdateSticker
	if errPL := c.ShouldBindJSON(&payload); errPL != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errPL.Error(),
		})
		return
	}

	err := models.UpdateSticker(&payload)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Cannot edit sticker topical and direction.",
			"data":    err.Error(),
		})
		return
	}

	models.AddLogStk(&structs.LogStk{
		Username:   c.Params.ByName("userEmail"),
		Log_type:   "Update sticker topical and direction",
		Log_text:   "Update sticker topical and direction",
		Log_create: time.Now().Format("2006-01-02 15:04:05"),
		Shop_id:    libs.StrToInt(c.Params.ByName("shopId")),
	})

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Edit sticker topical and direction success.",
		"data":    "",
	})

}

func CreatePrintDruglabel(c *gin.Context) {

	var request structs.PrintStickerDrugLabel
	// รับค่าจาก Request Body
	if errPL := c.ShouldBindJSON(&request); errPL != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errPL.Error(),
		})
		return
	}
	// productStickerUpdate := structs.PrintStickerDrugLabel{
	// 	Id:                  0,
	// 	Product_id:          data.Product_id,
	// 	Customer_id:         data.Customer_id,
	// 	User_id:             data.User_id,
	// 	ShopId:              shopId,
	// 	Sticker_code:        data.Sticker_code,
	// 	Sticker_name:        data.Sticker_name,
	// 	Sticker_name_acc:    data.Sticker_name_acc,
	// 	Sticker_amount:      data.Sticker_amount,
	// 	Sticker_unit:        data.Sticker_unit,
	// 	Sticker_unit_en:     data.Sticker_unit_en,
	// 	Sticker_price:       data.Sticker_price,
	// 	Sticker_expdate:     data.Sticker_expdate,
	// 	Sticker_active_id:   2,
	// 	Sticker_print_label: 0,
	// 	Sticker_print_order: 0,
	// 	Sticker_topical:     data.Sticker_topical,
	// 	Sticker_direction:   data.Sticker_direction,
	// 	Sticker_is_del:      0,
	// 	Sticker_modify:      time.Now().Format("2006-01-02 15:04:05"),
	// }

	// if err := models.CreateStickerDrugLabel(ctx, &productStickerUpdate); err != nil {
	//     fmt.Println("Error creating sticker label:", err)
	//     return
	// }
	stickerExpDate, err := time.Parse(time.RFC3339, request.Sticker_expdate)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid date format"})
		return
	}
	stickerExpDateFormatted := stickerExpDate.Format("2006-01-02")
	PrintSticker := structs.PrintStickerDrugLabel{
		Product_id:          request.Product_id,
		Customer_id:         request.Customer_id,
		User_id:             request.User_id,
		ShopId:              request.ShopId,
		Sticker_code:        request.Sticker_code,
		Sticker_name:        request.Sticker_name,
		Sticker_name_acc:    request.Sticker_name_acc,
		Sticker_amount:      request.Sticker_amount,
		Sticker_unit:        request.Sticker_unit,
		Sticker_unit_en:     request.Sticker_unit_en,
		Sticker_price:       request.Sticker_price,
		Sticker_expdate:     stickerExpDateFormatted,
		Sticker_active_id:   2,
		Sticker_print_label: 0,
		Sticker_print_order: 0,
		Sticker_topical:     request.Sticker_topical,
		Sticker_direction:   request.Sticker_direction,
		Sticker_is_del:      0,
		Sticker_modify:      time.Now().Format("2006-01-02 15:04:05"),
	}

	// บันทึกข้อมูลลงฐานข้อมูล
	err = models.CreateStickerDrugLabel(&PrintSticker)
	if err != nil {
		c.JSON(200, gin.H{"error": "Failed to create sticker"})
		return
	}

	models.AddLogStk(&structs.LogStk{
		Username:   c.Params.ByName("userEmail"),
		Log_type:   "Add Sticker Type 2",
		Log_text:   "Add Sticker Type 2 Name " + request.Sticker_name,
		Log_create: time.Now().Format("2006-01-02 15:04:05"),
		Shop_id:    request.ShopId,
	})

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Sticker created successfully",
		"data":    PrintSticker,
	})
}
