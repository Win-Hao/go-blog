package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"new_demo/middleware"
	"new_demo/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.Logger())
	ApiRoutersInit(router)
	//dao.DB.AutoMigrate(&models.Access{})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":" + utils.HttpPort)
}
