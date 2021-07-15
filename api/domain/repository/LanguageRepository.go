package repository

import "github.com/Atelier-Developers/alethia/domain/entity"

type LanguageRepository interface {
	SaveUserLanguage(userId uint64, languageId uint64) error
	DeleteUserLanguage(userId uint64, languageId uint64) error
	UserLanguageExists(userId uint64, languageId uint64) (bool, error)
	GetUserLanguages(userId uint64) ([]entity.Language, error)
}
