package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/LaviqueDias/supermarket/internal/usecases"
	"github.com/LaviqueDias/supermarket/pkg/utils"

	"github.com/gorilla/mux"
)

type CartHandler struct {
	useCase *usecases.CartUseCase
}

func NewCartHandler(uc *usecases.CartUseCase) *CartHandler {
	return &CartHandler{useCase: uc}
}

func (h *CartHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ClientID  int `json:"client_id"`
		ProductID int `json:"product_id"`
		Quantity  int `json:"quantity"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.useCase.AddToCart(input.ClientID, input.ProductID, input.Quantity); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "Produto adicionado ao carrinho"})
}

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	clientID, _ := strconv.Atoi(params["client_id"])

	cart, err := h.useCase.GetCart(clientID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, cart)
}

func (h *CartHandler) RemoveItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemID, _ := strconv.Atoi(params["item_id"])

	if err := h.useCase.RemoveFromCart(itemID); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}