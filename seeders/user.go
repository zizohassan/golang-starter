package seeders

import "fmt"

/***
*	Seed Function must Have the same file Name then Add Seeder key word
* 	Example :  file is user function must be UserSeeder
*/
func (s *Seeder) UserSeeder() {
	fmt.Println("seed data")
}
