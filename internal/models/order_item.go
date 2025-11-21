// order_item.go
package models

import (
	"github.com/google/uuid"
)

type OrderItem struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrderID   uuid.UUID `gorm:"type:uuid;not null;index"`
	ProductID uuid.UUID `gorm:"type:uuid;not null;index"`
	Quantity  int       `gorm:"not null"`
	UnitPrice float64   `gorm:"not null"`
	Subtotal  float64   `gorm:"not null"`

	// Associations
	Order   Order   `gorm:"foreignKey:OrderID"`
	Product Product `gorm:"foreignKey:ProductID"`
}
