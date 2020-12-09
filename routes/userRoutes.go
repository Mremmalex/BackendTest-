package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//User is a struct that maps to the fileds
//we have in our database
type User struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

//Jsonresponse is a struct for returning
//a json response after a function call
type Jsonresponse struct {
	Message string `json:"message"`
}

//Init is define the userRoute url routes
func Init() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/auth", registerFunc)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		res.Write([]byte("this is the homePage"))
	}
}

func registerFunc(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		res.Write([]byte("this is served from the main page"))
	}
	if req.Method == "POST" {
		res.Header().Set("content-type", "application/json")
		var person User
		reqBody, _ := ioutil.ReadAll(req.Body)
		json.Unmarshal(reqBody, &person)
		fmt.Println(person.FullName)
		response := Jsonresponse{"data recieved successfully"}
		jsonRes, _ := json.Marshal(response)
		res.Write(jsonRes)

	}

}
func createUser(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		res.Write([]byte("this is the function for creating User"))
	}
}
