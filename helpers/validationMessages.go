package helpers

import "github.com/bykovme/gotrans"

func Required(lang string) string {
	return "required:" + gotrans.Tr(lang, "required")
}

func Min(lang string, number string) string {
	return "min:" + gotrans.Tr(lang, "min") + " " + number
}

func Max(lang string, number string) string {
	return "max:" + gotrans.Tr(lang, "max") + " " + number
}

func Between(lang string, number string) string {
	return "between:" + gotrans.Tr(lang, "between") + " " + number
}
