package repository

import "github.com/Atelier-Developers/alethia/domain/entity"

type PostRepository interface {
	SavePost(post *entity.Post) error
	LikePost(post *entity.Post, userId uint64) error
	GetPostById(postId uint64, post *entity.Post) error
}
