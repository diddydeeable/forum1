package auth

import (
	"fmt"
	"forumhub/model"
)

var SessionMap map[string]model.User

func init() {
	SessionMap = make(map[string]model.User)
}

func SetCookieToMap(uuid string, user model.User) bool {
	_, exists := SessionMap[uuid]
	if exists {
		return false
	}
SessionMap[uuid]= user
fmt.Println("session created")
return true

}

//can you reuse this function in prev function?
func IsAuthenticated(user string) bool {
	for _, value := range SessionMap {
		if value.Username == user {
			return true
		}
	}
	return false
}


func GetUserName(name string)(model.User, bool){
	user, exists := SessionMap[name]
	return user, exists
}
