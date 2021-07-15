package persistance

import (
	"github.com/Atelier-Developers/alethia/Database"
	"github.com/Atelier-Developers/alethia/domain/entity"
	"log"
)

type ConversationRepository struct {
	dbClient *Database.MySQLDB
}

func NewConversationRepository(dbClient *Database.MySQLDB) *ConversationRepository {
	return &ConversationRepository{dbClient: dbClient}
}

func (conversationRepo ConversationRepository) ConversationExists(userId1 uint64, userId2 uint64) (bool, error) {
	db := conversationRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT Count(*) FROM CONVERSATION WHERE (user1_id=? AND user2_id=?) OR (user1_id=? AND user2_id=?)")
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

func (conversationRepo *ConversationRepository) SaveConversation(conversation *entity.Conversation) error {
	db := conversationRepo.dbClient.GetDB()
	stmt, err := db.Prepare("INSERT INTO CONVERSATION (user1_id, user2_id, is_archived, is_deleted, is_read) VALUES (?, ?, ?, ?, ?) ")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(conversation.UserId1, conversation.UserId2, conversation.IsArchived, conversation.IsDeleted, conversation.IsRead)

	if err != nil {
		return err
	}

	conversationId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	conversation.Id = uint64(conversationId)

	return nil
}

func (conversationRepo *ConversationRepository) UpdateConversation(conversation entity.Conversation) error {

	db := conversationRepo.dbClient.GetDB()

	stmt, err := db.Prepare("UPDATE CONVERSATION SET is_archived = CASE WHEN ? IS NOT NULL THEN ? ELSE is_archived END, is_deleted= CASE WHEN ? IS NOT NULL THEN ? ELSE is_deleted END, is_read= CASE WHEN ? IS NOT NULL THEN ? ELSE is_read END WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(conversation.IsArchived, conversation.IsArchived, conversation.IsDeleted, conversation.IsDeleted, conversation.IsRead, conversation.IsRead, conversation.Id)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (conversationRepo *ConversationRepository) DeleteConversation(conversation entity.Conversation) error {
	db := conversationRepo.dbClient.GetDB()

	stmt, err := db.Prepare("DELETE FROM CONVERSATION WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(conversation.Id)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (conversationRepo *ConversationRepository) GetConversation(userId1 uint64, userId2 uint64) (*entity.Conversation, error) {
	db := conversationRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT * FROM CONVERSATION WHERE user1_id=? AND user2_id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(userId1, userId2)

	var convo entity.Conversation
	err = row.Scan(&convo.Id, &convo.UserId1, &convo.UserId2, &convo.IsArchived, &convo.IsDeleted, &convo.IsRead)
	if err != nil {
		return nil, err
	}

	return &convo, nil
}

func (conversationRepo *ConversationRepository) GetUserConversations(userId uint64) ([]entity.Conversation, error) {
	db := conversationRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT * FROM CONVERSATION WHERE user1_id=? OR user2_id=?")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(userId, userId)
	if err != nil {
		log.Fatal(err)
	}

	var conversations []entity.Conversation
	for rows.Next() {
		var conversation entity.Conversation
		err = rows.Scan(&conversation.Id, &conversation.UserId1, &conversation.UserId2, &conversation.IsArchived, &conversation.IsDeleted, &conversation.IsRead)
		if err != nil {
			log.Fatal(err)
		}

		conversations = append(conversations, conversation)
	}

	return conversations, nil
}
