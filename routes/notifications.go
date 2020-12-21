package routes

import (
	"encoding/json"
	"net/http"
)

func SendFriendRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("content-type", "application/json")
		// param := r.Header.Get("Token")
		// token, err := middlewares.VerifyToken(param)
		// if err != nil {
		// 	log.Panic(err.Error())
		// }
		// if token.Valid {
		// 		}
		response := Jsonresponse{"this is a secure routes"}
		respData, _ := json.Marshal(response)
		w.Write(respData)

	}
}
