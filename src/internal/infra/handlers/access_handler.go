package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/CaiqueRibeiro/pizza-rute/src/internal/dtos"
	"github.com/CaiqueRibeiro/pizza-rute/src/internal/infra/repositories"
)

type AccessHandler struct {
	repo repositories.UserRepositoryInterface
}

func NewAccessHandler(repo repositories.UserRepositoryInterface) *AccessHandler {
	return &AccessHandler{
		repo: repo,
	}
}

func (ar *AccessHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginDto dtos.LoginInput
	err := json.NewDecoder(r.Body).Decode(&loginDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: "Some field were not sent or has invalid format"})
		return
	}
	user, err := ar.repo.FindByEmail(loginDto.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Error{Message: "Invalid Credentials"})
		return
	}
	isValid := user.ValidatePassword(loginDto.Password)
	if !isValid {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Error{Message: "Invalid Credentials"})
		return
	}
	json.NewEncoder(w).Encode(user)
}
