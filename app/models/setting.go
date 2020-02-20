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
type Setting struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);" json:"name"`
	Value       string `gorm:"type:varchar(255);" json:"value"`
	Status      string `gorm:"type:varchar(20);" json:"status"`
	SettingType string `gorm:"type:varchar(20);" json:"setting_type"`
	Slug        string `gorm:"type:varchar(50);" json:"slug"`
}

/*
* event run after add Setting
 */
func (u *Setting) AfterCreate(scope *gorm.Scope) (err error) {
	IncreaseOnCreate("settings")
	return
}

/*
* event run after delete Faq
 */
func (u *Setting) AfterDelete(tx *gorm.DB) (err error) {
	DecreaseOnDelete(u.Status, "settings")
	return
}

/**
* update status
 */
func (u *Setting) BeforeUpdate() (err error) {
	var setting Setting
	config.DB.First(&setting , u.ID)
	if setting.Status != u.Status{
		DecreaseRow(setting.Status, "settings")
		IncreaseRow(u.Status , "settings")
	}
	return
}

/**
* migration function must be the file name concat with Migrate
* key word Example : user will be UserMigrate
 */
func SettingMigrate() {
	config.DB.AutoMigrate(&Setting{})
}

/**
* you can update these column only
 */
func SettingFillAbleColumn() []string {
	return []string{"name", "value", "setting_type", "slug"}
}
