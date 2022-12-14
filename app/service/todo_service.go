package service

import (
	"todo-list-api/app/entity"
	"todo-list-api/app/model"
	"todo-list-api/app/repository"
	"todo-list-api/app/validation"
)

type TodoService struct {
	Repo *repository.TodoRepository
}

func NewTodoService(repo *repository.TodoRepository) BaseService[model.TodoRequest, model.TodoResponse, model.TodoFilter] {
	return &TodoService{
		Repo: repo,
	}
}

func (s *TodoService) Create(req model.TodoRequest) (resp model.TodoResponse) {
	validation.ValidateTodo(req)

	todo := entity.Todo{
		ActivityGroupId: req.ActivityGroupId,
		IsActive:        req.IsActive,
		Priority:        req.Priority,
		Title:           req.Title,
	}
	todo = s.Repo.Insert(todo)

	resp = model.TodoResponse{
		BasicData: model.BasicData{
			ID:        todo.ID,
			CreatedAt: todo.CreatedAt,
			UpdatedAt: todo.UpdatedAt,
			DeletedAt: todo.DeletedAt,
		},
		ActivityGroupId: todo.ActivityGroupId,
		IsActive:        todo.IsActive,
		Priority:        todo.Priority,
		Title:           todo.Title,
	}

	return resp
}

func (s *TodoService) List(filter *model.TodoFilter) (resps []model.TodoResponse) {
	todos := s.Repo.FindAll(filter)

	for _, todo := range todos {
		resps = append(resps, model.TodoResponse{
			BasicData: model.BasicData{
				ID:        todo.ID,
				CreatedAt: todo.CreatedAt,
				UpdatedAt: todo.UpdatedAt,
				DeletedAt: todo.DeletedAt,
			},
			ActivityGroupId: todo.ActivityGroupId,
			IsActive:        todo.IsActive,
			Priority:        todo.Priority,
			Title:           todo.Title,
		})
	}

	return resps
}

func (s *TodoService) FindById(ID uint) (resp model.TodoResponse) {
	todo := s.Repo.FindById(ID)

	resp = model.TodoResponse{
		BasicData: model.BasicData{
			ID:        todo.ID,
			CreatedAt: todo.CreatedAt,
			UpdatedAt: todo.UpdatedAt,
			DeletedAt: todo.DeletedAt,
		},
		ActivityGroupId: todo.ActivityGroupId,
		IsActive:        todo.IsActive,
		Priority:        todo.Priority,
		Title:           todo.Title,
	}

	return resp
}

func (s *TodoService) DeleteById(ID uint) {
	_ = s.Repo.FindById(ID)
	s.Repo.DeleteById(ID)
}

func (s *TodoService) UpdateById(ID uint, req model.TodoRequest) (resp model.TodoResponse) {
	validation.ValidateTodo(req)

	todo := entity.Todo{
		ActivityGroupId: req.ActivityGroupId,
		IsActive:        req.IsActive,
		Priority:        req.Priority,
		Title:           req.Title,
		BaseEntity: entity.BaseEntity{
			ID: ID,
		},
	}

	todo = s.Repo.UpdateById(todo)

	resp = model.TodoResponse{
		BasicData: model.BasicData{
			ID:        todo.ID,
			CreatedAt: todo.CreatedAt,
			UpdatedAt: todo.UpdatedAt,
			DeletedAt: todo.DeletedAt,
		},
		ActivityGroupId: todo.ActivityGroupId,
		IsActive:        todo.IsActive,
		Priority:        todo.Priority,
		Title:           todo.Title,
	}

	return resp
}
