package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Handle404(c *fiber.Ctx) error {
	return c.SendFile("./public/404_page.html")
}
