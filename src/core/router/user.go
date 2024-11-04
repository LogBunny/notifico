package router

import (
	"logbun/controllers"
	"logbun/middlewares"

	"github.com/gofiber/fiber/v2"
)

func userRouter(c fiber.Router) {
	c.Post("/notify", controllers.SendTelegramNotification)
	c.Use(middlewares.CheckAuth)
	c.Get("", controllers.GetUser)
	c.Post("", controllers.CreateUser)
	c.Post("/update", controllers.UpdateUser)
}
