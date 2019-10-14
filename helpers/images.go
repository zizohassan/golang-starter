package helpers

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
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
	return uploadHandler(g, file)
}

/***
* upload file with input name
* if there are resize width height
* it will resize image with this sizes
 */
func UploadImages(g *gin.Context, fileString string) (bool, []string) {
	var uploads []string
	form, err := g.MultipartForm()
	if err != nil {
		g.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return false, uploads
	}
	files := form.File[fileString]
	for _, file := range files {
		uploadHandler(g, file)
		uploads = append(uploads, file.Filename)
	}
	return true, uploads
}

/***
* upload handler
* take file then upload images
 */
func uploadHandler(g *gin.Context, file *multipart.FileHeader) (bool, string) {
	dis := "public/images/original/"
	fileName := file.Filename
	path := dis + fileName
	err := g.SaveUploadedFile(file, path)
	if err != nil {
		UploadError(g)
		return false, ""
	}
	if g.Query("method") != "" {
		if g.Query("width") != "" && g.Query("height") != "" {
			resizeImage(path, fileName, g.Query("width"), g.Query("height"))
		}
	}
	return true, fileName
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
		q, _ := strconv.Atoi(os.Getenv("QUALITY"))
		err = imgio.Save(path+"/"+fileName, resized, imgio.JPEGEncoder(q))
	}
	if strings.Contains(fileName, ".jpg") || strings.Contains(fileName, ".jpg") {
		q, _ := strconv.Atoi(os.Getenv("QUALITY"))
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

//// base 64 image handel

/**
* convert image to to base 64
 */
func ConvertImageToBase(g *gin.Context, filepath string) (bool, string) {
	rawData, err := getFileContents(filepath)
	if err != nil {
		return false, err.Error()
	}
	b, s := DecodeImage(g, string(rawData))
	return b, s
}

func MultiDecodeImage(g *gin.Context, images []string) []string {
	var uploadedImages []string
	for _, imageToUpload := range images {
		valid, imageName := DecodeImage(g, imageToUpload)
		if valid {
			uploadedImages = append(uploadedImages, imageName)
		}
	}
	return uploadedImages
}

/**
* Decode Image and save it on path
 */
func DecodeImage(g *gin.Context, filename string) (bool, string) {
	encData := string(filename)
	encData = strings.Replace(encData, "\n", "", -1)
	/// check if base64 is valid
	data, err := stripMime(encData)
	if err != nil {
		return false, "image_pattern_error " + err.Error()
	}
	/// decode base 64
	output, err := base64.StdEncoding.DecodeString(data[2])
	if err != nil {
		return false, "image_error_decode " + err.Error()
	}
	/// read image
	r := bytes.NewReader(output)
	var img image.Image
	var imgErorr error
	ext := ".png"
	/// decode image based on extentions
	if strings.Contains(data[1], "png") {
		img, imgErorr = png.Decode(r)
	} else {
		img, imgErorr = jpeg.Decode(r)
		ext = ".jpeg"
	}
	if imgErorr != nil {
		return false, "image_error_on_reader " + imgErorr.Error()
	}
	/// generate random string
	guid := xid.New()
	name := guid.String() + ext
	/// save image
	path := "public/images/original/"
	f, err := os.OpenFile(path+name, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return false, "image_error_save " + err.Error()
	}
	/// write image content based on extentions
	if strings.Contains(data[1], "png") {
		err = png.Encode(f, img)
		if err != nil {
			return false, "image_error_save " + err.Error()
		}
	} else {
		q, _ := strconv.Atoi(os.Getenv("QUALITY"))
		var opt jpeg.Options
		opt.Quality = q
		err = jpeg.Encode(f, img, &opt)
		if err != nil {
			return false, "image_error_save " + err.Error()
		}
	}

	if g.Query("method") != "" {
		if g.Query("width") != "" && g.Query("height") != "" {
			resizeImage(path+name, name, g.Query("width"), g.Query("height"))
		}
	}

	return true, name
}

/**
* convert image to based 64
 */
func getFileContents(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("Error opening file")
	}

	info, err := f.Stat()
	if err != nil {
		return nil, errors.New("Error getting file stats")
	}

	len := info.Size()
	data := make([]byte, len)
	n, err := f.Read(data)
	if err != nil {
		return nil, errors.New("Error reading file")
	}
	if int64(n) != len {
		return nil, errors.New("Could not read entire contents of file")
	}

	return data, nil
}

/**
* check if base 64 is valid
 */
func stripMime(combined string) ([]string, error) {
	re := regexp.MustCompile("data:(.*);base64,(.*)")
	parts := re.FindStringSubmatch(combined)
	if len(parts) < 3 {
		var s []string
		return s, errors.New("Invalid base64 input")
	}

	return parts, nil
}
