package controllers

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"linecrmapi/libs"
	"linecrmapi/middlewares"
	"linecrmapi/models"
	"linecrmapi/structs"
	"net/http"

	// "os"
	"strconv"
	"strings"

	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
)

func Checklineid(c *gin.Context) {

	var payload structs.PayloadChecklineid

	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	// get customer online by line id
	var customerOnline structs.CustomerOnline
	if errCO := models.GetCustomerOnlineByLineId(payload.Line_id, &customerOnline); errCO != nil {
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

	// get customer first in shop by citizen id
	var customer models.Customer
	if errC := models.GetCustomerFirstByCitizenId(customerOnline.Co_citizen_id, &customer); errC != nil {
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
			"message": "Customer Invalid!",
			"data":    "",
		})
		return
	}

	var coActk = middlewares.CreateAccessToken(customerOnline.ID, customer.ID, customer.ShopId, customer.ShopMotherId, customerOnline.Co_line_id, customerOnline.Co_citizen_id)

	//add log
	if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
		Line_id:        customerOnline.Co_line_id,
		Log_ip_address: c.ClientIP(),
		Log_browser:    c.GetHeader("User-Agent"),
		Log_text:       "Login by Line ID",
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

func Loginonline(c *gin.Context) {
	var payload structs.PayloadLoginOnline
	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	Co_email := strings.ToLower(payload.Co_email)
	var Co_password = Co_email + payload.Co_password
	Password := hashPassword(Co_password)

	var checkCustomerOnline structs.CustomerOnline
	if errCO := models.GetCustomerOnlineEmailPassword(payload.Co_email, Password, &checkCustomerOnline); errCO != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Customer Online Invalid!",
			"data":    errCO.Error(),
		})
		return
	}
	if checkCustomerOnline.ID == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Customer Online Invalid!",
			"data":    "",
		})
		return
	}

	// var coActk = middlewares.CreateAccessToken(checkCustomerOnline.ID, 0, 0, 0, checkCustomerOnline.Co_line_id, checkCustomerOnline.Co_citizen_id)

	// responseData := structs.ResponseOauth{
	// 	AccessToken: coActk,
	// }

	// c.JSON(200, gin.H{
	// 	"status":  true,
	// 	"message": "",
	// 	"data":    responseData,
	// })

	timeCreated := time.Now()
	timeExpire := timeCreated.Add(time.Duration(5) * time.Minute)

	randomNumber := rand.Intn(900000) + 100000

	//update customer online
	if errL := models.UpdateCustomerOnline(checkCustomerOnline.ID, &structs.CustomerOnlineUpdate{
		Co_otp:        strconv.Itoa(randomNumber),
		Co_otp_expire: timeExpire.Format("2006-01-02 15:04:05"),
		Co_update:     time.Now().Format("2006-01-02 15:04:05"),
	}); errL != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Update Customer Online Error.",
			"data":    "",
		})
		return
	}

	// send otp customer online -----------------
	s := libs.Sms{
		Msisdn:  checkCustomerOnline.Co_tel,
		Message: "(SMS OTP) APSX CRM-LINE Platform Clinic OTP " + strconv.Itoa(randomNumber),
	}
	res, _ := s.Send()
	if res.StatusCode != 201 {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "SMS Send Error.",
			"data":    "",
		})
		return
	}

	//add log
	if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
		Line_id:        "",
		Log_ip_address: c.ClientIP(),
		Log_browser:    c.GetHeader("User-Agent"),
		Log_text:       "Register Line & Update OTP",
		Log_create:     time.Now().Format("2006-01-02 15:04:05"),
	}); errL != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Create Log Customer Online OTP Error.",
			"data":    "",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":        true,
		"message":       "Login Customer Online Waiting OTP.",
		"data":          "OTP Online (otponline)",
		"co_tel":        checkCustomerOnline.Co_tel,
		"co_citizen_id": checkCustomerOnline.Co_citizen_id,
	})
	return
}

func LoginonlineTelCheck(c *gin.Context) {
	var payload structs.PayloadLoginTel
	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	var checkCustomerOnline structs.CustomerOnline
	if errCO := models.GetCustomerOnlineTelCheck(payload.Co_tel, &checkCustomerOnline); errCO != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Customer Online Invalid!",
			"data":    errCO.Error(),
		})
		return
	}
	if checkCustomerOnline.ID == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Customer Online Invalid!",
			"data":    "",
		})
		return
	}

	timeCreated := time.Now()
	timeExpire := timeCreated.Add(time.Duration(5) * time.Minute)

	randomNumber := rand.Intn(900000) + 100000

	//update customer online
	if errL := models.UpdateCustomerOnline(checkCustomerOnline.ID, &structs.CustomerOnlineUpdate{
		Co_otp:        strconv.Itoa(randomNumber),
		Co_otp_expire: timeExpire.Format("2006-01-02 15:04:05"),
		Co_update:     time.Now().Format("2006-01-02 15:04:05"),
	}); errL != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Update Customer Online Error.",
			"data":    "",
		})
		return
	}

	// send otp customer online -----------------
	s := libs.Sms{
		Msisdn:  checkCustomerOnline.Co_tel,
		Message: "(SMS OTP) APSX CRM-LINE Platform Clinic OTP " + strconv.Itoa(randomNumber),
	}
	res, _ := s.Send()
	if res.StatusCode != 201 {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "SMS Send Error.",
			"data":    "",
		})
		return
	}

	//add log
	if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
		Line_id:        "",
		Log_ip_address: c.ClientIP(),
		Log_browser:    c.GetHeader("User-Agent"),
		Log_text:       "Register Line & Update OTP",
		Log_create:     time.Now().Format("2006-01-02 15:04:05"),
	}); errL != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Create Log Customer Online OTP Error.",
			"data":    "",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Login Customer Online Waiting OTP.",
		"data":    "OTP Online (otponline)",
		"co_tel":  checkCustomerOnline.Co_tel,
	})
	return
}

func hashPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	b := h.Sum(nil)
	return string(base64.StdEncoding.EncodeToString(b))
}

func Login(c *gin.Context) {

	var payload structs.PayloadLogin

	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	// check line id
	// get customer online by line id
	var checkCustomerOnline structs.CustomerOnline
	if errCO := models.GetCustomerOnlineByLineId(payload.Line_id, &checkCustomerOnline); errCO != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Customer Online Invalid!",
			"data":    errCO.Error(),
		})
		return
	}
	if checkCustomerOnline.ID == 0 {

		// get customer online by line id
		var customerOnline structs.CustomerOnline
		if errCO := models.CheckCustomerOnlineLogin(payload.Line_id, payload.Citizen_id, payload.Tel, &customerOnline); errCO != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Error Get Customer Online Invalid!",
				"data":    errCO.Error(),
			})
			return
		}
		if customerOnline.ID == 0 {

			timeCreated := time.Now()
			timeExpire := timeCreated.Add(time.Duration(5) * time.Minute)

			randomNumber := rand.Intn(900000) + 100000

			//add customer online
			if errL := models.AddCustomerOnline(&structs.CustomerOnline{
				Co_line_name:  payload.Line_name,
				Co_line_email: payload.Line_email,
				Co_line_id:    payload.Line_id,
				Co_citizen_id: payload.Citizen_id,
				Co_tel:        payload.Tel,
				Co_otp:        strconv.Itoa(randomNumber),
				Co_otp_expire: timeExpire.Format("2006-01-02 15:04:05"),
				Co_is_active:  1,
				Co_is_del:     0,
				Co_create:     time.Now().Format("2006-01-02 15:04:05"),
				Co_update:     time.Now().Format("2006-01-02 15:04:05"),
			}); errL != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Create Customer Online Error.",
					"data":    "",
				})
				return
			}

			// send otp customer online -----------------
			s := libs.Sms{
				Msisdn:  payload.Tel,
				Message: "(SMS OTP) APSX CRM-LINE Platform Clinic OTP " + strconv.Itoa(randomNumber),
			}
			res, _ := s.Send()
			if res.StatusCode != 201 {
				c.JSON(http.StatusOK, gin.H{
					"status":  false,
					"message": "SMS Send Error.",
					"data":    "",
				})
				return
			}

			//add log
			if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
				Line_id:        customerOnline.Co_line_id,
				Log_ip_address: c.ClientIP(),
				Log_browser:    c.GetHeader("User-Agent"),
				Log_text:       "Register Line & Create OTP",
				Log_create:     time.Now().Format("2006-01-02 15:04:05"),
			}); errL != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Create Log Customer Online OTP Error.",
					"data":    "",
				})
				return
			}

			c.JSON(200, gin.H{
				"status":  true,
				"message": "Add Customer Online Success & Waiting OTP.",
				"data":    "",
			})
			return
		} else {

			timeCreated := time.Now()
			timeExpire := timeCreated.Add(time.Duration(5) * time.Minute)

			randomNumber := rand.Intn(900000) + 100000

			//update customer online
			if errL := models.UpdateCustomerOnline(customerOnline.ID, &structs.CustomerOnlineUpdate{
				// Co_citizen_id: payload.Citizen_id,
				// Co_tel:        payload.Tel,
				Co_otp:        strconv.Itoa(randomNumber),
				Co_otp_expire: timeExpire.Format("2006-01-02 15:04:05"),
				Co_update:     time.Now().Format("2006-01-02 15:04:05"),
			}); errL != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Update Customer Online Error.",
					"data":    "",
				})
				return
			}

			// send otp customer online -----------------
			s := libs.Sms{
				Msisdn:  payload.Tel,
				Message: "(SMS OTP) APSX CRM-LINE Platform Clinic OTP " + strconv.Itoa(randomNumber),
			}
			res, _ := s.Send()
			if res.StatusCode != 201 {
				c.JSON(http.StatusOK, gin.H{
					"status":  false,
					"message": "SMS Send Error.",
					"data":    "",
				})
				return
			}

			//add log
			if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
				Line_id:        customerOnline.Co_line_id,
				Log_ip_address: c.ClientIP(),
				Log_browser:    c.GetHeader("User-Agent"),
				Log_text:       "Register Line & Update OTP",
				Log_create:     time.Now().Format("2006-01-02 15:04:05"),
			}); errL != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Create Log Customer Online OTP Error.",
					"data":    "",
				})
				return
			}

			c.JSON(200, gin.H{
				"status":  true,
				"message": "Update Customer Online Waiting OTP.",
				"data":    "",
			})
			return
		}

	}

	// get customer first in shop by citizen id
	var customer models.Customer
	if errC := models.GetCustomerFirstByCitizenId(checkCustomerOnline.Co_citizen_id, &customer); errC != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Customer Invalid!",
			"data":    errC.Error(),
		})
		return
	}
	if customer.ID == 0 { //ลูกค้าออนไลน์ต้องมีร้านเท่านั้น
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Customer Invalid!",
			"data":    "",
		})
		return
	}

	var coActk = middlewares.CreateAccessToken(checkCustomerOnline.ID, customer.ID, customer.ShopId, customer.ShopMotherId, checkCustomerOnline.Co_line_id, checkCustomerOnline.Co_citizen_id)

	//add log
	if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
		Line_id:        checkCustomerOnline.Co_line_id,
		Log_ip_address: c.ClientIP(),
		Log_browser:    c.GetHeader("User-Agent"),
		Log_text:       "Login by Line ID",
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

	models.AddBetterstackLoginLog(structs.ObjLogCustomerLogin{
		Line_id:        checkCustomerOnline.Co_line_id,
		Log_ip_address: c.ClientIP(),
		Log_browser:    c.GetHeader("User-Agent"),
		Log_text:       "Login by Line ID",
		Log_create:     time.Now().Format("2006-01-02 15:04:05"),
	})

	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    responseData,
	})

}

