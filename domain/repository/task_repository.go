package repository

import (
	"github.com/Candrandika/be-todo-app-hmdtif/domain/entity"
	"gorm.io/gorm"
	"context"
)

type TaskRepository interface {
	GetAll() ([]entity.Task, error)
	CreateNew(task *entity.Task) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetAll() ([]entity.Task, error) {
	var tasks []entity.Task

	err := r.db.Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) CreateNew(task *entity.Task) error {
	return r.db.Create(task).Error
}
type TaskRepository interface {
	Create(t entity.Task) (entity.Task, error)
	GetAll() ([]entity.Task, error)
	GetByID(id uint) (entity.Task, error)
	Update(t entity.Task) (entity.Task, error)
	Delete(id uint) error
}
