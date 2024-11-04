package main

import (
	"logbun/config"
	"logbun/db"
	"logbun/migrations"
	"logbun/src/core/router"
	"logbun/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	utils.ImportEnv()
	config.LoadCfg()

	db.InitServices()

	if config.MIGRATE {
		migrations.Migrate()
	}

	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Next: func(c *fiber.Ctx) bool {
			return strings.HasPrefix(c.Path(), "api")
		},
	}))

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	router.MountRoutes(app)

	app.Listen(":3000")
	//test
}
