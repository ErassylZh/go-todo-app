package services

import (
	"errors"
	"fmt"
	"my_first_go_project/todo/models"
	"my_first_go_project/todo/repositories"
	"strconv"
)

type TaskService interface {
	AddTask(task models.Task, tokenString string) (*models.Task, error)
	GetTasks() []models.Task
	DeleteTask(id int)
	UpdateTask(taskID string, task models.Task, tokenString string) (*models.Task, error)
	GetTaskById(id int)
}

type taskService struct {
	taskRepository repositories.TaskRepository
	authService    AuthService
}

func NewTaskService(taskRepository *repositories.TaskRepositoryGORM, authService AuthService) TaskService {
	return &taskService{
		taskRepository: taskRepository,
		authService:    authService,
	}
}

func (ts *taskService) AddTask(task models.Task, tokenString string) (*models.Task, error) {
	userID, err := ts.authService.GetUserFromToken(tokenString)
	fmt.Println(userID)
	if err != nil {
		return nil, err
	}

	task.CreatorID = int(userID)
	return ts.taskRepository.Add(task)
}

func (ts *taskService) GetTasks() []models.Task {
	return ts.taskRepository.GetAll()
}

func (ts *taskService) UpdateTask(id string, task models.Task, token string) (*models.Task, error) {
	userID, err := ts.authService.GetUserFromToken(token)
	if err != nil {
		return nil, err
	}
	taskId, errConv := strconv.Atoi(id)
	if errConv != nil {
		panic(errConv)
		return nil, errConv
	}
	existingTask, err := ts.taskRepository.GetById(taskId)
	if err != nil {
		return nil, err
	}

	if existingTask.CreatorID != int(userID) {
		return nil, errors.New("You don't have permission to update this task")
	}
	existingTask.Title = task.Title
	ts.taskRepository.Update(existingTask)
	return &existingTask, nil
}

func (ts *taskService) DeleteTask(id int) {
	ts.taskRepository.Delete(id)
}

func (ts *taskService) GetTaskById(id int) {
	ts.taskRepository.GetById(id)
}
