package controllers

import (
	"fmt"
	"linecrmapi/libs"
	"linecrmapi/models"
	"linecrmapi/structs"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func AddOrder(c *gin.Context) {

	var payload structs.ObjPayloadAddOrder
	if errPL := c.ShouldBindJSON(&payload); errPL != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errPL.Error(),
		})
		return
	}

	if *payload.Or_total < 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "The net total cannot be negative.",
			"data":    "",
		})
		return
	}

	Shop_id := payload.Shop_id
	var ShopReadResponse structs.ShopReadResponse
	if errShop := models.GetShopById(Shop_id, &ShopReadResponse); errShop != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Shop Invalid!",
			"data":    errShop.Error(),
		})
		return
	}
	shopMId := ShopReadResponse.ShopMotherId

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

	var CustomerOnline structs.CustomerOnline
	if errCO := models.GetCustomerOnlineById(payload.Customer_online_id, &CustomerOnline); errCO != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Customer Online Invalid!",
			"data":    errCO.Error(),
		})
		return
	}
	var Customer_id int = 0
	var Customer_Group_Id int = 0
	var objQueryCustomer structs.ObjQueryCustomer
	errCustomer := models.GetCustomerShopById(CustomerOnline.Co_citizen_id, payload.Shop_id, &objQueryCustomer)
	if errCustomer != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Get data error.",
			"data":    "",
		})
		return
	} else {
		if objQueryCustomer.ID == 0 {
			//Add Customer
			var objQueryCustomerDocSetting structs.ObjQueryCustomerDocSetting
			errGetCustomerDocSetting := models.GetCustomerDocSetting(Shop_id, &objQueryCustomerDocSetting)
			if errGetCustomerDocSetting != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Get doc setting error.",
					"data":    errGetCustomerDocSetting,
				})
				return
			}
			CtmId := libs.SetDocSettingCode(objQueryCustomerDocSetting.CustomerIdDefault, objQueryCustomerDocSetting.CustomerNumberDigit, objQueryCustomerDocSetting.CustomerNumberDefault, objQueryCustomerDocSetting.CustomerType)
			nextNumberDefault := objQueryCustomerDocSetting.CustomerNumberDefault + 1

			var CustomerGroupId structs.CustomerGroupId
			if errCustomerGroup := models.GetCustomerGroup(Shop_id, &CustomerGroupId); errCustomerGroup != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Get Customer Group Error.",
					"data":    errGetCustomerDocSetting,
				})
				return
			} else {
				if CustomerGroupId.Id == 0 {
					AddCustomerGroup := structs.AddCustomerGroup{
						Shop_id:      Shop_id,
						Cg_name:      "Online",
						Cg_is_active: 1,
						Cg_is_online: 1,
						Cg_create:    time.Now().Format("2006-01-02 15:04:05"),
						Cg_update:    time.Now().Format("2006-01-02 15:04:05"),
					}
					errCreateCGroup := models.AddCustomerGroup(&AddCustomerGroup)
					if errCreateCGroup != nil {
						c.AbortWithStatusJSON(200, gin.H{
							"status":  false,
							"message": "Create Customer Group Online Error.",
							"data":    "",
						})
						return
					}
					Customer_Group_Id = AddCustomerGroup.Id
				} else {
					Customer_Group_Id = CustomerGroupId.Id
				}
			}

			objQueryCreateCustomer := models.Customer{
				ShopId:           Shop_id,
				ShopMotherId:     shopMId,
				CustomerGroupId:  Customer_Group_Id,
				UserId:           userId,
				CtmId:            CtmId,
				CtmCitizenId:     payload.Co_citizen_id,
				CtmPrefix:        payload.Co_prefix,
				CtmFname:         payload.Co_fname,
				CtmLname:         payload.Co_lname,
				CtmGender:        payload.Co_gender,
				CtmEmail:         payload.Co_email,
				CtmTel:           payload.Co_tel,
				CtmBirthdate:     payload.Co_birthdate,
				CtmAddress:       payload.Co_address,
				CtmDistrict:      payload.Co_district,
				CtmAmphoe:        payload.Co_amphoe,
				CtmProvince:      payload.Co_province,
				CtmZipcode:       payload.Co_zipcode,
				RightTreatmentId: 16,
				CtmPoint:         0,
				CtmCoin:          0,
				CtmIsActive:      1,
				CtmIsDel:         0,
				CtmCreate:        time.Now().Format("2006-01-02 15:04:05"),
				CtmUpdate:        time.Now().Format("2006-01-02 15:04:05"),
			}
			errCreate := models.CreateCustomer(&objQueryCreateCustomer)
			if errCreate != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Create customer error.",
					"data":    "",
				})
				return
			}
			Customer_id = objQueryCreateCustomer.ID
			var objDocSetting models.DocSetting
			errUpdateCustomerDocSetting := models.UpdateCustomerDocSetting(Shop_id, nextNumberDefault, &objDocSetting)
			if errUpdateCustomerDocSetting != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Update doc setting error.",
					"data":    errUpdateCustomerDocSetting.Error(),
				})
				return
			}
		} else {
			Customer_id = objQueryCustomer.ID
		}
	}

	var Queue_id int = 0
	var CreateQCheck int = 0
	for _, sub := range payload.Subs {
		if sub.Ord_type_id == 1 {
			CreateQCheck = 1
		}
	}
	if CreateQCheck == 1 {
		var RoomBedId structs.RoomBedId
		if errRoomBedId := models.GetRoomBedId(Shop_id, &RoomBedId); errRoomBedId != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Get Room Bed Id Error.",
				"data":    errRoomBedId.Error(),
			})
			return
		}
		if RoomBedId.Room_id == 0 || RoomBedId.Bed_id == 0 {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Get Room Bed Id Error.",
				"data":    "",
			})
			return
		}
		ObjPayloadCreateQueue := structs.ObjPayloadCreateQueueByOrder{
			ShopId:          Shop_id,
			CustomerId:      Customer_id,
			DoctorId:        userId,
			DoctorFullname:  User_fullname,
			RoomId:          RoomBedId.Room_id,
			BedId:           RoomBedId.Bed_id,
			QueUserId:       userId,
			QueUserFullname: User_fullname,
			QueTypeId:       1,
			QueAdmisId:      2,
			QuePriorityId:   1,
			QueDatetime:     time.Now().Format("2006-01-02 15:04:05"),
			QueTeleCode:     "",
			QueTeleUrl:      "",
		}
		Queue_id = CreateQueue(&ObjPayloadCreateQueue)
	}

	layout := "2006-01-02 15:04:05"
	time2 := time.Now().Format("15:04:05")
	ordate1, _ := time.Parse(layout, payload.Or_datetime)
	dateStr := fmt.Sprintf("%d-%02d-%02d", ordate1.Year(), ordate1.Month(), ordate1.Day())
	Ordate, errTfd := time.Parse(layout, dateStr+" "+time2)
	if errTfd != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "This Order Date.",
			"data":    errTfd.Error(),
		})
		return
	}

	POHAdd := structs.OrderDetail{
		Id:                  0,
		Shop_id:             Shop_id,
		User_id:             userId,
		Customer_id:         Customer_id,
		Customer_online_id:  payload.Customer_online_id,
		Or_fullname:         payload.Or_fullname,
		Or_tel:              payload.Or_tel,
		Or_email:            payload.Or_email,
		Or_address:          *payload.Or_address,
		Or_district:         payload.Or_district,
		Or_amphoe:           payload.Or_amphoe,
		Or_province:         payload.Or_province,
		Or_zipcode:          payload.Or_zipcode,
		Or_comment:          payload.Or_comment,
		Or_total_price:      *payload.Or_total_price,
		Or_discount_type_id: payload.Or_discount_type_id,
		Or_discount_item:    *payload.Or_discount_item,
		Or_discount_value:   *payload.Or_discount_value,
		Or_discount:         *payload.Or_discount,
		Or_befor_vat:        payload.Or_befor_vat,
		Tax_type_id:         payload.Tax_type_id,
		Tax_rate:            *payload.Tax_rate,
		Or_vat:              *payload.Or_vat,
		Or_total:            *payload.Or_total,
		Or_is_active:        1,
		Or_datetime:         Ordate.String(),
		Or_create:           time.Now().Format("2006-01-02 15:04:05"),
		Or_update:           time.Now().Format("2006-01-02 15:04:05"),
		Tags:                *&payload.Tags,
		Subs:                &payload.Subs,
		Queue_id:            Queue_id,
		DpmId:               payload.DpmId,
		Or_eclaim_id:        payload.Or_eclaim_id,
		Or_eclaim_rate:      payload.Or_eclaim_rate,
		Or_eclaim_over:      payload.Or_eclaim_over,
		Or_eclaim_total:     payload.Or_eclaim_total,
		Or_tele_code:        "",
	}

	orderId, err := models.AddOrder(&POHAdd)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Cannot Add Order",
			"data":    err.Error(),
		})
		return
	}

	if orderId == 0 {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Cannot Add Order",
			"data":    orderId,
		})
		return
	}

	//เช็คยา
	var Ord []structs.OrderDetailCheck
	if errOrd := models.GetOrderDetailProductCheck(orderId, &Ord); errOrd != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Order detail product invalid.",
			"data":    errOrd.Error(),
		})
		return
	}
	for _, Ordproduct := range Ord {
		var productStoreOrder structs.ProcessProductStoreOrderCheck
		if errOrdp := models.GetProcessProductStoreOrderCheck(*Ordproduct.Product_id, *Ordproduct.Product_store_id, &productStoreOrder); errOrdp != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Order detail product store invalid.",
				"data":    errOrdp.Error(),
			})
			return
		}
		if productStoreOrder.Id == 0 {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Order detail product empty.",
				"data":    Ordproduct.Ord_name + " ( " + Ordproduct.Ord_code + " ) ",
			})
			return
		}
		var OrderDetailCheckCourse structs.OrderDetailCheckCourse
		if errOrdpC := models.GetOrderDetailProductCheckCourse(orderId, *Ordproduct.Product_store_id, &OrderDetailCheckCourse); errOrdpC != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Order detail product invalid.",
				"data":    errOrdpC.Error(),
			})
			return
		}
		var Ord_qty float64 = Ordproduct.Ord_qty
		if OrderDetailCheckCourse.Ord_qty != 0 {
			Ord_qty = Ordproduct.Ord_qty - OrderDetailCheckCourse.Ord_qty
		}
		if productStoreOrder.Pdso_total < Ord_qty {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Order detail product are not enough to pay.",
				"data":    Ordproduct.Ord_name + " ( " + Ordproduct.Ord_code + " ) ",
			})
			return
		}
	}

	var LabplusRequestOrders structs.LabplusRequestOrder
	var RequestOrders *libs.RequestOrder
	if Queue_id != 0 {
		//CheckQueueStatusID
		var Queuefullname structs.CheckQueueStatusID
		if errQueuefullname := models.CheckQueueStatusID(Queue_id, &Queuefullname); errQueuefullname != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Err Queue fullname.",
				"data":    "",
			})
			return
		}
		//Labplus
		var CheckingLabplus []structs.CheckingLabplus
		if errCheckingLabplus := models.CheckingLabplus(Queue_id, &CheckingLabplus); errCheckingLabplus != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Something went wrong.",
				"data":    "",
			})
			return
		}
		if len(CheckingLabplus) > 0 {
			Labplus := structs.Labplus{}
			if errLabplus := models.GetShopLabplus(Shop_id, &Labplus); errLabplus != nil {
				c.JSON(200, gin.H{
					"status":  false,
					"message": "Cannot Queue Checking.",
					"data":    errLabplus.Error(),
				})
				return
			}
			if Labplus.Id != 0 {
				token := LabplusAuthen(Shop_id)
				if token != "token" {
					var CustomerPatient structs.CustomerPatient
					if errCustomerPatient := models.GetCustomerLabplus(Customer_id, &CustomerPatient); errCustomerPatient != nil {
						c.AbortWithStatusJSON(200, gin.H{
							"status":  false,
							"message": "Something went wrong.",
							"data":    "",
						})
						return
					}
					Ctm_birthdate, _ := time.Parse(time.RFC3339, CustomerPatient.Ctm_birthdate)
					var Ctm_gender string = "U"
					if CustomerPatient.Ctm_gender == "หญิง" {
						Ctm_gender = "F"
					} else if CustomerPatient.Ctm_gender == "Female" {
						Ctm_gender = "F"
					} else if CustomerPatient.Ctm_gender == "ชาย" {
						Ctm_gender = "M"
					} else if CustomerPatient.Ctm_gender == "Male" {
						Ctm_gender = "M"
					}
					Patient := structs.Patient{
						HN:          CustomerPatient.Ctm_id,
						FullName:    CustomerPatient.Ctm_prefix + " " + CustomerPatient.Ctm_fname + " " + CustomerPatient.Ctm_lname,
						TName:       CustomerPatient.Ctm_prefix,
						FName:       CustomerPatient.Ctm_fname,
						LName:       CustomerPatient.Ctm_lname,
						IDCard:      CustomerPatient.Ctm_citizen_id,
						Birthday:    Ctm_birthdate.Format("2006-01-02"),
						Sex:         Ctm_gender,
						DoctorName:  Queuefullname.Que_user_fullname,
						RequestNote: strconv.Itoa(Queue_id),
					}
					var LabOrders = []structs.LabOrder{}
					for _, dataCheckingLabplus := range CheckingLabplus {
						LabOrder := structs.LabOrder{
							Code: dataCheckingLabplus.Checking_code,
							Name: dataCheckingLabplus.Checking_name,
						}
						LabOrders = append(LabOrders, LabOrder)
					}
					LabplusRequestOrder := structs.LabplusRequestOrder{
						Patient:  Patient,
						LabOrder: LabOrders,
					}
					LabplusRequestOrders = LabplusRequestOrder
					RequestOrder, errRequestOrder := libs.LabplusRequestOrder(Labplus.Lapi_link, token, &LabplusRequestOrder)
					if errRequestOrder != nil {
						c.AbortWithStatusJSON(200, gin.H{
							"status":  false,
							"message": "Err RequestOrder.",
							"data":    errRequestOrder,
						})
						return
					}
					RequestOrders = RequestOrder
					if RequestOrder.Status_code == "201" {
						qlp_no := RequestOrder.Request_no
						AddQueuesLabplus := structs.AddQueuesLabplus{
							Shop_id:          Shop_id,
							Queue_id:         Queue_id,
							Qlp_no:           qlp_no,
							Qlp_message_th:   RequestOrder.Message_th,
							Qlp_process_code: "1",
							Qlp_process_name: "สถานะ: สั่งตรวจ",
							Qlp_datetime:     time.Now().Format("2006-01-02 15:04:05"),
							Qlp_update:       time.Now().Format("2006-01-02 15:04:05"),
						}
						models.AddQueuesLabplus(&AddQueuesLabplus)
					}
				}
			}
		}
	}

	models.AddLogOrder(&structs.LogOrders{
		Username:   User_email,
		Log_type:   "Add Order",
		Log_text:   "Add Order Id = " + strconv.Itoa(orderId),
		Log_create: time.Now().Format("2006-01-02 15:04:05"),
		Shop_id:    libs.StrToInt(c.Params.ByName("shopId")),
	})

	c.JSON(200, gin.H{
		"status":              true,
		"message":             "Add Order Success.",
		"data":                orderId,
		"LabplusRequestOrder": LabplusRequestOrders,
		"RequestOrder":        RequestOrders,
	})

}

