package Usecases

import (
	"context"
	"errors"
	"task_manager/Domain"
	"task_manager/Repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"

)

type TaskUsecase interface {
	GetAllTasks(ctx context.Context) ([]Domain.Task, error)
	GetTaskByID(ctx context.Context, id string) (*Domain.Task, error)
	CreateTask(ctx context.Context, task *Domain.Task) (*Domain.Task, error)
	UpdateTask(ctx context.Context, id string, task *Domain.Task) (*Domain.Task, error)
	DeleteTask(ctx context.Context, id string) error
}

type TaskUsecaseImpl struct {
	taskRepo Repositories.TaskRepository
}

func NewTaskUsecase(taskRepo Repositories.TaskRepository) TaskUsecase {
	return &TaskUsecaseImpl{taskRepo: taskRepo}
}

func (uc *TaskUsecaseImpl) GetAllTasks(ctx context.Context) ([]Domain.Task, error) {
	tasks, err := uc.taskRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (uc *TaskUsecaseImpl) GetTaskByID(ctx context.Context, id string) (*Domain.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task ID format")
	}

	task, err := uc.taskRepo.GetByID(ctx, objID)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (uc *TaskUsecaseImpl) CreateTask(ctx context.Context, task *Domain.Task) (*Domain.Task, error) {
	createdTask, err := uc.taskRepo.Create(ctx, task)
	if err != nil {
		return nil, err
	}
	return createdTask, nil
}

func (uc *TaskUsecaseImpl) UpdateTask(ctx context.Context, id string, task *Domain.Task) (*Domain.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task ID format")
	}

	updatedTask, err := uc.taskRepo.Update(ctx, objID, task)
	if err != nil {
		return nil, err
	}
	return updatedTask, nil
}

func (uc *TaskUsecaseImpl) DeleteTask(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return  errors.New("invalid task ID format")
	}

	err = uc.taskRepo.Delete(ctx, objID)
	if err != nil {
		return err
	}
	return nil
}