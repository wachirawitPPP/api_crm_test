package controllers

import (
	"fmt"
	"linecrmapi/libs"
	"linecrmapi/models"
	"linecrmapi/structs"
	"time"

	"github.com/gin-gonic/gin"
)

func AddProcess(c *gin.Context) {
	var payload structs.ObjPayloadProcessOrder
	if errPL := c.ShouldBindJSON(&payload); errPL != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errPL.Error(),
		})
		return
	}

	receipt_id := payload.Receipt_id
	var Rec structs.ProcessReceipt
	if errRec := models.GetProcessReceiptId(receipt_id, &Rec); errRec != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Receipt invalid.",
			"data":    errRec.Error(),
		})
		return
	}

	userId := Rec.User_id
	Shop_id := Rec.Shop_id
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

	var Ctm structs.CustomerIdProduct
	if errCtm := models.GetProcessCustomerId(Rec.Customer_id, &Ctm); errCtm != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Customer invalid.",
			"data":    errCtm.Error(),
		})
		return
	}

	if Rec.Rec_is_process == 1 {
		//สร้างรายการตรวจ
		var Recd []structs.ProcessReceiptDetail
		if errRecd := models.GetProcessReceiptDetailCheck(receipt_id, &Recd); errRecd != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Receipt detail checking invalid.",
				"data":    errRecd.Error(),
			})
			return
		}
		for _, Recdcheck := range Recd {
			var checking structs.ProcessCheckChecking
			if errChecking := models.GetProcessCheckingChecksId(*Recdcheck.Checking_id, &checking); errChecking != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Checking invalid.",
					"data":    errChecking.Error(),
				})
				return
			}
			if checking.Checking_type_id != 4 {
				var i float64
				for i = 0; i < Recdcheck.Recd_qty; i++ {
					var processCheck structs.ProcessCheck
					if errCCheck := models.GetProcessCheck(Recdcheck.Invoice_detail_id, *Recdcheck.Checking_id, &processCheck); errCCheck != nil {
						c.AbortWithStatusJSON(200, gin.H{
							"status":  false,
							"message": "Receipt detail check invalid.",
							"data":    errCCheck.Error(),
						})
						return
					}
					if errUID := models.UpdateCheckInvoice(processCheck.Id, receipt_id, Recdcheck.Id); errUID != nil {
						c.AbortWithStatusJSON(200, gin.H{
							"status":  false,
							"message": "Cannot Update Check",
							"data":    errUID.Error(),
						})
						return
					}
					//รายการยาที่ใช้
					var Recdcip []structs.ProcessReceiptDetail
					if errCheckingPro := models.GetProcessReceiptDetailCheckProduct(receipt_id, *Recdcheck.Checking_id, &Recdcip); errCheckingPro != nil {
						c.AbortWithStatusJSON(200, gin.H{
							"status":  false,
							"message": "Checking product invalid.",
							"data":    errCheckingPro.Error(),
						})
						return
					}
					for _, Recdcheckproduct := range Recdcip {
						var Check_product_id *int
						if processCheck.Id > 0 {
							Rec_pay_datetime, _ := time.Parse(time.RFC3339, Rec.Rec_pay_datetime)
							processCheckproduct := structs.ProcessCheckProduct{
								Id:               0,
								Shop_id:          Rec.Shop_id,
								Check_id:         processCheck.Id,
								Checking_id:      *Recdcheckproduct.Checking_id,
								Queue_id:         Recdcheckproduct.Queue_id,
								Customer_id:      Rec.Customer_id,
								Product_id:       *Recdcheckproduct.Product_id,
								Product_store_id: *Recdcheckproduct.Product_store_id,
								Product_unit_id:  *Recdcheckproduct.Product_unit_id,
								Receipt_id:       receipt_id,
								Chkp_code:        Recdcheckproduct.Recd_code,
								Chkp_name:        Recdcheckproduct.Recd_name,
								Chkp_qty:         Recdcheckproduct.Recd_qty,
								Chkp_unit:        Recdcheckproduct.Recd_unit,
								Chkp_is_active:   1,
								Chkp_datetime:    Rec_pay_datetime.Format("2006-01-02 15:04:05"),
								Chkp_create:      time.Now().Format("2006-01-02 15:04:05"),
								Chkp_modify:      time.Now().Format("2006-01-02 15:04:05"),
							}
							errCCheck := models.CreateChecksProduct(&processCheckproduct)
							if errCCheck != nil {
								c.AbortWithStatusJSON(200, gin.H{
									"status":  false,
									"message": "Create Check Product error.",
									"data":    "",
								})
								return
							}
							Check_product_id = &processCheckproduct.Id
						}
						var Receipt_id = receipt_id
						var Receipt_detail_id = Recdcheckproduct.Id
						processProductUsed := structs.ProcessProduct{
							Shop_id:                 Rec.Shop_id,
							Customer_id:             Rec.Customer_id,
							User_id:                 userId,
							Product_id:              *Recdcheckproduct.Product_id,
							Product_store_id:        *Recdcheckproduct.Product_store_id,
							Receipt_id:              &Receipt_id,
							Receipt_detail_id:       &Receipt_detail_id,
							Check_product_id:        Check_product_id,
							Service_product_used_id: nil,
							Queue_id:                Recdcheckproduct.Queue_id,
							Pdsh_out_id:             4, // 1 ขาย, 2 โอน, 3 ใช้บริการ, 4 ตรวจ
							Pdsh_ref_doc_no:         Rec.Rec_code,
							Pdso_qty:                Recdcheckproduct.Recd_qty,
							Recd_code:               Recdcheckproduct.Recd_code,
							Recd_name:               Recdcheckproduct.Recd_name,
							Recd_unit:               Recdcheckproduct.Recd_unit,
							Recd_price:              Recdcheckproduct.Recd_price,
							Recd_topical:            Recdcheckproduct.Recd_topical,
							Recd_direction:          Recdcheckproduct.Recd_direction,
							Pdsh_type_id:            15,
							Ctm_prefix:              Ctm.Ctm_prefix,
							Ctm_fname:               Ctm.Ctm_fname,
							Ctm_lname:               Ctm.Ctm_lname,
							Ctm_fname_en:            Ctm.Ctm_fname_en,
							Ctm_lname_en:            Ctm.Ctm_lname_en,
							Ctm_tel:                 Ctm.Ctm_tel,
							Ctm_gender:              Ctm.Ctm_gender,
						}
						// fmt.Println(processProductUsed)
						ProcessProductStoreUsed(&processProductUsed)
					}
					//จบรายการยาที่ใช้
				}
			}
		}
		//จบสร้างรายการตรวจ
	}

	if Rec.Rec_is_process == 1 {
		//สร้างรายการคอร์ส
		var Recd []structs.ProcessReceiptDetail
		if errRecdc := models.GetProcessReceiptDetailCourse(receipt_id, &Recd); errRecdc != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Receipt detail course invalid.",
				"data":    errRecdc.Error(),
			})
			return
		}
		for _, Recdcourse := range Recd {
			var course structs.ProcessCourseService
			if errCourse := models.GetCourseServicesId(*Recdcourse.Course_id, &course); errCourse != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Course invalid.",
					"data":    errCourse.Error(),
				})
				return
			}
			if Recdcourse.Queue_id != nil {
				//ผ่านห้องตรวจใช้คอร์สเสร็จ
				var Ser_is_active = 2
				var Ser_qty = int(Recdcourse.Recd_qty)
				if course.Course_lock_drug == 1 {
					Ser_qty = 1
				}
				Rec_pay_datetime, _ := time.Parse(time.RFC3339, Rec.Rec_pay_datetime)
				processServices := structs.ProcessService{
					Id:                0,
					Shop_id:           Rec.Shop_id,
					Shop_mother_id:    shopMId,
					Receipt_id:        receipt_id,
					Receipt_detail_id: Recdcourse.Id,
					User_id:           userId,
					Ser_customer_id:   Rec.Customer_id,
					Customer_id:       Rec.Customer_id,
					Course_id:         *Recdcourse.Course_id,
					Ser_code:          Recdcourse.Recd_code,
					Ser_name:          Recdcourse.Recd_name,
					Ser_lock_drug:     course.Course_lock_drug,
					Ser_qty:           int(Recdcourse.Recd_qty),
					Ser_unit:          Recdcourse.Recd_unit,
					Ser_use_date:      course.Course_use_date,
					Ser_exp:           course.Course_exp_date,
					Ser_use:           Ser_qty,
					Ser_price_total:   Recdcourse.Recd_total,
					Ser_is_active:     Ser_is_active,
					Ser_datetime:      Rec_pay_datetime.Format("2006-01-02 15:04:05"),
					Ser_create:        time.Now().Format("2006-01-02 15:04:05"),
					Ser_update:        time.Now().Format("2006-01-02 15:04:05"),
				}
				if errCS := models.CreateServices(&processServices); errCS != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Create Services error.",
						"data":    errCS.Error(),
					})
					return
				}
				//อัพเดทคิวคอร์ส
				var queueCourse structs.ProcessServiceQueueCourse
				if errQueueCourse := models.GetProcessServiceQueueCourse(Recdcourse.Id, &queueCourse); errQueueCourse != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Queue course invalid.",
						"data":    errQueueCourse.Error(),
					})
					return
				}
				processServiceQueueCourseUpdate := structs.ProcessServiceQueueCourseUpdate{
					Service_id: processServices.Id,
					Receipt_id: receipt_id,
				}
				if errUSQC := models.UpdateProcessServiceQueueCourseUpdate(queueCourse.Queue_course_id, &processServiceQueueCourseUpdate); errUSQC != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Update Queue course error.",
						"data":    errUSQC.Error(),
					})
					return
				}
				//จบอัพเดทคิวคอร์ส
				processServicesUsed := structs.ProcessServiceUsed{
					Id:             0,
					Shop_id:        Rec.Shop_id,
					Shop_mother_id: shopMId,
					Shop_used_id:   Shop_id,
					Service_id:     processServices.Id,
					Queue_id:       *Recdcourse.Queue_id,
					Receipt_id:     receipt_id,
					Course_id:      *Recdcourse.Course_id,
					Customer_id:    Rec.Customer_id,
					User_id:        userId,
					Seru_code:      Recdcourse.Recd_code,
					Seru_name:      Recdcourse.Recd_name,
					Seru_qty:       int(Recdcourse.Recd_qty),
					Seru_unit:      Recdcourse.Recd_unit,
					Seru_cost:      course.Course_cost,
					Seru_date:      Rec_pay_datetime.Format("2006-01-02 15:04:05"),
					Seru_is_active: 1,
					Seru_datetime:  Rec_pay_datetime.Format("2006-01-02 15:04:05"),
					Seru_create:    time.Now().Format("2006-01-02 15:04:05"),
					Seru_update:    time.Now().Format("2006-01-02 15:04:05"),
				}
				if errCSU := models.CreateServicesUsed(&processServicesUsed); errCSU != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Create Services error.",
						"data":    errCSU.Error(),
					})
					return
				}
				//จบผ่านห้องตรวจใช้คอร์สเสร็จ
				//รายการยาในคอร์ส
				var Recdcp []structs.ProcessReceiptDetail
				if errCoursePro := models.GetProcessReceiptDetailCourseProduct(receipt_id, *Recdcourse.Course_id, &Recdcp); errCoursePro != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Course product invalid.",
						"data":    errCoursePro.Error(),
					})
					return
				}
				for _, Recdserviceproduct := range Recdcp {
					Recd_qty := Recdserviceproduct.Recd_qty
					var Serp_is_active = 2

					if Recdserviceproduct.Recd_limit_qty > 0 { //เช็คคอร์สจำกัดยา
						Recd_qty = Recdserviceproduct.Recd_limit_qty
						Serp_is_active = 1
						if Recdserviceproduct.Recd_qty == Recdserviceproduct.Recd_limit_qty {
							Serp_is_active = 2
						}
					}
					var Serp_balance float64 = 0
					Serp_balance = Recd_qty - Recdserviceproduct.Recd_qty
					var Serp_lock_drug int = 1
					var Serp_use_set_qty float64 = 1
					if course.Course_lock_drug == 0 {
						Serp_lock_drug = 1
						Serp_use_set_qty = Recdserviceproduct.Recd_set_qty
					}
					Rec_pay_datetime, _ := time.Parse(time.RFC3339, Rec.Rec_pay_datetime)
					receiptServiceproduct := structs.ProcessServiceProduct{
						Id:                0,
						Shop_id:           Rec.Shop_id,
						Service_id:        processServices.Id,
						Course_id:         *Recdserviceproduct.Course_id,
						Receipt_id:        Recdserviceproduct.Receipt_id,
						Receipt_detail_id: Recdserviceproduct.Id,
						Product_id:        *Recdserviceproduct.Product_id,
						Product_store_id:  *Recdserviceproduct.Product_store_id,
						Product_unit_id:   *Recdserviceproduct.Product_unit_id,
						Serp_code:         Recdserviceproduct.Recd_code,
						Serp_name:         Recdserviceproduct.Recd_name,
						Serp_qty:          Recd_qty, //คอร์สจำกัดยาจะใช้ จำนวนยาเต็มในคอร์ส
						Serp_use:          Recdserviceproduct.Recd_qty,
						Serp_balance:      Serp_balance,
						Serp_unit:         Recdserviceproduct.Recd_unit,
						Serp_lock_drug:    Serp_lock_drug,
						Serp_use_set_qty:  Serp_use_set_qty,
						Serp_is_active:    Serp_is_active,
						Serp_datetime:     Rec_pay_datetime.Format("2006-01-02 15:04:05"),
						Serp_create:       time.Now().Format("2006-01-02 15:04:05"),
						Serp_modify:       time.Now().Format("2006-01-02 15:04:05"),
					}
					if errCSP := models.CreateServiceProduct(&receiptServiceproduct); errCSP != nil {
						c.AbortWithStatusJSON(200, gin.H{
							"status":  false,
							"message": "Create Service Product error.",
							"data":    errCSP.Error(),
						})
						return
					}
					receiptServiceproductused := structs.ProcessServiceProductUsed{
						Id:                 0,
						Shop_id:            Rec.Shop_id,
						Service_id:         processServices.Id,
						Service_used_id:    processServicesUsed.Id,
						Course_id:          *Recdserviceproduct.Course_id,
						Queue_id:           *Recdcourse.Queue_id,
						Receipt_id:         Recdserviceproduct.Receipt_id,
						Customer_id:        Rec.Customer_id,
						Service_product_id: receiptServiceproduct.Id,
						Product_id:         *Recdserviceproduct.Product_id,
						Product_store_id:   *Recdserviceproduct.Product_store_id,
						Product_unit_id:    *Recdserviceproduct.Product_unit_id,
						Serpu_code:         Recdserviceproduct.Recd_code,
						Serpu_name:         Recdserviceproduct.Recd_name,
						Serpu_qty:          Recdserviceproduct.Recd_qty,
						Serpu_unit:         Recdserviceproduct.Recd_unit,
						Serpu_is_active:    1,
						Serpu_datetime:     Rec_pay_datetime.Format("2006-01-02 15:04:05"),
						Serpu_create:       time.Now().Format("2006-01-02 15:04:05"),
						Serpu_modify:       time.Now().Format("2006-01-02 15:04:05"),
					}
					if errCSPU := models.CreateServiceProductUsed(&receiptServiceproductused); errCSPU != nil {
						c.AbortWithStatusJSON(200, gin.H{
							"status":  false,
							"message": "Create Service Product Used error.",
							"data":    errCSPU.Error(),
						})
						return
					}
					var Receipt_id = Recdserviceproduct.Receipt_id
					var Receipt_detail_id = Recdserviceproduct.Id
					var Service_product_used_id = receiptServiceproductused.Id
					processProductUsed := structs.ProcessProduct{
						Shop_id:                 Rec.Shop_id,
						Customer_id:             Rec.Customer_id,
						User_id:                 userId,
						Product_id:              *Recdserviceproduct.Product_id,
						Product_store_id:        *Recdserviceproduct.Product_store_id,
						Receipt_id:              &Receipt_id,
						Receipt_detail_id:       &Receipt_detail_id,
						Check_product_id:        nil,
						Service_product_used_id: &Service_product_used_id,
						Queue_id:                Recdcourse.Queue_id,
						Pdsh_out_id:             3, // 1 ขาย, 2 โอน, 3 ใช้บริการ, 4 ตรวจ
						Pdsh_ref_doc_no:         Rec.Rec_code,
						Pdso_qty:                Recdserviceproduct.Recd_qty,
						Recd_code:               Recdserviceproduct.Recd_code,
						Recd_name:               Recdserviceproduct.Recd_name,
						Recd_unit:               Recdserviceproduct.Recd_unit,
						Recd_price:              Recdserviceproduct.Recd_price,
						Recd_topical:            Recdserviceproduct.Recd_topical,
						Recd_direction:          Recdserviceproduct.Recd_direction,
						Pdsh_type_id:            13,
						Ctm_prefix:              Ctm.Ctm_prefix,
						Ctm_fname:               Ctm.Ctm_fname,
						Ctm_lname:               Ctm.Ctm_lname,
						Ctm_fname_en:            Ctm.Ctm_fname_en,
						Ctm_lname_en:            Ctm.Ctm_lname_en,
						Ctm_tel:                 Ctm.Ctm_tel,
						Ctm_gender:              Ctm.Ctm_gender,
					}
					// fmt.Println(processProductUsed)
					ProcessProductStoreUsed(&processProductUsed)
					//ปรับคอร์ส lock ยา
					if processServices.Ser_is_active == 2 {
						if Serp_is_active == 1 {
							processServicesupdate := structs.ProcessServiceUpdate{
								Ser_is_active: 1,
								Ser_update:    time.Now().Format("2006-01-02 15:04:05"),
							}
							if errUS := models.UpdateService(processServices.Id, &processServicesupdate); errUS != nil {
								c.AbortWithStatusJSON(200, gin.H{
									"status":  false,
									"message": "Update Services error.",
									"data":    errUS.Error(),
								})
								return
							}
						}
					}
					//จบปรับคอร์ส lock ยา
				}
				//จบรายการยาในคอร์ส
				//จบผ่านห้องตรวจใช้คอร์สเสร็จ
			} else {
				var Course_exp_date *string
				Course_exp_date = nil
				Rec_pay_datetime, _ := time.Parse(time.RFC3339, Rec.Rec_pay_datetime)
				processServices := structs.ProcessService{
					Id:                0,
					Shop_id:           Rec.Shop_id,
					Shop_mother_id:    shopMId,
					Receipt_id:        receipt_id,
					Receipt_detail_id: Recdcourse.Id,
					User_id:           userId,
					Ser_customer_id:   Rec.Customer_id,
					Customer_id:       Rec.Customer_id,
					Course_id:         *Recdcourse.Course_id,
					Ser_code:          Recdcourse.Recd_code,
					Ser_name:          Recdcourse.Recd_name,
					Ser_lock_drug:     course.Course_lock_drug,
					Ser_qty:           int(Recdcourse.Recd_qty),
					Ser_unit:          Recdcourse.Recd_unit,
					Ser_use_date:      course.Course_use_date,
					Ser_exp:           course.Course_exp_date,
					Ser_exp_date:      Course_exp_date,
					Ser_use:           0,
					Ser_price_total:   Recdcourse.Recd_total,
					Ser_is_active:     1,
					Ser_datetime:      Rec_pay_datetime.Format("2006-01-02 15:04:05"),
					Ser_create:        time.Now().Format("2006-01-02 15:04:05"),
					Ser_update:        time.Now().Format("2006-01-02 15:04:05"),
				}
				errCCheck := models.CreateServices(&processServices)
				if errCCheck != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Create Services error.",
						"data":    "",
					})
					return
				}
				//รายการยาในคอร์ส
				if course.Course_lock_drug == 1 {
					// ดึงยาแบบใหม่ในใบเสร็จ
					var Recdcp []structs.ProcessReceiptDetail
					if errCoursePro := models.GetProcessReceiptDetailCourseProduct(receipt_id, *Recdcourse.Course_id, &Recdcp); errCoursePro != nil {
						c.AbortWithStatusJSON(200, gin.H{
							"status":  false,
							"message": "Course product invalid.",
							"data":    errCoursePro.Error(),
						})
						return
					}
					for _, Recdserviceproduct := range Recdcp {
						Recd_qty := Recdserviceproduct.Recd_qty
						var Serp_lock_drug int = 1
						var Serp_use_set_qty float64 = 1
						Rec_pay_datetime, _ := time.Parse(time.RFC3339, Rec.Rec_pay_datetime)
						if *Recdserviceproduct.Product_id != 0 {
							serviceproduct := structs.ProcessServiceProduct{
								Id:                0,
								Shop_id:           Rec.Shop_id,
								Service_id:        processServices.Id,
								Course_id:         *Recdcourse.Course_id,
								Receipt_id:        receipt_id,
								Receipt_detail_id: Recdcourse.Id,
								Product_id:        *Recdserviceproduct.Product_id,
								Product_store_id:  *Recdserviceproduct.Product_store_id,
								Product_unit_id:   *Recdserviceproduct.Product_unit_id,
								Serp_code:         Recdserviceproduct.Recd_code,
								Serp_name:         Recdserviceproduct.Recd_name,
								Serp_qty:          Recd_qty,
								Serp_use:          0,
								Serp_balance:      Recd_qty,
								Serp_unit:         Recdserviceproduct.Recd_unit,
								Serp_lock_drug:    Serp_lock_drug,
								Serp_use_set_qty:  Serp_use_set_qty,
								Serp_is_active:    1,
								Serp_datetime:     Rec_pay_datetime.Format("2006-01-02 15:04:05"),
								Serp_create:       time.Now().Format("2006-01-02 15:04:05"),
								Serp_modify:       time.Now().Format("2006-01-02 15:04:05"),
							}
							if errCSP := models.CreateServiceProduct(&serviceproduct); errCSP != nil {
								c.AbortWithStatusJSON(200, gin.H{
									"status":  false,
									"message": "Create Service Product error.",
									"data":    errCSP.Error(),
								})
								return
							}
						}
					}
					// 	var productSet []structs.ProcessCourseProductSet
					// 	if errc := models.GetProcessCourseProduct(*Recdcourse.Course_id, &productSet); errc != nil {
					// 		c.AbortWithStatusJSON(200, gin.H{
					// 			"status":  false,
					// 			"message": "Course List invalid.",
					// 			"data":    errc,
					// 		})
					// 		return
					// 	}
					// 	for _, pSet := range productSet {
					// 		var itemproductset structs.ProcessProductSet
					// 		if errCO := models.GetProcessProductSetId(pSet.Product_id, Shop_id, &itemproductset); errCO != nil {
					// 			c.AbortWithStatusJSON(200, gin.H{
					// 				"status":  false,
					// 				"message": "Course List invalid.",
					// 				"data":    errCO,
					// 			})
					// 			return
					// 		}
					// 		var Serp_lock_drug int = 1
					// 		var Serp_use_set_qty float64 = 1
					// 		var Cp_amount float64 = pSet.Cp_amount
					// 		if course.Course_lock_drug == 1 {
					// 			Cp_amount = pSet.Cp_amount * Recdcourse.Recd_qty
					// 		}
					// 		if itemproductset.Product_id != 0 {
					// 			serviceproduct := structs.ProcessServiceProduct{
					// 				Id:                0,
					// 				Shop_id:           Rec.Shop_id,
					// 				Service_id:        processServices.Id,
					// 				Course_id:         *Recdcourse.Course_id,
					// 				Receipt_id:        receipt_id,
					// 				Receipt_detail_id: Recdcourse.Id,
					// 				Product_id:        itemproductset.Product_id,
					// 				Product_store_id:  itemproductset.Product_store_id,
					// 				Product_unit_id:   itemproductset.Product_units_id,
					// 				Serp_code:         itemproductset.Pd_code,
					// 				Serp_name:         itemproductset.Pd_name,
					// 				Serp_qty:          Cp_amount,
					// 				Serp_use:          0,
					// 				Serp_balance:      Cp_amount,
					// 				Serp_unit:         itemproductset.U_name,
					// 				Serp_lock_drug:    Serp_lock_drug,
					// 				Serp_use_set_qty:  Serp_use_set_qty,
					// 				Serp_is_active:    1,
					// 				Serp_datetime:     Rec_pay_datetime.Format("2006-01-02 15:04:05"),
					// 				Serp_create:       time.Now().Format("2006-01-02 15:04:05"),
					// 				Serp_modify:       time.Now().Format("2006-01-02 15:04:05"),
					// 			}
					// 			if errCSP := models.CreateServiceProduct(&serviceproduct); errCSP != nil {
					// 				c.AbortWithStatusJSON(200, gin.H{
					// 					"status":  false,
					// 					"message": "Create Service Product error.",
					// 					"data":    errCSP.Error(),
					// 				})
					// 				return
					// 			}
					// 		}
					// 	}
				}
				//จบรายการยาในคอร์ส
				//จบไม่ผ่านห้องตรวจ ซื้อคอร์สรอใช้
			}
		}
		//จบสร้างรายการคอร์ส
	}

	if Rec.Rec_is_process == 1 {
		//ตัดคลังสินค้า
		var Recd []structs.ProcessReceiptDetail
		if errRecdc := models.GetProcessReceiptDetailProduct(receipt_id, &Recd); errRecdc != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Receipt detail product invalid.",
				"data":    errRecdc.Error(),
			})
			return
		}
		for _, Recdproduct := range Recd {
			var Recd_qty float64 = 0
			Recd_qty = Recdproduct.Recd_qty * Recdproduct.Recd_rate
			var Receipt_id = Recdproduct.Receipt_id
			var Receipt_detail_id = Recdproduct.Id
			processProduct := structs.ProcessProduct{
				Shop_id:                 Rec.Shop_id,
				Customer_id:             Rec.Customer_id,
				User_id:                 userId,
				Product_id:              *Recdproduct.Product_id,
				Product_store_id:        *Recdproduct.Product_store_id,
				Receipt_id:              &Receipt_id,
				Receipt_detail_id:       &Receipt_detail_id,
				Check_product_id:        nil,
				Service_product_used_id: nil,
				Queue_id:                nil,
				Pdsh_out_id:             1,
				Pdsh_ref_doc_no:         Rec.Rec_code,
				Pdso_qty:                Recd_qty,
				Recd_code:               Recdproduct.Recd_code,
				Recd_name:               Recdproduct.Recd_name,
				Recd_unit:               Recdproduct.Recd_unit,
				Recd_price:              Recdproduct.Recd_price,
				Recd_topical:            Recdproduct.Recd_topical,
				Recd_direction:          Recdproduct.Recd_direction,
				Pdsh_type_id:            11,
				Ctm_prefix:              Ctm.Ctm_prefix,
				Ctm_fname:               Ctm.Ctm_fname,
				Ctm_lname:               Ctm.Ctm_lname,
				Ctm_fname_en:            Ctm.Ctm_fname_en,
				Ctm_lname_en:            Ctm.Ctm_lname_en,
				Ctm_tel:                 Ctm.Ctm_tel,
				Ctm_gender:              Ctm.Ctm_gender,
			}
			// fmt.Println(processProduct)
			ProcessProductStoreOut(&processProduct)
		}
		//จบตัดคลังสินค้า
	}

	if Rec.Rec_is_process == 1 {
		//รายการวงเงิน
		var Recd []structs.ProcessReceiptDetail
		if errRecdc := models.GetProcessReceiptDetailCoin(receipt_id, &Recd); errRecdc != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Receipt detail coin invalid.",
				"data":    errRecdc.Error(),
			})
			return
		}
		for _, Recdcoin := range Recd {
			var processCoin structs.ProcessCoin
			if errpcoin := models.GetProcessCustomerCoin(Rec.Customer_id, &processCoin); errpcoin != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Customer coin invalid.",
					"data":    errpcoin.Error(),
				})
				return
			}
			var Ctm_coin float64 = 0
			Ctm_coin = processCoin.Ctm_coin + Recdcoin.Recd_limit_qty
			var coinCustomerGroup structs.CoinCustomerGroupId
			if errpcoinid := models.GetProcessCoinId(*Recdcoin.Coin_id, &coinCustomerGroup); errpcoinid != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Get coin invalid.",
					"data":    errpcoinid.Error(),
				})
				return
			}
			processCoinUpdate := structs.ProcessCoin{
				Id:                Rec.Customer_id,
				Ctm_coin:          Ctm_coin,
				Customer_group_id: coinCustomerGroup.Customer_group_id,
				Ctm_update:        time.Now().Format("2006-01-02 15:04:05"),
			}
			if errpcoinupdate := models.UpdateProcessCustomerCoin(Rec.Customer_id, &processCoinUpdate); errpcoinupdate != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Update Services error.",
					"data":    errpcoinupdate.Error(),
				})
				return
			}
			// Add coin history
			AddCoinHistory := structs.AddCoinHistory{
				Shop_id:     Rec.Shop_id,
				Customer_id: Rec.Customer_id,
				Receipt_id:  receipt_id,
				Rec_code:    Rec.Rec_code,
				Ch_forward:  float64(processCoin.Ctm_coin),
				Ch_amount:   float64(Recdcoin.Recd_limit_qty),
				Ch_total:    float64(Ctm_coin),
				Ch_comment:  "Add Process: Add Coin History",
				Ch_create:   time.Now().Format("2006-01-02 15:04:05"),
			}
			if errCoinHistory := models.CreateCoinHistory(&AddCoinHistory); errCoinHistory != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Add coin history error.",
					"data":    errCoinHistory.Error(),
				})
				return
			}
		}
		//จบรายการวงเงิน
	}

	//อัพเดทคิว
	if Rec.Queue_id != nil {
		Queue_id := Rec.Queue_id
		var SerQ structs.ProcessQueue
		if errRecSerQ := models.GetProcessQueueById(*Queue_id, &SerQ); errRecSerQ != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Queue invalid.",
				"data":    errRecSerQ.Error(),
			})
			return
		}
		//หานาทีเริ่มถึงจบ Calculate the time difference in minutes
		startTimeStr := SerQ.Que_datetime
		layout := "2006-01-02T15:04:05-07:00"
		startTime, errTime := time.Parse(layout, startTimeStr)
		if errTime != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Error parsing start time.",
				"data":    errTime.Error(),
			})
			return
		}
		currentTime := time.Now()
		duration := currentTime.Sub(startTime)
		minutes := int(duration.Minutes())

		queueServiceupdate := structs.QueueProcessUpdate{
			Que_status_id:    4,
			Que_datetime_out: currentTime.Format("2006-01-02 15:04:05"),
			Que_time_end:     minutes,
			Que_update:       time.Now().Format("2006-01-02 15:04:05"),
		}
		if errQS := models.UpdateQueueProcessId(SerQ.Id, &queueServiceupdate); errQS != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Update Queue Services error.",
				"data":    errQS.Error(),
			})
			return
		}
	}

	Invoice_id := Rec.Invoice_id
	var PINV structs.ProcessInvoice
	if errRecPINV := models.GetProcessInvoiceById(Invoice_id, &PINV); errRecPINV != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Queue invalid.",
			"data":    errRecPINV.Error(),
		})
		return
	}

	//หานาทีเริ่มถึงจบ Calculate the time difference in minutes
	startTimeStr := PINV.Inv_datetime
	layout := "2006-01-02T15:04:05-07:00"
	startTime, errTime := time.Parse(layout, startTimeStr)
	if errTime != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error parsing start time.",
			"data":    errTime.Error(),
		})
		return
	}
	currentTime := time.Now()
	duration := currentTime.Sub(startTime)
	minutes := int(duration.Minutes())

	invoiceServiceupdate := structs.InvoiceProcessUpdate{
		Inv_time_end: minutes,
		Inv_update:   time.Now().Format("2006-01-02 15:04:05"),
	}
	if errQS := models.UpdateInvoiceProcessId(PINV.Id, &invoiceServiceupdate); errQS != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Update invoice error.",
			"data":    errQS.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Process Receipt Success.",
		"data":    receipt_id,
	})
}

