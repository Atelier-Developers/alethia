package notification

type EndorseSkill struct {
	UserId           uint64 `json:"user_id"`
	UserSkillUserId  uint64 `json:"user_skill_user_id"`
	UserSkillSkillId uint64 `json:"user_skill_skill_id"`
	Notification
}
