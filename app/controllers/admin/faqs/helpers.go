package faqs

import (
	"github.com/gin-gonic/gin"
	"golang-starter/app/models"
	"golang-starter/app/requests/admin/faq"
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
	if g.Query("question") != "" {
		filter = append(filter, `question like "%`+g.Query("question")+`%"`)
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
func validateRequest(g *gin.Context) (bool, *models.Faq) {
	// init struct to validate request
	row := new(models.Faq)
	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := faq.StoreUpdate(g.Request, row)
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
func FindOrFail(id interface{}) (models.Faq, bool) {
	var oldRow models.Faq
	config.DB.Where("id = ?", id).Find(&oldRow)
	if oldRow.ID != 0 {
		return oldRow, true
	}
	return oldRow, false
}

/**
* find or fail with preload
*/
func FindOrFailWithPreload(id interface{}) (models.Faq, bool) {
	var oldRow models.Faq
	config.DB.Where("id = ?", id).Preload("Answers").Find(&oldRow)
	if oldRow.ID != 0 {
		return oldRow, true
	}
	return oldRow, false
}


/**
* update row make sure you used UpdateOnlyAllowColumns to update allow columns
* use fill able method to only update what you need
 */
func updateColumns(row *models.Faq, oldRow models.Faq) models.Faq {
	onlyAllowData := helpers.UpdateOnlyAllowColumns(row, models.FaqFillAbleColumn())
	/// disable auto association (we will update only what we need)
	config.DB.Model(&oldRow).Updates(onlyAllowData)
	/// return data with preload relations
	newData, _ := FindOrFailWithPreload(oldRow.ID)

	return newData
}
