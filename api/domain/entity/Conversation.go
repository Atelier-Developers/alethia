package entity

type Conversation struct {
	Id         uint64 `json:"id"`
	UserId1    uint64 `json:"user_id_1"`
	UserId2    uint64 `json:"user_id_2"`
	IsArchived bool   `json:"is_archived"`
	IsDeleted  bool   `json:"is_deleted"`
	IsRead     bool   `json:"is_read"`
}
