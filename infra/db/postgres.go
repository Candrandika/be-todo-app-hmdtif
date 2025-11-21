package db

import (
	"fmt"
	"log"

	"github.com/Candrandika/be-todo-app-hmdtif/config"
	"github.com/Candrandika/be-todo-app-hmdtif/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"context"
	"errors"
)

func NewPostgresDatabase(env *config.Env) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		env.DBHost,
		env.DBUser,
		env.DBPass,
		env.DBName,
		env.DBPort,
		"disable",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established")
	return db
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&entity.Task{})
	if err != nil {
		log.Fatalf("Failed to auto migrate database: %v", err)
	}
	log.Println("Database migration completed")
}
type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) GetByID(ctx context.Context, id uint) (entity.Task, error) {
	var t entity.Task
	if err := r.db.WithContext(ctx).First(&t, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Task{}, err
		}
		return entity.Task{}, err
	}
	return t, nil
}

func (r *taskRepository) Update(ctx context.Context, t entity.Task) (entity.Task, error) {
	if err := r.db.WithContext(ctx).Save(&t).Error; err != nil {
		return entity.Task{}, err
	}
	return t, nil
}

func (r *taskRepository) Delete(ctx context.Context, id uint) error {
	res := r.db.WithContext(ctx).Delete(&entity.Task{}, id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
