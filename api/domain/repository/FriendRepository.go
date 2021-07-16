package repository

import "github.com/Atelier-Developers/alethia/domain/entity"

type FriendRepository interface {
	FriendExists(userId1 uint64, userId2 uint64) (bool, error)
	AddFriend(userId1 uint64, userId2 uint64) error
	DeleteFriend(userId1 uint64, userId2 uint64) error
	GetFriends(userId1 uint64) ([]entity.Friend, error)
}