func Logout(c *gin.Context) {
	lineId := c.Params.ByName("lineId")
	var checkCustomerOnline structs.CustomerOnline
	if errCO := models.GetCustomerOnlineByLineId2(lineId, &checkCustomerOnline); errCO != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Customer Online Invalid!",
			"data":    errCO.Error(),
		})
		return
	}

	if checkCustomerOnline.ID > 0 {
		if errCO2 := models.DeleteCustomerOnline(checkCustomerOnline.ID); errCO2 != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Something went wrong.",
				"data":    errCO2.Error(),
			})
			return
		}
	} else {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "ID not found",
			"data":    "",
		})
		return
	}
	//add log
	if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
		Line_id:        lineId,
		Log_ip_address: c.ClientIP(),
		Log_browser:    c.GetHeader("User-Agent"),
		Log_text:       "Logout by Line ID",
		Log_create:     time.Now().Format("2006-01-02 15:04:05"),
	}); errL != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Create Log Customer Error.",
			"data":    "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  true,
		"message": "success",
		"data":    "",
	})
}

func Otp(c *gin.Context) {

	var payload structs.PayloadOtp

	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	// get customer online by line id
	var customerOnline structs.CustomerOnline
	if errCO := models.GetCustomerOnlineByOtp(payload.Line_id, payload.Tel, payload.Otp, &customerOnline); errCO != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Customer Online OTP Invalid!",
			"data":    errCO.Error(),
		})
		return
	}
	if customerOnline.ID == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Submit OTP Not Match.",
			"data":    "",
		})
		return
	} else {
		//update customer online
		if errL := models.UpdateCustomerOnline(customerOnline.ID, &structs.CustomerOnlineUpdate{
			Co_otp:        "",
			Co_otp_expire: "",
			Co_update:     time.Now().Format("2006-01-02 15:04:05"),
		}); errL != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Update Customer Online Error.",
				"data":    "",
			})
			return
		}

		// get customer first in shop by citizen id
		var customer models.Customer
		if errC := models.GetCustomerFirstByCitizenId(customerOnline.Co_citizen_id, &customer); errC != nil {
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
				"message": "Customer Invalid!",
				"data":    "",
			})
			return
		}

		// create access token
		var coActk = middlewares.CreateAccessToken(customerOnline.ID, customer.ID, customer.ShopId, customer.ShopMotherId, customerOnline.Co_line_id, customerOnline.Co_citizen_id)

		//add log
		if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
			Line_id:        customerOnline.Co_line_id,
			Log_ip_address: c.ClientIP(),
			Log_browser:    c.GetHeader("User-Agent"),
			Log_text:       "Login OTP by Line ID",
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
}

func Register(c *gin.Context) {

	var payload structs.PayloadRegis

	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	var customerOnlineStructs []structs.CustomerOnline
	// var customerOnlineStructsCitizen []structs.CustomerOnline
	var customerOnlineStructsCheckTel []structs.CustomerOnline

	if errMD := models.CheckMail(payload.Co_email, &customerOnlineStructs); errMD != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    "",
		})
		return
	}

	// if errMDCheckCitizen := models.CheckCitizen(payload.Co_citizen_id, &customerOnlineStructsCitizen); errMDCheckCitizen != nil {
	// 	c.AbortWithStatusJSON(200, gin.H{
	// 		"status":  false,
	// 		"message": "Something went wrong.",
	// 		"data":    "",
	// 	})
	// 	return
	// }

	if errMDCheckTel := models.CheckTel(payload.Co_tel, &customerOnlineStructsCheckTel); errMDCheckTel != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    "",
		})
		return
	}

	if customerOnlineStructs == nil || len(customerOnlineStructs) < 1 {
		// if len(customerOnlineStructsCitizen) < 1 {
		if len(customerOnlineStructsCheckTel) < 1 {
			var password = strings.ToLower(payload.Co_email) + payload.Co_password
			hash := hashPassword(password)

			timeCreated := time.Now()
			timeExpire := timeCreated.Add(time.Duration(5) * time.Minute)

			//add customer online
			otp := middlewares.RandomNumberString(6)
			if errL := models.AddCustomerOnline(&structs.CustomerOnline{
				// Co_citizen_id: payload.Co_citizen_id,
				Co_password:   hash,
				Co_email:      payload.Co_email,
				Co_tel:        payload.Co_tel,
				Co_prefix:     payload.Co_prefix,
				Co_fname:      payload.Co_fname,
				Co_lname:      payload.Co_lname,
				Co_otp:        otp,
				Co_otp_expire: timeExpire.Format("2006-01-02 15:04:05"),
				Co_is_active:  1,
				Co_is_del:     0,
				Co_create:     time.Now().Format("2006-01-02 15:04:05"),
				Co_update:     time.Now().Format("2006-01-02 15:04:05"),
			}); errL != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Create Customer Online Error.",
					"data":    "",
				})
				return
			}

			// send otp customer online -----------------
			s := libs.Sms{
				Msisdn:  payload.Co_tel,
				Message: "(SMS OTP) APSX CRM-LINE Platform Clinic OTP " + otp,
			}
			res, _ := s.Send()
			if res.StatusCode != 201 {
				c.JSON(http.StatusOK, gin.H{
					"status":  false,
					"message": "SMS Send Error.",
					"data":    "",
				})
				return
			}

			//add log
			if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
				Line_id:        "",
				Log_ip_address: c.ClientIP(),
				Log_browser:    c.GetHeader("User-Agent"),
				Log_text:       "Register & Create OTP",
				Log_create:     time.Now().Format("2006-01-02 15:04:05"),
			}); errL != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Create Log Customer Online OTP Error.",
					"data":    "",
				})
				return
			}
			c.JSON(200, gin.H{
				"status":  true,
				"message": "Add Customer Online Success & Waiting OTP.",
				"data":    "OTP Online (otponline)",
			})
			return
		} else {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "This tel has Customers[0]. Please change the tel.",
				"data":    "",
			})
		}
		// } else {
		// 	c.AbortWithStatusJSON(200, gin.H{
		// 		"status":  false,
		// 		"message": "This citizen has Customers[0]. Please change the citizen.",
		// 		"data":    "",
		// 	})
		// }
	} else {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "This email has Customers[0]. Please change the email.",
			"data":    "",
		})
	}
}

