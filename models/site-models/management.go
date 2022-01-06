package sitemodels

import (
	"time"

	"gorm.io/gorm"
)

type Management struct {
	Id                  uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ImageUrl            string         `json:"image_url"`
	FullName            string         `json:"full_name"`
	Position            string         `json:"position"`
	Phone               string         `json:"phone"`
	Fax                 string         `json:"fax"`
	Email               string         `json:"email"`
	Biography           Biography      `json:"biography"`
	JobResponsibilities string         `json:"job_responsibilities"`
	WorkTime            WorkTime       `json:"work_time"`
}

type Biography struct {
	Id             uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ManagementId   uint           `json:"management_id"`
	WorkPlace      string         `json:"work_place"`
	BirthPlace     string         `json:"birth_place"`
	Education      string         `json:"education"`
	WorkExperience string         `json:"work_experience"`
	Nationality    string         `json:"nationality"`
	BirthYear      string         `json:"birth_year"`
	Speciality     string         `json:"speciality"`
	Phone          string         `json:"phone"`
}

type WorkTime struct {
	Id           uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ManagementId uint           `json:"management_id"`
	Monday       string         `json:"monday"`
	Tuesday      string         `json:"tuesday"`
	Wednesday    string         `json:"wednesday"`
	Thursday     string         `json:"thursday"`
	Friday       string         `json:"friday"`
	Saturday     string         `json:"saturday"`
	Sunday       string         `json:"sunday"`
}
