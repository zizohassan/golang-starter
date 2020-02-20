package models

import (
	"github.com/jinzhu/gorm"
	"golang-starter/config"
	"golang-starter/helpers"
)

/***
* model struct here we will build the main
* struct that connect to database
 */
type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50);" json:"name"`
	Email    string `gorm:"type:varchar(50);unique_index" json:"email"`
	Role     int    `gorm:"_" json:"role"`
	Password string `gorm:"size:255" json:"password"`
	Token    string `gorm:"size:255" json:"token"`
	Status   string `gorm:"type:varchar(20);" json:"status"`
}

/**
* use this struct when visitor login
 */
type Login struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

/**
* use this struct when reset email
 */
type Reset struct {
	Email string `json:"email"`
}

/**
* use this struct when reset email
 */
type Recover struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

/**
* event when user register
* create token
* hash password
 */
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	token, _ := helpers.HashPassword(user.Email + user.Password)
	password, _ := helpers.HashPassword(user.Password)
	scope.SetColumn("token", token)
	scope.SetColumn("password", password)

	return nil
}

/*
* event run after user register
 */
func (u *User) AfterCreate(scope *gorm.Scope) (err error) {
	IncreaseOnCreate("users")
	return
}

/*
* event run after delete user
 */
func (u *User) AfterDelete(tx *gorm.DB) (err error) {
	DecreaseOnDelete(u.Status , "users")
	return
}

/**
* update status
 */
func (u *User) BeforeUpdate() (err error) {
	var user User
	config.DB.First(&user , u.ID)
	if user.Status != u.Status{
		DecreaseRow(user.Status, "users")
		IncreaseRow(u.Status , "users")
	}
	return
}

/**
* migration function must be the file name concat with Migrate
* key word Example : user will be UserMigrate
 */
func UserMigrate() {
	config.DB.AutoMigrate(&User{})
}

/**
* you can update these column only
 */
func UserFillAbleColumn() []string {
	return []string{"name", "email", "role", "password", "status"}
}

/**
* active category only
 */
func ActiveUser(db *gorm.DB) *gorm.DB {
	return db.Where("status = " + ACTIVE)
}
