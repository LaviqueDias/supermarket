package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/LaviqueDias/supermarket/internal/usecases"
	"github.com/LaviqueDias/supermarket/pkg/utils"
)

type EmployeeHandler struct {
	useCase *usecases.EmployeeUseCase
}

func NewEmployeeHandler(uc *usecases.EmployeeUseCase) *EmployeeHandler {
	return &EmployeeHandler{useCase: uc}
}

func (h *EmployeeHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	employee, err := h.useCase.RegisterEmployee(input.Name, input.Email, input.Password, input.Role)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusCreated, employee)
}

func (h *EmployeeHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	employee, token, err := h.useCase.Login(input.Email, input.Password)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err.Error())
		return
	}

	response := map[string]interface{}{
		"employee": employee,
		"token":    token,
	}

	utils.RespondJSON(w, http.StatusOK, response)
}

func (h *EmployeeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	employees, err := h.useCase.GetAllEmployees()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(w, http.StatusOK, employees)
}