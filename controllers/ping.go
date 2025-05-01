package controllers

import (
	"time"

	models "web_app_fiber/models"

	"github.com/gofiber/fiber/v2"
)

func PingHandler(c *fiber.Ctx) error {

	payload := models.Payload{
		Success:      true,
		Message:      "Successfully received your message",
		UtcTimestamp: time.Now().UTC().Format("2006-01-02 15:04:05 UTC"),
		Data:         nil,
	}

	return c.JSON(payload)

}
