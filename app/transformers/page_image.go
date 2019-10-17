package transformers

import "golang-starter/app/models"

/**
* stander the single image page response
 */
func PageImageResponse(pageImage models.PageImage) map[string]interface{} {
	var u = make(map[string]interface{})
	u["image"] = pageImage.Image
	u["id"] = pageImage.ID
	u["page_id"] = pageImage.PageId

	return u
}

/**
* stander the Multi images page response
 */
func PageImagesResponse(pageImages []models.PageImage) []map[string]interface{} {
	var u = make([]map[string]interface{}, 0)
	for _, pageImage := range pageImages {
		u = append(u, PageImageResponse(pageImage))
	}
	return u
}
