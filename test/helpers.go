package test

import (
	"github.com/stretchr/testify/assert"
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
