package interfaces

import (
	"github.com/Atelier-Developers/alethia/domain/repository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type NotificationHandler struct {
	notificationRepository repository.NotificationRepository
}

func NewNotificationHandler(notificationRepository repository.NotificationRepository) NotificationHandler {
	return NotificationHandler{
		notificationRepository: notificationRepository,
	}
}

func (notificationHandler *NotificationHandler) GetUserNotifications(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		log.Fatal("User Id does not exist!")
	}
	result := map[string]interface{}{}
	var res interface{}
	res, err := notificationHandler.notificationRepository.GetCommentNotification(userId.(uint64))
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		result["comment"] = res
	}
	c.JSON(http.StatusOK, result)
}
