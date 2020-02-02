package seeders

import (
"golang-starter/app/models"
"golang-starter/config"
"strings"
)

/***
*	Seed Function must Have the same file Name then Add Seeder key word
* 	Example :  file is user function must be UserSeeder
 */
func  SettingSeeder() {
	settings := settings()
	for slug, setting := range settings {
		newSetting(slug, setting)
	}
}

/**
* fake data and create data base
 */
func newSetting(slug string, value string) {
	data := models.Setting{
		Slug:        slug,
		Value:       value,
		Name:        strings.Title(slug),
		SettingType: "text",
	}
	config.DB.Create(&data)
}

func settings() map[string]string {
	var m = make(map[string]string)
	m["twitter"] = "http://twitter.com"
	m["facebook"] = "http://facebook.com"
	m["youtube"] = "http://youtube.com"
	m["linkedin"] = "https://linkedin.com"

	return m
}

