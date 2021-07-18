package interfaces

import (
	"github.com/Atelier-Developers/alethia/domain/entity/notification"
	"github.com/Atelier-Developers/alethia/domain/repository"
	"github.com/Atelier-Developers/alethia/interfaces/bodyTemplates"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type SkillHandler struct {
	skillRepository        repository.SkillRepository
	notificationRepository repository.NotificationRepository
}

func NewSkillHandler(skillRepository repository.SkillRepository, notificationRepository repository.NotificationRepository) SkillHandler {
	return SkillHandler{
		skillRepository:        skillRepository,
		notificationRepository: notificationRepository,
	}
}

func (skillHandler *SkillHandler) EndorseSkill(c *gin.Context) {
	var endorseSkillRequestBody bodyTemplates.EndorseSkillRequestBody
	if err := c.ShouldBindJSON(&endorseSkillRequestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	userId, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	if endorseSkillRequestBody.UserId == userId.(uint64) {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	err := skillHandler.skillRepository.EndorseSkill(endorseSkillRequestBody.SkillId, endorseSkillRequestBody.UserId, userId.(uint64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	n := notification.EndorseSkill{
		UserId:           userId.(uint64),
		UserSkillUserId:  endorseSkillRequestBody.UserId,
		UserSkillSkillId: endorseSkillRequestBody.SkillId,
		Notification: notification.Notification{
			ReceiverId: endorseSkillRequestBody.UserId,
		},
	}
	err = skillHandler.notificationRepository.CreateEndorseSkillNotification(&n)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (skillHandler *SkillHandler) GetSkills(c *gin.Context) {
	skills, err := skillHandler.skillRepository.GetSkills()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, skills)
}

func (skillHandler *SkillHandler) AddUserSkill(c *gin.Context) {
	var addUserSkillRequestBody bodyTemplates.AddUserSkillRequestBody
	if err := c.ShouldBindJSON(&addUserSkillRequestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	userId, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	exists, err := skillHandler.skillRepository.UserSkillExist(addUserSkillRequestBody.SkillId, userId.(uint64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if exists {
		c.JSON(http.StatusConflict, err)
		return
	}

	err = skillHandler.skillRepository.AddUserSkill(addUserSkillRequestBody.SkillId, userId.(uint64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (skillHandler *SkillHandler) DeleteUserSkill(c *gin.Context) {
	var deleteUserSkillRequestBody bodyTemplates.DeleteUserSkillRequestBody
	skillId, err := strconv.ParseInt(c.Param("skill_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}
	deleteUserSkillRequestBody.SkillId = uint64(skillId)
	userId, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	exists, err = skillHandler.skillRepository.UserSkillExist(deleteUserSkillRequestBody.SkillId, userId.(uint64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, err)
		return
	}

	err = skillHandler.skillRepository.DeleteUserSkill(deleteUserSkillRequestBody.SkillId, userId.(uint64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (skillHandler *SkillHandler) GetUserSkills(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		log.Fatal("User Id does not exist!")
	}

	uId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	skills, err := skillHandler.skillRepository.GetUserSkills(userId.(uint64), uint64(uId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, skills)
}
