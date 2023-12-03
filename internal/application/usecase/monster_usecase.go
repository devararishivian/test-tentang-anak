package usecase

import (
	"github.com/devararishivian/test-tentang-anak/internal/domain/repository"
)

type MonsterUseCaseImpl struct {
	monsterRepository repository.MonsterRepository
}

func NewMonsterUseCase(monsterRepository repository.MonsterRepository) *MonsterUseCaseImpl {
	return &MonsterUseCaseImpl{
		monsterRepository: monsterRepository,
	}
}
