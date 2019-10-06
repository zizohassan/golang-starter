package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-starter/app/models"
	"golang-starter/config"
	"golang-starter/providers"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
)

/**
* define struct that carry the arg of request
 */
type RequestData struct {
	Migrate     bool
	RequestType string
	Url         string
	Data        interface{}
	Header      map[string]string
}

/**
* init gin and return gin engine
 */
func setupRouter(migrate bool) *gin.Engine {
	// drop database
	config.ConnectToDatabase()
	/***
	* migrate tables new instance control if we must drop all tables
	* or no may be you need to stay the data to check
	 */
	if migrate {
		models.MigrateAllTable(os.Getenv("TEST_MODEL_PATH"))
	}
	/// start gin
	r := providers.Gin()
	return providers.Routing(r)
}

/**
* post request
 */
func post(data interface{}, url string, migrate bool, headers map[string]string) *httptest.ResponseRecorder {
	return request(RequestData{
		Url:         url,
		Migrate:     migrate,
		RequestType: "POST",
		Data:        data,
		Header:      headers,
	})
}

func postWitOutHeader(data interface{}, url string, migrate bool) *httptest.ResponseRecorder {
	return request(RequestData{
		Url:         url,
		Migrate:     migrate,
		RequestType: "POST",
		Data:        data,
	})
}

/**
* Put request
 */
func put(data interface{}, url string, migrate bool, headers map[string]string) *httptest.ResponseRecorder {
	return request(RequestData{
		Url:         url,
		Migrate:     migrate,
		RequestType: "PUT",
		Data:        data,
		Header:      headers,
	})
}

func putWithOutHeader(data interface{}, url string, migrate bool) *httptest.ResponseRecorder {
	return request(RequestData{
		Url:         url,
		Migrate:     migrate,
		RequestType: "PUT",
		Data:        data,
	})
}

/**
* Get request
 */
func get(url string, migrate bool, headers map[string]string) *httptest.ResponseRecorder {
	return request(RequestData{
		Url:         url,
		Migrate:     migrate,
		RequestType: "GET",
		Header:      headers,
	})
}

func getWithOutHeader(url string, migrate bool) *httptest.ResponseRecorder {
	return request(RequestData{
		Url:         url,
		Migrate:     migrate,
		RequestType: "GET",
	})
}

/**
* Get request
 */
func deleter(url string, migrate bool, headers map[string]string) *httptest.ResponseRecorder {
	return request(RequestData{
		Url:         url,
		Migrate:     migrate,
		RequestType: "DELETE",
		Header:      headers,
	})
}

/**
* Make new request
 */
func request(request RequestData) *httptest.ResponseRecorder {
	router := setupRouter(request.Migrate)
	w := httptest.NewRecorder()
	sendData, _ := json.Marshal(&request.Data)
	req, _ := http.NewRequest(request.RequestType, request.Url, bytes.NewReader(sendData))
	if len(request.Header) > 0 {
		for headerName, headerValue := range request.Header {
			req.Header.Set(headerName, headerValue)
		}
	}
	router.ServeHTTP(w, req)
	fmt.Println(w)
	return w
}

/**
* return response as json
 */
func responseData(c io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c)
	return buf.String()
}
