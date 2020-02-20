package seeders

import (
	"golang-starter/app/models"
	"golang-starter/config"
)

/***
*	Seed Function must Have the same file Name then Add Seeder key word
* 	Example :  file is user function must be UserSeeder
 */
func  PageSeeder() {
	for _, page := range pages() {
		data := models.Page{
			Name:   page,
			Status: models.ACTIVE,
		}
		config.DB.Create(&data)
	}
}

/***
* list of pages
 */
func pages() []string {
	return []string{
		"home",
		"about",
		"contact",
		"terms",
		"police",
	}
}
