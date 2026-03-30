package router

import (
	"blogs/api"
	"blogs/views"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 适配器函数，将 http.HandlerFunc 转换为 gin 处理函数
func adaptHandler(h http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(c.Writer, c.Request)
	}
}

func Router() *gin.Engine {
	r := gin.Default()

	// 静态文件服务
	r.Static("/resource", "public/resource")

	// HTML 路由
	r.GET("/", adaptHandler(views.HTML.Index))
	r.GET("/c/*path", adaptHandler(views.HTML.Category))
	r.GET("/login/*path", adaptHandler(views.HTML.Login))
	r.GET("/p/*path", adaptHandler(views.HTML.Detail))
	r.GET("/pigeonhole", adaptHandler(views.HTML.Pigeonhole))
	r.GET("/writing/*path", adaptHandler(views.HTML.Writing))

	// API 路由
	r.GET("/api/v1/post/search", adaptHandler(api.API.SearchPost))
	r.POST("/api/v1/login", adaptHandler(api.API.Login))
	r.POST("/api/v1/post", adaptHandler(api.API.UpdateAndSavePost))
	r.GET("/api/v1/post/", adaptHandler(api.API.GetPost))
	r.GET("/api/v1/qiniu/token", adaptHandler(api.API.QiniuToken))

	return r
}