func RegisterBypassOTP(c *gin.Context) {

	var payload structs.PayloadExamedRegis

	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	var customerOnlineStructs []structs.CustomerOnline
	var customerOnlineStructsCitizen []structs.CustomerOnline
	var customerOnlineStructsCheckTel []structs.CustomerOnline

	if errMD := models.CheckMail(payload.Co_email, &customerOnlineStructs); errMD != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    "",
		})
		return
	}

	if errMDCheckCitizen := models.CheckCitizen(payload.Co_citizen_id, &customerOnlineStructsCitizen); errMDCheckCitizen != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    "",
		})
		return
	}

	if errMDCheckTel := models.CheckTel(payload.Co_tel, &customerOnlineStructsCheckTel); errMDCheckTel != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    "",
		})
		return
	}

	if customerOnlineStructs == nil || len(customerOnlineStructs) < 1 {
		if len(customerOnlineStructsCitizen) < 1 {
			if len(customerOnlineStructsCheckTel) < 1 {
				// password = 5 digit of last citizen_id
				var password = middlewares.GetSubstring(payload.Co_citizen_id, 5, false)
				hash := hashPassword(password)

				timeCreated := time.Now()
				timeExpire := timeCreated.Add(time.Duration(5) * time.Minute)

				//add customer online
				if errL := models.AddCustomerOnline(&structs.CustomerOnline{
					Co_citizen_id: payload.Co_citizen_id,
					Co_password:   hash,
					Co_email:      payload.Co_email,
					Co_tel:        payload.Co_tel,
					Co_prefix:     payload.Co_prefix,
					Co_fname:      payload.Co_fname,
					Co_lname:      payload.Co_lname,
					Co_otp:        "",
					Co_otp_expire: timeExpire.Format("2006-01-02 15:04:05"),
					Co_is_active:  1,
					Co_is_del:     0,
					Co_create:     time.Now().Format("2006-01-02 15:04:05"),
					Co_update:     time.Now().Format("2006-01-02 15:04:05"),
				}); errL != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Create Customer Online Error.",
						"data":    "",
					})
					return
				}

				// send otp customer online -----------------
				// s := libs.Sms{
				// 	Msisdn:  payload.Co_tel,
				// 	Message: "(SMS OTP) APSX CRM-LINE Platform Clinic OTP " + strconv.Itoa(randomNumber),
				// }
				// res, _ := s.Send()
				// if res.StatusCode != 201 {
				// 	c.JSON(http.StatusOK, gin.H{
				// 		"status":  false,
				// 		"message": "SMS Send Error.",
				// 		"data":    "",
				// 	})
				// 	return
				// }

				//add log
				if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
					Line_id:        "",
					Log_ip_address: c.ClientIP(),
					Log_browser:    c.GetHeader("User-Agent"),
					Log_text:       "Register ByPass",
					Log_create:     time.Now().Format("2006-01-02 15:04:05"),
				}); errL != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Create Log Customer Online",
						"data":    "",
					})
					return
				}
				c.JSON(200, gin.H{
					"status":  true,
					"message": "Add Customer Online Success",
					"data":    "",
				})
				return
			} else {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "This tel has Customers[0]. Please change the tel.",
					"data":    "",
				})
			}
		} else {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "This citizen has Customers[0]. Please change the citizen.",
				"data":    "",
			})
		}

	} else {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "This email has Customers[0]. Please change the email.",
			"data":    "",
		})
	}
}

func RegisterTelOnly(c *gin.Context) {
	var payload structs.PayloadRegisTelOnly

	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	// var customerOnlineStructsCitizen []structs.CustomerOnline
	var customerOnlineStructsCheckTel []structs.CustomerOnline
	var customerOnlineStructs []structs.CustomerOnline
	if errMDCheckTel := models.CheckTel(payload.Co_tel, &customerOnlineStructsCheckTel); errMDCheckTel != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    "",
		})
		return
	}

	if customerOnlineStructs == nil || len(customerOnlineStructs) < 1 {
		// if len(customerOnlineStructsCitizen) < 1 {
		if len(customerOnlineStructsCheckTel) < 1 {

			timeCreated := time.Now()
			timeExpire := timeCreated.Add(time.Duration(5) * time.Minute)

			randomNumber := rand.Intn(900000) + 100000

			//add customer online
			if errL := models.AddCustomerOnline(&structs.CustomerOnline{
				// Co_citizen_id: payload.Co_citizen_id,
				Co_tel:        payload.Co_tel,
				Co_otp:        strconv.Itoa(randomNumber),
				Co_otp_expire: timeExpire.Format("2006-01-02 15:04:05"),
				Co_is_active:  1,
				Co_is_del:     0,
				Co_create:     time.Now().Format("2006-01-02 15:04:05"),
				Co_update:     time.Now().Format("2006-01-02 15:04:05"),
			}); errL != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Create Customer Online Error.",
					"data":    "",
				})
				return
			}

			// send otp customer online -----------------
			s := libs.Sms{
				Msisdn:  payload.Co_tel,
				Message: "(SMS OTP) APSX CRM-LINE Platform Clinic OTP " + strconv.Itoa(randomNumber),
			}
			res, _ := s.Send()
			if res.StatusCode != 201 {
				c.JSON(http.StatusOK, gin.H{
					"status":  false,
					"message": "SMS Send Error.",
					"data":    "",
				})
				return
			}

			//add log
			if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
				Line_id:        "",
				Log_ip_address: c.ClientIP(),
				Log_browser:    c.GetHeader("User-Agent"),
				Log_text:       "Register & Create OTP",
				Log_create:     time.Now().Format("2006-01-02 15:04:05"),
			}); errL != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Create Log Customer Online OTP Error.",
					"data":    "",
				})
				return
			}
			c.JSON(200, gin.H{
				"status":  true,
				"message": "Add Customer Online Success & Waiting OTP.",
				"data":    "OTP Online (otponline)",
			})
			return
		} else {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "This tel has Customers[0]. Please change the tel.",
				"data":    "",
			})
		}
		// } else {
		// 	c.AbortWithStatusJSON(200, gin.H{
		// 		"status":  false,
		// 		"message": "This citizen has Customers[0]. Please change the citizen.",
		// 		"data":    "",
		// 	})
		// }
	} else {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "This email has Customers[0]. Please change the email.",
			"data":    "",
		})
	}
}

func Otponline(c *gin.Context) {

	var payload structs.PayloadOtpRegis

	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	// get customer online by line id
	var customerOnline structs.CustomerOnline
	if errCO := models.GetCustomerOnlineByOtpRegister(payload.Co_email, payload.Co_tel, payload.Otp, &customerOnline); errCO != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Customer Online OTP Invalid!",
			"data":    errCO.Error(),
		})
		return
	}
	if customerOnline.ID == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Submit OTP Not Match.",
			"data":    "",
		})
		return
	} else {
		//update customer online
		if errL := models.UpdateCustomerOnline(customerOnline.ID, &structs.CustomerOnlineUpdate{
			Co_otp:        "",
			Co_otp_expire: "",
			Co_update:     time.Now().Format("2006-01-02 15:04:05"),
		}); errL != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Update Customer Online Error.",
				"data":    "",
			})
			return
		}

		// get customer first in shop by citizen id
		var coActk string
		if customerOnline.Co_citizen_id == "" {
			coActk = middlewares.CreateAccessToken(customerOnline.ID, 0, 0, 0, "", "")
		} else {
			var customer models.Customer
			if errC := models.GetCustomerFirstByCitizenId(customerOnline.Co_citizen_id, &customer); errC != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Error Get Customer Invalid!",
					"data":    errC.Error(),
				})
				return
			}

			// add log
			if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
				Line_id:        "",
				Log_ip_address: c.ClientIP(),
				Log_browser:    c.GetHeader("User-Agent"),
				Log_text:       "Register & Send Verify Email.",
				Log_create:     time.Now().Format("2006-01-02 15:04:05"),
			}); errL != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Create Log Customer Online OTP Error.",
					"data":    "",
				})
				return
			}
			coActk = middlewares.CreateAccessToken(customerOnline.ID, customer.ID, customer.ShopId, customer.ShopMotherId, customerOnline.Co_line_id, customer.CtmCitizenId)
		}

		// create access token
		// var coActk = middlewares.CreateAccessToken(customerOnline.ID, customer.ID, customer.ShopId, customer.ShopMotherId, customerOnline.Co_line_id, customerOnline.Co_citizen_id)
		// if(customerOnline.Co_citizen_id == ""){
		// 	coActk = middlewares.CreateAccessToken(customerOnline.ID, 0, 0, 0, customerOnline.Co_line_id, "")
		// }else{
		// 	coActk = middlewares.CreateAccessToken(customerOnline.ID, 0, 0, 0, customerOnline.Co_line_id, customerOnline.Co_citizen_id)
		// }
		//add log
		// if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
		// 	Line_id:        customerOnline.Co_line_id,
		// 	Log_ip_address: c.ClientIP(),
		// 	Log_browser:    c.GetHeader("User-Agent"),
		// 	Log_text:       "Login OTP by Online",
		// 	Log_create:     time.Now().Format("2006-01-02 15:04:05"),
		// }); errL != nil {
		// 	c.AbortWithStatusJSON(200, gin.H{
		// 		"status":  false,
		// 		"message": "Create Log Customer Error.",
		// 		"data":    "",
		// 	})
		// 	return
		// }

		responseData := structs.ResponseOauth{
			AccessToken: coActk,
		}

		c.JSON(200, gin.H{
			"status":  true,
			"message": "",
			"data":    responseData,
		})

	}
}

