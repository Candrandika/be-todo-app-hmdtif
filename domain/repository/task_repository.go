package repository

import (
	"errors"

	"github.com/Candrandika/be-todo-app-hmdtif/domain/entity"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *entity.Task) (entity.Task, error)
	GetAll() ([]entity.Task, error)
	GetByID(id uint) (entity.Task, error)
	Update(task entity.Task) (entity.Task, error)
	Delete(id uint) error
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

func (r *taskRepository) Create(task *entity.Task) (entity.Task, error) {
	return *task, r.db.Create(task).Error
}

func (r *taskRepository) GetByID(id uint) (entity.Task, error) {
	var t entity.Task
	if err := r.db.First(&t, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Task{}, err
		}
		return entity.Task{}, err
	}
	return t, nil
}

func (r *taskRepository) Update(task entity.Task) (entity.Task, error) {
	if err := r.db.Save(&task).Error; err != nil {
		return entity.Task{}, err
	}
	return task, nil
}

func (r *taskRepository) Delete(id uint) error {
	res := r.db.Delete(&entity.Task{}, id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
