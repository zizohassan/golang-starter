package test

import (
	"golang-starter/app/models"
	"golang-starter/config"
	"os"
)

/**
* register new user
* return token as header
*/
func getTokenAsHeader(migrate bool) map[string]string {
	token := addAdminUser(migrate)
	var authToken = make(map[string]string)
	authToken["Authorization"] = token

	return authToken
}

/***
* create new admin
* return with admin token
*/
func addAdminUser(migrate bool) string {
	// connect database
	config.ConnectToDatabase()
	/// drop data base
	if migrate {
		models.MigrateAllTable(os.Getenv("TEST_MODEL_PATH"))
	}
	data := models.User{
		Name:     "Abdel Aziz",
		Role:     2,
		Block:    2,
		Email:    "zizo199988@gmail.com",
		Password: "1234567",
	}
	config.DB.Create(&data)

	return data.Token
}

