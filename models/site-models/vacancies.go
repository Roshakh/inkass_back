package sitemodels

import (
	"time"

	"gorm.io/gorm"
)

type Vacancies struct {
	Id            uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CityName      string         `json:"city_name"`
	CityVacancies CityVacancies  `json:"city_vacancies"`
}

type CityVacancies struct {
	Id          uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	VacanciesId uint           `json:"vacancies_id"`
	StuffTime   string         `json:"stuff_time"`
	TimeData    string         `json:"time_data"`
	Phone       string         `json:"phone"`
}
