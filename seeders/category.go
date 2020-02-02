package seeders

import (
	"golang-starter/app/models"
	"golang-starter/config"
	"syreclabs.com/go/faker"
)

/***
*	Seed Function must Have the same file Name then Add Seeder key word
* 	Example :  file is user function must be UserSeeder
 */
func CategorySeeder() {
	for i := 0 ; i < 10 ; i++ {
		newCategory()
	}
}

/**
* fake data and create data base
 */
func newCategory()  {
	data := models.Category{
		Name:     faker.Internet().UserName(),
		Status:   faker.RandomInt(1,2),
	}
	config.DB.Create(&data)
}
