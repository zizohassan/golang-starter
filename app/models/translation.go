package models

import (
	"github.com/jinzhu/gorm"
	"golang-starter/config"
)

/***
* model struct here we will build the main
* struct that connect to database
* status 1 active 2 is not active
 */
type Translation struct {
	gorm.Model
	Value  string `gorm:"type:varchar(255);" json:"value"`
	Slug   string `gorm:"type:varchar(50);" json:"slug"`
	Lang   string `gorm:"type:varchar(10);" json:"lang"`
	PageId int    `gorm:"type:varchar(30);" json:"page_id"`
}

/**
* migration function must be the file name concat with Migrate
* key word Example : user will be UserMigrate
 */
func (s *MigrationTables) TranslationMigrate() {
	config.DB.AutoMigrate(&Translation{})
}

/**
* you can update these column only
 */
func TranslationFillAbleColumn() []string {
	return []string{"value", "slug", "lang", "page_id"}
}
