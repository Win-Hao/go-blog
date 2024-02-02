package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"new_demo/models"
	"new_demo/utils"
	"new_demo/utils/errmsg"
)

type ManagerServer struct {
}

// UserExist 查询用户是否存在
func (u ManagerServer) UserExist(c *gin.Context) {

}

// AddUser 添加用户
func (u ManagerServer) AddUser(c *gin.Context) {
	var manager models.Manager
	_ = c.ShouldBind(&manager)
	code := models.CheckUser(manager.Username)
	if code == errmsg.ERROR_USERNAME_USED {
		c.JSON(http.StatusOK, gin.H{
			"code":    errmsg.ERROR_USERNAME_USED,
			"message": errmsg.GetCode(errmsg.ERROR_USERNAME_USED),
		})
		return
	}
	manager.CreatedAt = utils.GetUnix()
	code = models.AddUser(&manager)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    manager,
		"message": errmsg.GetCode(code),
	})
}

// GetUser 查询单个用户
func (u ManagerServer) GetUser(c *gin.Context) {
	id, _ := utils.Int(c.Param("id"))
	manager, code := models.GetUserById(id)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    manager,
		"message": errmsg.GetCode(code),
	})
}

// GetUsers 查询用户列表
func (u ManagerServer) GetUsers(c *gin.Context) {
	pageSize, _ := utils.Int(c.Query("pageSize"))
	pageNum, _ := utils.Int(c.Query("pageNum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = 2
	}
	users, code := models.GetUsers(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    users,
		"message": errmsg.GetCode(code),
	})
}

// EditUser 编辑用户
func (u ManagerServer) EditUser(c *gin.Context) {
	var manager models.Manager
	id, _ := utils.Int(c.Param("id"))
	_ = c.ShouldBind(&manager)
	fmt.Println(manager)
	status := models.UpdateUser(id, &manager)
	c.JSON(http.StatusOK, gin.H{
		"code":    status,
		"message": errmsg.GetCode(status),
	})
}

// DeleteUser 删除用户
func (u ManagerServer) DeleteUser(c *gin.Context) {
	id, _ := utils.Int(c.Param("id"))
	code := models.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetCode(code),
	})
}
