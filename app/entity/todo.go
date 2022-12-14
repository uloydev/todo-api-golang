package entity

type Todo struct {
	BaseEntity
	ActivityGroupId uint
	Title           string `gorm:"type:varchar(255)"`
	IsActive        bool
	Priority        string `gorm:"type:varchar(100)"`
}
