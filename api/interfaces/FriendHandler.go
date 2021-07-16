package interfaces

import (
	"github.com/Atelier-Developers/alethia/domain/entity"
	"github.com/Atelier-Developers/alethia/domain/repository"
	"github.com/Atelier-Developers/alethia/interfaces/bodyTemplates"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type FriendHandler struct {
	friendRepository repository.FriendRepository
	inviteRepository repository.InviteRepository
}

func NewFriendHandler(friendRepository repository.FriendRepository, inviteRepository repository.InviteRepository) FriendHandler {
	return FriendHandler{
		friendRepository: friendRepository,
		inviteRepository: inviteRepository,
	}
}

func (friendHandler *FriendHandler) AddInvite(c *gin.Context) {
	var userRequestBody bodyTemplates.InviteCreateRequestBody
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

	exists, err := friendHandler.friendRepository.FriendExists(userId.(uint64), userRequestBody.ReceiverId)

	if exists {
		log.Fatal("Friendship exists.")
	}

	exists, err = friendHandler.inviteRepository.InviteExists(userId.(uint64), userRequestBody.ReceiverId)

	if err != nil {
		log.Fatal(err)
	}

	if exists {
		log.Fatal("Invite Already Exists.")
	}

	invite := entity.Invite{
		UserId1: userId.(uint64),
		UserId2: userRequestBody.ReceiverId,
		Body:    userRequestBody.Body,
	}
	err = friendHandler.inviteRepository.AddInvite(invite)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (friendHandler *FriendHandler) AddFriend(c *gin.Context) {
	var userRequestBody bodyTemplates.FriendCreateRequestBody
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

	exists, err := friendHandler.friendRepository.FriendExists(userId.(uint64), userRequestBody.UserId)

	if err != nil {
		log.Fatal(err)
	}

	if exists {
		log.Fatal("Friendship exists.")
	}

	exists, err = friendHandler.inviteRepository.IsUserInvitedById(userId.(uint64), userRequestBody.UserId)
	if err != nil {
		log.Fatal(err)
	}

	if !exists {
		log.Fatal("This user has not been invited by the user.")
	}

	err = friendHandler.inviteRepository.DeleteInvite(userId.(uint64), userRequestBody.UserId)
	if err != nil {
		log.Fatal(err)
	}

	err = friendHandler.friendRepository.AddFriend(userId.(uint64), userRequestBody.UserId)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, nil)
}

func (friendHandler *FriendHandler) DeleteInvite(c *gin.Context) {
	var userRequestBody bodyTemplates.InviteDeleteRequestBody
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

	exists, err := friendHandler.inviteRepository.IsUserInvitedById(userId.(uint64), userRequestBody.UserId)
	if err != nil {
		log.Fatal(err)
	}

	if !exists {
		log.Fatal("Invite does not exist!")
	}

	err = friendHandler.inviteRepository.DeleteInvite(userId.(uint64), userRequestBody.UserId)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, nil)
}

func (friendHandler *FriendHandler) GetUserInvites(c *gin.Context) {

	userId, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	invites, err := friendHandler.inviteRepository.GetInvites(userId.(uint64))
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, invites)
}

func (friendHandler *FriendHandler) DeleteFriend(c *gin.Context) {
	var userRequestBody bodyTemplates.FriendDeleteRequestBody
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

	exists, err := friendHandler.friendRepository.FriendExists(userId.(uint64), userRequestBody.UserId)
	if err != nil {
		log.Fatal(err)
	}

	if !exists {
		log.Fatal("Friendship does not exist.")
	}

	err = friendHandler.friendRepository.DeleteFriend(userId.(uint64), userRequestBody.UserId)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, nil)
}

func (friendHandler *FriendHandler) GetFriends(c *gin.Context) {
	var userRequestBody bodyTemplates.FriendDeleteRequestBody
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

	friends, err := friendHandler.friendRepository.GetFriends(userId.(uint64))

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, friends)
}
