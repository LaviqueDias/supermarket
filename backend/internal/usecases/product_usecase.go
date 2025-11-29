package usecases

import (
	"errors"
	"github.com/LaviqueDias/supermarket/internal/domain/models"
	"github.com/LaviqueDias/supermarket/internal/domain/repositories"
)

type ProductUseCase struct {
	repo *repositories.ProductRepository
}

func NewProductUseCase(repo *repositories.ProductRepository) *ProductUseCase {
	return &ProductUseCase{repo: repo}
}

func (uc *ProductUseCase) CreateProduct(p *models.Product) error {
	if p.Price < 0 {
		return errors.New("preço não pode ser negativo")
	}
	if p.StockQuantity < 0 {
		return errors.New("quantidade em estoque não pode ser negativa")
	}
	return uc.repo.Create(p)
}

func (uc *ProductUseCase) GetAllProducts() ([]models.Product, error) {
	return uc.repo.FindAll()
}

func (uc *ProductUseCase) GetProductByID(id int) (*models.Product, error) {
	return uc.repo.FindByID(id)
}

func (uc *ProductUseCase) UpdateProduct(p *models.Product) error {
	if p.Price < 0 {
		return errors.New("preço não pode ser negativo")
	}
	return uc.repo.Update(p)
}

func (uc *ProductUseCase) DeleteProduct(id int) error {
	return uc.repo.Delete(id)
}
