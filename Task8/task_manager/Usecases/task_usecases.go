package Usecases

import (
	"task_manager/Domain"
	"task_manager/Repositories"
)

type TaskUsecase interface {
	GetAllTasks() ([]Domain.Task, error)
	GetTaskByID(id string) (*Domain.Task, error)
	CreateTask(task Domain.Task) (*Domain.Task, error)
	UpdateTask(id string, updatedTask Domain.Task) (*Domain.Task, error)
	DeleteTask(id string) error
}

type TaskUsecaseImpl struct {
	taskRepo Repositories.TaskRepository
}

func NewTaskUsecase(taskRepo Repositories.TaskRepository) TaskUsecase {
	return &TaskUsecaseImpl{taskRepo: taskRepo}
}

func (u *TaskUsecaseImpl) GetAllTasks() ([]Domain.Task, error) {
	return u.taskRepo.GetAll()
}

func (u *TaskUsecaseImpl) GetTaskByID(id string) (*Domain.Task, error) {
	return u.taskRepo.GetByID(id)
}

func (u *TaskUsecaseImpl) CreateTask(task Domain.Task) (*Domain.Task, error) {
	return u.taskRepo.Create(task)
}

func (u *TaskUsecaseImpl) UpdateTask(id string, updatedTask Domain.Task) (*Domain.Task, error) {
	return u.taskRepo.Update(id, updatedTask)
}

func (u *TaskUsecaseImpl) DeleteTask(id string) error {
	return u.taskRepo.Delete(id)
}

