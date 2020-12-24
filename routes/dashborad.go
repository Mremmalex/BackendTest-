package routes

import (
	"backEndTest/middlewares"
	"backEndTest/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Notify struct {
	Message  string `json:"message"`
	Username string `json:"username"`
	Accepted string `json:"accepted"`
}

func Notification(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		param := r.Header.Get("Token")
		username, _ := middlewares.DecodeToken(param)
		result, err := model.SelectOneUserByUsername(username)
		if err != nil {
			log.Panic(err.Error())
		}
		var current User
		err = result.Scan(&current.UserID, &current.Username, &current.Email)
		if err != nil {
			log.Panic(err.Error())
		}
		var friendlist FriendList
		query, err := model.GetFriendRequest(current.UserID)
		if err != nil {
			log.Panic(err.Error())
		}

		for query.Next() {
			err := query.Scan(&friendlist.TableID, &friendlist.UserToAdd, &friendlist.UserAdding, &friendlist.Accepted)
			if err != nil {
				log.Panic(err.Error())
			}
			fmt.Println(friendlist)

			var UserAdding User
			result, _ := model.GetUserById(friendlist.UserAdding)
			err = result.Scan(&UserAdding.Username)
			if friendlist.Accepted == "false" {
				notify := &Notify{
					Message:  "You Have A New Friend Request",
					Username: UserAdding.Username,
					Accepted: friendlist.Accepted,
				}
				response, _ := json.Marshal(notify)
				w.Write(response)
			} else {

				response := Jsonresponse{"You Have No Request At the Momment"}
				respData, _ := json.Marshal(response)
				w.Write(respData)

			}
		}
	}
}

// CreateNewEvent(w http.ResponseWriter, r *http.Request){
// 	if r.Method == "POST"{
// 		w.Header().Set("content-type", "application/json")
// 	}
// }
