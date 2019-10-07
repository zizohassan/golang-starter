package translations

import (
	"github.com/gin-gonic/gin"
	"golang-starter/app/models"
	"golang-starter/app/transformers"
	"golang-starter/config"
	"golang-starter/helpers"
)

var moduleName = "Translation"

/***
* get all rows with pagination
*/
func Index(g *gin.Context)  {
	// array of rows
	var rows []models.Translation
	// query before any thing
	paginator := helpers.Paging(&helpers.Param{
		DB:      config.DB,
		Page:    helpers.Page(g),
		Limit:   helpers.Limit(g),
		OrderBy: helpers.Order("id desc"),
		Filters : filter(g),
		Preload : preload(),
		ShowSQL: true,
	}, &rows)
	// transform slice
	paginator.Records = transformers.TranslationsResponse(rows)
	// return response
	helpers.OkResponseWithPaging(g , "Here is our "+moduleName , paginator)
}

/***
* return row with id
*/
func Show(g *gin.Context)  {
	// find this row or return 404
	row , find := findOrFail(g)
	if !find {
		helpers.ReturnNotFound(g , "we not found row id")
		return
	}
	// now return row data after transformers
	helpers.OkResponse(g, moduleName+" Created Successfully", transformers.TranslationResponse(row))
}

/**
* update category
*/
func Update(g *gin.Context)  {
	// check if request valid
	valid  , row := validateRequest(g)
	if !valid {
		return
	}
	// find this row or return 404
	oldRow , find := findOrFail(g)
	if !find {
		helpers.ReturnNotFound(g , "we not found row id")
		return
	}
	/// update allow columns
	oldRow = updateColumns(row , oldRow)
	// now return row data after transformers
	helpers.OkResponse(g, moduleName+" Updated Successfully", transformers.TranslationResponse(oldRow))
}