func OrdersSearch(c *gin.Context) {
	var filter structs.ObjPayloadSearchOrder
	if errSBJ := c.ShouldBindJSON(&filter); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errSBJ.Error(),
		})
		return
	}

	filter.Shop_id = libs.StrToInt(c.Params.ByName("shopId"))

	var countList []structs.OrderLists
	if filter.ActivePage < 1 {
		filter.ActivePage = 0
	} else {
		filter.ActivePage -= 1
	}

	err := models.GetOrderList(filter, false, &countList)
	emptySlice := []string{}
	if err != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": err.Error(),
			"data":    emptySlice,
		})
		c.Abort()
	} else {
		var ORList []structs.OrderLists
		if errMD := models.GetOrderList(filter, true, &ORList); errMD != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Something went wrong.",
				"data":    "",
			})
			return
		}

		if len(ORList) == 0 {
			emptyORList := []structs.OrderLists{}
			c.JSON(200, gin.H{
				"status":  true,
				"message": "",
				"data": structs.ResponsePaginationOrderLists{
					Result_data:   emptyORList,
					Count_of_page: 0,
					Count_all:     0,
				},
			})
			return
		}

		for i := range ORList {
			if ORList[i].User_id_cancel != 0 {
				var Users structs.UserCancel
				if errUsers := models.GetUserCancel(ORList[i].User_id_cancel, &Users); errUsers != nil {
					ORList[i].User_fullname_cancel = ""
					ORList[i].User_fullname_en_cancel = ""
				} else {
					ORList[i].User_fullname_cancel = Users.User_fullname
					ORList[i].User_fullname_en_cancel = Users.User_fullname_en
				}
			}
		}

		res := structs.ResponsePaginationOrderLists{
			Result_data:   ORList,
			Count_of_page: len(ORList),
			Count_all:     len(countList),
		}

		c.JSON(200, gin.H{
			"status":  true,
			"message": "",
			"data":    res,
		})
	}
}

