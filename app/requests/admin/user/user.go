package user

import (
	"github.com/thedevsaddam/govalidator"
	"golang-starter/app/models"
	"golang-starter/helpers"
	"net/http"
)

/**
* validate store user request
 */
func Store(r *http.Request, request *models.User) *govalidator.Validator {
	lang := helpers.GetCurrentLangFromHttp(r)
	/// Validation rules
	rules := govalidator.MapData{
		"name":     []string{"required", "min:6", "max:50"},
		"email":    []string{"required", "min:6", "max:50", "email"},
		"password": []string{"required", "min:6", "max:50"},
		"status":    []string{"required"},
		"role":     []string{"required", "between:1,2"},
	}

	messages := govalidator.MapData{
		"name":     []string{helpers.Required(lang), helpers.Min(lang, "6"), helpers.Max(lang, "50")},
		"email":    []string{helpers.Required(lang), helpers.Min(lang, "6"), helpers.Max(lang, "50"), helpers.Email(lang)},
		"password": []string{helpers.Required(lang), helpers.Min(lang, "6"), helpers.Max(lang, "50")},
		"status":    []string{helpers.Required(lang)},
		"role":     []string{helpers.Required(lang), helpers.Min(lang, "1"), helpers.Max(lang, "2")},
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

/**
* validate update user request
 */
func Update(r *http.Request, request *models.User) *govalidator.Validator {
	lang := helpers.GetCurrentLangFromHttp(r)
	/// Validation rules
	rules := govalidator.MapData{
		"name":     []string{"required", "min:6", "max:50"},
		"email":    []string{"required", "min:6", "max:50", "email"},
		"password": []string{"max:50"},
		"status":    []string{"required"},
		"role":     []string{"required", "between:1,2"},
	}

	messages := govalidator.MapData{
		"name":     []string{helpers.Required(lang), helpers.Min(lang, "6"), helpers.Max(lang, "50")},
		"email":    []string{helpers.Required(lang), helpers.Min(lang, "6"), helpers.Max(lang, "50"), helpers.Email(lang)},
		"password": []string{helpers.Max(lang, "50")},
		"status":    []string{helpers.Required(lang)},
		"role":     []string{helpers.Required(lang), helpers.Min(lang, "1"), helpers.Max(lang, "2")},
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
