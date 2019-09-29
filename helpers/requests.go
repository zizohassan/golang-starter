package helpers

import (
	"encoding/json"
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
