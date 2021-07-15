package repository

import "github.com/Atelier-Developers/alethia/domain/entity"

type BackgroundHistoryRepository interface {
	SaveBackgroundHistory(backgroundHistory *entity.BackgroundHistory) error
	UpdateBackgroundHistory(backgroundHistory *entity.BackgroundHistory) error
	DeleteBackgroundHistory(backgroundHistory *entity.BackgroundHistory) error
	GetUserBackgroundHistory(userId uint64) ([]entity.BackgroundHistory, error)
}
