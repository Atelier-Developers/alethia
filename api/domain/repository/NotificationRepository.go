package repository

import "github.com/Atelier-Developers/alethia/domain/entity/notification"

type NotificationRepository interface {
	CreateCommentNotification(comment *notification.Comment) error
	GetCommentNotifications() ([]notification.Comment, error)
}
