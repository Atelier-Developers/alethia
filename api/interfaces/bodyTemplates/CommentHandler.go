package bodyTemplates

type CommentCreateRequestBody struct {
	Text        string   `json:"text" binding:"required"`
	PostId      uint64 `json:"post_id" binding:"required"`
}

type CommentLikeRequestBody struct {
	CommentId uint64 `json:"comment_id" binding:"required"`
}

type CommentReplyRequestBody struct {
	CommentCreateRequestBody
	RepliedCommentId uint64 `json:"replied_comment_id" binding:"required"`
}
