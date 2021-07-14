package interfaces

import (
	"fmt"
	"github.com/Atelier-Developers/alethia/domain/entity"
	"github.com/Atelier-Developers/alethia/domain/repository"
	"github.com/Atelier-Developers/alethia/infrastructure/security"
	"github.com/gin-gonic/gin"
	"net/http"
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
	var user entity.User
	fmt.Println(c.Request.Body)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	fmt.Println(user)

	var err error
	var userPassword []byte

	userPassword, err = security.Hash(user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	user.Password = string(userPassword[:])
	err = userHandler.userRepository.SaveUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, nil)
}
