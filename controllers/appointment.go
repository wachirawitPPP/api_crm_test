package controllers

import (
	"linecrmapi/libs"
	"linecrmapi/models"
	"linecrmapi/structs"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// func DoctorTable(c *gin.Context) {

// 	var filter structs.PayloadDoctorTable
// 	if errSBJ := c.ShouldBindJSON(&filter); errSBJ != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Invalid request data.",
// 			"data":    errSBJ,
// 		})
// 		return
// 	}

// 	Shop_id := libs.StrToInt(c.Params.ByName("shopId"))

// 	var ShopTimeset structs.ShopTimeset
// 	if errMD := models.GetTimesetConfig(Shop_id, &ShopTimeset); errMD != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Something went wrong.",
// 			"data":    "",
// 		})
// 		return
// 	}

// 	if ShopTimeset.TimesetRange < 1 {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Shop Not set Time Range.",
// 			"data":    "",
// 		})
// 		return
// 	}

// 	startTimeStr := ShopTimeset.TimesetOpen
// 	endTimeStr := ShopTimeset.TimesetClose
// 	rangTime := time.Duration(ShopTimeset.TimesetRange)

// 	// Parse start time and end time strings into time.Time objects
// 	startTime, _ := time.Parse("15:04:05", startTimeStr)
// 	endTimeM, _ := time.Parse("15:04:05", endTimeStr)

// 	currentRound := startTime
// 	roundDuration := rangTime * time.Minute
// 	endTime := endTimeM.Add(rangTime * time.Minute)

// 	var members []structs.UserSettimeToDay
// 	if errMD := models.GetUserTimesetDoctor(Shop_id, startTimeStr, endTimeStr, &members); errMD != nil || len(members) < 1 {
// 		members = []structs.UserSettimeToDay{}
// 	}

// 	arr := []structs.TableTime{}
// 	var wg sync.WaitGroup
// 	var mu sync.Mutex

// 	round := 0
// 	for currentRound.Before(endTime) {
// 		round++
// 		wg.Add(1)
// 		go func(st string, i int) {
// 			defer wg.Done()
// 			mu.Lock()
// 			var Timeline structs.TableTime
// 			Timeline.TimeData = currentRound.Format("15:04:05")
// 			currentRound = currentRound.Add(roundDuration)

// 			dateStr1 := filter.DateStart
// 			dateStr2 := filter.DateEnd

// 			// Parse the date strings into time.Time objects
// 			startDate, _ := time.Parse("2006-01-02", dateStr1)
// 			endDate, _ := time.Parse("2006-01-02", dateStr2)

// 			// Iterate through the range of dates
// 			Timeline.DayDatas = []structs.TableDays{}

// 			TimeS := Timeline.TimeData
// 			TimeE := currentRound.Format("15:04:05")
// 			for currentDate := startDate; currentDate.Before(endDate) || currentDate.Equal(endDate); currentDate = currentDate.AddDate(0, 0, 1) {
// 				dateStr := currentDate.Format("2006-01-02")
// 				var TDays structs.TableDays
// 				TDays.Date = dateStr
// 				dayOfWeek := currentDate.Weekday()
// 				TDays.Day = dayOfWeek.String()
// 				TDays.Day_th = convertToThaiDay(dayOfWeek.String())
// 				TStart, _ := time.Parse("15:04:05", TimeS)
// 				TEnd, _ := time.Parse("15:04:05", TimeE)

// 				for _, m := range members {
// 					var day_id time.Weekday
// 					if dayOfWeek == 0 {
// 						day_id = 7
// 					} else {
// 						day_id = dayOfWeek
// 					}
// 					TS, _ := time.Parse("15:04:05", m.Time_start)
// 					TE, _ := time.Parse("15:04:05", m.Time_end)
// 					if (TStart.After(TS) && TEnd.Before(TE) || TS.Equal(TStart) || TStart.Equal(TE) || TEnd.Equal(TE)) && int(day_id) == m.Day_id {
// 						TDays.Members = append(TDays.Members, m)
// 					}
// 				}
// 				if TDays.Members == nil {
// 					TDays.Members = []structs.UserSettimeToDay{}
// 				}
// 				Timeline.DayDatas = append(Timeline.DayDatas, TDays)
// 			}
// 			mu.Unlock()
// 			arr = append(arr, Timeline)
// 		}("start"+strconv.Itoa(round), round)
// 		wg.Wait()
// 	}

// 	c.JSON(200, gin.H{
// 		"status":  true,
// 		"message": "",
// 		"data":    arr,
// 	})
// }

// func DoctorAppointmentToday(c *gin.Context) {

// 	Shop_id := libs.StrToInt(c.Params.ByName("shopId"))
// 	currentTime := time.Now()
// 	// Format the current time as YYYY-MM-DD
// 	formattedTime := currentTime.Format("2006-01-02")

// 	var apList []structs.AppointmentList
// 	if errMD := models.GetAppointmentByDate(Shop_id, formattedTime+" 00:00:00", formattedTime+" 23:59:59", &apList); errMD != nil || len(apList) == 0 {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "No Data!",
// 			"data":    []string{},
// 		})
// 		return
// 	}
// 	c.JSON(200, gin.H{
// 		"status":  true,
// 		"message": "",
// 		"data":    apList,
// 	})
// }

// func DoctorAppointmentTimesetToday(c *gin.Context) {

// 	Shop_id := libs.StrToInt(c.Params.ByName("shopId"))
// 	currentTime := time.Now()
// 	dayOfWeek := currentTime.Weekday()

// 	var ShopTimeset structs.ShopTimeset
// 	if errMD := models.GetTimesetConfig(Shop_id, &ShopTimeset); errMD != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Something went wrong.",
// 			"data":    "",
// 		})
// 		return
// 	}

// 	TimesetOpen := ShopTimeset.TimesetOpen
// 	TimesetClose := ShopTimeset.TimesetClose

// 	var apList []structs.UserSettimeToDay
// 	var day_id time.Weekday
// 	if dayOfWeek == 0 {
// 		day_id = 7
// 	} else {
// 		day_id = dayOfWeek
// 	}
// 	if errMD := models.GetUserTimesetByDay(Shop_id, int(day_id), TimesetOpen, TimesetClose, &apList); errMD != nil || len(apList) == 0 {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "No Data!",
// 			"data":    []string{},
// 		})
// 		return
// 	}
// 	c.JSON(200, gin.H{
// 		"status":  true,
// 		"message": "",
// 		"data":    apList,
// 	})
// }

// func AppointmentSearch(c *gin.Context) {

// 	var filter structs.PayloadSearchAppointment
// 	if errSBJ := c.ShouldBindJSON(&filter); errSBJ != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Invalid request data.",
// 			"data":    errSBJ,
// 		})
// 		return
// 	}

// 	filter.Shop_id = libs.StrToInt(c.Params.ByName("shopId"))

// 	var countApList []structs.AppointmentList

// 	if filter.ActivePage < 1 {
// 		filter.ActivePage = 0
// 	} else {
// 		filter.ActivePage -= 1
// 	}

// 	var Is_ipd int = 0
// 	var package_ipd = 11

// 	var Is_tele int = 0
// 	var package_tele = 15

// 	var ShopIpd models.ShopAddon
// 	if errCKA := models.GetShopAddonId(filter.Shop_id, package_ipd, &ShopIpd); errCKA != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Package addon something went wrong.",
// 			"data":    errCKA,
// 			"is_ipd":  Is_ipd,
// 			"is_tele": Is_tele,
// 		})
// 		return
// 	}

// 	if ShopIpd.Id > 0 {
// 		Is_ipd = 1
// 	}

// 	var ShopTele models.ShopAddon
// 	if errCKA2 := models.GetShopAddonId(filter.Shop_id, package_tele, &ShopTele); errCKA2 != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Package addon something went wrong.",
// 			"data":    errCKA2,
// 			"is_ipd":  Is_ipd,
// 			"is_tele": Is_tele,
// 		})
// 		return
// 	}

// 	if ShopTele.Id > 0 {
// 		Is_tele = 1
// 	}

// 	err := models.AppointmentSearch(filter, false, &countApList)
// 	emptySlice := []string{}
// 	if err != nil {
// 		c.JSON(200, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 			"data":    emptySlice,
// 			"is_ipd":  Is_ipd,
// 			"is_tele": Is_tele,
// 		})
// 		c.Abort()
// 	} else {
// 		var apList []structs.AppointmentList
// 		if errMD := models.AppointmentSearch(filter, true, &apList); errMD != nil {
// 			c.AbortWithStatusJSON(200, gin.H{
// 				"status":  false,
// 				"message": "Something went wrong.",
// 				"data":    "",
// 				"is_ipd":  Is_ipd,
// 				"is_tele": Is_tele,
// 			})
// 			return
// 		}

// 		if len(apList) == 0 {
// 			c.JSON(200, gin.H{
// 				"status":  true,
// 				"message": "",
// 				"data": models.ResponsePaginationEmpty{
// 					Result_data:   emptySlice,
// 					Count_of_page: 0,
// 					Count_all:     0,
// 				},
// 				"is_ipd":  Is_ipd,
// 				"is_tele": Is_tele,
// 			})
// 			return
// 		}

// 		res := structs.ResponseAppointment{
// 			Result_data:   apList,
// 			Count_of_page: len(apList),
// 			Count_all:     len(countApList),
// 		}

// 		c.JSON(200, gin.H{
// 			"status":  true,
// 			"message": "",
// 			"data":    res,
// 			"is_ipd":  Is_ipd,
// 			"is_tele": Is_tele,
// 		})
// 	}
// }

func AppointmentCalendar(c *gin.Context) {

	var filter structs.PayloadCalendarAppointment
	if errSBJ := c.ShouldBindJSON(&filter); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errSBJ,
		})
		return
	}

	customerId := libs.StrToInt(c.Params.ByName("customerId"))
	ShopId := libs.StrToInt(c.Params.ByName("shopId"))

	var apList []structs.AppointmentList

	if errMD := models.AppointmentSearchCalendar(customerId, ShopId, filter, &apList); errMD != nil {
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
			"data":    []string{},
		})
		return
	}

	var cv []structs.CalendarView
	// var wg sync.WaitGroup
	// for i, st := range apList {
	// 	wg.Add(1)
	// 	go func(st string, i int) {
	// 		defer wg.Done()
	// 		date, _ := time.Parse(time.RFC3339, apList[i].ApDatetime)
	// 		data := structs.CalendarView{
	// 			Id:              apList[i].ID,
	// 			Title:           apList[i].ApTopic,
	// 			Start:           date.Format("2006-01-02") + "T" + apList[i].ApStart,
	// 			End:             date.Format("2006-01-02") + "T" + apList[i].ApEnd,
	// 			BackgroundColor: apList[i].ApColor,
	// 			BorderColor:     apList[i].ApColor,
	// 		}
	// 		cv = append(cv, data)
	// 	}(st.ApCreate, i)
	// }
	// wg.Wait()
	for i, _ := range apList {
		date, _ := time.Parse(time.RFC3339, apList[i].ApDatetime)
		data := structs.CalendarView{
			Id:              apList[i].ID,
			Title:           apList[i].ApTopic,
			Start:           date.Format("2006-01-02T15:04:05"),
			End:             date.Format("2006-01-02T15:04:05"),
			BackgroundColor: apList[i].ApColor,
			BorderColor:     apList[i].ApColor,
			Color:           apList[i].ApColor,
		}
		cv = append(cv, data)
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    cv,
	})
}