func OrdersDetail(c *gin.Context) {
	ORId := libs.StrToInt(c.Params.ByName("id"))
	shopId := libs.StrToInt(c.Params.ByName("shopId"))
	var ORdetail structs.OrderDetailData
	var or_total_price float64
	err := models.GetOrderIDDetail(ORId, shopId, &ORdetail)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Order Invalid!",
			"data":    err.Error(),
		})
		return
	} else if ORdetail.Id == 0 {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Order Invalid!",
			"data":    "",
		})
		return
	}

	var RCShop structs.ReceiptShop
	if errShop := models.GetShopReceiptById(ORdetail.Shop_id, &RCShop); errShop != nil || RCShop.Id == 0 {
		ORdetail.Shop = structs.ReceiptShop{}
	} else {
		ORdetail.Shop = RCShop
	}

	var RCCus structs.ObjQueryCustomer
	if errCus := models.GetCustomerById(ORdetail.Customer_id, &RCCus); errCus != nil || RCCus.ID == 0 {
		ORdetail.Customer = structs.ObjQueryCustomer{}
	} else {
		ORdetail.Customer = RCCus
	}

	var ORSubs []structs.OrderSub
	if errORS := models.GetOrderSub(ORId, &ORSubs); errORS != nil || len(ORSubs) == 0 {
		ORdetail.Subs = &[]structs.OrderSub{}
	} else {
		for i := range ORSubs {
			ORSubs[i].Ord_name = strings.TrimSpace(ORSubs[i].Ord_name)
			ORSubs[i].Ord_code = strings.TrimSpace(ORSubs[i].Ord_code)
		}
		ORdetail.Subs = &ORSubs
	}

	var wg sync.WaitGroup
	ORSub := *ORdetail.Subs
	for i, _ := range ORSub {
		wg.Add(1)
		or_total_price += ORSub[i].Ord_qty * ORSub[i].Ord_price
		var units []structs.ProductUnitList
		go func(st string, i int) {
			defer wg.Done()
			if ORSub[i].Product_id != nil {
				if errU := models.GetProductUnit(*ORSub[i].Product_id, shopId, &units); errU != nil || len(units) == 0 {
					ORSub[i].Units = &[]structs.ProductUnitList{}
				} else {
					ORSub[i].Units = &units
				}
			} else {
				ORSub[i].Units = &[]structs.ProductUnitList{}
			}
		}(ORSub[i].Ord_code, i)
	}
	wg.Wait()

	var ORT []structs.OrderTags
	if errORT := models.GetOrderTags(ORId, &ORT); errORT != nil || len(ORT) == 0 {
		ORdetail.Tags = &[]structs.OrderTags{}
	} else {
		ORdetail.Tags = &ORT
	}

	if ORdetail.User_id_cancel != 0 {
		var Users structs.UserCancel
		if errUsers := models.GetUserCancel(ORdetail.User_id_cancel, &Users); errUsers != nil {
			ORdetail.User_fullname_cancel = ""
			ORdetail.User_fullname_en_cancel = ""
		} else {
			ORdetail.User_fullname_cancel = Users.User_fullname
			ORdetail.User_fullname_en_cancel = Users.User_fullname_en
		}
	}
	ORdetail.Or_total_price = or_total_price
	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    ORdetail,
	})
}

func OrdersDetailByOrder(c *gin.Context) {
	ORId := libs.StrToInt(c.Params.ByName("orderId"))
	shopId := libs.StrToInt(c.Params.ByName("shopId"))

	var objQueryShopStore models.ShopStore
	errShopStore := models.GetShopStoreByIdType1(shopId, &objQueryShopStore)
	if errShopStore != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Get data shop store error!",
			"data":    errShopStore,
		})
	}

	var ORdetail structs.OrderDetail
	err := models.GetOrderDetail(ORId, &ORdetail)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Order Invalid!",
			"data":    err.Error(),
		})
		return
	} else if ORdetail.Id == 0 {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Order Invalid!",
			"data":    "",
		})
		return
	}

	var ORSubs []structs.OrderSub
	if errORS := models.GetOrderSub(ORId, &ORSubs); errORS != nil || len(ORSubs) == 0 {
		ORdetail.Subs = &[]structs.OrderSub{}
	} else {
		ORdetail.Subs = &ORSubs
	}

	var wg sync.WaitGroup
	ORSub := *ORdetail.Subs
	for i, _ := range ORSub {
		wg.Add(1)
		var units []structs.ProductUnitList
		go func(i int) {
			defer wg.Done()
			if ORSub[i].Product_id != nil {
				if errU := models.GetProductUnit(*ORSub[i].Product_id, shopId, &units); errU != nil || len(units) == 0 {
					ORSub[i].Units = &[]structs.ProductUnitList{}
				} else {
					index := findIndexPDByID(&units, *ORSub[i].Product_unit_id)
					if index >= 0 && index < len(units) {
						ORSub[i].Claim_price_lgo = units[index].PspPriceLgo
						ORSub[i].Claim_price_nhs = units[index].PspPriceNhs
						ORSub[i].Claim_price_ofc = units[index].PspPriceOfc
						ORSub[i].Claim_price_ucs = units[index].PspPriceUcs
						ORSub[i].Claim_price_ssi = units[index].PspPriceSsi
						ORSub[i].Claim_price_sss = units[index].PspPriceSss
					}
					ORSub[i].Units = &units
				}
				ORSub[i].Balance = 0
				var productStoreBalance structs.ObjQueryProductStoreBalance
				if errT := models.GetProductStoreBalance(objQueryShopStore.ID, *ORSub[i].Product_id, &productStoreBalance); errT != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Product store balance invalid.",
						"data":    errT,
					})
					return
				}
				ORSub[i].Balance = productStoreBalance.Pds_balance
			} else {
				ORSub[i].Units = &[]structs.ProductUnitList{}
			}
			if ORSub[i].Ord_is_set == 1 {
				if ORSub[i].Ord_id_set == nil {
					ORSub[i].Label = "Set"
				} else {
					if ORSub[i].Ord_id_set != nil {
						if ORSub[i].Ord_type_id == 1 {
							var queueItemChecking structs.QueueItemChecking
							if errQICI := models.GetQueueItemCheckingId(*ORSub[i].Ord_id_set, &queueItemChecking); errQICI != nil {
								c.AbortWithStatusJSON(200, gin.H{
									"status":  false,
									"message": "Cannot Queue Item Checking.",
									"data":    errQICI.Error(),
								})
								return
							}
							ORSub[i].Label = queueItemChecking.Checking_name
						} else if ORSub[i].Ord_type_id == 2 {
							var queueItemCourse structs.QueueItemCourse
							if errQIC := models.GetQueueItemCourseId(*ORSub[i].Ord_id_set, &queueItemCourse); errQIC != nil {
								c.AbortWithStatusJSON(200, gin.H{
									"status":  false,
									"message": "Cannot Queue Item Course.",
									"data":    errQIC.Error(),
								})
								return
							}
							ORSub[i].Label = queueItemCourse.Course_name
						} else if ORSub[i].Ord_type_id == 3 {
							var queueItemProduct structs.QueueItemProduct
							if errQIP := models.GetQueueItemSetProductId(*ORSub[i].Ord_id_set, &queueItemProduct); errQIP != nil {
								c.AbortWithStatusJSON(200, gin.H{
									"status":  false,
									"message": "Cannot Queue Item Product.",
									"data":    errQIP.Error(),
								})
								return
							}
							ORSub[i].Label = queueItemProduct.Pd_name
						}
					}
				}
			}
			if ORSub[i].Ord_type_id == 5 {
				var Label string
				if ORSub[i].Course_id != nil {
					var itemCourse structs.ItemCourseOrder
					if errIC := models.GetItemCourseId(*ORSub[i].Course_id, &itemCourse); errIC != nil {
						c.AbortWithStatusJSON(200, gin.H{
							"status":  false,
							"message": "Cannot Order Course.",
							"data":    errIC.Error(),
						})
						return
					}
					Label = itemCourse.Course_code + ":" + itemCourse.Course_name + " (Limit " + strconv.Itoa(int(ORSub[i].Ord_limit_qty)) + ")"
					if ORSub[i].Ord_limit_qty == 0 {
						Label = itemCourse.Course_code + ":" + itemCourse.Course_name
					}
					ORSub[i].Label = Label
				} else if ORSub[i].Checking_id != nil {
					var itemChecking structs.ItemCheckingOrder
					if errICi := models.GetItemCheckingId(*ORSub[i].Checking_id, &itemChecking); errICi != nil {
						c.AbortWithStatusJSON(200, gin.H{
							"status":  false,
							"message": "Cannot Order Checking.",
							"data":    errICi.Error(),
						})
						return
					}
					Label = itemChecking.Checking_code + ":" + itemChecking.Checking_name
					ORSub[i].Label = Label
				}
			} else {
				if ORSub[i].Course_id != nil {
					var itemCourse structs.ItemCourseOrder
					if errIC := models.GetItemCourseId(*ORSub[i].Course_id, &itemCourse); errIC != nil {
						c.AbortWithStatusJSON(200, gin.H{
							"status":  false,
							"message": "Cannot Order Course.",
							"data":    errIC.Error(),
						})
						return
					}
					ORSub[i].Claim_price_lgo = itemCourse.Course_lgo
					ORSub[i].Claim_price_nhs = itemCourse.Course_nhs
					ORSub[i].Claim_price_ofc = itemCourse.Course_ofc
					ORSub[i].Claim_price_ucs = itemCourse.Course_ucs
					ORSub[i].Claim_price_ssi = itemCourse.Course_ssi
					ORSub[i].Claim_price_sss = itemCourse.Course_sss
				} else if ORSub[i].Checking_id != nil {
					var itemChecking structs.ItemCheckingOrder
					if errICi := models.GetItemCheckingId(*ORSub[i].Checking_id, &itemChecking); errICi != nil {
						c.AbortWithStatusJSON(200, gin.H{
							"status":  false,
							"message": "Cannot Order Checking.",
							"data":    errICi.Error(),
						})
						return
					}
					ORSub[i].Claim_price_lgo = itemChecking.Checking_lgo
					ORSub[i].Claim_price_nhs = itemChecking.Checking_nhs
					ORSub[i].Claim_price_ofc = itemChecking.Checking_ofc
					ORSub[i].Claim_price_ucs = itemChecking.Checking_ucs
					ORSub[i].Claim_price_ssi = itemChecking.Checking_ssi
					ORSub[i].Claim_price_sss = itemChecking.Checking_sss
				}
			}
		}(i)
	}
	wg.Wait()

	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    ORdetail.Subs,
	})
}

