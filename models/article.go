package models

import (
	"gorm.io/gorm"
	"new_demo/dao"
	"new_demo/utils/errmsg"
)

type Article struct {
	gorm.Model
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Content     string   `json:"content"`
	CategoryId  int      `json:"category_id"`
	Category    Category `json:"category" gorm:"foreignKey:category_id"`
}

func (Article) TableName() string {
	return "article"
}

func AddArticle(data *Article) int {
	err := dao.DB.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// todo 查询分类下所有文章

func GetCateArt(cid int, pageSize int, pageNum int) ([]Article, int) {
	var article []Article
	err := dao.DB.
		Preload("Category").
		Limit(pageSize).
		Offset((pageNum-1)*pageSize).
		Where("category_id", cid).
		Find(&article).Error
	if err != nil {
		return nil, errmsg.ERROR_CATETITLE_NOT_EXIST
	}
	return article, errmsg.SUCCESS
}

// todo 查询单个文章

func GetArtById(id int) (Article, int) {
	var article Article
	err := dao.DB.
		Preload("Category").
		Where("id=?", id).
		Find(&article).Error
	if err != nil {
		return Article{}, errmsg.ERROR_ART_NOT_EXIST
	}
	return article, errmsg.SUCCESS
}

// GetArticles 查询文章列表
func GetArticles(pageSize int, pageNum int) ([]Article, int) {
	var article []Article
	err := dao.DB.
		Preload("Category").
		Limit(pageSize).
		Offset((pageNum - 1) * pageSize).
		Find(&article).Error

	if err != nil {
		return nil, errmsg.ERROR
	}
	return article, errmsg.SUCCESS
}

// UpdateArticle 编辑文章
func UpdateArticle(id int, data *Article) int {
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["description"] = data.Description
	maps["content"] = data.Content
	maps["category_id"] = data.CategoryId
	err := dao.DB.Model(&Article{}).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteArticle  删除文章
func DeleteArticle(id int) int {
	err := dao.DB.Delete(&Article{}, id).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
