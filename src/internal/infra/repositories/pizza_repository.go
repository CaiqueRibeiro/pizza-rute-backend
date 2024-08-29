package repositories

import (
	"database/sql"

	"github.com/CaiqueRibeiro/pizza-rute/src/internal/entities"
)

type PizzaRepositoryInterface interface {
	Create(pizza *entities.Pizza) error
	List() ([]*entities.Pizza, error)
}

type PizzaRepository struct {
	db *sql.DB
}

func NewPizzaRepository(db *sql.DB) *PizzaRepository {
	return &PizzaRepository{
		db: db,
	}
}

// TODO: transform in Unit Of Work
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

func (pr *PizzaRepository) List() ([]*entities.Pizza, error) {
	rows, err := pr.db.Query("SELECT * from pizzas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var pizzas []*entities.Pizza
	for rows.Next() {
		var p entities.Pizza
		err = rows.Scan(
			&p.ID,
			&p.Name,
			&p.Sauce.ID,
			&p.Size,
			&p.Price,
		)
		if err != nil {
			return nil, err
		}

		// Fetch ingredients for this pizza
		ingredientRows, err := pr.db.Query(`
				SELECT ingredient_id, quantity FROM pizza_items WHERE pizza_id = ?
			`, p.ID.String())
		if err != nil {
			return nil, err
		}
		defer ingredientRows.Close()

		var ingredients []entities.PizzaItem
		for ingredientRows.Next() {
			var item entities.PizzaItem
			err = ingredientRows.Scan(
				&item.IngredientID,
				&item.Quantity,
			)
			if err != nil {
				return nil, err
			}
			ingredients = append(ingredients, item)
		}

		p.Ingredients = ingredients
		p.Price = p.Price / 100
		pizzas = append(pizzas, &p)
	}

	return pizzas, nil
}
