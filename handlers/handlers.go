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

func CreateTask(c *fiber.Ctx) error {
	db := database.DB
	task := new(models.Task)

	if err := c.BodyParser(task); err != nil {
		return c.Status(400).JSON(err)
	}

	db.Create(&task)
	return c.JSON(task)
}

func GetTask(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var task models.Task
	db.Find(&task, id)
	return c.JSON(task)
}

func EditTask(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var task models.Task

	if err := db.First(&task, id).Error; err != nil {
		return c.Status(404).JSON(err)
	}

	updateTask := new(models.Task)
	if err := c.BodyParser(updateTask); err != nil {
		return c.Status(400).JSON(err)
	}

	task.Title = updateTask.Title
	task.Description = updateTask.Description
	db.Save(&task)

	return c.JSON(task)

}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var task models.Task

	if err := db.First(&task, id).Error; err != nil {
		return c.Status(404).JSON(err)
	}

	db.Delete(&task)
	return c.JSON(fiber.Map{"status": "success", "message": "Task deleted", "id": task.ID})
}
