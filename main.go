package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"my_first_go_project/config"
	"my_first_go_project/handlers"
	"my_first_go_project/models"
	"my_first_go_project/repositories"
	"my_first_go_project/services"
)

func main() {
	r := gin.Default()
	dbConfig := config.NewDatabaseConfig()
	dsn := "user=" + dbConfig.User + " password=" + dbConfig.Password + " dbname=" + dbConfig.DBName + " sslmode=" + dbConfig.SSLMode
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect db")
	}
	db.AutoMigrate(&models.Task{})

	taskRepository := repositories.NewTaskRepositoryGORM(db)
	taskService := services.NewTaskService(taskRepository)
	taskHandler := handlers.NewTaskHandler(taskService)
	handlers.SetupRoutes(r, taskHandler)
	r.Run(config.Port)
}
