package persistance

import (
	"github.com/Atelier-Developers/alethia/Database"
	"github.com/Atelier-Developers/alethia/domain/entity/Comment"
	"log"
)

type CommentRepository struct {
	dbClient *Database.MySQLDB
}

func NewCommentRepository(dbClient *Database.MySQLDB) *CommentRepository {
	return &CommentRepository{dbClient: dbClient}
}

func (commentRepository *CommentRepository) SaveComment(comment *Comment.Comment) error {
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

func (commentRepository *CommentRepository) LikeComment(comment *Comment.Comment, UserId uint64) error {
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

func (commentRepository *CommentRepository) ReplyComment(comment *Comment.Comment, ReplyId uint64) (uint64, error) {
	db := commentRepository.dbClient.GetDB()
	stmt, err := db.Prepare("INSERT INTO REPLY_COMMENT (comment_id, replied_comment_id) VALUES (?, ?) ")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(comment.Id, ReplyId)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

func (commentRepository *CommentRepository) GetCommentLikes(commentId uint64) ([]Comment.CommentLike, error) {
	db := commentRepository.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT * FROM COMMENT_LIKE WHERE COMMENT_LIKE.comment_id=? ORDER BY created DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(commentId)
	if err != nil {
		log.Fatal(err)
	}

	var commentLikes []Comment.CommentLike
	for rows.Next() {
		var commentLike Comment.CommentLike
		err = rows.Scan(&commentLike.CommentId, &commentLike.UserId, &commentLike.Created)
		if err != nil {
			log.Fatal(err)
		}

		commentLikes = append(commentLikes, commentLike)
	}

	return commentLikes, nil
}

func (commentRepository *CommentRepository) GetCommentByID(id uint64, comment *Comment.Comment) error {
	db := commentRepository.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT COMMENT.*, U.username, RC.replied_comment_id FROM COMMENT LEFT OUTER JOIN REPLY_COMMENT RC on COMMENT.id = RC.comment_id INNER JOIN USER U on COMMENT.commenter_id = U.id WHERE COMMENT.id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	err = row.Scan(&comment.Id, &comment.Text, &comment.Created, &comment.CommenterId, &comment.PostId, &comment.CommenterUsername, &comment.RepliedCommentId)
	if err != nil {
		return err
	}
	return nil
}
