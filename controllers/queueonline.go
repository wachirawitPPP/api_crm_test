package controllers

import (
	"linecrmapi/libs"
	"linecrmapi/models"
	"linecrmapi/structs"

	"fmt"
	"time"
)

func LabplusAuthen(shop_id int) string {
	var token string = "token"
	var datetime string = time.Now().Format("2006-01-02 15:04:05")
	Labplus := structs.Labplus{}
	if errLabplus := models.GetShopLabplus(shop_id, &Labplus); errLabplus != nil {
		return "token"
	}
	if Labplus.Id != 0 {
		if Labplus.Lapi_token != "" {
			Lapi_token_expire, _ := time.Parse(time.RFC3339, Labplus.Lapi_token_expire)
			datetime, _ := time.Parse("2006-01-02 15:04:05", datetime)
			if Lapi_token_expire.After(datetime) {
				token = Labplus.Lapi_token
			} else {
				fmt.Println("UpdateLabplus1")
				var resultLabplusAuthen *libs.AuthResponse
				resultLabplusAuthen, errrLabplusAuthen := libs.LabplusAuthen(Labplus.Lapi_link, Labplus.Lapi_username, Labplus.Lapi_password)
				if errrLabplusAuthen != nil {
					return "token"
				}
				expire := time.Now().Add(10 * time.Minute)
				expireFormatted := expire.Format("2006-01-02 15:04:05")
				objUpdateLabplus := structs.UpdateLabplus{
					Lapi_token:        resultLabplusAuthen.Token,
					Lapi_token_expire: expireFormatted,
					Lapi_token_gen:    time.Now().Format("2006-01-02 15:04:05"),
				}
				err := models.UpdateLabplus(Labplus.Id, &objUpdateLabplus)
				if err != nil {
					return "token"
				} else {
					token = resultLabplusAuthen.Token
				}
			}
		} else {
			var resultLabplusAuthen *libs.AuthResponse
			resultLabplusAuthen, errrLabplusAuthen := libs.LabplusAuthen(Labplus.Lapi_link, Labplus.Lapi_username, Labplus.Lapi_password)
			if errrLabplusAuthen != nil {
				return resultLabplusAuthen.Error
			}
			expire := time.Now().Add(10 * time.Minute)
			expireFormatted := expire.Format("2006-01-02 15:04:05")
			objUpdateLabplus := structs.UpdateLabplus{
				Lapi_token:        resultLabplusAuthen.Token,
				Lapi_token_expire: expireFormatted,
				Lapi_token_gen:    time.Now().Format("2006-01-02 15:04:05"),
			}
			err := models.UpdateLabplus(Labplus.Id, &objUpdateLabplus)
			if err != nil {
				return "token"
			} else {
				token = resultLabplusAuthen.Token
			}
		}
	} else {
		return "token"
	}
	return token
}
