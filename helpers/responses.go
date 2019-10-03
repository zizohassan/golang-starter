package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

/**
* conflict
*/
func ReturnFoundRow(g *gin.Context, msg string) {
	var errors map[string]string
	var data map[string]interface{}
	response(g, msg, data, errors, http.StatusConflict, 409, false)
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
				"message": "something not valid in your request",
				"errors":  e,
				"code":    400,
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
func OkResponse(g *gin.Context, msg string, data map[string]interface{}) {
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
		"data":    data,
	})
	return
}