func CancelProcess(c *gin.Context) {
	RECId := libs.StrToInt(c.Params.ByName("id"))
	var REC structs.Receipt
	if errCK := models.GetReceiptId(RECId, &REC); errCK != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Receipt invalid.",
			"data":    errCK.Error(),
		})
		return
	}

	receipt_id := RECId
	var Rec structs.ProcessReceipt
	if errRec := models.GetProcessReceiptId(receipt_id, &Rec); errRec != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Receipt invalid.",
			"data":    errRec.Error(),
		})
		return
	}

	userId := Rec.User_id
	Shop_id := Rec.Shop_id

	if Rec.Rec_is_process == 1 {
		//check
		updateCheck := structs.CheckUpdate{
			Chk_is_active: 0,
			Chk_update:    time.Now().Format("2006-01-02 15:04:05"),
		}
		// if errQS := models.UpdateCaneclCheckId(receipt_id, &updateCheck); errQS != nil {
		if errCCRI := models.UpdateCaneclCheckReceiptId(receipt_id, &updateCheck); errCCRI != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Update check receipt error.",
				"data":    errCCRI.Error(),
			})
			return
		}
		if errCCR := models.UpdateCancalCheckReceipt(Rec.Invoice_id); errCCR != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Update check invoice error.",
				"data":    errCCR.Error(),
			})
			return
		}
		//check product
		updateCheckProduct := structs.CheckProductUpdate{
			Chkp_is_active: 0,
			Chkp_modify:    time.Now().Format("2006-01-02 15:04:05"),
		}
		if errQS := models.UpdateCaneclCheckProductId(receipt_id, &updateCheckProduct); errQS != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Update Check Product error.",
				"data":    errQS.Error(),
			})
			return
		}
		//service product used
		updateServiceProductUsed := structs.ServiceProductUsedUpdate{
			Serpu_is_active: 0,
			Serpu_modify:    time.Now().Format("2006-01-02 15:04:05"),
		}
		if errQS := models.UpdateCaneclServiceProductUsedId(receipt_id, &updateServiceProductUsed); errQS != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Update Service Product Used error.",
				"data":    errQS.Error(),
			})
			return
		}
		//service product
		updateServiceProduct := structs.ServiceProductUpdate{
			Serp_is_active: 0,
			Serp_modify:    time.Now().Format("2006-01-02 15:04:05"),
		}
		if errQS := models.UpdateCaneclServiceProductId(receipt_id, &updateServiceProduct); errQS != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Update Service Product error.",
				"data":    errQS.Error(),
			})
			return
		}
		//service used
		updateServiceUsed := structs.ServiceUsedUpdate{
			Seru_is_active: 0,
			Seru_update:    time.Now().Format("2006-01-02 15:04:05"),
		}
		if errQS := models.UpdateCaneclServiceUsedId(receipt_id, &updateServiceUsed); errQS != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Update Service Used error.",
				"data":    errQS.Error(),
			})
			return
		}
		//service
		updateService := structs.ServiceUpdate{
			Ser_is_active: 0,
			Ser_update:    time.Now().Format("2006-01-02 15:04:05"),
		}
		if errQS := models.UpdateCaneclServiceId(receipt_id, &updateService); errQS != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Update Service error.",
				"data":    errQS.Error(),
			})
			return
		}
		//queuecourse
		if errCQC := models.UpdateCaneclQueueCourseId(receipt_id); errCQC != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Update Queue Course error.",
				"data":    errCQC.Error(),
			})
			return
		}
		//sticker
		stickerUpdate := structs.StickerUpdate{
			Sticker_is_del: 1,
		}
		if errQS := models.UpdateCaneclStickerId(receipt_id, &stickerUpdate); errQS != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Update Sticker error.",
				"data":    errQS.Error(),
			})
			return
		}
		//coin
		var Recd []structs.ProcessReceiptDetail
		if errRecdc := models.GetProcessReceiptDetailCoin(receipt_id, &Recd); errRecdc != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Receipt detail product invalid.",
				"data":    errRecdc.Error(),
			})
			return
		}
		for _, Recdcoin := range Recd {
			var processCoin structs.ProcessCoin
			if errpcoin := models.GetProcessCustomerCoin(Rec.Customer_id, &processCoin); errpcoin != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Receipt detail product invalid.",
					"data":    errpcoin.Error(),
				})
				return
			}
			var Ctm_coin float64 = 0
			Ctm_coin = processCoin.Ctm_coin - Recdcoin.Recd_limit_qty
			processCoinUpdate := structs.ProcessCoinCancelUpdate{
				Id:         Rec.Customer_id,
				Ctm_coin:   Ctm_coin,
				Ctm_update: time.Now().Format("2006-01-02 15:04:05"),
			}
			if errpcoinupdate := models.UpdateProcessCustomerCoinCancel(Rec.Customer_id, &processCoinUpdate); errpcoinupdate != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Update Services error.",
					"data":    errpcoinupdate.Error(),
				})
				return
			}
			// Add coin history
			AddCoinHistory := structs.AddCoinHistory{
				Shop_id:     Rec.Shop_id,
				Customer_id: Rec.Customer_id,
				Receipt_id:  receipt_id,
				Rec_code:    Rec.Rec_code,
				Ch_forward:  float64(processCoin.Ctm_coin),
				Ch_amount:   float64(Recdcoin.Recd_limit_qty) * float64(-1),
				Ch_total:    float64(Ctm_coin),
				Ch_comment:  "Cancel Process: Minus Coin History",
				Ch_create:   time.Now().Format("2006-01-02 15:04:05"),
			}
			if errCoinHistory := models.CreateCoinHistory(&AddCoinHistory); errCoinHistory != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Cancel coin history error.",
					"data":    errCoinHistory.Error(),
				})
				return
			}
		}

		cancelProductUsed := structs.CancelProductReceipt{
			Shop_id:     Shop_id,
			Customer_id: Rec.Customer_id,
			User_id:     userId,
			Receipt_id:  receipt_id,
			Queue_id:    Rec.Queue_id,
		}
		CancelProductStore(&cancelProductUsed)
	}
	c.JSON(200, gin.H{
		"status":  true,
		"message": "Cancel Receipt Success.",
		"data":    receipt_id,
	})
}

