package libs

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	// "encoding/binary"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/go-mail/mail"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/juunini/simple-go-line-notify/notify"
)

func StrToInt(str string) int {
	str = strings.ReplaceAll(str, ",", "")
	result, err := strconv.Atoi(str)
	if err != nil {
		return 0
	} else {
		return result
	}
}

func StrToFloat(str string) float64 {
	str = strings.ReplaceAll(str, ",", "")
	result, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	} else {
		return result
	}
}

type S3DeleteObjectAPI interface {
	DeleteObject(ctx context.Context,
		params *s3.DeleteObjectInput,
		optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error)
}

func DeleteItem(c context.Context, api S3DeleteObjectAPI, input *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	return api.DeleteObject(c, input)
}

func DeleteS3(part string) int {
	// Load env variable
	evnErr := godotenv.Load()
	if evnErr != nil {
		log.Fatal("Error loading .env file")
	}
	// Setup s3 config
	cfg, _ := config.LoadDefaultConfig(context.TODO())
	client := s3.NewFromConfig(cfg)
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Key:    aws.String(part),
	}
	// Delete s3
	_, errd := DeleteItem(context.TODO(), client, input)
	if errd != nil {
		return 0
	} else {
		return 1
	}

}

func SendLineNoti(accessToken string, message string, imageURL string) bool {
	if imageURL != "" {
		if err := notify.SendImage(accessToken, message, imageURL); err != nil {
			return false
		} else {
			return true
		}
	} else {
		if err := notify.SendText(accessToken, message); err != nil {
			return false
		} else {
			return true
		}
	}
}

func SendMail(toEmail string, titleMsg string, msgHtml string) error {

	from := os.Getenv("EMAIL_NAME")
	password := os.Getenv("EMAIL_PWD")

	m := mail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", titleMsg)
	m.SetBody("text/html", msgHtml)

	d := mail.NewDialer("smtp.gmail.com", 587, from, password)

	if err := d.DialAndSend(m); err != nil {
		// panic(err)
		return err
	}
	return nil
}

// upload s3
func UploadImageS3(base64Image string, filename string, folder string) string {

	var imageUrl = ""

	decodeImage, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		fmt.Println("Decode string error : ", err)
		return imageUrl
	}

	contentType := http.DetectContentType(decodeImage)
	if !CheckImageExtension(contentType) {
		fmt.Println("Content type error.")
		return imageUrl
	}
	extension := strings.ReplaceAll(contentType, "image/", ".")

	newFilename := ""
	if filename == "" {
		prefix := ""
		switch folder {
		case "shop":
			prefix = "S"
		case "customer":
			prefix = "C"
		default:
			prefix = ""
		}
		uuid := uuid.New()
		newFilename = fmt.Sprintf("%s%s%s", prefix, uuid.String(), extension)
	} else {
		newFilename = fmt.Sprintf("%s%s", filename, extension)
	}

	errLoadEnv := godotenv.Load()
	if errLoadEnv != nil {
		fmt.Println("Get env error : ", errLoadEnv)
		return imageUrl
	}

	awsConfig, errLDC := config.LoadDefaultConfig(context.TODO())
	if errLDC != nil {
		fmt.Println("Aws config error : ", errLDC)
		return imageUrl
	}

	awsClient := s3.NewFromConfig(awsConfig)
	awsUpload := manager.NewUploader(awsClient)
	result, errUpload := awsUpload.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(os.Getenv("AWS_BUCKET")),
		Key:         aws.String(folder + "/" + newFilename),
		Body:        bytes.NewReader(decodeImage),
		ACL:         "public-read",
		ContentType: aws.String(contentType),
	})
	if errUpload != nil {
		fmt.Println("Aws upload error : ", errUpload)
		return imageUrl
	} else {
		imageUrl = result.Location
	}

	return imageUrl
}

