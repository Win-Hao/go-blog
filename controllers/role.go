package controllers

import (
	"github.com/gin-gonic/gin"
	"new_demo/dao"
	"new_demo/models"
	"new_demo/utils"
	"strings"
)

type RoleController struct {
	BaseController
}

// Detail
// @Summary 获取角色信息
// @Description 获取角色信息
// @Tags Role
// @Accept application/json
// @Produce application/json
// @Success 200 {object} models.RoleInfo
// @Router /role/detail [get]
func (Role RoleController) Detail(c *gin.Context) {
	roleInfo, err := models.GetRoleInfo()
	if err != nil {
		Role.Error(c, 400, "获取数据失败", err)
	}
	Role.Success(c, 200, "成功获取", roleInfo)
}

// Add
// @Summary 创建角色
// @Description 添加角色
// @Tags Role
// @Accept application/json
// @Produce application/json
// @Param roleInfo body models.RoleInfo true "Role Information"
// @Success 200 {object} models.RoleInfo
// @Router /role/add [post]
func (Role RoleController) Add(c *gin.Context) {
	var roleInfo models.RoleInfo
	err1 := c.ShouldBind(&roleInfo)
	if err1 != nil {
		Role.Error(c, 400, "解析数据失败", err1)
		return
	}

	if roleInfo.Title == "" {
		Role.Error(c, 400, "请输入角色名称", "")
		return
	}

	var role models.RoleInfo
	dao.DB.Where("title=?", roleInfo.Title).Find(&role)
	if role.ID != 0 {
		Role.Error(c, 400, "该角色已存在", "")
		return
	}

	title := strings.Trim(roleInfo.Title, "")
	description := strings.Trim(roleInfo.Description, "")
	createdAt := utils.GetDay()
	roleList := models.RoleInfo{
		Title:       title,
		Description: description,
		Status:      true,
		CreatedAt:   createdAt,
	}

	err3 := dao.DB.Create(&roleList).Error
	if err3 != nil {
		Role.Error(c, 400, "创建角色失败", err3)
		return
	}

	Role.Success(c, 200, "创建角色成功", roleList)
}

// Delete
// @Summary 删除角色
// @Description 删除角色
// @Tags Role
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {string} string "删除成功"
// @Failure 400 {string} string "删除失败"
// @Router /role/delete/{id} [delete]
func (Role RoleController) Delete(c *gin.Context) {
	var roleInfo models.RoleInfo
	id, _ := utils.Int(c.Param("id"))
	dao.DB.Where("id=?", id).Find(&roleInfo)
	if roleInfo.ID == 0 {
		Role.Error(c, 400, "该角色不存在", "")
		return
	}
	err := dao.DB.Delete(&models.RoleInfo{}, id).Error
	if err != nil {
		Role.Error(c, 400, "删除失败", err)
		return
	}
	Role.Success(c, 200, "删除成功", "")
}

// Edit
// @Summary 修改角色
// @Description 修改角色
// @Tags Role
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Param roleInfo body models.RoleInfo true "Role Information"
// @Success 200 {object} models.RoleInfo
// @Failure 400 {string} string "修改失败"
// @Router /role/edit/{id} [put]
func (Role RoleController) Edit(c *gin.Context) {
	var roleInfo models.RoleInfo
	err1 := c.ShouldBind(&roleInfo)
	if err1 != nil {
		Role.Error(c, 400, "解析数据失败", err1)
		return
	}
	if roleInfo.Title == "" {
		Role.Error(c, 400, "请输入角色名称", "")
		return
	}
	id, _ := utils.Int(c.Param("id"))
	roleList, _ := models.GetRoleInfoByID(id)
	if roleList.Title != roleInfo.Title {
		var roleByTitle models.RoleInfo
		dao.DB.Where("title=?", roleInfo.Title).Find(&roleByTitle)
		if roleByTitle.ID != 0 {
			Role.Error(c, 400, "该角色已存在", "")
			return
		}
	}
	role := models.RoleInfo{ID: id}
	createdAt := utils.GetDay()
	err2 := dao.DB.Model(&role).Updates(models.RoleInfo{
		Title:       roleInfo.Title,
		Description: roleInfo.Description,
		Status:      roleInfo.Status,
		CreatedAt:   createdAt,
	}).Error
	if err2 != nil {
		Role.Error(c, 400, "修改角色失败", err2)
	}
	Role.Success(c, 200, "修改角色成功", role)
}
