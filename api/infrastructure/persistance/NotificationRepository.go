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

func (notificationRepository *NotificationRepository) CreateEndorseSkillNotification(endorseSkill *notification.EndorseSkill) error {
	db := notificationRepository.dbClient.GetDB()
	err := notificationRepository.creatNotification(&endorseSkill.Notification)

	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("INSERT INTO NOTIFICATION_ENDORSE (notif_id, user_skill_skill_id, user_skill_user_id, user_id) VALUES (?, ?, ? , ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(endorseSkill.Notification.Id, endorseSkill.UserSkillSkillId, endorseSkill.UserSkillUserId, endorseSkill.UserId)

	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (notificationRepository *NotificationRepository) GetEndorseSkillNotification(userId uint64) ([]notification.EndorseSkill, error) {
	db := notificationRepository.dbClient.GetDB()
	var notifs []notification.EndorseSkill
	stmt, err := db.Prepare("SELECT NOTIFICATION_ENDORSE.*, receiver_id, created FROM NOTIFICATION_ENDORSE, NOTIFICATION WHERE NOTIFICATION.id = NOTIFICATION_ENDORSE.notif_id AND NOTIFICATION.receiver_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(userId)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var nc notification.EndorseSkill
		err := rows.Scan(&nc.Id, &nc.UserSkillSkillId, &nc.UserSkillUserId, &nc.UserId, &nc.ReceiverId, &nc.Created)
		if err != nil {
			log.Fatal(err)
		}
		notifs = append(notifs, nc)
	}
	return notifs, nil
}

func (notificationRepository *NotificationRepository) CreateChangeWorkNotification(changeWork *notification.ChangeWork) error {
	db := notificationRepository.dbClient.GetDB()
	err := notificationRepository.creatNotification(&changeWork.Notification)

	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("INSERT INTO NOTIFICATION_CHANGE_WORK (notif_id, user_history_id, type) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(changeWork.Notification.Id, changeWork.UserHistoryId, changeWork.Type)

	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (notificationRepository *NotificationRepository) GetChangeWorkNotification(userId uint64) ([]notification.ChangeWork, error) {
	db := notificationRepository.dbClient.GetDB()
	var notifs []notification.ChangeWork
	stmt, err := db.Prepare("SELECT NOTIFICATION_CHANGE_WORK.*, receiver_id, created FROM NOTIFICATION_CHANGE_WORK, NOTIFICATION WHERE NOTIFICATION.id = NOTIFICATION_CHANGE_WORK.notif_id AND NOTIFICATION.receiver_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(userId)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var nc notification.ChangeWork
		err := rows.Scan(&nc.Id, &nc.UserHistoryId, &nc.Type, &nc.ReceiverId, &nc.Created)
		if err != nil {
			log.Fatal(err)
		}
		notifs = append(notifs, nc)
	}
	return notifs, nil
}

func (notificationRepository *NotificationRepository) CreateLikeCommentNotification(likeComment *notification.LikeComment) error {
	db := notificationRepository.dbClient.GetDB()
	err := notificationRepository.creatNotification(&likeComment.Notification)

	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("INSERT INTO NOTIFICATION_LIKE_COMMENT (notif_id, comment_id, user_id) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(likeComment.Notification.Id, likeComment.CommentId, likeComment.UserId)

	if err != nil {
		log.Fatal(err)
	}
	return nil
}
func (notificationRepository *NotificationRepository) GetLikeCommentNotification(userId uint64) ([]notification.LikeComment, error) {
	db := notificationRepository.dbClient.GetDB()
	var notifs []notification.LikeComment
	stmt, err := db.Prepare("SELECT NOTIFICATION_LIKE_COMMENT.*, receiver_id, created FROM NOTIFICATION_LIKE_COMMENT, NOTIFICATION WHERE NOTIFICATION.id = NOTIFICATION_LIKE_COMMENT.notif_id AND NOTIFICATION.receiver_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(userId)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var nc notification.LikeComment
		err := rows.Scan(&nc.Id, &nc.CommentId, &nc.UserId, &nc.ReceiverId, &nc.Created)
		if err != nil {
			log.Fatal(err)
		}
		notifs = append(notifs, nc)
	}
	return notifs, nil
}
