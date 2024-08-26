package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/CaiqueRibeiro/pizza-rute/src/internal/dtos"
	"github.com/CaiqueRibeiro/pizza-rute/src/internal/entities"
	handlers "github.com/CaiqueRibeiro/pizza-rute/src/internal/infra/handlers/handler_helpers"
	"github.com/CaiqueRibeiro/pizza-rute/src/internal/infra/repositories"
	"github.com/CaiqueRibeiro/pizza-rute/src/pkg/errors"
	"github.com/CaiqueRibeiro/pizza-rute/src/pkg/utils"
)

type IngredientsHandler struct {
	repo repositories.IngredientRepositoryInterface
}

func NewIngredientsHandler(repo repositories.IngredientRepositoryInterface) *IngredientsHandler {
	return &IngredientsHandler{
		repo: repo,
	}
}

func (h *IngredientsHandler) CreateIngredient(w http.ResponseWriter, r *http.Request) {
	if !handlers.IsLoggedUserAllowed(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	var dto dtos.CreateIngredientInput
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.HandlerError{Message: "Some field were not sent or has invalid format"})
		return
	}
	ingredient, errs := entities.NewIngredient(dto)
	if len(errs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.HandlerError{Messages: utils.ErrorsToStrings(errs)})
		return
	}
	err = h.repo.Create(ingredient)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.HandlerError{Message: "Error while trying to create new ingredient"})
		return
	}
	json.NewEncoder(w).Encode(ingredient)
}

func (h *IngredientsHandler) UpdateIngredient(w http.ResponseWriter, r *http.Request) {
	if !handlers.IsLoggedUserAllowed(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	id := r.PathValue("id")
	ingredient, err := h.repo.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errors.HandlerError{Message: "Ingredient not found"})
		return
	}
	var dto dtos.UpdateIngredientInput
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.HandlerError{Message: "Some field were not sent or has invalid format"})
		return
	}
	if dto.Name != "" {
		ingredient.Name = dto.Name
	}
	if dto.Stock != nil {
		ingredient.Stock = *dto.Stock
	}
	err = h.repo.Update(ingredient)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.HandlerError{Message: "Error while trying to update an ingredient"})
		return
	}
	json.NewEncoder(w).Encode(ingredient)
}

func (h *IngredientsHandler) ListIngredients(w http.ResponseWriter, r *http.Request) {
	ingredients, err := h.repo.List()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errors.HandlerError{Message: "Error while trying to list ingredients"})
		return
	}
	json.NewEncoder(w).Encode(ingredients)
}

func (h *IngredientsHandler) GetIngredientById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	ingredient, err := h.repo.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errors.HandlerError{Message: "Ingredient not found"})
		return
	}
	json.NewEncoder(w).Encode(ingredient)
}
