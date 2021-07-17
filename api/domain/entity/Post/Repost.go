package Post

type Repost struct {
	PostId uint64 `json:"post_id"`
	RepostId uint64 `json:"repost_id"`
}
