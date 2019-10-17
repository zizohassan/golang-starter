package transformers

import "golang-starter/app/models"

/**
* stander the single faq response
 */
func FaqResponse(faq models.Faq) map[string]interface{} {
	var u = make(map[string]interface{})
	u["question"] = faq.Question
	u["id"] = faq.ID
	u["status"] = faq.Status
	u["answer"] = AnswersResponse(faq.Answers)

	return u
}

/**
* stander the Multi faqs response
 */
func FaqsResponse(faqs []models.Faq) []map[string]interface{} {
	var u = make([]map[string]interface{}, 0)
	for _, faq := range faqs {
		u = append(u, FaqResponse(faq))
	}
	return u
}
