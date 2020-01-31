package models

import (
	"github.com/jinzhu/gorm"
	"golang-starter/config"
)

/***
* model struct here we will build the main
* struct that connect to database
 */
type Answer struct {
	gorm.Model
	Text  string `gorm:"type:text" json:"text"`
	FaqId int    `gorm:"type:int" json:"faq_id"`
}

/**
* migration function must be the file name concat with Migrate
* key word Example : user will be UserMigrate
 */
func AnswerMigrate() {
	config.DB.AutoMigrate(&Answer{})
}

/**
* you can update these column only
 */
func AnswerFillAbleColumn() []string {
	return []string{"text"}
}
