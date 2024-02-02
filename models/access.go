package models

import (
	"new_demo/dao"
	"new_demo/utils/errmsg"
)

type Access struct {
	ID          int    `json:"id"`
	ModuleName  string `json:"module_name"`
	Description string `json:"description"`
	ModuleId    int    `json:"module_id"`
	ActionName  string `json:"action_name"`
	Type        int    `json:"type"`
	CreatedAt   int    `json:"created_at"`
	Url         string `json:"url"`
	Status      int    `json:"status"`
	Sort        int    `json:"sort"`
}

func (Access) TableName() string {
	return "access"
}

type Module struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func AddAccess(data *Access) int {
	err := dao.DB.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
func GetTopModule() ([]Module, int) {
	var module []Module
	err := dao.DB.Table("access").Select("id", "module_name").
		Where("module_id=?", 0).Scan(&module).Error
	if err != nil {
		return []Module{}, errmsg.ERROR
	}
	return module, errmsg.SUCCESS
}
