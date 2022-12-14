package entity

import "time"

type BaseEntity struct {
	ID        uint   `gorm:"primaryKey"`
	CreatedAt time.Time 
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}
