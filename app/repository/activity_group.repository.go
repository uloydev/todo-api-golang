package repository

import (
	"todo-list-api/app/entity"
	"todo-list-api/exception"

	"gorm.io/gorm"
)

type ActivityGroupRepository struct {
	DB *gorm.DB
}

func NewActivityGroupRepository(db *gorm.DB) BaseRepository[entity.ActivityGroup] {
	return &ActivityGroupRepository{
		DB: db,
	}
}

func (r *ActivityGroupRepository) Insert(activityGroup entity.ActivityGroup) entity.ActivityGroup {
	result := r.DB.Create(&activityGroup)
	exception.PanicWhenError(result.Error)
	return activityGroup
}

func (r *ActivityGroupRepository) FindAll() (activityGroups []entity.ActivityGroup) {
	result := r.DB.Find(&activityGroups)
	exception.PanicWhenError(result.Error)
	return activityGroups
}

func (r *ActivityGroupRepository) FindById(ID uint) (activityGroup entity.ActivityGroup) {
	result := r.DB.Where("id = ?", ID).First(&activityGroup)
	exception.PanicValidationWhenError(result.Error)
	return activityGroup
}

func (r *ActivityGroupRepository) DeleteById(ID uint) {
	result := r.DB.Delete(entity.ActivityGroup{}, ID)
	exception.PanicValidationWhenError(result.Error)
	return
}

func (r *ActivityGroupRepository) UpdateById(activityGroup entity.ActivityGroup) entity.ActivityGroup {
	result := r.DB.Model(&activityGroup).Updates(activityGroup).First(&activityGroup)
	exception.PanicValidationWhenError(result.Error)
	return activityGroup
}
