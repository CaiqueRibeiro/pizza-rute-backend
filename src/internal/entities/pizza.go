package entities

import (
	"github.com/CaiqueRibeiro/pizza-rute/src/internal/dtos"
	"github.com/CaiqueRibeiro/pizza-rute/src/pkg/entities"
)

type Size string

const (
	Small  Size = "SMALL"
	Medium Size = "MEDIUM"
	Big    Size = "BIG"
)

type SauceItem struct {
	IngredientID entities.ID `json:"ingredient_id"`
	Quantity     float32     `json:"quantity"`
}

type Sauce struct {
	ID          entities.ID `json:"id"`
	Name        string      `json:"name,omitempty"`
	Spicy       bool        `json:"is_spicy,omitempty"`
	Ingredients []SauceItem `json:"ingredients"`
}

type PizzaItem struct {
	IngredientID entities.ID `json:"ingredient_id"`
	Quantity     float32     `json:"quantity"`
}

type Pizza struct {
	ID          entities.ID `json:"id"`
	Name        string      `json:"name"`
	Sauce       Sauce       `json:"sauce"`
	Size        Size        `json:"size"`
	Price       float64     `json:"price"`
	Ingredients []PizzaItem `json:"ingredients"`
}

func validateSize(s Size) error {
	switch s {
	case Small, Medium, Big:
		return nil
	default:
		return ErrSizeIsRequired
	}
}

func NewPizza(dto dtos.CreatePizzaInput) (*Pizza, []error) {
	var errs []error

	if dto.Name == "" {
		errs = append(errs, ErrNameIsRequired)
	}

	if dto.Price <= 0 {
		errs = append(errs, ErrPriceIsInvalid)
	}

	if dto.Size == "" {
		errs = append(errs, ErrSizeIsRequired)
	}

	if err := validateSize(Size(dto.Size)); err != nil {
		errs = append(errs, ErrSizeIsRequired)
	}

	if len(dto.Ingredients) == 0 {
		errs = append(errs, ErrIngredientsAreRequired)
	}

	if dto.SauceId == "" {
		errs = append(errs, ErrSauceIsRequired)
	}

	var ingredients []PizzaItem

	for _, v := range dto.Ingredients {
		id, err := entities.ParseID(v.IngredientID)
		if err != nil {
			errs = append(errs, ErrIngredientIdIsInvalid)
		} else {
			ingredient := PizzaItem{
				IngredientID: id,
				Quantity:     v.Quantity,
			}
			ingredients = append(ingredients, ingredient)
		}
	}

	sauceId, err := entities.ParseID(dto.SauceId)

	if err != nil {
		errs = append(errs, ErrSauceIdIsInvalid)
	}

	if len(errs) > 0 {
		return nil, errs
	}

	return &Pizza{
		ID:          entities.NewID(),
		Name:        dto.Name,
		Price:       dto.Price,
		Size:        Size(dto.Size),
		Ingredients: ingredients,
		Sauce: Sauce{
			ID: sauceId,
		},
	}, []error{}
}
