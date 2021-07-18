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

func (postRepo *PostRepository) GetPostComments(postId uint64, userId uint64) ([]Comment.CommentWithLikeAndReplyCount, error) {
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

	var comments []Comment.CommentWithLikeAndReplyCount
	for rows.Next() {
		var comment Comment.CommentWithLikeAndReplyCount
		err = rows.Scan(&comment.Id, &comment.Text, &comment.Created, &comment.CommenterId, &comment.PostId, &comment.CommenterUsername, &comment.RepliedCommentId)

		if err != nil {
			log.Fatal(err)
		}

		stmt2, err := db.Prepare("SELECT COUNT(*) FROM COMMENT_LIKE WHERE comment_id=?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt2.Close()

		row := stmt2.QueryRow(comment.Id)
		var count uint64
		err = row.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}

		comment.LikeCount = count

		stmt3, err := db.Prepare("SELECT COUNT(*) FROM REPLY_COMMENT WHERE replied_comment_id=?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt3.Close()

		row = stmt3.QueryRow(comment.Id)
		err = row.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}

		comment.ReplyCount = count

		stmt4, err := db.Prepare("SELECT COUNT(*) FROM COMMENT_LIKE WHERE comment_id=? AND user_id=?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt4.Close()

		row = stmt4.QueryRow(comment.Id, userId)
		err = row.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}

		if count == 1 {
			comment.IsLikedByThisUser = true
		} else {
			comment.IsLikedByThisUser = false
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

func (postRepo *PostRepository) GetPostsByFriends(userId uint64) ([]Post.PostWithLikeAndCommentAndRepostCount, error) {
	db := postRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT POST.*, USER.username FROM POST, USER WHERE USER.id = POST.poster_id AND POST.poster_id IN ((SELECT user2_id FROM FRIEND WHERE user1_id=? ) UNION (SELECT user1_id FROM FRIEND WHERE user2_id=?)) ORDER BY POST.created DESC LIMIT 15")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId, userId)

	if err != nil {
		log.Fatal(err)
	}

	var posts []Post.PostWithLikeAndCommentAndRepostCount
	for rows.Next() {
		var post Post.PostWithLikeAndCommentAndRepostCount
		err = rows.Scan(&post.Id, &post.IsFeatured, &post.Description, &post.Created, &post.PosterId, &post.PosterUsername)

		stmt2, err := db.Prepare("SELECT COUNT(*) FROM POST_LIKE WHERE post_id=?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt2.Close()

		row := stmt2.QueryRow(post.Id)
		var count uint64
		err = row.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}

		post.LikeCount = count

		stmt3, err := db.Prepare("SELECT COUNT(*) FROM COMMENT WHERE post_id=?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt3.Close()

		row = stmt3.QueryRow(post.Id)
		err = row.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}

		post.CommentCount = count

		stmt4, err := db.Prepare("SELECT COUNT(*) FROM REPOST WHERE repost_id=?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt4.Close()

		row = stmt4.QueryRow(post.Id)
		err = row.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}

		post.RepostCount = count

		stmt5, err := db.Prepare("SELECT COUNT(*) FROM POST_LIKE WHERE post_id=? AND user_id=?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt5.Close()

		row = stmt5.QueryRow(post.Id, userId)
		err = row.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}

		if count == 1 {
			post.IsLikedByThisUser = true
		} else {
			post.IsLikedByThisUser = false
		}

		stmt6, err := db.Prepare("SELECT repost_id FROM REPOST WHERE  post_id=?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt6.Close()

		row = stmt6.QueryRow(post.Id)
		var repostId uint64
		err = row.Scan(&repostId)
		if err == nil {
			post.RepostId = repostId
		} else {
			post.RepostId = 0
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (postRepo *PostRepository) GetPostsLikedByFriends(userId uint64) ([]Post.LikedPost, error) {

	db := postRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT POST.*, UC1.username, POST_LIKE.user_id, UC2.username  FROM POST, USER UC1, USER UC2, POST_LIKE WHERE UC1.id = POST.poster_id AND UC2.id = POST_LIKE.user_id AND POST.id = POST_LIKE.post_id AND POST_LIKE.user_id IN ((SELECT user2_id FROM FRIEND WHERE user1_id=? ) UNION (SELECT user1_id FROM FRIEND WHERE user2_id=?)) ORDER BY POST_LIKE.created LIMIT 15")
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

		stmt2, err := db.Prepare("SELECT COUNT(*) FROM POST_LIKE WHERE post_id=?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt2.Close()

		row := stmt2.QueryRow(post.Id)
		var count uint64
		err = row.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}

		post.LikeCount = count

		stmt3, err := db.Prepare("SELECT COUNT(*) FROM COMMENT WHERE post_id=?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt3.Close()

		row = stmt3.QueryRow(post.Id)
		err = row.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}

		post.CommentCount = count

		stmt4, err := db.Prepare("SELECT COUNT(*) FROM REPOST WHERE repost_id=?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt4.Close()

		row = stmt4.QueryRow(post.Id)
		err = row.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}

		post.RepostCount = count

		stmt5, err := db.Prepare("SELECT COUNT(*) FROM POST_LIKE WHERE post_id=? AND user_id=?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt5.Close()

		row = stmt5.QueryRow(post.Id, userId)
		err = row.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}

		if count == 1 {
			post.IsLikedByThisUser = true
		} else {
			post.IsLikedByThisUser = false
		}

		stmt6, err := db.Prepare("SELECT repost_id FROM REPOST WHERE  post_id=?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt6.Close()

		row = stmt6.QueryRow(post.Id)
		var repostId uint64
		err = row.Scan(&repostId)
		if err == nil {
			post.RepostId = repostId
		} else {
			post.RepostId = 0
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (postRepo *PostRepository) GetPostsCommentedOnByFriends(userId uint64) ([]Post.CommentedOnPost, error) {

	db := postRepo.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT POST.*, UC1.username, COMMENT.commenter_id, UC2.username FROM POST, USER UC1, USER UC2, COMMENT WHERE UC1.id = POST.poster_id AND UC2.id = COMMENT.commenter_id AND POST.id = COMMENT.post_id AND COMMENT.commenter_id IN ((SELECT user2_id FROM FRIEND WHERE user1_id=? ) UNION (SELECT user1_id FROM FRIEND WHERE user2_id=?)) ORDER BY COMMENT.created LIMIT 15")
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

		stmt2, err := db.Prepare("SELECT COUNT(*) FROM POST_LIKE WHERE post_id=?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt2.Close()

		row := stmt2.QueryRow(post.Id)
		var count uint64
		err = row.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}

		post.LikeCount = count

		stmt3, err := db.Prepare("SELECT COUNT(*) FROM COMMENT WHERE post_id=?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt3.Close()

		row = stmt3.QueryRow(post.Id)
		err = row.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}

		post.CommentCount = count

		stmt4, err := db.Prepare("SELECT COUNT(*) FROM REPOST WHERE repost_id=?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt4.Close()

		row = stmt4.QueryRow(post.Id)
		err = row.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}

		post.RepostCount = count

		stmt5, err := db.Prepare("SELECT COUNT(*) FROM POST_LIKE WHERE post_id=? AND user_id=?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt5.Close()

		row = stmt5.QueryRow(post.Id, userId)
		err = row.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}

		if count == 1 {
			post.IsLikedByThisUser = true
		} else {
			post.IsLikedByThisUser = false
		}

		stmt6, err := db.Prepare("SELECT repost_id FROM REPOST WHERE  post_id=?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt6.Close()

		row = stmt6.QueryRow(post.Id)
		var repostId uint64
		err = row.Scan(&repostId)
		if err == nil {
			post.RepostId = repostId
		} else {
			post.RepostId = 0
		}

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

	stmt2, err := db.Prepare("SELECT repost_id FROM REPOST WHERE  post_id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt2.Close()

	row = stmt2.QueryRow(post.Id)
	var repostId uint64
	err = row.Scan(&repostId)
	if err == nil {
		post.RepostId = repostId
	} else {
		post.RepostId = 0
	}

	return nil
}
