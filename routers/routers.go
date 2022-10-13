package routers

import (
	"limitoo/apis/web"
	"limitoo/utils"

	"github.com/gin-gonic/gin"
)

func Configire(g *gin.Engine) *gin.Engine {
	// 模板引擎
	g.HTMLRender = utils.NewRender("website")

	// set public dir
	g.Static("/public", "./website/public/files")
	g.Static("/uploads", "./website/public/uploads")
	g.Static("/assets", "./website/public/assets")

	g.StaticFile("/favicon.ico", "./website/public/favicon.ico")
	g.StaticFile("/admin/mini", "./website/public/api/init.json")

	// website router

	weburl := g.Group("")
	{
		weburl.GET("/", web.IndexPage)
	}

	// admin router

	return g

}
