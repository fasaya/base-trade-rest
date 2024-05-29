package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UUID      uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()" json:"uuid"`
	Name      string     `gorm:"not null" json:"name"`
	ImageURL  *string    `json:"image_url"`
	UserID    uint       `gorm:"not null" json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"update_at"`

	Variants []Variant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"variants"`
}