func CancelProductStore(data *structs.CancelProductReceipt) int {
	var productStoreHistory []structs.ProcessProductStoreHistory
	models.GetReceiptProductStoreHistory(data.Receipt_id, &productStoreHistory)
	for _, History := range productStoreHistory {
		var Pdso_total float64 = 0
		var Pdsh_order_forward float64 = 0
		if History.Pdsh_out_id != 2 {
			if History.Pdsh_out_id != 1 {
				// ProductStoreOrder
				var productStoreOrder structs.ProcessProductStoreOrder
				models.GetProcessProductStoreOrders(History.Product_store_order_id, &productStoreOrder)
				var Pdso_use float64 = 0
				Pdso_use = productStoreOrder.Pdso_use - History.Pdsh_out
				Pdso_total = productStoreOrder.Pdso_total + History.Pdsh_out
				Pdsh_order_forward = productStoreOrder.Pdso_total
				productStoreOrderUpdate := structs.ProcessProductStoreOrder{
					Id:               productStoreOrder.Id,
					Product_store_id: productStoreOrder.Product_store_id,
					Pdso_expire:      productStoreOrder.Pdso_expire,
					Pdso_use:         Pdso_use,
					Pdso_out:         productStoreOrder.Pdso_out,
					Pdso_total:       Pdso_total,
					Pdso_update:      time.Now().Format("2006-01-02 15:04:05"),
				}
				models.UpdateProductStoreOrders(productStoreOrder.Id, &productStoreOrderUpdate)
			} else {
				// ProductStoreOrder
				var productStoreOrder structs.ProcessProductStoreOrder
				models.GetProcessProductStoreOrders(History.Product_store_order_id, &productStoreOrder)
				var Pdso_out float64 = 0
				Pdso_out = productStoreOrder.Pdso_out - History.Pdsh_out
				Pdso_total = productStoreOrder.Pdso_total + History.Pdsh_out
				Pdsh_order_forward = productStoreOrder.Pdso_total
				productStoreOrderUpdate := structs.ProcessProductStoreOrder{
					Id:               productStoreOrder.Id,
					Product_store_id: productStoreOrder.Product_store_id,
					Pdso_expire:      productStoreOrder.Pdso_expire,
					Pdso_use:         productStoreOrder.Pdso_use, //Pdso_use,
					Pdso_out:         Pdso_out,                   //productStoreOrder.Pdso_out,
					Pdso_total:       Pdso_total,
					Pdso_update:      time.Now().Format("2006-01-02 15:04:05"),
				}
				models.UpdateProductStoreOrders(productStoreOrder.Id, &productStoreOrderUpdate)
			}
			// ProductStore
			var productStore structs.ProcessProductStore
			models.GetProcessProductStoreId(History.Product_store_id, &productStore)
			var Pds_out float64 = 0
			Pds_out = productStore.Pds_out - History.Pdsh_out
			var Pds_total float64 = 0
			Pds_total = productStore.Pds_total + History.Pdsh_out
			productStoreUpdate := structs.ProcessProductStore{
				Id:            History.Product_store_id,
				Shop_store_id: productStore.Shop_store_id,
				Pds_out:       Pds_out,
				Pds_total:     Pds_total,
				Pds_update:    time.Now().Format("2006-01-02 15:04:05"),
			}
			models.UpdateProductStore(History.Product_store_id, &productStoreUpdate)
			var Pdsh_type_id int = 16
			if History.Pdsh_type_id == 11 {
				Pdsh_type_id = 12
			} else if History.Pdsh_type_id == 13 {
				Pdsh_type_id = 14
			}
			// ProductStoreHistory
			productStoreHistoryUpdate := structs.ProcessProductStoreHistory{
				Id:                      0,
				Shop_id:                 data.Shop_id,
				Shop_store_id:           History.Shop_store_id,
				Product_store_id:        History.Product_store_id,
				Product_store_order_id:  History.Product_store_order_id,
				Receipt_id:              History.Receipt_id,
				Receipt_detail_id:       History.Receipt_detail_id,
				Check_product_id:        History.Check_product_id,
				Service_product_used_id: History.Service_product_used_id,
				Queue_id:                History.Queue_id,
				Pdsh_in:                 History.Pdsh_out,
				Pdsh_out:                0,
				Pdsh_order_forward:      Pdsh_order_forward,
				Pdsh_total_forward:      productStore.Pds_total,
				Pdsh_amount:             History.Pdsh_out,
				Pdsh_order_total:        Pdso_total,
				Pdsh_total:              Pds_total,
				Pdsh_inout:              1,
				Pdsh_out_id:             0,
				Pdsh_ref_doc_no:         History.Pdsh_ref_doc_no,
				User_id:                 data.User_id,
				Product_id:              History.Product_id,
				Pd_code:                 History.Pd_code,
				Pd_name:                 History.Pd_name,
				Pdsh_type_id:            Pdsh_type_id,
				Customer_id:             History.Customer_id,
				Ctm_prefix:              History.Ctm_prefix,
				Ctm_fname:               History.Ctm_fname,
				Ctm_lname:               History.Ctm_lname,
				Ctm_fname_en:            History.Ctm_fname_en,
				Ctm_lname_en:            History.Ctm_lname_en,
				Ctm_tel:                 History.Ctm_tel,
				Ctm_gender:              History.Ctm_gender,
				Pdsh_comment:            "Cancel Product",
				Pdsh_date:               time.Now().Format("2006-01-02"),
				Pdsh_modify:             time.Now().Format("2006-01-02 15:04:05"),
			}
			models.CreateProductStoreHistory(&productStoreHistoryUpdate)
		}
	}
	return 1
}

