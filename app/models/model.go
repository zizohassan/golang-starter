package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang-starter/config"
	"golang-starter/helpers"
	"reflect"
)

/// actions
var ACTIVE = "activate"
var DEACTIVATE = "deactivate"
var BLOCK = "block"
var ALL = "all"
var TRASH = "trash"

type G struct {
	Gin *gin.Context
}

type ModelFunction func(db *gorm.DB)

/**
* id interface you can assign id as string or unit or int
* structBind interface you can assign any type of model
* we select the first id based on struct with gorm
* then  retrieve id from struct with refection
* if we not found id we will abort gin and return
 */
func (g G) FindOrFail(id interface{}, structBind interface{}, appendFunction ...ModelFunction) {
	appendFunctionsToQuery(appendFunction).Where("id = ? ", id).First(structBind)
	findId := reflect.ValueOf(structBind).Elem().FieldByName("ID").Uint()
	if findId == 0 {
		helpers.ReturnNotFound(g.Gin, helpers.ItemNotFound(g.Gin))
		g.Gin.Abort()
		return
	}
}

/**
* loop and handel where cases
* and preload
 */
func appendFunctionsToQuery(functions []ModelFunction) *gorm.DB {
	db := config.DB
	if len(functions) > 0 {
		for _, function := range functions {
			function(db)
		}
	}
	return db
}

/***
* short hand for find by struct
* it will be useful when you want to update
 */
func FindS(structBind interface{}, appendFunction ...ModelFunction) {
	appendFunctionsToQuery(appendFunction).First(structBind)
}

/***
* short hand for find by id
 */
func Find(id interface{}, structBind interface{}, appendFunction ...ModelFunction) {
	appendFunctionsToQuery(appendFunction).Where("id = ? ", id).First(structBind)
}

/**
* short hand to update data with fill able data
* then return with the new data
 */
func Update(data interface{}, row interface{}, allowColumns []string, preloads ...string) {
	onlyAllowData := helpers.UpdateOnlyAllowColumns(data, allowColumns)
	db := config.DB.Model(row).Updates(onlyAllowData)
	if len(preloads) > 0 {
		for _, preload := range preloads {
			db = db.Preload(preload)
		}
	}
	db.Scan(row)
}

/**
* Set incoming gin object to
 */
func InItApi(g *gin.Context) G {
	var k G
	k.Gin = g
	return k
}

/***
* increase 1 on column or with some conditions
 */
func Increase(tableName string, columnName string, id interface{}, where ...string) {
	len := len(where)
	if len > 0 {
		db := "UPDATE " + tableName + " SET " + columnName + " = " + columnName + " + 1 WHERE "
		for i, w := range where {
			db += ` ` + w + ` `
			if (len -1) != i{
				db += ` AND`
			}
		}
		config.DB.Exec(db)
		return
	}
	config.DB.Exec("UPDATE "+tableName+" SET "+columnName+" = "+columnName+" + 1 WHERE id = ?", id)
}


/***
* Decrease 1 on column or with some conditions
 */
func Decrease(tableName string, columnName string, id interface{}, where ...string) {
	if len(where) > 0 {
		db := "UPDATE " + tableName + " SET " + columnName + " = " + columnName + " - 1 WHERE "
		for _, w := range where {
			db += ` ` + w + ` `
			db += ` AND `
		}
		db += columnName + " > 0 "
		config.DB.Exec(db)
		return
	}
	config.DB.Exec("UPDATE "+tableName+" SET "+columnName+" = "+columnName+" - 1 WHERE id = ? AND WHERE "+columnName+" > 0", id)
}
