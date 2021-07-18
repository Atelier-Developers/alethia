package persistance

import (
	"github.com/Atelier-Developers/alethia/Database"
	"github.com/Atelier-Developers/alethia/domain/entity/Conversation"
	"log"
)

type ConversationRepository struct {
	dbClient *Database.MySQLDB
}

func NewConversationRepository(dbClient *Database.MySQLDB) *ConversationRepository {
	return &ConversationRepository{dbClient: dbClient}
}

func (conversationRepo *ConversationRepository) ConversationExists(userId1 uint64, userId2 uint64) (bool, error) {
	db := conversationRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT Count(*) FROM USER_CONVERSATION UC1 WHERE user_id=? AND EXISTS (SELECT * FROM USER_CONVERSATION WHERE USER_CONVERSATION.conversation_id = UC1.conversation_id AND USER_CONVERSATION.user_id=?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var count uint64
	row := stmt.QueryRow(userId1, userId2)
	err = row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count > 0, nil
}

func (conversationRepo *ConversationRepository) DoesConversationBelongToUser(conversationId uint64, userId uint64) (bool, error) {
	db := conversationRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT user_id FROM USER_CONVERSATION WHERE user_id=? AND conversation_id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(userId, conversationId)

	var uId uint64
	err = row.Scan(&uId)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (conversationRepo *ConversationRepository) SaveConversation(userId1 uint64, userId2 uint64) error {
	db := conversationRepo.dbClient.GetDB()
	stmt, err := db.Prepare("INSERT INTO CONVERSATION () VALUES () ")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec()

	if err != nil {
		return err
	}

	conversationId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	stmt1, err := db.Prepare("INSERT INTO USER_CONVERSATION (user_id, conversation_id, is_archived, is_deleted, is_read) VALUES (?, ?, 0, 0, 1), (?, ?, 0, 0, 1) ")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt1.Close()

	res, err = stmt1.Exec(userId1, conversationId, userId2, conversationId)

	if err != nil {
		return err
	}

	return nil
}

func (conversationRepo *ConversationRepository) GetConversationCorrespondents(conversationId uint64) (uint64, uint64, error) {
	db := conversationRepo.dbClient.GetDB()

	stmt, err := db.Prepare("SELECT UC1.user_id, UC2.user_id FROM USER_CONVERSATION UC1, USER_CONVERSATION UC2 WHERE UC1.user_id < UC2.user_id AND UC1.conversation_id = UC2.conversation_id AND UC1.conversation_id=?")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	row := stmt.QueryRow(conversationId)

	var userId1, userId2 uint64
	err = row.Scan(&userId1, &userId2)
	if err != nil {
		log.Fatal(err)
	}

	return userId1, userId2, nil
}

func (conversationRepo *ConversationRepository) GetUserConversation(conversationId uint64, userId uint64) (Conversation.UserConversation, error) {
	db := conversationRepo.dbClient.GetDB()

	stmt, err := db.Prepare("SELECT USER_CONVERSATION.*, USER.username FROM USER_CONVERSATION, USER WHERE USER_CONVERSATION.conversation_id=? AND USER_CONVERSATION.user_id=? AND USER.id = USER_CONVERSATION.user_id")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	row := stmt.QueryRow(conversationId, userId)

	var conversation Conversation.UserConversation
	err = row.Scan(&conversation.UserId, &conversation.ConversationId, &conversation.IsArchived, &conversation.IsDeleted, &conversation.IsRead, &conversation.Username)

	if err != nil {
		log.Fatal(err)
	}

	return conversation, err
}

func (conversationRepo *ConversationRepository) UpdateUserConversation(userConversation Conversation.UserConversation) error {

	db := conversationRepo.dbClient.GetDB()

	stmt, err := db.Prepare("UPDATE USER_CONVERSATION SET is_archived = CASE WHEN ? IS NOT NULL THEN ? ELSE is_archived END, is_deleted= CASE WHEN ? IS NOT NULL THEN ? ELSE is_deleted END, is_read= CASE WHEN ? IS NOT NULL THEN ? ELSE is_read END WHERE user_id=? AND conversation_id=?")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(userConversation.IsArchived, userConversation.IsArchived, userConversation.IsDeleted, userConversation.IsDeleted, userConversation.IsRead, userConversation.IsRead, userConversation.UserId, userConversation.ConversationId)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (conversationRepo *ConversationRepository) DeleteConversation(conversation Conversation.Conversation) error {
	db := conversationRepo.dbClient.GetDB()

	stmt2, err := db.Prepare("DELETE FROM USER_CONVERSATION WHERE conversation_id=?")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt2.Close()

	_, err = stmt2.Exec(conversation.Id)

	if err != nil {
		log.Fatal(err)
	}

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

func (conversationRepo *ConversationRepository) GetUserConversations(userId uint64) ([]Conversation.UserConversation, error) {
	db := conversationRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT UC2.user_id, USER.username, UC1.conversation_id, UC1.is_read, UC1.is_deleted, UC1.is_archived FROM USER_CONVERSATION UC1, USER_CONVERSATION UC2, USER WHERE UC1.user_id = ? AND UC1.is_deleted = 0 AND UC1.conversation_id = UC2.conversation_id AND UC1.user_id != UC2.user_id AND UC2.user_id = USER.id")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(userId)
	if err != nil {
		log.Fatal(err)
	}

	var userConversations []Conversation.UserConversation
	for rows.Next() {
		var userConversation Conversation.UserConversation
		err = rows.Scan(&userConversation.UserId, &userConversation.Username, &userConversation.ConversationId, &userConversation.IsRead, &userConversation.IsDeleted, &userConversation.IsArchived)
		if err != nil {
			log.Fatal(err)
		}

		userConversations = append(userConversations, userConversation)
	}

	return userConversations, nil
}
