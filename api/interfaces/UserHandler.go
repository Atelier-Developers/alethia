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

// TODO: AUTH AND TOKEN
func NewUserHandler(userRepository repository.UserRepository, authInterface  auth.AuthInterface, tokenInterface auth.TokenInterface) UserHandler {
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
	////validate request:
	//validateUser := user.Validate("login")
	//if len(validateUser) > 0 {
	//	c.JSON(http.StatusUnprocessableEntity, validateUser)
	//	return
	//}
	u, userErr := userHandler.userRepository.GetUserByUsernameAndPassword(userRequestBody.Username, userRequestBody.Password)

	fmt.Println(u)

	if userErr != nil {
		c.JSON(http.StatusInternalServerError, userErr)
		return
	}
	ts, tErr := userHandler.tokenInterface.CreateToken(u.ID)
	if tErr != nil {
		tokenErr["token_error"] = tErr.Error()
		c.JSON(http.StatusUnprocessableEntity, tErr.Error())
		return
	}
	saveErr := userHandler.authInterface.CreateAuth(u.ID, ts)
	if saveErr != nil {
		c.JSON(http.StatusInternalServerError, saveErr.Error())
		return
	}
	userData := make(map[string]interface{})
	userData["access_token"] = ts.AccessToken
	userData["refresh_token"] = ts.RefreshToken
	userData["id"] = u.ID
	userData["first_name"] = u.FirstName
	userData["last_name"] = u.LastName

	c.JSON(http.StatusOK, userData)
}
