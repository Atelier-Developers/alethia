package persistance

import (
	"github.com/Atelier-Developers/alethia/Database"
	"github.com/Atelier-Developers/alethia/domain/entity"
	"github.com/Atelier-Developers/alethia/infrastructure/security"
	"log"
)

type UserRepository struct {
	dbClient *Database.MySQLDB
}

func NewUserRepository(dbClient *Database.MySQLDB) *UserRepository {
	return &UserRepository{dbClient: dbClient}
}

func (userRepo *UserRepository) GetUsersWithSimilarUsername(username string) ([]entity.User, error) {
	db := userRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT * FROM USER WHERE username LIKE '%?%'")
	if err != nil {

		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(username)

	if err != nil {
		log.Fatal(err)
	}

	var users []entity.User
	for rows.Next() {
		var user entity.User
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Password, &user.Intro, &user.About, &user.Accomplishments, &user.AdditionalInfo, &user.JoinDate, &user.BirthDate)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	return users, nil
}

func (userRepo *UserRepository) GetUserByUsername(username string, user *entity.User) error {
	db := userRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT * FROM USER WHERE username=?")
	if err != nil {

		log.Fatal(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(username)

	err = row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Password, &user.Intro, &user.About, &user.Accomplishments, &user.AdditionalInfo, &user.JoinDate, &user.BirthDate)
	if err != nil {
		return err
	}

	return nil
}

func (userRepo *UserRepository) GetUserByID(id uint64, user *entity.User) error {
	db := userRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT * FROM USER WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	err = row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Password, &user.Intro, &user.About, &user.Accomplishments, &user.AdditionalInfo, &user.JoinDate, &user.BirthDate)
	if err != nil {
		return err
	}

	return nil
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

func (userRepo *UserRepository) UpdateUser(user *entity.User) error {
	db := userRepo.dbClient.GetDB()

	stmt, err := db.Prepare("UPDATE USER SET intro = CASE WHEN ? IS NOT NULL THEN ? ELSE intro END, about= CASE WHEN ? IS NOT NULL THEN ? ELSE about END, accomplishments= CASE WHEN ? IS NOT NULL THEN ? ELSE accomplishments END, additional_information= CASE WHEN ? IS NOT NULL THEN ? ELSE additional_information END, birthdate= CASE WHEN ? IS NOT NULL THEN ? ELSE birthdate END WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Intro, user.Intro, user.About, user.About, user.Accomplishments, user.Accomplishments, user.AdditionalInfo, user.AdditionalInfo, user.BirthDate, user.BirthDate, user.ID)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (userRepo *UserRepository) GetUserByUsernameAndPassword(username string, password string, user *entity.User) error {
	db := userRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT * FROM USER WHERE username=?")
	if err != nil {

		log.Fatal(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(username)

	err = row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Password, &user.Intro, &user.About, &user.Accomplishments, &user.AdditionalInfo, &user.JoinDate, &user.BirthDate)
	if err != nil {
		return err
	}

	err = security.VerifyPassword(string(user.Password), password)
	if err != nil {
		return err
	}

	return nil
}
