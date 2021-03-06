package repository

import (
	"github.com/Atelier-Developers/alethia/domain/entity/Comment"
	"github.com/Atelier-Developers/alethia/domain/entity/Post"
)

type PostRepository interface {
	SavePost(post *Post.Post) error
	LikePost(post *Post.Post, userId uint64) error
	UnlikePost(userId uint64, postId uint64) error
	IsPostLikedByThisUser(userId uint64, postId uint64) (bool, error)
	GetPostById(postId uint64, post *Post.Post) error
	GetPostLikes(postId uint64) ([]Post.Like, error)
	GetPostComments(postId uint64, userId uint64) ([]Comment.CommentWithLikeAndReplyCount, error)
	GetPostReposts(postId uint64) ([]Post.Repost, error)
	GetPostsByFriends(userId uint64) ([]Post.PostWithLikeAndCommentAndRepostCount, error)
	GetPostsLikedByFriends(userId uint64) ([]Post.LikedPost, error)
	GetPostsCommentedOnByFriends(userId uint64) ([]Post.CommentedOnPost, error)
}
