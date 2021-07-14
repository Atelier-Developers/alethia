package interfaces

import (
	"fmt"
	"github.com/Atelier-Developers/alethia/domain/entity"
	"github.com/Atelier-Developers/alethia/domain/repository"
	"github.com/Atelier-Developers/alethia/infrastructure/security"
	"github.com/Atelier-Developers/alethia/interfaces/bodyTemplates"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserHandler struct {
	userRepository repository.UserRepository
	//authInterface  auth.AuthInterface
	//tokenInterface auth.TokenInterface
}

// TODO: AUTH AND TOKEN
func NewUserHandler(userRepository repository.UserRepository) UserHandler {
	return UserHandler{
		userRepository: userRepository,
		//authInterface:  authInterface,
		//tokenInterface: tokenInterface,
	}
}

func (userHandler *UserHandler) SaveUser(c *gin.Context) {
	var userRequestBody bodyTemplates.UserRequestBody
	fmt.Println(c.Request.Body)
	if err := c.ShouldBindJSON(&userRequestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	fmt.Println(userRequestBody)
	fmt.Println(userRequestBody.JoinDate)

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

	fmt.Println(user)

	if time.Time.IsZero(user.JoinDate) {
		user.JoinDate = time.Now()
	}

	fmt.Println(user)

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
