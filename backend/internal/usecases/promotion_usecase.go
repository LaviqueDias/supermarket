package usecases

import (
	"github.com/LaviqueDias/supermarket/internal/domain/models"
	"github.com/LaviqueDias/supermarket/internal/domain/repositories"
)

type PromotionUseCase struct {
	repo *repositories.PromotionRepository
}

func NewPromotionUseCase(repo *repositories.PromotionRepository) *PromotionUseCase {
	return &PromotionUseCase{repo: repo}
}

func (uc *PromotionUseCase) CreatePromotion(p *models.Promotion) error {
	return uc.repo.Create(p)
}

func (uc *PromotionUseCase) GetAllPromotions() ([]models.Promotion, error) {
	return uc.repo.FindAll()
}

func (uc *PromotionUseCase) AddProductToPromotion(promotionID, productID int) error {
	return uc.repo.AddProduct(promotionID, productID)
}
