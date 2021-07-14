package persistance

import (
	"github.com/Atelier-Developers/alethia/Database"
	"github.com/Atelier-Developers/alethia/domain/entity"
	"log"
)

type UserRepository struct {
	dbClient *Database.MySQLDB
}

func NewUserRepository(dbClient *Database.MySQLDB) *UserRepository {
	return &UserRepository{dbClient: dbClient}
}

func (userRepo *UserRepository) SaveUser(user *entity.User) error {
	db := userRepo.dbClient.GetDB()
	stmt, err := db.Prepare("INSERT INTO USER (first_name, last_name, username, password, intro, about, accomplishments, additional_information, join_date, birthdate) VALUES  (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Username, user.Password, user.Intro, user.About, user.Accomplishments, user.AdditionalInfo, user.JoinDate, user.BirthDate)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
