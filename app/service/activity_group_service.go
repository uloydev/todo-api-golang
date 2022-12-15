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

func NewActivityGroupService(repo *repository.ActivityGroupRepository) BaseService[model.ActivityGroupRequest, model.ActivityGroupResponse, model.ActivityGroupFilter] {
	return &ActivityGroupService{
		Repo: repo,
	}
}

func (s *ActivityGroupService) Create(req model.ActivityGroupRequest) (resp model.ActivityGroupResponse) {
	validation.ValidateActivityGroup(req)

	activityGroup := entity.Activity{
		Title: req.Title,
		Email: req.Email,
	}
	activityGroup = s.Repo.Insert(activityGroup)

	resp = model.ActivityGroupResponse{
		BasicData: model.BasicData{
			ID:        activityGroup.ID,
			CreatedAt: activityGroup.CreatedAt,
			UpdatedAt: activityGroup.UpdatedAt,
			DeletedAt: activityGroup.DeletedAt,
		},
		Email: activityGroup.Email,
		Title: activityGroup.Title,
	}

	return resp
}

func (s *ActivityGroupService) List(filter *model.ActivityGroupFilter) (resps []model.ActivityGroupResponse) {
	activityGroups := s.Repo.FindAll(filter)

	for _, activityGroup := range activityGroups {
		resps = append(resps, model.ActivityGroupResponse{
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

	return resps
}

func (s *ActivityGroupService) FindById(ID uint) (resp model.ActivityGroupResponse) {
	activityGroup := s.Repo.FindById(ID)

	resp = model.ActivityGroupResponse{
		BasicData: model.BasicData{
			ID:        activityGroup.ID,
			CreatedAt: activityGroup.CreatedAt,
			UpdatedAt: activityGroup.UpdatedAt,
			DeletedAt: activityGroup.DeletedAt,
		},
		Title: activityGroup.Title,
		Email: activityGroup.Email,
	}

	return resp
}

func (s *ActivityGroupService) DeleteById(ID uint) {
	_ = s.Repo.FindById(ID)
	s.Repo.DeleteById(ID)
}

func (s *ActivityGroupService) UpdateById(ID uint, req model.ActivityGroupRequest) (resp model.ActivityGroupResponse) {
	validation.ValidateActivityGroup(req)

	activityGroup := entity.Activity{
		Email: req.Email,
		Title: req.Title,
		BaseEntity: entity.BaseEntity{
			ID: ID,
		},
	}

	activityGroup = s.Repo.UpdateById(activityGroup)

	resp = model.ActivityGroupResponse{
		BasicData: model.BasicData{
			ID:        activityGroup.ID,
			CreatedAt: activityGroup.CreatedAt,
			UpdatedAt: activityGroup.UpdatedAt,
			DeletedAt: activityGroup.DeletedAt,
		},
		Title: activityGroup.Title,
		Email: activityGroup.Email,
	}

	return resp
}
