package main

import (
	"fmt"
	"linecrmapi/configs"
	"linecrmapi/routes"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors" // นำเข้าแพ็คเกจ cors
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// connectDB เป็นฟังก์ชันช่วยสำหรับเชื่อมต่อฐานข้อมูล โดยรับค่า dbURL, dbName และ loggerMode
// คืนค่า *gorm.DB พร้อมกับฟังก์ชัน cleanup สำหรับปิด connection เมื่อโปรแกรมจบการทำงาน
func connectDB(dbURL, dbName string, loggerMode logger.LogLevel) (*gorm.DB, func()) {
	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{
		Logger: logger.Default.LogMode(loggerMode),
	})
	if err != nil {
		panic(fmt.Sprintf("ไม่สามารถเชื่อมต่อกับฐานข้อมูล %s: %v", dbName, err))
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("ไม่สามารถดึง *sql.DB สำหรับ %s: %v", dbName, err))
	}

	cleanup := func() {
		if err := sqlDB.Close(); err != nil {
			fmt.Printf("ไม่สามารถปิดการเชื่อมต่อกับ %s: %v\n", dbName, err)
		} else {
			fmt.Printf("การเชื่อมต่อกับ %s ถูกปิดแล้ว\n", dbName)
		}
	}
	return db, cleanup
}

func main() {
	// โหลดไฟล์ .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("เกิดข้อผิดพลาดในการโหลดไฟล์ .env")
	}

	// กำหนด logger mode ตาม environment
	var loggerMode logger.LogLevel
	if os.Getenv("ENV") == "PROD" {
		loggerMode = logger.Silent
	} else {
		loggerMode = logger.Info
	}

	// สร้าง slice สำหรับเก็บฟังก์ชัน cleanup ของการเชื่อมต่อฐานข้อมูล
	var cleanupFuncs []func()
	var cleanup func() // ประกาศตัวแปร cleanup

	// เชื่อมต่อฐานข้อมูล DB1
	db1URL := configs.SetDbUrl(configs.SetDb1Config())
	configs.DB1, cleanup = connectDB(db1URL, "DB1", loggerMode)
	cleanupFuncs = append(cleanupFuncs, cleanup)

	// เชื่อมต่อฐานข้อมูล DB2
	db2URL := configs.SetDbUrl(configs.SetDb2Config())
	configs.DB2, cleanup = connectDB(db2URL, "DB2", loggerMode)
	cleanupFuncs = append(cleanupFuncs, cleanup)

	// เชื่อมต่อฐานข้อมูลบันทึก DBL1
	dbl1URL := configs.SetDbUrl(configs.SetDbL1Config())
	configs.DBL1, cleanup = connectDB(dbl1URL, "DBL1", loggerMode)
	cleanupFuncs = append(cleanupFuncs, cleanup)

	// เชื่อมต่อฐานข้อมูลบันทึก DBL2
	dbl2URL := configs.SetDbUrl(configs.SetDbL2Config())
	configs.DBL2, cleanup = connectDB(dbl2URL, "DBL2", loggerMode)
	cleanupFuncs = append(cleanupFuncs, cleanup)

	// Ensure cleanup ของการเชื่อมต่อฐานข้อมูลเมื่อโปรแกรมจบการทำงาน
	defer func() {
		for _, fn := range cleanupFuncs {
			fn()
		}
	}()

	// สร้าง Gin router และเพิ่ม middleware สำหรับ Logger และ Recovery
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	// ตั้งค่า trusted proxies เป็น nil
	if err := router.SetTrustedProxies(nil); err != nil {
		log.Fatal(err)
	}
	// เพิ่ม CORS Middleware จาก gin-contrib/cors
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"https://www.app-apsx.com",
			"https://app-apsx.com",
			"https://www.clinic.app-apsx.com",
			"https://clinic.app-apsx.com",
			"https://www.admin.app-apsx.com",
			"https://admin.app-apsx.com",
			"https://apsx-clinic.vercel.app",
			"https://apsx.vercel.app",
			"https://apsx-admin.vercel.app",
			"*",
		},
		AllowMethods:     []string{"POST", "OPTIONS", "GET", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Requested-With", "X-CSRF-Token", "Accept", "Origin", "X-Forwarded-For", "X-Real-IP"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// เส้นทางหลัก
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "CLINIC API OK Version 2024 By " + os.Getenv("SMS_SENDER"),
		})
	})

	// ตั้งค่าเส้นทางเพิ่มเติม
	routes.SetRouterAuth(router)
	routes.SetRouterShop(router)
	routes.SetRouterCustomer(router)
	routes.SetRouterAppointment(router)
	routes.SetRouterQueue(router)
	routes.SetRouterReceipts(router)
	routes.SetRouterService(router)
	routes.SetRouterItem(router)
	routes.SetRouterOrders(router)
	routes.SetRouterInvoices(router)
	routes.SetRouterProcesss(router)
	routes.SetRouterOrderonline(router)

	// เริ่มต้นการทำงานของเซิร์ฟเวอร์
	if err := router.Run(os.Getenv("API_PORT")); err != nil {
		panic(fmt.Sprintf("[Error] ไม่สามารถเริ่มต้นเซิร์ฟเวอร์ Gin เนื่องจาก: %s", err.Error()))
	}
	// คำสั่งรัน: gin -p 8100 -a 8101 run main.go
}
