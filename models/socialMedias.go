package models

import (
	"time"

	"gorm.io/gorm"
)

type SocialMedias struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	Linkedin  string `gorm:"type:varchar(255);not null"`
	Github    string `gorm:"type:varchar(255);not null"` // Exported field
	CreatedAt time.Time
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Find the latest SocialMedias entry
func socialMediasFindLatest() (SocialMedias, error) {
	var socialMedias SocialMedias
	result := DB.Order("id desc").First(&socialMedias)
	return socialMedias, result.Error
}

// Get all SocialMedias entries
func socialMediasList() ([]SocialMedias, error) {
	var socialMedias []SocialMedias
	result := DB.Order("id desc").Find(&socialMedias)
	return socialMedias, result.Error
}

// Find a SocialMedias entry by ID
func socialMediasFind(id uint64) *SocialMedias {
	var socialMedia SocialMedias
	result := DB.First(&socialMedia, id)
	if result.Error != nil {
		return nil
	}
	return &socialMedia
}
