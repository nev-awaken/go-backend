package routes

import (
	"web_app_fiber/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func RegisterRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/ping", controllers.PingHandler)
	app.Get("/tasks", controllers.GetTaskHandler)
	app.Post("/tasks", controllers.AddTaskHandler)
	app.Use(controllers.Handle404)

}
