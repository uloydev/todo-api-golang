package entity

type ActivityGroup struct {
	BaseEntity
	Email string `gorm:"type:varchar(100)"`
	Title string `gorm:"type:varchar(255)"`
}
