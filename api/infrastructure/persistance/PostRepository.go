package persistance

import (
	"github.com/Atelier-Developers/alethia/Database"
	"github.com/Atelier-Developers/alethia/domain/entity/Comment"
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

func (postRepo *PostRepository) GetPostLikes(postId uint64) ([]Post.Like, error) {
	db := postRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT POST_LIKE.*, USER.username FROM POST_LIKE, USER WHERE post_id=? AND POST_LIKE.user_id=USER.id")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(postId)
	if err != nil {
		log.Fatal(err)
	}

	var likes []Post.Like
	for rows.Next() {
		var like Post.Like
		err = rows.Scan(&like.PostId, &like.UserId, &like.Created, &like.Username)

		if err != nil {
			log.Fatal(err)
		}

		likes = append(likes, like)
	}

	return likes, nil
}

func (postRepo *PostRepository) GetPostComments(postId uint64) ([]Comment.Comment, error) {
	db := postRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT COMMENT.*, U.username, RC.replied_comment_id FROM COMMENT LEFT OUTER JOIN REPLY_COMMENT RC on COMMENT.id = RC.comment_id INNER JOIN USER U on COMMENT.commenter_id = U.id WHERE COMMENT.post_id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(postId)
	if err != nil {
		log.Fatal(err)
	}

	var comments []Comment.Comment
	for rows.Next() {
		var comment Comment.Comment
		err = rows.Scan(&comment.Id, &comment.Text, &comment.Created, &comment.CommenterId, &comment.PostId, &comment.CommenterUsername, &comment.RepliedCommentId)

		if err != nil {
			log.Fatal(err)
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func (postRepo *PostRepository) GetPostReposts(postId uint64) ([]Post.Repost, error) {
	db := postRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT * FROM REPOST WHERE repost_id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(postId)
	if err != nil {
		log.Fatal(err)
	}

	var reposts []Post.Repost
	for rows.Next() {
		var repost Post.Repost
		err = rows.Scan(&repost.PostId, &repost.RepostId)

		if err != nil {
			log.Fatal(err)
		}

		reposts = append(reposts, repost)
	}

	return reposts, nil
}

func (postRepo *PostRepository) GetPostsByFriends(userId uint64) ([]Post.Post, error) {
	db := postRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT POST.*, USER.username FROM POST, USER WHERE USER.id = POST.poster_id AND POST.poster_id IN ((SELECT user2_id FROM FRIEND WHERE user1_id=? ) UNION (SELECT user1_id FROM FRIEND WHERE user2_id=?)) ORDER BY POST.created DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId, userId)

	if err != nil {
		log.Fatal(err)
	}

	var posts []Post.Post
	for rows.Next() {
		var post Post.Post
		err = rows.Scan(&post.Id, &post.IsFeatured, &post.Description, &post.Created, &post.PosterId, &post.PosterUsername)

		posts = append(posts, post)
	}

	return posts, nil
}

func (postRepo *PostRepository) GetPostsLikedByFriends(userId uint64) ([]Post.LikedPost, error) {

	db := postRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT POST.*, UC1.username, POST_LIKE.user_id, UC2.username  FROM POST, USER UC1, USER UC2, POST_LIKE WHERE UC1.id = POST.poster_id AND UC2.id = POST_LIKE.user_id AND POST.id = POST_LIKE.post_id AND POST_LIKE.user_id IN ((SELECT user2_id FROM FRIEND WHERE user1_id=? ) UNION (SELECT user1_id FROM FRIEND WHERE user2_id=?)) ORDER BY POST_LIKE.created")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId, userId)

	if err != nil {
		log.Fatal(err)
	}

	var posts []Post.LikedPost
	for rows.Next() {
		var post Post.LikedPost
		err = rows.Scan(&post.Id, &post.IsFeatured, &post.Description, &post.Created, &post.PosterId, &post.PosterUsername, &post.LikerId, &post.LikerUsername)

		posts = append(posts, post)
	}

	return posts, nil
}

func (postRepo *PostRepository) GetPostsCommentedOnByFriends(userId uint64) ([]Post.CommentedOnPost, error) {

	db := postRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT POST.*, UC1.username, COMMENT.commenter_id, UC2.username FROM POST, USER UC1, USER UC2, COMMENT WHERE UC1.id = POST.poster_id AND UC2.id = COMMENT.commenter_id AND POST.id = COMMENT.post_id AND COMMENT.commenter_id IN ((SELECT user2_id FROM FRIEND WHERE user1_id=? ) UNION (SELECT user1_id FROM FRIEND WHERE user2_id=?)) ORDER BY COMMENT.created")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId, userId)

	if err != nil {
		log.Fatal(err)
	}

	var posts []Post.CommentedOnPost
	for rows.Next() {
		var post Post.CommentedOnPost
		err = rows.Scan(&post.Id, &post.IsFeatured, &post.Description, &post.Created, &post.PosterId, &post.PosterUsername, &post.CommenterId, &post.CommenterUsername)

		posts = append(posts, post)
	}

	return posts, nil
}

func (postRepo *PostRepository) GetPostById(postId uint64, post *Post.Post) error {
	db := postRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT POST.*, USER.username FROM POST, USER WHERE POST.id=? AND USER.id = POST.poster_id")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(postId)

	err = row.Scan(&post.Id, &post.IsFeatured, &post.Description, &post.Created, &post.PosterId, &post.PosterUsername)
	if err != nil {
		return err
	}

	return nil
}
