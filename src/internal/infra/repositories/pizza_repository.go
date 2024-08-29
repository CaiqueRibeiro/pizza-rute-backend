package repositories

import (
	"database/sql"

	"github.com/CaiqueRibeiro/pizza-rute/src/internal/entities"
)

type PizzaRepositoryInterface interface {
	Create(pizza *entities.Pizza) error
}

type PizzaRepository struct {
	db *sql.DB
}

func NewPizzaRepository(db *sql.DB) *PizzaRepository {
	return &PizzaRepository{
		db: db,
	}
}

// TODO: transform in Unity Of Work
func (pr *PizzaRepository) Create(pizza *entities.Pizza) error {
	tx, err := pr.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO pizzas values(?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		pizza.ID.String(),
		pizza.Name,
		pizza.Sauce.ID.String(),
		pizza.Size,
		pizza.Price*100,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	relationStmt, err := tx.Prepare("INSERT INTO pizza_items values (?,?,?)")
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insert each ingredient into the pizza_items table
	for _, ingredient := range pizza.Ingredients {
		_, err = relationStmt.Exec(
			ingredient.IngredientID.String(),
			pizza.ID.String(),
			ingredient.Quantity,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
