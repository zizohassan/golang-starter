package transformers

import "golang-starter/app/models"

/**
* stander the single page response
 */
func PageResponse(page models.Page) map[string]interface{} {
	var u = make(map[string]interface{})
	u["name"] = page.Name
	u["id"] = page.ID
	u["status"] = page.Status
	u["created_at"] = page.CreatedAt
	u["updated_at"] = page.UpdatedAt
	u["translations"] = TranslationsResponse(page.Translations)
	u["images"] = PageImagesResponse(page.Images)

	return u
}

/**
* stander the Multi pages response
 */
func PagesResponse(pages []models.Page) []map[string]interface{} {
	var u = make([]map[string]interface{}, 0)
	for _, page := range pages {
		u = append(u, PageResponse(page))
	}

	return u
}
