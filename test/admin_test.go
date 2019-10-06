package test

import (
	"github.com/stretchr/testify/assert"
	"golang-starter/app/models"
	"golang-starter/config"
	"gopkg.in/thedevsaddam/gojsonq.v2"
	"testing"
)

/**
* register new user
* return token as header
*/
func getTokenAsHeader(t *testing.T , migrate bool) map[string]string {
	token := addAdminUser(t ,migrate)
	var authToken = make(map[string]string)
	authToken["Authorization"] = token
	return authToken
}


/***
* register new user
* update this user to admin access
* return with admin token
*/
func addAdminUser(t *testing.T , migrate bool) string {
	w, _ := registerNewUser(migrate)
	assert.Equal(t, 200, w.Code)
	responseData := responseData(w.Result().Body)
	token := gojsonq.New().JSONString(responseData).Find("data.token")
	var admin models.User
	config.DB.Model(&admin).Where("token = ?", token).Update("role", 2).Find(&admin)
	return admin.Token
}

