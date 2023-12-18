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

//can you reuse this function in prev function?
func IsAutheticated (user string)bool{
	//traverse through the map
  for _,value:= range sessionMap{

    //check if present value is equals to userValue
    if(value.Username == user){
      //if same return true
      return true
    }
  }

  //if value not found return false
  return false
}


func GetUserName(){
	if IsAutheticated(user){
		return sessionMap
	}
}
