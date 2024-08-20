package repositories

import (
	"database/sql"

	"github.com/CaiqueRibeiro/pizza-rute/src/internal/entities"
)

type IngredientRepositoryInterface interface {
	Create(ingredient *entities.Ingredient) error
}

type IngredientRepository struct {
	db *sql.DB
}

func NewIngredientRepository(db *sql.DB) *IngredientRepository {
	return &IngredientRepository{
		db: db,
	}
}

func (ir *IngredientRepository) Create(ingredient *entities.Ingredient) error {
	stmt, err := ir.db.Prepare("INSERT INTO ingredients values (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		ingredient.ID.String(),
		ingredient.Name,
		ingredient.Stock*100,
	)
	if err != nil {
		return err
	}
	return nil
}
