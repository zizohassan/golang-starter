package translation

import (
	"github.com/thedevsaddam/govalidator"
	"golang-starter/app/models"
	"net/http"
)

/**
* validate store category request
 */
func StoreUpdate(r *http.Request , request *models.Translation) *govalidator.Validator {
	/// Validation rules
	rules := govalidator.MapData{
		"value":    []string{"required", "min:2", "max:255"},
		"lang":  []string{"required", "min:2", "max:10"},
		"page":  []string{"required", "min:2", "max:30"},
		"slug":  []string{"required", "min:2", "max:50"},
	}
	opts := govalidator.Options{
		Request: r,     // request object
		Rules:   rules, // rules map
		Data:    request,
		RequiredDefault: true, // all the field to be pass the rules
	}
	return govalidator.New(opts)
}
