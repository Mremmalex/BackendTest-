package routes

import (
	"backEndTest/middlewares"
	"backEndTest/model"
	"backEndTest/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//User is a struct that maps to the fileds
//we have in our database
type User struct {
	UserID    int    `json:"userid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

//Token is a global toke variable
var Token string

//Jsonwponse is a struct for returning
//a json wponse after a function call
type Jsonresponse struct {
	Message string `json:"message"`
}

//UserSignUp handles the user registration process
//it also validates user input and saves the data to the DATABASE
func UserSignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("this is served from the main page"))
	}
	if r.Method == "POST" {
		w.Header().Set("content-type", "application/json")
		stmt, err := model.InsertUser()
		if err != nil {
			log.Panic(err.Error())
		}
		var user User
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &user)
		hashpwd, err := utils.HashMyPassword(user.Password)
		if err != nil {
			log.Panic(err.Error())
		}
		if !utils.CheckEmail(user.Email) {
			response := Jsonresponse{"Email Address is not Valid"}
			respData, _ := json.Marshal(response)
			w.Write(respData)
		} else if !utils.CheckLenOfPassword(user.Password) {
			response := Jsonresponse{"please provide a strong password"}
			respData, _ := json.Marshal(response)
			w.Write(respData)
		} else {
			_, err = stmt.Exec(user.FirstName, user.LastName, user.Username, user.Email, hashpwd)
			rponse := Jsonresponse{"Account Created successfully"}
			jsonRes, _ := json.Marshal(rponse)
			w.Write(jsonRes)
		}
	}

}

func UserSignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("content-type", "application/json")
		var user User
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &user)
		result, err := model.SelectOneUser(user.Username, user.Email)
		if err != nil {
			log.Panic(err.Error())
		}

		userr := User{}
		err = result.Scan(&userr.Username, &userr.Email, &userr.Password)
		if err != nil {
			log.Panic(err.Error())
		}
		isValidPassword := utils.CheckHashPassword(user.Password, userr.Password)
		if isValidPassword == false {
			response := Jsonresponse{"Login Credential Incorrect"}
			respData, _ := json.Marshal(response)
			w.Write(respData)
		} else {
			token, err := middlewares.CreateJwtToken(userr.Username)
			if err != nil {
				log.Panic(err.Error())
			}
			response := Jsonresponse{token}
			respData, _ := json.Marshal(response)
			w.Write(respData)
		}
	}
}
