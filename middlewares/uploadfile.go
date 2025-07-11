package middlewares

import (
	"context"
	"fmt"
	"image"
	"log"
	"strconv"

	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func UploadS3(c *gin.Context) {
	// Load env variable
	evnErr := godotenv.Load()
	if evnErr != nil {
		log.Fatal("Error loading .env file")
	}
	// Setup gin app
	folder, _ := c.GetPostForm("folder")
	file, err := c.FormFile("path")
	if err != nil {
		fmt.Println("------ no image ------")
		c.AddParam("filename", "")
		c.AddParam("location", "")
		c.AddParam("size", "")
		c.Next()
		return
	}
	// Setup s3 uploader
	cfg, cfgErr := config.LoadDefaultConfig(context.TODO())
	if cfgErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"massage": "Uplaod is config error.",
			"data":    err,
		})
		return
	}
	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)
	// Save the file
	f, openErr := file.Open()
	if openErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"massage": "Uplaod is open error.",
			"data":    err,
		})
		return
	}
	contentType := file.Header.Values("Content-Type")[0]
	size := file.Size
	result, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(os.Getenv("AWS_BUCKET")),
		Key:         aws.String(folder + "/" + file.Filename),
		Body:        f,
		ACL:         "public-read",
		ContentType: aws.String(contentType),
	})
	if uploadErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"massage": "Uplaod is error.",
			"data":    uploadErr,
		})
	} else {
		c.AddParam("filename", file.Filename)
		c.AddParam("location", result.Location)
		c.AddParam("size", strconv.FormatInt(size, 10))
		c.Next()
	}
}

func UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"massage": err.Error(),
		})
		return
	}
	fileExt := filepath.Ext(header.Filename)

	if allowFileType(fileExt) == false {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"massage": "File Type not allow " + fileExt,
		})
		return
	} else {
		fmt.Println(fileExt)
		// originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), filepath.Ext(header.Filename))
		now := time.Now()
		filename := fmt.Sprintf("%v", now.Unix()) + fmt.Sprintf("%v", fileExt)
		filePath := "/Upload/images/" + filename

		imageFile, _, err := image.Decode(file)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"massage": err.Error(),
			})
			return
		}
		src := imaging.Resize(imageFile, 800, 0, imaging.Lanczos)
		err = imaging.Save(src, fmt.Sprintf("Upload/images/%v", filename))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"massage": "failed to save image " + err.Error(),
			})
			return
		}
		c.AddParam("imgPath", filePath)
		c.Next()
	}
}

func allowFileType(mimeType string) bool {
	var result bool
	switch mimeType {
	case ".jpeg", ".jpg":
		result = true
		break
	case ".gif":
		result = true
		break
	case ".png":
		result = true
		break
	default:
		result = true
		break
	}
	return result
}

func UploadExcel(c *gin.Context) {

	fileExcel, header, err := c.Request.FormFile("file_excel")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"massage": "Read file excel error.",
		})
		return
	}

	fileType := filepath.Ext(header.Filename)
	if fileType != ".xls" && fileType != ".xlsx" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"massage": "File type error.",
		})
		return
	} else {

		uuid := uuid.New()
		filePath := fmt.Sprintf("uploads/excels/%s%s", uuid.String(), fileType)
		fileCreate, err := os.Create(filePath)
		if err != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Create file error.",
			})
			return
		}

		defer os.Remove(filePath)
		defer fileCreate.Close()

		_, err = io.Copy(fileCreate, fileExcel)
		if err != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status":  false,
				"message": "Copy file error.",
			})
			return
		}

		c.AddParam("file_path", filePath)

		c.Next()
	}
}

// fix
func UploadFileS3(c *gin.Context) {

	folder, _ := c.GetPostForm("folder")
	path, errFF := c.FormFile("path")
	if errFF != nil {
		c.AddParam("filename", "")
		c.AddParam("location", "")
		c.Next()
		return
	}

	extension := filepath.Ext(path.Filename)
	if !checkExtension(extension) {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"massage": "Extension file error.",
			"data":    "",
		})
		return
	}

	uuid := uuid.New()
	filename := ""
	if folder == "customer" {
		filename = fmt.Sprintf("C%s%s", uuid.String(), extension)
	}

	file, errOF := path.Open()
	if errOF != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"massage": "Load file error.",
			"data":    "",
		})
		return
	}

	errLoadEnv := godotenv.Load()
	if errLoadEnv != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"massage": "Load env error.",
			"data":    "",
		})
		return
	}

	awsConfig, errLDC := config.LoadDefaultConfig(context.TODO())
	if errLDC != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"massage": "Load config error.",
			"data":    "",
		})
		return
	}

	awsClient := s3.NewFromConfig(awsConfig)
	awsUpload := manager.NewUploader(awsClient)
	result, errUpload := awsUpload.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(os.Getenv("AWS_BUCKET")),
		Key:         aws.String(folder + "/" + filename),
		Body:        file,
		ACL:         "public-read",
		ContentType: aws.String(path.Header.Values("Content-Type")[0]),
	})
	if errUpload != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"massage": "Upload aws error.",
			"data":    "",
		})
	} else {
		c.AddParam("filename", filename)
		c.AddParam("location", result.Location)
		c.Next()
	}
}

func checkExtension(extension string) bool {
	var result bool
	switch extension {
	case ".jpg":
		result = true
	case ".jpeg":
		result = true
	case ".gif":
		result = true
	case ".png":
		result = true
	default:
		result = false
	}
	return result
}
