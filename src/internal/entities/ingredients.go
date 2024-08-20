package entities

import "github.com/CaiqueRibeiro/pizza-rute/src/pkg/entities"

type Ingredients struct {
	ID    entities.ID `json:"id"`
	Name  string      `json:"name,omitempty"`
	Stock float32     `json:"stock"`
}
