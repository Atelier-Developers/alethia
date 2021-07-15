package repository

import "github.com/Atelier-Developers/alethia/domain/entity/notification"

type NotificationRepository interface {
	CreateCommentNotification(comment *notification.Comment) error
	GetCommentNotification(userId uint64) ([]notification.Comment, error)
	CreateEndorseSkillNotification(endorseSkill *notification.EndorseSkill) error
	GetEndorseSkillNotification(userId uint64) ([]notification.EndorseSkill, error)
}
