package service

import (
	"base-trade-rest/api/request"
	"base-trade-rest/core/helpers"
	"base-trade-rest/core/model"
	"base-trade-rest/core/repository"
	"errors"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type UserService struct {
	userRepo repository.IUserRepository
}

type IUserService interface {
	CreateUser(*request.AuthRegisterRequest) (*model.User, error)
	GetListUser() ([]model.User, error)
	GetDetailUser(int) (*model.User, error)
	UpdateUser(*model.User) (*model.User, error)
	DeleteUser(int) error
	ValidateCredentials(request.AuthLoginRequest) (*model.User, error)
}

func NewUserService(userRepo repository.IUserRepository) *UserService {
	var userService = UserService{}
	userService.userRepo = userRepo
	return &userService
}

func (s *UserService) CreateUser(request *request.AuthRegisterRequest) (*model.User, error) {

	// Transform
	var user model.User
	copier.Copy(&user, &request)

	// Check if user already exist
	_, err := s.userRepo.GetUserByKey("email", user.Email)
	if err == nil {
		return nil, errors.New("user already exist")
	}

	// Generate UUID
	user.UUID = uuid.New().String()

	// Hash Password
	user.Password = helpers.HashPass(request.Password)

	// Create user
	result, err := s.userRepo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *UserService) GetListUser() ([]model.User, error) {
	result, err := s.userRepo.GetAllUser()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *UserService) GetDetailUser(id int) (*model.User, error) {
	result, err := s.userRepo.GetDetailUser(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *UserService) UpdateUser(user *model.User) (*model.User, error) {
	result, err := s.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *UserService) DeleteUser(id int) error {
	err := s.userRepo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) ValidateCredentials(request request.AuthLoginRequest) (*model.User, error) {
	// Check if user already exist
	user, err := s.userRepo.GetUserByKey("email", request.Email)
	if err != nil {
		return nil, errors.New("user not found or invalid password")
	}

	comparePass := helpers.ComparePass([]byte(user.Password), []byte(request.Password))
	if !comparePass {
		return nil, errors.New("user not found or invalid password")
	}

	return user, nil
}
