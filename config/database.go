package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/subosito/gotenv"
	"os"
	"strconv"
)

/**
* global params DB Will hold the connection with Singleton pattern
* and error params will hold any error in the package
 */

var DB *gorm.DB = nil
var err error = nil

/**
* connect with data base with env file params
* just edit all data in .env file
 */
func ConnectToDatabase() {
	err = gotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	if DB == nil {
		DB, err = gorm.Open("mysql", os.Getenv("DATABASE_USERNAME")+":"+os.Getenv("DATABASE_PASSWORD")+"@tcp("+os.Getenv("DATABASE_HOST")+":"+os.Getenv("DATABASE_PORT")+")/"+os.Getenv("DATABASE_NAME")+"?charset=utf8&parseTime=True&loc=Local")
	}
	debug, _ := strconv.ParseBool(os.Getenv("DEBUG_DATABASE"))
	if os.Getenv("APP_ENV") == "local" {
		DB.LogMode(debug)
	}
}
