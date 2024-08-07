package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/CaiqueRibeiro/pizza-rute/src/configs"
	"github.com/CaiqueRibeiro/pizza-rute/src/internal/infra/handlers"
	"github.com/CaiqueRibeiro/pizza-rute/src/internal/infra/middlewares"
	"github.com/CaiqueRibeiro/pizza-rute/src/internal/infra/repositories"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var cfg *configs.Conf

func init() {
	var err error
	cfg, err = configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err = sql.Open(cfg.DBDriver, dbUrl)
	if err != nil {
		panic(err.Error())
	}
	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS users(
		id varchar(255),
		name varchar(80),
		surname varchar(255),
		email varchar(80),
		photo_url varchar(255),
		job_position varchar(255),
		password varchar(255),
		primary key (id))
	`)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err.Error())
	}
}

func main() {
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
