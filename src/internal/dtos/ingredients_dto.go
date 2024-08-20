package dtos

type CreateIngredientInput struct {
	Name  string  `json:"name"`
	Stock float32 `json:"stock"`
}
