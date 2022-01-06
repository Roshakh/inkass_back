package models

import (
	"time"

	"gorm.io/gorm"
)

type Reviews struct {
	Id          uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ImageUrl    string         `json:"image_url"`
	UserName    string         `json:"user_name"`
	Content     string         `json:"content"`
	Description string         `json:"description"`
	Rates       uint           `json:"rates"`
}
