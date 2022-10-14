package models

import (
	"fmt"
)

type Details struct {
	Id         int    `gorm:"id" json:"id"`           // ID序号
	NewsId     int    `gorm:"news_id" json:"news_id"` // new ID
	Figure     string `gorm:"figure" json:"figure"`   // 列表图
	PList      string `gorm:"p_list" json:"p_list"`
	Alt        string `gorm:"alt" json:"alt"`
	Src        string `gorm:"src" json:"src"` // 新闻内容原URL
	LocalSrc   string `gorm:"local_src" json:"local_src"`
	CreateTime string `gorm:"create_time" json:"create_time"` // 创建时间
	Content    string `gorm:"content" json:"content"`         // 新闻内容
}

func (t *Details) TableName() string {
	return "details"
}

type News struct {
	Id          int    `gorm:"id" json:"id"`           // ID
	Href        string `gorm:"href" json:"href"`       // 新闻源页面URL
	ImgUrl      string `gorm:"img_url" json:"img_url"` // 图片URL
	LoadImg     string `gorm:"load_img" json:"load_img"`
	Title       string `gorm:"title" json:"title"`             // 新闻标题
	Country     string `gorm:"country" json:"country"`         // 国家缩写
	Source      string `gorm:"source" json:"source"`           // 来源的网站名称
	Description string `gorm:"description" json:"description"` // 新闻摘要
	Status      int    `gorm:"status" json:"status"`           // 状态, 0=未处理获取内容, 1=处理内容, 9=处理失败
	CreateTime  string `gorm:"create_time" json:"create_time"` // 获取时间
	Menu        string `gorm:"menu" json:"menu"`               // 新分类
	HrefHash    string `gorm:"href_hash" json:"href_hash"`     // hash值
}

func (t *News) TableName() string {
	return "news"
}

func GetListsNewsModel(PageNum int, PageSize int, maps interface{}) ([]News, error) {

	var lists []News

	err := DB.Self.Where(maps).Offset((PageNum - 1) * PageSize).Limit(PageSize).Order("id desc").Find(&lists).Error
	// err := DB.Self.Where(maps).Offset((PageNum - 1) * PageSize).Limit(PageSize).Order("id desc").Select("news.*, details.content").Joins(condition).Find(&lists).Error

	if err != nil {
		fmt.Println("mysql: ", err)
		return lists, err
	}
	return lists, nil
}

func GetIndexOneModel(id int) (*News, error) {
	u := &News{}
	d := DB.Self.Model(&News{}).Where("id = ?", id).First(&u)
	return u, d.Error
}

func GetNewsOneModel(id int) (*Details, error) {
	u := &Details{}
	d := DB.Self.Where("news_id = ?", id).First(&u)
	return u, d.Error
}

func GetNewesModel(id int) ([]*Details, error) {

	lists := make([]*Details, 0)
	fmt.Println("model", id)
	where := fmt.Sprintf("news_id =  %d", id)
	err := DB.Self.Model(&Details{}).Where(where).Find(&lists).Error
	if err != nil {
		return lists, err
	}

	return lists, nil
}

type blog struct {
}

var Blog *blog

// BlogList 获取 blog 列表
func (b *blog) BlogList(PageNum int, PageSize int, maps interface{}) (blog []News) {
	condition := "left join news on news.id = details.news_id"
	DB.Self.Where(maps).Offset((PageNum - 1) * PageSize).Limit(PageSize).Order("id desc").Select("news.*, details.content").Joins(condition).Find(&blog)
	return
}
