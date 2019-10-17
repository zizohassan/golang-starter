package requests

import (
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"errors"
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
}
