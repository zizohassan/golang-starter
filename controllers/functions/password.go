package functions

import (
	"golang.org/x/crypto/bcrypt"
	"investment-users/helpers"
	"math/rand"
	"golang-starter/models"
	"time"
)

/**
* hash passwords
*/
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 3)
	return string(bytes), err
}

/**
* check if password is valid
 */
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

/**
* generate token based on user data
*/
func GenerateToken(user models.User) (string, error) {
	rand.Seed(time.Now().UnixNano())
	token, _ := helpers.HashPassword(user.Password + user.Name + helpers.RandomString(10))
	return token, nil
}
