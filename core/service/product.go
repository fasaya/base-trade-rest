package service

import (
	"base-trade-rest/api/request"
	"base-trade-rest/core/model"
	"base-trade-rest/core/repository"

	"github.com/google/uuid"
)

type ProductService struct {
	productRepo repository.IProductRepository
}

type IProductService interface {
	CreateProduct(*model.Product) (*model.Product, error)
	GetListProduct() ([]model.Product, error)
	GetDetailProduct(int) (*model.Product, error)
	UpdateProduct(*model.Product) (*model.Product, error)
	DeleteProduct(int) error
	GetProductByOwner(uint, uint) (*model.Product, error)
}

func NewProductService(productRepo repository.IProductRepository) *ProductService {
	var productService = ProductService{productRepo: productRepo}
	return &productService
}

func (s *ProductService) CreateProduct(product *model.Product) (*model.Product, error) {
	// Generate UUID
	product.UUID = uuid.New().String()

	// Create product
	result, err := s.productRepo.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *ProductService) GetListProduct() ([]model.Product, error) {
	result, err := s.productRepo.GetAllProduct()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ProductService) GetDetailProduct(id int) (*model.Product, error) {
	result, err := s.productRepo.GetDetailProduct(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ProductService) UpdateProduct(product *model.Product) (*model.Product, error) {
	result, err := s.productRepo.UpdateProduct(product)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ProductService) DeleteProduct(id int) error {
	err := s.productRepo.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) GetProductByOwner(creatorID uint, productUUID uint) (*model.Product, error) {
	var find = []request.FieldValueRequest{
		{Field: "uuid", Value: productUUID},
		{Field: "user_id", Value: creatorID},
	}

	result, err := s.productRepo.GetProductByMultipleKey(find)
	if err != nil {
		return nil, err
	}
	return result, nil
}
