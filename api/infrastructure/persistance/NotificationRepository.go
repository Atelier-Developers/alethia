package persistance

import (
	"github.com/Atelier-Developers/alethia/Database"
	"github.com/Atelier-Developers/alethia/domain/entity/notification"
	"log"
)

type NotificationRepository struct {
	dbClient *Database.MySQLDB
}

func NewNotificationRepository(dbClient *Database.MySQLDB) *NotificationRepository {
	return &NotificationRepository{dbClient: dbClient}
}

func (notificationRepository *NotificationRepository) creatNotification(notification *notification.Notification) error {
	db := notificationRepository.dbClient.GetDB()
	stmt, err := db.Prepare("INSERT INTO NOTIFICATION (receiver_id) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(notification.ReceiverId)

	if err != nil {
		log.Fatal(err)
	}

	notifId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	notification.Id = uint64(notifId)
	return nil
}


func (notificationRepository *NotificationRepository) CreateCommentNotification(comment *notification.Comment) error {
	db := notificationRepository.dbClient.GetDB()
	err := notificationRepository.creatNotification(&comment.Notification)

	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("INSERT INTO NOTIFICATION_COMMENT (notif_id, comment_id) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(comment.Notification.Id, comment.CommentId)

	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (notificationRepository *NotificationRepository) GetCommentNotification(userId uint64) ([]notification.Comment, error) {
	db := notificationRepository.dbClient.GetDB()
	var notifs []notification.Comment
	stmt, err := db.Prepare("SELECT NOTIFICATION_COMMENT.*, receiver_id, created FROM NOTIFICATION_COMMENT, NOTIFICATION WHERE NOTIFICATION.id = NOTIFICATION_COMMENT.notif_id AND NOTIFICATION.receiver_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(userId)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var nc notification.Comment
		err := rows.Scan(&nc.Id, &nc.CommentId, &nc.ReceiverId, &nc.Created)
		if err != nil {
			log.Fatal(err)
		}
		notifs = append(notifs, nc)
	}
	return notifs, nil
}
