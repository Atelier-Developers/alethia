package repository

import "github.com/Atelier-Developers/alethia/domain/entity"

type BackgroundHistoryRepository interface {
	SaveBackgroundHistory(backgroundHistory *entity.BackgroundHistory) error
}
