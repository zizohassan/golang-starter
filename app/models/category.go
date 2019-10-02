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
type Category struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50);" json:"name"`
	Status   int    `gorm:"type:tinyint(1);" json:"status"`
}

/**
* migration function must be the file name concat with Migrate
* key word Example : user will be UserMigrate
 */
func (s *MigrationTables) CategoryMigrate() {
	config.DB.AutoMigrate(&Category{})
}
