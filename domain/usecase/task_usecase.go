package usecase

import (
	"errors"

	"github.com/Candrandika/be-todo-app-hmdtif/domain/dto"
	"github.com/Candrandika/be-todo-app-hmdtif/domain/entity"
	"github.com/Candrandika/be-todo-app-hmdtif/domain/repository"
)

var ErrTaskNotFound = errors.New("task not found")

type TaskUsecase interface {
	GetAllTask() ([]entity.Task, error)
	CreateNewTask(req dto.CreateTaskRequest) (entity.Task, error)
	GetByID(id uint) (entity.Task, error)
	Update(id uint, req dto.TaskUpdateRequest) (entity.Task, error)
	Delete(id uint) error
}

type taskUsecase struct {
	repo repository.TaskRepository
}

func NewTaskUsecase(r repository.TaskRepository) TaskUsecase {
	return &taskUsecase{repo: r}
}

func (u *taskUsecase) GetAllTask() ([]entity.Task, error) {
	return u.repo.GetAll()
}

func (u *taskUsecase) CreateNewTask(req dto.CreateTaskRequest) (entity.Task, error) {
	t := entity.Task{
		Title:       req.Title,
		Description: req.Description,
		Done:        false,
	}
	return u.repo.Create(t)
}

func (u *taskUsecase) GetByID(id uint) (entity.Task, error) {
	t, err := u.repo.GetByID(id)
	if err != nil {
		return entity.Task{}, ErrTaskNotFound
	}
	return t, nil
}

func (u *taskUsecase) Update(id uint, req dto.TaskUpdateRequest) (entity.Task, error) {
	t, err := u.repo.GetByID(id)
	if err != nil {
		return entity.Task{}, ErrTaskNotFound
	}
	t.Title = req.Title
	t.Description = req.Description
	t.Done = req.Done
	updated, err := u.repo.Update(t)
	if err != nil {
		return entity.Task{}, err
	}
	return updated, nil
}

func (u *taskUsecase) Delete(id uint) error {
	_, err := u.repo.GetByID(id)
	if err != nil {
		return ErrTaskNotFound
	}
	return u.repo.Delete(id)
}