func OtponlineTelOnly(c *gin.Context) {
	var payload structs.PayloadOtpTelOnlyRegis

	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	var customerOnline structs.CustomerOnline
	if errCO := models.GetCustomerOnlineByOtpTelOnlyRegister(payload.Co_tel, payload.Otp, &customerOnline); errCO != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Customer Online OTP Invalid!",
			"data":    errCO.Error(),
		})
		return
	}
	if customerOnline.ID == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Submit OTP Not Match.",
			"data":    "",
		})
		return
	} else {
		//update customer online
		if errL := models.UpdateCustomerOnline(customerOnline.ID, &structs.CustomerOnlineUpdate{
			Co_otp:        "",
			Co_otp_expire: "",
			Co_update:     time.Now().Format("2006-01-02 15:04:05"),
		}); errL != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Update Customer Online Error.",
				"data":    "",
			})
			return
		}

		var coActk = middlewares.CreateAccessToken(customerOnline.ID, 0, 0, 0, customerOnline.Co_line_id, "")
		//add log
		if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
			Line_id:        customerOnline.Co_line_id,
			Log_ip_address: c.ClientIP(),
			Log_browser:    c.GetHeader("User-Agent"),
			Log_text:       "Login OTP by Online",
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

}

func Onlinesync(c *gin.Context) {

	var payload structs.PayloadSync

	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	var customerOnlineStructsCitizen []structs.CustomerOnline
	if errMDCheckCitizen := models.CheckCitizen(payload.Co_citizen_id, &customerOnlineStructsCitizen); errMDCheckCitizen != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    "",
		})
		return
	}
	Id := libs.StrToInt(c.Params.ByName("customerOnlineId"))
	if len(customerOnlineStructsCitizen) < 1 {

		var CustomerOnline structs.CustomerOnline
		if errMDCheckCitizen := models.GetCustomerOnlineById(Id, &CustomerOnline); errMDCheckCitizen != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Something went wrong.",
				"data":    "",
			})
			return
		}
		timeCreated := time.Now()
		timeExpire := timeCreated.Add(time.Duration(5) * time.Minute)
		randomNumber := rand.Intn(900000) + 100000
		//update customer online
		if CustomerOnline.ID != 0 {
			if errL := models.UpdateCustomerOnlineSync(CustomerOnline.ID, &structs.CustomerOnlineUpdateSync{
				Co_citizen_id: payload.Co_citizen_id,
				Co_otp:        strconv.Itoa(randomNumber),
				Co_otp_expire: timeExpire.Format("2006-01-02 15:04:05"),
				Co_update:     time.Now().Format("2006-01-02 15:04:05"),
			}); errL != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Update Customer Online Error.",
					"data":    "",
				})
				return
			}
			// send otp customer online -----------------
			s := libs.Sms{
				Msisdn:  CustomerOnline.Co_tel,
				Message: "(SMS OTP) APSX CRM-LINE Platform Clinic OTP " + strconv.Itoa(randomNumber),
			}
			res, _ := s.Send()
			if res.StatusCode != 201 {
				c.JSON(http.StatusOK, gin.H{
					"status":  false,
					"message": "SMS Send Error.",
					"data":    "",
				})
				return
			}

			//add log
			if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
				Line_id:        "",
				Log_ip_address: c.ClientIP(),
				Log_browser:    c.GetHeader("User-Agent"),
				Log_text:       "Register & Create OTP",
				Log_create:     time.Now().Format("2006-01-02 15:04:05"),
			}); errL != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Create Log Customer Online OTP Error.",
					"data":    "",
				})
				return
			}
			c.JSON(200, gin.H{
				"status":  true,
				"message": "Add Customer Online Success & Waiting OTP.",
				"data":    "OTP Online Sync (Otponlinesync)",
				"co_tel":  CustomerOnline.Co_tel,
			})
			return
		} else {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Get Customer Online Error.",
				"data":    "",
			})
			return
		}
	} else {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "This citizen has Customers[0]. Please change the citizen.",
			"data":    "",
		})
	}

}

func Otponlinesync(c *gin.Context) {

	var payload structs.PayloadOtpSync

	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}
	Id := libs.StrToInt(c.Params.ByName("customerOnlineId"))
	// get customer online by line id
	var customerOnline structs.CustomerOnline
	if errCO := models.GetCustomerOnlineByOtpSync(Id, payload.Otp, &customerOnline); errCO != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Customer Online OTP Invalid!",
			"data":    errCO.Error(),
		})
		return
	}
	if customerOnline.ID == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Submit OTP Not Match.",
			"data":    "",
		})
		return
	} else {
		// get customer first in shop by citizen id
		var customer models.Customer
		if errC := models.GetCustomerFirstByCitizenId(customerOnline.Co_citizen_id, &customer); errC != nil {
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
				"message": "Customer Invalid!",
				"data":    "",
			})
			return
		}

		//update customer online
		if errL := models.UpdateCustomerOnlineTel(customerOnline.ID, &structs.CustomerOnlineUpdateTel{
			Co_tel:        customer.CtmTel,
			Co_otp:        "",
			Co_otp_expire: "",
			Co_update:     time.Now().Format("2006-01-02 15:04:05"),
		}); errL != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Update Customer Online Error.",
				"data":    "",
			})
			return
		}

		// create access token
		var coActk = middlewares.CreateAccessToken(customerOnline.ID, customer.ID, customer.ShopId, customer.ShopMotherId, customerOnline.Co_line_id, customerOnline.Co_citizen_id)
		//add log
		if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
			Line_id:        customerOnline.Co_line_id,
			Log_ip_address: c.ClientIP(),
			Log_browser:    c.GetHeader("User-Agent"),
			Log_text:       "Login OTP by Online",
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
}

func Onlineunsync(c *gin.Context) {

	Id := libs.StrToInt(c.Params.ByName("customerOnlineId"))
	// get customer online by line id
	var customerOnline structs.CustomerOnline
	if errCO := models.GetCustomerOnlineById(Id, &customerOnline); errCO != nil {
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
			"message": "Submit OTP Not Match.",
			"data":    "",
		})
		return
	} else {
		//update customer online
		if errL := models.UpdateCustomerOnlineSync(customerOnline.ID, &structs.CustomerOnlineUpdateSync{
			Co_citizen_id: "",
			Co_otp:        "",
			Co_otp_expire: "",
			Co_update:     time.Now().Format("2006-01-02 15:04:05"),
		}); errL != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Update Customer Online Error.",
				"data":    "",
			})
			return
		}

		// create access token
		var coActk = middlewares.CreateAccessToken(customerOnline.ID, 0, 0, 0, customerOnline.Co_line_id, "")
		//add log
		if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
			Line_id:        customerOnline.Co_line_id,
			Log_ip_address: c.ClientIP(),
			Log_browser:    c.GetHeader("User-Agent"),
			Log_text:       "Login OTP by Online",
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
}

// EXA AUTHENTICATION----------------------------------------------------------------

// เชื่อมต่อบัญชี EXA กับ LINE (EXA MED)
func SyncAccountWithLine(c *gin.Context) {
	var payload structs.PayloadSyncAccountWithEmail

	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}
	// Co_id := c.Params.ByName("co_id")

	// get customer online by line id
	var customerOnline structs.CustomerOnline
	if errCO := models.GetCustomerOnlineByLineId(payload.Line_id, &customerOnline); errCO != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Customer Online Invalid!",
			"data":    errCO.Error(),
		})
		return
	}
	if customerOnline.ID != 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "This line id has registered with another account.",
			"data":    "",
		})
		return
	}

	// // get customer by email
	var customer structs.CustomerOnline
	if errC := models.GetCustomerOnlineByEmail(payload.Co_email, &customer); errC != nil {
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
			"message": "Customer not found by email.",
			"data":    "",
		})
		return
	}
	// fmt.Println("Sync Account With Email", customer)
	var updateCustomer structs.SyncAccountWithEmail
	updateCustomer.Co_Line_id = payload.Line_id
	updateCustomer.Co_Line_email = payload.Line_email
	updateCustomer.Co_Line_name = payload.Line_name
	updateCustomer.Co_update = time.Now().Format("2006-01-02 15:04:05")

	if errU := models.SyncAccountWithEmail(customer.ID, &updateCustomer); errU != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Update Customer Online Error.",
			"data":    "",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "sync successfully",
	})

}

