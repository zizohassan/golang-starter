package test

import (
	"github.com/stretchr/testify/assert"
	"golang-starter/models"
	"testing"
	"time"
)

/***
* check if user not send any required data
 */
func TestResetWithoutEmail(t *testing.T)  {
	data := models.User{
	}
	w := post(data , "reset" , true)
	assert.Equal(t, 400, w.Code)
}


/***
* first register new user
* send valid request to reset email
*/
func TestResetValidRequest(t *testing.T)  {
	data := models.User{
		Email:"zizo199988@gmail.com",
		Password:"123457",
		Name:"Abdel Aziz Hassan",
	}
	w := post(data , "register" , true)
	assert.Equal(t, 200, w.Code)
	resetRequest := models.Reset{
		Email:"zizo199988@gmail.com",
	}
	r := post(resetRequest , "reset" , false)
	time.Sleep(100000)
	assert.Equal(t, 200, r.Code)
}
