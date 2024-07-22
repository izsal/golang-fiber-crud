package handles

import (
	"my-fiber-app/entities"
	"my-fiber-app/repositories"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TaskHandle struct {
	taskRepository repositories.TaskRepository
}

func NewTaskHandle(taskRepository repositories.TaskRepository) *TaskHandle {
	return &TaskHandle{taskRepository: taskRepository}
}

func (h *TaskHandle) CreateTask(c *fiber.Ctx) error {
	task := new(entities.Task)

	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.taskRepository.CreateTask(task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(task)
}

func (h *TaskHandle) GetTasks(c *fiber.Ctx) error {
	task, err := h.taskRepository.GetAllTasks()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(task)
}

func (h *TaskHandle) GetTasksById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	task, err := h.taskRepository.GetTaskByID(uint(id))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(task)
}

func (h *TaskHandle) UpdateTask(c *fiber.Ctx) error {
	task := new(entities.Task)

	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	id, _ := strconv.Atoi(c.Params("id"))
	task.ID = uint(id)

	if err := h.taskRepository.UpdateTask(task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(task)
}

func (h *TaskHandle) DeleteTasksById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	err := h.taskRepository.DeleteTask(uint(id))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"success": true})
}
