package service

import "github.com/devararishivian/test-tentang-anak/internal/domain/entity"

type MonsterService interface {
	Add(monster *entity.Monster) error
	List(req *entity.FetchMonstersRequest) (*entity.Monsters, error)
}
