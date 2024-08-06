package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/CaiqueRibeiro/pizza-rute/src/internal/dtos"
	"github.com/CaiqueRibeiro/pizza-rute/src/internal/entities"
	"github.com/CaiqueRibeiro/pizza-rute/src/pkg/utils"
)

type UserHandler struct{}

type Error struct {
	Message  string   `json:"error,omitempty"`
	Messages []string `json:"errors,omitempty"`
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var newUser dtos.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: "Some field were not sent or has invalid format"})
		return
	}
	user, errs := entities.NewUser(newUser)
	if len(errs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Messages: utils.ErrorsToStrings(errs)})
		return
	}
	json.NewEncoder(w).Encode(user)
}
