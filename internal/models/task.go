package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Title       string    `db:"title" json:"title" binding:"required"`
	Description string    `db:"description" json:"description"`
	Status      bool      `db:"status" json:"status"` // Ótima escolha usar int (iota) aqui!
	//UserID    uint       `db:"user_id" json:"user_id"`
	DueDate   *time.Time `db:"due_date" json:"due_date"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"-"`
}