func UploadImageS3Filename(base64Image string, filename string, folder string) string {

	var imageUrl = ""

	decodeImage, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		fmt.Println("Decode string error : ", err)
		return imageUrl
	}

	contentType := http.DetectContentType(decodeImage)
	if !CheckImageExtension(contentType) {
		fmt.Println("Content type error.")
		return imageUrl
	}
	extension := strings.ReplaceAll(contentType, "image/", ".")

	newFilename := ""
	if filename == "" {
		prefix := ""
		switch folder {
		case "shop":
			prefix = "S"
		case "customer":
			prefix = "C"
		default:
			prefix = ""
		}
		uuid := uuid.New()
		newFilename = fmt.Sprintf("%s%s%s", prefix, uuid.String(), extension)
	} else {
		newFilename = fmt.Sprintf("%s%s", filename, extension)
	}

	errLoadEnv := godotenv.Load()
	if errLoadEnv != nil {
		fmt.Println("Get env error : ", errLoadEnv)
		return imageUrl
	}

	awsConfig, errLDC := config.LoadDefaultConfig(context.TODO())
	if errLDC != nil {
		fmt.Println("Aws config error : ", errLDC)
		return imageUrl
	}

	awsClient := s3.NewFromConfig(awsConfig)
	awsUpload := manager.NewUploader(awsClient)
	result, errUpload := awsUpload.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(os.Getenv("AWS_BUCKET")),
		Key:         aws.String(folder + "/" + newFilename),
		Body:        bytes.NewReader(decodeImage),
		ACL:         "public-read",
		ContentType: aws.String(contentType),
	})
	if errUpload != nil {
		fmt.Println("Aws upload error : ", errUpload)
		return imageUrl
	} else {
		imageUrl = result.Location
	}

	return imageUrl
}

func CheckImageExtension(extension string) bool {
	fmt.Println("===========================================================")
	fmt.Println(extension)
	fmt.Println("===========================================================")
	var result bool
	switch extension {
	case "image/jpg":
		result = true
	case "image/jpeg":
		result = true
	case "image/gif":
		result = true
	case "image/png":
		result = true
	default:
		result = false
	}
	return result
}

func UploadFileS3(base64File string, folder string) (string, int) {

	var fileUrl = ""
	var fileSize = 0

	decodeFile, err := base64.StdEncoding.DecodeString(base64File)
	if err != nil {
		fmt.Println("Decode string error : ", err)
		return fileUrl, fileSize
	}

	fileSize = len(decodeFile) / 1024

	contentType := http.DetectContentType(decodeFile)

	if !CheckFileExtension(contentType) {
		fmt.Println("Content type error.")
		return fileUrl, fileSize
	}

	extension := strings.Split(contentType, "/")

	prefix := ""
	switch folder {
	case "shop":
		prefix = "S"
	case "customer":
		prefix = "C"
	default:
		prefix = ""
	}
	uuid := uuid.New()
	filename := fmt.Sprintf("%s%s%s", prefix, uuid.String(), "."+extension[1])

	errLoadEnv := godotenv.Load()
	if errLoadEnv != nil {
		fmt.Println("Get env error : ", errLoadEnv)
		return fileUrl, fileSize
	}

	awsConfig, errLDC := config.LoadDefaultConfig(context.TODO())
	if errLDC != nil {
		fmt.Println("Aws config error : ", errLDC)
		return fileUrl, fileSize
	}

	awsClient := s3.NewFromConfig(awsConfig)
	awsUpload := manager.NewUploader(awsClient)
	result, errUpload := awsUpload.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(os.Getenv("AWS_BUCKET")),
		Key:         aws.String(folder + "/" + filename),
		Body:        bytes.NewReader(decodeFile),
		ACL:         "public-read",
		ContentType: aws.String(contentType),
	})
	if errUpload != nil {
		fmt.Println("Aws upload error : ", errUpload)
		return fileUrl, fileSize
	} else {
		fileUrl = result.Location
	}
	fmt.Println(fileUrl)

	return fileUrl, fileSize

}

