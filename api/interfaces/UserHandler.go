package interfaces

import (
	"fmt"
	"github.com/Atelier-Developers/alethia/domain/entity"
	"github.com/Atelier-Developers/alethia/domain/entity/notification"
	"github.com/Atelier-Developers/alethia/domain/repository"
	"github.com/Atelier-Developers/alethia/infrastructure/auth"
	"github.com/Atelier-Developers/alethia/infrastructure/security"
	"github.com/Atelier-Developers/alethia/interfaces/bodyTemplates"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type UserHandler struct {
	conversationRepository repository.ConversationRepository
	inviteRepository       repository.InviteRepository
	friendRepository       repository.FriendRepository
	userRepository         repository.UserRepository
	authInterface          auth.AuthInterface
	tokenInterface         auth.TokenInterface
	notificationRepository repository.NotificationRepository
}

func NewUserHandler(conversationRepository repository.ConversationRepository, inviteRepository repository.InviteRepository, userRepository repository.UserRepository, friendRepository repository.FriendRepository, authInterface auth.AuthInterface, tokenInterface auth.TokenInterface, notificationRepository repository.NotificationRepository) UserHandler {
	return UserHandler{
		conversationRepository: conversationRepository,
		inviteRepository:       inviteRepository,
		userRepository:         userRepository,
		friendRepository:       friendRepository,
		authInterface:          authInterface,
		tokenInterface:         tokenInterface,
		notificationRepository: notificationRepository,
	}
}

func (userHandler *UserHandler) SaveUser(c *gin.Context) {
	var userRequestBody bodyTemplates.UserCreateRequestBody
	if err := c.ShouldBindJSON(&userRequestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	user := entity.User{
		Username:        userRequestBody.Username,
		FirstName:       userRequestBody.FirstName,
		LastName:        userRequestBody.LastName,
		Intro:           userRequestBody.Intro,
		About:           userRequestBody.About,
		Accomplishments: userRequestBody.Accomplishments,
		AdditionalInfo:  userRequestBody.AdditionalInfo,
		JoinDate:        userRequestBody.JoinDate,
		BirthDate:       userRequestBody.BirthDate,
		Location:        userRequestBody.Location,
	}

	if time.Time.IsZero(user.JoinDate) {
		user.JoinDate = time.Now()
	}
	var err error

	user.Password, err = security.Hash(userRequestBody.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var tmp entity.User
	err = userHandler.userRepository.GetUserByUsername(user.Username, &tmp)
	if err == nil {
		c.JSON(http.StatusConflict, err)
		return
	}

	err = userHandler.userRepository.SaveUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func (userHandler *UserHandler) GetUser(c *gin.Context) {
	userId, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	var user entity.UserWithMutualConnectionAndFriendshipStatus
	err := userHandler.userRepository.GetUserByID(userId.(uint64), userId.(uint64), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	response := bodyTemplates.UserGetResponseBody{
		UserID:          userId.(uint64),
		Username:        user.Username,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Intro:           user.Intro,
		About:           user.About,
		Accomplishments: user.Accomplishments,
		AdditionalInfo:  user.AdditionalInfo,
		BirthDate:       user.BirthDate,
		Location:        user.Location,
		JoinDate:        user.JoinDate,
	}
	c.JSON(http.StatusCreated, response)
}

func (userHandler *UserHandler) ViewProfile(c *gin.Context) {

	var userViewProfileRequestBody bodyTemplates.UserViewProfileRequestBody
	if err := c.ShouldBindJSON(&userViewProfileRequestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	userId, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}
	n := notification.ViewProfile{
		UserId: userId.(uint64),
		Notification: notification.Notification{
			ReceiverId: userViewProfileRequestBody.Id,
		},
	}
	err := userHandler.notificationRepository.CreateViewProfileNotification(&n)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (userHandler *UserHandler) GetUsersWithMutualConnection(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		log.Fatal("User Id does not exist!")
	}

	users, err := userHandler.userRepository.GetUsersWithMutualConnection(userId.(uint64))
	if err != nil {
		log.Fatal(err)
	}

	var responseUsers []bodyTemplates.UserGetMutualConnectionsResponseBody
	for _, u := range users {
		responseUser := bodyTemplates.UserGetMutualConnectionsResponseBody{
			UserID:           u.ID,
			Username:         u.Username,
			FirstName:        u.FirstName,
			LastName:         u.LastName,
			Intro:            u.Intro,
			About:            u.About,
			Accomplishments:  u.Accomplishments,
			AdditionalInfo:   u.AdditionalInfo,
			BirthDate:        u.BirthDate,
			Location:         u.Location,
			JoinDate:         u.JoinDate,
			MutualConnection: uint64(u.MutualConnection.Int64),
		}

		responseUsers = append(responseUsers, responseUser)
	}

	c.JSON(http.StatusOK, responseUsers)
}

func (userHandler *UserHandler) GetUsersWithSimilarUsername(c *gin.Context) {
	var userRequestBody bodyTemplates.UsersGetByUsernameRequestBody
	if err := c.ShouldBindJSON(&userRequestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	username := userRequestBody.Username

	userId, exists := c.Get("user_id")
	if !exists {
		log.Fatal("User Id does not exist!")
	}

	users, err := userHandler.userRepository.GetUsersWithSimilarUsername(userId.(uint64), string(username))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var responseUsers []bodyTemplates.UsersGetByUsernameResponseBody
	for _, u := range users {
		responseUser := bodyTemplates.UsersGetByUsernameResponseBody{
			UserID:                u.ID,
			Username:              u.Username,
			FirstName:             u.FirstName,
			LastName:              u.LastName,
			Intro:                 u.Intro,
			About:                 u.About,
			Accomplishments:       u.Accomplishments,
			AdditionalInfo:        u.AdditionalInfo,
			BirthDate:             u.BirthDate,
			Location:              u.Location,
			JoinDate:              u.JoinDate,
			MutualConnection:      uint64(u.MutualConnection.Int64),
			IsFriendsWithThisUser: u.IsFriendsWithThisUser,
		}

		responseUsers = append(responseUsers, responseUser)
	}

	if userRequestBody.Location != "" {
		var tmp []bodyTemplates.UsersGetByUsernameResponseBody
		for _, u := range responseUsers {
			if u.Location == userRequestBody.Location {
				tmp = append(tmp, u)
			}
		}
		responseUsers = tmp
	}
	if userRequestBody.Company != "" {
		var tmp []bodyTemplates.UsersGetByUsernameResponseBody
		for _, u := range responseUsers {
			result, err := userHandler.userRepository.DoesUserWorkAtCompany(u.UserID, userRequestBody.Company)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				return
			}
			if result {
				tmp = append(tmp, u)
			}
		}
		responseUsers = tmp
	}
	if userRequestBody.Language != "" {
		var tmp []bodyTemplates.UsersGetByUsernameResponseBody
		for _, u := range responseUsers {
			result, err := userHandler.userRepository.DoesUserHaveLanguage(u.UserID, userRequestBody.Language)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				return
			}
			if result {
				tmp = append(tmp, u)
			}
		}
		responseUsers = tmp
	}

	c.JSON(http.StatusOK, responseUsers)
}

func (userHandler *UserHandler) GetUserById(c *gin.Context) {
	var userRequestBody bodyTemplates.UserGetRequestBody
	if err := c.ShouldBindJSON(&userRequestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		log.Fatal("User Id does not exist!")
	}

	var u entity.UserWithMutualConnectionAndFriendshipStatus
	err := userHandler.userRepository.GetUserByID(userId.(uint64), userRequestBody.Id, &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	responseUser := bodyTemplates.UserGetByIdResponseBody{
		UserID:                u.ID,
		Username:              u.Username,
		FirstName:             u.FirstName,
		LastName:              u.LastName,
		Intro:                 u.Intro,
		About:                 u.About,
		Accomplishments:       u.Accomplishments,
		AdditionalInfo:        u.AdditionalInfo,
		BirthDate:             u.BirthDate,
		Location:              u.Location,
		JoinDate:              u.JoinDate,
		MutualConnection:      uint64(u.MutualConnection.Int64),
		IsFriendsWithThisUser: u.IsFriendsWithThisUser,
	}

	isInvited, err := userHandler.inviteRepository.IsUserInvitedById(userId.(uint64), userRequestBody.Id)
	if err != nil {
		log.Fatal(err)
	}

	responseUser.IsInvitedByThisUser = isInvited

	conversationExists, err := userHandler.conversationRepository.ConversationExists(userId.(uint64), userRequestBody.Id)
	if err != nil {
		log.Fatal(err)
	}

	responseUser.HasConversationWithThisUser = conversationExists

	c.JSON(http.StatusOK, responseUser)
}

func (userHandler *UserHandler) Login(c *gin.Context) {
	var userRequestBody *bodyTemplates.UserLoginRequestBody
	var tokenErr = map[string]string{}

	if err := c.ShouldBindJSON(&userRequestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	var user entity.User
	userErr := userHandler.userRepository.GetUserByUsernameAndPassword(userRequestBody.Username, userRequestBody.Password, &user)

	if userErr != nil {
		c.JSON(http.StatusBadRequest, userErr)
		return
	}
	ts, tErr := userHandler.tokenInterface.CreateToken(user.ID)
	if tErr != nil {
		tokenErr["token_error"] = tErr.Error()
		c.JSON(http.StatusUnprocessableEntity, tErr.Error())
		return
	}
	saveErr := userHandler.authInterface.CreateAuth(user.ID, ts)
	if saveErr != nil {
		c.JSON(http.StatusInternalServerError, saveErr.Error())
		return
	}
	userData := make(map[string]interface{})
	userData["access_token"] = ts.AccessToken
	userData["refresh_token"] = ts.RefreshToken
	userData["id"] = user.ID
	userData["first_name"] = user.FirstName
	userData["last_name"] = user.LastName

	c.JSON(http.StatusOK, userData)
}

func (userHandler *UserHandler) Logout(c *gin.Context) {
	//check is the user is authenticated first
	metadata, err := userHandler.tokenInterface.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
	//if the access token exist and it is still valid, then delete both the access token and the refresh token
	deleteErr := userHandler.authInterface.DeleteTokens(metadata)
	if deleteErr != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, deleteErr.Error())
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}

func (userHandler *UserHandler) EditUser(c *gin.Context) {
	var userRequestBody bodyTemplates.UserUpdateRequestBody
	if err := c.ShouldBindJSON(&userRequestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	userId, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	user := entity.User{
		ID:              userId.(uint64),
		Intro:           userRequestBody.Intro,
		About:           userRequestBody.About,
		Accomplishments: userRequestBody.Accomplishments,
		AdditionalInfo:  userRequestBody.AdditionalInfo,
		BirthDate:       userRequestBody.BirthDate,
		Location:        userRequestBody.Location,
	}

	err := userHandler.userRepository.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, nil)
}
