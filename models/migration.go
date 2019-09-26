package models

import (
	"github.com/jinzhu/inflection"
	"os"
	"reflect"
	"starter/config"
	"starter/helpers"
	"strconv"
	"strings"
)

type MigrationTables struct{}

/**
* first loop in all migrations files
* get all migration methods
* drop related table if .env have delete attribute
* Call migration function
 */
func MigrateAllTable() {
	var t MigrationTables
	migrateFiles := helpers.ReadAllFiles("models")
	for _, file := range migrateFiles {
		filepath := strings.Split(file, ".")
		fileName := filepath[0]
		if fileName != "migration" {
			functionName := strings.Title(filepath[0]) + "Migrate"
			deleteTables, _ := strconv.ParseBool(os.Getenv("DROP_ALL_TABLES"))
			if deleteTables {
				config.DB.DropTableIfExists(inflection.Plural(filepath[0]))
			}
			reflect.ValueOf(&t).MethodByName(functionName).Call([]reflect.Value{})
		}
	}
}
