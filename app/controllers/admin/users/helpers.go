package users

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"golang-starter/app/models"
	"golang-starter/app/requests/admin/user"
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
	if  g.Query("status") != ""{
		if g.Query("status") != "all"{
			filter = append(filter, `status = "` + g.Query("status") + `"`)
		}
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
* update row make sure you used UpdateOnlyAllowColumns to update allow columns
* use fill able method to only update what you need
 */
func updateColumns(data *models.User, oldRow *models.User) {
	////check if password not empty
	if data.Password != "" {
		password, _ := helpers.HashPassword(data.Password)
		data.Password = password
	}
	// update based on fill able data and assign the new data
	// the new data will set in the same pointer
	models.Update(data, oldRow, models.UserFillAbleColumn())
}