// ลงทะเบียนบัญชี EXA
func RegisterExa(c *gin.Context) {

	var payload structs.PayloadRegis

	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	var customerOnlineStructs []structs.CustomerOnline
	// var customerOnlineStructsCitizen []structs.CustomerOnline
	var customerOnlineStructsCheckTel []structs.CustomerOnline

	if errMD := models.CheckMail(payload.Co_email, &customerOnlineStructs); errMD != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    "",
		})
		return
	}

	// if errMDCheckCitizen := models.CheckCitizen(payload.Co_citizen_id, &customerOnlineStructsCitizen); errMDCheckCitizen != nil {
	// 	c.AbortWithStatusJSON(200, gin.H{
	// 		"status":  false,
	// 		"message": "Something went wrong.",
	// 		"data":    "",
	// 	})
	// 	return
	// }

	if errMDCheckTel := models.CheckTel(payload.Co_tel, &customerOnlineStructsCheckTel); errMDCheckTel != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    "",
		})
		return
	}

	if customerOnlineStructs == nil || len(customerOnlineStructs) < 1 {
		// if len(customerOnlineStructsCitizen) < 1 {
		if len(customerOnlineStructsCheckTel) < 1 {
			var password = strings.ToLower(payload.Co_email) + payload.Co_password
			hash := hashPassword(password)

			timeCreated := time.Now()
			timeExpire := timeCreated.Add(time.Duration(5) * time.Minute)
			message := "แจ้งรหัสยืนยันตัวตน"

			otpCode := libs.RandStringBytesMaskImpr("0123456789", 6)
			otpKey := libs.RandStringBytesMaskImpr("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ", 6)
			otpMessage := libs.RenderEmailRequestOtp(message, payload.Co_fname, otpCode, otpKey)

			// randomNumber := rand.Intn(900000) + 100000

			//add customer online
			if errL := models.AddCustomerOnline(&structs.CustomerOnline{
				// Co_citizen_id: payload.Co_citizen_id,
				Co_password:   hash,
				Co_email:      payload.Co_email,
				Co_tel:        payload.Co_tel,
				Co_prefix:     payload.Co_prefix,
				Co_fname:      payload.Co_fname,
				Co_lname:      payload.Co_lname,
				Co_otp:        otpCode,
				Co_otp_key:    otpKey,
				Co_otp_expire: timeExpire.Format("2006-01-02 15:04:05"),
				Co_is_active:  1,
				Co_is_del:     0,
				Co_create:     time.Now().Format("2006-01-02 15:04:05"),
				Co_update:     time.Now().Format("2006-01-02 15:04:05"),
			}); errL != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Create Customer Online Error.",
					"data":    "",
				})
				return
			}

			// send email customer online -----------------
			if payload.Co_email == "" {
				c.JSON(200, gin.H{
					"status":  false,
					"message": "Email not found.",
					"data":    "",
				})
				return
			}

			errSendmail := libs.SendMail(payload.Co_email, "แจ้งรหัสยืนยันตัวตน", otpMessage)

			if errSendmail != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Error Send Mail.",
					"data":    "",
				})
				return
			}

			// send otp customer online -----------------
			// s := libs.Sms{
			// 	Msisdn:  payload.Co_tel,
			// 	Message: "(SMS OTP) APSX CRM-LINE Platform Clinic OTP " + strconv.Itoa(randomNumber),
			// }
			// res, _ := s.Send()
			// if res.StatusCode != 201 {
			// 	c.JSON(http.StatusOK, gin.H{
			// 		"status":  false,
			// 		"message": "SMS Send Error.",
			// 		"data":    "",
			// 	})
			// 	return
			// }

			//add log
			if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
				Line_id:        "",
				Log_ip_address: c.ClientIP(),
				Log_browser:    c.GetHeader("User-Agent"),
				Log_text:       "Register & Send Verify Email.",
				Log_create:     time.Now().Format("2006-01-02 15:04:05"),
			}); errL != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Create Log Customer Online OTP Error.",
					"data":    "",
				})
				return
			}
			c.JSON(200, gin.H{
				"status":  true,
				"message": "Add Customer Online Success & Waiting Verify Email.",
				"data":    gin.H{"otp_key": otpKey},
			})
			return
		} else {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "This tel has Customers[0]. Please change the tel.",
				"data":    "",
			})
		}
		// } else {
		// 	c.AbortWithStatusJSON(200, gin.H{
		// 		"status":  false,
		// 		"message": "This citizen has Customers[0]. Please change the citizen.",
		// 		"data":    "",
		// 	})
		// }
	} else {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "This email has Customers[0]. Please change the email.",
			"data":    "",
		})
	}
}

// เข้าสู่ระบบบัญชี EXA  ด้วย LINE
func LoginExa(c *gin.Context) {

	var payload structs.PayloadLoginTest

	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}
	fmt.Printf("Received payload: %+v\n", payload)

	// check line id
	// get customer online by line id
	var checkCustomerOnline structs.CustomerOnline
	if errCO := models.GetCustomerOnlineByLineId(payload.Line_id, &checkCustomerOnline); errCO != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Customer Online Invalid!",
			"data":    errCO.Error(),
		})
		return
	}
	if checkCustomerOnline.ID == 0 {

		fmt.Printf("checkCustomerOnline: %+v\n", checkCustomerOnline)

		// get customer online by line id
		var customerOnline structs.CustomerOnline
		if errCO := models.CheckCustomerOnlineLogin(payload.Line_id, payload.Citizen_id, payload.Tel, &customerOnline); errCO != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Error Get Customer Online Invalid!",
				"data":    errCO.Error(),
			})
			return
		}
		if customerOnline.ID == 0 {

			timeCreated := time.Now()
			timeExpire := timeCreated.Add(time.Duration(5) * time.Minute)
			// hash := hashPassword(payload.Co_password)
			randomNumber := rand.Intn(900000) + 100000

			//add customer online
			if errL := models.AddCustomerOnline(&structs.CustomerOnline{
				Co_line_name:  payload.Line_name,
				Co_line_email: payload.Line_email,
				Co_line_id:    payload.Line_id,
				Co_citizen_id: payload.Citizen_id,
				Co_tel:        payload.Tel,
				Co_otp:        strconv.Itoa(randomNumber),
				Co_otp_expire: timeExpire.Format("2006-01-02 15:04:05"),
				Co_is_active:  1,
				Co_is_del:     0,
				Co_create:     time.Now().Format("2006-01-02 15:04:05"),
				Co_update:     time.Now().Format("2006-01-02 15:04:05"),
			}); errL != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Create Customer Online Error.",
					"data":    "",
				})
				return
			}

			// send otp customer online -----------------
			s := libs.Sms{
				Msisdn:  payload.Tel,
				Message: "(SMS OTP) APSX CRM-LINE Platform Clinic OTP " + strconv.Itoa(randomNumber),
			}
			res, _ := s.Send()
			if res.StatusCode != 201 {
				c.JSON(http.StatusOK, gin.H{
					"status":  false,
					"message": "SMS Send Error.",
					"data":    "",
				})
				return
			}

			//add log
			if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
				Line_id:        customerOnline.Co_line_id,
				Log_ip_address: c.ClientIP(),
				Log_browser:    c.GetHeader("User-Agent"),
				Log_text:       "Register Line & Create OTP",
				Log_create:     time.Now().Format("2006-01-02 15:04:05"),
			}); errL != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Create Log Customer Online OTP Error.",
					"data":    "",
				})
				return
			}

			c.JSON(200, gin.H{
				"status":  true,
				"message": "Add Customer Online Success & Waiting OTP.",
				"data":    "",
			})
			return
		} else {

			timeCreated := time.Now()
			timeExpire := timeCreated.Add(time.Duration(5) * time.Minute)

			randomNumber := rand.Intn(900000) + 100000

			//update customer online
			if errL := models.UpdateCustomerOnline(customerOnline.ID, &structs.CustomerOnlineUpdate{
				// Co_citizen_id: payload.Citizen_id,
				// Co_tel:        payload.Tel,
				Co_otp:        strconv.Itoa(randomNumber),
				Co_otp_expire: timeExpire.Format("2006-01-02 15:04:05"),
				Co_update:     time.Now().Format("2006-01-02 15:04:05"),
			}); errL != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Update Customer Online Error.",
					"data":    "",
				})
				return
			}

			// send otp customer online -----------------
			s := libs.Sms{
				Msisdn:  payload.Tel,
				Message: "(SMS OTP) APSX CRM-LINE Platform Clinic OTP " + strconv.Itoa(randomNumber),
			}
			res, _ := s.Send()
			if res.StatusCode != 201 {
				c.JSON(http.StatusOK, gin.H{
					"status":  false,
					"message": "SMS Send Error.",
					"data":    "",
				})
				return
			}

			//add log
			if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
				Line_id:        customerOnline.Co_line_id,
				Log_ip_address: c.ClientIP(),
				Log_browser:    c.GetHeader("User-Agent"),
				Log_text:       "Register Line & Update OTP",
				Log_create:     time.Now().Format("2006-01-02 15:04:05"),
			}); errL != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Create Log Customer Online OTP Error.",
					"data":    "",
				})
				return
			}

			c.JSON(200, gin.H{
				"status":  true,
				"message": "Update Customer Online Waiting OTP.",
				"data":    "",
			})
			return
		}

	}
	if checkCustomerOnline.Co_citizen_id == payload.Citizen_id {
		timeCreated := time.Now()
		timeExpire := timeCreated.Add(time.Duration(5) * time.Minute)

		randomNumber := rand.Intn(900000) + 100000
		if errL := models.UpdateCustomerOnline(checkCustomerOnline.ID, &structs.CustomerOnlineUpdate{
			// Co_citizen_id: payload.Citizen_id,
			// Co_tel:        payload.Tel,
			Co_otp:        strconv.Itoa(randomNumber),
			Co_otp_expire: timeExpire.Format("2006-01-02 15:04:05"),
			Co_update:     time.Now().Format("2006-01-02 15:04:05"),
		}); errL != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Update Customer Online Error.",
				"data":    "",
			})
			return
		}
		s := libs.Sms{
			Msisdn:  payload.Tel,
			Message: "(SMS OTP) APSX CRM-LINE Platform Clinic OTP " + strconv.Itoa(randomNumber),
		}
		res, _ := s.Send()
		if res.StatusCode != 201 {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "SMS Send Error.",
				"data":    "",
			})
			return
		}

		//add log
		if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
			Line_id:        checkCustomerOnline.Co_line_id,
			Log_ip_address: c.ClientIP(),
			Log_browser:    c.GetHeader("User-Agent"),
			Log_text:       "Register Line & Update OTP",
			Log_create:     time.Now().Format("2006-01-02 15:04:05"),
		}); errL != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Create Log Customer Online OTP Error.",
				"data":    "",
			})
			return
		}

		c.JSON(200, gin.H{
			"status":  true,
			"message": "Update Customer Online Waiting OTP.",
			"data":    "",
		})
		return
	}
	if checkCustomerOnline.Co_line_id != payload.Citizen_id {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Line ID Invalid!",
			"data":    "",
		})
		return
	}

	// get customer first in shop by citizen id
	var customer models.Customer
	if errC := models.GetCustomerFirstByCitizenId(checkCustomerOnline.Co_citizen_id, &customer); errC != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Customer Invalid!",
			"data":    errC.Error(),
		})
		return
	}
	var coActk string

	if customer.ID == 0 {
		coActk = middlewares.CreateAccessToken(
			checkCustomerOnline.ID, 0, 0, 0,
			checkCustomerOnline.Co_line_id, checkCustomerOnline.Co_citizen_id,
		)
	} else {
		coActk = middlewares.CreateAccessToken(
			checkCustomerOnline.ID,
			customer.ID,
			customer.ShopId,
			customer.ShopMotherId,
			checkCustomerOnline.Co_line_id,
			checkCustomerOnline.Co_citizen_id,
		)
	}

	//add log
	if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
		Line_id:        checkCustomerOnline.Co_line_id,
		Log_ip_address: c.ClientIP(),
		Log_browser:    c.GetHeader("User-Agent"),
		Log_text:       "Login by Line ID",
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

	models.AddBetterstackLoginLog(structs.ObjLogCustomerLogin{
		Line_id:        checkCustomerOnline.Co_line_id,
		Log_ip_address: c.ClientIP(),
		Log_browser:    c.GetHeader("User-Agent"),
		Log_text:       "Login by Line ID",
		Log_create:     time.Now().Format("2006-01-02 15:04:05"),
	})

	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    responseData,
	})

}

