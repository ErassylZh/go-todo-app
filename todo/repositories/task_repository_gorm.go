package repositories

import (
	"gorm.io/gorm"
	"my_first_go_project/todo/models"
)

type TaskRepository interface {
	Add(task models.Task) (*models.Task, error)
	GetAll() []models.Task
	Update(task models.Task)
	Delete(id int)
	GetById(id int) (models.Task, error)
}

type TaskRepositoryGORM struct {
	db *gorm.DB
}

func NewTaskRepositoryGORM(db *gorm.DB) *TaskRepositoryGORM {
	return &TaskRepositoryGORM{db: db}
}

func (tr *TaskRepositoryGORM) Add(task models.Task) (*models.Task, error) {
	tr.db.Create(&task)
	return &task, nil
}

func (tr *TaskRepositoryGORM) GetAll() []models.Task {
	var tasks []models.Task
	tr.db.Find(&tasks)
	return tasks
}

func (tr *TaskRepositoryGORM) Update(task models.Task) {
	tr.db.Save(&task)
}

func (tr *TaskRepositoryGORM) Delete(id int) {
	tr.db.Delete(&models.Task{}, id)
}

func (tr *TaskRepositoryGORM) GetById(id int) (models.Task, error) {
	var task models.Task
	result := tr.db.First(&task, id)
	if result.Error != nil {
		return models.Task{}, result.Error
	}
	return task, nil
}
