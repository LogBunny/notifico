package controllers

import (
	"errors"
	"fmt"
	"log"
	"logbun/db"
	"logbun/models"
	"logbun/schemas"
	"logbun/utils"
	"logbun/views"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateUser(c *fiber.Ctx) error {
	userEmail := c.Locals("userEmail").(string)
	userName := c.Locals("userName").(string)
	data, err := db.UsersSvc.FetchProfileByEmail(userEmail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user, err := db.UsersSvc.CreateUser(models.User{Email: userEmail, UserName: userName, Name: userName})
			if err != nil {
				return views.InternalServerError(c, err)
			}
			data = user
		} else {
			return views.InternalServerError(c, err)
		}

	}
	return views.StatusOK(c, data)
}

func UpdateUser(c *fiber.Ctx) error {
	userEmail := c.Locals("userEmail").(string)
	user, err := db.UsersSvc.FetchProfileByEmail(userEmail)
	if err != nil {
		return views.InternalServerError(c, err)
	}
	var userData schemas.UserUpdateSchema
	newUser := models.User{
		ID:         user.ID,
		UserName:   userData.UserName,
		Email:      userData.Email,
		Name:       userData.Name,
		TelegramId: userData.TelegramId,
	}
	if err := c.BodyParser(&userData); err != nil {
		return views.InternalServerError(c, err)
	}

	err = db.UsersSvc.UpdateUser(newUser)
	if err != nil {
		return views.InternalServerError(c, err)
	}

	return views.StatusOK(c, newUser)
}

func GetUser(c *fiber.Ctx) error {
	userEmail := c.Locals("userEmail").(string)
	data, err := db.UsersSvc.FetchProfileByEmail(userEmail)
	if err != nil {

		return views.InternalServerError(c, err)

	}
	return views.StatusOK(c, data)
}

func SendTelegramNotification(c *fiber.Ctx) error {
	var userEmailData schemas.UserEmailSchema

	if err := c.BodyParser(&userEmailData); err != nil {
		return views.InternalServerError(c, err)
	}
	log.Println(userEmailData)
	user, err := db.UsersSvc.FetchProfileByEmail(userEmailData.From)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return views.ForbiddenView(c)
		}
		return views.InternalServerError(c, err)
	}
	message := fmt.Sprintf("New Mail recevied:\nFrom:%s\nSubject:%s\nBody:%s", userEmailData.From, userEmailData.Subject, userEmailData.Body)
	go utils.SendNotificationMessage(user.TelegramId, message)

	return views.StatusOK(c, "ok")
}
