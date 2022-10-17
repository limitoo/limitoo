package web

import (
	"limitoo/models"
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

func IndexPage(c *gin.Context) {
	data := make(map[string]interface{})
	totals := 100

	lists, err := models.GetListsNewsModel(1, totals, "status=1")
	if err != nil {
		c.HTML(500, "/web/404.html", pongo2.Context{
			"error": err,
		})
		return
	}

	data["title"] = "Limitoo news"
	data["lists"] = lists
	data["totals"] = totals

	c.HTML(http.StatusOK, "/web/index/content.html", pongo2.Context{
		"title": "Limitoo news",
		"lists": lists,
	})
}
