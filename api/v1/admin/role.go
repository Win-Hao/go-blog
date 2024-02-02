package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"new_demo/models"
	"new_demo/utils"
	"new_demo/utils/errmsg"
)

type RoleServer struct {
}

// GetRole 获取角色
func (ro RoleServer) GetRole(c *gin.Context) {
	pageSize, _ := utils.Int(c.Query("pageSize"))
	pageNum, _ := utils.Int(c.Query("pageNum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = 2
	}
	roles, code := models.GetRoles(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"data":    roles,
		"code":    code,
		"message": errmsg.GetCode(code),
	})
}

// AddRole 添加角色
func (ro RoleServer) AddRole(c *gin.Context) {
	role := models.Role{}
	_ = c.ShouldBind(&role)
	code := models.CheckRole(role.Title)
	if code == errmsg.ERROR_ROLE_USED {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": errmsg.GetCode(code),
		})
		return
	}
	code = models.AddRole(&role)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetCode(code),
	})
}

// EditRole 修改角色
func (ro RoleServer) EditRole(c *gin.Context) {
	roleInfo := models.Role{}
	_ = c.ShouldBind(&roleInfo)
	id := roleInfo.ID
	role, _ := models.CheckRoleById(id)
	if role.Title != roleInfo.Title {
		code := models.CheckRole(roleInfo.Title)
		if code == errmsg.ERROR_ROLE_USED {
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetCode(code),
			})
			return
		}
	}
	code := models.EditRole(&roleInfo, id)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetCode(code),
	})
}

// DeleteRole 删除角色
func (ro RoleServer) DeleteRole(c *gin.Context) {
	id, _ := utils.Int(c.Param("id"))
	_, code := models.CheckRoleById(id)
	if code == errmsg.ERROR_ROLE_NOT_EXIST {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": errmsg.GetCode(code),
		})
		return
	}
	code = models.DeleteRole(id)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetCode(code),
	})
}
