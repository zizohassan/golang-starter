package setting

import (
	"github.com/thedevsaddam/govalidator"
	"golang-starter/app/models"
	"net/http"
)

/**
* validate store setting request
 */
func StoreUpdate(r *http.Request, request *models.Setting) *govalidator.Validator {
	/// Validation rules
	rules := govalidator.MapData{
		"value":        []string{"required", "min:2", "max:255"},
		"name":         []string{"required", "min:2", "max:50"},
		"setting_type": []string{"required", "min:2", "max:50"},
		"slug":         []string{"required", "min:2", "max:50"},
	}
	opts := govalidator.Options{
		Request:         r,     // request object
		Rules:           rules, // rules map
		Data:            request,
		RequiredDefault: true, // all the field to be pass the rules
	}
	return govalidator.New(opts)
}
