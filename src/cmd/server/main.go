package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"

	"github.com/CaiqueRibeiro/pizza-rute/src/cmd/scripts"
	"github.com/CaiqueRibeiro/pizza-rute/src/configs"
	"github.com/CaiqueRibeiro/pizza-rute/src/internal/infra/handlers"
	"github.com/CaiqueRibeiro/pizza-rute/src/internal/infra/middlewares"
	"github.com/CaiqueRibeiro/pizza-rute/src/internal/infra/repositories"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	ct := flag.Bool("ct", false, "Creates tables before initialization")
	flag.Parse()
	if *ct {
		scripts.CreateTables()
	}
}

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open(cfg.DBDriver, dbUrl)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepository)
	accessHandler := handlers.NewAccessHandler(userRepository)

	mux := http.NewServeMux()

	mux.Handle("POST /login", middlewares.WithContext(accessHandler.Login))

	mux.Handle("POST /users", middlewares.Authorized(userHandler.CreateUser))
	mux.Handle("GET /users", middlewares.Authorized(userHandler.ListUsers))
	mux.Handle("GET /users/{id}", middlewares.Authorized(userHandler.GetUserByID))

	http.ListenAndServe(cfg.WebServerPort, mux)
}
