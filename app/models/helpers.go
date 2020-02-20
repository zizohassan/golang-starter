package models

import (
	"golang-starter/config"
)

/**
* add to active and all
 */
func IncreaseOnCreate(moduleName string) {
	IncreaseRow(ACTIVE , moduleName)
	Increase("actions", "count", nil, `slug = "`+ALL+`_`+moduleName+`"`)
}

/**
* remove from all and status
 */

func DecreaseOnDelete(status string , moduleName string)  {
	DecreaseRow(status , moduleName)
	Decrease("actions", "count", nil, `slug = "`+ALL+`_user"`)
}

func DecreaseRow(status string , moduleName string)  {
	Decrease("actions", "count", nil, `module_name =  "`+moduleName+`"` , `verb = "`+status+`"`)
}

func IncreaseRow(status string , moduleName string)  {
	Increase("actions", "count", nil, `module_name =  "`+moduleName+`"` , `verb = "`+status+`"`)
}

func GetActionByModule(moduleName string) []Action {
	var actions []Action
	config.DB.Where("module_name = ? ", moduleName).Find(&actions)
	return actions
}

