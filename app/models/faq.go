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
type Faq struct {
	gorm.Model
	Question string   `gorm:"type:varchar(255);" json:"question"`
	Status   int      `gorm:"type:tinyint(1);" json:"status"`
	Answer   []string `gorm:"-" json:"answer"`
	Answers  []Answer `gorm:"association_autoupdate:false;association_autocreate:false" json:"answers"`
}

/**
* migration function must be the file name concat with Migrate
* key word Example : user will be UserMigrate
 */
func  FaqMigrate() {
	config.DB.AutoMigrate(&Faq{})
}

/**
* you can update these column only
 */
func FaqFillAbleColumn() []string {
	return []string{"question", "status"}
}

/**
* active questions only
*/
func ActiveFaq(db *gorm.DB) *gorm.DB {
	return db.Where("status = 2")
}
