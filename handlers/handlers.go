package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mnsh5/fiber-crud-api/database"
	"github.com/mnsh5/fiber-crud-api/models"
)

func GetTasks(c *fiber.Ctx) error {
	db := database.DB
	var tasks []models.Task
	db.Find(&tasks)
	return c.JSON(tasks)
}
