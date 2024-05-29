package models

import (
	"time"

	"github.com/google/uuid"
)

type Variant struct {
	ID        uint       `gorm:"primaryKey"`
	UUID      uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()" json:"uuid"`
	Name      string     `gorm:"not null"`
	Quantity  int        `gorm:"not null"`
	ProductID uint       `gorm:"not null"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"update_at"`
}