func GetAppointmentList(c *gin.Context) {
	customerId := libs.StrToInt(c.Params.ByName("customerId"))
	ShopId := libs.StrToInt(c.Params.ByName("shopId"))

	var apDetail []structs.AppointmentDetail
	err := models.GetAppointmentDetailList(customerId, ShopId, &apDetail)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Invalid Data!",
			"data":    []structs.AppointmentDetail{},
		})
		return
	} else {

		c.JSON(200, gin.H{
			"status":  true,
			"message": "",
			"data":    apDetail,
		})
	}
}

// func GetAppointmentDetail(c *gin.Context) {
// 	apId := libs.StrToInt(c.Params.ByName("apId"))
// 	ShopId := libs.StrToInt(c.Params.ByName("shopId"))
// 	var apDetail structs.AppointmentDetail
// 	err := models.GetAppointmentDetail(apId, ShopId, &apDetail)
// 	if err != nil {
// 		c.JSON(200, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 			"data":    "",
// 		})
// 		return
// 	} else if apDetail.ID == 0 {
// 		c.JSON(200, gin.H{
// 			"status":  false,
// 			"message": "",
// 			"data":    "",
// 		})
// 		return
// 	} else {
// 		var Tags []structs.AppointmentTags
// 		if errShop := models.GetAppointmentTagList(apId, &Tags); errShop != nil || len(Tags) == 0 {
// 			apDetail.ApTags = []structs.AppointmentTags{}
// 		} else {
// 			apDetail.ApTags = Tags
// 		}

