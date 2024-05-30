package repository

import (
	"base-trade-rest/core/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	CreateUser(*model.User) (*model.User, error)
	GetDetailUser(int) (*model.User, error)
	GetAllUser() ([]model.User, error)
	UpdateUser(*model.User) (*model.User, error)
	DeleteUser(int) error
	GetUserByKey(string, interface{}) (*model.User, error)
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	var userRepo = UserRepository{}
	userRepo.db = db
	return &userRepo
}

func (r *UserRepository) CreateUser(user *model.User) (*model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetDetailUser(id int) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAllUser() ([]model.User, error) {
	var users []model.User
	err := r.db.Order("id desc").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) UpdateUser(user *model.User) (*model.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) DeleteUser(id int) error {
	var user model.User
	err := r.db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUserByKey(field string, value interface{}) (*model.User, error) {
	var user model.User
	err := r.db.Where(field+" = ?", value).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
