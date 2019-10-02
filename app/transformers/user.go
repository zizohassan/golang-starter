package transformers

import "golang-starter/app/models"

/**
* stander the single user response
 */
func UserResponse(user models.User) map[string]interface{} {
	var u = make(map[string]interface{})
	u["name"] = user.Name
	u["email"] = user.Email
	u["role"] = user.Role
	u["token"] = user.Token

	return u
}

/**
* stander the Multi users response
 */
func UsersResponse(users []models.User) []map[string]interface{} {
	var u  = make([]map[string]interface{} , 0)
	for _ , user := range users {
		u = append(u , UserResponse(user))
	}
	return u

}

