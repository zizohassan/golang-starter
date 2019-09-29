package visitor

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"golang-starter/models"
)

/**
* validate login request
*/
func Login(r *http.Request , login *models.Login) *govalidator.Validator {
	/**
	*  Validation rules
	*/
	rules := govalidator.MapData{
		"email":    []string{"required", "min:4", "max:50", "email"},
		"password": []string{"required", "between:6,20"},
	}
	opts := govalidator.Options{
		Request: r,     // request object
		Rules:   rules, // rules map
		Data:    login,
		RequiredDefault: true, // all the field to be pass the rules
	}
	return govalidator.New(opts)
}

/**
* validate register request
 */
func Register(r *http.Request , user *models.User) *govalidator.Validator {
	/**
	*  Validation rules
	 */
	rules := govalidator.MapData{
		"email":    []string{"required", "min:4", "max:50", "email"},
		"name":     []string{"required", "min:4", "max:50"},
		"password": []string{"required", "between:6,20"},
	}
	opts := govalidator.Options{
		Request: r,     // request object
		Rules:   rules, // rules map
		Data:    user,
		RequiredDefault: true, // all the field to be pass the rules
	}
	return govalidator.New(opts)
}
