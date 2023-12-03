package persistence

import (
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
