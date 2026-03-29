package api

import (
	"blogs/common"
	"blogs/config"
	"net/http"

	"github.com/qiniu/api.v7/v7/auth"
	"github.com/qiniu/api.v7/v7/storage"
)

var API = &Api{}

type Api struct {
}

func (*Api) QiniuToken(w http.ResponseWriter, r *http.Request) {
	mac := auth.New(config.Cfg.System.QiniuAccessKey, config.Cfg.System.QiniuSecretKey)
	putPolicy := storage.PutPolicy{
		Scope: config.Cfg.System.QiniuBucket,
	}
	token := putPolicy.UploadToken(mac)
	common.Success(w, token)
}
