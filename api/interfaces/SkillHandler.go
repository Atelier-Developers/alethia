package interfaces

import (
	"github.com/Atelier-Developers/alethia/domain/repository"
	"github.com/Atelier-Developers/alethia/interfaces/bodyTemplates"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type SkillHandler struct {
	skillRepository repository.SkillRepository
}

func NewSkillHandler(skillRepository repository.SkillRepository) SkillHandler {
	return SkillHandler{
		skillRepository: skillRepository,
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

	err := skillHandler.skillRepository.EndorseSkill(endorseSkillRequestBody.SkillId, endorseSkillRequestBody.UserId, userId.(uint64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
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
	if err := c.ShouldBindJSON(&deleteUserSkillRequestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	userId, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	exists, err := skillHandler.skillRepository.UserSkillExist(deleteUserSkillRequestBody.SkillId, userId.(uint64))
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

	skills, err := skillHandler.skillRepository.GetUserSkills(userId.(uint64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, skills)
}