package sitemodels

import (
	"time"

	"inkass/inkassback/models"

	"gorm.io/gorm"
)

type Contacts struct {
	Id        uint               `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	DeletedAt gorm.DeletedAt     `gorm:"index" json:"deleted_at"`
	Address   string             `json:"address"`
	Phone     string             `json:"phone"`
	Email     string             `json:"email"`
	Fax       string             `json:"strig"`
	Links     []*Link            `json:"links"`
	Branches  []*models.Location `json:"location"`
}

type Link struct {
	Id        uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Value     string         `json:"value"`
	IconUrl   string         `json:"icon_url"`
}
