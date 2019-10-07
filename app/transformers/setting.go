package transformers

import "golang-starter/app/models"

/**
* stander the single user response
 */
func TranslationResponse(translation models.Translation) map[string]interface{} {
	var u = make(map[string]interface{})
	u["value"] = translation.Value
	u["id"] = translation.ID
	u["lang"] = translation.Lang
	u["slug"] = translation.Slug
	u["page"] = translation.Page

	return u
}

/**
* stander the Multi users response
 */
func TranslationsResponse(translations []models.Translation) []map[string]interface{} {
	var u  = make([]map[string]interface{} , 0)
	for _ , translation := range translations {
		u = append(u , TranslationResponse(translation))
	}
	return u
}
