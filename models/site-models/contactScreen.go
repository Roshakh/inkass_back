package sitemodels

import (
	"time"

	"gorm.io/gorm"
)

type ContactScreen struct {
	Id          uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Description string         `json:"description"`
	Phone       string         `json:"phone"`
	PhoneInfo   string         `json:"phone_info"`
	Email       string         `json:"email"`
	EmailInfo   string         `json:"email_info"`
	Regions     Regions        `json:"regions"`
}

type Regions struct {
	Id              uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ContactScreenId uint           `json:"contact_screen_id"`
	RegionName      string         `json:"region_name"`
	Branches        Branch         `json:"branches"`
}

type Branch struct {
	Id              uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ContactScreenId string         `json:"contact_screen_id"`
	RegionsId       string         `json:"regions_id"`
	NameFilial      string         `json:"name_filial"`
	Address         string         `json:"address"`
	Phone           string         `json:"phone"`
	PhoneInfo       string         `json:"phone_info"`
	Email           string         `json:"email"`
	EmailInfo       string         `json:"email_info"`
}
