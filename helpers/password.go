package helpers

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
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
func GenerateToken(stringToHash string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	token, _ := HashPassword(stringToHash + RandomString(10))
	return token, nil
}


/**
* generate random string
 */
func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
