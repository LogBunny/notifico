package main

import (
	"log"
	"logbun/config"
	"logbun/db"
	"logbun/migrations"
	"logbun/src/core/router"
	"logbun/utils"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func deleteOldMails() {
	ticker := time.NewTicker(12 * time.Hour)
	defer ticker.Stop()

	for {
		<-ticker.C // Wait for the next tick
		err := db.EmailSvc.DeleteEmails()
		if err != nil {
			log.Printf("Failed to delete old mails: %v", err)
		} else {
			log.Println("Old mails deleted successfully.")
		}
	}
}

func main() {
	utils.ImportEnv()
	config.LoadCfg()

	db.InitServices()

	if config.MIGRATE {
		migrations.Migrate()
	}

	go deleteOldMails()

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
