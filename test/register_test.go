package test

import (
	"golang-starter/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

/***
* check if user not send any required data
*/
func TestRegisterWithoutEmail(t *testing.T)  {
	data := models.User{
		Password:"123457",
		Name:"Abdel Aziz Hassan",
	}
	w := post(data , "register" , true)
	assert.Equal(t, 400, w.Code)
}

func TestRegisterWithoutPassword(t *testing.T)  {
	data := models.User{
		Email:"zizo1999988@gmail.com",
		Name:"Abdel Aziz Hassan",
	}
	w := post(data , "register" , true)
	assert.Equal(t, 400, w.Code)
}

func TestRegisterWithoutName(t *testing.T)  {
	data := models.User{
		Email:"zizo1999988@gmail.com",
		Password:"123457",
	}
	w := post(data , "register" , true)
	assert.Equal(t, 400, w.Code)
}

/***
* test if user send not valid email it will return bad request
*/
func TestRegisterWithWrongEmailContest(t *testing.T)  {
	data := models.User{
		Email:"Abdel Aziz Hassan",
		Password:"123457",
		Name:"Abdel Aziz Hassan",
	}
	w := post(data , "register" , true)
	assert.Equal(t, 400, w.Code)
}

/**
* check input length
*/
func TestRegisterWithMoreThan50Email(t *testing.T)  {
	data := models.User{
		Name:"Abdel Aziz Hassan",
		Email:randomString(50)+"@gmail.com",
		Password:"123457",
	}
	w := post(data , "register" , true)
	assert.Equal(t, 400, w.Code)
}

func TestRegisterWithMoreThan50Name(t *testing.T)  {
	data := models.User{
		Name:randomString(70),
		Email:"zizo19999@gmail.com",
		Password:"123457",
	}
	w := post(data , "register" , true)
	assert.Equal(t, 400, w.Code)
}

func TestRegisterWithLessThan7Email(t *testing.T)  {
	data := models.User{
		Email:"m@m.i",
		Password:"123457",
	}
	w := post(data , "register" , true)
	assert.Equal(t, 400, w.Code)
}

func TestRegisterWithMoreThan20Password(t *testing.T)  {
	data := models.User{
		Name:"Abdel Aziz hassan",
		Email:"zizo19999@gmail.com",
		Password:randomString(30),
	}
	w := post(data , "register" , true)
	assert.Equal(t, 400, w.Code)
}

func TestRegisterWithLessThan4Password(t *testing.T)  {
	data := models.User{
		Name:"Abdel Aziz hassan",
		Email:"zizo19999@gmail.com",
		Password:randomString(2),
	}
	w := post(data , "register" , true)
	assert.Equal(t, 400, w.Code)
}

/**
* check if user has register with email before
 */

func TestRegisterWithValidCase(t *testing.T)  {
	data := models.User{
		Email:"zizo199988@gmail.com",
		Password:"123457",
		Name:"Abdel Aziz Hassan",
	}
	w := post(data , "register" , true)
	assert.Equal(t, 200, w.Code)
}

/**
* check if user has register with email before
*/
func TestRegisterWithExistEmail(t *testing.T)  {
	data := models.User{
		Email:"zizo199988@gmail.com",
		Password:"123457",
		Name:"Abdel Aziz Hassan",
	}
	w := post(data , "register" , true)
	k := post(data , "register" , false)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 409, k.Code)
}
