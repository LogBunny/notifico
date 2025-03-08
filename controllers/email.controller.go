package controllers

import (
	"logbun/db"
	"logbun/views"

	"github.com/gofiber/fiber/v2"
)

func GetEmails(c *fiber.Ctx) error {
	email := c.Params("email", "0")
	if email == "0" {
		return views.BadRequestWithMessage(c, "email not provided")
	}

	emails, err := db.EmailSvc.GetEmails(email)
	if err != nil {
		return views.InternalServerError(c, err)
	}

	return views.StatusOK(c, emails)
}
