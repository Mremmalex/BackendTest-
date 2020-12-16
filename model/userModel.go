package model

import (
	"database/sql"
	"log"
)

func CreateUserTable() {
    db , _ := Dbcon()
    _, err := db.Exec("CREATE TABLE IF NOT EXISTS users (userID INT AUTO_INCREMENT, FirstName CHAR(74), LastName CHAR(74), Username CHAR(74) ,Email CHAR(90), Password CHAR(225), DateCreated TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,PRIMARY KEY(userID) )")
    if err != nil  {
        log.Panic(err.Error())
    }
}


func InsertUser() (result *sql.Stmt, err error){
    db, err := Dbcon()
    result, err = db.Prepare("INSERT INTO users(FirstName, LastName, Username, Email, Password) VALUES (?,?,?,?,?)")
    return result, err
}
func SelectOneUser(username string, email string) (result *sql.Row, err error) {
    db, err := Dbcon()
    result = db.QueryRow("SELECT Username, Email, Password FROM users WHERE Username=? AND email=?" , username, email)
    return result, err
}
