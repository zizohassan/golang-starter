package transformers

import "golang-starter/app/models"

/**
* stander the single Answer response
 */
func AnswerResponse(answer models.Answer) map[string]interface{} {
	var u = make(map[string]interface{})
	u["id"] = answer.ID
	u["answer"] = answer.Text

	return u
}

/**
* stander the Multi Answers response
 */
func AnswersResponse(answers []models.Answer) []map[string]interface{} {
	var u  = make([]map[string]interface{} , 0)
	for _ , answer := range answers {
		u = append(u , AnswerResponse(answer))
	}
	return u
}
