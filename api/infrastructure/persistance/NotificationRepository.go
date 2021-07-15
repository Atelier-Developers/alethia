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

func (notificationRepository *NotificationRepository) getNotification(notificationId uint64) (notification.Notification, error) {
	db := notificationRepository.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT * FROM NOTIFICATION WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(notificationId)
	var notif notification.Notification
	err = row.Scan(&notif.Id, &notif.ReceiverId, &notif.Created)
	if err != nil {
		log.Fatal(err)
	}
	return notif, nil
}

func (notificationRepository *NotificationRepository) CreateCommentNotification(comment *notification.Comment) error {
	db := notificationRepository.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT * FROM NOTIFICATION_COMMENT")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	err = notificationRepository.creatNotification(&comment.Notification)

	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare("INSERT INTO NOTIFICATION_COMMENT (notif_id, comment_id) VALUES (?, ?)")
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

func (notificationRepository *NotificationRepository) GetCommentNotification(comment *notification.Comment) error {
	//db := notificationRepository.dbClient.GetDB()
	//notif, err := notificationRepository.getNotification(&comment.Notification)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//stmt, err := db.Prepare("INSERT INTO NOTIFICATION_COMMENT (notif_id, comment_id) VALUES (?, ?)")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer stmt.Close()
	//
	//_, err = stmt.Exec(comment.Notification.Id, comment.CommentId)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	return nil
}
