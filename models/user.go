package models

import (
	"github.com/jinzhu/gorm"
	"starter/config"
)

/***
* model struct here we will build the main
* struct that connect to database
 */
type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50);" json:"name"`
	Email    string `gorm:"type:varchar(100);unique_index" json:"email"`
	Role     string `gorm:"size:20" json:"role"`
	Password string `gorm:"size:255" json:"_"`
	Token    string `gorm:"size:255" json:"token"`
}

/**
* use this struct when user login
 */
type Login struct {
	Password string `gorm:"size:255" json:"password"`
	Email    string `gorm:"type:varchar(100);unique_index" json:"email"`
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
* migration function must be the file name concnate with Migrate
* key word Example : user will be UserMigrate
 */
func (s *MigrationTables) UserMigrate() {
	config.DB.AutoMigrate(&User{})
}
