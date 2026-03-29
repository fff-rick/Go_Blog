package views

import (
	"blogs/common"
	"blogs/service"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail

	// 获取访问的路径参数
	path := r.URL.Path
	// /p/3.html
	path = strings.TrimPrefix(path, "/p/")
	pIDStr := strings.TrimSuffix(path, ".html")
	pID, err := strconv.Atoi(pIDStr)
	if err != nil {
		log.Println("文章参数错误")
		detail.WriteErr(w, errors.New("路径参数错误"))
		return
	}
	postRes, err := service.GetPostDetail(pID)
	if err != nil {
		log.Println("获取文章详情错误：", err)
		detail.WriteErr(w, errors.New("获取文章详情错误"))
		return
	}
	detail.WriteData(w, postRes)
}
