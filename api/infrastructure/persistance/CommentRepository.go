package persistance

import (
	"github.com/Atelier-Developers/alethia/Database"
	"github.com/Atelier-Developers/alethia/domain/entity"
	"log"
)

type CommentRepository struct {
	dbClient *Database.MySQLDB
}

func NewCommentRepository(dbClient *Database.MySQLDB) *CommentRepository {
	return &CommentRepository{dbClient: dbClient}
}

func (commentRepository *CommentRepository) SaveComment(comment *entity.Comment) error {
	db := commentRepository.dbClient.GetDB()
	stmt, err := db.Prepare("INSERT INTO COMMENT (text, commenter_id, post_id) VALUES (?, ?, ?) ")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(comment.Text, comment.CommenterId, comment.PostId)

	if err != nil {
		return err
	}

	commentId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	comment.Id = uint64(commentId)

	return nil
}

func (commentRepository *CommentRepository) LikeComment(comment *entity.Comment, UserId uint64) error {
	db := commentRepository.dbClient.GetDB()
	stmt, err := db.Prepare("INSERT INTO COMMENT_LIKE (comment_id, user_id) VALUES (?, ?) ")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(comment.Id, UserId)
	if err != nil {
		return err
	}
	return nil
}

func (commentRepository *CommentRepository) ReplyComment(comment *entity.Comment, ReplyId uint64) error {
	db := commentRepository.dbClient.GetDB()
	stmt, err := db.Prepare("INSERT INTO REPLY_COMMENT (comment_id, replied_comment_id) VALUES (?, ?) ")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(comment.Id, ReplyId)
	if err != nil {
		return err
	}
	return nil
}

func (commentRepository *CommentRepository) GetCommentByID(id uint64, comment *entity.Comment) error {
	db := commentRepository.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT * FROM COMMENT WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	err = row.Scan(&comment.Id, &comment.Text, &comment.Created, &comment.CommenterId, &comment.PostId)
	if err != nil {
		return err
	}
	return nil
}
