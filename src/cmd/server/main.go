package main

import (
	"database/sql"
	"net/http"

	"github.com/CaiqueRibeiro/pizza-rute/src/internal/infra/handlers"
	"github.com/CaiqueRibeiro/pizza-rute/src/internal/infra/repositories"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/test")
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

	mux := http.NewServeMux()
	mux.HandleFunc("/users", userHandler.CreateUser)

	http.ListenAndServe(":3000", mux)
}