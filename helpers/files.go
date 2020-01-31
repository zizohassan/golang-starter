package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"io/ioutil"
	"log"
	"net/http"
)

/**
* return with []strings that contains
* files names inside path what you set
 */
func ReadAllFiles(root string) []string {
	var filesList []string
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		filesList = append(filesList, f.Name())
	}
	return filesList
}

/**
valid csv files
*/

func ValidCsvFile(g *gin.Context) string {
	err := CheckFile(g.Request)
	if ReturnNotValidRequestFile(err, g) {
		return ""
	}
	file, _ := g.FormFile("file")
	fileName := "public/files/" + file.Filename
	_ = g.SaveUploadedFile(file, fileName)
	return fileName
}

func CheckFile(r *http.Request) *govalidator.Validator {
	lang := GetCurrentLangFromHttp(r)
	rules := govalidator.MapData{
		"file:file": []string{"ext:csv", "required"},
	}
	messages := govalidator.MapData{
		"file:file": []string{NotValidExt(lang), Required(lang)},
	}
	opts := govalidator.Options{
		Request:  r,     // request object
		Rules:    rules, // rules map,
		Messages: messages,
	}
	return govalidator.New(opts)
}
