package visitor

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"starter/models"
)

func Login(r *http.Request , login models.Login) *govalidator.Validator {
	/**
	*  Validation rules
	 */
	rules := govalidator.MapData{
		"email":    []string{"required", "min:4", "max:20", "email"},
		"password": []string{"required", "between:6,20"},
	}
	opts := govalidator.Options{
		Request: r,     // request object
		Rules:   rules, // rules map
		Data:    &login,
		//Messages:        messages, // custom message map (Optional)
		RequiredDefault: true, // all the field to be pass the rules
	}
	return govalidator.New(opts)
}
