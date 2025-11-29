package usecases

import (
	"errors"
	"github.com/LaviqueDias/supermarket/internal/domain/repositories"
)

type CartUseCase struct {
	cartRepo    *repositories.CartRepository
	productRepo *repositories.ProductRepository
}

func NewCartUseCase(cartRepo *repositories.CartRepository, productRepo *repositories.ProductRepository) *CartUseCase {
	return &CartUseCase{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (uc *CartUseCase) AddToCart(clientID, productID, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantidade deve ser maior que zero")
	}

	cartID, err := uc.cartRepo.GetCartByClientID(clientID)
	if err != nil {
		return errors.New("carrinho não encontrado")
	}

	product, err := uc.productRepo.FindByID(productID)
	if err != nil {
		return errors.New("produto não encontrado")
	}

	return uc.cartRepo.AddItem(cartID, productID, quantity, product.Price)
}

func (uc *CartUseCase) GetCart(clientID int) (map[string]interface{}, error) {
	items, total, err := uc.cartRepo.GetCartItems(clientID)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"client_id": clientID,
		"items":     items,
		"total":     total,
	}, nil
}

func (uc *CartUseCase) RemoveFromCart(itemID int) error {
	return uc.cartRepo.RemoveItem(itemID)
}
