package categories

import (
	"github.com/gin-gonic/gin"
	"golang-starter/app/models"
	"golang-starter/app/requests/admin"
	"golang-starter/app/transformers"
	"golang-starter/config"
	"golang-starter/helpers"
)

/***
* get all Categories
*/
func Index(g *gin.Context)  {
	/**
	* array of categories
	*/
	var categories []models.Category
	/**
	* query before any thing
	*/
	paginator := helpers.Paging(&helpers.Param{
		DB:      config.DB,
		Page:    helpers.Page(g),
		Limit:   helpers.Limit(g),
		OrderBy: helpers.Order("id desc"),
		Filters : filter(g),
		Preload : preload(),
		ShowSQL: true,
	}, &categories)

	/**
	* transform slice
	*/
	paginator.Records = transformers.CategoriesResponse(categories)
	/**
	* return response
	 */
	helpers.OkResponseWithPaging(g , "here is our categories" , paginator)
}

/**
* store new category
*/
func Store(g *gin.Context)  {
	/**
	* init visitor login struct to validate request
	 */
	category := new(models.Category)
	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := admin.Store(g.Request, category)
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if helpers.ReturnNotValidRequest(err, g) {
		return
	}
	/**
	* create new category
	*/
	config.DB.Create(&category)
	/**
	* now return category data after transformers
	*/
	helpers.OkResponse(g, "Category Created Successfully", transformers.CategoryResponse(*category))
}

/**
* filter categories with some columns
*/
func filter(g *gin.Context) []string {
	var filter []string
	if  g.Query("status") != ""{
		filter = append(filter, "status = " + g.Query("status"))
	}
	return filter
}

/**
* filter categories with some preload conditions
 */
func preload() []string {
	return []string{}
}

