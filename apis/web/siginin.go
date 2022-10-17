package web

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

func SinginPage(c *gin.Context) {

	c.HTML(200, "/web/signin.html", pongo2.Context{
		"title": "Limitoo news",
	})
}
