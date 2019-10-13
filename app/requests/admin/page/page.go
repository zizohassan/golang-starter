package page

import (
	"github.com/thedevsaddam/govalidator"
	"golang-starter/app/models"
	"golang-starter/helpers"
	"net/http"
)

/**
* validate store category request
 */
func StoreUpdate(r *http.Request, request *models.Page) *govalidator.Validator {
	lang := helpers.GetCurrentLangFromHttp(r)
	/// Validation rules
	rules := govalidator.MapData{
		"name":   []string{"required", "min:6", "max:50"},
		"status": []string{"required", "between:1,2"},
	}

	 messages  := govalidator.MapData{
		"name":   []string{helpers.Required(lang), helpers.Min(lang , "6"), helpers.Max(lang , "50")},
		"status": []string{helpers.Required(lang) , helpers.Between(lang , "1,2")},
	}

	opts := govalidator.Options{
		Request:         r,     // request object
		Rules:           rules, // rules map
		Data:            request,
		Messages:        messages, // custom message map (Optional)
		RequiredDefault: true,     // all the field to be pass the rules
	}
	return govalidator.New(opts)
}
