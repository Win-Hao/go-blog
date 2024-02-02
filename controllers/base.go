package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
}

func (Base BaseController) Success(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
func (Base BaseController) Error(c *gin.Context, code int, msg string, err interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":  code,
		"msg":   msg,
		"error": err,
	})
}
