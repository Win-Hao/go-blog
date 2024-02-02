package models

import (
	"new_demo/dao"
	"new_demo/utils/errmsg"
)

type Category struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CategoriesWithArticles struct {
	CategoryId    int    `json:"category_id"`
	CategoryTitle string `json:"category_title"`
	ArticlesCount int    `json:"articles_count"`
}

func (Category) TableName() string {
	return "category"
}

func CheckCategory(title string) int {
	var category Category
	dao.DB.Where("title=?", title).Find(&category)
	if category.ID != 0 {
		return errmsg.ERROR_CATETITLE_USED
	}
	return errmsg.SUCCESS
}

func AddCategory(data *Category) int {

	err := dao.DB.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
func GetCategoryById(id int) Category {
	var cate Category
	dao.DB.Where("id=?", id).Find(&cate)
	return cate
}
func GetCategories(pageSize int, pageNum int) ([]Category, int) {
	var cate []Category
	err := dao.DB.
		Limit(pageSize).
		Offset((pageNum - 1) * pageSize).
		Find(&cate).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return cate, errmsg.SUCCESS
}

// UpdateCategory 编辑分类
func UpdateCategory(id int, data *Category) int {
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["description"] = data.Description
	err := dao.DB.Model(&Category{}).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteCategory  删除用户
func DeleteCategory(id int) int {
	err := dao.DB.Delete(&Category{}, id).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetCategoriesWithArticlesCount 获取分类下文章数
func GetCategoriesWithArticlesCount() ([]CategoriesWithArticles, error, int) {
	var categoryCounts []CategoriesWithArticles
	err := dao.DB.Model(&Article{}).
		Select("category.id as category_id, category.title as category_title , count(article.id) as articles_count").
		Joins("JOIN category ON category.id = article.category_id").
		Group("category.id , category.title").Scan(&categoryCounts).Error
	if err != nil {
		return nil, err, errmsg.ERROR
	}
	return categoryCounts, nil, errmsg.SUCCESS
}
