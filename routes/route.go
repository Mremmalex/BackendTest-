package routes

import (
	"net/http"
)

//Init is define the userRoute url routes
func Init() {
	http.HandleFunc("/", Index)
    http.HandleFunc("/event", Events)
    http.HandleFunc("/event/region", GetEventByRegion)
    http.HandleFunc("/event/details", EventDetails)
	http.HandleFunc("/auth/signup", UserSignUp)
	http.HandleFunc("/auth/login", UserSignIn)
	http.HandleFunc("/auth/addfriend", SendFriendRequest)
}


