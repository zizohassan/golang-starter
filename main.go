package main

import (
	"github.com/bykovme/gotrans"
	"golang-starter/app/models"
	"golang-starter/app/requests"
	"golang-starter/config"
	"golang-starter/providers"
	"golang-starter/seeders"
	"os"
)

func main() {

	/**
	* start multi language
	 */
	err := gotrans.InitLocales("public/trans")
	if err != nil {
		panic(err)
	}
	/**
	* add custom role to validation
	*
	 */
	requests.Init()
	/**
	* connect with data base logic you can edit .env file to
	* change any connection params
	 */
	config.ConnectToDatabase()
	/**
	* start container will carry all models
	* and database connection
	 */
	providers.StartContainer()
	/**
	* drop All tables and migrate
	* to stop delete tables make DROP_ALL_TABLES false in env file
	* if you need to stop auto migration just stop this line
	 */
	models.MigrateAllTable(os.Getenv("PRODUCTION_MODEL_PATH"))
	/**
	* this function will open seeders folder look inside all files
	* search for seeders function and seed execute these function
	* if you need to stop seeding you can stop this line
	 */
	seeders.Seed()
	/**
	* Run gin framework
	* add middleware
	* run routing
	* serve app
	 */
	providers.Run()
}
