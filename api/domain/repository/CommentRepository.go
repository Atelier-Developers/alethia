package repository

import "github.com/Atelier-Developers/alethia/domain/entity"

type CommentRepository interface {
	SaveComment(comment *entity.Comment) error
	LikeComment(comment *entity.Comment, UserId uint64) error
	ReplyComment(comment *entity.Comment, ReplyId uint64) error
	GetCommentByID(id uint64, comment *entity.Comment) error
	GetCommentNumberOfLikes(commentId uint64) (*entity.CommentLikeCount, error)
}
