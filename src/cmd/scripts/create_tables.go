package scripts

import (
	"database/sql"
	"fmt"

	"github.com/CaiqueRibeiro/pizza-rute/src/configs"
)

func CreateTables() {
	fmt.Println("Creating tables")

	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open(cfg.DBDriver, dbUrl)
	if err != nil {
		panic(err.Error())
	}

	createUserTable(db)
	createIngredientsTable(db)
	createSauceTables(db)
	createPizzaTables(db)

	fmt.Println("Finished tables creation")
}

func createUserTable(db *sql.DB) {
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

func createIngredientsTable(db *sql.DB) {
	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS ingredients(
			id varchar(255),
			name varchar(255),
			stock int,
			primary key (id)
		)`)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err.Error())
	}
}

func createSauceTables(db *sql.DB) {
	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS sauces(
			id varchar(255),
			name varchar(255),
			is_spicy tinyint,
			primary key (id)
		)`)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err.Error())
	}

	stmt, err = db.Prepare(`
	CREATE TABLE IF NOT EXISTS sauce_items(
		ingredient_id varchar(255),
		sauce_id varchar(255),
		quantity int,
		primary key (ingredient_id, sauce_id),
		FOREIGN KEY (ingredient_id) REFERENCES ingredients(id),
		FOREIGN KEY (sauce_id) REFERENCES sauces(id)
	)
	`)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err.Error())
	}
}

func createPizzaTables(db *sql.DB) {
	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS pizzas(
			id varchar(255),
			name varchar(255),
			sauce_id varchar(255),
			size varchar(10),
			price int,
			primary key (id),
			FOREIGN KEY (sauce_id) REFERENCES sauces(id)
		)`)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err.Error())
	}

	stmt, err = db.Prepare(`
	CREATE TABLE IF NOT EXISTS pizza_items(
		ingredient_id varchar(255),
		pizza_id varchar(255),
		quantity int,
		primary key (ingredient_id, pizza_id),
		FOREIGN KEY (ingredient_id) REFERENCES ingredients(id),
		FOREIGN KEY (pizza_id) REFERENCES pizzas(id)
	)`)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err.Error())
	}
}
