package routes

import (
	"linecrmapi/controllers"
	"linecrmapi/middlewares"

	"github.com/gin-gonic/gin"
)

func SetRouterAuth(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("checklineid", middlewares.CheckPublicKey, controllers.Checklineid)
		auth.POST("login", middlewares.CheckPublicKey, controllers.Login)
		auth.GET("logout/:lineId", middlewares.CheckPublicKey, controllers.Logout)
		auth.POST("otp", middlewares.CheckPublicKey, controllers.Otp)
		auth.POST("register", middlewares.CheckPublicKey, controllers.Register)
		auth.POST("registertelonly", middlewares.CheckPublicKey, controllers.RegisterTelOnly)
		auth.POST("otponline", middlewares.CheckPublicKey, controllers.Otponline)
		auth.POST("otponlinetelonly", middlewares.CheckPublicKey, controllers.OtponlineTelOnly)
		auth.POST("loginonline", middlewares.CheckPublicKey, controllers.Loginonline)
		auth.POST("loginonlinetel", middlewares.CheckPublicKey, controllers.LoginonlineTelCheck)
		auth.POST("otponlinesync", middlewares.CheckAccessToken, controllers.Otponlinesync)
		auth.POST("onlinesync", middlewares.CheckAccessToken, controllers.Onlinesync)
		auth.GET("onlineunsync", middlewares.CheckAccessToken, controllers.Onlineunsync)

		  auth.POST("testchecklineid", middlewares.CheckPublicKey, controllers.TestChecklineid)
  auth.POST("removeaccountsyncwithmail", middlewares.CheckPublicKey, controllers.RemoveSyncAccountWithLine)

		// Exa med
		auth.POST("registerexa", middlewares.CheckPublicKey, controllers.RegisterExa)
		//no 2fa login
		auth.POST("loginexa", middlewares.CheckPublicKey, controllers.LoginExaNoOtp)
		// login with 2fa
		auth.POST("loginonlineexa", middlewares.CheckPublicKey, controllers.LoginonlineExa)
		auth.POST("accountsyncwithmail", middlewares.CheckPublicKey, controllers.SyncAccountWithLine)
		auth.POST("verifyemail", middlewares.CheckPublicKey, controllers.VerifyEmail)
		auth.POST("onlinesyncexa", middlewares.CheckAccessToken, controllers.OnlinesyncExa)
		auth.POST("otponlinesyncexa", middlewares.CheckAccessToken, controllers.OtponlinesyncExa)
		auth.POST("otponlinesyncexa-lock", middlewares.CheckAccessToken, controllers.OtponlinesyncExaLockShop)
		auth.POST("registerCid", middlewares.CheckPublicKey, controllers.RegisterBypassOTP)

		

		//Health Survey
		auth.POST("checkhealthsurvey", middlewares.CheckPublicKey, controllers.HealthSurveyAuth)
		auth.POST("otphealthsurvey", middlewares.CheckPublicKey, controllers.OtpHealthSurvey)
		
		// auth.POST("callback", middlewares.CheckPublicKey, controllers.Oauth)
	}
	return
}

/**
 * @apiGroup Authen
 * @api {post} auth/checklineid 1.Authen Check LineId
 * @apiDescription เข้าสู่ระบบ Authen check line id
 * @apiVersion 1.0.0
 * @apiHeader   {String}    Access-Token   Public Token
 * @apiBody {String}   line_id ไอดีไลน์
 * @apiSuccess {Boolean}    response สถานะ : true,false
 * @apiSuccess {String} message ข้อความตอบกลับ
 * @apiSuccess {Object[]} data
 * @apiSuccess {String} data.access_token 	access token สำหรับการเข้าถึงข้อมูลของ customer
 */

/**
 * @apiGroup Authen
 * @api {post} auth/login 2.Authen Login
 * @apiDescription register line id & login Waiting OTP
 * @apiVersion 1.0.0
 * @apiHeader   {String}    Access-Token   Public Token
 * @apiBody {String}   line_id ไอดีไลน์
 * @apiBody {String}   line_name ชื่อไลน์
 * @apiBody {String}   line_email อีเมลไลน์
 * @apiBody {String}   citizen_id  รหัสประชาชน
 * @apiBody {String}   tel	เบอร์โทรศัพท์
 * @apiSuccess {Boolean}    response 	สถานะ : true,false
 * @apiSuccess {String} 	message 	Add Customer Online Waiting OTP
 * @apiSuccess {String} 	data	""
 */

