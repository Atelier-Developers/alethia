package bodyTemplates

type PostCreateRequestBody struct {
	Description string `json:"description" binding:"required"`
	RepostId    uint64 `json:"repost_id" `
}
