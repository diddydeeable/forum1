package auth

import (
	"fmt"
	"forumhub/model"
)

var sessionMap map[string]model.User

func init() {
	sessionMap = make(map[string]model.User)
}

func SetCookieToMap(uuid string, user model.User) bool {
	_, exists := sessionMap[uuid]
	if exists {
		return false
	}
sessionMap[uuid]= user
fmt.Println("session created")
return true

}

