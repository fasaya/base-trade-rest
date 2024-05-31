package model

import (
	"time"
)

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UUID      string     `gorm:"not null" json:"uuid"`
	Name      string     `gorm:"not null" json:"name"`
	Email     string     `gorm:"unique;not null" json:"email"`
	Password  string     `gorm:"not null" json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"update_at"`

	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"products,omitempty"`
}

// Override the table name
func (User) TableName() string {
	return "admins"
}
