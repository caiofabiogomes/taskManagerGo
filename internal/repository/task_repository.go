package repository

import (
	"context"
	"errors"
	"example/tasksManager/internal/models"
	"time"

	"database/sql"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type TaskRepository struct {
	taskContext TaskContext
}

// Add implements [ITaskRepository].
func (r *TaskRepository) Add(task models.Task) error {
	err := r.taskContext.Connect()
	if err != nil {
		return errors.New("failed to connect to database")
	}

	defer r.taskContext.Close()

	if _, err := r.taskContext.Dbx.Exec(
		`INSERT INTO tasks(title,description,status,due_date,created_at) 
		VALUES($1,$2,$3,$4,$5)`,
		task.Title, task.Description, task.Status, task.DueDate, time.Now()); err != nil {
		return err
	}

	return nil
}

// Delete implements [ITaskRepository].
func (r *TaskRepository) Delete(id uuid.UUID) error {

	err := r.taskContext.Connect()
	if err != nil {
		return errors.New("failed to connect to database")
	}

	defer r.taskContext.Close()

	result, err := r.taskContext.Dbx.Exec(
		`UPDATE tasks SET deleted_at = $1 WHERE ID = $2 AND deleted_at IS NULL`, time.Now(), id)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("task not found")
	}

	return nil
}

// FindById implements [ITaskRepository].
func (r *TaskRepository) FindById(ctx context.Context, id uuid.UUID) (*models.Task, error) {
	err := r.taskContext.Connect()
	if err != nil {
		return nil, err
	}
	defer r.taskContext.Close()

	var task models.Task
	err = r.taskContext.Dbx.GetContext(
		ctx,
		&task,
		`SELECT *
		 FROM tasks 
		 WHERE id = $1 AND deleted_at IS NULL`,
		id,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &task, nil
}

// Update implements [ITaskRepository].
func (r *TaskRepository) Update(task models.Task) (affectedRows uint16, err error) {
	err = r.taskContext.Connect()
	if err != nil {
		return 0, errors.New("failed to connect to database")
	}
	defer r.taskContext.Close()

	result, err := r.taskContext.Dbx.Exec(
		`UPDATE tasks SET  
						title = $1,
						description = $2,
						status = $3,
						due_date = $4,
						updated_at = $5
		 WHERE ID = $6 AND deleted_at IS NULL`,
		task.Title, task.Description, task.Status, task.DueDate, time.Now(), task.ID)

	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rows == 0 {
		return 0, nil
	}

	return uint16(rows), nil
}

type ITaskRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (*models.Task, error)
	Add(task models.Task) error
	Update(task models.Task) (affectedRows uint16, err error)
	Delete(id uuid.UUID) error
	GetAll(ctx context.Context) ([]models.Task, error)
}

func NewTaskRepository(taskContext TaskContext) ITaskRepository {
	return &TaskRepository{
		taskContext: taskContext,
	}
}

func (r *TaskRepository) GetAll(ctx context.Context) ([]models.Task, error) {
	err := r.taskContext.Connect()
	if err != nil {
		return nil, err
	}
	defer r.taskContext.Close()

	var tasks []models.Task
	if err := r.taskContext.Dbx.SelectContext(
		ctx,
		&tasks,
		`SELECT
             *
        FROM tasks WHERE deleted_at IS NULL`); err != nil {
		return nil, err
	}

	if tasks == nil {
		return []models.Task{}, nil
	}

	return tasks, nil
}