func ProcessProduct(c *gin.Context) {
	// var payload structs.ObjPayloadProcessProduct
	// if errPL := c.ShouldBindJSON(&payload); errPL != nil {
	// 	c.AbortWithStatusJSON(200, gin.H{
	// 		"status":  false,
	// 		"message": "Invalid request data.",
	// 		"data":    errPL.Error(),
	// 	})
	// 	return
	// }
	// Shop_id := libs.StrToInt(c.Params.ByName("shopId"))
	// var shopData []structs.ShopReadResponse
	// errS := models.GetShopById(Shop_id, &shopData)
	// if errS != nil {
	// 	c.JSON(200, gin.H{
	// 		"status":  false,
	// 		"message": "Shop invalid",
	// 		"data":    errS.Error(),
	// 	})
	// 	return
	// }
	// var temp_amount float64 = payload.Pdso_qty
	// var amount float64 = payload.Pdso_qty
	// var expire string = ""
	// for amount > 0 {
	// 	temp_amount = amount
	// 	var productStoreOrder structs.ProcessProductStoreOrder
	// 	models.GetProcessProductStoreOrder(payload.Product_id, payload.Product_store_id, &productStoreOrder)
	// 	if productStoreOrder.Id != 0 {
	// 		fmt.Println(productStoreOrder.Id)
	// 		if productStoreOrder.Pdso_total < amount {
	// 			// fmt.Println(productStoreOrder.Pdso_total)
	// 			// fmt.Println(amount)
	// 			amount = amount - productStoreOrder.Pdso_total
	// 			temp_amount = temp_amount - amount
	// 			// fmt.Println(amount)
	// 			// fmt.Println(temp_amount)
	// 			// fmt.Println(1)
	// 			// ProductStoreOrder
	// 			var processProductStoreOrder structs.ProcessProductStoreOrder
	// 			models.GetProcessProductStoreOrderId(productStoreOrder.Id, &processProductStoreOrder)
	// 			var Pdso_out float64 = 0
	// 			Pdso_out = productStoreOrder.Pdso_out + temp_amount
	// 			var Pdso_total float64 = 0
	// 			Pdso_total = productStoreOrder.Pdso_total - temp_amount
	// 			if productStoreOrder.Pdso_expire < expire || expire == "" {
	// 				expire = productStoreOrder.Pdso_expire
	// 			}
	// 			productStoreOrderUpdate := structs.ProcessProductStoreOrder{
	// 				Id:               productStoreOrder.Id,
	// 				Product_store_id: productStoreOrder.Product_store_id,
	// 				Pdso_expire:      productStoreOrder.Pdso_expire,
	// 				Pdso_out:         Pdso_out,
	// 				Pdso_use:         productStoreOrder.Pdso_use,
	// 				Pdso_total:       Pdso_total,
	// 				Pdso_update:      time.Now().Format("2006-01-02 15:04:05"),
	// 			}
	// 			models.UpdateProductStoreOrders(productStoreOrder.Id, &productStoreOrderUpdate)
	// 			// ProductStore
	// 			var processProductStore structs.ProcessProductStore
	// 			models.GetProcessProductStoreId(payload.Product_store_id, &processProductStore)
	// 			var Pds_out float64 = 0
	// 			Pds_out = processProductStore.Pds_out + temp_amount
	// 			var Pds_total float64 = 0
	// 			Pds_total = processProductStore.Pds_total - temp_amount
	// 			productStoreUpdate := structs.ProcessProductStore{
	// 				Id:            payload.Product_store_id,
	// 				Shop_store_id: processProductStore.Shop_store_id,
	// 				Pds_out:       Pds_out,
	// 				Pds_total:     Pds_total,
	// 				Pds_update:    time.Now().Format("2006-01-02 15:04:05"),
	// 			}
	// 			models.UpdateProductStore(payload.Product_store_id, &productStoreUpdate)
	// 			// ProductStoreHistorys
	// 			productStoreHistoryUpdate := structs.ProcessProductStoreHistory{
	// 				Id:                      0,
	// 				Shop_id:                 Shop_id,
	// 				Shop_store_id:           processProductStore.Shop_store_id,
	// 				Product_store_id:        payload.Product_store_id,
	// 				Product_store_order_id:  processProductStoreOrder.Id,
	// 				Receipt_id:              payload.Receipt_id,
	// 				Receipt_detail_id:       payload.Receipt_detail_id,
	// 				Check_product_id:        payload.Check_product_id,
	// 				Service_product_used_id: payload.Service_product_used_id,
	// 				Queue_id:                payload.Queue_id,
	// 				Pdsh_in:                 0,
	// 				Pdsh_out:                temp_amount,
	// 				Pdsh_amount:             0 - temp_amount,
	// 				Pdsh_order_total:        Pdso_total,
	// 				Pdsh_total:              Pds_total,
	// 				Pdsh_inout:              2,
	// 				Pdsh_out_id:             payload.Pdsh_out_id,
	// 				Pdsh_ref_doc_no:         payload.Pdsh_ref_doc_no,
	// 				Pdsh_comment:            "Process Product",
	// 				Pdsh_date:               time.Now().Format("2006-01-02"),
	// 				Pdsh_modify:             time.Now().Format("2006-01-02 15:04:05"),
	// 			}
	// 			models.CreateProductStoreHistory(&productStoreHistoryUpdate)

	// 		} else {
	// 			// fmt.Println(productStoreOrder.Pdso_total)
	// 			// fmt.Println(amount)
	// 			amount = amount - productStoreOrder.Pdso_total
	// 			if amount < 0 {
	// 				amount = 0
	// 			}
	// 			temp_amount = temp_amount - amount
	// 			// fmt.Println(amount)
	// 			// fmt.Println(temp_amount)
	// 			// fmt.Println(2)
	// 			// ProductStoreOrder
	// 			var processProductStoreOrder structs.ProcessProductStoreOrder
	// 			models.GetProcessProductStoreOrderId(productStoreOrder.Id, &processProductStoreOrder)
	// 			var Pdso_out float64 = 0
	// 			Pdso_out = productStoreOrder.Pdso_out + temp_amount
	// 			var Pdso_total float64 = 0
	// 			Pdso_total = productStoreOrder.Pdso_total - temp_amount
	// 			if productStoreOrder.Pdso_expire < expire || expire == "" {
	// 				expire = productStoreOrder.Pdso_expire
	// 			}
	// 			productStoreOrderUpdate := structs.ProcessProductStoreOrder{
	// 				Id:               productStoreOrder.Id,
	// 				Product_store_id: productStoreOrder.Product_store_id,
	// 				Pdso_expire:      productStoreOrder.Pdso_expire,
	// 				Pdso_out:         Pdso_out,
	// 				Pdso_use:         productStoreOrder.Pdso_use,
	// 				Pdso_total:       Pdso_total,
	// 				Pdso_update:      time.Now().Format("2006-01-02 15:04:05"),
	// 			}
	// 			models.UpdateProductStoreOrders(productStoreOrder.Id, &productStoreOrderUpdate)
	// 			// ProductStore
	// 			var processProductStore structs.ProcessProductStore
	// 			models.GetProcessProductStoreId(payload.Product_store_id, &processProductStore)
	// 			var Pds_out float64 = 0
	// 			Pds_out = processProductStore.Pds_out + temp_amount
	// 			var Pds_total float64 = 0
	// 			Pds_total = processProductStore.Pds_total - temp_amount
	// 			productStoreUpdate := structs.ProcessProductStore{
	// 				Id:            payload.Product_store_id,
	// 				Shop_store_id: processProductStore.Shop_store_id,
	// 				Pds_out:       Pds_out,
	// 				Pds_total:     Pds_total,
	// 				Pds_update:    time.Now().Format("2006-01-02 15:04:05"),
	// 			}
	// 			models.UpdateProductStore(payload.Product_store_id, &productStoreUpdate)

	// 			// ProductStoreHistorys
	// 			productStoreHistoryUpdate := structs.ProcessProductStoreHistory{
	// 				Id:                      0,
	// 				Shop_id:                 Shop_id,
	// 				Shop_store_id:           processProductStore.Shop_store_id,
	// 				Product_store_id:        payload.Product_store_id,
	// 				Product_store_order_id:  processProductStoreOrder.Id,
	// 				Receipt_id:              payload.Receipt_id,
	// 				Receipt_detail_id:       payload.Receipt_detail_id,
	// 				Check_product_id:        payload.Check_product_id,
	// 				Service_product_used_id: payload.Service_product_used_id,
	// 				Queue_id:                payload.Queue_id,
	// 				Pdsh_in:                 0,
	// 				Pdsh_out:                temp_amount,
	// 				Pdsh_amount:             0 - temp_amount,
	// 				Pdsh_order_total:        Pdso_total,
	// 				Pdsh_total:              Pds_total,
	// 				Pdsh_inout:              2,
	// 				Pdsh_out_id:             payload.Pdsh_out_id,
	// 				Pdsh_ref_doc_no:         payload.Pdsh_ref_doc_no,
	// 				Pdsh_comment:            "Process Product",
	// 				Pdsh_date:               time.Now().Format("2006-01-02"),
	// 				Pdsh_modify:             time.Now().Format("2006-01-02 15:04:05"),
	// 			}
	// 			models.CreateProductStoreHistory(&productStoreHistoryUpdate)
	// 		}
	// 	} else {
	// 		// fmt.Println(productStoreOrder.Pdso_total)
	// 		// fmt.Println(amount)
	// 		amount = 0
	// 		// fmt.Println(amount)
	// 		// fmt.Println(temp_amount)
	// 		// fmt.Println(3)
	// 		// ProductStoreOrder
	// 		var productStoreOrder structs.ProcessProductStoreOrder
	// 		models.GetProcessProductStoreOrderLast(payload.Product_id, payload.Product_store_id, &productStoreOrder)
	// 		// var processProductStoreOrder structs.ProcessProductStoreOrder
	// 		// models.GetProcessProductStoreOrderId(productStoreOrder.Id, &processProductStoreOrder)
	// 		var Pdso_out float64 = 0
	// 		Pdso_out = productStoreOrder.Pdso_out + temp_amount
	// 		var Pdso_total float64 = 0
	// 		Pdso_total = productStoreOrder.Pdso_total - temp_amount
	// 		if productStoreOrder.Pdso_expire < expire || expire == "" {
	// 			expire = productStoreOrder.Pdso_expire
	// 		}
	// 		productStoreOrderUpdate := structs.ProcessProductStoreOrder{
	// 			Id:               productStoreOrder.Id,
	// 			Product_store_id: productStoreOrder.Product_store_id,
	// 			Pdso_expire:      productStoreOrder.Pdso_expire,
	// 			Pdso_out:         Pdso_out,
	// 			Pdso_use:         productStoreOrder.Pdso_use,
	// 			Pdso_total:       Pdso_total,
	// 			Pdso_update:      time.Now().Format("2006-01-02 15:04:05"),
	// 		}
	// 		models.UpdateProductStoreOrders(productStoreOrder.Id, &productStoreOrderUpdate)

	// 		// ProductStore
	// 		var processProductStore structs.ProcessProductStore
	// 		models.GetProcessProductStoreId(payload.Product_store_id, &processProductStore)
	// 		var Pds_out float64 = 0
	// 		Pds_out = processProductStore.Pds_out + temp_amount
	// 		var Pds_total float64 = 0
	// 		Pds_total = processProductStore.Pds_total - temp_amount
	// 		productStoreUpdate := structs.ProcessProductStore{
	// 			Id:            payload.Product_store_id,
	// 			Shop_store_id: processProductStore.Shop_store_id,
	// 			Pds_out:       Pds_out,
	// 			Pds_total:     Pds_total,
	// 			Pds_update:    time.Now().Format("2006-01-02 15:04:05"),
	// 		}
	// 		models.UpdateProductStore(payload.Product_store_id, &productStoreUpdate)
	// 		// ProductStoreHistorys
	// 		productStoreHistoryUpdate := structs.ProcessProductStoreHistory{
	// 			Id:                      0,
	// 			Shop_id:                 Shop_id,
	// 			Shop_store_id:           processProductStore.Shop_store_id,
	// 			Product_store_id:        payload.Product_store_id,
	// 			Product_store_order_id:  productStoreOrder.Id,
	// 			Receipt_id:              payload.Receipt_id,
	// 			Receipt_detail_id:       payload.Receipt_detail_id,
	// 			Check_product_id:        payload.Check_product_id,
	// 			Service_product_used_id: payload.Service_product_used_id,
	// 			Queue_id:                payload.Queue_id,
	// 			Pdsh_in:                 0,
	// 			Pdsh_out:                temp_amount,
	// 			Pdsh_amount:             0 - temp_amount,
	// 			Pdsh_order_total:        Pdso_total,
	// 			Pdsh_total:              Pds_total,
	// 			Pdsh_inout:              2,
	// 			Pdsh_out_id:             payload.Pdsh_out_id,
	// 			Pdsh_ref_doc_no:         payload.Pdsh_ref_doc_no,
	// 			Pdsh_comment:            "Process Product",
	// 			Pdsh_date:               time.Now().Format("2006-01-02"),
	// 			Pdsh_modify:             time.Now().Format("2006-01-02 15:04:05"),
	// 		}
	// 		models.CreateProductStoreHistory(&productStoreHistoryUpdate)
	// 	}
	// }
	// // Sticker
	// productStickerUpdate := structs.ProcessProductSticker{
	// 	Id:                  0,
	// 	Product_id:          payload.Product_id,
	// 	Customer_id:         payload.Customer_id,
	// 	User_id:             payload.User_id,
	// 	Receipt_id:          payload.Receipt_id,
	// 	Receipt_detail_id:   payload.Receipt_detail_id,
	// 	Sticker_code:        payload.Recd_code,
	// 	Sticker_name:        payload.Recd_name,
	// 	Sticker_amount:      payload.Pdso_qty,
	// 	Sticker_unit:        payload.Recd_unit,
	// 	Sticker_price:       payload.Recd_price,
	// 	Sticker_expdate:     expire,
	// 	Sticker_active_id:   1,
	// 	Sticker_print_label: 1,
	// 	Sticker_print_order: 1,
	// 	Sticker_topical:     payload.Recd_topical,
	// 	Sticker_direction:   payload.Recd_direction,
	// 	Sticker_is_del:      0,
	// 	Sticker_modify:      time.Now().Format("2006-01-02 15:04:05"),
	// }
	// models.CreateProductSticker(&productStickerUpdate)
	// c.JSON(200, gin.H{
	// 	"status":      true,
	// 	"message":     "Process Product Success.",
	// 	"data":        "",
	// 	"amount":      amount,
	// 	"temp_amount": temp_amount,
	// })
}

