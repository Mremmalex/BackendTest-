package model

import (
	"database/sql"
	"log"

	//importing sql driver as blank
	_ "github.com/go-sql-driver/mysql"
)

//Dbcon is a global sql.DB instance
func Dbcon() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "mremmalex:password@tcp(localhost:3306)/backendtest")
	return db, err
}

//Main initilization and configoration sql database
func Init() {
	createDB("backendtest")
	useDB("backendtest")
	CreateUserTable()
	CreateEventTable()
}

func createDB(dbname string) {
	db, _ := Dbcon()
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname + " ")
	if err != nil {
		log.Panic(err.Error())
	}
}

//useDB tells sql package the database to focus on
func useDB(dbname string) {
	db, _ := Dbcon()
	_, err := db.Exec("USE " + dbname)
	if err != nil {
		log.Panic(err.Error())
	}
}
