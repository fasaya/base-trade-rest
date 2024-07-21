package service

import (
	"base-trade-rest/api/request"
	"base-trade-rest/core/helpers"
	"base-trade-rest/core/model"
	"base-trade-rest/core/repository"

	"github.com/google/uuid"
)

type VariantService struct {
	variantRepo repository.IVariantRepository
}

type IVariantService interface {
	CreateVariant(*model.Variant) (*model.Variant, error)
	GetListVariant(req request.PaginateRequest) ([]model.Variant, helpers.PaginationResponse, error)
	GetDetailVariantByUUID(string) (*model.Variant, error)
	UpdateVariant(*model.Variant) (*model.Variant, error)
	DeleteVariantByUUID(string) error
	GetVariantByOwner(creatorID uint, variantUUID string) (*model.Variant, error)
}

func NewVariantService(variantRepo repository.IVariantRepository) *VariantService {
	var variantService = VariantService{variantRepo: variantRepo}
	return &variantService
}

func (s *VariantService) CreateVariant(variant *model.Variant) (*model.Variant, error) {
	// Generate UUID
	variant.UUID = uuid.New().String()

	// Create variant
	result, err := s.variantRepo.CreateVariant(variant)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *VariantService) GetListVariant(req request.PaginateRequest) ([]model.Variant, helpers.PaginationResponse, error) {
	result, paginationResponse, err := s.variantRepo.GetAllVariant(req)
	if err != nil {
		return nil, helpers.PaginationResponse{}, err
	}
	return result, paginationResponse, nil
}

func (s *VariantService) GetDetailVariantByUUID(uuid string) (*model.Variant, error) {
	result, err := s.variantRepo.GetVariantByKey("uuid", uuid)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *VariantService) UpdateVariant(variant *model.Variant) (*model.Variant, error) {
	result, err := s.variantRepo.UpdateVariant(variant)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *VariantService) DeleteVariantByUUID(uuid string) error {
	variant, err := s.variantRepo.GetVariantByKey("uuid", uuid)
	if err != nil {
		return err
	}

	err = s.variantRepo.DeleteVariant(variant.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *VariantService) GetVariantByOwner(creatorID uint, variantUUID string) (*model.Variant, error) {
	var find = []request.FieldValueRequest{
		{Field: "uuid", Value: variantUUID},
		{Field: "user_id", Value: creatorID},
	}

	result, err := s.variantRepo.GetVariantByMultipleKey(find)
	if err != nil {
		return nil, err
	}
	return result, nil
}
