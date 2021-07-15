package bodyTemplates


type ConversationCreateRequestBody struct {
	UserId2        uint64   `json:"user_id_2" binding:"required"`
	IsArchived     bool  `json:"is_archived" binding:"required"`
	IsDeleted     bool  `json:"is_deleted" binding:"required"`
	IsRead     bool  `json:"is_read" binding:"required"`
}
