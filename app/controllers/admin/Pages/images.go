package pages

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	_, find := FindOrFail(id)
	if !find {
		helpers.ReturnNotFound(g, helpers.ItemNotFound(g))
		return
	}
	/// upload images and insert in database
	insertImageInDataBase(g, row.Images, id)
	/// get the new data with images
	newPage, _ := FindOrFailWithPreload(id, helpers.LangHeader(g))
	helpers.OkResponse(g, helpers.DoneUpdate(g), transformers.PageResponse(newPage))
}

/***
* Delete images
* Delete image by id
 */
func DeleteImage(g *gin.Context) {
	// find this row or return 404
	row, find := FindOrFailImage(g.Param("id"))
	if !find {
		helpers.ReturnNotFound(g, helpers.ItemNotFound(g))
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
	row, find := FindOrFail(g.Param("id"))
	if !find {
		helpers.ReturnNotFound(g, helpers.ItemNotFound(g))
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
		fmt.Println("Images upload ", images)
		uploadedImages := helpers.MultiDecodeImage(g, images)
		fmt.Println("Images upload ", uploadedImages)
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
