package repository

import (
	"fmt"
	"todo-list-api/app/entity"
	"todo-list-api/app/model"
	"todo-list-api/exception"

	"gorm.io/gorm"
)

type ActivityGroupRepository struct {
	DB *gorm.DB
}

func NewActivityGroupRepository(db *gorm.DB) BaseRepository[entity.Activity, model.ActivityGroupFilter] {
	return &ActivityGroupRepository{
		DB: db,
	}
}

func (r *ActivityGroupRepository) Insert(activityGroup entity.Activity) entity.Activity {
	result := r.DB.Create(&activityGroup)
	exception.PanicWhenError(result.Error)
	return activityGroup
}

func (r *ActivityGroupRepository) FindAll(filter *model.ActivityGroupFilter) (activityGroups []entity.Activity) {
	result := r.DB.Find(&activityGroups, filter)
	exception.PanicWhenError(result.Error)
	return activityGroups
}

func (r *ActivityGroupRepository) FindById(ID uint) (activityGroup entity.Activity) {
	result := r.DB.Where("id = ?", ID).First(&activityGroup)
	if result.Error != nil {
		exception.PanicNotFoundWhenError(fmt.Errorf("Activity with ID %d Not Found", ID))
	}
	return activityGroup
}

func (r *ActivityGroupRepository) DeleteById(ID uint) {
	result := r.DB.Delete(entity.Activity{}, ID)
	if result.Error != nil {
		exception.PanicNotFoundWhenError(fmt.Errorf("Activity with ID %d Not Found", ID))
	}
}

func (r *ActivityGroupRepository) UpdateById(activityGroup entity.Activity) entity.Activity {
	result := r.DB.Model(&activityGroup).Updates(activityGroup).First(&activityGroup)
	if result.Error != nil {
		exception.PanicNotFoundWhenError(fmt.Errorf("Activity with ID %d Not Found", activityGroup.ID))
	}
	return activityGroup
}
