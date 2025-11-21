package models

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"uniqueIndex;not null"`
	Phone     *string
	CreatedAt time.Time `gorm:"not null"`

	// Associations
	Orders []Order `gorm:"foreignKey:CustomerID"`
}
