package persistence

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/devararishivian/test-tentang-anak/internal/domain/entity"
	"github.com/devararishivian/test-tentang-anak/internal/domain/repository"
	"github.com/devararishivian/test-tentang-anak/internal/infrastructure"
	"strings"
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
	var insertMonsterStmt = `INSERT INTO "monsters"(name, category_name, description, height, weight, stats, image_path) 
							VALUES ($1, $2, $3, $4, $5, $6, $7)`

	monsterStatByte, err := json.Marshal(monster.Stats)
	if err != nil {
		return err
	}

	monsterStat := string(monsterStatByte)

	_, err = r.db.Conn.Exec(
		context.Background(),
		insertMonsterStmt,
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

func (r *MonsterRepositoryImpl) Fetch(req *entity.FetchMonstersRequest) (result *entity.Monsters, err error) {
	var (
		selectMonsterStmt = `SELECT m.id, m.name, m.category_name, m.description, m.height, m.weight, m.stats, m.image_path 
                             FROM "monsters" AS m 
                             WHERE 1=1`
		params []any
		order  = "id"
		sort   = "ASC"
	)

	i := 1
	if req.Name != "" {
		selectMonsterStmt += fmt.Sprintf(" AND m.name = $%d", i)
		params = append(params, req.Name)
		i++
	}

	if len(req.TypeIDs) > 0 {
		selectMonsterStmt += ` JOIN "monster_typeids" AS mt ON m.id = mt.monster_id 
                               WHERE mt.type_id IN(`
		for _, id := range req.TypeIDs {
			selectMonsterStmt += fmt.Sprintf("$%d,", i)
			params = append(params, id)
			i++
		}

		selectMonsterStmt = strings.TrimRight(selectMonsterStmt, ",") + ")"
	}

	if req.Order != "" {
		order = req.Order
	}

	if req.Sort != "" {
		if strings.ToLower(req.Sort) == "desc" {
			sort = "DESC"
		}
	}

	selectMonsterStmt += fmt.Sprintf(" ORDER BY %s %s", order, sort)

	rows, err := r.db.Conn.Query(context.Background(), selectMonsterStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	monsters := make([]*entity.Monster, 0)

	for rows.Next() {
		var monster entity.Monster
		var stats string

		err = rows.Scan(
			&monster.ID,
			&monster.Name,
			&monster.CategoryName,
			&monster.Description,
			&monster.Height,
			&monster.Weight,
			&stats,
			&monster.ImagePath,
		)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal([]byte(stats), &monster.Stats)
		if err != nil {
			return nil, err
		}

		monsters = append(monsters, &monster)
	}

	if len(monsters) == 0 {
		return nil, errors.New("no monsters found")
	}

	return result, nil
}
