package controllers

import (
	"linecrmapi/libs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadTest(c *gin.Context) {
	folder, _ := c.GetPostForm("folder")
	filename := c.Params.ByName("filename")
	location := c.Params.ByName("location")
	size := c.Params.ByName("size")
	delete := libs.DeleteS3("admin/1670376952782.jpg") // Test Delete File S3
	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"message":  "Uplaod is success.",
		"data":     folder,
		"filename": filename,
		"location": location,
		"size":     size,
		"delete":   delete,
	})
}

func SmsTest(c *gin.Context) {
	s := libs.Sms{
		Msisdn:  "0997102829",
		Message: "ทดสอบส่ง",
	}
	res, _ := s.Send()
	if res.StatusCode == 201 {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "SMS is success.",
			"code":    res.StatusCode,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "SMS is error.",
			"code":    res,
		})
	}

}
