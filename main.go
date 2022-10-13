package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"limitoo/initcfg"
	"limitoo/routers"

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

	fmt.Println("Boot limitoo server...")

	gin.SetMode(viper.GetString("runmode"))
	g := gin.Default()

	routers.Configire(g)

	fmt.Println("Start to listening the incoming requests on http address:", viper.GetString("addr"))
	fmt.Println(http.ListenAndServe(viper.GetString("addr"), g).Error())
}
