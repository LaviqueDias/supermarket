package usecases

import (
	"github.com/LaviqueDias/supermarket/internal/domain/models"
	"github.com/LaviqueDias/supermarket/internal/domain/repositories"
	"github.com/LaviqueDias/supermarket/pkg/utils"
)

type ClientUseCase struct {
	repo *repositories.ClientRepository
}

func NewClientUseCase(repo *repositories.ClientRepository) *ClientUseCase {
	return &ClientUseCase{repo: repo}
}

func (uc *ClientUseCase) RegisterClient(name, email, password string) (*models.Client, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	client := &models.Client{
		Name:         name,
		Email:        email,
		PasswordHash: hashedPassword,
	}

	err = uc.repo.Create(client)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (uc *ClientUseCase) GetAllClients() ([]models.Client, error) {
	return uc.repo.FindAll()
}
