package helpers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
* decode body then return interface
*/
func DecodeAndReturn(req  *http.Request, structBind interface{}) (interface{} , error)  {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&structBind)
	if err != nil {
		return structBind  , err
	}
	return structBind , nil
}

/**
* get lang header
*/
func LangHeader(g *gin.Context) string  {
	 s := ""
	 s = g.Request.Header.Get("Accept-Language")
	return s
}
