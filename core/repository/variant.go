package repository

import (
	"base-trade-rest/api/request"
	"base-trade-rest/core/helpers"
	"base-trade-rest/core/model"

	"gorm.io/gorm"
)

type VariantRepository struct {
	db *gorm.DB
}

type IVariantRepository interface {
	CreateVariant(*model.Variant) (*model.Variant, error)
	GetDetailVariant(uint) (*model.Variant, error)
	GetAllVariant(req request.PaginateRequest) ([]model.Variant, helpers.PaginationResponse, error)
	UpdateVariant(*model.Variant) (*model.Variant, error)
	DeleteVariant(uint) error
	GetVariantByKey(string, interface{}) (*model.Variant, error)
	GetVariantByMultipleKey([]request.FieldValueRequest) (*model.Variant, error)
}

func NewVariantRepository(db *gorm.DB) *VariantRepository {
	var variantRepo = VariantRepository{}
	variantRepo.db = db
	return &variantRepo
}

func (r *VariantRepository) CreateVariant(variant *model.Variant) (*model.Variant, error) {
	err := r.db.Create(&variant).Error
	if err != nil {
		return nil, err
	}
	return variant, nil
}

func (r *VariantRepository) GetDetailVariant(id uint) (*model.Variant, error) {
	var variant model.Variant
	err := r.db.First(&variant, id).Error
	if err != nil {
		return nil, err
	}
	return &variant, nil
}

func (r *VariantRepository) GetAllVariant(req request.PaginateRequest) ([]model.Variant, helpers.PaginationResponse, error) {
	var variants []model.Variant
	var total int64
	var lastPage int64

	// Set default limit
	if req.Limit <= 0 {
		req.Limit = 10
	}

	// Set default page number
	if req.Page <= 0 {
		req.Page = 1
	}

	// if req.Sort != "" && req.Sort[0] == '-' {
	// 	req.Sort = req.Sort[1:] + " DESC"
	// }

	query := r.db.Model(&model.Variant{})

	if req.Search != "" {
		query = query.Where("name LIKE ?", "%"+req.Search+"%")
	}

	// Count the total number of records matching the search criteria
	err := query.Count(&total).Error
	if err != nil {
		return nil, helpers.PaginationResponse{}, err
	}

	if req.Limit > 0 {
		query = query.Limit(req.Limit).Offset((req.Page - 1) * req.Limit)
	}

	query = query.Order("id desc")

	err = query.Find(&variants).Error
	if err != nil {
		return nil, helpers.PaginationResponse{}, err
	}

	// Calculate the last page, ensure req.Limit is greater than zero
	if req.Limit > 0 {
		lastPage = (total + int64(req.Limit) - 1) / int64(req.Limit)
	} else {
		lastPage = 0
	}

	paginationResponse := helpers.PaginationResponse{
		Page:     int64(req.Page),
		Limit:    int64(req.Limit),
		LastPage: lastPage,
		Total:    total,
	}

	return variants, paginationResponse, nil
}

func (r *VariantRepository) UpdateVariant(variant *model.Variant) (*model.Variant, error) {
	err := r.db.Save(&variant).Error
	if err != nil {
		return nil, err
	}

	return variant, nil
}

func (r *VariantRepository) DeleteVariant(id uint) error {
	var variant model.Variant
	err := r.db.Where("id = ?", id).Delete(&variant).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *VariantRepository) GetVariantByKey(field string, value interface{}) (*model.Variant, error) {
	var variant model.Variant
	err := r.db.Preload("Product").Where(field+" = ?", value).First(&variant).Error
	if err != nil {
		return nil, err
	}
	return &variant, nil
}

func (r *VariantRepository) GetVariantByMultipleKey(req []request.FieldValueRequest) (*model.Variant, error) {
	var variant model.Variant

	query := r.db
	for _, v := range req {
		query = query.Where(v.Field+" = ?", v.Value)
	}

	err := query.First(&variant).Error
	if err != nil {
		return nil, err
	}
	return &variant, nil
}
