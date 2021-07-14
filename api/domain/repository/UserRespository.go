package repository

import (
	"github.com/Atelier-Developers/alethia/domain/entity"
)

type UserRepository interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	GetUser(uint64) (*entity.User, error)
	GetUsers() ([]entity.User, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, map[string]string)
}