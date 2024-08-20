package entities

import "github.com/CaiqueRibeiro/pizza-rute/src/pkg/entities"

const (
	Small = iota + 1
	Medium
	Big
)

type SauceItem struct {
	IngredientID entities.ID `json:"ingredient_id"`
	Quantity     float32     `json:"quantity"`
}

type Sauce struct {
	ID          entities.ID `json:"id"`
	Name        string      `json:"name"`
	Spicy       bool        `json:"is_spicy"`
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
	Size        string      `json:"size"`
	Price       float64     `json:"price"`
	Ingredients []PizzaItem `json:"ingredients"`
}
