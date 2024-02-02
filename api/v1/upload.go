package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"new_demo/models"
	"new_demo/utils/errmsg"
)

type UploadServer struct {
}

func (up UploadServer) Upload(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	url, code := models.UploadFile(file, fileSize)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetCode(code),
		"url":     url,
	})
}
