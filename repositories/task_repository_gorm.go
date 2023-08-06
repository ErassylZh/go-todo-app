package repositories

import (
	"gorm.io/gorm"
	"my_first_go_project/models"
)

type TaskRepository interface {
	Add(task models.Task)
	GetAll() []models.Task
}

type TaskRepositoryGORM struct {
	db *gorm.DB
}

func NewTaskRepositoryGORM(db *gorm.DB) TaskRepository {
	return &TaskRepositoryGORM{db: db}
}

func (tr *TaskRepositoryGORM) Add(task models.Task) {
	tr.db.Create(&task)
}

func (tr *TaskRepositoryGORM) GetAll() []models.Task {
	var tasks []models.Task
	tr.db.Find(&tasks)
	return tasks
}