func ProcessProductExpire(c *gin.Context) {
	//ProductStoreOrder Expire
	var processProductStoreOrderExpire []structs.ProcessProductStoreOrderExpire
	errPPSO := models.GetProcessProductStoreOrderExpire(&processProductStoreOrderExpire)
	if errPPSO != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Shop invalid",
			"data":    errPPSO.Error(),
		})
		return
	}
	ItemList := []structs.ProcessProductStoreOrderExpire{}
	for _, productStoreOrder := range processProductStoreOrderExpire {
		//ProductStoreOrder
		var Pdso_exp float64 = 0
		Pdso_exp = productStoreOrder.Pdso_exp + productStoreOrder.Pdso_total
		var Pdso_total float64 = 0
		Pdso_total = 0
		Pdso_expire, _ := time.Parse(time.RFC3339, productStoreOrder.Pdso_expire)
		productStoreOrderUpdate := structs.ProcessProductStoreOrderExpire{
			Id:               productStoreOrder.Id,
			Product_store_id: productStoreOrder.Product_store_id,
			Pdso_expire:      Pdso_expire.Format("2006-01-02"),
			Pdso_exp:         Pdso_exp,
			Pdso_total:       Pdso_total,
			Pdso_update:      time.Now().Format("2006-01-02 15:04:05"),
		}
		ItemList = append(ItemList, productStoreOrderUpdate)
		models.UpdateProductStoreOrdersExpire(productStoreOrder.Id, &productStoreOrderUpdate)
		// ProductStore
		var processProductStore structs.ProcessProductStoreExpire
		models.GetProcessProductStoreExpire(productStoreOrder.Product_store_id, &processProductStore)
		var Pds_exp float64 = 0
		Pds_exp = processProductStore.Pds_exp + productStoreOrder.Pdso_total
		var Pds_total float64 = 0
		Pds_total = processProductStore.Pds_total - productStoreOrder.Pdso_total
		productStoreUpdate := structs.UpdateProcessProductStoreExpire{
			Id:            productStoreOrder.Product_store_id,
			Shop_store_id: processProductStore.Shop_store_id,
			Pds_exp:       Pds_exp,
			Pds_total:     Pds_total,
			Pds_update:    time.Now().Format("2006-01-02 15:04:05"),
		}
		models.UpdateProductStoreExpire(productStoreOrder.Product_store_id, &productStoreUpdate)
		var productId structs.ProductIdProcess
		models.GetProductIdProcess(processProductStore.Product_id, &productId)
		// ProductStoreHistorys
		productStoreHistoryUpdate := structs.ProcessProductStoreHistoryExp{
			Id:                      0,
			Shop_id:                 processProductStore.Shop_id,
			Shop_store_id:           processProductStore.Shop_store_id,
			Product_store_id:        productStoreOrder.Product_store_id,
			Product_store_order_id:  productStoreOrder.Id,
			Receipt_id:              nil,
			Receipt_detail_id:       nil,
			Check_product_id:        nil,
			Service_product_used_id: nil,
			Queue_id:                nil,
			Pdsh_in:                 0,
			Pdsh_out:                Pdso_exp,
			Pdsh_order_forward:      productStoreOrder.Pdso_total,
			Pdsh_total_forward:      processProductStore.Pds_total,
			Pdsh_amount:             0 - Pdso_exp,
			Pdsh_order_total:        Pdso_total,
			Pdsh_total:              Pds_total,
			Pdsh_inout:              2,
			Pdsh_out_id:             5, //ปะเภทการเคลื่อนไหวยาหรืออุปกรณ์ : 1 ขาย, 2 โอน, 3 ใช้บริการ, 4 ตรวจ, 5 หมดอายุ
			Pdsh_ref_doc_no:         "EXP",
			Product_id:              productId.Id,
			Pd_code:                 productId.Pd_code,
			Pd_name:                 productId.Pd_name,
			Pdsh_type_id:            17,
			Pdsh_comment:            "Process Product Expire",
			Pdsh_date:               Pdso_expire.Format("2006-01-02"),
			Pdsh_modify:             time.Now().Format("2006-01-02 15:04:05"),
		}
		models.CreateProductStoreHistoryExp(&productStoreHistoryUpdate)
	}
	c.JSON(200, gin.H{
		"status":  true,
		"message": "Process Product Expire Success.",
		"data":    "", //processProductStoreOrderExpire,
		// "dataupdate": ItemList,
	})
}

