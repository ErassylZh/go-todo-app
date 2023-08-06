package services

import (
	"my_first_go_project/models"
	"my_first_go_project/repositories"
)

type TaskService interface {
	AddTask(task models.Task)
	GetTasks() []models.Task
	DeleteTask(id int)
	UpdateTask(task models.Task)
	GetTaskById(id int)
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

func (ts *taskService) UpdateTask(task models.Task) {
	ts.taskRepository.Update(task)
}

func (ts *taskService) DeleteTask(id int) {
	ts.taskRepository.Delete(id)
}

func (ts *taskService) GetTaskById(id int) {
	ts.taskRepository.GetById(id)
}