// 		var RCShop structs.ReceiptShop
// 		if errShop := models.GetShopReceiptById(apDetail.ShopID, &RCShop); errShop != nil || RCShop.Id == 0 {
// 			apDetail.Shop = structs.ReceiptShop{}
// 		} else {
// 			apDetail.Shop = RCShop
// 		}

// 		c.JSON(200, gin.H{
// 			"status":  true,
// 			"message": "",
// 			"data":    apDetail,
// 		})
// 	}
// }

func isCurrentTimeBetween(start, end time.Time, currentTime time.Time) bool {
	if currentTime == start || currentTime == end {
		return true
	} else {
		return currentTime.After(start) && currentTime.Before(end)
	}
}

func AddAppointment(c *gin.Context) {
	var payload structs.ObjPayloadCreateAppointment
	if errPL := c.ShouldBindJSON(&payload); errPL != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errPL,
		})
		return
	}

	customerId := libs.StrToInt(c.Params.ByName("customerId"))
	shopId := libs.StrToInt(c.Params.ByName("shopId"))

	cus_fullname := ""
	cus_tel := ""

	if customerId > 0 {
		var objQueryCustomer structs.ObjQueryCustomer
		errCustomer := models.GetCustomerById(customerId, &objQueryCustomer)
		if errCustomer != nil {
			c.JSON(200, gin.H{
				"status":  false,
				"message": "Get customer error.",
				"data":    "",
			})
		}
		cus_fullname = objQueryCustomer.CtmFname + " " + objQueryCustomer.CtmLname
		cus_tel = objQueryCustomer.CtmTel
	}

	//get config shop
	var ShopTimeset structs.ShopTimeset
	if errMD := models.GetTimesetConfig(shopId, &ShopTimeset); errMD != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    "",
		})
		return
	}
	check_dup := 0
	if ShopTimeset.TimesetDayId == 1 {
		check_dup = ShopTimeset.TimesetDayAmount
	}

	// for _, item := range payload.ApDatetime {
	//check shop settime
	dateString := payload.ApDatetime
	parsedTime, errD := time.Parse("2006-01-02 15:04:05", dateString)
	if errD != nil {
		// fmt.Println("Error parsing date:", errD)
		return
	}
	dayOfWeek := parsedTime.Weekday()
	day_id := int(dayOfWeek)
	if int(dayOfWeek) == 0 {
		day_id = 7
	}

	shop_day := 0
	if ShopTimeset.TimesetSunday == 1 && int(dayOfWeek) == 0 {
		shop_day = 1
	}
	if ShopTimeset.TimesetMonday == 1 && int(dayOfWeek) == 1 {
		shop_day = 1
	}
	if ShopTimeset.TimesetTuesday == 1 && int(dayOfWeek) == 2 {
		shop_day = 1
	}
	if ShopTimeset.TimesetWednesday == 1 && int(dayOfWeek) == 3 {
		shop_day = 1
	}
	if ShopTimeset.TimesetThursday == 1 && int(dayOfWeek) == 4 {
		shop_day = 1
	}
	if ShopTimeset.TimesetFriday == 1 && int(dayOfWeek) == 5 {
		shop_day = 1
	}
	if ShopTimeset.TimesetSaturday == 1 && int(dayOfWeek) == 6 {
		shop_day = 1
	}

	if shop_day == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "ทางร้านไม่ได้กำหนดวันนี้",
		})
		return
	}

	//check shop holiday
	dateSelect := parsedTime.Format("2006-01-02")
	var ShopHoliday []structs.ShopHoliday
	if errSH := models.GetShopHoliday(shopId, dateSelect, &ShopHoliday); errSH != nil || len(ShopHoliday) == 1 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "วันนี้เป็นวันหยุด",
		})
		return
	}

	//check doctor settime
	var UserSettime []structs.UserSettimeList
	if errUS := models.GetUserSettime(payload.UserID, day_id, &UserSettime); errUS != nil || len(UserSettime) == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "คุณหมอไม่ได้กำหนดวันนี้",
		})
		return
	}
	if len(UserSettime) == 1 {
		date_string := parsedTime.Format("2006-01-02")
		startTime, _ := time.Parse("2006-01-02 15:04:05", date_string+" "+UserSettime[0].TimeStart)
		endTime, _ := time.Parse("2006-01-02 15:04:05", date_string+" "+UserSettime[0].TimeEnd)

		if isCurrentTimeBetween(startTime, endTime, parsedTime) {
			// fmt.Println("The current time is between the start and end times.")
		} else {
			// fmt.Println("The current time is outside the specified range.")
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "หมอไม่ได้กำหนดเวลานี้",
			})
			return
		}
	}

	//check doctor duplicate
	var AppointmentCheck []structs.AppointmentDetail
	if errAC := models.CheckAppointmentDupicate(payload.UserID, dateString, &AppointmentCheck); errAC != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Check Duplicate",
		})
		return
	}
	if len(AppointmentCheck) != 0 && len(AppointmentCheck) >= check_dup {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "การนัดพบแพทย์ครั้งนี้มีการนัดอยู่แล้ว",
		})
		return
	}

	ap := structs.AppointmentAction{
		ID:               0,
		ShopID:           shopId,
		UserID:           payload.UserID,
		CustomerID:       customerId,
		CustomerFullname: cus_fullname,
		ApType:           2,
		ApTopic:          payload.ApTopic,
		ApTel:            cus_tel,
		ApDatetime:       payload.ApDatetime,
		ApNote:           "เพิ่มนัดหมายจาก Line CRM",
		ApComment:        "",
		ApColor:          "#49C96D",
		ApConfirm:        0,
		ApStatusID:       1,
		ApStatusSMS:      0,
		ApStatusLine:     0,
		ApSms:            "",
		ApIsGcalendar:    0,
		ApGid:            "",
		ApUserID:         0,
		ApIsDel:          0,
		ApOpdType:        1,
		ApIsTele:         0,
		ApCreate:         time.Now().Format("2006-01-02 15:04:05"),
		ApUpdate:         time.Now().Format("2006-01-02 15:04:05"),
	}

	err := models.AddAppointment(&ap)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Cannot Add Appointment",
			"data":    err.Error(),
		})
		return
	}

	appointmentId := ap.ID

	models.AddLogAppointment(&structs.LogAppointment{
		Username:   c.Params.ByName("userEmail"),
		Log_type:   "Add",
		Log_text:   "Add CRM Appointment Id = " + strconv.Itoa(appointmentId),
		Log_create: time.Now().Format("2006-01-02 15:04:05"),
	})
	// }

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Add Appointment Success.",
		"data":    "",
	})
}

