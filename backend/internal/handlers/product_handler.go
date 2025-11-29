package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/LaviqueDias/supermarket/internal/domain/models"
	"github.com/LaviqueDias/supermarket/internal/usecases"
	"github.com/LaviqueDias/supermarket/pkg/utils"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	useCase *usecases.ProductUseCase
}

func NewProductHandler(uc *usecases.ProductUseCase) *ProductHandler {
	return &ProductHandler{useCase: uc}
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.useCase.CreateProduct(&p); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusCreated, p)
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.useCase.GetAllProducts()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(w, http.StatusOK, products)
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	product, err := h.useCase.GetProductByID(id)
	if err != nil {
		utils.RespondError(w, http.StatusNotFound, "Produto não encontrado")
		return
	}

	utils.RespondJSON(w, http.StatusOK, product)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var p models.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	p.ID = id
	if err := h.useCase.UpdateProduct(&p); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, p)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if err := h.useCase.DeleteProduct(id); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}