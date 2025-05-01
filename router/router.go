package routes

import (
	"web_app_fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/ping", controllers.PingHandler)
	app.Get("/tasks", controllers.GetTaskHandler)
	app.Post("/tasks", controllers.AddTaskHandler)

}
