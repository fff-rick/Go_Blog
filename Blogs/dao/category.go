package dao

import (
	models "blogs/models"
	"log"
)

func GetAllCategory() ([]models.Category, error) {
	var categorys []models.Category
	result := DB.Find(&categorys)
	if result.Error != nil {
		log.Println("查询所有category失败：", result.Error)
		return nil, result.Error
	}
	return categorys, nil
}

func GetCategoryNameByID(id int) string {
	var category models.Category
	result := DB.Where("cid = ?", id).First(&category)
	if result.Error != nil {
		log.Println("获取分类名称失败：", result.Error)
		return ""
	}
	return category.Name
}
