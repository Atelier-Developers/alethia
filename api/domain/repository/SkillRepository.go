package repository

import "github.com/Atelier-Developers/alethia/domain/entity"

type SkillRepository interface {
	EndorseSkill(userSkillId uint64, userId uint64, endorserId uint64) error
	AddUserSkill(skillId uint64, userId uint64) error
	DeleteUserSkill(skillId uint64, userId uint64) error
	GetUserSkills(userId uint64) ([]entity.Skill, error)
	UserSkillExist(skillId uint64, userId uint64) (bool, error)
	GetSkills() ([]entity.Skill, error)
	GetSkillById(id uint64) (entity.Skill, error)
}
