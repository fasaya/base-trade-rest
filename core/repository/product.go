package repository

import (
	"base-trade-rest/api/request"
	"base-trade-rest/core/helpers"
	"base-trade-rest/core/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

type IProductRepository interface {
	CreateProduct(*model.Product) (*model.Product, error)
	GetDetailProduct(uint) (*model.Product, error)
	GetAllProduct(req request.PaginateRequest) ([]model.Product, helpers.PaginationResponse, error)
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

func (r *ProductRepository) GetAllProduct(req request.PaginateRequest) ([]model.Product, helpers.PaginationResponse, error) {
	var products []model.Product
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

	query := r.db.Model(&model.Product{})

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

	err = query.Preload("Variants").Find(&products).Error
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

	return products, paginationResponse, nil
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
