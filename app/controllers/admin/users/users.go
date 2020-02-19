package users

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
	var rows []models.User
	// query before any thing
	paginator := helpers.Paging(&helpers.Param{
		DB:      config.DB,
		Page:    helpers.Page(g),
		Limit:   helpers.Limit(g),
		OrderBy: helpers.Order(g, "id desc"),
		Filters: filter(g),
		Preload: preload(),
		ShowSQL: true,
	}, &rows)
	// transform slice
	paginator.Records = transformers.UsersResponse(rows)
	// return response
	helpers.OkResponseWithPaging(g, helpers.DoneGetAllItems(g), paginator)
}

/**
* store new user
 */
func Store(g *gin.Context) {
	// check if request valid
	valid, row := validateRequest(g, "store")
	if !valid {
		return
	}
	/// check if this email exists
	var count int
	config.DB.Model(models.User{}).Where("email = ? ", row.Email).Count(&count)
	if count > 0 {
		helpers.ReturnDuplicateData(g, "email")
		return
	}
	// create new row
	config.DB.Create(&row)
	//now return row data after transformers
	helpers.OkResponse(g, helpers.DoneCreateItem(g), transformers.UserResponse(*row))
}

/***
* return row with id
 */
func Show(g *gin.Context) {
	var row models.User
	// check if this id exits , abort if not
	if models.InItApi(g).FindOrFail(g.Param("id"), &row); row.ID == 0 {
		return
	}
	// now return row data after transformers
	helpers.OkResponse(g, helpers.DoneGetItem(g), transformers.UserResponse(row))
}

/***
* delete row with id
 */
func Delete(g *gin.Context) {
	var row models.User
	// check if this id exits , abort if not
	if models.InItApi(g).FindOrFail(g.Param("id"), &row); row.ID == 0 {
		return
	}
	config.DB.Unscoped().Delete(&row)
	// now return row data after transformers
	helpers.OkResponseWithOutData(g, helpers.DoneDelete(g))
}

/**
* update user
 */
func Update(g *gin.Context) {
	// check if request valid
	valid, data := validateRequest(g, "update")
	if !valid {
		return
	}
	// find this row or return 404
	var oldRow models.User
	// check if this id exits , abort if not
	if models.InItApi(g).FindOrFail(g.Param("id"), &oldRow); oldRow.ID == 0 {
		return
	}
	/// check if this email exists
	var count int
	config.DB.Model(models.User{}).Where("email = ? AND email != ?", data.Email, oldRow.Email).Count(&count)
	if count > 0 {
		helpers.ReturnDuplicateData(g, "email")
		return
	}
	/// update allow columns
	updateColumns(data, &oldRow)
	// now return row data after transformers
	helpers.OkResponse(g, helpers.DoneUpdate(g), transformers.UserResponse(oldRow))
}
