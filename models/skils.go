package models

import (
	"time"

	"gorm.io/gorm"
)

type Skils struct {
	ID        uint64 `gorm:"primaryKey"`
	Skils     string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func SkilsFindLatest() (Skils, error) {
	var skils Skils
	result := DB.Order("id desc").First(&skils)
	return skils, result.Error
}
