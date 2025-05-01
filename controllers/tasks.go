package controllers

import (
	"context"
	utils "web_app_fiber/helpers"

	"fmt"
	models "web_app_fiber/models"

	"github.com/gofiber/fiber/v2"
)

func GetTaskHandler(c *fiber.Ctx) error {
	conn := utils.CheckoutConnection()

	selectQuery := "SELECT id, title, description, completed, created_at FROM tasks"
	rows, err := conn.QueryContext(context.Background(), selectQuery)
	if err != nil {
		return c.Status(500).JSON(
			utils.CreateResponsePayload(
				false,
				fmt.Sprintf("Error at GetTaskHandler (query context): %v", err),
				nil))

	}

	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt)
		if err != nil {
			return c.Status(500).JSON(
				utils.CreateResponsePayload(
					false,
					fmt.Sprintf("Error at GetTaskHandler (row scan): %v", err),
					nil))
		}

		tasks = append(tasks, task)
	}
	payload := models.Payload{
		Success:      true,
		Message:      "Succesfully fetched tasks",
		UtcTimestamp: utils.UtilsNowUtc(),
		Data:         tasks,
	}

	return c.JSON(payload)

}

type TaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func AddTaskHandler(c *fiber.Ctx) error {
	var taskReq TaskRequest

	err := c.BodyParser(&taskReq)
	if err != nil {
		return c.Status(500).
			JSON(utils.CreateResponsePayload(false, fmt.Sprintf("Error at AddTaskHandler (request invalid): %v", err), nil))
	}

	conn := utils.CheckoutConnection()

	insertQuery := "INSERT INTO tasks (title, description, completed) VALUES ($1, $2, $3)"
	_, err = conn.ExecContext(context.Background(), insertQuery, taskReq.Title, taskReq.Description, false)
	if err != nil {
		return c.Status(500).
			JSON(utils.CreateResponsePayload(
				false,
				fmt.Sprintf("Error at AddTaskHandler (Insert error): %v", err),
				nil))
	}

	return c.JSON(utils.CreateResponsePayload(true, "Task inserted Successfully", nil))
}
