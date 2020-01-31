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
type Page struct {
	gorm.Model
	Name         string              `gorm:"type:varchar(50);" json:"name"`
	Status       int                 `gorm:"type:tinyint(1);" json:"status"`
	Image        []string            `gorm:"-" json:"image"`
	Translation  []map[string]string `gorm:"-" json:"translation"`
	Translations []Translation       `gorm:"association_autoupdate:false;association_autocreate:false" json:"translations"`
	Images       []PageImage         `gorm:"association_autoupdate:false;association_autocreate:false" json:"images"`
}

/**
* migration function must be the file name concat with Migrate
* key word Example : user will be UserMigrate
 */
func  PageMigrate() {
	config.DB.AutoMigrate(&Page{})
}

/**
* you can update these column only
 */
func PageFillAbleColumn() []string {
	return []string{"name", "status"}
}


/**
* active Page only
 */
func ActivePage(db *gorm.DB) *gorm.DB {
	return db.Where("status = 2")
}
