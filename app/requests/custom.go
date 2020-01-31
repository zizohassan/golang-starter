package requests

import (
	"errors"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"new_back_end/app/models"
	"new_back_end/config"
	"reflect"
	"strings"
)

func Init() {
	/**
	* this role check if slice of strings is not empty
	 */
	govalidator.AddCustomRule("strings_slice", func(field string, rule string, message string, value interface{}) error {
		err := fmt.Errorf("The %s field is required", field)
		fmt.Println(message)
		if message != "" {
			err = errors.New(message)
		}
		if value == nil {
			return err
		}
		if len(value.([]string)) == 0 {
			if message != "" {
				return errors.New(message)
			}
		}
		return nil
	})

	/**
	* this role check if slice of int is not empty
	 */
	govalidator.AddCustomRule("int_slice", func(field string, rule string, message string, value interface{}) error {
		err := fmt.Errorf("The %s field is required", field)
		if message != "" {
			err = errors.New(message)
		}
		if value == nil {
			return err
		}
		if len(value.([]int)) == 0 {
			if message != "" {
				return errors.New(message)
			}
		}
		return nil
	})

	/**
	* this role check if slice of int is not empty
	 */
	govalidator.AddCustomRule("permission_slice", func(field string, rule string, message string, value interface{}) error {
		err := fmt.Errorf("The %s field is required", field)
		if message != "" {
			err = errors.New(message)
		}
		if value == nil {
			return err
		}
		permissions := value.([]models.PermissionForm)
		if len(permissions) == 0 {
			return err
		}
		return nil
	})

	govalidator.AddCustomRule("int_array_slice", func(field string, rule string, message string, value interface{}) error {
		err := fmt.Errorf("The %s field is required", field)
		if message != "" {
			err = errors.New(message)
		}
		if value == nil {
			return err
		}
		var i []int
		if reflect.TypeOf(value) == reflect.TypeOf(i) {
			i = value.([]int)
			if len(i) == 0 {
				return err
			}
		}
		return nil
	})

	govalidator.AddCustomRule("unique", func(field string, rule string, message string, value interface{}) error {
		err := fmt.Errorf("The %s field is unique", field)
		table := strings.Split(rule, ":")
		fmt.Println(message)
		if message != "" {
			err = errors.New(message)
		}
		if value == nil {
			return err
		}
		var count int
		config.DB_READ.Table(table[1]).Where(field+" = ?", value.(string)).Count(&count)
		if count != 0 {
			return err
		}
		return nil
	})

	govalidator.AddCustomRule("unique_update", func(field string, rule string, message string, value interface{}) error {
		err := fmt.Errorf("The %s field is unique", field)
		rules := strings.Split(rule, ":")
		table := strings.Split(rules[1], ",")
		fmt.Println(message)
		if message != "" {
			err = errors.New(message)
		}
		if value == nil {
			return err
		}
		var count int
		config.DB_READ.Table(table[0]).Where(field+" = ?", value.(string)).Where("id != ?", table[1]).Count(&count)
		if count != 0 {
			return err
		}
		return nil
	})
}
