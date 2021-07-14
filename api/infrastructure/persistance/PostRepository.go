package persistance

import (
	"github.com/Atelier-Developers/alethia/Database"
	"github.com/Atelier-Developers/alethia/domain/entity"
	"log"
)

type PostRepository struct {
	dbClient *Database.MySQLDB
}

func NewPostRepository(dbClient *Database.MySQLDB) *PostRepository {
	return &PostRepository{dbClient: dbClient}
}

func (postRepo *PostRepository) SavePost(post *entity.Post) error {
	db := postRepo.dbClient.GetDB()
	stmt, err := db.Prepare("INSERT INTO POST (is_featured, description, created, repost_id, poster_id) VALUES (?, ?, ?, ?, ?) ")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(post.IsFeatured, post.Description, post.Created, post.RepostId, post.PosterId)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
