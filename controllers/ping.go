package controllers

import (
	helpers "web_app_fiber/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func PingHandler(c *fiber.Ctx) error {
	log.Info("Stop hitting me")
	response := helpers.CreateResponsePayload(true, "Go Web Server is alive & running", nil)

	return c.JSON(response)

}
