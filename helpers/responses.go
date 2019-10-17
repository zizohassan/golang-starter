package helpers

import (
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

/**
* conflict
 */
func ReturnBadRequest(g *gin.Context) {
	var errors map[string]string
	var data map[string]interface{}
	var msg = gotrans.Tr(GetCurrentLang(g), "400")
	response(g, msg, data, errors, 400, 400, false)
	return
}

/**
* Duplicate data
 */
func ReturnDuplicateData(g *gin.Context, inputName string) {
	var errors map[string]string
	var data map[string]interface{}
	var msg = T(g, "duplicate_data_part_one", inputName, "duplicate_data_part_two")
	response(g, msg, data, errors, 409, 409, false)
	return
}

/**
* upload error
 */
func UploadError(g *gin.Context) {
	var errors map[string]string
	var data map[string]interface{}
	var msg = T(g, "upload_error_code")
	response(g, msg, data, errors, 415, 415, false)
	return
}

/**
* multi upload error
 */
func MultiUploadError(g *gin.Context) {
	var errors map[string]string
	var data map[string]interface{}
	var msg = T(g, "upload_multi_images_error_code")
	response(g, msg, data, errors, 415, 415, false)
	return
}

/**
* NotValidRequest response
 */
func ReturnNotValidRequest(error *govalidator.Validator, g *gin.Context) bool {
	e := error.ValidateJSON()
	if len(e) > 0 {
		g.JSON(
			http.StatusBadRequest, gin.H{
				"status":  false,
				"message": gotrans.Tr(GetCurrentLang(g), "400"),
				"errors":  e,
				"code":    400,
				"payload": nil,
			})
		return true
	}
	return false
}

/**
* NotValidRequest response
 */
func ReturnNotValidRequestFormData(error *govalidator.Validator, g *gin.Context) bool {
	e := error.Validate()
	if len(e) > 0 {
		g.JSON(
			http.StatusBadRequest, gin.H{
				"status":  false,
				"message": gotrans.Tr(GetCurrentLang(g), "400"),
				"errors":  e,
				"code":    400,
				"payload": nil,
			})
		return true
	}
	return false
}

/**
* NotFound response
 */
func ReturnNotFound(g *gin.Context, msg string) {
	var errors map[string]string
	var data map[string]interface{}
	response(g, msg, data, errors, http.StatusNotFound, 404, false)
	return
}

/**
* Forbidden response
 */
func ReturnForbidden(g *gin.Context, msg string) {
	var errors map[string]string
	var data map[string]interface{}
	response(g, msg, data, errors, http.StatusForbidden, 403, false)
	return
}

/**
* ok response with data
 */
func OkResponse(g *gin.Context, msg string, data interface{}) {
	var errors map[string]string
	response(g, msg, data, errors, http.StatusOK, 200, true)
	return
}

/**
* ok response without data
 */
func OkResponseWithOutData(g *gin.Context, msg string) {
	var errors map[string]string
	var data map[string]interface{}
	response(g, msg, data, errors, http.StatusOK, 200, true)
	return
}

/**
* Not Authorize
 */
func ReturnYouAreNotAuthorize(g *gin.Context) {
	var errors map[string]string
	var data map[string]interface{}
	var msg = gotrans.Tr(GetCurrentLang(g), "401")
	response(g, msg, data, errors, 401, 401, true)
	return
}

/**
* ok with paging
 */
func OkResponseWithPaging(g *gin.Context, msg string, data *Paginator) {
	var errors map[string]string
	response(g, msg, data, errors, http.StatusOK, 200, true)
	return
}

/**
* stander response
 */
func response(g *gin.Context, msg string, data interface{}, errors map[string]string, httpStatus int, code int, status bool) {
	g.JSON(httpStatus, gin.H{
		"status":  status,
		"message": msg,
		"errors":  errors,
		"code":    code,
		"payload": data,
	})
	return
}
