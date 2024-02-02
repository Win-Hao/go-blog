package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"new_demo/models"
	"new_demo/utils"
	"new_demo/utils/errmsg"
)

type CateServer struct {
}

// CateExist 查询用户是否存在
func (ca CateServer) CateExist(c *gin.Context) {

}

// AddCate  添加分类
func (ca CateServer) AddCate(c *gin.Context) {
	var cate models.Category
	_ = c.ShouldBind(&cate)
	status := models.CheckCategory(cate.Title)
	if status == errmsg.ERROR_CATETITLE_USED {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR_CATETITLE_USED,
			"message": errmsg.GetCode(errmsg.ERROR_CATETITLE_USED),
		})
		return
	}
	code := models.AddCategory(&cate)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    cate,
		"message": errmsg.GetCode(code),
	})

}

//查询单个用户

// GetCategories 查询分类列表
func (ca CateServer) GetCategories(c *gin.Context) {
	pageSize, _ := utils.Int(c.Query("pageSize"))
	pageNum, _ := utils.Int(c.Query("pageNum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = 2
	}
	categories, code := models.GetCategories(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    categories,
		"message": errmsg.GetCode(code),
	})
}
func (ca CateServer) GetCategory(c *gin.Context) {
	cateId, _ := utils.Int(c.Param("cateId"))
	category := models.GetCategoryById(cateId)
	c.JSON(http.StatusOK, gin.H{
		"data":    category,
		"message": "获取成功",
	})
}

// EditCate 编辑分类
func (ca CateServer) EditCate(c *gin.Context) {
	var cate models.Category
	id, _ := utils.Int(c.Param("id"))
	_ = c.ShouldBind(&cate)
	cateInfo := models.GetCategoryById(id)
	if cate.Title != cateInfo.Title {
		code := models.CheckCategory(cate.Title)
		if code == errmsg.ERROR_CATETITLE_USED {
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetCode(code),
			})
			return
		}
	}
	status := models.UpdateCategory(id, &cate)
	c.JSON(http.StatusOK, gin.H{
		"code":    status,
		"data":    cate,
		"message": errmsg.GetCode(status),
	})
}

// DeleteCate 删除分类
func (ca CateServer) DeleteCate(c *gin.Context) {
	id, _ := utils.Int(c.Param("id"))
	code := models.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetCode(code),
	})
}

// GetCateArtCount 查询分类下的文章数
func (ca CateServer) GetCateArtCount(c *gin.Context) {
	CateArtCount, _, code := models.GetCategoriesWithArticlesCount()
	if code != errmsg.ERROR {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": errmsg.GetCode(code),
			"data":    CateArtCount,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetCode(code),
	})

}
