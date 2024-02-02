package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"new_demo/models"
	"new_demo/utils"
	"new_demo/utils/errmsg"
)

type ArtServer struct {
}

// CateExist 查询用户是否存在
func (a ArtServer) CateExist(c *gin.Context) {

}

// AddArticle  添加文章
func (a ArtServer) AddArticle(c *gin.Context) {
	var article models.Article
	_ = c.ShouldBind(&article)
	code := models.AddArticle(&article)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    article,
		"message": errmsg.GetCode(code),
	})
}

// todo 查询分类下所有文章

func (a ArtServer) GetCateArt(c *gin.Context) {
	pageSize, _ := utils.Int(c.Query("pageSize"))
	pageNum, _ := utils.Int(c.Query("pageNum"))
	cid, _ := utils.Int(c.Query("cid"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = 2
	}

	data, code := models.GetCateArt(cid, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"message": errmsg.GetCode(code),
	})
}

// todo 查询单个文章信息

func (a ArtServer) GetArticle(c *gin.Context) {
	id, _ := utils.Int(c.Param("id"))
	article, code := models.GetArtById(id)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    article,
		"message": errmsg.GetCode(code),
	})
}

// GetArticles 查询分类列表
func (a ArtServer) GetArticles(c *gin.Context) {
	pageSize, _ := utils.Int(c.Query("pageSize"))
	pageNum, _ := utils.Int(c.Query("pageNum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = 2
	}
	articles, code := models.GetArticles(pageSize, pageNum)
	fmt.Println(len(articles))
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    articles,
		"message": errmsg.GetCode(code),
	})
}

// EditArticle 编辑文章
func (a ArtServer) EditArticle(c *gin.Context) {
	var article models.Article
	id, _ := utils.Int(c.Param("id"))
	_ = c.ShouldBind(&article)
	status := models.UpdateArticle(id, &article)
	c.JSON(http.StatusOK, gin.H{
		"code":    status,
		"data":    article,
		"message": errmsg.GetCode(status),
	})
}

// DeleteArticle 删除文章
func (a ArtServer) DeleteArticle(c *gin.Context) {
	id, _ := utils.Int(c.Param("id"))
	code := models.DeleteArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetCode(code),
	})
}
