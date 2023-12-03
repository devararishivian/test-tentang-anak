package persistence

import (
	"context"
	"encoding/json"
	"github.com/devararishivian/test-tentang-anak/internal/domain/entity"
	"github.com/devararishivian/test-tentang-anak/internal/domain/repository"
	"github.com/devararishivian/test-tentang-anak/internal/infrastructure"
)

type MonsterRepositoryImpl struct {
	db *infrastructure.Database
}

func NewMonsterRepository(db *infrastructure.Database) repository.MonsterRepository {
	return &MonsterRepositoryImpl{
		db: db,
	}
}
func (r *MonsterRepositoryImpl) Store(monster *entity.Monster) error {
	var insertUserStmt = `INSERT INTO "monsters"(name, category_name, description, height, weight, stats, image_path) 
							VALUES ($1, $2, $3, $4, $5, $6, $7)`

	monsterStatByte, err := json.Marshal(monster.Stats)
	if err != nil {
		return err
	}

	monsterStat := string(monsterStatByte)

	_, err = r.db.Conn.Exec(
		context.Background(),
		insertUserStmt,
		monster.Name,
		monster.CategoryName,
		monster.Description,
		monster.Height,
		monster.Weight,
		monsterStat,
		monster.ImagePath,
	)
	if err != nil {
		return err
	}

	return nil
}
