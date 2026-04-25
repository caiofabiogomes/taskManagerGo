package main

import "time"

type Task struct {
	ID          uint   `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Status      int    `json:"status"` // Usando iota para os estados
	//UserID      uint      `json:"user_id"`
	DueDate   time.Time `json:"due_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}
