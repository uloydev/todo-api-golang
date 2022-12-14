package service

import (
	"todo-list-api/app/entity"
	"todo-list-api/app/model"
	"todo-list-api/app/repository"
	"todo-list-api/app/validation"
)

type ActivityGroupService struct {
	Repo *repository.ActivityGroupRepository
}

func NewActivityGroupService(repo *repository.ActivityGroupRepository) BaseService[model.ActivityGroupRequest, model.ActivityGroupResponse] {
	return &ActivityGroupService{
		Repo: repo,
	}
}

func (s *ActivityGroupService) Create(request model.ActivityGroupRequest) (response model.ActivityGroupResponse) {
	validation.ValidateActivityGroup(request)

	activityGroup := entity.ActivityGroup{
		Email: request.Email,
		Title: request.Title,
	}
	activityGroup = s.Repo.Insert(activityGroup)

	response = model.ActivityGroupResponse{
		BasicData: model.BasicData{
			ID:        activityGroup.ID,
			CreatedAt: activityGroup.CreatedAt,
			UpdatedAt: activityGroup.UpdatedAt,
			DeletedAt: activityGroup.DeletedAt,
		},
		Email: activityGroup.Email,
		Title: activityGroup.Title,
	}

	return response
}

func (s *ActivityGroupService) List() (responses []model.ActivityGroupResponse) {
	activityGroups := s.Repo.FindAll()

	for _, activityGroup := range activityGroups {
		responses = append(responses, model.ActivityGroupResponse{
			BasicData: model.BasicData{
				ID:        activityGroup.ID,
				CreatedAt: activityGroup.CreatedAt,
				UpdatedAt: activityGroup.UpdatedAt,
				DeletedAt: activityGroup.DeletedAt,
			},
			Title: activityGroup.Title,
			Email: activityGroup.Email,
		})
	}

	return responses
}

func (s *ActivityGroupService) FindById(ID uint) (response model.ActivityGroupResponse) {
	activityGroup := s.Repo.FindById(ID)

	response = model.ActivityGroupResponse{
		BasicData: model.BasicData{
			ID:        activityGroup.ID,
			CreatedAt: activityGroup.CreatedAt,
			UpdatedAt: activityGroup.UpdatedAt,
			DeletedAt: activityGroup.DeletedAt,
		},
		Title: activityGroup.Title,
		Email: activityGroup.Email,
	}

	return response
}

func (s *ActivityGroupService) DeleteById(ID uint) {
	_ = s.Repo.FindById(ID)
	s.Repo.DeleteById(ID)
}

func (s *ActivityGroupService) UpdateById(ID uint, request model.ActivityGroupRequest) (response model.ActivityGroupResponse) {
	validation.ValidateActivityGroup(request)

	activityGroup := entity.ActivityGroup{
		Email: request.Email,
		Title: request.Title,
		BaseEntity: entity.BaseEntity{
			ID: ID,
		},
	}

	activityGroup = s.Repo.UpdateById(activityGroup)

	response = model.ActivityGroupResponse{
		BasicData: model.BasicData{
			ID:        activityGroup.ID,
			CreatedAt: activityGroup.CreatedAt,
			UpdatedAt: activityGroup.UpdatedAt,
			DeletedAt: activityGroup.DeletedAt,
		},
		Title: activityGroup.Title,
		Email: activityGroup.Email,
	}

	return response
}
