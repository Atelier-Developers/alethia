package bodyTemplates

type UserLanguageCreateRequestBody struct {
	LanguageId uint64 `json:"language_id"  binding:"required"`
}

type UserLanguageDeleteRequestBody struct {
	LanguageId uint64 `json:"language_id" binding:"required"`
}

type UserLanguageGetRequestBody struct {
	UserId uint64 `json:"user_id" binding:"required"`
}