func ProcessProductExpireHistory(c *gin.Context) {
	//ProductStoreOrder Expire
	var processProductStoreOrderExpire []structs.ProcessProductStoreOrderExpire
	errPPSO := models.GetProcessProductStoreOrderExpireHistory(&processProductStoreOrderExpire)
	if errPPSO != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Shop invalid",
			"data":    errPPSO.Error(),
		})
		return
	}
	for _, productStoreOrder := range processProductStoreOrderExpire {
		//ProductStoreOrder
		var Pdso_exp float64 = 0
		Pdso_exp = productStoreOrder.Pdso_exp + productStoreOrder.Pdso_total
		var Pdsh_order_forward float64 = 0
		Pdsh_order_forward = productStoreOrder.Pdso_total + Pdso_exp
		var Pdsh_order_total float64 = 0
		Pdsh_order_total = Pdsh_order_forward - Pdso_exp
		Pdso_expire, _ := time.Parse(time.RFC3339, productStoreOrder.Pdso_expire)
		// ProductStore
		var processProductStore structs.ProcessProductStoreExpire
		models.GetProcessProductStoreExpire(productStoreOrder.Product_store_id, &processProductStore)
		var Pdsh_total_forward float64 = 0
		Pdsh_total_forward = processProductStore.Pds_total + Pdso_exp
		var Pdsh_total float64 = 0
		Pdsh_total = Pdsh_total_forward - Pdso_exp
		// Product
		var productId structs.ProductIdProcess
		models.GetProductIdProcess(processProductStore.Product_id, &productId)
		// ProductStoreHistorys
		productStoreHistoryUpdate := structs.ProcessProductStoreHistoryExp{
			Id:                      0,
			Shop_id:                 processProductStore.Shop_id,
			Shop_store_id:           processProductStore.Shop_store_id,
			Product_store_id:        productStoreOrder.Product_store_id,
			Product_store_order_id:  productStoreOrder.Id,
			Receipt_id:              nil,
			Receipt_detail_id:       nil,
			Check_product_id:        nil,
			Service_product_used_id: nil,
			Queue_id:                nil,
			Pdsh_in:                 0,
			Pdsh_out:                Pdso_exp,
			Pdsh_order_forward:      Pdsh_order_forward,
			Pdsh_total_forward:      Pdsh_total_forward,
			Pdsh_amount:             0 - Pdso_exp,
			Pdsh_order_total:        Pdsh_order_total,
			Pdsh_total:              Pdsh_total,
			Pdsh_inout:              2,
			Pdsh_out_id:             5, //ปะเภทการเคลื่อนไหวยาหรืออุปกรณ์ : 1 ขาย, 2 โอน, 3 ใช้บริการ, 4 ตรวจ, 5 หมดอายุ
			Pdsh_ref_doc_no:         "EXP",
			Product_id:              productId.Id,
			Pd_code:                 productId.Pd_code,
			Pd_name:                 productId.Pd_name,
			Pdsh_type_id:            17,
			Pdsh_comment:            "Process Product Expire",
			Pdsh_date:               Pdso_expire.Format("2006-01-02"),
			Pdsh_modify:             time.Now().Format("2006-01-02 15:04:05"),
		}
		models.CreateProductStoreHistoryExp(&productStoreHistoryUpdate)
	}
	c.JSON(200, gin.H{
		"status":  true,
		"message": "Process Product Expire Success.",
		"data":    "",
	})
}

func ProcessProductStoreOut(data *structs.ProcessProduct) int {
	var temp_amount float64 = data.Pdso_qty
	var amount float64 = data.Pdso_qty
	var expire string = ""
	for amount > 0 {
		temp_amount = amount
		var productStoreOrder structs.ProcessProductStoreOrder
		models.GetProcessProductStoreOrder(data.Product_id, data.Product_store_id, &productStoreOrder)
		if productStoreOrder.Id != 0 {
			fmt.Println(productStoreOrder.Id)
			if productStoreOrder.Pdso_total < amount {
				// fmt.Println(productStoreOrder.Pdso_total)
				// fmt.Println(amount)
				amount = amount - productStoreOrder.Pdso_total
				temp_amount = temp_amount - amount
				// fmt.Println(amount)
				// fmt.Println(temp_amount)
				// fmt.Println(1)
				// ProductStoreOrder
				var processProductStoreOrder structs.ProcessProductStoreOrder
				models.GetProcessProductStoreOrderId(productStoreOrder.Id, &processProductStoreOrder)
				var Pdso_out float64 = 0
				Pdso_out = productStoreOrder.Pdso_out + temp_amount
				var Pdso_total float64 = 0
				Pdso_total = productStoreOrder.Pdso_total - temp_amount
				if productStoreOrder.Pdso_expire < expire || expire == "" {
					expire = productStoreOrder.Pdso_expire
				}
				productStoreOrderUpdate := structs.ProcessProductStoreOrder{
					Id:               productStoreOrder.Id,
					Product_store_id: productStoreOrder.Product_store_id,
					Pdso_expire:      productStoreOrder.Pdso_expire,
					Pdso_out:         Pdso_out,
					Pdso_use:         productStoreOrder.Pdso_use,
					Pdso_total:       Pdso_total,
					Pdso_update:      time.Now().Format("2006-01-02 15:04:05"),
				}
				models.UpdateProductStoreOrders(productStoreOrder.Id, &productStoreOrderUpdate)
				// ProductStore
				var processProductStore structs.ProcessProductStore
				models.GetProcessProductStoreId(data.Product_store_id, &processProductStore)
				var Pds_out float64 = 0
				Pds_out = processProductStore.Pds_out + temp_amount
				var Pds_total float64 = 0
				Pds_total = processProductStore.Pds_total - temp_amount
				productStoreUpdate := structs.ProcessProductStore{
					Id:            data.Product_store_id,
					Shop_store_id: processProductStore.Shop_store_id,
					Pds_out:       Pds_out,
					Pds_total:     Pds_total,
					Pds_update:    time.Now().Format("2006-01-02 15:04:05"),
				}
				models.UpdateProductStore(data.Product_store_id, &productStoreUpdate)
				// ProductStoreHistorys
				productStoreHistoryUpdate := structs.ProcessProductStoreHistory{
					Id:                      0,
					Shop_id:                 data.Shop_id,
					Shop_store_id:           processProductStore.Shop_store_id,
					Product_store_id:        data.Product_store_id,
					Product_store_order_id:  processProductStoreOrder.Id,
					Receipt_id:              data.Receipt_id,
					Receipt_detail_id:       data.Receipt_detail_id,
					Check_product_id:        data.Check_product_id,
					Service_product_used_id: data.Service_product_used_id,
					Queue_id:                data.Queue_id,
					Pdsh_in:                 0,
					Pdsh_out:                temp_amount,
					Pdsh_order_forward:      productStoreOrder.Pdso_total,
					Pdsh_total_forward:      processProductStore.Pds_total,
					Pdsh_amount:             0 - temp_amount,
					Pdsh_order_total:        Pdso_total,
					Pdsh_total:              Pds_total,
					Pdsh_inout:              2,
					Pdsh_out_id:             data.Pdsh_out_id,
					Pdsh_ref_doc_no:         data.Pdsh_ref_doc_no,
					User_id:                 data.User_id,
					Product_id:              data.Product_id,
					Pd_code:                 data.Recd_code,
					Pd_name:                 data.Recd_name,
					Pdsh_type_id:            data.Pdsh_type_id,
					Customer_id:             data.Customer_id,
					Ctm_prefix:              data.Ctm_prefix,
					Ctm_fname:               data.Ctm_fname,
					Ctm_lname:               data.Ctm_lname,
					Ctm_fname_en:            data.Ctm_fname_en,
					Ctm_lname_en:            data.Ctm_lname_en,
					Ctm_tel:                 data.Ctm_tel,
					Ctm_gender:              data.Ctm_gender,
					Pdsh_comment:            "Process Product",
					Pdsh_date:               time.Now().Format("2006-01-02"),
					Pdsh_modify:             time.Now().Format("2006-01-02 15:04:05"),
				}
				models.CreateProductStoreHistory(&productStoreHistoryUpdate)

			} else {
				// fmt.Println(productStoreOrder.Pdso_total)
				// fmt.Println(amount)
				amount = amount - productStoreOrder.Pdso_total
				if amount < 0 {
					amount = 0
				}
				temp_amount = temp_amount - amount
				// fmt.Println(amount)
				// fmt.Println(temp_amount)
				// fmt.Println(2)
				// ProductStoreOrder
				var processProductStoreOrder structs.ProcessProductStoreOrder
				models.GetProcessProductStoreOrderId(productStoreOrder.Id, &processProductStoreOrder)
				var Pdso_out float64 = 0
				Pdso_out = productStoreOrder.Pdso_out + temp_amount
				var Pdso_total float64 = 0
				Pdso_total = productStoreOrder.Pdso_total - temp_amount
				if productStoreOrder.Pdso_expire < expire || expire == "" {
					expire = productStoreOrder.Pdso_expire
				}
				productStoreOrderUpdate := structs.ProcessProductStoreOrder{
					Id:               productStoreOrder.Id,
					Product_store_id: productStoreOrder.Product_store_id,
					Pdso_expire:      productStoreOrder.Pdso_expire,
					Pdso_out:         Pdso_out,
					Pdso_use:         productStoreOrder.Pdso_use,
					Pdso_total:       Pdso_total,
					Pdso_update:      time.Now().Format("2006-01-02 15:04:05"),
				}
				models.UpdateProductStoreOrders(productStoreOrder.Id, &productStoreOrderUpdate)
				// ProductStore
				var processProductStore structs.ProcessProductStore
				models.GetProcessProductStoreId(data.Product_store_id, &processProductStore)
				var Pds_out float64 = 0
				Pds_out = processProductStore.Pds_out + temp_amount
				var Pds_total float64 = 0
				Pds_total = processProductStore.Pds_total - temp_amount
				productStoreUpdate := structs.ProcessProductStore{
					Id:            data.Product_store_id,
					Shop_store_id: processProductStore.Shop_store_id,
					Pds_out:       Pds_out,
					Pds_total:     Pds_total,
					Pds_update:    time.Now().Format("2006-01-02 15:04:05"),
				}
				models.UpdateProductStore(data.Product_store_id, &productStoreUpdate)

				// ProductStoreHistorys
				productStoreHistoryUpdate := structs.ProcessProductStoreHistory{
					Id:                      0,
					Shop_id:                 data.Shop_id,
					Shop_store_id:           processProductStore.Shop_store_id,
					Product_store_id:        data.Product_store_id,
					Product_store_order_id:  processProductStoreOrder.Id,
					Receipt_id:              data.Receipt_id,
					Receipt_detail_id:       data.Receipt_detail_id,
					Check_product_id:        data.Check_product_id,
					Service_product_used_id: data.Service_product_used_id,
					Queue_id:                data.Queue_id,
					Pdsh_in:                 0,
					Pdsh_out:                temp_amount,
					Pdsh_order_forward:      productStoreOrder.Pdso_total,
					Pdsh_total_forward:      processProductStore.Pds_total,
					Pdsh_amount:             0 - temp_amount,
					Pdsh_order_total:        Pdso_total,
					Pdsh_total:              Pds_total,
					Pdsh_inout:              2,
					Pdsh_out_id:             data.Pdsh_out_id,
					Pdsh_ref_doc_no:         data.Pdsh_ref_doc_no,
					User_id:                 data.User_id,
					Product_id:              data.Product_id,
					Pd_code:                 data.Recd_code,
					Pd_name:                 data.Recd_name,
					Pdsh_type_id:            data.Pdsh_type_id,
					Customer_id:             data.Customer_id,
					Ctm_prefix:              data.Ctm_prefix,
					Ctm_fname:               data.Ctm_fname,
					Ctm_lname:               data.Ctm_lname,
					Ctm_fname_en:            data.Ctm_fname_en,
					Ctm_lname_en:            data.Ctm_lname_en,
					Ctm_tel:                 data.Ctm_tel,
					Ctm_gender:              data.Ctm_gender,
					Pdsh_comment:            "Process Product",
					Pdsh_date:               time.Now().Format("2006-01-02"),
					Pdsh_modify:             time.Now().Format("2006-01-02 15:04:05"),
				}
				models.CreateProductStoreHistory(&productStoreHistoryUpdate)
			}
		} else {
			// fmt.Println(productStoreOrder.Pdso_total)
			// fmt.Println(amount)
			amount = 0
			// fmt.Println(amount)
			// fmt.Println(temp_amount)
			// fmt.Println(3)
			// ProductStoreOrder
			var productStoreOrder structs.ProcessProductStoreOrder
			models.GetProcessProductStoreOrderLast(data.Product_id, data.Product_store_id, &productStoreOrder)
			// var processProductStoreOrder structs.ProcessProductStoreOrder
			// models.GetProcessProductStoreOrderId(productStoreOrder.Id, &processProductStoreOrder)
			var Pdso_out float64 = 0
			Pdso_out = productStoreOrder.Pdso_out + temp_amount
			var Pdso_total float64 = 0
			Pdso_total = productStoreOrder.Pdso_total - temp_amount
			if productStoreOrder.Pdso_expire < expire || expire == "" {
				expire = productStoreOrder.Pdso_expire
			}
			productStoreOrderUpdate := structs.ProcessProductStoreOrder{
				Id:               productStoreOrder.Id,
				Product_store_id: productStoreOrder.Product_store_id,
				Pdso_expire:      productStoreOrder.Pdso_expire,
				Pdso_out:         Pdso_out,
				Pdso_use:         productStoreOrder.Pdso_use,
				Pdso_total:       Pdso_total,
				Pdso_update:      time.Now().Format("2006-01-02 15:04:05"),
			}
			models.UpdateProductStoreOrders(productStoreOrder.Id, &productStoreOrderUpdate)

			// ProductStore
			var processProductStore structs.ProcessProductStore
			models.GetProcessProductStoreId(data.Product_store_id, &processProductStore)
			var Pds_out float64 = 0
			Pds_out = processProductStore.Pds_out + temp_amount
			var Pds_total float64 = 0
			Pds_total = processProductStore.Pds_total - temp_amount
			productStoreUpdate := structs.ProcessProductStore{
				Id:            data.Product_store_id,
				Shop_store_id: processProductStore.Shop_store_id,
				Pds_out:       Pds_out,
				Pds_total:     Pds_total,
				Pds_update:    time.Now().Format("2006-01-02 15:04:05"),
			}
			models.UpdateProductStore(data.Product_store_id, &productStoreUpdate)
			// ProductStoreHistorys
			productStoreHistoryUpdate := structs.ProcessProductStoreHistory{
				Id:                      0,
				Shop_id:                 data.Shop_id,
				Shop_store_id:           processProductStore.Shop_store_id,
				Product_store_id:        data.Product_store_id,
				Product_store_order_id:  productStoreOrder.Id,
				Receipt_id:              data.Receipt_id,
				Receipt_detail_id:       data.Receipt_detail_id,
				Check_product_id:        data.Check_product_id,
				Service_product_used_id: data.Service_product_used_id,
				Queue_id:                data.Queue_id,
				Pdsh_in:                 0,
				Pdsh_out:                temp_amount,
				Pdsh_order_forward:      productStoreOrder.Pdso_total,
				Pdsh_total_forward:      processProductStore.Pds_total,
				Pdsh_amount:             0 - temp_amount,
				Pdsh_order_total:        Pdso_total,
				Pdsh_total:              Pds_total,
				Pdsh_inout:              2,
				Pdsh_out_id:             data.Pdsh_out_id,
				Pdsh_ref_doc_no:         data.Pdsh_ref_doc_no,
				User_id:                 data.User_id,
				Product_id:              data.Product_id,
				Pd_code:                 data.Recd_code,
				Pd_name:                 data.Recd_name,
				Pdsh_type_id:            data.Pdsh_type_id,
				Customer_id:             data.Customer_id,
				Ctm_prefix:              data.Ctm_prefix,
				Ctm_fname:               data.Ctm_fname,
				Ctm_lname:               data.Ctm_lname,
				Ctm_fname_en:            data.Ctm_fname_en,
				Ctm_lname_en:            data.Ctm_lname_en,
				Ctm_tel:                 data.Ctm_tel,
				Ctm_gender:              data.Ctm_gender,
				Pdsh_comment:            "Process Product",
				Pdsh_date:               time.Now().Format("2006-01-02"),
				Pdsh_modify:             time.Now().Format("2006-01-02 15:04:05"),
			}
			models.CreateProductStoreHistory(&productStoreHistoryUpdate)
		}
	}
	// Sticker ปริ้นสติ๊กเกอร์ตอนชำระเงิน
	// var productRateUnit structs.ProductRate
	// models.GetProductRate1(data.Shop_id, data.Product_id, &productRateUnit)
	// if productRateUnit.Pd_type_id == 1 {
	// 	var ProcessInvoiceDetailById structs.ProcessInvoiceDetailById
	// 	if data.Receipt_id != nil && data.Receipt_detail_id != nil {
	// 		models.GetProcessInvoiceDetailById(*data.Receipt_id, *data.Receipt_detail_id, &ProcessInvoiceDetailById)
	// 	}
	// 	productStickerUpdate := structs.ProcessProductSticker{
	// 		Id:                  0,
	// 		Product_id:          data.Product_id,
	// 		Customer_id:         data.Customer_id,
	// 		User_id:             data.User_id,
	// 		Invoice_id:          &ProcessInvoiceDetailById.Invoice_id,
	// 		Invoice_detail_id:   &ProcessInvoiceDetailById.Invoice_detail_id,
	// 		Receipt_id:          data.Receipt_id,
	// 		Receipt_detail_id:   data.Receipt_detail_id,
	// 		Sticker_code:        data.Recd_code,
	// 		Sticker_name:        data.Recd_name,
	// 		Sticker_name_acc:    productRateUnit.Pd_name_acc,
	// 		Sticker_amount:      data.Pdso_qty,
	// 		Sticker_unit:        productRateUnit.U_name,
	// 		Sticker_unit_en:     productRateUnit.U_name_en,
	// 		Sticker_price:       data.Recd_price,
	// 		Sticker_expdate:     expire,
	// 		Sticker_active_id:   1,
	// 		Sticker_print_label: 0,
	// 		Sticker_print_order: 0,
	// 		Sticker_topical:     data.Recd_topical,
	// 		Sticker_direction:   data.Recd_direction,
	// 		Sticker_is_del:      0,
	// 		Sticker_modify:      time.Now().Format("2006-01-02 15:04:05"),
	// 	}
	// 	models.CreateProductSticker(&productStickerUpdate)
	// }
	return 1
}

