package providers

import (
	"github.com/gin-gonic/gin"
)

/**
* start gin framework
* add cross origin middleware
*/
func Gin() *gin.Engine {
	/// init gin
	r := gin.Default()
	/**
	* Run Default middleware
	* this means all requests will go throw this middleware
	*/
	middlewares(r)

	return r
}
