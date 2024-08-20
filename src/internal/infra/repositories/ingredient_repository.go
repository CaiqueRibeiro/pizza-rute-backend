package repositories

import (
	"context"
	"database/sql"

	"github.com/CaiqueRibeiro/pizza-rute/src/internal/entities"
)

type IngredientRepositoryInterface interface {
	Create(ingredient *entities.Ingredient) error
	List() ([]*entities.Ingredient, error)
	FindByID(id string) (*entities.Ingredient, error)
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

func (ir *IngredientRepository) List() ([]*entities.Ingredient, error) {
	rows, err := ir.db.Query("SELECT * FROM ingredients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ingredients []*entities.Ingredient
	for rows.Next() {
		var i entities.Ingredient
		err = rows.Scan(&i.ID, &i.Name, &i.Stock)
		if err != nil {
			return nil, err
		}
		ingredients = append(ingredients, &i)
	}
	return ingredients, nil
}

func (ir *IngredientRepository) FindByID(id string) (*entities.Ingredient, error) {
	stmt, err := ir.db.Prepare("SELECT * FROM ingredients WHERE id=?")
	if err != nil {
		return nil, err
	}
	var i entities.Ingredient
	err = stmt.QueryRowContext(context.Background(), id).Scan(&i.ID, &i.Name, &i.Stock)
	if err != nil {
		return nil, err
	}
	return &i, nil
}
