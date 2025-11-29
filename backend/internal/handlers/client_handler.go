package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/LaviqueDias/supermarket/internal/usecases"
	"github.com/LaviqueDias/supermarket/pkg/utils"
)

type ClientHandler struct {
	useCase *usecases.ClientUseCase
}

func NewClientHandler(uc *usecases.ClientUseCase) *ClientHandler {
	return &ClientHandler{useCase: uc}
}

func (h *ClientHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	client, err := h.useCase.RegisterClient(input.Name, input.Email, input.Password)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusCreated, client)
}

func (h *ClientHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	clients, err := h.useCase.GetAllClients()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(w, http.StatusOK, clients)
}