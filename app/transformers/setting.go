package transformers

import "golang-starter/app/models"

/**
* stander the single setting response
 */
func SettingResponse(translation models.Setting) map[string]interface{} {
	var u = make(map[string]interface{})
	u["value"] = translation.Value
	u["id"] = translation.ID
	u["name"] = translation.Name
	u["slug"] = translation.Slug
	u["setting_type"] = translation.SettingType

	return u
}

/**
* stander the Multi settings response
 */
func SettingsResponse(translations []models.Setting) []map[string]interface{} {
	var u  = make([]map[string]interface{} , 0)
	for _ , translation := range translations {
		u = append(u , SettingResponse(translation))
	}
	return u
}

