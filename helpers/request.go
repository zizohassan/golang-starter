package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

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

func OkResponse(g *gin.Context, msg string, data map[string]interface{}, code int) {
	if msg == "" {
		msg = "request is valid "
	}
	if code == 0 {
		code = 200
	}
	var errors map[string]string
	g.JSON(
		http.StatusOK, gin.H{
			"status":  true,
			"message": msg,
			"errors":  errors,
			"code":    code,
			"data":    data,
		})
	return
}
