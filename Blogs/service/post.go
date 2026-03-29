package service

import (
	"blogs/config"
	"blogs/dao"
	"blogs/models"
	"html/template"
	"log"
)

func GetPostDetail(id int) (*models.PostRes, error) {
	post, err := dao.GetPostByID(id)
	if err != nil {
		log.Println("获取post失败：", err)
		return nil, err
	}

	pr := &models.PostRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Article: models.PostMore{
			Pid:          post.Pid,
			Title:        post.Title,
			Slug:         post.Slug,
			Content:      template.HTML(post.Content),
			CategoryId:   post.CategoryId,
			CategoryName: dao.GetCategoryNameByID(post.CategoryId),
			UserId:       post.UserId,
			UserName:     dao.GetUserNameByID(post.UserId),
			ViewCount:    post.ViewCount,
			Type:         post.Type,
			CreateAt:     models.DateDay(post.CreateAt),
			UpdateAt:     models.DateDay(post.UpdateAt),
		},
	}
	return pr, nil

}

func Writing() *models.WritingRes {
	wr := &models.WritingRes{}
	wr.Title = config.Cfg.Viewer.Title
	wr.CdnURL = config.Cfg.System.CdnURL
	wr.QiniuDomain = config.Cfg.System.QiniuDomain
	Categorys, err := dao.GetAllCategory()
	if err != nil {
		log.Println("取值所有category失败：", err)
		return nil
	}
	wr.Categorys = Categorys
	return wr
}

func SavePost(post *models.Post) {
	dao.SavePost(post)
}

func UpdatePost(post *models.Post) {
	dao.UpdatePost(post)
}

func SearchPost(condition string) []models.SearchResp {
	posts, err := dao.SearchPost(condition)
	if err != nil {
		log.Println("搜索失败：", err)
		return nil
	}
	var sr []models.SearchResp
	for _, post := range posts {
		sr = append(sr, models.SearchResp{
			Pid:   post.Pid,
			Title: post.Title,
		})
	}
	return sr
}
