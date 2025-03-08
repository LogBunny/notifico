package router

import (
	"logbun/controllers"

	"github.com/gofiber/fiber/v2"
)

func emailRouter(c fiber.Router) {
	c.Get("/:email", controllers.GetEmails)
}
