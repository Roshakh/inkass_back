package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id             uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name           string         `json:"name" gorm:"not null" validate:"required"`
	Phone          uint           `json:"phone"`
	Login          string         `json:"login"`
	Password       string         `json:"password"`
	Email          string         `json:"email"`
	Price          uint           `json:"price"`
	INN            uint           `json:"inn"`
	NumberContract string         `json:"number_contract"`
}