// เข้าสู่ระบบบัญชี EXA  ด้วย Email & Password
func LoginExaNoOtp(c *gin.Context) {
	var payload structs.PayloadLoginOnline
	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	Co_email := strings.ToLower(payload.Co_email)
	var Co_password = Co_email + payload.Co_password
	Password := hashPassword(Co_password)

	var checkCustomerOnline structs.CustomerOnline
	if errCO := models.GetCustomerOnlineEmailPassword(payload.Co_email, Password, &checkCustomerOnline); errCO != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Customer Online Invalid!",
			"data":    errCO.Error(),
		})
		return
	}
	if checkCustomerOnline.ID == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Customer Online Invalid!",
			"data":    "",
		})
		return
	}

	var coActk string

	if checkCustomerOnline.Co_citizen_id == "" {
		coActk = middlewares.CreateAccessToken(checkCustomerOnline.ID, 0, 0, 0, "", "")
	} else {
		var customer models.Customer
		///กลับมาแก้ให้รับ shop_id เป็น payload ด้วย ในอนาคต
		if errC := models.GetCustomerLockShop(checkCustomerOnline.Co_citizen_id, 950, &customer); errC != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Error Get Customer Invalid!",
				"data":    errC.Error(),
			})
			return
		}

		// add log
		if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
			Line_id:        "",
			Log_ip_address: c.ClientIP(),
			Log_browser:    c.GetHeader("User-Agent"),
			Log_text:       "Register & Send Verify Email.",
			Log_create:     time.Now().Format("2006-01-02 15:04:05"),
		}); errL != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Create Log Customer Online OTP Error.",
				"data":    "",
			})
			return
		}
		coActk = middlewares.CreateAccessToken(checkCustomerOnline.ID, customer.ID, customer.ShopId, customer.ShopMotherId, checkCustomerOnline.Co_line_id, customer.CtmCitizenId)
	}

	c.AbortWithStatusJSON(200, gin.H{
		"status":  true,
		"message": "Email Verify Success.",
		"data":    gin.H{"access_token": coActk},
	})
}
func LoginonlineExa(c *gin.Context) {
	var payload structs.PayloadLoginOnline
	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	Co_email := strings.ToLower(payload.Co_email)
	var Co_password = Co_email + payload.Co_password
	Password := hashPassword(Co_password)

	var checkCustomerOnline structs.CustomerOnline
	if errCO := models.GetCustomerOnlineEmailPassword(payload.Co_email, Password, &checkCustomerOnline); errCO != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Customer Online Invalid!",
			"data":    errCO.Error(),
		})
		return
	}
	if checkCustomerOnline.ID == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Customer Online Invalid!",
			"data":    "",
		})
		return
	}

	message := "เข้าสู่ระบบ"

	otpCode := libs.RandStringBytesMaskImpr("0123456789", 6)
	otpKey := libs.RandStringBytesMaskImpr("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ", 6)
	otpMessage := libs.RenderEmailRequestOtp(message, checkCustomerOnline.Co_fname, otpCode, otpKey)

	timeCreated := time.Now()
	timeExpire := timeCreated.Add(time.Duration(5) * time.Minute)

	//update customer online
	if errL := models.UpdateOtpCustomerOnline(checkCustomerOnline.ID, &structs.CustomerOnlineOtpUpdate{
		// Co_citizen_id: checkCustomerOnline.Co_citizen_id,
		Co_otp_key:    otpKey,
		Co_otp:        otpCode,
		Co_otp_expire: timeExpire.Format("2006-01-02 15:04:05"),
		Co_update:     time.Now().Format("2006-01-02 15:04:05"),
	}); errL != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Update Customer Online Error.",
			"data":    "",
		})
		return
	}

	// send email otp

	errSendmail := libs.SendMail(payload.Co_email, message, otpMessage)
	if errSendmail != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Send Mail.",
			"data":    "",
		})
		return
	}

	//add log
	if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
		Line_id:        "",
		Log_ip_address: c.ClientIP(),
		Log_browser:    c.GetHeader("User-Agent"),
		Log_text:       "Register Line & Update OTP",
		Log_create:     time.Now().Format("2006-01-02 15:04:05"),
	}); errL != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Create Log Customer Online OTP Error.",
			"data":    "",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Login Customer Online Waiting OTP Email.",
		"data":    "OTP Online (otponline)",
		"otp_key": otpKey,
	})
	return
}

