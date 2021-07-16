package persistance

import (
	"github.com/Atelier-Developers/alethia/Database"
	"github.com/Atelier-Developers/alethia/domain/entity/Post"
	"log"
)

type PostRepository struct {
	dbClient *Database.MySQLDB
}

func NewPostRepository(dbClient *Database.MySQLDB) *PostRepository {
	return &PostRepository{dbClient: dbClient}
}

func (postRepo *PostRepository) SavePost(post *Post.Post) error {
	db := postRepo.dbClient.GetDB()
	stmt, err := db.Prepare("INSERT INTO POST (is_featured, description, created, poster_id) VALUES (?, ?, ?, ?) ")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(post.IsFeatured, post.Description, post.Created, post.PosterId)

	if err != nil {
		log.Fatal(err)
	}

	if post.RepostId != 0 {
		postId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}

		post.Id = uint64(postId)

		stmt, err := db.Prepare("INSERT INTO REPOST (post_id, repost_id) VALUES (?, ?) ")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(post.Id, post.RepostId)

		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func (postRepo *PostRepository) LikePost(post *Post.Post, userId uint64) error {
	db := postRepo.dbClient.GetDB()
	stmt, err := db.Prepare("INSERT INTO POST_LIKE (post_id, user_id) VALUES (?, ?) ")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(post.Id, userId)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

//func (postRepo *PostRepository) GetPostLikes(postId uint64) ([]Post.PostLike, error) {
//	db := postRepo.dbClient.GetDB()
//	stmt, err := db.Prepare("SELECT * FROM POST_LIKE WHERE post_id=?")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer stmt.Close()
//
//	rows, err := stmt.Query(postId)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var likes []Post.PostLike
//	for rows.Next() {
//		var like Post.PostLike
//		err = rows.Scan(&like.PostId, &like.UserId)
//
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		likes = append(likes, like)
//	}
//
//	return likes, nil
//}

func (postRepo *PostRepository) GetPostById(postId uint64, post *Post.Post) error {
	db := postRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT * FROM POST WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(postId)

	err = row.Scan(&post.Id, &post.IsFeatured, &post.Description, &post.Created, &post.PosterId)
	if err != nil {
		return err
	}

	return nil
}
