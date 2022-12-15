package repository

import (
	"fmt"
	"todo-list-api/app/entity"
	"todo-list-api/app/model"
	"todo-list-api/exception"

	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) BaseRepository[entity.Todo, model.TodoFilter] {
	return &TodoRepository{
		DB: db,
	}
}

func (r *TodoRepository) Insert(todo entity.Todo) entity.Todo {
	result := r.DB.Create(&todo)
	exception.PanicWhenError(result.Error)
	return todo
}

func (r *TodoRepository) FindAll(filter *model.TodoFilter) (todos []entity.Todo) {
	result := r.DB.Find(&todos, filter)
	exception.PanicWhenError(result.Error)
	return todos
}

func (r *TodoRepository) FindById(ID uint) (todo entity.Todo) {
	result := r.DB.Where("id = ?", ID).First(&todo)
	if result.Error != nil {
		exception.PanicNotFoundWhenError(fmt.Errorf("Todo with ID %d Not Found", ID))
	}
	return todo
}

func (r *TodoRepository) DeleteById(ID uint) {
	result := r.DB.Delete(entity.Todo{}, ID)
	if result.Error != nil {
		exception.PanicNotFoundWhenError(fmt.Errorf("Todo with ID %d Not Found", ID))
	}
}

func (r *TodoRepository) UpdateById(todo entity.Todo) entity.Todo {
	result := r.DB.Model(&todo).Updates(todo).First(&todo)
	if result.Error != nil {
		exception.PanicNotFoundWhenError(fmt.Errorf("Todo with ID %d Not Found", todo.ID))
	}
	return todo
}