func UploadFileFilenameS3(base64File string, folder string, fname string) (string, int) {

	var fileUrl = ""
	var fileSize = 0

	decodeFile, err := base64.StdEncoding.DecodeString(base64File)
	if err != nil {
		fmt.Println("Decode string error : ", err)
		return fileUrl, fileSize
	}

	fileSize = len(decodeFile) / 1024

	contentType := http.DetectContentType(decodeFile)

	if !CheckFileExtension(contentType) {
		fmt.Println("Content type error.")
		return fileUrl, fileSize
	}

	extension := strings.Split(contentType, "/")
	filename := fmt.Sprintf("%s%s", fname, "."+extension[1])

	errLoadEnv := godotenv.Load()
	if errLoadEnv != nil {
		fmt.Println("Get env error : ", errLoadEnv)
		return fileUrl, fileSize
	}

	awsConfig, errLDC := config.LoadDefaultConfig(context.TODO())
	if errLDC != nil {
		fmt.Println("Aws config error : ", errLDC)
		return fileUrl, fileSize
	}

	awsClient := s3.NewFromConfig(awsConfig)
	awsUpload := manager.NewUploader(awsClient)
	result, errUpload := awsUpload.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(os.Getenv("AWS_BUCKET")),
		Key:         aws.String(folder + "/" + filename),
		Body:        bytes.NewReader(decodeFile),
		ACL:         "public-read",
		ContentType: aws.String(contentType),
	})
	if errUpload != nil {
		fmt.Println("Aws upload error : ", errUpload)
		return fileUrl, fileSize
	} else {
		fileUrl = result.Location
	}
	fmt.Println(fileUrl)

	return fileUrl, fileSize

}

func CheckFileExtension(extension string) bool {
	var result bool
	switch extension {
	case "image/jpg":
		result = true
	case "image/jpeg":
		result = true
	case "image/gif":
		result = true
	case "image/png":
		result = true
	case "application/pdf":
		result = true
	default:
		result = false
	}
	return result
}

// array diff
func CheckArrayInt(slice []int, elem int) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}

func GetArrayIntDiff(arr1, arr2 []int) []int {
	var diff []int
	for _, v := range arr2 {
		if !CheckArrayInt(arr1, v) {
			diff = append(diff, v)
		}
	}
	return diff
}

func CheckArrayStr(slice []string, elem string) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}

func GetArrayStrDiff(arr1, arr2 []string) []string {
	var diff []string
	for _, v := range arr2 {
		if !CheckArrayStr(arr1, v) {
			diff = append(diff, v)
		}
	}
	return diff
}

// doc code
func SetDocSettingCode(codeDefault string, numberDigit int, numberDefault int, dateFormatId int) string {
	dateFormat := ""
	switch dateFormatId {
	case 1:
		dateFormat = ""
		break
	case 2: // YYYYx
		dateFormat = time.Now().Format("2006")
		break
	case 3: // YYYYMMx
		dateFormat = time.Now().Format("200601")
		break
	case 4: // YYYYMMDDx
		dateFormat = time.Now().Format("20060102")
		break
	}
	return codeDefault + dateFormat + (fmt.Sprintf("%0"+strconv.Itoa(numberDigit)+"d", numberDefault))
}

func UploadExcel(base64File string) string {

	// decodeFile, decodeFileError := base64.StdEncoding.DecodeString(base64File)
	// if decodeFileError != nil {
	// 	fmt.Println("Decode base64 error.")
	return ""
	// }

	// // contentType := http.DetectContentType(decodeFile)
	// // extension := strings.Split(contentType, "/")
	// // fmt.Println(decodeFile)
	// // fmt.Println(extension)
	// // if extension[1] != ".xls" && extension[1] != ".xlsx" {
	// // 	fmt.Println("Content type error.")
	// // 	return ""
	// // }

	// uuid := uuid.New()
	// filePath := fmt.Sprintf("uploads/excels/%s%s", uuid.String(), ".xls")
	// file, err := os.Create(filePath)
	// if err != nil {
	// 	fmt.Println("Create file error.")
	// 	return ""
	// }
	// defer file.Close()

	// _, err = file.Write(decodeFile)
	// if err != nil {
	// 	fmt.Println("Write file error.")
	// 	return ""
	// }

	// fmt.Println(filePath)
	// excelFile, err := excelize.OpenFile(filePath)
	// if err != nil {
	// 	fmt.Println("Open file error.")
	// 	return ""
	// }

	// defer func() {
	// 	if err := excelFile.Close(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()

	// sheets := excelFile.GetSheetList()
	// for _, sheetName := range sheets {
	// 	rows, err := excelFile.GetRows(sheetName)
	// 	if err != nil {
	// 		fmt.Println("error reading sheet", sheetName, ":", err)
	// 		return ""
	// 	}

	// 	fmt.Println(rows)

	// 	// saveAsJSON(d, sheetName+".json")
	// 	// saveAsJSONWithHeaders(d, sheetName+"_with_headers.json")
	// }

	// // return ""

	// return filePath

}

