package pages

import (
	"github.com/gin-gonic/gin"
	"golang-starter/app/models"
)

/**
* get faq with answers
*/
func Faqs(g *gin.Context) {
	///// declare variables
	//var rows []models.Faq
	//config.DB.Scopes(models.ActiveFaq).Preload("Answers").Find(&rows)
	///// now return row data after transformers
	//helpers.OkResponse(g, helpers.DoneGetItem(g), transformers.FaqsResponse(rows))

	models.InItApi(g)

}
