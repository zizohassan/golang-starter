package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang-starter/config"
	"golang-starter/models"
	"testing"
)

/**
* register with valid data and
* login with this data
*/
func TestLoginWithValidData(t *testing.T)  {
	data := models.User{
		Email:"zizo19999@gmail.com",
		Password:"123457",
		Name:"Abdel Aziz Hassan",
	}
	w := post(data , "register" , true)
	assert.Equal(t, 200, w.Code)
	l := post(data , "login" , false)
	assert.Equal(t, 200, l.Code)
}

/**
* register with valid data
* login with not valid data
*/
func TestLoginWithNotValidEmail(t *testing.T)  {
	data := models.User{
		Email:"zizo19999@gmail.com",
		Password:"123457",
		Name:"Abdel Aziz Hassan",
	}
	w := post(data , "register" , true)
	assert.Equal(t, 200, w.Code)
	loginData := models.User{
		Email:"zizo@gmail.com",
		Password:"123457",
	}
	l := post(loginData , "login" , false)
	assert.Equal(t, 404, l.Code)
}


/**
* register with valid data
* login with not valid data
 */
func TestLoginWithNotValidPassword(t *testing.T)  {
	data := models.User{
		Email:"zizo19999@gmail.com",
		Password:"123457",
		Name:"Abdel Aziz Hassan",
	}
	w := post(data , "register" , true)
	assert.Equal(t, 200, w.Code)
	loginData := models.User{
		Email:"zizo19999@gmail.com",
		Password:"12345745",
	}
	l := post(loginData , "login" , false)
	fmt.Println(l)
	assert.Equal(t, 404, l.Code)
}

/**
* register with valid data
* login with block user
 */
func TestLoginWithBlockUser(t *testing.T)  {
	data := models.User{
		Email:"zizo19999@gmail.com",
		Password:"123457",
		Name:"Abdel Aziz Hassan",
	}
	w := post(data , "register" , true)
	assert.Equal(t, 200, w.Code)
	config.DB.Model(&models.User{}).Where("email = ? " , data.Email).Update("block" , 1)
	loginData := models.User{
		Email:"zizo19999@gmail.com",
		Password:"123457",
	}
	l := post(loginData , "login" , false)
	assert.Equal(t, 403, l.Code)
}

/**
* check input length
 */
func TestLoginWithMoreThan50Email(t *testing.T)  {
	data := models.User{
		Name:"Abdel Aziz Hassan",
		Email:randomString(50)+"@gmail.com",
		Password:"123457",
	}
	w := post(data , "login" , true)
	assert.Equal(t, 400, w.Code)
}

func TestLoginWithLessThan7Email(t *testing.T)  {
	data := models.User{
		Email:"m@m.i",
		Password:"123457",
	}
	w := post(data , "login" , true)
	assert.Equal(t, 400, w.Code)
}

func TestLoginWithMoreThan20Password(t *testing.T)  {
	data := models.User{
		Name:"Abdel Aziz hassan",
		Email:"zizo19999@gmail.com",
		Password:randomString(30),
	}
	w := post(data , "login" , true)
	assert.Equal(t, 400, w.Code)
}

func TestLoginWithLessThan4Password(t *testing.T)  {
	data := models.User{
		Name:"Abdel Aziz hassan",
		Email:"zizo19999@gmail.com",
		Password:randomString(2),
	}
	w := post(data , "login" , true)
	assert.Equal(t, 400, w.Code)
}