// ยืนยัน OTP ของผู้ใช้งานผ่าน Email OTP (EXA MED)
func VerifyEmail(c *gin.Context) {
	var payload structs.PayloadEmailVerify

	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	var customerOnline structs.CustomerOnline

	if errCheckOtp := models.GetCustomerOnlineByEmailOtpKeyRegister(payload.Co_email, payload.Co_otp_key, payload.Co_otp, &customerOnline); errCheckOtp != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    "",
		})
		return
	}
	if customerOnline.ID == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid OTP or Email.",
			"data":    "",
		})
		return
	}
	if errL := models.UpdateOtpCustomerOnline(customerOnline.ID, &structs.CustomerOnlineOtpUpdate{
		// Co_citizen_id: customerOnline.Co_citizen_id,
		Co_otp:        "",
		Co_otp_key:    "",
		Co_otp_expire: "",
		Co_update:     time.Now().Format("2006-01-02 15:04:05"),
	}); errL != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Update Customer Online Error.",
			"data":    "",
		})
		return
	}

	var coActk string

	if customerOnline.Co_citizen_id == "" {
		coActk = middlewares.CreateAccessToken(customerOnline.ID, 0, 0, 0, "", "")
	} else {
		var customer models.Customer
		if errC := models.GetCustomerFirstByCitizenId(customerOnline.Co_citizen_id, &customer); errC != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Error Get Customer Invalid!",
				"data":    errC.Error(),
			})
			return
		}

		// add log
		if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
			Line_id:        "",
			Log_ip_address: c.ClientIP(),
			Log_browser:    c.GetHeader("User-Agent"),
			Log_text:       "Register & Send Verify Email.",
			Log_create:     time.Now().Format("2006-01-02 15:04:05"),
		}); errL != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Create Log Customer Online OTP Error.",
				"data":    "",
			})
			return
		}
		coActk = middlewares.CreateAccessToken(customerOnline.ID, customer.ID, customer.ShopId, customer.ShopMotherId, customerOnline.Co_line_id, customer.CtmCitizenId)
	}

	c.AbortWithStatusJSON(200, gin.H{
		"status":  true,
		"message": "Email Verify Success.",
		"data":    gin.H{"access_token": coActk},
	})

}
func ResendOtpEmail(c *gin.Context) {
	var payload structs.PayloadResendEmailOtp

	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	var customerOnline structs.CustomerOnline

	if errCheckOtp := models.GetCustomerOnlineByEmail(payload.Co_email, &customerOnline); errCheckOtp != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    "",
		})
		return
	}
	if customerOnline.ID == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Email not found.",
			"data":    "",
		})
		return
	}
	var otpCode, otpKey string
	message := "แจ้งรหัสยืนยันตัวตน"

	otpCode = libs.RandStringBytesMaskImpr("0123456789", 6)
	otpKey = libs.RandStringBytesMaskImpr("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ", 6)
	otpMessage := libs.RenderEmailRequestOtp(message, customerOnline.Co_fname, otpCode, otpKey)

	errSendmail := libs.SendMail(payload.Co_email, "แจ้งรหัสยืนยันตัวตน", otpMessage)
	if errSendmail != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Send Mail.",
			"data":    "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  true,
		"message": "Add Customer Online Success & Waiting Verify Email.",
		"data":    gin.H{"otp_key": otpKey},
	})
}

// sync เลขบัตร ปชช. และ ส่งOTP ผ่าน Email (EXA MED)
func OnlinesyncExa(c *gin.Context) {

	var payload structs.PayloadSync

	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	var customerOnlineStructsCitizen []structs.CustomerOnline
	if errMDCheckCitizen := models.CheckCitizen(payload.Co_citizen_id, &customerOnlineStructsCitizen); errMDCheckCitizen != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    "",
		})
		return
	}
	Id := libs.StrToInt(c.Params.ByName("customerOnlineId"))
	if len(customerOnlineStructsCitizen) < 1 {

		var CustomerOnline structs.CustomerOnline
		if errMDCheckCitizen := models.GetCustomerOnlineById(Id, &CustomerOnline); errMDCheckCitizen != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Something went wrong.",
				"data":    "",
			})
			return
		}
		timeCreated := time.Now()
		timeExpire := timeCreated.Add(time.Duration(5) * time.Minute)
		randomNumber := rand.Intn(900000) + 100000
		//update customer online
		if CustomerOnline.ID != 0 {
			if errL := models.UpdateCustomerOnlineSync(CustomerOnline.ID, &structs.CustomerOnlineUpdateSync{
				// Co_citizen_id: payload.Co_citizen_id,
				Co_otp:        strconv.Itoa(randomNumber),
				Co_otp_expire: timeExpire.Format("2006-01-02 15:04:05"),
				Co_update:     time.Now().Format("2006-01-02 15:04:05"),
			}); errL != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Update Customer Online Error.",
					"data":    "",
				})
				return
			}
			// send otp customer online -----------------
			s := libs.Sms{
				Msisdn:  CustomerOnline.Co_tel,
				Message: "(SMS OTP) APSX CRM-LINE Platform Clinic OTP " + strconv.Itoa(randomNumber),
			}
			res, _ := s.Send()
			if res.StatusCode != 201 {
				c.JSON(http.StatusOK, gin.H{
					"status":  false,
					"message": "SMS Send Error.",
					"data":    "",
				})
				return
			}

			//add log
			if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
				Line_id:        "",
				Log_ip_address: c.ClientIP(),
				Log_browser:    c.GetHeader("User-Agent"),
				Log_text:       "Register & Create OTP",
				Log_create:     time.Now().Format("2006-01-02 15:04:05"),
			}); errL != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Create Log Customer Online OTP Error.",
					"data":    "",
				})
				return
			}
			c.JSON(200, gin.H{
				"status":  true,
				"message": "Add Customer Online Success & Waiting OTP.",
				"data":    "OTP Online Sync (Otponlinesync)",
				"co_tel":  CustomerOnline.Co_tel,
			})
			return
		} else {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Get Customer Online Error.",
				"data":    "",
			})
			return
		}
	} else {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "This citizen has Customers[0]. Please change the citizen.",
			"data":    "",
		})
	}
}

// ยืนยัน OTP การ sync เลขบัตร ปชช. ของผู้ใช้งานผ่าน Email OTP (EXA MED)
func OtponlinesyncExa(c *gin.Context) {

	var payload structs.PayloadOtpSync

	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}
	Id := libs.StrToInt(c.Params.ByName("customerOnlineId"))
	// get customer online by line id
	var customerOnline structs.CustomerOnline
	if errCO := models.GetCustomerOnlineByOtpSync(Id, payload.Otp, &customerOnline); errCO != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Customer Online OTP Invalid!",
			"data":    errCO.Error(),
		})
		return
	}
	if customerOnline.ID == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Submit OTP Not Match.",
			"data":    "",
		})
		return
	} else {
		if errL := models.UpdateCustomerOnlineSync(customerOnline.ID, &structs.CustomerOnlineUpdateSync{
			Co_citizen_id: payload.Cid,
			// Co_otp:        strconv.Itoa(randomNumber),
			// Co_otp_expire: timeExpire.Format("2006-01-02 15:04:05"),
			Co_update: time.Now().Format("2006-01-02 15:04:05"),
		}); errL != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Update Customer Online Error.",
				"data":    "",
			})
			return
		}
		// get customer first in shop by citizen id
		var customer models.Customer
		if errC := models.GetCustomerFirstByCitizenId(payload.Cid, &customer); errC != nil {
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
				"message": "Customer Invalid!",
				"data":    "",
			})
			return
		}

		//update customer online
		// if errL := models.UpdateCustomerOnlineTel(customerOnline.ID, &structs.CustomerOnlineUpdateTel{
		// 	Co_tel:        customer.CtmTel,
		// 	Co_otp:        "",
		// 	Co_otp_expire: "",
		// 	Co_update:     time.Now().Format("2006-01-02 15:04:05"),
		// }); errL != nil {
		// 	c.AbortWithStatusJSON(200, gin.H{
		// 		"status":  false,
		// 		"message": "Update Customer Online Error.",
		// 		"data":    "",
		// 	})
		// 	return
		// }

		// create access token
		var coActk = middlewares.CreateAccessToken(customerOnline.ID, customer.ID, customer.ShopId, customer.ShopMotherId, customerOnline.Co_line_id, customerOnline.Co_citizen_id)
		//add log
		if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
			Line_id:        customerOnline.Co_line_id,
			Log_ip_address: c.ClientIP(),
			Log_browser:    c.GetHeader("User-Agent"),
			Log_text:       "Login OTP by Online",
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
}
func OtponlinesyncExaLockShop(c *gin.Context) {

	var payload structs.PayloadOtpSyncLockShop

	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}
	Id := libs.StrToInt(c.Params.ByName("customerOnlineId"))
	// get customer online by line id
	var customerOnline structs.CustomerOnline
	if errCO := models.GetCustomerOnlineByOtpSync(Id, payload.Otp, &customerOnline); errCO != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Error Get Customer Online OTP Invalid!",
			"data":    errCO.Error(),
		})
		return
	}
	if customerOnline.ID == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Submit OTP Not Match.",
			"data":    "",
		})
		return
	} else {
		if errL := models.UpdateCustomerOnlineSync(customerOnline.ID, &structs.CustomerOnlineUpdateSync{
			// Co_Tel: payload.co
			Co_citizen_id: payload.Cid,
			Co_otp:        "",
			Co_otp_expire: "",
			Co_update:     time.Now().Format("2006-01-02 15:04:05"),
		}); errL != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Update Customer Online Error.",
				"data":    "",
			})
			return
		}
		// get customer first in shop by citizen id
		var customer models.Customer
		if errC := models.GetCustomerLockShop(payload.Cid, payload.Shop_Id, &customer); errC != nil {
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
				"message": "Customer Invalid!",
				"data":    "",
			})
			return
		}

		//update customer online
		// if errL := models.UpdateCustomerOnlineTel(customerOnline.ID, &structs.CustomerOnlineUpdateTel{
		// 	// Co_tel:        customer.CtmTel,
		// 	Co_otp:        "",
		// 	Co_otp_expire: "",
		// 	Co_update:     time.Now().Format("2006-01-02 15:04:05"),
		// }); errL != nil {
		// 	c.AbortWithStatusJSON(200, gin.H{
		// 		"status":  false,
		// 		"message": "Update Customer Online Error.",
		// 		"data":    "",
		// 	})
		// 	return
		// }

		// create access token
		var coActk = middlewares.CreateAccessToken(customerOnline.ID, customer.ID, customer.ShopId, customer.ShopMotherId, customerOnline.Co_line_id, customerOnline.Co_citizen_id)
		//add log
		if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
			Line_id:        customerOnline.Co_line_id,
			Log_ip_address: c.ClientIP(),
			Log_browser:    c.GetHeader("User-Agent"),
			Log_text:       "Login OTP by Online",
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
}

