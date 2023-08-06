package handlers

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, th *TaskHandler) {
	api := r.Group("/api")
	{
		api.POST("/task", th.AddTask)
		api.GET("/tasks", th.GetTasks)
	}
}
