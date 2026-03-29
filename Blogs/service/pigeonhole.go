package service

import (
	"blogs/config"
	"blogs/dao"
	models "blogs/models"
	"log"
)

func FindPostPigeonhole() *models.PigenoholeRes {
	//查询所有文章，进行月份整理
	//查询所有文章进行分类整理
	categorys, err := dao.GetAllCategory()
	if err != nil {
		log.Println("查询所有分类失败：", err)
	}
	post, err := dao.GetAllPost()
	lines := make(map[string][]models.Post)
	for _, post := range post {
		at := post.CreateAt
		month := at.Format("2006-01")
		lines[month] = append(lines[month], post)
	}
	return &models.PigenoholeRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Categorys:    categorys,
		Lines:        lines,
	}
}
