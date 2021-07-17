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

type LikedPost struct {
	Id             uint64    `json:"id"`
	IsFeatured     bool      `json:"is-featured"`
	Description    string    `json:"description"`
	Created        time.Time `json:"created"`
	RepostId       uint64    `json:"repost_id"`
	PosterId       uint64    `json:"poster_id"`
	PosterUsername string    `json:"poster_username"`
	LikerId        uint64    `json:"liker_id"`
	LikerUsername   string    `json:"like_username"`
}

type CommentedOnPost struct {
	Id                uint64    `json:"id"`
	IsFeatured        bool      `json:"is-featured"`
	Description       string    `json:"description"`
	Created           time.Time `json:"created"`
	RepostId          uint64    `json:"repost_id"`
	PosterId          uint64    `json:"poster_id"`
	PosterUsername    string    `json:"poster_username"`
	CommenterId       uint64    `json:"commenter_id"`
	CommenterUsername string    `json:"commenter_username"`
}
