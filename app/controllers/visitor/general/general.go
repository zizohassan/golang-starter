package general

import (
	"github.com/gin-gonic/gin"
	"golang-starter/app/models"
	"golang-starter/app/transformers"
	"golang-starter/config"
	"golang-starter/helpers"
)

/***
* this will be the first api
* call in the hole system
* return all translations
* all setting
*/
//TODO::will cache this response in redis
func Init(g *gin.Context)  {
	/// declare variables
	var settings []models.Setting
	var pages []models.Page
	/// queries
	config.DB.Preload("Translations").Preload("Images").Find(&pages)
	config.DB.Find(&settings)
	/// build response
	var response = make(map[string]interface{})
	response["pages"] = transformers.PagesResponse(pages)
	response["settings"] = transformers.SettingsResponse(settings)
	/// return with data
	helpers.OkResponse(g , helpers.T(g , "init_project") , response)
	return
}
