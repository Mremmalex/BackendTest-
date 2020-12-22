package routes

import (
	"backEndTest/middlewares"
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
	http.HandleFunc("/auth/addfriend", middlewares.IsAuthorised(SendFriendRequest))
	http.HandleFunc("/auth/notification", middlewares.IsAuthorised(Notification))
	http.HandleFunc("/auth/accept", middlewares.IsAuthorised(AcceptFriendRequest))
}
