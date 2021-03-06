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
	Name   string `gorm:"type:varchar(50);" json:"name"`
	Status string `gorm:"type:varchar(20);" json:"status"`
}

/**
* migration function must be the file name concat with Migrate
* key word Example : user will be UserMigrate
 */
func CategoryMigrate() {
	config.DB.AutoMigrate(&Category{})
}

/*
* event run after add Category
 */
func (u *Category) AfterCreate(scope *gorm.Scope) (err error) {
	IncreaseOnCreate("categories")
	return
}

/*
* event run after delete Category
 */
func (u *Category) AfterDelete(tx *gorm.DB) (err error) {
	DecreaseOnDelete(u.Status , "categories")
	return
}

/**
* update status
 */
func (u *Category) BeforeUpdate() (err error) {
	var category Category
	config.DB.First(&category , u.ID)
	if category.Status != u.Status{
		DecreaseRow(category.Status, "categories")
		IncreaseRow(u.Status , "categories")
	}
	return
}

/**
* you can update these column only
 */
func CategoryFillAbleColumn() []string {
	return []string{"name", "status"}
}

/**
* active category only
 */
func ActiveCategory(db *gorm.DB) *gorm.DB {
	return db.Where("status = " + ACTIVE)
}
