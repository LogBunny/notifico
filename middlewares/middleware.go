package middlewares

import (
	"logbun/utils"
	"logbun/views"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CheckAuth(c *fiber.Ctx) error {
	return CoreAuth(c, false)
}

func CoreAuth(c *fiber.Ctx, isPartialAuth bool) error {
	authToken := strings.TrimSpace(strings.TrimPrefix(c.Get("Authorization"), "Bearer"))
	if authToken == "" {
		if isPartialAuth {
			c.Locals("userEmail", "")
			return c.Next()
		}
		return views.UnAuthorisedView(c)
	}

	claims, err := utils.VerifyFirebaseJWT(authToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Verification failed")
	}

	userEmail := claims["email"].(string)
	var userName = ""
	if claims["name"] != nil {
		userName = claims["name"].(string)
	}
	c.Locals("userEmail", userEmail)
	c.Locals("userName", userName)
	return c.Next()
}
