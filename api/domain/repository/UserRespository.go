package repository

import (
	"github.com/Atelier-Developers/alethia/domain/entity"
)

type UserRepository interface {
	SaveUser(user *entity.User) error
	UpdateUser(user *entity.User) error
	GetUserByUsernameAndPassword(username string, password string, user *entity.User) error
	GetUserByUsername(username string, user *entity.User) error
	GetUserByID(id uint64, user *entity.User) error
	GetUsersWithSimilarUsername(username string) ([]entity.User, error)
}
