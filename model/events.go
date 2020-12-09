package model

import (
	"database/sql"
	"log"
)

func CreateEventTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS Events (eventID INT AUTO_INCREMENT, EventName CHAR(78), EventDesc CHAR(255), EventDate CHAR(78), EventTime CHAR(78), PRIMARY KEY (eventntID))")
    if err != nil {
        log.Panic(err.Error())
    }
}
