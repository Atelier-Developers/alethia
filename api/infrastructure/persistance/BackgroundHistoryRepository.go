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


func (backgroundRepo *BackgroundHistoryRepository) SaveBackgroundHistory(backgroundHistory *entity.BackgroundHistory) error {
	db := backgroundRepo.dbClient.GetDB()
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


