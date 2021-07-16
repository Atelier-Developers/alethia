package repository

import "github.com/Atelier-Developers/alethia/domain/entity"

type InviteRepository interface {
	InviteExists(userId1 uint64, userId2 uint64) (bool, error)
	IsUserInvitedById(userId1 uint64, userId2 uint64) (bool, error)
	AddInvite(invite entity.Invite) error
	DeleteInvite(userId1 uint64, userId2 uint64) error
	GetInvites(userId uint64) ([]entity.Invite, error)
}
