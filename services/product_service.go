package services

import (
	"inventory/models"
	"inventory/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepo
}

func NewProductService() *ProductService {
	return &ProductService{repo: repositories.NewProductRepo()}
}

func (s *ProductService) SaveProduct(product *models.Product) (*models.Product, error) {
	return s.repo.SaveProduct(product)
}

func (s *ProductService) GetProduct(id uint64) (*models.Product, error) {
	return s.repo.GetProduct(id)
}

func (s *ProductService) UpdateProduct(id uint64, product *models.Product) (*models.Product, error) {
	return s.repo.UpdateProduct(id, product)
}

func (s *ProductService) DeleteProduct(id uint64) error {
	return s.repo.DeleteProduct(id)
}

func (s *ProductService) ListProducts() ([]models.Product, error) {
	return s.repo.ListProducts()
}
