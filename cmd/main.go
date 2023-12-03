package main

import (
	"errors"
	"github.com/devararishivian/test-tentang-anak/internal/config"
	"github.com/devararishivian/test-tentang-anak/internal/infrastructure"
	"github.com/devararishivian/test-tentang-anak/internal/interface/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

var fiberConfig = fiber.Config{
	// Override default error handler
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		// Status code defaults to 500
		code := fiber.StatusInternalServerError

		// Retrieve the custom status code if it's a *fiber.Error
		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		return ctx.Status(code).JSON(fiber.Map{
			"message": e.Message,
		})
	},
}

func main() {
	errConfig := config.LoadConfig("./internal/config/config.json")
	if errConfig != nil {
		panic(errConfig)
	}

	db, errDB := infrastructure.NewDatabase()
	if errDB != nil {
		panic(errDB)
	}

	app := fiber.New(fiberConfig)
	app.Use(logger.New())

	api.RegisterRoutes(app.Group("api"), db)

	log.Fatal(app.Listen(config.Server.Address))
}
