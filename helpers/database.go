package helpers

import (
	"github.com/fatih/structs"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/gorm"
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
func UpdateOnlyAllowColumns(structNeedToMap interface{}, fillAble []string) interface{} {
	row := structs.Map(structNeedToMap)
	var data = make(map[string]interface{})
	for _, value := range fillAble {
		if row[strings.Title(value)] != "" {
			data[value] = row[strcase.ToCamel(value)]
		}
	}
	return data
}

/**
* add preload dynamic
* this will allow to add more than one preload
 */
func PreloadD(db *gorm.DB, preload []string) *gorm.DB {
	if len(preload) > 0 {
		for _, p := range preload {
			db = db.Preload(p)
		}
		return db
	}
	return db
}
