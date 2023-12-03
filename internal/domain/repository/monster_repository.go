package repository

import "github.com/devararishivian/test-tentang-anak/internal/domain/entity"

type MonsterRepository interface {
	Store(monster *entity.Monster) error
}
