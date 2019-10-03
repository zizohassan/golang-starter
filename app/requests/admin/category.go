package admin

import (
	"github.com/thedevsaddam/govalidator"
	"golang-starter/app/models"
	"net/http"
)

/**
* validate store category request
 */
func StoreUpdate(r *http.Request , category *models.Category) *govalidator.Validator {
	/// Validation rules
	rules := govalidator.MapData{
		"name":    []string{"required", "min:6", "max:50"},
		"status":  []string{"required", "between:1,2"},
	}
	opts := govalidator.Options{
		Request: r,     // request object
		Rules:   rules, // rules map
		Data:    category,
		RequiredDefault: true, // all the field to be pass the rules
	}
	return govalidator.New(opts)
}
