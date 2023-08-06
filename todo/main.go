package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"my_first_go_project/config"
	"my_first_go_project/todo/handlers"
	"my_first_go_project/todo/models"
	"my_first_go_project/todo/repositories"
	"my_first_go_project/todo/services"
)

func main() {
	r := gin.Default()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	dsn := "user=" + cfg.Todo.DBUser + " password=" + cfg.Todo.DBPassword +
		" dbname=" + cfg.Todo.DBName + " sslmode=" + cfg.Todo.SSLMode
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&models.Task{})

	taskRepository := repositories.NewTaskRepositoryGORM(db)
	authService := services.NewAuthService(cfg.JWT.SecretKey)
	taskService := services.NewTaskService(taskRepository, authService)
	taskHandler := handlers.NewTaskHandler(taskService)

	handlers.SetupRoutes(r, taskHandler)
	r.Run(":" + cfg.Todo.Port)
}
