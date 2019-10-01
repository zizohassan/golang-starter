package models

import (
	"github.com/jinzhu/gorm"
	"golang-starter/helpers"
	"golang-starter/config"
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
	Block    int   `gorm:"_" json:"block"`
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
	Email    string `json:"email"`
}

/**
* event when user register
* create token
* hash password
* set user role
* set block user to not block (1 is blocked 2 is not blocked)
*/
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	token, _ := helpers.HashPassword(user.Email + user.Password)
	password, _ := helpers.HashPassword(user.Password)
	scope.SetColumn("token", token)
	scope.SetColumn("password", password)
	scope.SetColumn("role", 1)
	scope.SetColumn("block", 2)
	return nil
}

/**
* migration function must be the file name concat with Migrate
* key word Example : user will be UserMigrate
 */
func (s *MigrationTables) UserMigrate() {
	config.DB.AutoMigrate(&User{})
}