package sitemodels

import (
	"time"

	"gorm.io/gorm"
)

type News struct {
	Id           uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Description  string         `json:"description"`
	SingleNew    []SingleNew    `json:"single_new"`
}

type SingleNew struct {
	Id          uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	NewsId      uint           `json:"news_id"`
	ImageUrl    string         `json:"image_url"`
	Title       string         `json:"title"`
	Time        string         `json:"time"`
	Description string         `json:"description"`
}
