package test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"starter/config"
	"starter/models"
	"starter/providers"
)

/**
* init gin and return gin engine
 */
func setupRouter() *gin.Engine {
	/**
	* drop database
	*/
	config.ConnectToDatabase()
	/***
	* migrate tables
	*/
	models.MigrateAllTable(os.Getenv("TEST_MODEL_PATH"))
	/**
	* start gin
	*/
	r := providers.Gin()
	return providers.Routing(r)
}

/**
* post request
 */
func post(data interface{}, url string) *httptest.ResponseRecorder {
	router := setupRouter()
	w := httptest.NewRecorder()
	sendData, _ := json.Marshal(&data)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(sendData))
	router.ServeHTTP(w, req)
	return w
}

/**
* generate random string
*/
func randomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
