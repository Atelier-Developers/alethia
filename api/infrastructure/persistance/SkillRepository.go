package persistance

import (
	"github.com/Atelier-Developers/alethia/Database"
	"github.com/Atelier-Developers/alethia/domain/entity"
	"log"
)

type SkillRepository struct {
	dbClient *Database.MySQLDB
}

func NewSkillRepository(dbClient *Database.MySQLDB) *SkillRepository {
	return &SkillRepository{dbClient: dbClient}
}

func (skillRepository *SkillRepository) EndorseSkill(userSkillId uint64, userId uint64, endorserId uint64) error {
	db := skillRepository.dbClient.GetDB()
	stmt, err := db.Prepare("INSERT INTO ENDORSE (user_skill_skill_id, user_skill_user_id, endorser_id) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(userSkillId, userId, endorserId)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
func (skillRepository *SkillRepository) AddUserSkill(skillId uint64, userId uint64) error {
	db := skillRepository.dbClient.GetDB()
	stmt, err := db.Prepare("INSERT INTO USER_SKILL (skill_id, user_id) VALUES (?, ?) ")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(skillId, userId)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (skillRepository *SkillRepository) GetSkillById(id uint64) (entity.Skill, error) {
	db := skillRepository.dbClient.GetDB()

	stmt, err := db.Prepare("SELECT * FROM SKILL WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	var skill entity.Skill
	err = row.Scan(&skill.Id, &skill.Title, &skill.Category)

	if err != nil {
		log.Fatal(err)
	}

	return skill, nil
}

func (skillRepository *SkillRepository) GetSkills() ([]entity.Skill, error) {
	db := skillRepository.dbClient.GetDB()

	stmt, err := db.Prepare("SELECT * FROM SKILL")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	var skills []entity.Skill
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var skill entity.Skill
		err = rows.Scan(&skill.Id, &skill.Title, &skill.Category)
		if err != nil {
			log.Fatal(err)
		}

		skills = append(skills, skill)
	}

	return skills, nil
}

func (skillRepository *SkillRepository) UserSkillExist(skillId uint64, userId uint64) (bool, error) {
	db := skillRepository.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT COUNT(*) FROM USER_SKILL WHERE skill_id = ? AND user_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var count uint64
	row := stmt.QueryRow(skillId, userId)
	err = row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count > 0, nil
}

//TODO Duplicate User skill

func (skillRepository *SkillRepository) DeleteUserSkill(skillId uint64, userId uint64) error {
	db := skillRepository.dbClient.GetDB()
	stmt, err := db.Prepare("DELETE FROM USER_SKILL WHERE skill_id=? AND user_id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(skillId, userId)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (skillRepository *SkillRepository) GetUserSkills(userId uint64) ([]entity.Skill, error) {
	db := skillRepository.dbClient.GetDB()
	stmt, err := db.Prepare("SELECT SKILL.* FROM SKILL, USER_SKILL WHERE SKILL.id = USER_SKILL.skill_id AND USER_SKILL.user_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId)
	if err != nil {
		log.Fatal(err)
	}

	var skills []entity.Skill
	for rows.Next() {
		var skill entity.Skill
		if err := rows.Scan(&skill.Id, &skill.Title, &skill.Category); err != nil {
			log.Fatal(err)
		}
		skills = append(skills, skill)
	}

	return skills, nil
}
