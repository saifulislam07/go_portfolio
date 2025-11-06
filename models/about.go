package models

import (
	"time"

	"gorm.io/gorm"
)

type About struct {
	ID         uint64 `gorm:"primaryKey"`
	AboutMe    string `gorm:"type:text"`
	AboutImage string `gorm:"type:text"`
	CreatedAt  time.Time
	UpdatedAt  time.Time      `gorm:"index"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func AboutFindLatest() (About, error) {
	var about About
	result := DB.Order("id desc").First(&about)
	return about, result.Error
}
