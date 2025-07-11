package routes

import (
	"linecrmapi/controllers"
	"linecrmapi/middlewares"

	"github.com/gin-gonic/gin"
)

func SetRouterShop(router *gin.Engine) {
	r := router.Group("/shop")
	{
		r.GET("/getshop", middlewares.CheckAccessToken, controllers.GetInShopList)
		r.POST("/oauth", middlewares.CheckAccessToken, controllers.ShopOauth)
		r.GET("/getshoponline", middlewares.CheckAccessToken, controllers.GetInShopListOnline)
		r.POST("/oauthonline", middlewares.CheckAccessToken, controllers.ShopOauthOnline)
	}
	return
}

/**
 * @apiGroup Shop
 * @api {get} shop/getshop 1.My Shop List
 * @apiDescription get My Shop List
 * @apiVersion 1.0.0
 * @apiHeader   {String}    Authorization Bearer   Access Token
 * @apiSuccess {Boolean}    response สถานะ : true,false
 * @apiSuccess {String} message ข้อความตอบกลับ
 * @apiSuccess {Object[]} data
 * @apiSuccess {Int} data.id 	รหัสหลักร้านค้า
 * @apiSuccess {String} data.code 	รหัสร้านค้า
 * @apiSuccess {String} data.name 	ชื่อร้านค้า
 * @apiSuccess {String} data.nature_type 	หมวดประเภท
 * @apiSuccess {String} data.nature 	ประเภท
 * @apiSuccess {String} data.province 	จังหวัด
 * @apiSuccess {String} data.latlong 	latlong
 * @apiSuccess {String} data.image 	โลโก้ร้านค้า
 */

/**
 * @apiGroup Shop
 * @api {post} shop/oauth 2.Shop Oauth
 * @apiDescription Authen Shop Oauth Switch shop
 * @apiVersion 1.0.0
 * @apiHeader   {String}    Authorization Bearer   Access Token
 * @apiBody {Int}   shop_id รหัสร้านค้า
 * @apiSuccess {Boolean}    response 	สถานะ : true,false
 * @apiSuccess {String} 	message 	Add Customer Online Waiting OTP
 * @apiSuccess {Object[]} data
 * @apiSuccess {String} data.access_token 	access token สำหรับการเข้าถึงข้อมูลของ customer
 */

/**
 * @apiGroup Online Shop
 * @api {get} shop/getshoponline 1.My Shop List
 * @apiDescription get My Shop List Online
 * @apiVersion 1.0.0
 * @apiHeader   {String}    Authorization Bearer   Access Token
 * @apiSuccess {Boolean}    response สถานะ : true,false
 * @apiSuccess {String} message ข้อความตอบกลับ
 * @apiSuccess {Object[]} data
 * @apiSuccess {Int} data.id 	รหัสหลักร้านค้า
 * @apiSuccess {String} data.code 	รหัสร้านค้า
 * @apiSuccess {String} data.name 	ชื่อร้านค้า
 * @apiSuccess {String} data.nature_type 	หมวดประเภท
 * @apiSuccess {String} data.nature 	ประเภท
 * @apiSuccess {String} data.province 	จังหวัด
 * @apiSuccess {String} data.latlong 	latlong
 * @apiSuccess {String} data.image 	โลโก้ร้านค้า
 */

/**
 * @apiGroup Online Shop
 * @api {post} shop/oauthonline 2.Shop Oauth
 * @apiDescription Authen Shop Oauth Switch Shop Online
 * @apiVersion 1.0.0
 * @apiHeader   {String}    Authorization Bearer   Access Token
 * @apiBody {Int}   shop_id รหัสร้านค้า
 * @apiSuccess {Boolean}    response 	สถานะ : true,false
 * @apiSuccess {String} 	message 	Add Customer Online Waiting OTP
 * @apiSuccess {Object[]} data
 * @apiSuccess {String} data.access_token 	access token สำหรับการเข้าถึงข้อมูลของ customer
 */
