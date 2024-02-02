package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"new_demo/middleware"
	"new_demo/models"
	"new_demo/utils/errmsg"
)

type LoginServer struct {
}

func (L LoginServer) Login(c *gin.Context) {
	var data models.Manager
	_ = c.ShouldBind(&data)
	code1 := models.CheckLogin(
		data.Username,
		data.Password)
	if code1 == errmsg.ERROR_USER_NOT_EXIST {
		c.JSON(http.StatusOK, gin.H{
			"code":    code1,
			"message": errmsg.GetCode(code1),
		})
		return
	}
	if code1 == errmsg.ERROR_PASSWORD_WRONG {
		c.JSON(http.StatusOK, gin.H{
			"code":    code1,
			"message": errmsg.GetCode(code1),
		})
		return
	}
	token, code2 := middleware.GetToken(data.Username)
	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"code":    code2,
		"message": errmsg.GetCode(code2),
	})
}
