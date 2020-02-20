package transformers

import "golang-starter/app/models"

/**
* stander the single user response
 */
func ActionResponse(action models.Action) map[string]interface{} {
	var u = make(map[string]interface{})
	u["id"] = action.ID
	u["noun"] = action.Noun
	u["verb"] = action.Verb
	u["slug"] = action.Slug
	u["module_name"] = action.ModuleName
	u["slug"] = action.Slug
	u["count"] = action.Count

	return u
}

/**
* stander the Multi users response
 */
func ActionsResponse(actions []models.Action) []map[string]interface{} {
	var u  = make([]map[string]interface{} , 0)
	for _ , action := range actions {
		u = append(u , ActionResponse(action))
	}
	return u
}

