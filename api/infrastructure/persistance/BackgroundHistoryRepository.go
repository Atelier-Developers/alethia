package persistance

import (
	"github.com/Atelier-Developers/alethia/Database"
	"github.com/Atelier-Developers/alethia/domain/entity"
	"log"
)

type BackgroundHistoryRepository struct {
	dbClient *Database.MySQLDB
}

func NewBackgroundHistoryRepository(dbClient *Database.MySQLDB) *BackgroundHistoryRepository {
	return &BackgroundHistoryRepository{dbClient: dbClient}
}

func (backgroundHistoryRepo *BackgroundHistoryRepository) SaveBackgroundHistory(backgroundHistory *entity.BackgroundHistory) error {
	db := backgroundHistoryRepo.dbClient.GetDB()
	stmt, err := db.Prepare("INSERT INTO USER_HISTORY (user_id, start_date, end_date, category, title, description, location) VALUES (?, ?, ?, ?, ?, ?, ?) ")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(backgroundHistory.UserID, backgroundHistory.StartDate, backgroundHistory.EndDate, backgroundHistory.Category, backgroundHistory.Title, backgroundHistory.Description, backgroundHistory.Location)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (backgroundHistoryRepo *BackgroundHistoryRepository) UpdateBackgroundHistory(backgroundHistory *entity.BackgroundHistory) error {
	db := backgroundHistoryRepo.dbClient.GetDB()

	stmt, err := db.Prepare("UPDATE USER_HISTORY SET start_date = CASE WHEN ? IS NOT NULL THEN ? ELSE start_date END, end_date= CASE WHEN ? IS NOT NULL THEN ? ELSE end_date END, category= CASE WHEN ? IS NOT NULL THEN ? ELSE category END, title= CASE WHEN ? IS NOT NULL THEN ? ELSE title END, description= CASE WHEN ? IS NOT NULL THEN ? ELSE description END, location= CASE WHEN ? IS NOT NULL THEN ? ELSE location END WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(backgroundHistory.StartDate, backgroundHistory.StartDate, backgroundHistory.EndDate, backgroundHistory.EndDate, backgroundHistory.Category, backgroundHistory.Category, backgroundHistory.Title, backgroundHistory.Title, backgroundHistory.Description, backgroundHistory.Description, backgroundHistory.Location, backgroundHistory.Location, backgroundHistory.ID)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (backgroundHistoryRepo *BackgroundHistoryRepository) DeleteBackgroundHistory(backgroundHistory *entity.BackgroundHistory) error {
	db := backgroundHistoryRepo.dbClient.GetDB()

	stmt, err := db.Prepare("DELETE FROM USER_HISTORY WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(backgroundHistory.ID)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
