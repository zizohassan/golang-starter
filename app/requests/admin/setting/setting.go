package setting

import (
	"github.com/thedevsaddam/govalidator"
	"golang-starter/app/models"
	"golang-starter/helpers"
	"net/http"
)

/**
* validate store setting request
 */
func StoreUpdate(r *http.Request, request *models.Setting) *govalidator.Validator {
	lang := helpers.GetCurrentLangFromHttp(r)
	/// Validation rules
	rules := govalidator.MapData{
		"value":        []string{"required", "min:2", "max:255"},
		"name":         []string{"required", "min:2", "max:50"},
		"setting_type": []string{"required", "min:2", "max:50"},
		"slug":         []string{"required", "min:2", "max:50"},
	}

	messages := govalidator.MapData{
		"value": []string{helpers.Required(lang), helpers.Min(lang, "2"), helpers.Max(lang, "255")},
		"name":  []string{helpers.Required(lang), helpers.Min(lang, "2"), helpers.Max(lang, "50")},
		"setting_type":  []string{helpers.Required(lang), helpers.Min(lang, "2"), helpers.Max(lang, "50")},
		"slug":  []string{helpers.Required(lang), helpers.Min(lang, "2"), helpers.Max(lang, "50")},
	}

	opts := govalidator.Options{
		Request:         r,     // request object
		Rules:           rules, // rules map
		Data:            request,
		Messages:        messages, // custom message map (Optional)
		RequiredDefault: true, // all the field to be pass the rules
	}
	return govalidator.New(opts)
}
