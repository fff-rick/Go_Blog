package models

import (
	"html/template"
	"io"
	"log"
	"time"
)

type TemplateBlog struct {
	*template.Template
}

type HtmlTemplate struct {
	Index      TemplateBlog
	Category   TemplateBlog
	Custom     TemplateBlog
	Detail     TemplateBlog
	Login      TemplateBlog
	Pigeonhole TemplateBlog
	Writing    TemplateBlog
}

func InitTemplate(templateDir string) (*HtmlTemplate, error) {
	tp, err := readTemplate([]string{"index",
		"category",
		"custom",
		"detail",
		"login",
		"pigeonhole",
		"writing"},
		templateDir,
	)
	if err != nil {
		log.Println("解析模板出错：", err)
		return nil, err
	}
	var htmltemplate HtmlTemplate
	htmltemplate.Index = tp[0]
	htmltemplate.Category = tp[1]
	htmltemplate.Custom = tp[2]
	htmltemplate.Detail = tp[3]
	htmltemplate.Login = tp[4]
	htmltemplate.Pigeonhole = tp[5]
	htmltemplate.Writing = tp[6]
	return &htmltemplate, nil

}

func readTemplate(teplates []string, templateDir string) ([]TemplateBlog, error) {
	var tbs []TemplateBlog
	for _, view := range teplates {
		viewName := view + ".html"
		// 名为 `"*.html"` 的模板实例
		t := template.New(view + ".html")
		//解析所有模板文件
		home := templateDir + "home.html"
		header := templateDir + "layout/header.html"
		footer := templateDir + "layout/footer.html"
		personal := templateDir + "layout/personal.html"
		post := templateDir + "layout/post-list.html"
		pagination := templateDir + "layout/pagination.html"
		//映射函数
		t.Funcs(template.FuncMap{
			"isODD":       IsODD,
			"getNextName": GetNextName,
			"date":        Date,
			"dateDay":     DateDay,
		})

		t, err := t.ParseFiles(templateDir+viewName,
			home,
			header,
			footer,
			personal, post,
			pagination)
		if err != nil {
			log.Println("模板解析错误：", err)
			return nil, err
		}
		var tb TemplateBlog
		tb.Template = t
		tbs = append(tbs, tb)
	}
	return tbs, nil
}

func (t *TemplateBlog) WriteData(w io.Writer, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		_, err = w.Write([]byte("写数据出错"))
		if err != nil {
			log.Println("写回数据出错：", err)
			return
		}
	}
}

func (t *TemplateBlog) WriteErr(w io.Writer, err error) {
	if err != nil {
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			log.Println("写回错误出错：", err)
			return
		}
	}
}

func IsODD(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}
func GetNextName(str []string, index int) string {
	return str[index+1]
}
func Date(layout string) string {
	return time.Now().Format(layout)
}
func DateDay(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}
