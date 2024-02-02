package models

import (
	"new_demo/dao"
	"new_demo/utils/errmsg"
	"time"
)

type Manager struct {
	ID        int ` json:"id"`
	CreatedAt int `json:"created_at"`
	UpdatedAt time.Time
	Username  string `json:"username"`
	Password  string `json:"password"`
	Status    bool   `json:"status"`
	Email     string `json:"email"`
	Mobile    int    `json:"mobile"`
	RoleId    int    `json:"role_id"`
	IsSuper   int    `json:"is_super"`
	Role      Role   `json:"role" ,gorm:"foreignKey:role_id"`
}

func (Manager) TableName() string {
	return "manager"
}

func CheckUser(username string) int {
	var manager Manager
	dao.DB.Where("username=?", username).Find(&manager)
	if manager.ID != 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

func AddUser(data *Manager) int {
	err := dao.DB.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
func GetUserById(id int) (Manager, int) {
	var manager Manager
	err := dao.DB.Preload("Role").Where("id=?", id).Find(&manager).Error
	if err != nil {
		return Manager{}, errmsg.ERROR
	}
	return manager, errmsg.SUCCESS
}
func GetUsers(pageSize int, pageNum int) ([]Manager, int) {
	var managers []Manager
	err := dao.DB.
		Preload("Role").
		Limit(pageSize).
		Offset((pageNum - 1) * pageSize).
		Find(&managers).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return managers, errmsg.SUCCESS
}

// UpdateUser 编辑用户
func UpdateUser(id int, data *Manager) int {
	var maps = make(map[string]interface{})
	maps["status"] = data.Status
	maps["email"] = data.Email
	maps["mobile"] = data.Mobile
	maps["role_id"] = data.RoleId
	maps["is_super"] = data.IsSuper
	err := dao.DB.Model(&Manager{}).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	err := dao.DB.Delete(&Manager{}, id).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//登录验证

func CheckLogin(username string, password string) int {
	var manager Manager
	dao.DB.Where("username=?", username).Find(&manager)
	if manager.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if password != manager.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	return errmsg.SUCCESS
}
