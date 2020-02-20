package models

import (
	"golang-starter/config"
)

/**
* add to active and all
 */
func IncreaseOnCreate(moduleName string) {
	Increase("actions", "count", nil, `slug =  "`+ACTIVE+`_`+moduleName+`"`)
	Increase("actions", "count", nil, `slug = "`+ALL+`_`+moduleName+`"`)
}

/**
* remove from all and status
 */

func DecreaseOnDelete(status string , moduleName string)  {
	Decrease("actions", "count", nil, `module_name =  "`+moduleName+`"` , `verb = "`+status+`"`)
	Decrease("actions", "count", nil, `slug = "`+ALL+`_user"`)
}

func GetActionByModule(moduleName string) []Action {
	var actions []Action
	config.DB.Where("module_name = ? ", moduleName).Find(&actions)
	return actions
}

