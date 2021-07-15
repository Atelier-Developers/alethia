package interfaces

import (
	"fmt"
	"github.com/Atelier-Developers/alethia/domain/entity"
	"github.com/Atelier-Developers/alethia/domain/repository"
	"github.com/Atelier-Developers/alethia/infrastructure/auth"
	"github.com/Atelier-Developers/alethia/infrastructure/security"
	"github.com/Atelier-Developers/alethia/interfaces/bodyTemplates"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserHandler struct {
	userRepository repository.UserRepository
	authInterface  auth.AuthInterface
	tokenInterface auth.TokenInterface
}

func NewUserHandler(userRepository repository.UserRepository, authInterface auth.AuthInterface, tokenInterface auth.TokenInterface) UserHandler {
	return UserHandler{
		userRepository: userRepository,
		authInterface:  authInterface,
		tokenInterface: tokenInterface,
	}
}

func (userHandler *UserHandler) SaveUser(c *gin.Context) {
	var userRequestBody bodyTemplates.UserRequestBody
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

