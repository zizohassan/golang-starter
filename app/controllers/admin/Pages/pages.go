package pages

import (
	"github.com/gin-gonic/gin"
	"github.com/iancoleman/strcase"
	"golang-starter/app/models"
	"golang-starter/app/transformers"
	"golang-starter/config"
	"golang-starter/helpers"
)

/***
* get all rows with pagination
 */
func Index(g *gin.Context) {
	// array of rows
	var rows []models.Page
	// query before any thing
	paginator := helpers.Paging(&helpers.Param{
		DB:      config.DB,
		Page:    helpers.Page(g),
		Limit:   helpers.Limit(g),
		OrderBy: helpers.Order("id desc"),
		Filters: filter(g),
		Preload: preload(),
		ShowSQL: true,
	}, &rows)
	// transform slice
	paginator.Records = transformers.PagesResponse(rows)
	// return response
	helpers.OkResponseWithPaging(g, helpers.DoneGetAllItems(g), paginator)
}

/***
* return row with id
 */
func Show(g *gin.Context) {
	// find this row or return 404
	row, find := FindOrFail(g.Param("id"))
	if !find {
		helpers.ReturnNotFound(g, helpers.ItemNotFound(g))
		return
	}
	// now return row data after transformers
	helpers.OkResponse(g, helpers.DoneGetItem(g), transformers.PageResponse(row))
}

/**
* update page
 */
func Update(g *gin.Context) {
	// check if request valid
	valid, row := validateRequest(g)
	if !valid {
		return
	}
	// find this row or return 404
	oldRow, find := FindOrFail(g.Param("id"))
	if !find {
		helpers.ReturnNotFound(g, helpers.ItemNotFound(g))
		return
	}
	/// delete all images if reset flag in the url
	if g.Query("reset") != "" {
		deleteAllPageImage(g)
		return
	}
	/// upload images
	insertImageInDataBase(g, row.Image, int(oldRow.ID))
	/// insert translations
	insertTranslationsInDataBase(row.Translation, int(oldRow.ID))
	/// update allow columns
	oldRow = updateColumns(row, oldRow)
	// now return row data after transformers
	helpers.OkResponse(g, helpers.DoneUpdate(g), transformers.PageResponse(oldRow))
}

/***
* loop on translation
* insert in database
 */
func insertTranslationsInDataBase(translations []map[string]string, id int) {
	if len(translations) > 0 {
		for _, translation := range translations {
			config.DB.Create(&models.Translation{
				PageId: id,
				Value:  translation["value"],
				Slug:   strcase.ToSnake(translation["slug"]),
				Lang:   translation["lang"],
			})
		}
	}
}
