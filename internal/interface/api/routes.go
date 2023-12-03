package api

import (
	"github.com/devararishivian/test-tentang-anak/internal/application/usecase"
	"github.com/devararishivian/test-tentang-anak/internal/infrastructure"
	"github.com/devararishivian/test-tentang-anak/internal/infrastructure/persistence"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, db *infrastructure.Database) {
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	v1Route := router.Group("v1")

	registerMonsterRoutesV1(v1Route, db)
}

func registerMonsterRoutesV1(router fiber.Router, db *infrastructure.Database) {
	repository := persistence.NewMonsterRepository(db)
	useCase := usecase.NewMonsterUseCase(repository)
	_ = NewMonsterHandler(useCase)

	route := router.Group("monsters")
	route.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendString("OK")
	})
}