func ProcessProductStoreUsed(data *structs.ProcessProduct) int {
	var temp_amount float64 = data.Pdso_qty
	var amount float64 = data.Pdso_qty
	var expire string = ""
	for amount > 0 {
		temp_amount = amount
		var productStoreOrder structs.ProcessProductStoreOrder
		models.GetProcessProductStoreOrder(data.Product_id, data.Product_store_id, &productStoreOrder)
		if productStoreOrder.Id != 0 {
			if productStoreOrder.Pdso_total < amount {
				// fmt.Println(productStoreOrder.Pdso_total)
				// fmt.Println(amount)
				amount = amount - productStoreOrder.Pdso_total
				temp_amount = temp_amount - amount
				// fmt.Println(amount)
				// fmt.Println(temp_amount)
				// ProductStoreOrder
				var processProductStoreOrder structs.ProcessProductStoreOrder
				models.GetProcessProductStoreOrderId(productStoreOrder.Id, &processProductStoreOrder)
				var Pdso_use float64 = 0
				Pdso_use = productStoreOrder.Pdso_use + temp_amount
				var Pdso_total float64 = 0
				Pdso_total = productStoreOrder.Pdso_total - temp_amount
				if productStoreOrder.Pdso_expire < expire || expire == "" {
					expire = productStoreOrder.Pdso_expire
				}
				productStoreOrderUpdate := structs.ProcessProductStoreOrder{
					Id:               productStoreOrder.Id,
					Product_store_id: productStoreOrder.Product_store_id,
					Pdso_expire:      productStoreOrder.Pdso_expire,
					Pdso_use:         Pdso_use,
					Pdso_out:         productStoreOrder.Pdso_out,
					Pdso_total:       Pdso_total,
					Pdso_update:      time.Now().Format("2006-01-02 15:04:05"),
				}
				models.UpdateProductStoreOrders(productStoreOrder.Id, &productStoreOrderUpdate)
				// ProductStore
				var processProductStore structs.ProcessProductStore
				models.GetProcessProductStoreId(data.Product_store_id, &processProductStore)
				var Pds_out float64 = 0
				Pds_out = processProductStore.Pds_out + temp_amount
				var Pds_total float64 = 0
				Pds_total = processProductStore.Pds_total - temp_amount
				productStoreUpdate := structs.ProcessProductStore{
					Id:            data.Product_store_id,
					Shop_store_id: processProductStore.Shop_store_id,
					Pds_out:       Pds_out,
					Pds_total:     Pds_total,
					Pds_update:    time.Now().Format("2006-01-02 15:04:05"),
				}
				models.UpdateProductStore(data.Product_store_id, &productStoreUpdate)
				// ProductStoreHistorys
				productStoreHistoryUpdate := structs.ProcessProductStoreHistory{
					Id:                      0,
					Shop_id:                 data.Shop_id,
					Shop_store_id:           processProductStore.Shop_store_id,
					Product_store_id:        data.Product_store_id,
					Product_store_order_id:  processProductStoreOrder.Id,
					Receipt_id:              data.Receipt_id,
					Receipt_detail_id:       data.Receipt_detail_id,
					Check_product_id:        data.Check_product_id,
					Service_product_used_id: data.Service_product_used_id,
					Queue_id:                data.Queue_id,
					Pdsh_in:                 0,
					Pdsh_out:                temp_amount,
					Pdsh_order_forward:      productStoreOrder.Pdso_total,
					Pdsh_total_forward:      processProductStore.Pds_total,
					Pdsh_amount:             0 - temp_amount,
					Pdsh_order_total:        Pdso_total,
					Pdsh_total:              Pds_total,
					Pdsh_inout:              2,
					Pdsh_out_id:             data.Pdsh_out_id,
					Pdsh_ref_doc_no:         data.Pdsh_ref_doc_no,
					User_id:                 data.User_id,
					Product_id:              data.Product_id,
					Pd_code:                 data.Recd_code,
					Pd_name:                 data.Recd_name,
					Pdsh_type_id:            data.Pdsh_type_id,
					Customer_id:             data.Customer_id,
					Ctm_prefix:              data.Ctm_prefix,
					Ctm_fname:               data.Ctm_fname,
					Ctm_lname:               data.Ctm_lname,
					Ctm_fname_en:            data.Ctm_fname_en,
					Ctm_lname_en:            data.Ctm_lname_en,
					Ctm_tel:                 data.Ctm_tel,
					Ctm_gender:              data.Ctm_gender,
					Pdsh_comment:            "Process Product",
					Pdsh_date:               time.Now().Format("2006-01-02"),
					Pdsh_modify:             time.Now().Format("2006-01-02 15:04:05"),
				}
				models.CreateProductStoreHistory(&productStoreHistoryUpdate)

			} else {
				// fmt.Println(productStoreOrder.Pdso_total)
				// fmt.Println(amount)
				amount = amount - productStoreOrder.Pdso_total
				if amount < 0 {
					amount = 0
				}
				temp_amount = temp_amount - amount
				// fmt.Println(amount)
				// fmt.Println(temp_amount)
				// ProductStoreOrder
				var processProductStoreOrder structs.ProcessProductStoreOrder
				models.GetProcessProductStoreOrderId(productStoreOrder.Id, &processProductStoreOrder)
				var Pdso_use float64 = 0
				Pdso_use = productStoreOrder.Pdso_use + temp_amount
				var Pdso_total float64 = 0
				Pdso_total = productStoreOrder.Pdso_total - temp_amount
				if productStoreOrder.Pdso_expire < expire || expire == "" {
					expire = productStoreOrder.Pdso_expire
				}
				productStoreOrderUpdate := structs.ProcessProductStoreOrder{
					Id:               productStoreOrder.Id,
					Product_store_id: productStoreOrder.Product_store_id,
					Pdso_expire:      productStoreOrder.Pdso_expire,
					Pdso_use:         Pdso_use,
					Pdso_out:         productStoreOrder.Pdso_out,
					Pdso_total:       Pdso_total,
					Pdso_update:      time.Now().Format("2006-01-02 15:04:05"),
				}
				models.UpdateProductStoreOrders(productStoreOrder.Id, &productStoreOrderUpdate)
				// ProductStore
				var processProductStore structs.ProcessProductStore
				models.GetProcessProductStoreId(data.Product_store_id, &processProductStore)
				var Pds_out float64 = 0
				Pds_out = processProductStore.Pds_out + temp_amount
				var Pds_total float64 = 0
				Pds_total = processProductStore.Pds_total - temp_amount
				productStoreUpdate := structs.ProcessProductStore{
					Id:            data.Product_store_id,
					Shop_store_id: processProductStore.Shop_store_id,
					Pds_out:       Pds_out,
					Pds_total:     Pds_total,
					Pds_update:    time.Now().Format("2006-01-02 15:04:05"),
				}
				models.UpdateProductStore(data.Product_store_id, &productStoreUpdate)

				// ProductStoreHistorys
				productStoreHistoryUpdate := structs.ProcessProductStoreHistory{
					Id:                      0,
					Shop_id:                 data.Shop_id,
					Shop_store_id:           processProductStore.Shop_store_id,
					Product_store_id:        data.Product_store_id,
					Product_store_order_id:  processProductStoreOrder.Id,
					Receipt_id:              data.Receipt_id,
					Receipt_detail_id:       data.Receipt_detail_id,
					Check_product_id:        data.Check_product_id,
					Service_product_used_id: data.Service_product_used_id,
					Queue_id:                data.Queue_id,
					Pdsh_in:                 0,
					Pdsh_out:                temp_amount,
					Pdsh_order_forward:      productStoreOrder.Pdso_total,
					Pdsh_total_forward:      processProductStore.Pds_total,
					Pdsh_amount:             0 - temp_amount,
					Pdsh_order_total:        Pdso_total,
					Pdsh_total:              Pds_total,
					Pdsh_inout:              2,
					Pdsh_out_id:             data.Pdsh_out_id,
					Pdsh_ref_doc_no:         data.Pdsh_ref_doc_no,
					User_id:                 data.User_id,
					Product_id:              data.Product_id,
					Pd_code:                 data.Recd_code,
					Pd_name:                 data.Recd_name,
					Pdsh_type_id:            data.Pdsh_type_id,
					Customer_id:             data.Customer_id,
					Ctm_prefix:              data.Ctm_prefix,
					Ctm_fname:               data.Ctm_fname,
					Ctm_lname:               data.Ctm_lname,
					Ctm_fname_en:            data.Ctm_fname_en,
					Ctm_lname_en:            data.Ctm_lname_en,
					Ctm_tel:                 data.Ctm_tel,
					Ctm_gender:              data.Ctm_gender,
					Pdsh_comment:            "Process Product",
					Pdsh_date:               time.Now().Format("2006-01-02"),
					Pdsh_modify:             time.Now().Format("2006-01-02 15:04:05"),
				}
				models.CreateProductStoreHistory(&productStoreHistoryUpdate)

			}
		} else {
			// fmt.Println(productStoreOrder.Pdso_total)
			// fmt.Println(amount)
			amount = 0
			// fmt.Println(amount)
			// fmt.Println(temp_amount)
			// ProductStoreOrder
			var productStoreOrder structs.ProcessProductStoreOrder
			models.GetProcessProductStoreOrderLast(data.Product_id, data.Product_store_id, &productStoreOrder)
			// var processProductStoreOrder structs.ProcessProductStoreOrder
			// models.GetProcessProductStoreOrderId(productStoreOrder.Id, &processProductStoreOrder)
			var Pdso_use float64 = 0
			Pdso_use = productStoreOrder.Pdso_use + temp_amount
			var Pdso_total float64 = 0
			Pdso_total = productStoreOrder.Pdso_total - temp_amount
			if productStoreOrder.Pdso_expire < expire || expire == "" {
				expire = productStoreOrder.Pdso_expire
			}
			productStoreOrderUpdate := structs.ProcessProductStoreOrder{
				Id:               productStoreOrder.Id,
				Product_store_id: productStoreOrder.Product_store_id,
				Pdso_expire:      productStoreOrder.Pdso_expire,
				Pdso_use:         Pdso_use,
				Pdso_out:         productStoreOrder.Pdso_out,
				Pdso_total:       Pdso_total,
				Pdso_update:      time.Now().Format("2006-01-02 15:04:05"),
			}
			models.UpdateProductStoreOrders(productStoreOrder.Id, &productStoreOrderUpdate)

			// ProductStore
			var processProductStore structs.ProcessProductStore
			models.GetProcessProductStoreId(data.Product_store_id, &processProductStore)
			var Pds_out float64 = 0
			Pds_out = processProductStore.Pds_out + temp_amount
			var Pds_total float64 = 0
			Pds_total = processProductStore.Pds_total - temp_amount
			productStoreUpdate := structs.ProcessProductStore{
				Id:            data.Product_store_id,
				Shop_store_id: processProductStore.Shop_store_id,
				Pds_out:       Pds_out,
				Pds_total:     Pds_total,
				Pds_update:    time.Now().Format("2006-01-02 15:04:05"),
			}
			models.UpdateProductStore(data.Product_store_id, &productStoreUpdate)
			// ProductStoreHistorys
			productStoreHistoryUpdate := structs.ProcessProductStoreHistory{
				Id:                      0,
				Shop_id:                 data.Shop_id,
				Shop_store_id:           processProductStore.Shop_store_id,
				Product_store_id:        data.Product_store_id,
				Product_store_order_id:  productStoreOrder.Id,
				Receipt_id:              data.Receipt_id,
				Receipt_detail_id:       data.Receipt_detail_id,
				Check_product_id:        data.Check_product_id,
				Service_product_used_id: data.Service_product_used_id,
				Queue_id:                data.Queue_id,
				Pdsh_in:                 0,
				Pdsh_out:                temp_amount,
				Pdsh_order_forward:      productStoreOrder.Pdso_total,
				Pdsh_total_forward:      processProductStore.Pds_total,
				Pdsh_amount:             0 - temp_amount,
				Pdsh_order_total:        Pdso_total,
				Pdsh_total:              Pds_total,
				Pdsh_inout:              2,
				Pdsh_out_id:             data.Pdsh_out_id,
				Pdsh_ref_doc_no:         data.Pdsh_ref_doc_no,
				User_id:                 data.User_id,
				Product_id:              data.Product_id,
				Pd_code:                 data.Recd_code,
				Pd_name:                 data.Recd_name,
				Pdsh_type_id:            data.Pdsh_type_id,
				Customer_id:             data.Customer_id,
				Ctm_prefix:              data.Ctm_prefix,
				Ctm_fname:               data.Ctm_fname,
				Ctm_lname:               data.Ctm_lname,
				Ctm_fname_en:            data.Ctm_fname_en,
				Ctm_lname_en:            data.Ctm_lname_en,
				Ctm_tel:                 data.Ctm_tel,
				Ctm_gender:              data.Ctm_gender,
				Pdsh_comment:            "Process Product",
				Pdsh_date:               time.Now().Format("2006-01-02"),
				Pdsh_modify:             time.Now().Format("2006-01-02 15:04:05"),
			}
			models.CreateProductStoreHistory(&productStoreHistoryUpdate)
		}
	}
	// Sticker
	var productRateUnit structs.ProductRate
	models.GetProductRate1(data.Shop_id, data.Product_id, &productRateUnit)
	if productRateUnit.Pd_type_id == 1 {
		if data.Service_product_used_id != nil {
			var ProcessInvoiceDetailById structs.ProcessInvoiceDetailById
			if data.Receipt_id != nil && data.Receipt_detail_id != nil {
				models.GetProcessInvoiceDetailById(*data.Receipt_id, *data.Receipt_detail_id, &ProcessInvoiceDetailById)
			}
			productStickerUpdate := structs.ProcessProductSticker{
				Id:                  0,
				Shop_id:             data.Shop_id,
				Product_id:          data.Product_id,
				Customer_id:         data.Customer_id,
				User_id:             data.User_id,
				Invoice_id:          &ProcessInvoiceDetailById.Invoice_id,
				Invoice_detail_id:   &ProcessInvoiceDetailById.Invoice_detail_id,
				Receipt_id:          data.Receipt_id,
				Receipt_detail_id:   data.Receipt_detail_id,
				Sticker_code:        data.Recd_code,
				Sticker_name:        data.Recd_name,
				Sticker_name_acc:    productRateUnit.Pd_name_acc,
				Sticker_amount:      data.Pdso_qty,
				Sticker_unit:        productRateUnit.U_name,
				Sticker_unit_en:     productRateUnit.U_name_en,
				Sticker_price:       data.Recd_price,
				Sticker_expdate:     expire,
				Sticker_active_id:   1,
				Sticker_print_label: 0,
				Sticker_print_order: 0,
				Sticker_topical:     data.Recd_topical,
				Sticker_direction:   data.Recd_direction,
				Sticker_is_del:      0,
				Sticker_modify:      time.Now().Format("2006-01-02 15:04:05"),
			}
			models.CreateProductSticker(&productStickerUpdate)
			var ObjUserEmail structs.ObjUserEmail
			models.GetUserEmail(data.User_id, &ObjUserEmail)
			models.AddLogStk(&structs.LogStk{
				Username:   ObjUserEmail.User_email,
				Log_type:   "Add Sticker Receipt",
				Log_text:   "Add Sticker Name " + data.Recd_name,
				Log_create: time.Now().Format("2006-01-02 15:04:05"),
				Shop_id:    data.Shop_id,
			})
		}
	}
	return 1
}

