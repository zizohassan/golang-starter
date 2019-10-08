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
	Block    int    `gorm:"_" json:"block"`
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

/**
* migration function must be the file name concat with Migrate
* key word Example : user will be UserMigrate
 */
func (s *MigrationTables) UserMigrate() {
	config.DB.AutoMigrate(&User{})
}

/**
* you can update these column only
 */
func UserFillAbleColumn() []string {
	return []string{"name", "email", "role", "password", "block"}
}
