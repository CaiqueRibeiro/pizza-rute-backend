package entities

import (
	"github.com/CaiqueRibeiro/pizza-rute/src/internal/dtos"
	"github.com/CaiqueRibeiro/pizza-rute/src/pkg/entities"
)

type Ingredient struct {
	ID    entities.ID `json:"id"`
	Name  string      `json:"name,omitempty"`
	Stock float32     `json:"stock"`
}

func NewIngredient(newIngredient dtos.CreateIngredientInput) (*Ingredient, []error) {
	var errs []error

	if newIngredient.Name == "" {
		errs = append(errs, ErrNameIsRequired)
	}

	if newIngredient.Stock <= 0 {
		errs = append(errs, ErrStockIsRequired)
	}

	if len(errs) > 0 {
		return nil, errs
	}

	return &Ingredient{
		ID:    entities.NewID(),
		Name:  newIngredient.Name,
		Stock: newIngredient.Stock,
	}, []error{}
}
