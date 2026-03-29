package common

import (
	"blogs/config"
	models "blogs/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
)

var Template *models.HtmlTemplate

func LoadTemplate() {
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			panic(err)
		}
		w.Done()
	}()
	w.Wait()
}

func Success(w http.ResponseWriter, data interface{}) {
	var result models.Result
	result = models.Result{
		Code:  200,
		Error: "",
		Data:  data,
	}
	resultJson, err := json.Marshal(result)
	if err != nil {
		log.Println("json格式转化失败: ", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resultJson)
}

func Error(w http.ResponseWriter, err error) {
	var result models.Result
	result = models.Result{
		Code:  -1,
		Error: err.Error(),
	}
	resultJson, err := json.Marshal(result)
	if err != nil {
		log.Println("json格式转化失败: ", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resultJson)
}

func GetResponseJsonParams(r *http.Request) (map[string]string, error) {
	var params map[string]string

	// 读取body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("读取请求体失败: ", err)
		return nil, err
	}

	// 尝试解析JSON
	err = json.Unmarshal(body, &params)
	if err != nil {
		log.Printf("JSON解析失败: %v", err)
		return nil, err
	}
	return params, nil
}
