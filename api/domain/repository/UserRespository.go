package repository

import (
	"github.com/Atelier-Developers/alethia/domain/entity"
)

type UserRepository interface {
	DoesUserWorkAtCompany(userId uint64, companyName string) (bool, error)
	DoesUserHaveLanguage(userId uint64, language string) (bool, error)
	SaveUser(user *entity.User) error
	UpdateUser(user *entity.User) error
	GetUserByUsernameAndPassword(username string, password string, user *entity.User) error
	GetUserByUsername(username string, user *entity.User) error
	GetUserByID(userId uint64, id uint64, user *entity.UserWithMutualConnectionAndFriendshipStatus) error
	GetUsersWithSimilarUsername(userId uint64, username string) ([]entity.UserWithMutualConnectionAndFriendshipStatus, error)
	GetUsersWithMutualConnection(userId uint64) ([]entity.UserWithMutualConnection, error)
}
