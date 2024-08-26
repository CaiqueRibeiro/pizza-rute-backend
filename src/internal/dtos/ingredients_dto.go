package dtos

type CreateIngredientInput struct {
	Name  string  `json:"name"`
	Stock float32 `json:"stock"`
}

type UpdateIngredientInput struct {
	Name  string   `json:"name,omitempty"`
	Stock *float32 `json:"stock,omitempty"`
}
