package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/CaiqueRibeiro/pizza-rute/src/internal/dtos"
	"github.com/CaiqueRibeiro/pizza-rute/src/internal/entities"
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

func (ih *IngredientsHandler) CreateIngredient(w http.ResponseWriter, r *http.Request) {
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
	err = ih.repo.Create(ingredient)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.HandlerError{Message: "Error while trying to create new ingredient"})
		return
	}
	json.NewEncoder(w).Encode(ingredient)
}
