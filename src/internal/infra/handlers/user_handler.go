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

type UserHandler struct {
	repo repositories.UserRepositoryInterface
}

func NewUserHandler(repo repositories.UserRepositoryInterface) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if !handlers.IsLoggedUserAllowed(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	var newUser dtos.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.HandlerError{Message: "Some field were not sent or has invalid format"})
		return
	}
	user, errs := entities.NewUser(newUser)
	if len(errs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.HandlerError{Messages: utils.ErrorsToStrings(errs)})
		return
	}
	err = h.repo.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.HandlerError{Message: "Error while trying to register user"})
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.repo.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.HandlerError{Message: "Error while trying to list users"})
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	user, err := h.repo.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errors.HandlerError{Message: "User not found"})
		return
	}
	json.NewEncoder(w).Encode(user)
}
