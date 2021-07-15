package main

import (
	"context"
	"github.com/Atelier-Developers/alethia/Database"
	"github.com/Atelier-Developers/alethia/infrastructure/auth"
	"github.com/Atelier-Developers/alethia/infrastructure/persistance"
	"github.com/Atelier-Developers/alethia/interfaces"
	"github.com/Atelier-Developers/alethia/middleware"
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

	languageRepo := persistance.NewLanguageRepository(&db)
	backgroundHistoryRepo := persistance.NewBackgroundHistoryRepository(&db)
	userRepo := persistance.NewUserRepository(&db)
	postRepo := persistance.NewPostRepository(&db)
	commentRepo := persistance.NewCommentRepository(&db)
	skillRepo := persistance.NewSkillRepository(&db)

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	var error error
	redisService, error := auth.NewRedisDB(redisHost, redisPort, redisPassword, context.Background())
	if error != nil {
		log.Fatal(error)
	}

	token := auth.NewToken()

	userBackgroundHistoryHandler := interfaces.NewUserBackgroundHistoryHandler(backgroundHistoryRepo, redisService.Auth, token)
	userLanguageHandler := interfaces.NewUserLanguageHandler(languageRepo, redisService.Auth, token)
	userHandler := interfaces.NewUserHandler(userRepo, redisService.Auth, token)
	postHandler := interfaces.NewPostHandler(postRepo)
	commentHandler := interfaces.NewCommentHandler(commentRepo)
	skillHandler := interfaces.NewSkillHandler(skillRepo)

	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	router.POST("/login", userHandler.Login)
	router.POST("/logout", userHandler.Logout)
	router.POST("/register", userHandler.SaveUser)

	userGroup := router.Group("/users", middleware.AuthMiddleware(redisService.Auth))
	{
		userGroup.PUT("", userHandler.EditUser)

		languageGroup := userGroup.Group("/language", middleware.AuthMiddleware(redisService.Auth))
		{
			languageGroup.POST("", userLanguageHandler.AddUserLanguage)
			languageGroup.DELETE("", userLanguageHandler.DeleteUserLanguage)
			languageGroup.GET("", userLanguageHandler.GetUserLanguages)
		}
		backgroundGroup := userGroup.Group("/background")
		{
			backgroundGroup.POST("", userBackgroundHistoryHandler.AddBackgroundHistory)
			backgroundGroup.PUT("", userBackgroundHistoryHandler.EditBackgroundHistory)
			backgroundGroup.DELETE("", userBackgroundHistoryHandler.DeleteBackgroundHistory)
		}

		userSkillGroup := userGroup.Group("/skill")
		{
			userSkillGroup.GET("", skillHandler.GetUserSkills)
			userSkillGroup.POST("", skillHandler.AddUserSkill)
			userSkillGroup.DELETE("", skillHandler.DeleteUserSkill)
			userSkillGroup.POST("/endorse", skillHandler.EndorseSkill)
		}
	}

	postGroup := router.Group("/post", middleware.AuthMiddleware(redisService.Auth))
	{
		postGroup.POST("", postHandler.SavePost)
		postGroup.POST("/like", postHandler.LikePost)
		commentGroup := postGroup.Group("/comment")
		{
			commentGroup.POST("", commentHandler.SaveComment)
			commentGroup.POST("/reply", commentHandler.ReplyComment)
			//TODO multiple like not allowed
			commentGroup.POST("/like", commentHandler.LikeComment)
		}
	}

	router.Run(":3000")
}
