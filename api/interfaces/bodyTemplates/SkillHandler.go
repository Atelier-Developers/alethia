package bodyTemplates

type EndorseSkillRequestBody struct {
	SkillId uint64 `json:"skill_id"  binding:"required"`
	UserId  uint64 `json:"user_id"  binding:"required"`
}

type AddUserSkillRequestBody struct {
	SkillId uint64 `json:"skill_id"  binding:"required" `
}

type GetUserSkillRequestBody struct {
	UserId uint64 `json:"user_id" binding:"request"`
}

type DeleteUserSkillRequestBody struct {
	SkillId uint64 `json:"skill_id"  binding:"required" `
}
