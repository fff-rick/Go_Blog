package service

import (
	"blogs/config"
	"blogs/dao"
	models "blogs/models"
	"html/template"
	"log"
)

func GetAllIndexInfo(page, pageSize int) (*models.HomeResponse, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		log.Println("取值所有category失败：", err)
		return nil, err
	}
	posts, err := dao.GetPostPage(page, pageSize)
	var postsMore []models.PostMore
	for _, post := range posts {
		content := []rune(post.Content)
		if len(content) > 50 {
			content = content[:50]
		}
		categoryName := dao.GetCategoryNameByID(post.Pid)
		userName := dao.GetUserNameByID(post.UserId)
		pm := models.PostMore{
			Pid:          post.Pid,
			Title:        post.Title,
			Slug:         post.Slug,
			Content:      template.HTML(content),
			CategoryId:   post.CategoryId,
			CategoryName: categoryName,
			UserId:       post.UserId,
			UserName:     userName,
			CreateAt:     models.DateDay(post.CreateAt),
			UpdateAt:     models.DateDay(post.UpdateAt),
		}
		postsMore = append(postsMore, pm)
	}

	Total := dao.GetPostCount()
	pages := dao.GetPages(Total, page, pageSize)
	var hr = &models.HomeResponse{
		Viewer:    config.Cfg.Viewer,
		Categorys: categorys,
		Posts:     postsMore,
		Total:     Total,
		Page:      page,
		Pages:     pages,
		PageEnd:   page == Total,
	}
	return hr, nil
}
