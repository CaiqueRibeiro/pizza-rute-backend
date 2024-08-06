package main

import (
	"net/http"

	"github.com/CaiqueRibeiro/pizza-rute/src/internal/infra/handlers"
)

func main() {
	mux := http.NewServeMux()
	userHandler := handlers.NewUserHandler()
	mux.HandleFunc("/users", userHandler.CreateUser)
	http.ListenAndServe(":3000", mux)
}
