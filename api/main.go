package main

import (
	"github.com/Atelier-Developers/alethia/Database"
	"github.com/Atelier-Developers/alethia/infrastructure/persistance"
	"github.com/Atelier-Developers/alethia/interfaces"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
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

	//redis_host := os.Getenv("REDIS_HOST")
	//redis_port := os.Getenv("REDIS_PORT")
	//redis_password := os.Getenv("REDIS_PASSWORD")
	//
	//redisService, err := auth.NewRedisDB(redis_host, redis_port, redis_password)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//token := auth.NewToken()
	//
	//authenticate := auth.NewAuth(services.User, redisService.Auth, tk)

	userHandler := interfaces.NewUserHandler(userRepo)
	postHandler := interfaces.NewPostHandler(postRepo)

	router := gin.Default()

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
