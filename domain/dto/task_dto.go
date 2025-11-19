package dto

import "time"

type TaskResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsDone      bool      `json:"is_done"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}
