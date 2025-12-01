
// ============================================
// internal/usecases/employee_usecase.go
// ============================================
package usecases

import (
	"errors"
	"github.com/LaviqueDias/supermarket/internal/domain/models"
	"github.com/LaviqueDias/supermarket/internal/domain/repositories"
	"github.com/LaviqueDias/supermarket/pkg/utils"
)

type EmployeeUseCase struct {
	repo *repositories.EmployeeRepository
}

func NewEmployeeUseCase(repo *repositories.EmployeeRepository) *EmployeeUseCase {
	return &EmployeeUseCase{repo: repo}
}

func (uc *EmployeeUseCase) RegisterEmployee(name, email, password, role string) (*models.Employee, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	employee := &models.Employee{
		Name:         name,
		Email:        email,
		PasswordHash: hashedPassword,
		Role:         role,
	}

	err = uc.repo.Create(employee)
	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (uc *EmployeeUseCase) GetAllEmployees() ([]models.Employee, error) {
	return uc.repo.FindAll()
}

func (uc *EmployeeUseCase) Login(email, password string) (*models.Employee, string, error) {
	employee, err := uc.repo.FindByEmail(email)
	if err != nil {
		return nil, "", errors.New("credenciais inválidas")
	}

	if !utils.CheckPasswordHash(password, employee.PasswordHash) {
		return nil, "", errors.New("credenciais inválidas")
	}

	token, err := utils.GenerateToken(employee.ID, employee.Email, "employee", employee.Role)
	if err != nil {
		return nil, "", err
	}

	return employee, token, nil
}