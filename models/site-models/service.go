package sitemodels

import (
	"time"

	"gorm.io/gorm"
)

type Service struct {
	Id             uint              `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt      time.Time         `json:"created_at"`
	UpdatedAt      time.Time         `json:"updated_at"`
	DeletedAt      gorm.DeletedAt    `gorm:"index" json:"deleted_at"`
	Content        string            `json:"content"`
	ServiceCatalog []ServiceCatalog `json:"service_catalog"`
	SingleService  SingleService     `json:"single_services"`
}
type ServiceCatalog struct {
	Id          uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ServiceId   uint           `json:"service_id"`
	ImageUrl    string         `json:"image_url"`
	Content     string         `json:"content"`
	Description string         `json:"description"`
}

type SingleService struct {
	Id                   uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
	DeletedAt            gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ServiceId            uint           `json:"service_id"`
	Content              string         `json:"content"`
	PriceService         string         `json:"price_service"`
	ServiceMethodsNumber string         `json:"service_methods_number"`
	ServiceMethods       string         `json:"service_methods"`
}
