package repository

import (
	"github.com/Atelier-Developers/alethia/domain/entity"
)

type UserRepository interface {
	SaveUser(user *entity.User) error
	UpdateUser(user *entity.User) error
	GetUserByUsernameAndPassword(username string, password string, user *entity.User) error
	GetUserByUsername(username string, user *entity.User) error
	GetUserByID(userId uint64, id uint64, user *entity.UserWithMutualConnection) error
	GetUsersWithSimilarUsername(userId uint64, username string) ([]entity.UserWithMutualConnection, error)
	GetUsersWithMutualConnection(userId uint64) ([]entity.UserWithMutualConnection, error)
}
