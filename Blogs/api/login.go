package api

import (
	"blogs/common"
	"blogs/service"
	"net/http"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	result, _ := common.GetResponseJsonParams(r)
	username := result["username"]
	password := result["password"]
	//fmt.Println("get: ", password)
	lr, err := service.Login(username, password)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, lr)
}
