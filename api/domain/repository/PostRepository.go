package repository

import (
	"github.com/Atelier-Developers/alethia/domain/entity/Post"
)

type PostRepository interface {
	SavePost(post *Post.Post) error
	LikePost(post *Post.Post, userId uint64) error
	GetPostById(postId uint64, post *Post.Post) error
	//GetPostLikes(postId uint64) ([])
}
