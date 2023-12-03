package api

import (
	"github.com/devararishivian/test-tentang-anak/internal/domain/service"
	"github.com/devararishivian/test-tentang-anak/internal/interface/validator"
)

type MonsterHandler struct {
	monsterService service.MonsterService
	validator      *validator.Validator
}

func NewMonsterHandler(useCase service.MonsterService) MonsterHandler {
	return MonsterHandler{
		monsterService: useCase,
		validator:      validator.NewValidator(),
	}
}
