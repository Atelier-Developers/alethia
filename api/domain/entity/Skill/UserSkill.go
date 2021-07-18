package Skill

type UserSkill struct {
	Skill
	EndorseCount         uint64 `json:"endorse_count"`
	IsEndorsedByThisUser bool   `json:"is_endorsed_by_this_user"`
}
