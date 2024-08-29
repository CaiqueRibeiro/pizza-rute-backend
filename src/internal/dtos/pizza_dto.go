package dtos

type PizzaItemDTO struct {
	IngredientID string  `json:"ingredient_id"`
	Quantity     float32 `json:"quantity"`
}

type CreatePizzaInput struct {
	Name        string         `json:"name"`
	SauceId     string         `json:"sauce_id"`
	Size        string         `json:"size"`
	Price       float64        `json:"price"`
	Ingredients []PizzaItemDTO `json:"ingredients"`
}
