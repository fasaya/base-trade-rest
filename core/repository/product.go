package repository

import (
	"base-trade-rest/api/request"
	"base-trade-rest/core/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

type IProductRepository interface {
	CreateProduct(*model.Product) (*model.Product, error)
	GetDetailProduct(uint) (*model.Product, error)
	GetAllProduct() ([]model.Product, error)
	UpdateProduct(*model.Product) (*model.Product, error)
	DeleteProduct(uint) error
	GetProductByKey(string, interface{}) (*model.Product, error)
	GetProductByMultipleKey([]request.FieldValueRequest) (*model.Product, error)
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	var productRepo = ProductRepository{}
	productRepo.db = db
	return &productRepo
}

func (r *ProductRepository) CreateProduct(product *model.Product) (*model.Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) GetDetailProduct(id uint) (*model.Product, error) {
	var product model.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) GetAllProduct() ([]model.Product, error) {
	var products []model.Product
	err := r.db.Order("id desc").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) UpdateProduct(product *model.Product) (*model.Product, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *ProductRepository) DeleteProduct(id uint) error {
	var product model.Product
	err := r.db.Where("id = ?", id).Delete(&product).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) GetProductByKey(field string, value interface{}) (*model.Product, error) {
	var product model.Product
	err := r.db.Where(field+" = ?", value).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) GetProductByMultipleKey(req []request.FieldValueRequest) (*model.Product, error) {
	var product model.Product

	query := r.db
	for _, v := range req {
		query = query.Where(v.Field+" = ?", v.Value)
	}

	err := query.First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
