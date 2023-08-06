package handlers

import (
	"github.com/gin-gonic/gin"
	"my_first_go_project/models"
	"my_first_go_project/services"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	taskService services.TaskService
}

func NewTaskHandler(taskService services.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (th *TaskHandler) AddTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	th.taskService.AddTask(task)
	c.Status(http.StatusCreated)
}

func (th *TaskHandler) GetTasks(c *gin.Context) {
	tasks := th.taskService.GetTasks()
	c.JSON(http.StatusOK, tasks)
}

func (th *TaskHandler) UpdateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	th.taskService.UpdateTask(task)
	c.Status(http.StatusNoContent)
}

func (th *TaskHandler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	taskId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	th.taskService.DeleteTask(taskId)
	c.Status(http.StatusOK)
}

func (th *TaskHandler) GetTaskById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}
	taskId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	th.taskService.GetTaskById(taskId)
	c.Status(http.StatusOK)
}
