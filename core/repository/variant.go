package repository

import (
	"base-trade-rest/api/request"
	"base-trade-rest/core/model"

	"gorm.io/gorm"
)

type VariantRepository struct {
	db *gorm.DB
}

type IVariantRepository interface {
	CreateVariant(*model.Variant) (*model.Variant, error)
	GetDetailVariant(uint) (*model.Variant, error)
	GetAllVariant() ([]model.Variant, error)
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

func (r *VariantRepository) GetAllVariant() ([]model.Variant, error) {
	var variants []model.Variant
	err := r.db.Order("id desc").Find(&variants).Error
	if err != nil {
		return nil, err
	}
	return variants, nil
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
