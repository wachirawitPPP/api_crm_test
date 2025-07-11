package configs

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/gorm"
)

var DB1 *gorm.DB
var DB2 *gorm.DB
var DBL1 *gorm.DB
var DBL2 *gorm.DB

// DBConfig เก็บการตั้งค่าการเชื่อมต่อฐานข้อมูล
type DBConfig struct {
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
}

// newDBConfig อ่านค่า environment variables จาก prefix ที่กำหนด และคืนค่า *DBConfig
func newDBConfig(prefix string) *DBConfig {
	// อ่านค่าจาก environment variables
	host := os.Getenv(prefix + "_HOST")
	portStr := os.Getenv(prefix + "_PORT")
	dbName := os.Getenv(prefix + "_NAME")
	user := os.Getenv(prefix + "_USER")
	password := os.Getenv(prefix + "_PWD")

	// แปลงค่า port และใช้ default 3306 หากแปลงไม่สำเร็จ
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Printf("Invalid port value '%s' for %s, defaulting to 3306", portStr, prefix)
		port = 3306
	}

	return &DBConfig{
		Host:     host,
		Port:     port,
		DBName:   dbName,
		User:     user,
		Password: password,
	}
}

// ฟังก์ชันสำหรับสร้าง config ของฐานข้อมูลแต่ละตัว โดยระบุ prefix ของ environment variables
func SetDb1Config() *DBConfig {
	return newDBConfig("DB")
}

func SetDb2Config() *DBConfig {
	return newDBConfig("DBR")
}

func SetDbL1Config() *DBConfig {
	return newDBConfig("DBL")
}

func SetDbL2Config() *DBConfig {
	return newDBConfig("DBLR")
}

// SetDbUrl สร้าง MySQL connection URL จาก DBConfig ที่ได้รับ
func SetDbUrl(config *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)
}