func DeleteAppointment(c *gin.Context) {
	shopId := libs.StrToInt(c.Params.ByName("shopId"))
	apId := libs.StrToInt(c.Params.ByName("apId"))
	var ap structs.AppointmentAction
	if errCO := models.GetAppointmentByID(apId, shopId, &ap); errCO != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Appointment Invalid.",
			"data":    errCO,
		})
		return
	}

	if ap.ID < 1 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Appointment Item Invalid.",
			"data":    apId,
		})
		return
	}

	ap.ApIsDel = 1
	if errAD := models.UpdateAppointment(shopId, ap.ID, &ap); errAD != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Cannot Remove Appointment",
			"data":    errAD,
		})
		return
	}

	models.AddLogAppointment(&structs.LogAppointment{
		Username:   c.Params.ByName("userEmail"),
		Log_type:   "Delete",
		Log_text:   "Delete CRM Appointment Id = " + strconv.Itoa(apId),
		Log_create: time.Now().Format("2006-01-02 15:04:05"),
	})

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Delete Appointment ID " + strconv.Itoa(apId) + " success",
		"data":    "",
	})

}

func GetAppointmentDoctor(c *gin.Context) {

	shopId := libs.StrToInt(c.Params.ByName("shopId"))

	var objResponse []structs.ObjResponseSearchAppointmentDoctor
	err := models.GetAppointmentDoctor(shopId, &objResponse)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Get data error.",
			"data":    "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Get data successful.",
			"data":    objResponse,
		})
	}
}

