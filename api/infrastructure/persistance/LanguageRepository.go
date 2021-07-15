package persistance

import (
	"github.com/Atelier-Developers/alethia/Database"
	"github.com/Atelier-Developers/alethia/domain/entity"
	"log"
)

type LanguageRepository struct {
	dbClient *Database.MySQLDB
}

func NewLanguageRepository(dbClient *Database.MySQLDB) *LanguageRepository {
	return &LanguageRepository{dbClient: dbClient}
}

func (languageRepo *LanguageRepository) SaveUserLanguage(userId uint64, languageId uint64) error {
	db := languageRepo.dbClient.GetDB()

	stmt, err := db.Prepare("INSERT INTO USER_LANGUAGE (user_id, language_id) VALUES (?, ?) ")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(userId, languageId)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (languageRepo *LanguageRepository) DeleteUserLanguage(userId uint64, languageId uint64) error {
	db := languageRepo.dbClient.GetDB()

	stmt, err := db.Prepare("DELETE FROM USER_LANGUAGE WHERE user_id=? AND language_id=?")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(userId, languageId)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (languageRepo *LanguageRepository) UserLanguageExists(userId uint64, languageId uint64) (bool, error) {
	db := languageRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT Count(*) FROM USER_LANGUAGE WHERE user_id=? AND language_id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var count uint64
	row := stmt.QueryRow(userId, languageId)
	err = row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count > 0, nil
}

func (languageRepo *LanguageRepository) GetUserLanguages(userId uint64) ([]entity.Language, error) {
	db := languageRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT LANGUAGE.* FROM LANGUAGE, USER_LANGUAGE WHERE LANGUAGE.id = USER_LANGUAGE.language_id AND USER_LANGUAGE.user_id=?")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(userId)
	if err != nil {
		log.Fatal(err)
	}

	var languages []entity.Language
	for rows.Next() {
		var language entity.Language
		err = rows.Scan(&language.ID, &language.Language)
		if err != nil {
			log.Fatal(err)
		}

		languages = append(languages, language)
	}

	return languages, nil
}
