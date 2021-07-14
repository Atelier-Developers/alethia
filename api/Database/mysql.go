package Database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)


type MySQLDB struct {
	db *sql.DB
}


func (dbClient *MySQLDB) Init() error {
	DbUsername := os.Getenv("DB_USERNAME")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbName := os.Getenv("DB_NAME")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DbUsername, DbPassword, DbHost, DbPort, DbName)
	var err error
	dbClient.db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = dbClient.db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print(fmt.Sprintf("Connected to Mysql or port %s", DbPort))
	return nil
}

func (dbClient *MySQLDB) Close() error {
	return dbClient.db.Close()
}

func (dbClient *MySQLDB) GetDB() *sql.DB {
	return dbClient.db
}
