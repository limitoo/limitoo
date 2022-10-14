package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"limitoo/initcfg"
	"limitoo/models"
	"limitoo/routers"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var cfig = pflag.StringP("evn", "c", "", "apiserver config file path.")

func main() {
	pflag.Parse()

	if err := initcfg.ConfigInit(*cfig); err != nil {
		fmt.Println("Init config file error.")
		panic(err)
	}

	// init mysql db
	models.DB.Init()
	defer models.DB.Close()

	fmt.Println("Boot limitoo server...")

	gin.SetMode(viper.GetString("runmode"))
	g := gin.Default()

	g.Use(gin.Recovery())

	g.Use(gzip.Gzip(gzip.DefaultCompression))

	routers.Configire(g)

	fmt.Println("Start to listening the incoming requests on http address:", viper.GetString("addr"))
	fmt.Println(http.ListenAndServe(viper.GetString("addr"), g).Error())
}
