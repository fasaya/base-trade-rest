package resource

import (
	"time"
)

type AuthResource struct {
	ID        uint       `json:"id"`
	UUID      string     `json:"uuid"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"update_at"`

	Token string `json:"token,omitempty"`
}
