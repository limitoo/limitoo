package utils

import (
	"log"

	"gopkg.in/ini.v1"
)

var (
	Config *ini.File

	// PageSize RunMode string
	PageSize int
	Site     *SiteMap
)

type SiteMap struct {
	Domain      string `ini:"domain"`
	Title       string `ini:"title"`
	Slogon      string `ini:"slogon"`
	Keyword     string `ini:"keyword"`
	Description string `ini:"description"`
	Email       string `ini:"email"`
	Contact     string `ini:"contact"`
	Company     string `ini:"company"`
	Phone       string `ini:"phone"`
	Icp         string `ini:"icp"`
	Address     string `ini:"address"`
	Tpl         string `ini:"tpl"`
}

func init() {
	// SetApp()
	// SetMeta()
	// SetSite()
}

// LoadInI 载入配置
func LoadInI() {
	var err error
	Config, err = ini.Load("config/env.ini")
	// Config, err = ini.Load("env.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'config/env.ini': %v", err)
	}
}

func SetApp() {
	sec := Config.Section("app")
	PageSize = sec.Key("page_size").MustInt(10)
}

func SetMeta() {
	sec := Config.Section("site")
	_ = sec.Key("keyword").MustString("")
}

func SetSite() {
	Site = &SiteMap{}
	_ = Config.Section("site").MapTo(Site)
}
