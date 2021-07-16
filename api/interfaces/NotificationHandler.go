package interfaces

import (
	"github.com/Atelier-Developers/alethia/domain/entity/notification"
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

	res, err := notificationHandler.notificationRepository.GetBirthdayNotification(userId.(uint64))
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		res = []notification.Birthday{}
	}
	result["birthday"] = res

	res, err = notificationHandler.notificationRepository.GetChangeWorkNotification(userId.(uint64))
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		res = []notification.ChangeWork{}
	}
	result["change_work"] = res

	res, err = notificationHandler.notificationRepository.GetCommentNotification(userId.(uint64))
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		res = []notification.Comment{}
	}
	result["comment"] = res

	res, err = notificationHandler.notificationRepository.GetEndorseSkillNotification(userId.(uint64))
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		res = []notification.EndorseSkill{}
	}
	result["endorse"] = res

	res, err = notificationHandler.notificationRepository.GetLikeCommentNotification(userId.(uint64))
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		res = []notification.LikeComment{}
	}
	result["like_comment"] = res

	res, err = notificationHandler.notificationRepository.GetLikePostNotification(userId.(uint64))
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		res = []notification.LikePost{}
	}
	result["like_post"] = res

	res, err = notificationHandler.notificationRepository.GetReplyCommentNotification(userId.(uint64))
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		res = []notification.LikeComment{}
	}
	result["reply_comment"] = res

	res, err = notificationHandler.notificationRepository.GetViewProfileNotification(userId.(uint64))
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		res = []notification.ViewProfile{}
	}
	result["view_profile"] = res

	c.JSON(http.StatusOK, result)
}
