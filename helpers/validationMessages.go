package helpers

import "github.com/bykovme/gotrans"

func Required(lang string) string {
	return "required:" + gotrans.Tr(lang, "required")
}

func StringsSlice(lang string) string {
	return "strings_slice:" + gotrans.Tr(lang, "strings_slice")
}

func IntSlice(lang string) string {
	return "int_slice:" + gotrans.Tr(lang, "strings_slice")
}

func Email(lang string) string {
	return "email:" + gotrans.Tr(lang, "email_not_valid")
}

func Min(lang string, number string) string {
	return "min:" + gotrans.Tr(lang, "min") + " " + number
}

func In(lang string, ableStrings ...string) string {
	returnString := "( "
	for i := 0; i < len(ableStrings)-1; i++ {
		returnString += ableStrings[i] + ", "
	}
	returnString += ableStrings[len(ableStrings)-1] + " )"
	return "in:" + gotrans.Tr(lang, "in") + " " + returnString
}

func Ext(lang string, extentions string) string {
	return "ext:" + gotrans.Tr(lang, "ext") + " " + extentions
}

func Mime(lang string, extentions string) string {
	return "mime:" + gotrans.Tr(lang, "ext") + " " + extentions
}

func Size(lang string, size string) string {
	return "size:" + gotrans.Tr(lang, "size") + " " + size
}

func Numeric(lang string) string {
	return "numeric:" + gotrans.Tr(lang, "numeric")
}

func Digits(lang string) string {
	return "digits:" + gotrans.Tr(lang, "numeric")
}

func Url(lang string) string {
	return "url:" + gotrans.Tr(lang, "url")
}

func Bool(lang string) string {
	return "boolean:" + gotrans.Tr(lang, "boolean")
}

func Max(lang string, number string) string {
	return "max:" + gotrans.Tr(lang, "max") + " " + number
}

func Unique(lang string, key... string) string {
	s := ""
	for _, k := range key {
		s += gotrans.Tr(lang, k)
	}
	return s
}

func Between(lang string, number string) string {
	return "between:" + gotrans.Tr(lang, "between") + " " + number
}
func NotValidExt(lang string) string {
	return "ext:" + gotrans.Tr(lang, "error_read_file")
}