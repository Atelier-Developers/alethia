package persistance

import (
	"github.com/Atelier-Developers/alethia/Database"
	"github.com/Atelier-Developers/alethia/domain/entity"
	"log"
)

type InviteRepository struct {
	dbClient *Database.MySQLDB
}

func NewInviteRepository(dbClient *Database.MySQLDB) *InviteRepository {
	return &InviteRepository{dbClient: dbClient}
}

func (inviteRepo *InviteRepository) AddInvite(invite entity.Invite) error {
	db := inviteRepo.dbClient.GetDB()
	stmt, err := db.Prepare("INSERT INTO INVITE (user1_id, user2_id, body) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(invite.UserId1, invite.UserId2, invite.Body)

	if err != nil {
		return err
	}

	return nil
}

func (inviteRepo *InviteRepository) DeleteInvite(userId1 uint64, userId2 uint64) error {
	db := inviteRepo.dbClient.GetDB()
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

func (inviteRepo *InviteRepository) IsUserInvitedById(userId1 uint64, userId2 uint64) (bool, error) {
	db := inviteRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT Count(*) FROM INVITE WHERE user1_id=? AND user2_id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var count uint64
	row := stmt.QueryRow(userId2, userId1)
	err = row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count > 0, nil
}

func (inviteRepo *InviteRepository) InviteExists(userId1 uint64, userId2 uint64) (bool, error) {
	db := inviteRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT Count(*) FROM INVITE WHERE (user1_id=? AND user2_id=?) OR (user1_id=? AND user2_id=?)")
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

func (inviteRepo *InviteRepository) GetInvites(userId uint64) ([]entity.Invite, error) {
	db := inviteRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT INVITE.*, U1.username, U2.username FROM INVITE, USER U1, USER U2 WHERE (user1_id=? OR user2_id=?) AND U1.id = INVITE.user1_id AND U2.id = INVITE.user2_id ")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId, userId)

	if err != nil {
		log.Fatal(err)
	}

	var invites []entity.Invite
	for rows.Next() {
		var invite entity.Invite
		err = rows.Scan(&invite.UserId1, &invite.UserId2, &invite.Created, &invite.Body, &invite.Username1, &invite.Username2)
		if err != nil {
			log.Fatal(err)
		}

		invites = append(invites, invite)
	}

	return invites, nil
}
