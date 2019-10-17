package faqs

import (
	"golang-starter/app/models"
	"golang-starter/config"
	"golang-starter/helpers"
)

/**
* loop on the answers
* insert all answer in database
*/
func insertAnswerToDataBase(answers []string, id uint) {
	if len(answers) > 0 {
		for _, answer := range answers {
			if answer != ""{
				config.DB.Create(&models.Answer{
					Text: helpers.ClearText(answer),
					FaqId:int(id),
				})
			}
		}
	}
}

/**
*  delete assign pages
*/

func deleteAnswers(faqId uint)  {
	var rows models.Answer
	config.DB.Unscoped().Where("faq_id = ? ", faqId).Delete(&rows)
}