func UpdateOrder(c *gin.Context) {

	var payload structs.ObjPayloadEditOrder
	if errPL := c.ShouldBindJSON(&payload); errPL != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errPL.Error(),
		})
		return
	}

	if payload.Or_is_active == 0 {
		c.JSON(200, gin.H{
			"status":  true,
			"message": "Cancel Order.",
			"data":    "",
		})
		return
	}

	var user models.User
	userId := libs.StrToInt(c.Params.ByName("userID"))
	shopId := libs.StrToInt(c.Params.ByName("shopId"))
	if errMD := models.GetUserById(userId, shopId, &user); errMD != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    "",
		})
		return
	}

	layout := "2006-01-02 15:04:05"
	time2 := time.Now().Format("15:04:05")
	ordate1, _ := time.Parse(layout, payload.Or_datetime)
	dateStr := fmt.Sprintf("%d-%02d-%02d", ordate1.Year(), ordate1.Month(), ordate1.Day())
	Ordate, errTfd := time.Parse(layout, dateStr+" "+time2)
	if errTfd != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "This Order Date.",
			"data":    errTfd.Error(),
		})
		return
	}

	var or_total_price float64
	if len(payload.Subs) > 0 {
		for _, item := range payload.Subs {
			or_total_price += item.Ord_amount
		}
	}

	POHEdit := structs.OrderDetail{
		Id:                  payload.Id,
		Shop_id:             payload.Shop_id,
		User_id:             payload.User_id,
		DpmId:               payload.DpmId,
		Or_eclaim_id:        payload.Or_eclaim_id,
		Customer_id:         payload.Customer_id,
		Queue_id:            *payload.Queue_id,
		Or_tel:              payload.Or_tel,
		Or_email:            payload.Or_email,
		Or_address:          *payload.Or_address,
		Or_district:         payload.Or_district,
		Or_amphoe:           payload.Or_amphoe,
		Or_province:         payload.Or_province,
		Or_zipcode:          payload.Or_zipcode,
		Or_comment:          payload.Or_comment,
		Or_total_price:      or_total_price,
		Or_discount_type_id: payload.Or_discount_type_id,
		Or_discount_item:    *payload.Or_discount_item,
		Or_discount_value:   *payload.Or_discount_value,
		Or_discount:         *payload.Or_discount,
		Or_befor_vat:        payload.Or_befor_vat,
		Tax_type_id:         payload.Tax_type_id,
		Tax_rate:            *payload.Tax_rate,
		Or_vat:              *payload.Or_vat,
		Or_total:            *payload.Or_total,
		Or_is_active:        payload.Or_is_active,
		Or_datetime:         Ordate.String(),
		Or_update:           time.Now().Format("2006-01-02 15:04:05"),
		Tags:                *&payload.Tags,
		Subs:                &payload.Subs,
		Or_eclaim_rate:      payload.Or_eclaim_rate,
		Or_eclaim_over:      payload.Or_eclaim_over,
		Or_eclaim_total:     payload.Or_eclaim_total,
	}

	err := models.UpdateOrder(&POHEdit)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Cannot Edit Order",
			"data":    err.Error(),
		})
		return
	}

	models.AddLogOrder(&structs.LogOrders{
		Username:   c.Params.ByName("userEmail"),
		Log_type:   "Edit Order",
		Log_text:   "Edit Order Id = " + strconv.Itoa(payload.Id),
		Log_create: time.Now().Format("2006-01-02 15:04:05"),
		Shop_id:    libs.StrToInt(c.Params.ByName("shopId")),
	})

	Noti_id, Noti_product_detail, Noti_product_name := CheckOrderTopicalProduct(shopId, payload.Id)

	c.JSON(200, gin.H{
		"status":              true,
		"message":             "Edit Order Success.",
		"data":                "",
		"noti_product":        Noti_id,
		"noti_product_detail": Noti_product_detail,
		"noti_product_name":   Noti_product_name,
	})

}

