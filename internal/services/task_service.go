package services

import (
	"context"
	"example/tasksManager/internal/dto"
	"example/tasksManager/internal/repository"

	"example/tasksManager/internal/models"

	"github.com/google/uuid"
)

type TaskService struct {
	repo repository.ITaskRepository
}

type ITaskService interface {
	FindById(ctx context.Context, id uuid.UUID) (*dto.GetTaskResponse, error)
	Add(task dto.AddTaskRequest) error
	Update(task dto.UpdateTaskRequest) (affectedRows uint16, err error)
	Delete(task dto.DeleteTaskRequest) error
	GetAll(ctx context.Context) ([]dto.GetTaskResponse, error)
}

func NewTaskService(repo repository.ITaskRepository) ITaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) FindById(ctx context.Context, id uuid.UUID) (*dto.GetTaskResponse, error) {
	result, err := s.repo.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	return &dto.GetTaskResponse{
		ID:          result.ID.String(),
		Title:       result.Title,
		Description: result.Description,
		Status:      result.Status,
		DueDate:     result.DueDate,
		UpdatedAt:   result.UpdatedAt,
	}, nil
}

func (s *TaskService) Add(task dto.AddTaskRequest) error {
	taskToAdd := models.Task{
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		DueDate:     task.DueDate,
	}

	err := s.repo.Add(taskToAdd)

	return err
}

func (s *TaskService) Update(task dto.UpdateTaskRequest) (affectedRows uint16, err error) {
	taskToUpdate := models.Task{
		ID:          uuid.MustParse(task.ID),
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		DueDate:     task.DueDate,
	}

	return s.repo.Update(taskToUpdate)
}

func (s *TaskService) Delete(task dto.DeleteTaskRequest) error {
	err := s.repo.Delete(uuid.MustParse(task.ID))

	return err
}

func (s *TaskService) GetAll(ctx context.Context) ([]dto.GetTaskResponse, error) {
	result, err := s.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	response := []dto.GetTaskResponse{}

	for _, item := range result {
		response = append(response, dto.GetTaskResponse{
			ID:          item.ID.String(),
			Title:       item.Title,
			Description: item.Description,
			Status:      item.Status,
			DueDate:     item.DueDate,
			UpdatedAt:   item.UpdatedAt,
		})
	}
	return response, nil
}
