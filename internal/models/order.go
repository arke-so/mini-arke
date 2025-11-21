package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	CustomerID  uuid.UUID `gorm:"type:uuid;not null;index"`
	OrderDate   time.Time `gorm:"not null;index"`
	Status      string    `gorm:"not null;default:'pending';index"`
	TotalAmount float64   `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   *time.Time

	// Associations
	Customer   Customer    `gorm:"foreignKey:CustomerID"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
}
