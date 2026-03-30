package dao

import (
	models "blogs/models"
	"log"
)

func GetPostPage(page int, pageSize int) ([]models.Post, error) {
	var posts []models.Post
	offset := (page - 1) * pageSize
	result := DB.Limit(pageSize).Offset(offset).Find(&posts)
	if result.Error != nil {
		log.Println("查询所有post失败：", result.Error)
		return nil, result.Error
	}
	return posts, nil
}

func GetPostCount() int {
	var count int64
	DB.Model(&models.Post{}).Count(&count)
	return int(count)
}

func GetPages(total, page, pageSize int) []int {
	if total == 0 || pageSize == 0 {
		return []int{}
	}
	I := total / pageSize
	if total%pageSize != 0 {
		I++
	}
	var pages []int
	for i := 1; i <= I; i++ {
		pages = append(pages, i)
	}
	return pages
}

func GetPostPageByCID(cid, page, pageSize int) ([]models.Post, error) {
	var posts []models.Post
	offset := (page - 1) * pageSize
	result := DB.Where("category_id = ?", cid).Limit(pageSize).Offset(offset).Find(&posts)
	if result.Error != nil {
		log.Println("查询所有post失败：", result.Error)
		return nil, result.Error
	}
	return posts, nil
}

func GetPostCountByCID(cid int) int {
	var count int64
	DB.Model(&models.Post{}).Where("category_id = ?", cid).Count(&count)
	return int(count)
}

func GetPostByID(id int) (*models.Post, error) {
	var post models.Post
	result := DB.First(&post, id)
	if result.Error != nil {
		log.Println("获取post失败：", result.Error)
		return nil, result.Error
	}
	return &post, nil
}

func SavePost(p *models.Post) error {
	result := DB.Create(p)
	if result.Error != nil {
		log.Println("保存post失败：", result.Error)
		return result.Error
	}
	return nil
}

func UpdatePost(post *models.Post) {
	result := DB.Save(post)
	if result.Error != nil {
		log.Println("更新文章失败：", result.Error)
		return
	}
}

func GetAllPost() ([]models.Post, error) {
	var posts []models.Post
	result := DB.Find(&posts)
	if result.Error != nil {
		log.Println("获取所有post失败：", result.Error)
		return nil, result.Error
	}
	return posts, nil
}

func SearchPost(condition string) ([]models.Post, error) {
	var posts []models.Post
	result := DB.Where("title LIKE ? OR content LIKE ?", "%"+condition+"%", "%"+condition+"%").Find(&posts)
	if result.Error != nil {
		log.Println("获取post失败：", result.Error)
		return nil, result.Error
	}
	return posts, nil
}
