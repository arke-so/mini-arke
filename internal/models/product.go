package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	SKU           string    `gorm:"uniqueIndex;not null"`
	Name          string    `gorm:"not null"`
	Description   *string   `gorm:"type:text"`
	Price         float64   `gorm:"not null"`
	StockQuantity int       `gorm:"not null;default:0"`
	CreatedAt     time.Time `gorm:"not null"`
	UpdatedAt     *time.Time
}
