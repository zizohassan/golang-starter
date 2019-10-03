package categories

import (
	"github.com/gin-gonic/gin"
	"golang-starter/app/models"
	"golang-starter/app/requests/admin"
	"golang-starter/config"
	"golang-starter/helpers"
)

/**
* filter categories with some columns
 */
func filter(g *gin.Context) []string {
	var filter []string
	if  g.Query("status") != ""{
		filter = append(filter, "status = " + g.Query("status"))
	}
	if  g.Query("name") != ""{
		filter = append(filter, `name like "%` + g.Query("name") + `%"`)
	}
	return filter
}

/**
* filter categories with some preload conditions
 */
func preload() []string {
	return []string{}
}

/**
* here we will check if request valid or not
 */
func validateRequest(g *gin.Context) (bool , *models.Category)   {
	// init struct to validate request
	row := new(models.Category)
	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := admin.StoreUpdate(g.Request, row)
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
* findOrFail Data
 */
func findOrFail(g *gin.Context) (models.Category , bool)  {
	var oldRow models.Category
	config.DB.Find(&oldRow , "id = "+g.Param("id"))
	if oldRow.ID != 0{
		return   oldRow , true
	}
	return  oldRow , false
}

/**
* update row make sure you used UpdateOnlyAllowColumns to update allow columns
* use fill able method to only update what you need
*/
func updateColumns(row *models.Category , oldRow models.Category)  {
	onlyAllowData := helpers.UpdateOnlyAllowColumns(row , models.FillAbleColumn())
	config.DB.Model(&oldRow).Select(models.FillAbleColumn()).Updates(onlyAllowData)
}
