package usecase

import (
	"github.com/Candrandika/be-todo-app-hmdtif/domain/dto"
	"github.com/Candrandika/be-todo-app-hmdtif/domain/entity"
	"github.com/Candrandika/be-todo-app-hmdtif/domain/repository"
	"context"
)

type TaskUsecase interface {
	GetAllTask() ([]dto.TaskResponse, error)
	CreateNewTask(req dto.CreateTaskRequest) (*dto.TaskResponse, error)
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
type TaskUsecase interface {
	GetAll(ctx context.Context) ([]entity.Task, error)
	Create(ctx context.Context, req dto.TaskCreateRequest) (entity.Task, error)
	GetByID(ctx context.Context, id uint) (entity.Task, error)
	Update(ctx context.Context, id uint, req dto.TaskUpdateRequest) (entity.Task, error)
	Delete(ctx context.Context, id uint) error
}
