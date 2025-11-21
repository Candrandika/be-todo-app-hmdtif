package usecase

import (
	"github.com/Candrandika/be-todo-app-hmdtif/domain/dto"
	"github.com/Candrandika/be-todo-app-hmdtif/domain/entity"
	"github.com/Candrandika/be-todo-app-hmdtif/domain/repository"
)

type TaskUsecase interface {
	GetAllTask() ([]dto.TaskResponse, error)
	CreateNewTask(req dto.CreateTaskRequest) (*dto.TaskResponse, error)
	GetTaskByID(id uint) (*dto.TaskResponse, error)
	UpdateTask(id uint, req dto.UpdateTaskRequest) (*dto.TaskResponse, error)
	DeleteTask(id uint) error
}

type taskUsecase struct {
	repo repository.TaskRepository
}

func NewTaskUsecase(repo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{repo}
}

func (u *taskUsecase) GetAllTask() ([]dto.TaskResponse, error) {
	tasks, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}

	response := make([]dto.TaskResponse, 0)

	for _, t := range tasks {
		response = append(response, dto.TaskResponse{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			IsDone:      t.IsDone,
			CreatedAt:   t.CreatedAt,
		})
	}
	return (response), nil
}

func (u *taskUsecase) CreateNewTask(req dto.CreateTaskRequest) (*dto.TaskResponse, error) {
	newTask := entity.Task{
		Title:       req.Title,
		Description: req.Description,
	}

	if err := u.repo.CreateNew(&newTask); err != nil {
		return nil, err
	}

	return &dto.TaskResponse{
		ID:          newTask.ID,
		Title:       newTask.Title,
		Description: newTask.Description,
		IsDone:      newTask.IsDone,
		CreatedAt:   newTask.CreatedAt,
	}, nil
}

func (u *taskUsecase) GetTaskByID(id uint) (*dto.TaskResponse, error) {
	task, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		IsDone:      task.IsDone,
		CreatedAt:   task.CreatedAt,
	}, nil
}

func (u *taskUsecase) UpdateTask(id uint, req dto.UpdateTaskRequest) (*dto.TaskResponse, error) {
	task, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Apply fields from request when non-empty/non-nil
	if req.Title != "" {
		task.Title = req.Title
	}
	if req.Description != "" {
		task.Description = req.Description
	}
	if req.IsDone != nil {
		task.IsDone = *req.IsDone
	}

	if err := u.repo.Update(task); err != nil {
		return nil, err
	}

	return &dto.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		IsDone:      task.IsDone,
		CreatedAt:   task.CreatedAt,
	}, nil
}

func (u *taskUsecase) DeleteTask(id uint) error {
	// Verify task exists first
	_, err := u.repo.GetByID(id)
	if err != nil {
		return err
	}
	return u.repo.Delete(id)
}
