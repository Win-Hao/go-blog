package models

import (
	"new_demo/dao"
	"new_demo/utils/errmsg"
	"time"
)

type Role struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

func (Role) TableName() string {
	return "role"
}

func CheckRole(title string) int {
	role := Role{}
	err := dao.DB.Where("title=?", title).Find(&role).Error
	if err != nil {
		return errmsg.ERROR
	}
	if role.ID != 0 {
		return errmsg.ERROR_ROLE_USED
	}
	return errmsg.SUCCESS
}

func GetRoles(pageSize int, pageNum int) ([]Role, int) {
	var role []Role
	err := dao.DB.Limit(pageSize).
		Offset((pageNum - 1) * pageSize).
		Find(&role).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return role, errmsg.SUCCESS
}

func CheckRoleById(id int) (Role, int) {
	role := Role{}
	dao.DB.Where("id=?", id).Find(&role)
	if role.ID == 0 {
		return role, errmsg.ERROR_ROLE_NOT_EXIST
	}
	return role, errmsg.SUCCESS
}

func AddRole(data *Role) int {
	err := dao.DB.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func EditRole(data *Role, id int) int {
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["description"] = data.Description
	maps["status"] = data.Status
	err := dao.DB.Model(&Role{}).Where("id=?", id).Updates(&maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteRole(id int) int {
	err := dao.DB.Delete(&Role{}, id).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
