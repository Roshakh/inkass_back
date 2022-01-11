package sitemodels

import (
	"time"

	"gorm.io/gorm"
)

type TarifTable struct {
	Id            uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Monthly       string         `json:"monthly"`
	MonthlyPrc    string         `json:"monthly_prc"`
	ForCustom     string         `json:"for_custom"`
	ForCustomOnce string         `json:"for_custome_once"`
	LimitMonthly  string         `json:"limit_monthly"`
}
