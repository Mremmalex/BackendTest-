package model

import (
	"log"
)

func CreateUserTable() {
    db , _ := Dbcon()
    _, err := db.Exec("CREATE TABLE IF NOT EXISTS users (userID INT AUTO_INCREMENT, FirstName CHAR(74), LastName CHAR(74), Email CHAR(90), Password CHAR(225), DateCreated TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,PRIMARY KEY(userID) )")
    if err != nil  {
        log.Panic(err.Error())
    }
}
