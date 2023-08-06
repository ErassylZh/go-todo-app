package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"my_first_go_project/auth/handlers"
	"my_first_go_project/auth/models"
	"my_first_go_project/auth/repositories"
	"my_first_go_project/auth/services"
	"my_first_go_project/config"
)

func main() {
	r := gin.Default()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	dsn := "user=" + cfg.Auth.DBUser + " password=" + cfg.Auth.DBPassword +
		" dbname=" + cfg.Auth.DBName + " sslmode=" + cfg.Auth.SSLMode
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&models.User{})

	userRepository := repositories.NewUserRepositoryGORM(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	handlers.SetupRoutes(r, userHandler)
	r.Run(":" + cfg.Auth.Port)
}
