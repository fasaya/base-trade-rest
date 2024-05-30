package model

import (
	"base-trade-rest/core/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UUID      string     `gorm:"not null" json:"uuid"`
	Name      string     `gorm:"not null" json:"name" form:"name" valid:"required~Your name is required"`
	Email     string     `gorm:"unique;not null" json:"email" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password  string     `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"update_at"`

	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"products,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errValidate := govalidator.ValidateStruct(u)

	if errValidate != nil {
		err = errValidate
		return err
	}

	u.Password = helpers.HashPass(u.Password)

	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errValidate := govalidator.ValidateStruct(u)

	if errValidate != nil {
		err = errValidate
		return err
	}

	u.Password = helpers.HashPass(u.Password)

	return nil
}
