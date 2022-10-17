package web

import (
	"limitoo/models"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

func ContentPage(c *gin.Context) {
	newsId := c.Param("id")
	id, err := strconv.Atoi(newsId)
	if err != nil {
		id = 1
	}

	info, err := models.GetIndexOneModel(id)
	if err != nil {
		c.HTML(500, "/web/404.html", pongo2.Context{
			"error": err,
		})
		return
	}

	content, err := models.GetNewsOneModel(id)
	if err != nil {
		c.HTML(500, "/web/404.html", pongo2.Context{
			"error": err,
		})
		return
	}
	// err := json.Unmarshal([]byte(content.Content), &x)
	// if err != nil {
	// 	panic(err)
	// }
	tmp := StringToBytes(content.Content)
	// str := strings.Split(tmp, "\", \"")
	// fmt.Println("1111111111", str)
	// output := []string{tmp}
	// fmt.Println("sss", reflect.TypeOf(str))
	// fmt.Println("sss", reflect.TypeOf(output))
	// fmt.Println("content: %v", output)

	read := ReadNum()
	ReadNum()

	c.HTML(http.StatusOK, "/web/posts/index.html", pongo2.Context{
		"title":   "Limitoo news",
		"info":    info,
		"content": tmp,
		"read":    read,
	})
}

func ReadNum() int64 {
	rand.Seed(2)

	return rand.Int63n(10000)
}
func StringToBytes(str string) []string {
	// data := make(map[string]interface{})

	abc := str[1:]
	tmp := abc[:len(abc)-1]

	nstr := strings.Split(tmp, "\", \"")

	// json.Unmarshal([]byte(tmp), &data)
	// newStr, err := json.Marshal(data)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("===", string(newStr))

	return nstr
}