func GetAppointmentTopic(c *gin.Context) {
	shopId := libs.StrToInt(c.Params.ByName("shopId"))
	var at []structs.AppointmentTopic
	errDN := models.GetAppointmentTopic(shopId, &at)
	if errDN != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    errDN.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    at,
	})
}

func GetAppointmentTag(c *gin.Context) {
	shopId := libs.StrToInt(c.Params.ByName("shopId"))
	var at []structs.AppointmentTag
	errDN := models.GetAppointmentTag(shopId, &at)
	if errDN != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    errDN.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    at,
	})
}

// func GetShopConfig(c *gin.Context) {
// 	shopId := libs.StrToInt(c.Params.ByName("shopId"))
// 	var at []structs.AppointmentTag
// 	errDN := models.GetAppointmentTag(shopId, &at)
// 	if errDN != nil {
// 		c.JSON(200, gin.H{
// 			"status":  false,
// 			"message": "Something went wrong.",
// 			"data":    errDN.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"status":  true,
// 		"message": "",
// 		"data":    at,
// 	})
// }

func ShopTimeOpen(c *gin.Context) {
	shopId := libs.StrToInt(c.Params.ByName("shopId"))

	var ShopTimeset structs.ShopTimeset
	if errMD := models.GetTimesetConfig(shopId, &ShopTimeset); errMD != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    "",
		})
		return
	}

	startTimeStr := ShopTimeset.TimesetOpen
	endTimeStr := ShopTimeset.TimesetClose
	rangTime := time.Duration(ShopTimeset.TimesetRange)

	// Parse start time and end time strings into time.Time objects
	startTime, _ := time.Parse("15:04:05", startTimeStr)
	endTimeM, _ := time.Parse("15:04:05", endTimeStr)

	currentRound := startTime
	roundDuration := rangTime * time.Minute
	endTime := endTimeM.Add(rangTime * time.Minute)

	arr := []structs.TableTime{}
	// rounds := 0
	for currentRound.Before(endTime) {
		// rounds++
		// if rounds > 7 {
		// 	break
		// } else {
		var Timeline structs.TableTime
		Timeline.TimeData = currentRound.Format("15:04:05")
		currentRound = currentRound.Add(roundDuration)
		arr = append(arr, Timeline)
		// }
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    arr,
	})
}

