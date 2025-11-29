package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/LaviqueDias/supermarket/internal/domain/models"
	"github.com/LaviqueDias/supermarket/internal/usecases"
	"github.com/LaviqueDias/supermarket/pkg/utils"
)

type PromotionHandler struct {
	useCase *usecases.PromotionUseCase
}

func NewPromotionHandler(uc *usecases.PromotionUseCase) *PromotionHandler {
	return &PromotionHandler{useCase: uc}
}

func (h *PromotionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var p models.Promotion
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.useCase.CreatePromotion(&p); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusCreated, p)
}

func (h *PromotionHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	promotions, err := h.useCase.GetAllPromotions()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(w, http.StatusOK, promotions)
}

func (h *PromotionHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	var input struct {
		PromotionID int `json:"promotion_id"`
		ProductID   int `json:"product_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.useCase.AddProductToPromotion(input.PromotionID, input.ProductID); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "Produto adicionado à promoção"})
}