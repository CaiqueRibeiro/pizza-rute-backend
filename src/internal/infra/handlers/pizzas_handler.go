package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/CaiqueRibeiro/pizza-rute/src/internal/dtos"
	"github.com/CaiqueRibeiro/pizza-rute/src/internal/entities"
	"github.com/CaiqueRibeiro/pizza-rute/src/internal/infra/handlers/handler_helpers"
	"github.com/CaiqueRibeiro/pizza-rute/src/internal/infra/repositories"
	"github.com/CaiqueRibeiro/pizza-rute/src/pkg/errors"
	"github.com/CaiqueRibeiro/pizza-rute/src/pkg/utils"
)

type PizzasHandler struct {
	repo repositories.PizzaRepositoryInterface
}

func NewPizzaHandler(repo repositories.PizzaRepositoryInterface) *PizzasHandler {
	return &PizzasHandler{
		repo: repo,
	}
}

func (h *PizzasHandler) CreatePizza(w http.ResponseWriter, r *http.Request) {
	if !handler_helpers.IsLoggedUserAllowed(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	var dto dtos.CreatePizzaInput
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.HandlerError{Message: "Some field were not sent or has invalid format"})
		return
	}
	pizza, errs := entities.NewPizza(dto)
	if len(errs) > 0 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.HandlerError{Messages: utils.ErrorsToStrings(errs)})
		return
	}
	err = h.repo.Create(pizza)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.HandlerError{Message: "Error while trying to create a pizza receipt"})
		return
	}
	json.NewEncoder(w).Encode(pizza)
}

func (h *PizzasHandler) ListPizzas(w http.ResponseWriter, r *http.Request) {
	pizzas, err := h.repo.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.HandlerError{Message: "Error while trying to list pizza receipts"})
		return
	}
	json.NewEncoder(w).Encode(pizzas)
}
