package interfaces

import (
	"github.com/Atelier-Developers/alethia/domain/entity"
	"github.com/Atelier-Developers/alethia/domain/entity/notification"
	"github.com/Atelier-Developers/alethia/domain/repository"
	"github.com/Atelier-Developers/alethia/interfaces/bodyTemplates"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type UserBackgroundHistoryHandler struct {
	backgroundHistoryRepository repository.BackgroundHistoryRepository
	notificationRepository      repository.NotificationRepository
	friendRepository            repository.FriendRepository
}

func NewUserBackgroundHistoryHandler(backgroundHistoryRepository repository.BackgroundHistoryRepository, notificationRepository repository.NotificationRepository, friendRepository repository.FriendRepository) UserBackgroundHistoryHandler {
	return UserBackgroundHistoryHandler{
		backgroundHistoryRepository: backgroundHistoryRepository,
		notificationRepository:      notificationRepository,
		friendRepository:            friendRepository,
	}
}

func (userBackgroundHistoryHandler *UserBackgroundHistoryHandler) AddBackgroundHistory(c *gin.Context) {
	var userRequestBody bodyTemplates.UserBackgroundHistoryCreateRequestBody
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

	backgroundHistory := entity.BackgroundHistory{
		UserID:      userId.(uint64),
		StartDate:   userRequestBody.StartDate,
		EndDate:     userRequestBody.EndDate,
		Category:    userRequestBody.Category,
		Title:       userRequestBody.Title,
		Description: userRequestBody.Description,
		Location:    userRequestBody.Location,
	}

	err := userBackgroundHistoryHandler.backgroundHistoryRepository.SaveBackgroundHistory(&backgroundHistory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	friends, err := userBackgroundHistoryHandler.friendRepository.GetFriends(backgroundHistory.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	for _, f := range friends {
		var recieverId uint64
		if f.UserId1 == backgroundHistory.UserID {
			recieverId = f.UserId2
		} else {
			recieverId = f.UserId1
		}
		n := notification.ChangeWork{
			UserHistoryId: backgroundHistory.ID,
			Type:          "add",
			Notification: notification.Notification{
				ReceiverId: recieverId,
			},
		}
		err = userBackgroundHistoryHandler.notificationRepository.CreateChangeWorkNotification(&n)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	}
	c.JSON(http.StatusCreated, nil)
}

func (userBackgroundHistoryHandler *UserBackgroundHistoryHandler) EditBackgroundHistory(c *gin.Context) {
	var userRequestBody bodyTemplates.UserBackgroundHistoryUpdateRequestBody
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
	bh, err := userBackgroundHistoryHandler.backgroundHistoryRepository.GetBackgroundHistory(userRequestBody.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	//TODO TEST this
	if bh.UserID != userId.(uint64) {
		c.JSON(http.StatusForbidden, err)
		return
	}
	backgroundHistory := entity.BackgroundHistory{
		ID:          userRequestBody.ID,
		StartDate:   userRequestBody.StartDate,
		EndDate:     userRequestBody.EndDate,
		Category:    userRequestBody.Category,
		Title:       userRequestBody.Title,
		Description: userRequestBody.Description,
		Location:    userRequestBody.Location,
	}

	err = userBackgroundHistoryHandler.backgroundHistoryRepository.UpdateBackgroundHistory(&backgroundHistory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	friends, err := userBackgroundHistoryHandler.friendRepository.GetFriends(backgroundHistory.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	for _, f := range friends {
		var recieverId uint64
		if f.UserId1 == backgroundHistory.UserID {
			recieverId = f.UserId2
		} else {
			recieverId = f.UserId1
		}
		n := notification.ChangeWork{
			UserHistoryId: backgroundHistory.ID,
			Type:          "update",
			Notification: notification.Notification{
				ReceiverId: recieverId,
			},
		}
		err = userBackgroundHistoryHandler.notificationRepository.CreateChangeWorkNotification(&n)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	}

	c.JSON(http.StatusCreated, nil)
}

func (userBackgroundHistoryHandler *UserBackgroundHistoryHandler) DeleteBackgroundHistory(c *gin.Context) {
	var userRequestBody bodyTemplates.UserBackgroundHistoryDeleteRequestBody
	if err := c.ShouldBindJSON(&userRequestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	backgroundId, err := strconv.ParseInt(c.Param("background_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}
	userRequestBody.ID = uint64(backgroundId)

	_, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	backgroundHistory := entity.BackgroundHistory{
		ID: userRequestBody.ID,
	}

	err = userBackgroundHistoryHandler.backgroundHistoryRepository.DeleteBackgroundHistory(&backgroundHistory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (userBackgroundHistoryHandler *UserBackgroundHistoryHandler) GetUserBackgroundHistories(c *gin.Context) {
	_, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	uId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	backgrounds, err := userBackgroundHistoryHandler.backgroundHistoryRepository.GetUserBackgroundHistory(uint64(uId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, backgrounds)
}
