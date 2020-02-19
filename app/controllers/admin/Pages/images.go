package pages

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang-starter/app/models"
	"golang-starter/app/requests/admin/page"
	"golang-starter/app/transformers"
	"golang-starter/config"
	"golang-starter/helpers"
	"strconv"
)

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
	var page models.Page
	// check if this id exits , abort if not
	if models.InItApi(g).FindOrFail(g.Param("id"), &page); page.ID == 0 {
		return
	}
	/// upload images and insert in database
	insertImageInDataBase(g, row.Images, id)
	/// get the new data with images
	var newPage models.Page
	if models.InItApi(g).FindOrFail(g.Param("id"), &row , (func(db *gorm.DB) {
		db.Preload("Translations" , "lang = ?", helpers.LangHeader(g))
	}) , (func(db *gorm.DB) {
		db.Preload("Images")
	})); newPage.ID == 0 {
		return
	}
	helpers.OkResponse(g, helpers.DoneUpdate(g), transformers.PageResponse(newPage))
}

/***
* Delete images
* Delete image by id
 */
func DeleteImage(g *gin.Context) {
	// find this row or return 404
	var row models.PageImage
	// check if this id exits , abort if not
	if models.InItApi(g).FindOrFail(g.Param("id"), &row); row.ID == 0 {
		return
	}
	config.DB.Unscoped().Delete(&row)
	// now return ok response
	helpers.OkResponseWithOutData(g, helpers.DoneDelete(g))
}

/***
* Delete images assign to page by page id
 */
func DeletePageImages(g *gin.Context) {
	// find this row or return 404
	var row models.Page
	// check if this id exits , abort if not
	if models.InItApi(g).FindOrFail(g.Param("id"), &row); row.ID == 0 {
		return
	}
	deleteAllPageImage(row.ID)
	// now return ok response
	helpers.OkResponseWithOutData(g, helpers.DoneDelete(g))
	return
}

/***
*  get all page image and delete
 */
func deleteAllPageImage(id uint) {
	var rows models.PageImage
	config.DB.Unscoped().Where("page_id = ? ", id).Delete(&rows)
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
