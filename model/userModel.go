package model

import (
	"database/sql"
	"log"
)

type User struct {
    PubID uint `json:"pubID"`
    FirstName string `json:"firstname"`
    LastName string `json:"ladtname"`
    Email string `json:"email"` 
    Password string `json:"password"`
}

func CreateUserTable(db *sql.DB) {
    _, err := db.Exec("CREATE TABLE IF NOT EXISTS user (userID INT AUTO_INCREMENT, FirstName CHAR(74), LastName CHAR(74), Email CHAR(90), Password CHAR(225), DateCreated TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,PRIMARY KEY(userID) )")
    if err != nil  {
        log.Panic(err.Error())
    }
}
