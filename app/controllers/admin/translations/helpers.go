package translations

import (
	"github.com/gin-gonic/gin"
	"golang-starter/app/models"
	"golang-starter/app/requests/admin/translation"
	"golang-starter/config"
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
	if  g.Query("page") != ""{
		filter = append(filter, `page like "%` + g.Query("page") + `%"`)
	}
	if  g.Query("slug") != ""{
		filter = append(filter, `slug like "%` + g.Query("slug") + `%"`)
	}
	if  g.Query("lang") != ""{
		filter = append(filter, `lang like "%` + g.Query("lang") + `%"`)
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
func validateRequest(g *gin.Context) (bool , *models.Translation)   {
	// init struct to validate request
	row := new(models.Translation)
	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := translation.StoreUpdate(g.Request, row)
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
func findOrFail(g *gin.Context) (models.Translation , bool)  {
	var oldRow models.Translation
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
func updateColumns(row *models.Translation , oldRow models.Translation) models.Translation {
	onlyAllowData := helpers.UpdateOnlyAllowColumns(row , models.TranslationFillAbleColumn())
	config.DB.Model(&oldRow).Updates(onlyAllowData).Find(&oldRow)
	return oldRow
}
