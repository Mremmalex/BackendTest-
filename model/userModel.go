package model

import (
	"database/sql"
	"log"
)

func CreateUserTable() {
	db, _ := Dbcon()
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (userID INT AUTO_INCREMENT, FirstName CHAR(74), LastName CHAR(74), Username CHAR(74) ,Email CHAR(90), Password CHAR(225), DateCreated TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,PRIMARY KEY(userID) )")
	if err != nil {
		log.Panic(err.Error())
	}
}
func CreateAddFriendTable() {
	db, _ := Dbcon()
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS friendlist (tableID INT AUTO_INCREMENT, UserToAdd INT , UserAdding INT, Accepted CHAR(29) DEFAULT 'false' ,PRIMARY KEY(tableID))")
	if err != nil {
		log.Panic(err.Error())
	}
}

func InsertUser() (result *sql.Stmt, err error) {
	db, err := Dbcon()
	result, err = db.Prepare("INSERT INTO users(FirstName, LastName, Username, Email, Password) VALUES (?,?,?,?,?)")
	return result, err
}
func SelectOneUser(username string, email string) (result *sql.Row, err error) {
	db, err := Dbcon()
	result = db.QueryRow("SELECT Username, Email, Password FROM users WHERE Username=? AND email=?", username, email)
	return result, err
}

func SelectOneUserByUsername(username string) (result *sql.Row, err error) {
	db, err := Dbcon()
	result = db.QueryRow("SELECT userID, Username,  Email FROM users WHERE Username=?", username)
	return result, err
}

//this is the query for adding friends
func AddFriend() (result *sql.Stmt, err error) {
	db, _ := Dbcon()
	result, err = db.Prepare("INSERT INTO friendlist(UserToAdd , UserAdding) VALUES(?,?) ")
	return result, err
}

func GetFriendRequest(userid int) (*sql.Rows, error) {
	db, _ := Dbcon()
	result, err := db.Query("SELECT * FROM friendlist WHERE UserToAdd=?", userid)
	return result, err
}
func GetUserById(userid int) (*sql.Row, error) {
	db, err := Dbcon()
	result := db.QueryRow("SELECT Username FROM users WHERE userID=?", userid)
	return result, err
}

func AcceptFriendRequest() (result *sql.Stmt, err error) {
	db, err := Dbcon()
	result, err = db.Prepare("UPDATE friendlist SET Accepted=? WHERE UserToAdd=?")
	return result, err
}
