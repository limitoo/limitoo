package models

import (
	"limitoo/initcfg"

	"github.com/jinzhu/gorm"
)

type Database struct {
	Self *gorm.DB
}

var DB *Database

func (db *Database) Init() {
	DB = &Database{
		Self: initcfg.GetSelfDB(),
	}

}

func (db *Database) Close() {
	DB.Self.Close()
}