func ProductSticker(data *structs.StickerProcessProduct) int {
	var temp_amount float64 = data.Pdso_qty
	var amount float64 = data.Pdso_qty
	var expire string = ""
	for amount > 0 {
		temp_amount = amount
		var productStoreOrder structs.ProcessProductStoreOrder
		models.GetProcessProductStoreOrder(data.Product_id, data.Product_store_id, &productStoreOrder)
		if productStoreOrder.Id != 0 {
			if productStoreOrder.Pdso_total < amount {
				amount = amount - productStoreOrder.Pdso_total
				temp_amount = temp_amount - amount
				var processProductStoreOrder structs.ProcessProductStoreOrder
				models.GetProcessProductStoreOrderId(productStoreOrder.Id, &processProductStoreOrder)
				if productStoreOrder.Pdso_expire < expire || expire == "" {
					expire = productStoreOrder.Pdso_expire
				}
			} else {
				amount = amount - productStoreOrder.Pdso_total
				if amount < 0 {
					amount = 0
				}
				temp_amount = temp_amount - amount
				var processProductStoreOrder structs.ProcessProductStoreOrder
				models.GetProcessProductStoreOrderId(productStoreOrder.Id, &processProductStoreOrder)
				if productStoreOrder.Pdso_expire < expire || expire == "" {
					expire = productStoreOrder.Pdso_expire
				}
			}
		} else {
			amount = 0
			var productStoreOrder structs.ProcessProductStoreOrder
			models.GetProcessProductStoreOrderLast(data.Product_id, data.Product_store_id, &productStoreOrder)
			if productStoreOrder.Pdso_expire < expire || expire == "" {
				expire = productStoreOrder.Pdso_expire
			}
		}
	}
	// Sticker
	var productRateUnit structs.ProductRate
	models.GetProductRate1(data.Shop_id, data.Product_id, &productRateUnit)
	if productRateUnit.Pd_type_id == 1 {
		productStickerUpdate := structs.ProcessProductStickerInvoice{
			Id:                  0,
			Product_id:          data.Product_id,
			Shop_id:             data.Shop_id,
			Customer_id:         data.Customer_id,
			User_id:             data.User_id,
			Invoice_id:          data.Invoice_id,
			Invoice_detail_id:   data.Invoice_detail_id,
			Receipt_id:          nil,
			Receipt_detail_id:   nil,
			Sticker_code:        data.Invd_code,
			Sticker_name:        data.Invd_name,
			Sticker_name_acc:    productRateUnit.Pd_name_acc,
			Sticker_amount:      data.Pdso_qty,
			Sticker_unit:        productRateUnit.U_name,
			Sticker_unit_en:     productRateUnit.U_name_en,
			Sticker_price:       data.Invd_price,
			Sticker_expdate:     expire,
			Sticker_active_id:   1,
			Sticker_print_label: 0,
			Sticker_print_order: 0,
			Sticker_topical:     data.Invd_topical,
			Sticker_direction:   data.Invd_direction,
			Sticker_is_del:      0,
			Sticker_modify:      time.Now().Format("2006-01-02 15:04:05"),
		}
		models.CreateProductStickerInvoice(&productStickerUpdate)
		var ObjUserEmail structs.ObjUserEmail
		models.GetUserEmail(data.User_id, &ObjUserEmail)
		models.AddLogStk(&structs.LogStk{
			Username:   ObjUserEmail.User_email,
			Log_type:   "Add Sticker Invoice",
			Log_text:   "Add Sticker Name " + data.Invd_name,
			Log_create: time.Now().Format("2006-01-02 15:04:05"),
			Shop_id:    data.Shop_id,
		})
	}
	return 1
}
