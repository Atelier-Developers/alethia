package bodyTemplates

type InviteCreateRequestBody struct {
	ReceiverId uint64 `json:"receiver_id" binding:"required"`
	Body       string `json:"body" binding:"required"`
}

type FriendCreateRequestBody struct {
	UserId uint64 `json:"user_id" binding:"required"`
}

type InviteDeleteRequestBody struct {
	UserId uint64 `json:"user_id" binding:"required"`
}

type FriendDeleteRequestBody struct {
	UserId uint64 `json:"user_id" binding:"required"`
}

type FriendGetRequestBody struct {
	UserId uint64 `json:"user_id" binding:"required"`
}
