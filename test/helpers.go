package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/thedevsaddam/gojsonq.v2"
	"testing"
	_ "testing"
)

func checkPostRequestWithOutHeadersDataIsValid(t *testing.T , data interface{} , url string , migrate bool)  {
		w := postWitOutHeader(data , url , migrate)
		assert.Equal(t, 400, w.Code)
}

func checkPostRequestWithHeadersDataIsValid(t *testing.T , data interface{} , url string , migrate bool , headers map[string]string)  {
	w := post(data , url , migrate ,headers)
	assert.Equal(t, 400, w.Code)
}

func checkPutRequestWithOutHeadersDataIsValid(t *testing.T , data interface{} , url string , migrate bool)  {
	w := putWithOutHeader(data , url , migrate)
	assert.Equal(t, 400, w.Code)
}

func checkPutRequestWithHeadersDataIsValid(t *testing.T , data interface{} , url string , migrate bool , headers map[string]string)  {
	w := put(data , url , migrate ,headers)
	assert.Equal(t, 400, w.Code)
}

func filter(t *testing.T, url string, value interface{}, key string, method string, token map[string]string) {
	fmt.Println("url log", url+"?"+key+"="+fmt.Sprintf("%v", value))
	k := get(url+"?"+key+"="+fmt.Sprintf("%v", value), false, token)
	responseData := responseData(k.Result().Body)
	recoverResponse := gojsonq.New().JSONString(responseData)
	if method == "equal" {
		assert.EqualValues(t, value, recoverResponse.Find("data.records.[0]."+key))
	} else {
		assert.NotEqual(t, value, recoverResponse.Find("data.records.[0]."+key))
	}
	assert.Equal(t, 200, k.Code)
}
