package repository

import (
	"github.com/Atelier-Developers/alethia/domain/entity"
)

type UserRepository interface {
	SaveUser(user *entity.User) error
}
