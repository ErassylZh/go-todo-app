package handlers

import (
	"github.com/gin-gonic/gin"
	"my_first_go_project/todo/models"
	"my_first_go_project/todo/services"
	"net/http"
	"strconv"
	"strings"
)

type TaskHandler struct {
	taskService services.TaskService
}

func NewTaskHandler(taskService services.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (th *TaskHandler) AddTask(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token missing"})
		return
	}
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTask, err := th.taskService.AddTask(task, tokenString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task  " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTask)
}

func (th *TaskHandler) GetTasks(c *gin.Context) {
	tasks := th.taskService.GetTasks()
	c.JSON(http.StatusOK, tasks)
}

func (th *TaskHandler) UpdateTask(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token missing"})
		return
	}

	taskID := c.Param("id")
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTask, err := th.taskService.UpdateTask(taskID, task, tokenString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
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
