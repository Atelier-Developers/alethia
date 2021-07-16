package persistance

import (
	"github.com/Atelier-Developers/alethia/Database"
	"github.com/Atelier-Developers/alethia/domain/entity/Conversation"
	"log"
)

type MessageRepository struct {
	dbClient *Database.MySQLDB
}

func NewMessageRepository(dbClient *Database.MySQLDB) *MessageRepository {
	return &MessageRepository{dbClient: dbClient}
}

func (messageRepository *MessageRepository) AddMessage(message Conversation.Message) error {
	db := messageRepository.dbClient.GetDB()
	stmt, err := db.Prepare("INSERT INTO MESSAGE (user_id, reply_id, conversation_id, body) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(message.UserId, message.ReplyId, message.ConversationId, message.Body)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (messageRepository *MessageRepository) GetMessage(messageID uint64) (Conversation.Message, error) {
	db := messageRepository.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT * FROM MESSAGE WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var message Conversation.Message
	row := stmt.QueryRow(messageID)

	err = row.Scan(&message.Id, &message.UserId, &message.ReplyId, &message.ConversationId, &message.Body, &message.Created)
	if err != nil {
		log.Fatal(err)
	}

	return message, nil
}

func (messageRepository *MessageRepository) GetMessages(conversationId uint64) ([]Conversation.Message, error) {
	db := messageRepository.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT * FROM MESSAGE WHERE conversation_id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var messages []Conversation.Message
	rows, err := stmt.Query(conversationId)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var message Conversation.Message
		err := rows.Scan(&message.Id, &message.UserId, &message.ReplyId, &message.ConversationId, &message.Body, &message.Created)
		if err != nil {
			log.Fatal(err)
		}

		messages = append(messages, message)
	}
	return messages, nil
}
