package bodyTemplates

type ConversationCreateRequestBody struct {
	UserId2    uint64 `json:"user_id_2" binding:"required"`
	IsArchived *bool  `json:"is_archived" binding:"required"`
	IsDeleted  *bool  `json:"is_deleted" binding:"required"`
	IsRead     *bool  `json:"is_read" binding:"required"`
}

type ConversationUpdateRequestBody struct {
	Id         uint64 `json:"id" binding:"required"`
	IsArchived *bool  `json:"is_archived" `
	IsDeleted  *bool  `json:"is_deleted" `
	IsRead     *bool  `json:"is_read"`
}

type ConversationDeleteRequestBody struct {
	Id         uint64 `json:"id" binding:"required"`
}

type ConversationGetRequestBody struct {
	UserId1         uint64 `json:"user_id_1" binding:"required"`
	UserId2         uint64 `json:"user_id_2" binding:"required"`
}
