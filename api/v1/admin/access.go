package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"new_demo/models"
	"new_demo/utils"
	"new_demo/utils/errmsg"
)

type AccessServer struct {
}

// AddAccess 添加权限
func (a AccessServer) AddAccess(c *gin.Context) {
	access := models.Access{}
	_ = c.ShouldBind(&access)
	access.CreatedAt = utils.GetUnix()
	code := models.AddAccess(&access)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetCode(code),
	})
}

// EditAccess 修改权限
func (a AccessServer) EditAccess(c *gin.Context) {

}

// GetAccess 获取权限
func (a AccessServer) GetAccess(c *gin.Context) {

}

// GetTopModule 获取顶级模块名称
func (a AccessServer) GetTopModule(c *gin.Context) {
	module, code := models.GetTopModule()
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    module,
		"message": errmsg.GetCode(code),
	})
}

// DeleteAccess 删除权限
func (a AccessServer) DeleteAccess(c *gin.Context) {

}
