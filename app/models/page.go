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
	Status       string              `gorm:"type:varchar(20);" json:"status"`
	Image        []string            `gorm:"-" json:"image"`
	Translation  []map[string]string `gorm:"-" json:"translation"`
	Translations []Translation       `gorm:"association_autoupdate:false;association_autocreate:false" json:"translations"`
	Images       []PageImage         `gorm:"association_autoupdate:false;association_autocreate:false" json:"images"`
}

/**
* migration function must be the file name concat with Migrate
* key word Example : user will be UserMigrate
 */
func PageMigrate() {
	config.DB.AutoMigrate(&Page{})
}

/*
* event run after add Page
 */
func (u *Page) AfterCreate(scope *gorm.Scope) (err error) {
	IncreaseOnCreate("pages")
	return
}

/*
* event run after delete Faq
 */
func (u *Page) AfterDelete(tx *gorm.DB) (err error) {
	DecreaseOnDelete(u.Status, "pages")
	return
}

/**
* update status
 */
func (u *Page) BeforeUpdate() (err error) {
	var page Page
	config.DB.First(&page , u.ID)
	if page.Status != u.Status{
		DecreaseRow(page.Status, "pages")
		IncreaseRow(u.Status , "pages")
	}
	return
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
	return db.Where("status = " + ACTIVE)
}
