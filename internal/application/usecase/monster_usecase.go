package usecase

import (
	"github.com/devararishivian/test-tentang-anak/internal/domain/entity"
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

func (u *MonsterUseCaseImpl) Add(monster *entity.Monster) error {
	err := u.monsterRepository.Store(monster)
	if err != nil {
		return err
	}

	return nil
}
