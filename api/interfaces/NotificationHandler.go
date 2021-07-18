package interfaces

import (
	"github.com/Atelier-Developers/alethia/domain/entity"
	"github.com/Atelier-Developers/alethia/domain/entity/Comment"
	"github.com/Atelier-Developers/alethia/domain/entity/Post"
	"github.com/Atelier-Developers/alethia/domain/entity/notification"
	"github.com/Atelier-Developers/alethia/domain/repository"
	"github.com/Atelier-Developers/alethia/interfaces/bodyTemplates"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type NotificationHandler struct {
	postRepository              repository.PostRepository
	skillRepository             repository.SkillRepository
	commentRepository           repository.CommentRepository
	backgroundHistoryRepository repository.BackgroundHistoryRepository
	userRepository              repository.UserRepository
	notificationRepository      repository.NotificationRepository
}

func NewNotificationHandler(postRepository repository.PostRepository, skillRepository repository.SkillRepository, commentRepository repository.CommentRepository, backgroundHistoryRepository repository.BackgroundHistoryRepository, userRepository repository.UserRepository, notificationRepository repository.NotificationRepository) NotificationHandler {
	return NotificationHandler{
		postRepository:              postRepository,
		skillRepository:             skillRepository,
		commentRepository:           commentRepository,
		backgroundHistoryRepository: backgroundHistoryRepository,
		userRepository:              userRepository,
		notificationRepository:      notificationRepository,
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

	var birthdayNotifs []bodyTemplates.NotificationBirthdayResponseBody
	for _, n := range res.([]notification.Birthday) {
		var birthdayNotif bodyTemplates.NotificationBirthdayResponseBody

		birthdayNotif.Id = n.Id
		birthdayNotif.ReceiverId = n.ReceiverId
		birthdayNotif.Created = n.Created
		birthdayNotif.IsSeen = n.IsSeen

		var user entity.UserWithMutualConnectionAndFriendshipStatus
		err = notificationHandler.userRepository.GetUserByID(userId.(uint64), n.UserId, &user)
		if err != nil {
			log.Fatal(err)
		}

		birthdayNotif.Creator.UserID = user.ID
		birthdayNotif.Creator.Username = user.Username
		birthdayNotif.Creator.FirstName = user.FirstName
		birthdayNotif.Creator.LastName = user.LastName
		birthdayNotif.Creator.Intro = user.Intro
		birthdayNotif.Creator.About = user.About
		birthdayNotif.Creator.Accomplishments = user.Accomplishments
		birthdayNotif.Creator.AdditionalInfo = user.AdditionalInfo
		birthdayNotif.Creator.BirthDate = user.BirthDate
		birthdayNotif.Creator.JoinDate = user.JoinDate

		birthdayNotifs = append(birthdayNotifs, birthdayNotif)
	}

	result["birthday"] = birthdayNotifs

	res, err = notificationHandler.notificationRepository.GetChangeWorkNotification(userId.(uint64))
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		res = []notification.ChangeWork{}
	}

	var changeWorkNotifs []bodyTemplates.NotificationChangeWorkResponseBody
	for _, n := range res.([]notification.ChangeWork) {
		var changeWorkNotif bodyTemplates.NotificationChangeWorkResponseBody

		changeWorkNotif.Id = n.Id
		changeWorkNotif.ReceiverId = n.ReceiverId
		changeWorkNotif.Created = n.Created
		changeWorkNotif.IsSeen = n.IsSeen

		background, err := notificationHandler.backgroundHistoryRepository.GetBackgroundHistory(n.UserHistoryId)
		if err != nil {
			log.Fatal(err)
		}

		changeWorkNotif.BackgroundHistory = background

		var user entity.UserWithMutualConnectionAndFriendshipStatus
		err = notificationHandler.userRepository.GetUserByID(userId.(uint64), background.UserID, &user)
		if err != nil {
			log.Fatal(err)
		}

		changeWorkNotif.Creator.UserID = user.ID
		changeWorkNotif.Creator.Username = user.Username
		changeWorkNotif.Creator.FirstName = user.FirstName
		changeWorkNotif.Creator.LastName = user.LastName
		changeWorkNotif.Creator.Intro = user.Intro
		changeWorkNotif.Creator.About = user.About
		changeWorkNotif.Creator.Accomplishments = user.Accomplishments
		changeWorkNotif.Creator.AdditionalInfo = user.AdditionalInfo
		changeWorkNotif.Creator.BirthDate = user.BirthDate
		changeWorkNotif.Creator.JoinDate = user.JoinDate

		changeWorkNotifs = append(changeWorkNotifs, changeWorkNotif)
	}

	result["change_work"] = changeWorkNotifs

	res, err = notificationHandler.notificationRepository.GetCommentNotification(userId.(uint64))
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		res = []notification.Comment{}
	}

	var commentNotifs []bodyTemplates.NotificationCommentResponseBody
	for _, n := range res.([]notification.Comment) {
		var commentNotif bodyTemplates.NotificationCommentResponseBody

		commentNotif.Id = n.Id
		commentNotif.ReceiverId = n.ReceiverId
		commentNotif.Created = n.Created
		commentNotif.IsSeen = n.IsSeen

		var comment Comment.Comment
		err := notificationHandler.commentRepository.GetCommentByID(n.CommentId, &comment)
		if err != nil {
			log.Fatal(err)
		}

		var user entity.UserWithMutualConnectionAndFriendshipStatus
		err = notificationHandler.userRepository.GetUserByID(userId.(uint64), comment.CommenterId, &user)
		if err != nil {
			log.Fatal(err)
		}

		commentNotif.Creator.UserID = user.ID
		commentNotif.Creator.Username = user.Username
		commentNotif.Creator.FirstName = user.FirstName
		commentNotif.Creator.LastName = user.LastName
		commentNotif.Creator.Intro = user.Intro
		commentNotif.Creator.About = user.About
		commentNotif.Creator.Accomplishments = user.Accomplishments
		commentNotif.Creator.AdditionalInfo = user.AdditionalInfo
		commentNotif.Creator.BirthDate = user.BirthDate
		commentNotif.Creator.JoinDate = user.JoinDate

		commentNotifs = append(commentNotifs, commentNotif)
	}

	result["comment"] = commentNotifs

	res, err = notificationHandler.notificationRepository.GetEndorseSkillNotification(userId.(uint64))
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		res = []notification.EndorseSkill{}
	}

	var endorseSkillNotifs []bodyTemplates.NotificationEndorseResponseBody
	for _, n := range res.([]notification.EndorseSkill) {
		var endorseSkillNotif bodyTemplates.NotificationEndorseResponseBody

		endorseSkillNotif.Id = n.Id
		endorseSkillNotif.ReceiverId = n.ReceiverId
		endorseSkillNotif.Created = n.Created
		endorseSkillNotif.IsSeen = n.IsSeen

		skill, err := notificationHandler.skillRepository.GetSkillById(n.UserSkillSkillId)

		endorseSkillNotif.UserSkill = skill

		var user entity.UserWithMutualConnectionAndFriendshipStatus
		err = notificationHandler.userRepository.GetUserByID(userId.(uint64), n.UserId, &user)
		if err != nil {
			log.Fatal(err)
		}

		endorseSkillNotif.Creator.UserID = user.ID
		endorseSkillNotif.Creator.Username = user.Username
		endorseSkillNotif.Creator.FirstName = user.FirstName
		endorseSkillNotif.Creator.LastName = user.LastName
		endorseSkillNotif.Creator.Intro = user.Intro
		endorseSkillNotif.Creator.About = user.About
		endorseSkillNotif.Creator.Accomplishments = user.Accomplishments
		endorseSkillNotif.Creator.AdditionalInfo = user.AdditionalInfo
		endorseSkillNotif.Creator.BirthDate = user.BirthDate
		endorseSkillNotif.Creator.JoinDate = user.JoinDate

		endorseSkillNotifs = append(endorseSkillNotifs, endorseSkillNotif)
	}

	result["endorse"] = endorseSkillNotifs

	res, err = notificationHandler.notificationRepository.GetLikeCommentNotification(userId.(uint64))
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		res = []notification.LikeComment{}
	}

	var likeCommentNotifs []bodyTemplates.NotificationLikeComment
	for _, n := range res.([]notification.LikeComment) {
		var likeCommentNotif bodyTemplates.NotificationLikeComment

		likeCommentNotif.Id = n.Id
		likeCommentNotif.ReceiverId = n.ReceiverId
		likeCommentNotif.Created = n.Created
		likeCommentNotif.IsSeen = n.IsSeen

		var comment Comment.Comment
		err := notificationHandler.commentRepository.GetCommentByID(n.CommentId, &comment)
		if err != nil {
			log.Fatal(err)
		}

		var user entity.UserWithMutualConnectionAndFriendshipStatus
		err = notificationHandler.userRepository.GetUserByID(userId.(uint64), n.UserId, &user)
		if err != nil {
			log.Fatal(err)
		}

		likeCommentNotif.Creator.UserID = user.ID
		likeCommentNotif.Creator.Username = user.Username
		likeCommentNotif.Creator.FirstName = user.FirstName
		likeCommentNotif.Creator.LastName = user.LastName
		likeCommentNotif.Creator.Intro = user.Intro
		likeCommentNotif.Creator.About = user.About
		likeCommentNotif.Creator.Accomplishments = user.Accomplishments
		likeCommentNotif.Creator.AdditionalInfo = user.AdditionalInfo
		likeCommentNotif.Creator.BirthDate = user.BirthDate
		likeCommentNotif.Creator.JoinDate = user.JoinDate

		likeCommentNotifs = append(likeCommentNotifs, likeCommentNotif)
	}

	result["like_comment"] = likeCommentNotifs

	res, err = notificationHandler.notificationRepository.GetLikePostNotification(userId.(uint64))
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		res = []notification.LikePost{}
	}

	var likePostNotifs []bodyTemplates.NotificationLikePost
	for _, n := range res.([]notification.LikePost) {
		var likePostNotif bodyTemplates.NotificationLikePost

		likePostNotif.Id = n.Id
		likePostNotif.ReceiverId = n.ReceiverId
		likePostNotif.Created = n.Created
		likePostNotif.IsSeen = n.IsSeen

		var post Post.Post
		err := notificationHandler.postRepository.GetPostById(n.PostId, &post)
		if err != nil {
			log.Fatal(err)
		}

		var user entity.UserWithMutualConnectionAndFriendshipStatus
		err = notificationHandler.userRepository.GetUserByID(userId.(uint64), n.UserId, &user)
		if err != nil {
			log.Fatal(err)
		}

		likePostNotif.Creator.UserID = user.ID
		likePostNotif.Creator.Username = user.Username
		likePostNotif.Creator.FirstName = user.FirstName
		likePostNotif.Creator.LastName = user.LastName
		likePostNotif.Creator.Intro = user.Intro
		likePostNotif.Creator.About = user.About
		likePostNotif.Creator.Accomplishments = user.Accomplishments
		likePostNotif.Creator.AdditionalInfo = user.AdditionalInfo
		likePostNotif.Creator.BirthDate = user.BirthDate
		likePostNotif.Creator.JoinDate = user.JoinDate

		likePostNotifs = append(likePostNotifs, likePostNotif)
	}

	result["like_post"] = likePostNotifs

	res, err = notificationHandler.notificationRepository.GetReplyCommentNotification(userId.(uint64))
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		res = []notification.LikeComment{}
	}

	var replyCommentNotifs []bodyTemplates.NotificationReply
	for _, n := range res.([]notification.ReplyComment) {
		var replyCommentNotif bodyTemplates.NotificationReply

		replyCommentNotif.Id = n.Id
		replyCommentNotif.ReceiverId = n.ReceiverId
		replyCommentNotif.Created = n.Created
		replyCommentNotif.IsSeen = n.IsSeen

		var comment Comment.Comment
		err := notificationHandler.commentRepository.GetCommentByID(n.CommentId, &comment)
		if err != nil {
			log.Fatal(err)
		}

		var user entity.UserWithMutualConnectionAndFriendshipStatus
		err = notificationHandler.userRepository.GetUserByID(userId.(uint64), comment.CommenterId, &user)
		if err != nil {
			log.Fatal(err)
		}

		replyCommentNotif.Creator.UserID = user.ID
		replyCommentNotif.Creator.Username = user.Username
		replyCommentNotif.Creator.FirstName = user.FirstName
		replyCommentNotif.Creator.LastName = user.LastName
		replyCommentNotif.Creator.Intro = user.Intro
		replyCommentNotif.Creator.About = user.About
		replyCommentNotif.Creator.Accomplishments = user.Accomplishments
		replyCommentNotif.Creator.AdditionalInfo = user.AdditionalInfo
		replyCommentNotif.Creator.BirthDate = user.BirthDate
		replyCommentNotif.Creator.JoinDate = user.JoinDate

		replyCommentNotifs = append(replyCommentNotifs, replyCommentNotif)
	}

	result["reply_comment"] = replyCommentNotifs

	res, err = notificationHandler.notificationRepository.GetViewProfileNotification(userId.(uint64))
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		res = []notification.ViewProfile{}
	}

	var viewProfileNotifs []bodyTemplates.NotificationViewProfile
	for _, n := range res.([]notification.ViewProfile) {
		var viewProfileNotif bodyTemplates.NotificationViewProfile

		viewProfileNotif.Id = n.Id
		viewProfileNotif.ReceiverId = n.ReceiverId
		viewProfileNotif.Created = n.Created
		viewProfileNotif.IsSeen = n.IsSeen

		var user entity.UserWithMutualConnectionAndFriendshipStatus
		err = notificationHandler.userRepository.GetUserByID(userId.(uint64), n.UserId, &user)
		if err != nil {
			log.Fatal(err)
		}

		viewProfileNotif.Creator.UserID = user.ID
		viewProfileNotif.Creator.Username = user.Username
		viewProfileNotif.Creator.FirstName = user.FirstName
		viewProfileNotif.Creator.LastName = user.LastName
		viewProfileNotif.Creator.Intro = user.Intro
		viewProfileNotif.Creator.About = user.About
		viewProfileNotif.Creator.Accomplishments = user.Accomplishments
		viewProfileNotif.Creator.AdditionalInfo = user.AdditionalInfo
		viewProfileNotif.Creator.BirthDate = user.BirthDate
		viewProfileNotif.Creator.JoinDate = user.JoinDate

		viewProfileNotifs = append(viewProfileNotifs, viewProfileNotif)
	}
	result["view_profile"] = viewProfileNotifs

	c.JSON(http.StatusOK, result)
}
