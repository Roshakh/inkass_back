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
	Monthly       uint           `json:"monthly"`
	MonthlyPrc    uint           `json:"monthly_prc"`
	ForCustom     uint           `json:"for_custom"`
	ForCustomOnce uint           `json:"for_custome_once"`
	LimitMonthly  uint           `json:"limit_monthly"`
}
