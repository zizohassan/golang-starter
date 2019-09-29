package seeders

import (
	"reflect"
	"golang-starter/helpers"
	"strings"
)

type Seeder struct {}

/***
* Open seed folder get All seed Files
* run seed functions inside these files
 */
func Seed() {
	var t Seeder
	seedFiles := helpers.ReadAllFiles("seeders")
	for _, file := range seedFiles {
		filepath := strings.Split(file, ".")
		fileName := filepath[0]
		if fileName != "seed" {
			functionName := strings.Title(filepath[0])+"Seeder"
			reflect.ValueOf(&t).MethodByName(functionName).Call([]reflect.Value{})
		}
	}
}
