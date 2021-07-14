package bodyTemplates

type PostCreateRequestBody struct {
	IsFeatured  bool   `json:"is-featured" binding:"required"`
	Description string `json:"description" binding:"required"`
	RepostId    uint64 `json:"repost_id" binding:"required"`
}
