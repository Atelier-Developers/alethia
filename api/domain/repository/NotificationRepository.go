package repository

import "github.com/Atelier-Developers/alethia/domain/entity/notification"

type NotificationRepository interface {
	CreateCommentNotification(comment *notification.Comment) error
	GetCommentNotification(userId uint64) ([]notification.Comment, error)
	CreateEndorseSkillNotification(endorseSkill *notification.EndorseSkill) error
	GetEndorseSkillNotification(userId uint64) ([]notification.EndorseSkill, error)
	CreateChangeWorkNotification(changeWork *notification.ChangeWork) error
	DeleteChangeWorkNotifications(changeWork *notification.ChangeWork) error
	GetChangeWorkNotification(userId uint64) ([]notification.ChangeWork, error)
	CreateLikeCommentNotification(likeComment *notification.LikeComment) error
	GetLikeCommentNotification(userId uint64) ([]notification.LikeComment, error)
	CreateLikePostNotification(likePost *notification.LikePost) error
	GetLikePostNotification(userId uint64) ([]notification.LikePost, error)
	CreateReplyCommentNotification(likeComment *notification.ReplyComment) error
	GetReplyCommentNotification(userId uint64) ([]notification.ReplyComment, error)
	CreateViewProfileNotification(likeComment *notification.ViewProfile) error
	GetViewProfileNotification(userId uint64) ([]notification.ViewProfile, error)
	CreateBirthdayNotification(likeComment *notification.Birthday) error
	GetBirthdayNotification(userId uint64) ([]notification.Birthday, error)
}
