package dto

import "time"

type AddTaskRequest struct {
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description"`
	Status      bool       `json:"status"`
	DueDate     *time.Time `json:"due_date"`
}

type AddTaskResponse struct {
	ID string `json:"id"`
}

type UpdateTaskRequest struct {
	ID          string     `json:"id" binding:"required"`
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description"`
	Status      bool       `json:"status"`
	DueDate     *time.Time `json:"due_date"`
}

type UpdateTaskResponse struct {
	ID string `json:"id"`
}

type DeleteTaskRequest struct {
	ID string `json:"id" binding:"required"`
}

type GetTaskResponse struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      bool       `json:"status"`
	DueDate     *time.Time `json:"due_date"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
