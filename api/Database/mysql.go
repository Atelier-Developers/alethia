package Database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var db *sql.DB

func Init() error {
	DbUsername := os.Getenv("DB_USERNAME")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbName := os.Getenv("DB_NAME")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DbUsername, DbPassword, DbHost, DbPort, DbName)
	var err error
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print(fmt.Sprintf("Connected to Mysql or port %s", DbPort))
	return nil
}

func GetDB() *sql.DB {
	return db
}
