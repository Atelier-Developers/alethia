package repository

import (
	"github.com/Atelier-Developers/alethia/domain/entity/Skill"
)

type SkillRepository interface {
	IsThisUserSkillEndorsed(userSkillId uint64, userId uint64, endorserId uint64) (bool, error)
	UnEndorseSkill(userSkillId uint64, userId uint64, endorserId uint64) error
	EndorseSkill(userSkillId uint64, userId uint64, endorserId uint64) error
	AddUserSkill(skillId uint64, userId uint64) error
	DeleteUserSkill(skillId uint64, userId uint64) error
	GetUserSkills(visitorId uint64, userId uint64) ([]Skill.UserSkill, error)
	UserSkillExist(skillId uint64, userId uint64) (bool, error)
	GetSkills() ([]Skill.Skill, error)
	GetSkillById(id uint64) (Skill.Skill, error)
}