// func GetAppointmentDetailByToken(c *gin.Context) {

// 	token := c.Params.ByName("token")
// 	str, errPr := libs.Decrypt(token)
// 	if errPr != nil {
// 		c.AbortWithStatusJSON(200, gin.H{
// 			"status":  false,
// 			"message": "Get medical cert invalid.",
// 			"data":    errPr.Error(),
// 		})
// 		return
// 	}
// 	obj := strings.Split(str, "&")
// 	apId := libs.StrToInt(obj[0])
// 	ShopId := libs.StrToInt(obj[1])

// 	var apDetail structs.AppointmentDetail
// 	err := models.GetAppointmentDetail(apId, ShopId, &apDetail)
// 	if err != nil {
// 		c.JSON(200, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 			"data":    "",
// 		})
// 		return
// 	} else if apDetail.ID == 0 {
// 		c.JSON(200, gin.H{
// 			"status":  false,
// 			"message": "",
// 			"data":    obj,
// 		})
// 		return
// 	} else {
// 		var RCShop structs.ReceiptShop
// 		if errShop := models.GetShopReceiptById(apDetail.ShopID, &RCShop); errShop != nil || RCShop.Id == 0 {
// 			apDetail.Shop = structs.ReceiptShop{}
// 		} else {
// 			apDetail.Shop = RCShop
// 		}

// 		var Tags []structs.AppointmentTags
// 		if errShop := models.GetAppointmentTagList(apId, &Tags); errShop != nil || len(Tags) == 0 {
// 			apDetail.ApTags = []structs.AppointmentTags{}
// 		} else {
// 			apDetail.ApTags = Tags
// 		}
// 		c.JSON(200, gin.H{
// 			"status":  true,
// 			"message": "",
// 			"data":    apDetail,
// 		})
// 	}
// }

// Function to convert English day names to Thai
func convertToThaiDay(day string) string {
	// Mapping of English to Thai day names
	thaiDayNames := map[string]string{
		"Monday":    "วันจันทร์",
		"Tuesday":   "วันอังคาร",
		"Wednesday": "วันพุธ",
		"Thursday":  "วันพฤหัสบดี",
		"Friday":    "วันศุกร์",
		"Saturday":  "วันเสาร์",
		"Sunday":    "วันอาทิตย์",
	}

	// Return the Thai day name if found, otherwise return the original English name
	if thaiName, ok := thaiDayNames[day]; ok {
		return thaiName
	}
	return day
}
