package model

import (
	"database/sql"
	"log"

	//importing sql driver as blank
	_ "github.com/go-sql-driver/mysql"
)

//DB is a global sql.DB instance
var db *sql.DB

//Main initilization and configoration sql database
func Init() {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/")
		if err != nil {
		    log.Panic(err.Error())
		}
    defer db.Close()

    _, err = db.Exec("CREATE DATABASE IF NOT EXISTS backendtest")
    if err  != nil {
        log.Panic(err)
    }
    useDB(db, "backendtest")
    CreateUserTable(db) 
}

func createDB(dbname string) {
    _, err := db.Exec("CREATE DATABASE IF NOT EXITS" +dbname+ " ") 
    if err  != nil {
        panic(err)
    }
}
func useDB(db *sql.DB, dbname string) {
    _, err := db.Exec("USE " + dbname + " ")
    if err != nil {
        log.Panic(err)
    }
}
