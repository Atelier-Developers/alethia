package repository

import "github.com/Atelier-Developers/alethia/domain/entity"

type PostRepository interface {
	SavePost(post *entity.Post) error
	LikePost(post *entity.Post) error
}
