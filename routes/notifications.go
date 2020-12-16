package routes

import (
	"fmt"
	"net/http"
)


func SendFriendRequest(w http.ResponseWriter, r *http.Request)  {
    // param := r.URL.Query().Get("Token") 
    param := r.Header.Get("Authorization")   
    fmt.Println(param)
   if r.Method == "GET" {
        w.Header().Set("content-type", "application/json")    
        w.Write([]byte("this is a secure routes"))
    } 
    w.Write([]byte("this is a protected route"))
}
