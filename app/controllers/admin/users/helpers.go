package users

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"golang-starter/app/models"
	"golang-starter/app/requests/admin/user"
	"golang-starter/config"
	"golang-starter/helpers"
)

/**
* filter module with some columns
 */
func filter(g *gin.Context) []string {
	var filter []string
	if g.Query("block") != "" {
		filter = append(filter, "block = "+g.Query("block"))
	}
	if g.Query("name") != "" {
		filter = append(filter, `name like "%`+g.Query("name")+`%"`)
	}
	if g.Query("email") != "" {
		filter = append(filter, `email like "%`+g.Query("email")+`%"`)
	}
	if g.Query("role") != "" {
		filter = append(filter, `role like "%`+g.Query("role")+`%"`)
	}
	return filter
}

/**
* preload module with some preload conditions
 */
func preload() []string {
	return []string{}
}

/**
* here we will check if request valid or not
 */
func validateRequest(g *gin.Context, action string) (bool, *models.User) {
	var err *govalidator.Validator
	// init struct to validate request
	row := new(models.User)
	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	if action == "store" {
		err = user.Store(g.Request, row)
	} else {
		err = user.Update(g.Request, row)
	}
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if helpers.ReturnNotValidRequest(err, g) {
		return false, row
	}
	return true, row
}

/**
* findOrFail Data
 */
func findOrFail(g *gin.Context) (models.User, bool) {
	var oldRow models.User
	config.DB.Find(&oldRow, "id = "+g.Param("id"))
	if oldRow.ID != 0 {
		return oldRow, true
	}
	return oldRow, false
}

/**
* update row make sure you used UpdateOnlyAllowColumns to update allow columns
* use fill able method to only update what you need
 */
func updateColumns(row *models.User, oldRow models.User) models.User {
	if row.Password != "" {
		password, _ := helpers.HashPassword(row.Password)
		row.Password = password
	}
	onlyAllowData := helpers.UpdateOnlyAllowColumns(row, models.UserFillAbleColumn())
	config.DB.Model(&oldRow).Updates(onlyAllowData).Find(&oldRow)
	return oldRow
}
