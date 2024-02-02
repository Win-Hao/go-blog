package routers

import (
	"github.com/gin-gonic/gin"
	v "new_demo/api/v1"
	"new_demo/api/v1/admin"
	"new_demo/middleware"
)

func ApiRoutersInit(router *gin.Engine) {
	router.Use(middleware.Cors())

	router.POST("/login", v.LoginServer{}.Login)
	v1 := router.Group("/api/v1", middleware.JWTAuthMiddleware())
	{
		v1.POST("/upload", v.UploadServer{}.Upload)
		//User模块的路由接口
		manager := v1.Group("/manager")
		{
			manager.GET("/", v.ManagerServer{}.GetUsers)
			manager.GET("/:id", v.ManagerServer{}.GetUser)
			manager.POST("/", v.ManagerServer{}.AddUser)
			manager.PUT("/:id", v.ManagerServer{}.EditUser)
			manager.DELETE("/delete/:id", v.ManagerServer{}.DeleteUser)
		}
		//Category模块的路由接口
		category := v1.Group("/category")
		{
			category.GET("/detail", v.CateServer{}.GetCategories)
			category.GET("/oneCate/:cateId", v.CateServer{}.GetCategory)
			category.POST("/add", v.CateServer{}.AddCate)
			category.PUT("/edit/:id", v.CateServer{}.EditCate)
			category.DELETE("/delete/:id", v.CateServer{}.DeleteCate)
			category.GET("/cateArtCount", v.CateServer{}.GetCateArtCount)
		}
		//Article模块的路由接口
		article := v1.Group("/article")
		{
			article.GET("/detail", v.ArtServer{}.GetArticles)
			article.GET("/cateArt", v.ArtServer{}.GetCateArt)
			article.GET("/oneArt/:id", v.ArtServer{}.GetArticle)
			article.POST("/add", v.ArtServer{}.AddArticle)
			article.PUT("/edit/:id", v.ArtServer{}.EditArticle)
			article.DELETE("/delete/:id", v.ArtServer{}.DeleteArticle)
		}

		role := v1.Group("/role")
		{
			role.GET("/", admin.RoleServer{}.GetRole)
			role.POST("/", admin.RoleServer{}.AddRole)
			role.PUT("/", admin.RoleServer{}.EditRole)
			role.DELETE("/:id", admin.RoleServer{}.DeleteRole)
		}
		access := v1.Group("/access")
		{
			access.GET("/module", admin.AccessServer{}.GetTopModule)
		}
	}

}
