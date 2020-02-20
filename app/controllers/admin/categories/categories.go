package categories

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
	var rows []models.Category
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
	response := make(map[string]interface{})
	response["status"] = models.GetActionByModule("categories")
	response["records"] = transformers.CategoriesResponse(rows)
	// return response
	helpers.OkResponseWithPaging(g, helpers.DoneGetAllItems(g), paginator)
}

/**
* store new category
 */
func Store(g *gin.Context) {
	// check if request valid
	valid, row := validateRequest(g)
	if !valid {
		return
	}
	// create new row
	config.DB.Create(&row)
	//now return row data after transformers
	helpers.OkResponse(g, helpers.DoneCreateItem(g), transformers.CategoryResponse(*row))
}

/***
* return row with id
 */
func Show(g *gin.Context) {
	var row models.Category
	// check if this id exits , abort if not
	if models.InItApi(g).FindOrFail(g.Param("id"), &row); row.ID == 0 {
		return
	}
	// now return row data after transformers
	helpers.OkResponse(g, helpers.DoneGetItem(g), transformers.CategoryResponse(row))
}

/***
* delete row with id
 */
func Delete(g *gin.Context) {
	var row models.Category
	// check if this id exits , abort if not
	if models.InItApi(g).FindOrFail(g.Param("id"), &row); row.ID == 0 {
		return
	}
	config.DB.Unscoped().Delete(&row)
	// now return ok response
	helpers.OkResponseWithOutData(g, helpers.DoneDelete(g))
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
	var oldRow models.Category
	// check if this id exits , abort if not
	if models.InItApi(g).FindOrFail(g.Param("id"), &oldRow); oldRow.ID == 0 {
		return
	}
	/// update allow columns
	updateColumns(row, &oldRow)
	// now return row data after transformers
	helpers.OkResponse(g, helpers.DoneUpdate(g), transformers.CategoryResponse(oldRow))
}