func DelOrder(c *gin.Context) {
	ORDId := libs.StrToInt(c.Params.ByName("id"))
	Shop_id := libs.StrToInt(c.Params.ByName("shopId"))
	var ORD structs.OrderDetail
	if errCK := models.GetOrderDetail(ORDId, &ORD); errCK != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Order invalid.",
			"data":    errCK.Error(),
		})
		return
	}

	if ORD.Id < 1 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Order invalid.",
			"data":    ORD,
		})
		return
	}
	userId := libs.StrToInt(c.Params.ByName("userID"))
	if errAD := models.CancelOrder(ORDId, ORD.Queue_id, userId, &ORD); errAD != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Cannot Remove Order",
			"data":    errAD.Error(),
		})
		return
	}
	models.AddLogOrder(&structs.LogOrders{
		Username:   c.Params.ByName("userEmail"),
		Log_type:   "Cancel Order",
		Log_text:   "Cancel Order Id = " + strconv.Itoa(ORDId),
		Log_create: time.Now().Format("2006-01-02 15:04:05"),
		Shop_id:    libs.StrToInt(c.Params.ByName("shopId")),
	})

	//Labplus
	var ResponseCancels *libs.ResponseCancel
	if ORD.Queue_id != 0 {
		Labplus := structs.Labplus{}
		if errLabplus := models.GetShopLabplus(Shop_id, &Labplus); errLabplus != nil {
			c.JSON(200, gin.H{
				"status":  false,
				"message": "Cannot Queue Checking.",
				"data":    errLabplus.Error(),
			})
			return
		}
		if Labplus.Id != 0 {
			var QLabplus structs.QueuesLabplus
			if errQLabplus := models.GetQueueIDLabplus(ORD.Queue_id, &QLabplus); errQLabplus != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Queue Labplus invalid.",
					"data":    errQLabplus.Error(),
				})
				return
			}
			if QLabplus.Id != 0 {
				token := LabplusAuthen(Shop_id)
				if token != "token" {
					ResponseCancel, errResponseCancel := libs.LabpluscancelOrder(Labplus.Lapi_link, token, QLabplus.Qlp_no)
					if errResponseCancel != nil {
						c.AbortWithStatusJSON(200, gin.H{
							"status":  false,
							"message": "Err ResponseCancel.",
							"data":    errResponseCancel,
						})
						return
					}
					ResponseCancels = ResponseCancel
				}
				QueuesLabplus := structs.QueuesLabplus{}
				models.DeleteQueuesLabplus(ORD.Queue_id, &QueuesLabplus)
			}
		}
	}

	c.JSON(200, gin.H{
		"status":          true,
		"message":         "Cancel Order ID " + strconv.Itoa(ORDId) + " success",
		"data":            "",
		"ResponseCancels": ResponseCancels,
	})

}

func DelOrderSub(c *gin.Context) {
	ORSId := libs.StrToInt(c.Params.ByName("id"))
	var ORDS structs.OrderSub
	if errCK := models.GetOrderSubId(ORSId, &ORDS); errCK != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Order Sub invalid.",
			"data":    errCK.Error(),
		})
		return
	}

	if ORDS.Id == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Order Sub invalid.",
			"data":    ORDS,
		})
		return
	}

	var Ordsub structs.OrderSub
	if errAD := models.CancelOrderSub(ORSId, &Ordsub); errAD != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Cannot Remove Order Sub",
			"data":    errAD.Error(),
		})
		return
	}
	models.AddLogOrder(&structs.LogOrders{
		Username:   c.Params.ByName("userEmail"),
		Log_type:   "Cancel Order Sub",
		Log_text:   "Cancel Order Sub Id = " + strconv.Itoa(ORSId),
		Log_create: time.Now().Format("2006-01-02 15:04:05"),
		Shop_id:    libs.StrToInt(c.Params.ByName("shopId")),
	})

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Cancel Order Sub ID " + strconv.Itoa(ORSId) + " success",
		"data":    "",
	})

}

func DelOrderTag(c *gin.Context) {
	TGId := libs.StrToInt(c.Params.ByName("id"))
	var ORT structs.OrderTags
	if errCK := models.GetOrderTagId(TGId, &ORT); errCK != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Order Sub invalid.",
			"data":    errCK.Error(),
		})
		return
	}

	if ORT.Id == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Order Sub invalid.",
			"data":    ORT,
		})
		return
	}

	var Ordsub structs.OrderSub
	if errAD := models.CancelOrderTag(TGId, &Ordsub); errAD != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Cannot Remove Order Tag",
			"data":    errAD.Error(),
		})
		return
	}
	models.AddLogOrder(&structs.LogOrders{
		Username:   c.Params.ByName("userEmail"),
		Log_type:   "Cancel Order Tag",
		Log_text:   "Cancel Order Tag Id = " + strconv.Itoa(TGId),
		Log_create: time.Now().Format("2006-01-02 15:04:05"),
		Shop_id:    libs.StrToInt(c.Params.ByName("shopId")),
	})

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Cancel Order Tag ID " + strconv.Itoa(TGId) + " success",
		"data":    "",
	})

}

func CheckOrderTopical(c *gin.Context) {
	Id := libs.StrToInt(c.Params.ByName("id"))
	shopId := libs.StrToInt(c.Params.ByName("shopId"))
	var OR structs.Order
	if errOR := models.GetOrderByID(Id, &OR); errOR != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Order Sub invalid.",
			"data":    errOR.Error(),
		})
		return
	}
	var ItemList = structs.OrderTopicalNotiProduct{}
	if OR.Or_is_active == 1 {
		var OT []structs.OrderTopical
		if err := models.GetOrderTopical(shopId, &OT); err != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Order Sub invalid.",
				"data":    err.Error(),
			})
			return
		}

		for _, data := range OT {
			var topical []structs.OrderTopical
			errCT := models.CheckTopical(data.Id, &topical)
			if errCT != nil {
				c.JSON(200, gin.H{
					"status":  false,
					"message": "Appoint Tag Invalid!",
					"data":    errCT.Error(),
				})
				return
			}
			if len(topical) > 0 {
				var orederTopical []structs.OrderTopical
				errCOT := models.CheckOrderTopical(Id, data.Id, &orederTopical)
				if errCOT != nil {
					c.JSON(200, gin.H{
						"status":  false,
						"message": "Appoint Tag Invalid!",
						"data":    errCOT.Error(),
					})
					return
				}
				if len(topical) == len(orederTopical) {
					ItemList.Id = data.Id
					ItemList.Noti_product_detail = data.Topical_detail
					ItemList.Noti_product_name = data.Topical_name
				}
			}
		}
	}
	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    ItemList,
	})
}

func CheckOrderTopicalProduct(shopId int, Id int) (int, string, string) {
	var OT []structs.OrderTopical
	models.GetOrderTopical(shopId, &OT)
	var ItemList = structs.OrderTopicalNotiProduct{}
	for _, data := range OT {
		var topical []structs.OrderTopical
		models.CheckTopical(data.Id, &topical)
		if len(topical) > 0 {
			var orederTopical []structs.OrderTopical
			models.CheckOrderTopical(Id, data.Id, &orederTopical)
			if len(topical) == len(orederTopical) {
				ItemList.Id = 1
				ItemList.Noti_product_detail = data.Topical_detail
				ItemList.Noti_product_name = data.Topical_name
			}
		}
	}
	return ItemList.Id, ItemList.Noti_product_detail, ItemList.Noti_product_name
}

func CheckOrderTopicalProductItem(shopId int, Id int, Pd_id int) (int, string, string) {
	var CkPd int = 0
	var OT []structs.OrderTopical
	models.GetOrderTopical(shopId, &OT)
	var ItemList = structs.OrderTopicalNotiProduct{}
	for _, data := range OT {
		var topical []structs.OrderTopical
		models.CheckTopical(data.Id, &topical)
		if len(topical) > 0 {
			var orederTopical []structs.OrderTopical
			models.CheckOrderTopical(Id, data.Id, &orederTopical)
			if len(topical) == len(orederTopical) {
				for _, dataP := range orederTopical {
					if dataP.Id == Pd_id {
						CkPd = 1
					}
				}
				if CkPd == 1 {
					ItemList.Id = 1
					ItemList.Noti_product_detail = data.Topical_detail
					ItemList.Noti_product_name = data.Topical_name
				}
			}
		}
	}
	return ItemList.Id, ItemList.Noti_product_detail, ItemList.Noti_product_name
}

