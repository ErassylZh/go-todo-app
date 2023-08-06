package services

import (
	"my_first_go_project/models"
	"my_first_go_project/repositories"
)

type TaskService interface {
	AddTask(task models.Task)
	GetTasks() []models.Task
}

type taskService struct {
	taskRepository repositories.TaskRepository
}

func NewTaskService(tr repositories.TaskRepository) TaskService {
	return &taskService{taskRepository: tr}
}

func (ts *taskService) AddTask(task models.Task) {
	ts.taskRepository.Add(task)
}

func (ts *taskService) GetTasks() []models.Task {
	return ts.taskRepository.GetAll()
}
