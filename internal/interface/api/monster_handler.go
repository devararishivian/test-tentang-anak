package api

import (
	"github.com/devararishivian/test-tentang-anak/internal/domain/entity"
	"github.com/devararishivian/test-tentang-anak/internal/domain/service"
	"github.com/devararishivian/test-tentang-anak/internal/interface/validator"
	"github.com/devararishivian/test-tentang-anak/internal/presentation/model"
	"github.com/gofiber/fiber/v2"
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

func (h *MonsterHandler) Store(c *fiber.Ctx) error {
	req := new(model.StoreMonsterRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.CommonResponseWithErrors{
			Message: "bad request",
			Errors:  []any{err.Error()},
		})
	}

	if err := h.validator.Validate(req); err.ValidationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.CommonResponseWithErrors{
			Message: "undefined property",
			Errors:  err.ValidationError,
		})
	}

	monster := new(entity.Monster)
	monster.Name = req.Name
	monster.CategoryName = req.CategoryName
	monster.TypeIDs = req.TypeIDs
	monster.Description = req.Description
	monster.Height = req.Height
	monster.Weight = req.Weight
	monster.Stats = entity.MonsterStat(req.Stats)
	monster.ImagePath = req.ImagePath

	err := h.monsterService.Add(monster)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(model.CommonResponse{Message: "success add monster"})
}

func (h *MonsterHandler) Fetch(c *fiber.Ctx) error {
	req := new(model.FetchMonsterRequest)

	if err := c.QueryParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.CommonResponseWithErrors{
			Message: "bad request",
			Errors:  []any{err.Error()},
		})
	}

	if err := h.validator.Validate(req); err.ValidationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.CommonResponseWithErrors{
			Message: "undefined property",
			Errors:  err.ValidationError,
		})
	}

	entityReq := new(entity.FetchMonstersRequest)
	entityReq.Name = req.Name
	entityReq.TypeIDs = req.TypeIDs
	entityReq.Order = req.Order
	entityReq.Sort = req.Order

	res, err := h.monsterService.List(entityReq)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(model.CommonResponseWithData{
		Message: "success add monster",
		Data:    res,
	})
}
