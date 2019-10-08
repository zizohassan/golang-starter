package helpers

import (
	"github.com/fatih/structs"
	"golang-starter/config"
	"strings"
)

/***
* truncate tables
*/
func DbTruncate(tableName ...string) {
	for _, table := range tableName {
		config.DB.Exec("TRUNCATE " + table)
	}
}

/**
* this function get struct and return with only
* Available column that allow to updated depend on FillAbleColumn function
* this for security
* map struct to update
 */
func UpdateOnlyAllowColumns(structNeedToMap interface{} , fillAble []string)  interface{} {
	row := structs.Map(structNeedToMap)
	var data = make(map[string]interface{})
	for _ , value  := range fillAble{
		if row[strings.Title(value)] != ""{
			data[value] = row[strings.Title(value)]
		}
	}
	return  data
}
