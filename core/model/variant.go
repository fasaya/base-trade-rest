package model

import (
	"time"
)

type Variant struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UUID      string     `gorm:"not null" json:"uuid"`
	Name      string     `gorm:"not null"`
	Quantity  int        `gorm:"not null"`
	ProductID uint       `gorm:"not null"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"update_at"`

	Product Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product,omitempty"`
}
