package main

import (
	routes "web_app_fiber/router"
	setup "web_app_fiber/setup"

	"github.com/gofiber/fiber/v2"
)


func main() {
	app := fiber.New()
	routes.RegisterRoutes(app)
	go setup.InitSetup()
	app.Listen(":8080")
}
