package transformers

import "golang-starter/app/models"

/**
* stander the single user response
 */
func PageResponse(page models.Page) map[string]interface{} {
	var u = make(map[string]interface{})
	u["name"] = page.Name
	u["id"] = page.ID
	u["status"] = page.Status
	u["translations"] = TranslationsResponse(page.Translations)
	u["images"] = PageImagesResponse(page.Images)

	return u
}

/**
* stander the Multi users response
 */
func PagesResponse(pages []models.Page) []map[string]interface{} {
	var u = make([]map[string]interface{}, 0)
	for _, page := range pages {
		u = append(u, PageResponse(page))
	}

	return u
}
