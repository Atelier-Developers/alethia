package main

import (
	"github.com/Atelier-Developers/alethia/Database"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	err := Database.Init()
	if err != nil {
		log.Fatalln(err.Error())
	}

	db := Database.GetDB()
	defer db.Close()

	router := gin.Default()

	router.Run(":3000")
}
