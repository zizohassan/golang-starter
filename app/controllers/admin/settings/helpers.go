package settings

import (
	"github.com/gin-gonic/gin"
	"golang-starter/app/models"
	"golang-starter/app/requests/admin/setting"
	"golang-starter/helpers"
)

/**
* filter module with some columns
 */
func filter(g *gin.Context) []string {
	var filter []string
	if  g.Query("value") != ""{
		filter = append(filter, `value like "%` + g.Query("value") + `%"`)
	}
	if  g.Query("name") != ""{
		filter = append(filter, `name like "%` + g.Query("name") + `%"`)
	}
	if  g.Query("slug") != ""{
		filter = append(filter, `slug like "%` + g.Query("slug") + `%"`)
	}
	if  g.Query("setting_type") != ""{
		filter = append(filter, `setting_type like "%` + g.Query("setting_type") + `%"`)
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
func validateRequest(g *gin.Context) (bool , *models.Setting)   {
	// init struct to validate request
	row := new(models.Setting)
	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := setting.StoreUpdate(g.Request, row)
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if helpers.ReturnNotValidRequest(err, g) {
		return false , row
	}
	return true , row
}

/**
* update row make sure you used UpdateOnlyAllowColumns to update allow columns
* use fill able method to only update what you need
*/
func updateColumns(data *models.Setting , oldRow *models.Setting) {
	models.Update(data, oldRow, models.SettingFillAbleColumn())
}
