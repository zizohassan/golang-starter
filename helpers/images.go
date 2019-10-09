package helpers

import (
	"fmt"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"strings"
)

/***
* upload file with input name
* if there are resize width height
* it will resize image with this sizes
 */
func UploadImage(g *gin.Context, fileString string) (bool, string) {
	file, _ := g.FormFile(fileString)
	dis := "public/images/original/"
	fileName := file.Filename
	path := dis + fileName
	err := g.SaveUploadedFile(file, path)
	if err != nil {
		UploadError(g)
		return false, ""
	}
	if g.Query("resize") != "" {
		if g.Query("width") != "" && g.Query("height") != "" {
			resizeImage(path, fileName, g.Query("width"), g.Query("height"))
		}
	}
	return true, path
}

/***
* resize and save in new folder
 */
func resizeImage(image string, fileName string, width string, height string) {
	img, err := imgio.Open(image)
	if err != nil {
		fmt.Println(err)
		return
	}
	folder := width + "X" + height
	path := "public/images/" + folder
	os.MkdirAll(path, os.ModePerm)
	widthInt, _ := strconv.Atoi(width)
	heightInt, _ := strconv.Atoi(height)
	resized := transform.Resize(img, widthInt, heightInt, transform.Linear)
	if strings.Contains(fileName, ".png") || strings.Contains(fileName, ".PNG") {
		err = imgio.Save(path+"/"+fileName, resized, imgio.PNGEncoder())
	}
	if strings.Contains(fileName, ".jpeg") || strings.Contains(fileName, ".JPEG") {
		q , _ := strconv.Atoi(os.Getenv("QUALITY"))
		err = imgio.Save(path+"/"+fileName, resized, imgio.JPEGEncoder(q))
	}
	if strings.Contains(fileName, ".jpg") || strings.Contains(fileName, ".jpg") {
		q , _ := strconv.Atoi(os.Getenv("QUALITY"))
		err = imgio.Save(path+"/"+fileName, resized, imgio.JPEGEncoder(q))
	}
	if strings.Contains(fileName, ".BMP") || strings.Contains(fileName, ".bmp") {
		err = imgio.Save(path+"/"+fileName, resized, imgio.BMPEncoder())
	}
	if err != nil {
		fmt.Println(err)
		return
	}
}
