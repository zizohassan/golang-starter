package models

import (
	"golang-starter/config"
	"os"
	"strconv"
)

func MigrateAllTable() {
	deleteTables, _ := strconv.ParseBool(os.Getenv("DROP_ALL_TABLES"))
	if deleteTables {
		DbTruncate()
	}
	AnswerMigrate()
	CategoryMigrate()
	FaqMigrate()
	PageMigrate()
	PageImageMigrate()
	SettingMigrate()
	TranslationMigrate()
	UserMigrate()
	ActionMigrate()
}

/**
* drop all tables
*/

type query struct {
	Query string
}

func DbTruncate() {
	var query []query
	config.DB.Table("information_schema.tables").Select("concat('DROP TABLE IF EXISTS `', table_name, '`;') as query").Where("table_schema = ? " , "starter").Find(&query)
	for _ , q :=range query{
		config.DB.Exec(q.Query)
	}
}