func AddOrderQueueIPD(c *gin.Context) {
	var payload structs.ObjPayloadAddOrderQueueIPD
	if errPL := c.ShouldBindJSON(&payload); errPL != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errPL.Error(),
		})
		return
	}

	Or_tele_code := ""

	var Queue structs.CheckQueueStatusID
	if errQueue := models.CheckQueueStatusID(*payload.Queue_id, &Queue); errQueue != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    "",
		})
		return
	}

	if Queue.Que_status_id > 2 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Queue is active invalid. Please refresh.",
			"data":    "",
		})
		return
	}

	Or_tele_code = Queue.Que_tele_code

	var user models.User
	userId := libs.StrToInt(c.Params.ByName("userID"))
	Shop_id := libs.StrToInt(c.Params.ByName("shopId"))
	if errMD := models.GetUserById(userId, Shop_id, &user); errMD != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    "",
		})
		return
	}

	layout := "2006-01-02 15:04:05"
	Ordate, errTfd := time.Parse(layout, payload.Or_datetime)
	if errTfd != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "This Order Date.",
			"data":    errTfd.Error(),
		})
		return
	}

	POHAdd := structs.OrderDetail{
		Id:             0,
		Shop_id:        Shop_id,
		User_id:        userId,
		Customer_id:    payload.Customer_id,
		Queue_id:       *payload.Queue_id,
		Or_fullname:    payload.Or_fullname,
		Or_tel:         payload.Or_tel,
		Or_email:       payload.Or_email,
		Or_address:     *payload.Or_address,
		Or_district:    payload.Or_district,
		Or_amphoe:      payload.Or_amphoe,
		Or_province:    payload.Or_province,
		Or_zipcode:     payload.Or_zipcode,
		Or_comment:     payload.Or_comment,
		Or_total_price: *payload.Or_total_price,
		Or_discount:    *payload.Or_discount,
		Or_befor_vat:   payload.Or_befor_vat,
		Tax_type_id:    payload.Tax_type_id,
		Tax_rate:       *payload.Tax_rate,
		Or_vat:         *payload.Or_vat,
		Or_total:       *payload.Or_total,
		Or_is_active:   1,
		Or_datetime:    Ordate.String(),
		Or_create:      time.Now().Format("2006-01-02 15:04:05"),
		Or_update:      time.Now().Format("2006-01-02 15:04:05"),
		Tags:           *&payload.Tags,
		Subs:           &payload.Subs,
		Or_tele_code:   Or_tele_code,
	}

	orderId, err := models.AddOrderIPD(&POHAdd)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Cannot Add Order",
			"data":    err.Error(),
		})
		return
	}

	if orderId == 0 {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Cannot Add Order",
			"data":    "",
		})
		return
	}

	models.AddLogOrder(&structs.LogOrders{
		Username:   c.Params.ByName("userEmail"),
		Log_type:   "Add Order",
		Log_text:   "Add Order Id = " + strconv.Itoa(orderId),
		Log_create: time.Now().Format("2006-01-02 15:04:05"),
		Shop_id:    libs.StrToInt(c.Params.ByName("shopId")),
	})
	var message string = "Add order success."
	//QueueIPD
	if payload.Queue_ipd.Queue_id != -1 {
		var getQueue structs.GetQueueId
		if errGQ := models.GetQueueIPDId(payload.Queue_ipd.Queue_id, &getQueue); errGQ != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Queue invalid.",
				"data":    errGQ.Error(),
			})
			return
		}

		if getQueue.Que_status_id != 3 {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Queue is active invalid. Please refresh.",
				"data":    "",
			})
			return
		}
		var QueRefId int
		var QueRefIpd int
		var CountQ int
		var queCode string
		if getQueue.Que_ref_ipd > 0 {
			var queueCount structs.QueueMax
			if errCQ := models.GetCountQueue(getQueue.Que_ref_ipd, &queueCount); errCQ != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Count queue invalid.",
					"data":    errCQ.Error(),
				})
				return
			}
			QueRefId = getQueue.ID
			QueRefIpd = getQueue.Que_ref_ipd
			CountQ = queueCount.Id + 1
			var getQueueCode structs.GetQueueId
			if errGQM := models.GetQueueIPDId(getQueue.Que_ref_ipd, &getQueueCode); errGQM != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Queue invalid.",
					"data":    errGQM.Error(),
				})
				return
			}
			queCode = getQueueCode.Que_code + "-" + strconv.Itoa(CountQ)
		} else {
			var queueCount structs.QueueMax
			if errCQ := models.GetCountQueue(payload.Queue_ipd.Queue_id, &queueCount); errCQ != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Count queue invalid.",
					"data":    errCQ.Error(),
				})
				return
			}
			QueRefId = getQueue.ID
			QueRefIpd = getQueue.ID
			CountQ = queueCount.Id + 1
			queCode = getQueue.Que_code + "-" + strconv.Itoa(CountQ)
		}

		objQueryCreateQueue := structs.AddMoveQueueIPD{
			ShopId:          Shop_id,
			CustomerId:      payload.Queue_ipd.CustomerId,
			RoomId:          payload.Queue_ipd.RoomId,
			BedId:           payload.Queue_ipd.BedId,
			UserId:          userId,
			QueUserId:       payload.Queue_ipd.DoctorId,
			QueUserFullname: payload.Queue_ipd.DoctorFullname,
			QueCode:         queCode,
			QueTypeId:       1,
			QueAdmisId:      1,
			QueRefIpd:       QueRefIpd,
			QueRefId:        QueRefId,
			QuePriorityId:   payload.Queue_ipd.QuePriorityId,
			QueStatusId:     1,
			QueDatetime:     payload.Queue_ipd.QueDatetime,
			QueUpdate:       time.Now().Format("2006-01-02 15:04:05"),
			QueCreate:       time.Now().Format("2006-01-02 15:04:05"),
		}

		errCQIPD := models.CreateQueueIPD(&objQueryCreateQueue)
		if errCQIPD != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Create queue error.",
				"data":    "",
			})
			return
		}

		//Get QueueIPD detail
		Queue_id := payload.Queue_ipd.Queue_id
		var ItemListChecking []structs.QueueChecking
		if errMD := models.GetQueueCheckingIPD(Queue_id, &ItemListChecking); errMD != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Cannot Queue Checking.",
				"data":    errMD.Error(),
			})
			return
		}

		if len(ItemListChecking) > 0 {
			for _, qChecking := range ItemListChecking {
				queuechecking := structs.QueueCheckingIPD{
					Queue_id:        objQueryCreateQueue.ID,
					User_id:         userId,
					Checking_id:     qChecking.Checking_id,
					Queci_code:      qChecking.Queci_code,
					Queci_name:      qChecking.Queci_name,
					Queci_qty:       qChecking.Queci_qty,
					Queci_unit:      qChecking.Queci_unit,
					Queci_cost:      qChecking.Queci_cost,
					Queci_price:     qChecking.Queci_price,
					Queci_discount:  qChecking.Queci_discount,
					Queci_total:     qChecking.Queci_total,
					Queci_is_set:    qChecking.Queci_is_set,
					Queci_id_ref:    *qChecking.Id,
					Queci_ipd_order: 1,
					Queci_is_active: qChecking.Queci_is_active,
					Queci_modify:    time.Now().Format("2006-01-02 15:04:05"),
				}

				errACI := models.CreateQueueItemCheckingIPD(&queuechecking)
				if errACI != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Cannot Add Queue Checking.",
						"data":    errACI.Error(),
					})
					return
				}
				errUCI := models.UpdateQueueCheckingIPD(*qChecking.Id)
				if errUCI != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Cannot Update Queue Checking.",
						"data":    errUCI.Error(),
					})
					return
				}
			}
		}

		var ItemListCourse []structs.QueueCourse
		if errMD := models.GetQueueCourseIPD(Queue_id, &ItemListCourse); errMD != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Cannot Queue Course.",
				"data":    errMD.Error(),
			})
			return
		}

		if len(ItemListCourse) > 0 {
			for _, qCourse := range ItemListCourse {
				queuecourseset := structs.QueueCourseIPD{
					Queue_id:       objQueryCreateQueue.ID,
					User_id:        userId,
					Course_id:      qCourse.Course_id,
					Quec_code:      qCourse.Quec_code,
					Quec_name:      qCourse.Quec_name,
					Quec_qty:       qCourse.Quec_qty,
					Quec_unit:      qCourse.Quec_unit,
					Quec_cost:      qCourse.Quec_cost,
					Quec_price:     qCourse.Quec_price,
					Quec_discount:  qCourse.Quec_discount,
					Quec_total:     qCourse.Quec_total,
					Quec_is_set:    qCourse.Quec_is_set,
					Quec_is_active: qCourse.Quec_is_active,
					Quec_id_ref:    *qCourse.Id,
					Quec_ipd_order: 1,
					Quec_modify:    time.Now().Format("2006-01-02 15:04:05"),
				}
				errAC := models.CreateQueueItemCourseIPD(&queuecourseset)
				if errAC != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Cannot Add Queue Course",
						"data":    errAC.Error(),
					})
					return
				}
				errUC := models.UpdateQueueCourseIPD(*qCourse.Id)
				if errUC != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Cannot Update Queue Course",
						"data":    errUC.Error(),
					})
					return
				}
			}
		}

		var ItemListProduct []structs.QueueProduct
		if errMD := models.GetQueueProductIPD(Queue_id, &ItemListProduct); errMD != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Cannot Queue Product.",
				"data":    errMD.Error(),
			})
			return
		}

		if len(ItemListProduct) > 0 {
			for _, qProduct := range ItemListProduct {
				if qProduct.Quep_type_id == 1 {
					queueproduct := structs.QueueProductIPD{
						Queue_id:          objQueryCreateQueue.ID,
						User_id:           userId,
						Product_id:        qProduct.Product_id,
						Product_store_id:  qProduct.Product_store_id,
						Product_unit_id:   qProduct.Product_unit_id,
						Queue_checking_id: qProduct.Queue_checking_id,
						Checking_id:       qProduct.Checking_id,
						Quep_type_id:      qProduct.Quep_type_id,
						Quep_code:         qProduct.Quep_code,
						Quep_name:         qProduct.Quep_name,
						Quep_qty:          qProduct.Quep_qty,
						Quep_set_qty:      qProduct.Quep_set_qty,
						Quep_limit_qty:    qProduct.Quep_limit_qty,
						Quep_unit:         qProduct.Quep_unit,
						Quep_cost:         qProduct.Quep_cost,
						Quep_price:        qProduct.Quep_price,
						Quep_discount:     qProduct.Quep_discount,
						Quep_total:        qProduct.Quep_total,
						Topical_id:        qProduct.Topical_id,
						Quep_topical:      qProduct.Quep_topical,
						Quep_direction:    qProduct.Quep_direction,
						Quep_is_set:       qProduct.Quep_is_set,
						Quep_id_ref:       *qProduct.Id,
						Quep_ipd_order:    1,
						Quep_is_active:    qProduct.Quep_is_active,
						Quep_modify:       time.Now().Format("2006-01-02 15:04:05"),
					}
					errAP := models.CreateQueueItemProductIPD(&queueproduct)
					if errAP != nil {
						c.JSON(200, gin.H{
							"status":  false,
							"message": "Cannot Add Queue Product.",
							"data":    errAP.Error(),
						})
						return
					}
					errUP := models.UpdateQueueProductIPD(*qProduct.Id)
					if errUP != nil {
						c.JSON(200, gin.H{
							"status":  false,
							"message": "Cannot Update Queue Product.",
							"data":    errUP.Error(),
						})
						return
					}
				} else {
					if qProduct.Queue_checking_id != nil {
						var ItemListCheckingRef structs.QueueCheckingIPD
						if errMDCI := models.GetQueueCheckingIPDRef(*qProduct.Queue_checking_id, &ItemListCheckingRef); errMDCI != nil {
							c.AbortWithStatusJSON(200, gin.H{
								"status":  false,
								"message": "Cannot Queue Checking Ref.",
								"data":    errMDCI.Error(),
							})
							return
						}
						queueproduct := structs.QueueProductIPD{
							Queue_id:          objQueryCreateQueue.ID,
							User_id:           userId,
							Product_id:        qProduct.Product_id,
							Product_store_id:  qProduct.Product_store_id,
							Product_unit_id:   qProduct.Product_unit_id,
							Queue_checking_id: ItemListCheckingRef.Id,
							Checking_id:       &ItemListCheckingRef.Checking_id,
							Quep_type_id:      qProduct.Quep_type_id,
							Quep_code:         qProduct.Quep_code,
							Quep_name:         qProduct.Quep_name,
							Quep_qty:          qProduct.Quep_qty,
							Quep_set_qty:      qProduct.Quep_set_qty,
							Quep_limit_qty:    qProduct.Quep_limit_qty,
							Quep_unit:         qProduct.Quep_unit,
							Quep_cost:         qProduct.Quep_cost,
							Quep_price:        qProduct.Quep_price,
							Quep_discount:     qProduct.Quep_discount,
							Quep_total:        qProduct.Quep_total,
							Topical_id:        qProduct.Topical_id,
							Quep_topical:      qProduct.Quep_topical,
							Quep_direction:    qProduct.Quep_direction,
							Quep_is_set:       qProduct.Quep_is_set,
							Quep_id_ref:       *qProduct.Id,
							Quep_ipd_order:    1,
							Quep_is_active:    qProduct.Quep_is_active,
							Quep_modify:       time.Now().Format("2006-01-02 15:04:05"),
						}
						errAP := models.CreateQueueItemProductIPD(&queueproduct)
						if errAP != nil {
							c.JSON(200, gin.H{
								"status":  false,
								"message": "Cannot Add Checking Ref Product.",
								"data":    errAP.Error(),
							})
							return
						}
						errUP := models.UpdateQueueProductIPD(*qProduct.Id)
						if errUP != nil {
							c.JSON(200, gin.H{
								"status":  false,
								"message": "Cannot Update Checking Ref Product.",
								"data":    errUP.Error(),
							})
							return
						}
					} else if qProduct.Queue_course_id != nil {
						var ItemListCourseRef structs.QueueCourseIPD
						if errMDC := models.GetQueueCourseIPDRef(*qProduct.Queue_course_id, &ItemListCourseRef); errMDC != nil {
							c.AbortWithStatusJSON(200, gin.H{
								"status":  false,
								"message": "Cannot Queue Course Ref.",
								"data":    errMDC.Error(),
							})
							return
						}
						queueproduct := structs.QueueProductIPD{
							Queue_id:         objQueryCreateQueue.ID,
							User_id:          userId,
							Product_id:       qProduct.Product_id,
							Product_store_id: qProduct.Product_store_id,
							Product_unit_id:  qProduct.Product_unit_id,
							Queue_course_id:  ItemListCourseRef.Id,
							Course_id:        &ItemListCourseRef.Course_id,
							Quep_type_id:     qProduct.Quep_type_id,
							Quep_code:        qProduct.Quep_code,
							Quep_name:        qProduct.Quep_name,
							Quep_qty:         qProduct.Quep_qty,
							Quep_set_qty:     qProduct.Quep_set_qty,
							Quep_limit_qty:   qProduct.Quep_limit_qty,
							Quep_unit:        qProduct.Quep_unit,
							Quep_cost:        qProduct.Quep_cost,
							Quep_price:       qProduct.Quep_price,
							Quep_discount:    qProduct.Quep_discount,
							Quep_total:       qProduct.Quep_total,
							Topical_id:       qProduct.Topical_id,
							Quep_topical:     qProduct.Quep_topical,
							Quep_direction:   qProduct.Quep_direction,
							Quep_is_set:      qProduct.Quep_is_set,
							Quep_id_ref:      *qProduct.Id,
							Quep_ipd_order:   1,
							Quep_is_active:   qProduct.Quep_is_active,
							Quep_modify:      time.Now().Format("2006-01-02 15:04:05"),
						}
						errAP := models.CreateQueueItemProductIPD(&queueproduct)
						if errAP != nil {
							c.JSON(200, gin.H{
								"status":  false,
								"message": "Cannot Add Course Ref Product.",
								"data":    errAP.Error(),
							})
							return
						}
						errUP := models.UpdateQueueProductIPD(*qProduct.Id)
						if errUP != nil {
							c.JSON(200, gin.H{
								"status":  false,
								"message": "Cannot Update Course Ref Product.",
								"data":    errUP.Error(),
							})
							return
						}
					}
				}
			}
		}
		message = "Add order & Move queue IPD success."
	}

	var LabplusRequestOrders structs.LabplusRequestOrder
	var RequestOrders *libs.RequestOrder
	if *payload.Queue_id != 0 {
		//CheckQueueStatusID
		var Queuefullname structs.CheckQueueStatusID
		if errQueuefullname := models.CheckQueueStatusID(*payload.Queue_id, &Queuefullname); errQueuefullname != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Err Queue fullname.",
				"data":    "",
			})
			return
		}
		//Labplus
		var CheckingLabplus []structs.CheckingLabplus
		if errCheckingLabplus := models.CheckingLabplus(*payload.Queue_id, &CheckingLabplus); errCheckingLabplus != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Something went wrong.",
				"data":    "",
			})
			return
		}
		if len(CheckingLabplus) > 0 {
			Labplus := structs.Labplus{}
			if errLabplus := models.GetShopLabplus(Shop_id, &Labplus); errLabplus != nil {
				c.JSON(200, gin.H{
					"status":  false,
					"message": "Cannot Queue Checking.",
					"data":    errLabplus.Error(),
				})
				return
			}
			if Labplus.Id != 0 {
				token := LabplusAuthen(Shop_id)
				if token != "token" {
					var CustomerPatient structs.CustomerPatient
					if errCustomerPatient := models.GetCustomerLabplus(payload.Customer_id, &CustomerPatient); errCustomerPatient != nil {
						c.AbortWithStatusJSON(200, gin.H{
							"status":  false,
							"message": "Something went wrong.",
							"data":    "",
						})
						return
					}
					Ctm_birthdate, _ := time.Parse(time.RFC3339, CustomerPatient.Ctm_birthdate)
					var Ctm_gender string = "U"
					if CustomerPatient.Ctm_gender == "หญิง" {
						Ctm_gender = "F"
					} else if CustomerPatient.Ctm_gender == "Female" {
						Ctm_gender = "F"
					} else if CustomerPatient.Ctm_gender == "ชาย" {
						Ctm_gender = "M"
					} else if CustomerPatient.Ctm_gender == "Male" {
						Ctm_gender = "M"
					}
					Patient := structs.Patient{
						HN:          CustomerPatient.Ctm_id,
						FullName:    CustomerPatient.Ctm_prefix + " " + CustomerPatient.Ctm_fname + " " + CustomerPatient.Ctm_lname,
						TName:       CustomerPatient.Ctm_prefix,
						FName:       CustomerPatient.Ctm_fname,
						LName:       CustomerPatient.Ctm_lname,
						IDCard:      CustomerPatient.Ctm_citizen_id,
						Birthday:    Ctm_birthdate.Format("2006-01-02"),
						Sex:         Ctm_gender,
						DoctorName:  Queuefullname.Que_user_fullname,
						RequestNote: strconv.Itoa(*payload.Queue_id),
					}
					var LabOrders = []structs.LabOrder{}
					for _, dataCheckingLabplus := range CheckingLabplus {
						LabOrder := structs.LabOrder{
							Code: dataCheckingLabplus.Checking_code,
							Name: dataCheckingLabplus.Checking_name,
						}
						LabOrders = append(LabOrders, LabOrder)
					}
					LabplusRequestOrder := structs.LabplusRequestOrder{
						Patient:  Patient,
						LabOrder: LabOrders,
					}
					LabplusRequestOrders = LabplusRequestOrder
					RequestOrder, errRequestOrder := libs.LabplusRequestOrder(Labplus.Lapi_link, token, &LabplusRequestOrder)
					if errRequestOrder != nil {
						c.AbortWithStatusJSON(200, gin.H{
							"status":  false,
							"message": "Err RequestOrder.",
							"data":    errRequestOrder,
						})
						return
					}
					RequestOrders = RequestOrder
					if RequestOrder.Status_code == "201" {
						qlp_no := RequestOrder.Request_no
						AddQueuesLabplus := structs.AddQueuesLabplus{
							Shop_id:          Shop_id,
							Queue_id:         *payload.Queue_id,
							Qlp_no:           qlp_no,
							Qlp_message_th:   RequestOrder.Message_th,
							Qlp_process_code: "1",
							Qlp_process_name: "สถานะ: สั่งตรวจ",
							Qlp_datetime:     time.Now().Format("2006-01-02 15:04:05"),
							Qlp_update:       time.Now().Format("2006-01-02 15:04:05"),
						}
						models.AddQueuesLabplus(&AddQueuesLabplus)
					}
				}
			}
		}
	}

	c.JSON(200, gin.H{
		"status":               true,
		"message":              message,
		"data":                 "",
		"LabplusRequestOrders": LabplusRequestOrders,
		"RequestOrders":        RequestOrders,
	})
}

