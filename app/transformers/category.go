package transformers

import "golang-starter/app/models"

/**
* stander the single user response
 */
func CategoryResponse(category models.Category) map[string]interface{} {
	var u = make(map[string]interface{})
	u["name"] = category.Name
	u["id"] = category.ID
	u["status"] = category.Status

	return u
}

/**
* stander the Multi users response
 */
func CategoriesResponse(categories []models.Category) []map[string]interface{} {
	var u  = make([]map[string]interface{} , 0)
	for _ , category := range categories {
		u = append(u , CategoryResponse(category))
	}
	return u
}

