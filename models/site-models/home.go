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
	SPhone        string         `json:"s_phone"`
	CPhone        string         `json:"c_phone"`
	Content       string         `json:"content"`
	Description   string         `json:"description"`
	Address       string         `json:"address"`
	Phone         string         `json:"phone"`
	Email         string         `json:"email"`
	Online        uint           `json:"online"`
	LastUpdate    string         `json:"last_update"`
	Link          Links          `json:"links"`
	CurrencyRates CurrencyRates  `json:"currency_rates"`
}

type CurrencyRates struct {
	ID           int    `json:"id"`
	Code         string `json:"Code"`
	Ccy          string `json:"Ccy"`
	CcyNmRU      string `json:"CcyNm_RU"`
	CcyNmUZ      string `json:"CcyNm_UZ"`
	CcyNmUZC     string `json:"CcyNm_UZC"`
	CcyNmEN      string `json:"CcyNm_EN"`
	Nominal      string `json:"Nominal"`
	Rate         string `json:"Rate"`
	Diff         string `json:"Diff"`
	Date         string `json:"Date"`
	HomeScreenId uint   `json:"home_screen_id"`
}

type Links struct {
	Id           uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	HomeScreenId uint           `json:"home_screen_id"`
	OkUrl        string         `json:"okru_url"`
	FbUrl        string         `json:"fb_url"`
	VkUrl        string         `json:"vk_url"`
	TgUrl        string         `json:"tg_url"`
}
