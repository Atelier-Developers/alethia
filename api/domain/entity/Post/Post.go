package Post

import "time"

type Post struct {
	Id             uint64    `json:"id"`
	IsFeatured     bool      `json:"is-featured"`
	Description    string    `json:"description"`
	Created        time.Time `json:"created"`
	RepostId       uint64    `json:"repost_id"`
	PosterId       uint64    `json:"poster_id"`
	PosterUsername string    `json:"poster_username"`
}

type PostWithLikeAndCommentAndRepostCount struct {
	Id                uint64    `json:"id"`
	IsFeatured        bool      `json:"is-featured"`
	Description       string    `json:"description"`
	Created           time.Time `json:"created"`
	RepostId          uint64    `json:"repost_id"`
	PosterId          uint64    `json:"poster_id"`
	PosterUsername    string    `json:"poster_username"`
	LikeCount         uint64    `json:"like_count"`
	CommentCount      uint64    `json:"comment_count"`
	RepostCount       uint64    `json:"repost_count"`
	IsLikedByThisUser bool      `json:"is_liked_by_this_user"`
}

type LikedPost struct {
	PostWithLikeAndCommentAndRepostCount
	LikerId       uint64 `json:"liker_id"`
	LikerUsername string `json:"like_username"`
}

type CommentedOnPost struct {
	PostWithLikeAndCommentAndRepostCount
	CommenterId       uint64 `json:"commenter_id"`
	CommenterUsername string `json:"commenter_username"`
}