func findIndexPDByID(units *[]structs.ProductUnitList, id int) int {
	for i, units := range *units {
		if units.Product_units_id == id {
			return i
		}
	}
	return -1 // Return -1 if the ID is not found
}

func CreateQueue(objPayload *structs.ObjPayloadCreateQueueByOrder) int {
	var CheckError int = 0
	var objQueryQueueDocSetting structs.ObjQueryQueueDocSetting

	errGetQueueDocSetting := models.GetQueueDocSetting(objPayload.ShopId, &objQueryQueueDocSetting)
	if errGetQueueDocSetting != nil {
		CheckError = 0
	}
	maxQueue, _ := models.GetMaxQueueNumber(objPayload)
	if maxQueue > 9999 {
		maxQueue = 1
	} else {
		maxQueue++
	}

	queCode := ""
	queDpmCode := ""
	docKey := 0
	docValue := 0
	if objPayload.QueTypeId == 1 {
		if objPayload.QueAdmisId == 1 {
			queCode = libs.SetDocSettingCode(objQueryQueueDocSetting.IpdIdDefault, objQueryQueueDocSetting.IpdNumberDigit, objQueryQueueDocSetting.IpdNumberDefault, objQueryQueueDocSetting.IpdType)
			docKey = 1
			docValue = objQueryQueueDocSetting.IpdNumberDefault + 1
			queDpmCode = libs.SetDocSettingCode(objQueryQueueDocSetting.IpdIdDefault, objQueryQueueDocSetting.IpdNumberDigit, maxQueue, 1)
		} else {
			queCode = libs.SetDocSettingCode(objQueryQueueDocSetting.OpdIdDefault, objQueryQueueDocSetting.OpdNumberDigit, objQueryQueueDocSetting.OpdNumberDefault, objQueryQueueDocSetting.OpdType)
			docKey = 2
			docValue = objQueryQueueDocSetting.OpdNumberDefault + 1
			queDpmCode = libs.SetDocSettingCode(objQueryQueueDocSetting.OpdIdDefault, objQueryQueueDocSetting.OpdNumberDigit, maxQueue, 1)
		}
	} else {
		queCode = libs.SetDocSettingCode(objQueryQueueDocSetting.ServeIdDefault, objQueryQueueDocSetting.ServeNumberDigit, objQueryQueueDocSetting.ServeNumberDefault, objQueryQueueDocSetting.ServeType)
		docKey = 3
		docValue = objQueryQueueDocSetting.ServeNumberDefault + 1
		queDpmCode = libs.SetDocSettingCode(objQueryQueueDocSetting.ServeIdDefault, objQueryQueueDocSetting.ServeNumberDigit, maxQueue, 1)
	}

	objQueryCreateQueue := structs.QueueByOrder{
		ShopId:          objPayload.ShopId,
		CustomerId:      objPayload.CustomerId,
		RoomId:          objPayload.RoomId,
		BedId:           objPayload.BedId,
		UserId:          objPayload.QueUserId,
		QueUserId:       objPayload.DoctorId,
		QueUserFullname: objPayload.DoctorFullname,
		QueCode:         queCode,
		QueTypeId:       objPayload.QueTypeId,
		QueAdmisId:      objPayload.QueAdmisId,
		QuePriorityId:   objPayload.QuePriorityId,
		QueStatusId:     3,
		QueDatetime:     objPayload.QueDatetime,
		QueTeleCode:     objPayload.QueTeleCode,
		QueTeleUrl:      objPayload.QueTeleUrl,
		QueUpdate:       time.Now().Format("2006-01-02 15:04:05"),
		QueCreate:       time.Now().Format("2006-01-02 15:04:05"),
		DpmId:           nil,
		QueNumber:       maxQueue,
		QueDpmCode:      queDpmCode,
	}
	err := models.CreateQueueByOrder(&objQueryCreateQueue)
	if err != nil {
		CheckError = 0
	}
	queueId := objQueryCreateQueue.ID

	var objDocSetting models.DocSetting
	errUpdateQueueDocSetting := models.UpdateQueueDocSetting(objPayload.ShopId, docKey, docValue, &objDocSetting)
	if errUpdateQueueDocSetting != nil {
		CheckError = 0
	}

	if CheckError == 1 {
		return 0
	} else {
		return queueId
	}
}
