package repository

import (
	"github.com/Candrandika/be-todo-app-hmdtif/domain/entity"
	"gorm.io/gorm"
)

type TaskRepository interface {
	GetAll() ([]entity.Task, error)
	CreateNew(task *entity.Task) error
	GetByID(id uint) (*entity.Task, error)
	Update(task *entity.Task) error
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
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) CreateNew(task *entity.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) GetByID(id uint) (*entity.Task, error) {
	var task entity.Task
	if err := r.db.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) Update(task *entity.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) Delete(id uint) error {
	result := r.db.Delete(&entity.Task{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
