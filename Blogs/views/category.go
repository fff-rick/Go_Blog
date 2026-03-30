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

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	categoryTemplate := common.Template.Category
	cIDStr := strings.TrimPrefix(path, "/c/")
	// 提取数字部分作为分类 ID
	cIDStr = strings.Split(cIDStr, ".")[0]
	cID, err := strconv.Atoi(cIDStr)
	if err != nil {
		log.Println("分类参数错误")
		categoryTemplate.WriteErr(w, errors.New("路径参数错误"))
		return
	}

	if err := r.ParseForm(); err != nil {
		categoryTemplate.WriteErr(w, errors.New("解析表单数据错误"))
		return
	}

	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	pageSize := 5
	categoryResponse, err := service.GetPostsByCategoryId(cID, page, pageSize)
	if err != nil {
		categoryTemplate.WriteErr(w, err)
		return
	}
	categoryTemplate.WriteData(w, categoryResponse)

}
