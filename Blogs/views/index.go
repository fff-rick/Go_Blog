package views

import (
	"blogs/common"
	"blogs/service"
	"errors"
	"log"
	"net/http"
	"strconv"
)

// 从数据库传入数据
func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	// 解析表单数据
	if err := r.ParseForm(); err != nil {
		index.WriteErr(w, errors.New("解析表单数据错误"))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	pageSize := 5

	hr, err := service.GetAllIndexInfo(page, pageSize)
	if err != nil {
		log.Println("获取首页数据错误：", err)
		index.WriteErr(w, errors.New("获取首页数据错误"))
	}
	index.WriteData(w, hr)
}
