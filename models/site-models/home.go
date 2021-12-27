package sitemodels

import (
	"time"

	"gorm.io/gorm"
)

type HomeScreen struct {
	Id            uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	AboutUs       AboutUs
	Service       Service
	Contacts      Contacts
	CurrencyRates []CurrencyRates
}

type CurrencyRates struct {
	ID       int    `json:"id"`
	Code     string `json:"Code"`
	Ccy      string `json:"Ccy"`
	CcyNmRU  string `json:"CcyNm_RU"`
	CcyNmUZ  string `json:"CcyNm_UZ"`
	CcyNmUZC string `json:"CcyNm_UZC"`
	CcyNmEN  string `json:"CcyNm_EN"`
	Nominal  string `json:"Nominal"`
	Rate     string `json:"Rate"`
	Diff     string `json:"Diff"`
	Date     string `json:"Date"`
}

type AboutUs struct {
	Id        uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ImageUrl  string         `json:"image_url"`
	Content   string         `json:"content"`
}

type ServiceCatogory struct {
	Id        uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Service struct {
	Id                uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Services          []*SingleService
	ServiceCatogories []*ServiceCatogory
}

type SingleService struct {
	Id          uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CategoryId  uint
	Title       string
	ImageUrl    string `json:"image_url"`
	Description string
}

type SingleNew struct {
	Id          uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CategoryId  uint           `json:"category_id"`
	Title       string         `json:"title"`
	ImageUrl    string         `json:"image_url"`
	Description string         `json:"description"`
}

type News struct {
	Id          uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CategoryId  uint           `json:"category_id"`
	Title       string         `json:"title"`
	ImageUrl    string         `json:"image_url"`
	Description string         `json:"description"`
	Content     string         `json:"content"`
	NewsList    SingleNew
}
