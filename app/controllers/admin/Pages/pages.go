package pages

import (
	"github.com/gin-gonic/gin"
	"github.com/iancoleman/strcase"
	"golang-starter/app/models"
	"golang-starter/app/requests/admin/page"
	"golang-starter/app/transformers"
	"golang-starter/config"
	"golang-starter/helpers"
	"strconv"
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
	insertImageInDataBase(g, row.Image, int(oldRow.ID))
	insertTranslationsInDataBase(g, row.Translation, int(oldRow.ID))
	/// update allow columns
	oldRow = updateColumns(row, oldRow)
	// now return row data after transformers
	helpers.OkResponse(g, helpers.DoneUpdate(g), transformers.PageResponse(oldRow))
}

/**
* update image form page
 */
func UploadImage(g *gin.Context) {
	// init struct to validate request
	var row models.PageImageRequest
	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := page.UploadStoreUpdate(g.Request, &row)
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if helpers.ReturnNotValidRequest(err, g) {
		return
	}
	///get id
	id, _ := strconv.Atoi(g.Param("id"))
	// find this row or return 404
	_, find := FindOrFail(id)
	if !find {
		helpers.ReturnNotFound(g, helpers.ItemNotFound(g))
		return
	}
	insertImageInDataBase(g, row.Images, id)
	/// get the new data with images
	newPage, _ := FindOrFail(id)
	helpers.OkResponse(g, helpers.DoneUpdate(g), transformers.PageResponse(newPage))
}

/**
* upload images
* loop and insert images with id
 */
func insertImageInDataBase(g *gin.Context, images []string, id int) {
	if len(images) > 0 {
		uploadedImages := helpers.MultiDecodeImage(g, images)
		///// loop and insert image in database
		if len(uploadedImages) > 0 {
			for _, upload := range uploadedImages {
				config.DB.Create(&models.PageImage{
					PageId: id,
					Image:  upload,
				})
			}
		}
	}
}

/***
* loop on translation
* insert in database
*/
func insertTranslationsInDataBase(g *gin.Context, translations []map[string]string, id int) {
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
