package handler

import (
	"errors"
	"strconv"

	"github.com/Candrandika/be-todo-app-hmdtif/domain/dto"
	"github.com/Candrandika/be-todo-app-hmdtif/domain/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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

func NewTaskHandler(usecase usecase.TaskUsecase, validator *validator.Validate) TaskHandler {
	return &taskHandler{usecase, validator}
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

func (h *taskHandler) Show(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"payload": fiber.Map{
				"error": fiber.Map{
					"code":    fiber.StatusBadRequest,
					"message": "Invalid task ID",
					"error":   err.Error(),
				},
			},
		})
	}

	task, err := h.usecase.GetTaskByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"payload": fiber.Map{
					"error": fiber.Map{
						"code":    fiber.StatusNotFound,
						"message": "Task not found",
						"error":   err.Error(),
					},
				},
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"payload": fiber.Map{
				"error": fiber.Map{
					"code":    fiber.StatusInternalServerError,
					"message": "Failed to get task",
					"error":   err.Error(),
				},
			},
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"payload": fiber.Map{
			"message": "Success get task",
			"task":    task,
		},
	})
}

func (h *taskHandler) Update(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"payload": fiber.Map{
				"error": fiber.Map{
					"code":    fiber.StatusBadRequest,
					"message": "Invalid task ID",
					"error":   err.Error(),
				},
			},
		})
	}

	var req dto.UpdateTaskRequest
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

	task, err := h.usecase.UpdateTask(uint(id), req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"payload": fiber.Map{
					"error": fiber.Map{
						"code":    fiber.StatusNotFound,
						"message": "Task not found",
						"error":   err.Error(),
					},
				},
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"payload": fiber.Map{
				"error": fiber.Map{
					"code":    fiber.StatusInternalServerError,
					"message": "Failed to update task",
					"error":   err.Error(),
				},
			},
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"payload": fiber.Map{
			"message": "Success update task",
			"task":    task,
		},
	})
}

func (h *taskHandler) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"payload": fiber.Map{
				"error": fiber.Map{
					"code":    fiber.StatusBadRequest,
					"message": "Invalid task ID",
					"error":   err.Error(),
				},
			},
		})
	}

	err = h.usecase.DeleteTask(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"payload": fiber.Map{
				"error": fiber.Map{
					"code":    fiber.StatusInternalServerError,
					"message": "Failed to delete task",
					"error":   err.Error(),
				},
			},
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"payload": fiber.Map{
			"message": "Success delete task",
		},
	})
}
