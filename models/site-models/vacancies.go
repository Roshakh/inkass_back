package sitemodels

import (
	"time"

	"gorm.io/gorm"
)

type Vacancies struct {
	Id        uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Limit         string          `json:"limit"`
	CityVacancies []CityVacancies `json:"city_vacancies"`
}

type CityVacancies struct {
	Id              uint              `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	DeletedAt       gorm.DeletedAt    `gorm:"index" json:"deleted_at"`
	VacanciesId     uint              `json:"vacancies_id"`
	CityName        string            `json:"city_name"`
	Content         string            `json:"content"`
	Location        string            `json:"location"`
	StaffTime       string            `json:"staff_time"`
	Time            string            `json:"time"`
	Phone           string            `json:"phone"`
	SimpleVacancies []SimpleVacancies `json:"simple_vacancies"`
}

type SimpleVacancies struct {
	Id                        uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt                 time.Time      `json:"created_at"`
	UpdatedAt                 time.Time      `json:"updated_at"`
	DeletedAt                 gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CityVacanciesId           uint           `json:"city_vacancies_id"`
	Content                   string         `json:"content"`
	Location                  string         `json:"location"`
	StaffTime                 string         `json:"staff_time"`
	Time                      string         `json:"time"`
	Requirements              string         `json:"requirements"`
	ImportantRequirementTitle string         `json:"important_requirements_title"`
	ImportantRequirements     string         `json:"important_requirements"`
}
