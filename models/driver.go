package models

import (
	"time"

	"gorm.io/gorm"
)

type Driver struct {
	Id        uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name      string         `json:"name" gorm:"not null" validate:"required"`
	Phone     uint           `json:"phone"`
	PhotoUrl  string         `json:"photo_url"`
	Login     string         `json:"login"`
	Password  string         `json:"password"`
	Bank      []*Bank        `json:"bank"`
	Route     []*Route       `json:"route"`
	Company   []*Company     `json:"company"`
}

type Bank struct {
	Id        uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name      string         `json:"name" gorm:"not null" validate:"required"`
	Phone     uint           `json:"phone"`
	Locations []Location     `json:"locations" gorm:"locations"`
}

type Route struct {
	Id        uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Locations []Location     `json:"locations" gorm:"locations"`
}

type Location struct {
	Id        uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Latitude  float32        `json:"latitude"`
	Longitude float32        `json:"longitude"`
}

type Company struct {
	Id        uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
