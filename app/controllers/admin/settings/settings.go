package settings

import (
	"github.com/gin-gonic/gin"
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
	var rows []models.Setting
	// query before any thing
	paginator := helpers.Paging(&helpers.Param{
		DB:      config.DB,
		Page:    helpers.Page(g),
		Limit:   helpers.Limit(g),
		OrderBy: helpers.Order(g,"id desc"),
		Filters: filter(g),
		Preload: preload(),
		ShowSQL: true,
	}, &rows)
	// transform slice
	paginator.Records = transformers.SettingsResponse(rows)
	// return response
	helpers.OkResponseWithPaging(g, helpers.DoneGetAllItems(g), paginator)
}

/***
* return row with id
 */
func Show(g *gin.Context) {
	// find this row or return 404
	var row models.Setting
	// check if this id exits , abort if not
	if models.InItApi(g).FindOrFail(g.Param("id"), &row); row.ID == 0 {
		return
	}
	// now return row data after transformers
	helpers.OkResponse(g, helpers.DoneGetItem(g), transformers.SettingResponse(row))
}

/**
* update category
 */
func Update(g *gin.Context) {
	// check if request valid
	valid, row := validateRequest(g)
	if !valid {
		return
	}
	// find this row or return 404
	var oldRow models.Setting
	// check if this id exits , abort if not
	if models.InItApi(g).FindOrFail(g.Param("id"), &row); row.ID == 0 {
		return
	}
	/// update allow columns
	updateColumns(row, &oldRow)
	// now return row data after transformers
	helpers.OkResponse(g, helpers.DoneUpdate(g), transformers.SettingResponse(oldRow))
}
