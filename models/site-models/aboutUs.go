package sitemodels

import (
	"time"

	"gorm.io/gorm"
)

type AboutUs struct {
	Id          uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ImageUrl    string         `json:"image_url"`
	Content     string         `json:"content"`
	Description string         `json:"description"`
	WhyUs       string         `json:"why_us"`
	Name1       string         `json:"name1"`
	Name1Uint   uint           `json:"name1_uint"`
	Name2       string         `json:"name2"`
	Name2Uint   uint           `json:"name2_uint"`
	Name3       string         `json:"name3"`
	Name3Uint   uint           `json:"name3_uint"`
	Name4       string         `json:"name4"`
	Name4Uint   uint           `json:"name4_uint"`
	HowWorking  string         `json:"how_working"`
	Reviews     []Review       `json:"review"`
}

type Review struct {
	Id          uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	AboutUsId   uint           `json:"about_us_id"`
	Content     string         `json:"content"`
	Description string         `json:"description"`
	ImageUrl    string         `json:"image_url"`
	UserName    string         `json:"user_name"`
	Rates       uint           `json:"rates"`
}
