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
func (s *Seeder) UserSeeder() {
	newUser(true)
	for i := 0 ; i < 10 ; i++ {
		newUser(false)
	}
}

/**
* fake data and create data base
*/
func newUser(admin bool)  {
	data := models.User{
		Email:    faker.Internet().Email()  ,
		Password: faker.Internet().Password(8, 14),
		Name:     faker.Internet().UserName(),
	}
	if admin {
		data.Role = 2
	}
	config.DB.Create(&data)
}