func CalculateAge(birthDate time.Time) (int, int) {
	currentDate := time.Now()

	// Calculate the difference in years.
	ageYears := currentDate.Year() - birthDate.Year()

	// Adjust the age if the birth month hasn't occurred yet in the current year.
	if birthDate.Month() > currentDate.Month() ||
		(birthDate.Month() == currentDate.Month() && birthDate.Day() > currentDate.Day()) {
		ageYears--
	}

	// Calculate the difference in months.
	months := int(currentDate.Month()) - int(birthDate.Month())

	// Adjust the months to be positive.
	if months < 0 {
		months += 12
	}

	return ageYears, months
}

func Encrypt(text string) (string, error) {
	key := []byte("702b6628c691204ea5385602cafad01f")
	plaintext := []byte(text)

	// PKCS#7 padding
	blockSize := aes.BlockSize
	padding := blockSize - len(plaintext)%blockSize
	padText := append(plaintext, bytes.Repeat([]byte{byte(padding)}, padding)...)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, blockSize+len(padText))
	iv := ciphertext[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[blockSize:], padText)

	// Encode ciphertext using Base64 URL encoding
	encoded := base64.StdEncoding.EncodeToString(ciphertext)

	// Replace '/' with another character
	strRp := strings.ReplaceAll(encoded, "/", "-")
	strRp = strings.ReplaceAll(strRp, "+", "__")

	return strRp, nil
}

func Decrypt(encoded string) (string, error) {
	key := []byte("702b6628c691204ea5385602cafad01f")
	strRp := strings.ReplaceAll(encoded, "-", "/")
	strRp = strings.ReplaceAll(strRp, "__", "+")

	ciphertext, err := base64.StdEncoding.DecodeString(strRp)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	str := strings.ReplaceAll(string(ciphertext), "\r", "")
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "\n", "")

	return str, nil
}
func RenderEmailRequestOtp(message, fullname, otpCode, otpKey string) string {
	html := "<div>"
	html = html + "<p><b>เรียนคุณ " + fullname + "</b></p>"
	html = html + "<p>นี่เป็นรหัสยืนยันตัวตนสำหรับ" + message + "บัญชี EXA MED+ ของคุณ ( โปรดเก็บเป็นความลับ )</p>"
	html = html + fmt.Sprintf("<p><h1>รหัสยืนยันตัวตนของคุณคือ: %s</h1></p>", otpCode)
	html = html + fmt.Sprintf("<p>รหัสอ้างอิง: %s</p>", otpKey)
	html = html + "<p>รหัสยืนยันตัวตนนี้จะใช้ได้เพียงครั้งเดียวเท่านั้น และจะหมดอายุภายใน 5 นาที</p>"
	html = html + "<p>หากคุณไม่ได้ทำรายการนี้ กรุณาติดต่อเราทันที</p>"
	html = html + "<p>ขอบคุณ<br>---</p>"
	html = html + "<p>ติดต่อ<br>โทร: --- | อีเมล: ---</p>"
	html = html + "<div>"
	return html
}



func RandStringBytesMaskImpr(letterBytes string, n int) string {
    b := make([]byte, n)
    // สร้าง buffer สำหรับเก็บข้อมูลสุ่มทั้งหมดที่ต้องการ
    randomBytes := make([]byte, n*2) // ใช้ n*2 เพื่อให้มีข้อมูลเพียงพอ
    if _, err := rand.Read(randomBytes); err != nil {
        panic(err)
    }
    
    for i := 0; i < n; i++ {
        // ใช้ 2 bytes สำหรับแต่ละตัวอักษร เพื่อให้มีค่าสุ่มที่มากพอ
        randomValue := uint16(randomBytes[i*2]) | uint16(randomBytes[i*2+1])<<8
        // เลือกตัวอักษรจาก letterBytes โดยใช้ modulo
        b[i] = letterBytes[int(randomValue)%len(letterBytes)]
    }
    
    return string(b)
}