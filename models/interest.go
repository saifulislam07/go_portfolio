package models

import (
	"time"

	"gorm.io/gorm"
)

type Interests struct {
	ID        uint64 `gorm:"primaryKey"`
	Interest  string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func InterestsFindLatest() (Interests, error) {
	var interests Interests
	result := DB.Order("id desc").First(&interests)
	return interests, result.Error
}

func InterestList() ([]Interests, error) {
	var interests []Interests
	result := DB.Order("id desc").Find(&interests)
	return interests, result.Error
}

func InterestFind(id uint64) *Interests {
	var interest Interests
	result := DB.First(&interest, id)
	if result.Error != nil {
		return nil
	}
	return &interest
}
