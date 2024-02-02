package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"new_demo/utils"
)

var DB *gorm.DB
var err error

func init() {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		utils.Username,
		utils.Password,
		utils.Ip,
		utils.Port,
		utils.Database,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("连接数据库失败，请重试:%s", err))
	}
}
