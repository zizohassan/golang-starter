package faqs

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
	/// array of rows
	var rows []models.Faq
	/// query before any thing
	paginator := helpers.Paging(&helpers.Param{
		DB:      config.DB,
		Page:    helpers.Page(g),
		Limit:   helpers.Limit(g),
		OrderBy: helpers.Order(g,"id desc"),
		Filters: filter(g),
		Preload: preload(),
		ShowSQL: true,
	}, &rows)
	/// transform slice
	paginator.Records = transformers.FaqsResponse(rows)
	/// return response
	helpers.OkResponseWithPaging(g, helpers.DoneGetAllItems(g), paginator)
}

/**
* store new category
 */
func Store(g *gin.Context) {
	/// check if request valid
	valid, row := validateRequest(g)
	if !valid {
		return
	}
	/// create new row
	config.DB.Create(&row)
	/// insert in to answer
	insertAnswerToDataBase(row.Answer, row.ID)
	/// get with new data
	data , _ := FindOrFailWithPreload(row.ID)
	/// now return row data after transformers
	helpers.OkResponse(g, helpers.DoneCreateItem(g), transformers.FaqResponse(data))
}

/***
* return row with id
 */
func Show(g *gin.Context) {
	/// find this row or return 404
	row, find := FindOrFailWithPreload(g.Param("id"))
	if !find {
		helpers.ReturnNotFound(g, helpers.ItemNotFound(g))
		return
	}
	/// now return row data after transformers
	helpers.OkResponse(g, helpers.DoneGetItem(g), transformers.FaqResponse(row))
}

/***
* delete row with id
 */
func Delete(g *gin.Context) {
	/// find this row or return 404
	row, find := FindOrFail(g.Param("id"))
	if !find {
		helpers.ReturnNotFound(g, helpers.ItemNotFound(g))
		return
	}
	/// delete related answers
	deleteAnswers(row.ID)
	/// delete row
	config.DB.Unscoped().Delete(&row)
	/// now return ok response
	helpers.OkResponseWithOutData(g, helpers.DoneDelete(g))
}

/**
* update faq
 */
func Update(g *gin.Context) {
	/// check if request valid
	valid, row := validateRequest(g)
	if !valid {
		return
	}
	/// find this row or return 404
	oldRow, find := FindOrFail(g.Param("id"))
	if !find {
		helpers.ReturnNotFound(g, helpers.ItemNotFound(g))
		return
	}
	/// delete old Answer if user need
	if g.Query("reset") == "true" {
		deleteAnswers(oldRow.ID)
	}
	/// insert in to answer
	insertAnswerToDataBase(row.Answer, oldRow.ID)
	/// update allow columns
	oldRow = updateColumns(row, oldRow)
	/// now return row data after transformers
	helpers.OkResponse(g, helpers.DoneUpdate(g), transformers.FaqResponse(oldRow))
}
