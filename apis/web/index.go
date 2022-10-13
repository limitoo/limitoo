package web

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

func IndexPage(c *gin.Context) {

	c.HTML(200, "/web/index.html", pongo2.Context{
		"title": "Limitoo news",
	})
}
