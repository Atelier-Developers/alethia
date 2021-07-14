package main

import (
	"context"
	"github.com/Atelier-Developers/alethia/Database"
	"github.com/Atelier-Developers/alethia/infrastructure/auth"
	"github.com/Atelier-Developers/alethia/infrastructure/persistance"
	"github.com/Atelier-Developers/alethia/interfaces"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func main() {

	var db Database.MySQLDB
	err := db.Init()

	defer db.Close()

	if err != nil {
		log.Fatalln(err.Error())
	}

	userRepo := persistance.NewUserRepository(&db)
	postRepo := persistance.NewPostRepository(&db)

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	var error error
	redisService, error := auth.NewRedisDB(redisHost, redisPort, redisPassword, context.Background())
	if error != nil {
		log.Fatal(error)
	}

	token := auth.NewToken()

	userHandler := interfaces.NewUserHandler(userRepo, redisService.Auth, token)
	postHandler := interfaces.NewPostHandler(postRepo)

	router := gin.Default()

	router.POST("/login", userHandler.Login)
	userGroup := router.Group("/users")
	{
		userGroup.POST("", userHandler.SaveUser)
	}

	postGroup := router.Group("/post")
	{
		postGroup.POST("", postHandler.SavePost)
	}

	router.Run(":3000")
}
