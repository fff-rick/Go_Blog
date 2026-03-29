package dao

import (
	models "blogs/models"
	"log"
)

func GetAllCategory() ([]models.Category, error) {
	rows, err := DB.Query("SELECT * FROM category")
	if err != nil {
		log.Println("查询所有category失败：", err)
		return nil, err
	}
	var categorys []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.Cid,
			&category.Name,
			&category.CreateAt,
			&category.UpdateAt)
		if err != nil {
			log.Println("取值category失败：", err)
			return nil, err
		}
		categorys = append(categorys, category)

	}
	return categorys, nil
}

func GetCategoryNameByID(id int) string {
	sql := `select name from category where cid = ?`
	row := DB.QueryRow(sql, id)
	var name string
	row.Scan(&name)
	return name
}
