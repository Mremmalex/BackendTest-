package model

import (
	"database/sql"
	"log"
)

func CreateEventTable() {
    db, _ := Dbcon()
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS events (EventID INT AUTO_INCREMENT, EventName CHAR(78), EventDesc CHAR(255), EventDate CHAR(78), EventTime CHAR(78),EventRegion CHAR(78) ,PRIMARY KEY(EventID))")
    if err != nil {
        log.Panic(err.Error())
    }
}


func InsertEvent() (result *sql.Stmt, err error) {
    db, err := Dbcon()
    result,err = db.Prepare("INSERT INTO events (EventName, EventDesc, EventDate,EventTime, EventRegion) VALUES (?, ?, ?, ?, ?)")
    return result, err 
}

func GetAllEvent()(result *sql.Rows, err error ) {
        db, err := Dbcon()
        result, err = db.Query("SELECT * FROM events")
        return result, err 
}
