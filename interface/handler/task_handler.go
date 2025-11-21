package handler

import (
	"strconv"

	"github.com/Candrandika/be-todo-app-hmdtif/domain/dto"
	"github.com/Candrandika/be-todo-app-hmdtif/domain/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type TaskHandler interface {
	Index(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Show(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type taskHandler struct {
	usecase   usecase.TaskUsecase
	validator *validator.Validate
}

func NewTaskHandler(u usecase.TaskUsecase, v *validator.Validate) TaskHandler {
	return &taskHandler{usecase: u, validator: v}
}

func (h *taskHandler) Index(ctx *fiber.Ctx) error {
	tasks, err := h.usecase.GetAllTask()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"payload": fiber.Map{
				"code":    fiber.StatusInternalServerError,
				"message": "Failed to get all tasks",
				"errors":  err.Error(),
			},
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"payload": fiber.Map{
			"message": "Success get all tasks",
			"tasks":   tasks,
		},
	})
}

func (h *taskHandler) Create(ctx *fiber.Ctx) error {
	var req dto.CreateTaskRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"payload": fiber.Map{
				"error": fiber.Map{
					"code":    fiber.StatusBadRequest,
					"message": "Invalid request",
					"error":   err.Error(),
				},
			},
		})
	}

	if err := h.validator.Struct(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"payload": fiber.Map{
				"error": fiber.Map{
					"code":    fiber.StatusBadRequest,
					"message": "Invalid request",
					"error":   err.Error(),
				},
			},
		})
	}

	newTask, err := h.usecase.CreateNewTask(req)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"payload": fiber.Map{
				"error": fiber.Map{
					"code":    fiber.StatusInternalServerError,
					"message": "Invalid request",
					"error":   err.Error(),
				},
			},
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"payload": fiber.Map{
			"message": "Success create new task",
			"task":    newTask,
		},
	})
}

func (h *taskHandler) Show(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"payload": fiber.Map{
				"error": fiber.Map{
					"code":    fiber.StatusBadRequest,
					"message": "invalid id",
				},
			},
		})
	}

	task, err := h.usecase.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"payload": fiber.Map{
				"error": fiber.Map{
					"code":    fiber.StatusNotFound,
					"message": "task not found",
				},
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"payload": fiber.Map{
			"message": "Success get task",
			"task":    task,
		},
	})
}

func (h *taskHandler) Update(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"payload": fiber.Map{
				"error": fiber.Map{
					"code":    fiber.StatusBadRequest,
					"message": "invalid id",
				},
			},
		})
	}

	var req dto.TaskUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"payload": fiber.Map{
				"error": fiber.Map{
					"code":    fiber.StatusBadRequest,
					"message": "invalid body",
					"error":   err.Error(),
				},
			},
		})
	}

	if err := h.validator.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"payload": fiber.Map{
				"error": fiber.Map{
					"code":    fiber.StatusBadRequest,
					"message": "validation failed",
					"error":   err.Error(),
				},
			},
		})
	}

	updated, err := h.usecase.Update(uint(id), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"payload": fiber.Map{
				"error": fiber.Map{
					"code":    fiber.StatusInternalServerError,
					"message": "unable to update",
					"error":   err.Error(),
				},
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"payload": fiber.Map{
			"message": "Success update task",
			"task":    updated,
		},
	})
}

func (h *taskHandler) Delete(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"payload": fiber.Map{
				"error": fiber.Map{
					"code":    fiber.StatusBadRequest,
					"message": "invalid id",
				},
			},
		})
	}

	if err := h.usecase.Delete(uint(id)); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"payload": fiber.Map{
				"error": fiber.Map{
					"code":    fiber.StatusNotFound,
					"message": "task not found",
				},
			},
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
