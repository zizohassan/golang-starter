package models

import (
	"github.com/jinzhu/gorm"
	"golang-starter/config"
)

/***
* model struct here we will build the main
* struct that connect to database
 */
type Action struct {
	gorm.Model
	Noun       string `gorm:"type:varchar(50)" json:"noun"`
	Verb       string `gorm:"type:varchar(50)" json:"verb"`
	Slug       string `gorm:"type:varchar(50)" json:"slug"`
	ModuleName string `gorm:"type:varchar(50)" json:"module_name"`
	Count      int    `gorm:"type:int" json:"count"`
}


/**
* migration function must be the file name concat with Migrate
* key word Example : user will be UserMigrate
 */
func ActionMigrate() {
	config.DB.AutoMigrate(&Action{})
}