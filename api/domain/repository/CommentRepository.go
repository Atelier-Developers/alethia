package repository

import (
	"github.com/Atelier-Developers/alethia/domain/entity/Comment"
)

type CommentRepository interface {
	SaveComment(comment *Comment.Comment) error
	LikeComment(comment *Comment.Comment, UserId uint64) error
	ReplyComment(comment *Comment.Comment, ReplyId uint64) error
	GetCommentByID(id uint64, comment *Comment.Comment) error
	GetCommentLikes(commentId uint64) ([]Comment.CommentLike, error)
}
