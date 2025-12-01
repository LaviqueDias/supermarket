

// ============================================
// internal/usecases/client_usecase.go
// ============================================
package usecases

import (
	"errors"
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

func (uc *ClientUseCase) Login(email, password string) (*models.Client, string, error) {
	client, err := uc.repo.FindByEmail(email)
	if err != nil {
		return nil, "", errors.New("credenciais inválidas")
	}

	if !utils.CheckPasswordHash(password, client.PasswordHash) {
		return nil, "", errors.New("credenciais inválidas")
	}

	token, err := utils.GenerateToken(client.ID, client.Email, "client", "")
	if err != nil {
		return nil, "", err
	}

	return client, token, nil
}
