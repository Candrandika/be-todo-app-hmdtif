package entity

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255);not null"`
	Description string
	IsDone      bool `gorm:"default:false"`
}
