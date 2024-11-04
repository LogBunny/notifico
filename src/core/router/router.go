package router

import "github.com/gofiber/fiber/v2"

func MountRoutes(app *fiber.App) {
	app.Get("/api", func(c *fiber.Ctx) error { return c.SendString("Miaow") })
	api := app.Group("/api")
	v1 := api.Group("/v1")
	userRoute := v1.Group("/user")
	userRouter(userRoute)
}