/**
 * @apiGroup Authen
 * @api {post} auth/otp 3.Authen OTP
 * @apiDescription เข้าสู่ระบบ OTP Authen
 * @apiVersion 1.0.0
 * @apiHeader   {String}    Access-Token   Public Token
 * @apiBody {String}   line_id ไอดีไลน์
 * @apiBody {String}   tel	เบอร์โทรศัพท์
 * @apiBody {String}   otp	รหัส OTP
 * @apiSuccess {Boolean}    response สถานะ : true,false
 * @apiSuccess {String} message ข้อความตอบกลับ
 * @apiSuccess {String} 	data	""
 * @apiSuccess {String} data.access_token 	access token สำหรับการเข้าถึงข้อมูลของ customer
 */

/**
 * @apiGroup Online Authen
 * @api {post} auth/register 1.Register
 * @apiDescription Register สมัครสมาชิก
 * @apiVersion 2.0.0
 * @apiHeader   {String}    Authorization   Bearer Token   Public Token
 * @apiBody {String}   co_email อีเมลลูกค้า
 * @apiBody {String}   co_password	รหัสผ่าน
 * @apiBody {String}   co_tel	เบอร์โทรศัพท์
 * @apiBody {String}   co_prefix คำนำหน้า
 * @apiBody {String}   co_fname	ชื่อ
 * @apiBody {String}   co_lname สกุล
 * @apiSuccess {Boolean}    response สถานะ : true,false
 * @apiSuccess {String} message ข้อความตอบกลับ
 * @apiSuccess {String} 	data	""
 */

/**
 * @apiGroup Online Authen
 * @api {post} auth/otponline 2.OTP Online
 * @apiDescription OTP Online
 * @apiVersion 2.0.0
 * @apiHeader   {String}   Authorization   Bearer Token   Public Token
 * @apiBody {String}   co_email อีเมลลูกค้า
 * @apiBody {String}   co_tel	เบอร์โทรศัพท์
 * @apiBody {String}   otp รหัส OTP
 * @apiSuccess {Boolean}    response สถานะ : true,false
 * @apiSuccess {String} message ข้อความตอบกลับ
 * @apiSuccess {Object[]} data
 * @apiSuccess {String} data.access_token 	access token สำหรับการเข้าถึงข้อมูลของ customer
 */

/**
 * @apiGroup Online Authen
 * @api {post} auth/loginonline 3.Login Online
 * @apiDescription Login Online เข้าสู่ระบบ
 * @apiVersion 2.0.0
 * @apiHeader   {String}    Authorization   Bearer Token  Public Token
 * @apiBody {String}   co_email อีเมลลูกค้า
 * @apiBody {String}   co_password	รหัสผ่าน
 * @apiSuccess {Boolean}    response สถานะ : true,false
 * @apiSuccess {String} message ข้อความตอบกลับ
 * @apiSuccess {String} co_tel เบอร์โทรศัพท์
 * @apiSuccess {String} data
 */

/**
 * @apiGroup Online Sync
 * @api {post} auth/onlinesync 1.Sync
 * @apiDescription Sync ช้อมูล
 * @apiVersion 2.0.0
 * @apiHeader   {String}    Authorization   Bearer Token "Access Tokens"
 * @apiBody {String}   co_citizen_id เลขบัตรประชาชน
 * @apiSuccess {Boolean}    response สถานะ : true,false
 * @apiSuccess {String} message ข้อความตอบกลับ
 * @apiSuccess {String} co_tel เบอร์โทรศัพท์
 * @apiSuccess {String} data
 */

/**
 * @apiGroup Online Sync
 * @api {post} auth/onlinesync 2.OTP Online Sync
 * @apiDescription OTP Online Sync ช้อมูล
 * @apiVersion 2.0.0
 * @apiHeader   {String}    Authorization   Bearer Token "Access Tokens"
 * @apiBody {String}   otp รหัส OTP
 * @apiSuccess {Boolean}    response สถานะ : true,false
 * @apiSuccess {String} message ข้อความตอบกลับ
 * @apiSuccess {Object[]} data
 * @apiSuccess {String} data.access_token 	access token สำหรับการเข้าถึงข้อมูลของ customer
 */

/**
 * @apiGroup Online Sync
 * @api {get} auth/onlineunsync 2.Unsync
 * @apiDescription Unsync ปลดเลขบัตรประชาชน
 * @apiVersion 2.0.0
 * @apiHeader   {String}    Authorization   Bearer Token "Access Tokens"
 * @apiSuccess {Boolean}    response สถานะ : true,false
 * @apiSuccess {String} message ข้อความตอบกลับ
 * @apiSuccess {Object[]} data
 * @apiSuccess {String} data.access_token 	access token สำหรับการเข้าถึงข้อมูลของ customer
 */
