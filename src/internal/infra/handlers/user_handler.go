package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/CaiqueRibeiro/pizza-rute/src/internal/dtos"
	"github.com/CaiqueRibeiro/pizza-rute/src/internal/entities"
)

type UserHandler struct{}

type Error struct {
	Message string
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
		json.NewEncoder(w).Encode(Error{Message: "Fields sent seem invalid"})
		return
	}
	user := &entities.User{
		ID:          "1234",
		Name:        newUser.Name,
		Surname:     newUser.Surname,
		Email:       newUser.Email,
		PhotoUrl:    newUser.PhotoUrl,
		JobPosition: newUser.JobPosition,
	}
	json.NewEncoder(w).Encode(user)
}
