package persistance

import (
	"github.com/Atelier-Developers/alethia/Database"
	"github.com/Atelier-Developers/alethia/domain/entity"
	"log"
)

type FriendRepository struct {
	dbClient *Database.MySQLDB
}

func NewFriendRepository(dbClient *Database.MySQLDB) *FriendRepository {
	return &FriendRepository{dbClient: dbClient}
}


func (friendRepository *FriendRepository) FriendExists(userId1 uint64, userId2 uint64) (bool, error) {
	db := friendRepository.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT Count(*) FROM FRIEND WHERE (user1_id=? AND user2_id=?) OR (user1_id=? AND user2_id=?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var count uint64
	row := stmt.QueryRow(userId1, userId2, userId2, userId1)
	err = row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count > 0, nil
}

func (friendRepository *FriendRepository) AddFriend(userId1 uint64, userId2 uint64) error {
	db := friendRepository.dbClient.GetDB()
	stmt, err := db.Prepare("INSERT INTO FRIEND (user1_id, user2_id) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(userId1, userId2)

	if err != nil {
		return err
	}

	return nil
}

func (friendRepository *FriendRepository) DeleteFriend(userId1 uint64, userId2 uint64) error {
	db := friendRepository.dbClient.GetDB()
	stmt, err := db.Prepare("DELETE FROM INVITE WHERE user1_id=? AND user2_id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(userId1, userId2)

	if err != nil {
		return err
	}

	return nil
}

func (friendRepository *FriendRepository) GetFriends(userId uint64) ([]entity.Friend, error) {
	db := friendRepository.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT * FROM FRIEND WHERE user1_id=? OR user2_id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId, userId)

	if err != nil {
		log.Fatal(err)
	}

	var friends []entity.Friend
	for rows.Next() {
		var friend entity.Friend
		err = rows.Scan(&friend.UserId1, &friend.UserId2, &friend.Created)
		if err != nil {
			log.Fatal(err)
		}

		friends = append(friends, friend)
	}

	return friends, nil
}
