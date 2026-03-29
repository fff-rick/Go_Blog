package api

import (
	"blogs/common"
	"blogs/dao"
	"blogs/models"
	"blogs/service"
	"blogs/utils"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (*Api) UpdateAndSavePost(w http.ResponseWriter, r *http.Request) {
	//验证用户是否合法
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		log.Println("token解析失败：", err)
		common.Error(w, errors.New("登陆已过期"))
		return
	}
	uid := claim.Uid
	method := r.Method
	switch method {
	case http.MethodPost:
		params, err := common.GetResponseJsonParams(r)
		if err != nil {
			log.Println("获取前端json数据失败：", err)
			return
		}
		cIdStr := params["categoryId"]
		cId, _ := strconv.Atoi(cIdStr)
		content := params["content"]
		markdown := params["markdown"]
		slug := params["slug"]
		title := params["title"]
		pT := params["type"]
		postType, _ := strconv.Atoi(pT)
		uid := uid
		viweCount := 0
		post := &models.Post{
			CategoryId: cId,
			Content:    content,
			Markdown:   markdown,
			Slug:       slug,
			Title:      title,
			Type:       postType,
			ViewCount:  viweCount,
			UserId:     uid,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.SavePost(post)
		//返回成功信息给前端
		common.Success(w, post)
	case http.MethodPut:
		params, err := common.GetResponseJsonParams(r)
		if err != nil {
			log.Println("获取前端json数据失败：", err)
			return
		}
		cIdStr := params["categoryId"]
		cId, _ := strconv.Atoi(cIdStr)
		content := params["content"]
		markdown := params["markdown"]
		slug := params["slug"]
		title := params["title"]
		pT := params["type"]
		postType, _ := strconv.Atoi(pT)
		pIDStr := params["pid"]
		pID, _ := strconv.Atoi(pIDStr)
		uid := uid
		viweCount := 0
		post := &models.Post{
			Pid:        pID,
			CategoryId: cId,
			Content:    content,
			Markdown:   markdown,
			Slug:       slug,
			Title:      title,
			Type:       postType,
			ViewCount:  viweCount,
			UserId:     uid,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.UpdatePost(post)
		common.Success(w, post)
	}
}

func (*Api) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pIDStr := strings.TrimPrefix(path, "/api/v1/post/")
	pID, err := strconv.Atoi(pIDStr)
	if err != nil {
		log.Println("文章参数错误")
		common.Error(w, errors.New("路径参数错误"))
		return
	}
	post, err := dao.GetPostByID(pID)
	if err != nil {
		log.Println("获取文章详情错误：", err)
		common.Error(w, errors.New("获取文章详情错误"))
		return
	}
	common.Success(w, post)
}

func (*Api) SearchPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("解析表单数据出错：", err)
		return
	}
	condition := r.Form.Get("val")
	searchRes := service.SearchPost(condition)
	common.Success(w, searchRes)
}