//health survey auth

func HealthSurveyAuth(c *gin.Context) {
	// get customer online by line id
	var payload structs.PayloadLoginHealthSurvey
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	var customerOnline structs.CustomerOnline

	if errorCO := models.GetCustomerOnlineByCitizenId(payload.Co_citizen_id, &customerOnline); errorCO != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Customer Online Not Found!",
			"data":    errorCO.Error(),
		})
		return
	}
	if customerOnline.ID == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Customer Online Not Found!",
			"data":    "",
		})
		return
	}
	timeCreated := time.Now()
	timeExpire := timeCreated.Add(time.Duration(5) * time.Minute)
	randomNumber := rand.Intn(900000) + 100000

	if errL := models.UpdateCustomerOnline(customerOnline.ID, &structs.CustomerOnlineUpdate{
		Co_otp:        strconv.Itoa(randomNumber),
		Co_otp_expire: timeExpire.Format("2006-01-02 15:04:05"),
		Co_update:     time.Now().Format("2006-01-02 15:04:05"),
	}); errL != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Update Customer Online Error.",
			"data":    "",
		})
		return
	}

	s := libs.Sms{
		Msisdn:  customerOnline.Co_tel,
		Message: "(SMS OTP) APSX CRM-LINE Platform Clinic OTP " + strconv.Itoa(randomNumber),
	}
	res, _ := s.Send()
	if res.StatusCode != 201 {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "SMS Send Error.",
			"data":    "",
		})
		return
	}
	if res.StatusCode != 201 {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "SMS Send Error.",
			"data":    "",
		})
		return
	}

	//add log
	if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
		Line_id:        customerOnline.Co_line_id,
		Log_ip_address: c.ClientIP(),
		Log_browser:    c.GetHeader("User-Agent"),
		Log_text:       "Auth Health Survey & Update OTP",
		Log_create:     time.Now().Format("2006-01-02 15:04:05"),
	}); errL != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Create Log Customer Online OTP Error.",
			"data":    "",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Waiting OTP.",
		"data":    gin.H{"co_tel": customerOnline.Co_tel},
	})

}
func OtpHealthSurvey(c *gin.Context) {

	var payload structs.PayloadVerifyHealthSurvey

	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	// get customer online by line id
	var customerOnline structs.CustomerOnline
	if errorCO := models.GetCustomerOnlineByCitizenId(payload.Co_citizen_id, &customerOnline); errorCO != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Customer Online Not Found!",
			"data":    errorCO.Error(),
		})
		return
	}
	if customerOnline.ID == 0 {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Submit OTP Not Match.",
			"data":    "",
		})
		return
	}
	if customerOnline.Co_otp != payload.Co_otp {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Submit OTP Not Match.",
			"data":    "",
		})
		return
	} else {
		//update customer online
		if errL := models.UpdateCustomerOnline(customerOnline.ID, &structs.CustomerOnlineUpdate{
			Co_otp:        "",
			Co_otp_expire: "",
			Co_update:     time.Now().Format("2006-01-02 15:04:05"),
		}); errL != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Update Customer Online Error.",
				"data":    "",
			})
			return
		}

		// get customer first in shop by citizen id
		var coActk string
		if customerOnline.Co_citizen_id == "" {
			coActk = middlewares.CreateAccessToken(customerOnline.ID, 0, 0, 0, "", "")
		} else {
			var customer models.Customer
			if errC := models.GetCustomerLockShop(customerOnline.Co_citizen_id, payload.Shop_id, &customer); errC != nil {
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
					"message": "Customer Invalid!",
					"data":    "",
				})
				return
			}

			// add log
			if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
				Line_id:        "",
				Log_ip_address: c.ClientIP(),
				Log_browser:    c.GetHeader("User-Agent"),
				Log_text:       "Register & Send Verify Email.",
				Log_create:     time.Now().Format("2006-01-02 15:04:05"),
			}); errL != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Create Log Customer Online OTP Error.",
					"data":    "",
				})
				return
			}
			coActk = middlewares.CreateAccessToken(customerOnline.ID, customer.ID, customer.ShopId, customer.ShopMotherId, customerOnline.Co_line_id, customer.CtmCitizenId)
		}

		// create access token
		// var coActk = middlewares.CreateAccessToken(customerOnline.ID, customer.ID, customer.ShopId, customer.ShopMotherId, customerOnline.Co_line_id, customerOnline.Co_citizen_id)
		// if(customerOnline.Co_citizen_id == ""){
		// 	coActk = middlewares.CreateAccessToken(customerOnline.ID, 0, 0, 0, customerOnline.Co_line_id, "")
		// }else{
		// 	coActk = middlewares.CreateAccessToken(customerOnline.ID, 0, 0, 0, customerOnline.Co_line_id, customerOnline.Co_citizen_id)
		// }
		//add log
		// if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
		// 	Line_id:        customerOnline.Co_line_id,
		// 	Log_ip_address: c.ClientIP(),
		// 	Log_browser:    c.GetHeader("User-Agent"),
		// 	Log_text:       "Login OTP by Online",
		// 	Log_create:     time.Now().Format("2006-01-02 15:04:05"),
		// }); errL != nil {
		// 	c.AbortWithStatusJSON(200, gin.H{
		// 		"status":  false,
		// 		"message": "Create Log Customer Error.",
		// 		"data":    "",
		// 	})
		// 	return
		// }

		responseData := structs.ResponseOauth{
			AccessToken: coActk,
		}

		c.JSON(200, gin.H{
			"status":  true,
			"message": "",
			"data":    responseData,
		})

	}
}

// ----------------test remove บัญชี EXA กับ LINE (EXA MED) ----------------------//
func RemoveSyncAccountWithLine(c *gin.Context) {
	var payload structs.TestRemoveAccountWithLineId

	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	var co_remove structs.TestRemoveAccountWithLineId

	if errU := models.TestRemoveAccountWithLine(payload.Co_line_id, &co_remove); errU != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "remove Customer Online Error.",
			"data":    "",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		// "message": "remove Customer Online successfully.",
		"message": payload,
	})

}

// ----------------------- test TestChecklineid ----------------------------
func TestChecklineid(c *gin.Context) {

	var payload structs.PayloadChecklineid

	if errSBJ := c.ShouldBindJSON(&payload); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    "",
		})
		return
	}

	// get customer online by line id
	var customerOnline structs.CustomerOnline
	if errCO := models.GetCustomerOnlineByLineId(payload.Line_id, &customerOnline); errCO != nil {
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

	// get customer first in shop by citizen id
	// var customer models.Customer
	// if errC := models.GetCustomerFirstByCitizenId(customerOnline.Co_citizen_id, &customer); errC != nil {
	// 	c.AbortWithStatusJSON(200, gin.H{
	// 		"status":  false,
	// 		"message": "Error Get Customer Invalid!",
	// 		"data":    errC.Error(),
	// 	})
	// 	return
	// }
	// if customer.ID == 0 {
	// 	c.AbortWithStatusJSON(200, gin.H{
	// 		"status":  false,
	// 		"message": "Customer Invalid!",
	// 		"data":    "",
	// 	})
	// 	return
	// }

	// var coActk = middlewares.CreateAccessToken(customerOnline.ID, customer.ID, customer.ShopId, customer.ShopMotherId, customerOnline.Co_line_id, customerOnline.Co_citizen_id)
	var coActk = middlewares.CreateAccessToken(customerOnline.ID, 0, 0, 0, customerOnline.Co_line_id, customerOnline.Co_citizen_id)

	//add log
	if errL := models.CreateLogOauth(&structs.ObjLogCustomerLogin{
		Line_id:        customerOnline.Co_line_id,
		Log_ip_address: c.ClientIP(),
		Log_browser:    c.GetHeader("User-Agent"),
		Log_text:       "Login by Line ID",
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

	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    payload,
	})
}

