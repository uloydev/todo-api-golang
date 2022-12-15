package entity

type Activity struct {
	BaseEntity
	Email string `gorm:"type:varchar(100)"`
	Title string `gorm:"type:varchar(255)"`
}

func (Activity) TableName() string {
	return "activities"
}
