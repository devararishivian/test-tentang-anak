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
							VALUES ($1, $2, $3, $4, $5, $6, $7)
							RETURNING id`

	monsterStatByte, err := json.Marshal(monster.Stats)
	if err != nil {
		return err
	}

	monsterStat := string(monsterStatByte)

	tx, err := r.db.Conn.Begin(context.Background())
	if err != nil {
		return err
	}

	var monsterID int
	err = tx.QueryRow(
		context.Background(),
		insertMonsterStmt,
		monster.Name,
		monster.CategoryName,
		monster.Description,
		monster.Height,
		monster.Weight,
		monsterStat,
		monster.ImagePath,
	).Scan(&monsterID)
	if err != nil {
		errRollback := tx.Rollback(context.Background())
		if errRollback != nil {
			return errRollback
		}
		return err
	}

	for _, typeID := range monster.TypeIDs {
		_, errExec := tx.Exec(
			context.Background(),
			`INSERT INTO "monster_typeids" (monster_id, type_id) VALUES ($1, $2)`,
			monsterID,
			typeID,
		)
		if errExec != nil {
			errRollback := tx.Rollback(context.Background())
			if errRollback != nil {
				return errRollback
			}
			return errExec
		}
	}

	return tx.Commit(context.Background())
}

func (r *MonsterRepositoryImpl) Fetch(req *entity.FetchMonstersRequest) (result *entity.Monsters, err error) {
	var (
		selectMonsterStmt = `SELECT m.id, m.name, m.category_name, m.description, m.height, m.weight, m.stats, m.image_path 
                             FROM "monsters" AS m 
                                 JOIN "monster_typeids" AS mt ON m.id = mt.monster_id 
                             WHERE 1=1`
		params []any
		order  = "id"
		sort   = "ASC"
	)

	i := 0
	if req.Name != "" {
		selectMonsterStmt += fmt.Sprintf(" AND m.name = $%d", i)
		params = append(params, req.Name)
		i++
	}

	if req.TypeIDs != nil {
		if len(*req.TypeIDs) > 0 {
			for index, id := range *req.TypeIDs {
				if id != 0 {
					if index == 0 {
						selectMonsterStmt += ` AND mt.type_id IN(`
					}
					selectMonsterStmt += fmt.Sprintf("$%d,", i)
					params = append(params, id)
					i++

					selectMonsterStmt = strings.TrimRight(selectMonsterStmt, ",") + ")"
				}
			}

		}
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

	monsters := make(entity.Monsters, 0)

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

		monsters = append(monsters, monster)
	}

	if len(monsters) == 0 {
		return nil, errors.New("no monsters found")
	}

	return &monsters, nil
}
