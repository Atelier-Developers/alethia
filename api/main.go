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

	inviteRepo := persistance.NewInviteRepository(&db)
	friendRepo := persistance.NewFriendRepository(&db)
	messageRepo := persistance.NewMessageRepository(&db)
	languageRepo := persistance.NewLanguageRepository(&db)
	backgroundHistoryRepo := persistance.NewBackgroundHistoryRepository(&db)
	userRepo := persistance.NewUserRepository(&db)
	postRepo := persistance.NewPostRepository(&db)
	commentRepo := persistance.NewCommentRepository(&db)
	skillRepo := persistance.NewSkillRepository(&db)
	notificationRepo := persistance.NewNotificationRepository(&db)
	conversationRepo := persistance.NewConversationRepository(&db)

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	var error error
	redisService, error := auth.NewRedisDB(redisHost, redisPort, redisPassword, context.Background())
	if error != nil {
		log.Fatal(error)
	}

	token := auth.NewToken()

	notificationHandler := interfaces.NewNotificationHandler(notificationRepo)
	friendHandler := interfaces.NewFriendHandler(friendRepo, inviteRepo)
	messageHandler := interfaces.NewMessageHandler(conversationRepo, messageRepo)
	userBackgroundHistoryHandler := interfaces.NewUserBackgroundHistoryHandler(backgroundHistoryRepo, notificationRepo)
	userLanguageHandler := interfaces.NewUserLanguageHandler(languageRepo)
	userHandler := interfaces.NewUserHandler(userRepo, redisService.Auth, token)
	postHandler := interfaces.NewPostHandler(postRepo, notificationRepo)
	commentHandler := interfaces.NewCommentHandler(commentRepo, notificationRepo)
	skillHandler := interfaces.NewSkillHandler(skillRepo, notificationRepo)
	conversationHandler := interfaces.NewConversationHandler(conversationRepo)

	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	router.POST("/login", userHandler.Login)
	router.POST("/logout", userHandler.Logout)
	router.POST("/register", userHandler.SaveUser)

	router.GET("/language", userLanguageHandler.GetLanguages)
	router.GET("/skill", skillHandler.GetSkills)

	userGroup := router.Group("/users", middleware.AuthMiddleware(redisService.Auth))
	{
		userGroup.PUT("", userHandler.EditUser)
		userGroup.GET("", userHandler.GetUser)
		userGroup.POST("", userHandler.GetUserById)

		inviteGroup := userGroup.Group("/invite")
		{
			inviteGroup.POST("", friendHandler.AddInvite)
			inviteGroup.DELETE("", friendHandler.DeleteInvite)
			inviteGroup.GET("", friendHandler.GetUserInvites)
		}

		friendGroup := userGroup.Group("/friend")
		{
			friendGroup.POST("", friendHandler.AddFriend)
			friendGroup.GET("", friendHandler.GetFriends)
			friendGroup.DELETE("", friendHandler.DeleteFriend)
		}

		languageGroup := userGroup.Group("/language")
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
			backgroundGroup.GET("", userBackgroundHistoryHandler.GetUserBackgroundHistories)
		}

		userSkillGroup := userGroup.Group("/skill")
		{
			userSkillGroup.GET("", skillHandler.GetUserSkills)
			userSkillGroup.POST("", skillHandler.AddUserSkill)
			userSkillGroup.DELETE("", skillHandler.DeleteUserSkill)
			userSkillGroup.POST("/endorse", skillHandler.EndorseSkill)
		}

		userNotificationGroup := userGroup.Group("/notification")
		{
			userNotificationGroup.GET("", notificationHandler.GetUserNotifications)
		}
	}

	conversationGroup := router.Group("/conversation", middleware.AuthMiddleware(redisService.Auth))
	{
		conversationGroup.POST("", conversationHandler.AddConversation)
		conversationGroup.PUT("", conversationHandler.UpdateUserConversation)
		//conversationGroup.DELETE("", conversationHandler.DeleteConversation)
		conversationGroup.GET("", conversationHandler.GetUserConversations)

		messageGroup := conversationGroup.Group("/message")
		{
			messageGroup.POST("", messageHandler.AddMessage)
			messageGroup.GET("", messageHandler.GetMessages)
			// TODO: Naming should be different
			messageGroup.GET("/singleMessage", messageHandler.GetMessage)
		}
		//conversationGroup.GET("/singleConversation", conversationHandler.GetConversation)
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
