package scripts

import (
	"database/sql"
	"fmt"

	"github.com/CaiqueRibeiro/pizza-rute/src/configs"
)

func CreateTables() {
	fmt.Println("Creating Tables")

	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open(cfg.DBDriver, dbUrl)
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
