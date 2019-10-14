package pages

import (
	"github.com/gin-gonic/gin"
	"golang-starter/app/models"
	"golang-starter/app/requests/admin/page"
	"golang-starter/config"
	"golang-starter/helpers"
)

/**
* filter module with some columns
 */
func filter(g *gin.Context) []string {
	var filter []string
	if g.Query("status") != "" {
		filter = append(filter, "status = "+g.Query("status"))
	}
	if g.Query("name") != "" {
		filter = append(filter, `name like "%`+g.Query("name")+`%"`)
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
* preload function when findOrFail
 */


/**
* here we will check if request valid or not
 */
func validateRequest(g *gin.Context) (bool, *models.Page) {
	// init struct to validate request
	row := new(models.Page)
	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := page.StoreUpdate(g.Request, row)
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
func FindOrFail(id interface{}) (models.Page, bool) {
	var oldRow models.Page
	db := config.DB.Where("id = ? ", id)
	db = helpers.PreloadD(db  , []string{"Translations", "Images"})
	db.Find(&oldRow)
	if oldRow.ID != 0 {
		return oldRow, true
	}
	return oldRow, false
}

/**
* update row make sure you used UpdateOnlyAllowColumns to update allow columns
* use fill able method to only update what you need
 */
func updateColumns(row *models.Page, oldRow models.Page) models.Page {
	onlyAllowData := helpers.UpdateOnlyAllowColumns(row, models.PageFillAbleColumn())
	config.DB.Model(&oldRow).Updates(onlyAllowData)
	newData  , _ :=  FindOrFail(oldRow.ID)
	return newData
}
