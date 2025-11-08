package models

import (
	"time"

	"gorm.io/gorm"
)

type Portfolios struct {
	ID        uint64 `gorm:"primaryKey"`
	Image     string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func PortfolioFindLatest() (Portfolios, error) {
	var portfolio Portfolios
	result := DB.Order("id desc").First(&portfolio)
	return portfolio, result.Error
}

func PortfolioList() ([]Portfolios, error) {
	var portfolio []Portfolios
	result := DB.Order("id desc").Find(&portfolio)
	return portfolio, result.Error
}

func PortfolioFind(id uint64) *Portfolios {
	var portfolio Portfolios
	result := DB.First(&portfolio, id)
	if result.Error != nil {
		return nil
	}
	return &portfolio
}
