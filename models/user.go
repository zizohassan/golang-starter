package models

import (
	"github.com/jinzhu/gorm"
	"investment-users/helpers"
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
* stander the single user response
 */
func UserResponse(user User) map[string]interface{} {
	var u = make(map[string]interface{})
	u["name"] = user.Name
	u["email"] = user.Email
	u["role"] = user.Role
	u["token"] = user.Token

	return u
}

/**
* stander the Multi users response
 */
func UsersResponse(users []User) map[uint]map[string]interface{} {
	var u = make(map[uint]map[string]interface{})
	for _, user := range users {
		u[user.ID] = UserResponse(user)
	}
	return u
}

/**
* migration function must be the file name concat with Migrate
* key word Example : user will be UserMigrate
 */
func (s *MigrationTables) UserMigrate() {
	config.DB.AutoMigrate(&User{})
}
