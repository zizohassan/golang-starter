package test

import (
	"bytes"
	"encoding/json"
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
* init gin and return gin engine
 */
func setupRouter(migrate bool) *gin.Engine {
	/**
	* drop database
	 */
	config.ConnectToDatabase()
	/***
	* migrate tables new instance control if we must drop all tables
	* or no may be you need to stay the data to check
	 */
	if migrate {
		models.MigrateAllTable(os.Getenv("TEST_MODEL_PATH"))
	}
	/**
	* start gin
	 */
	r := providers.Gin()
	return providers.Routing(r)
}

/**
* post request
 */
func post(data interface{}, url string, migrate bool) *httptest.ResponseRecorder {
	router := setupRouter(migrate)
	w := httptest.NewRecorder()
	sendData, _ := json.Marshal(&data)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(sendData))
	router.ServeHTTP(w, req)
	return w
}

/**
* return response as json
 */
func responseData(c io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c)
	return   buf.String()
}
