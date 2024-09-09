package entities

import (
	"testing"

	"github.com/CaiqueRibeiro/pizza-rute/src/internal/dtos"
)

type Assertions struct {
	desc    string
	igName  string
	igStock float32
	wantErr bool
	errMsg  string
}

func TestIngredientsEntity(t *testing.T) {
	tests := []Assertions{
		{
			desc:    "should create a new ingredient",
			igName:  "Tomatoe",
			igStock: 500,
			wantErr: false,
		},
		{
			desc:    "should return error when name is empty",
			igName:  "",
			igStock: 500,
			wantErr: true,
			errMsg:  "name is required",
		},
		{
			desc:    "should return error when name if stock is 0",
			igName:  "Tomatoe",
			igStock: 0,
			wantErr: true,
			errMsg:  "stock is required",
		},
		{
			desc:    "should return error when name if stock is negative",
			igName:  "Tomatoe",
			igStock: -1,
			wantErr: true,
			errMsg:  "stock is required",
		},
	}

	for _, scenario := range tests {
		t.Run(scenario.desc, func(t *testing.T) {
			dto := dtos.CreateIngredientInput{
				Name:  scenario.igName,
				Stock: scenario.igStock,
			}
			ingredient, errs := NewIngredient(dto)

			if scenario.wantErr {
				if len(errs) == 0 {
					t.Errorf("expected errors, got nil")
				}
				for _, err := range errs {
					if err.Error() != scenario.errMsg {
						t.Errorf("expected %s, got %s", scenario.errMsg, err.Error())
					}
				}
			} else {
				if len(errs) > 0 {
					t.Errorf("expected no errors, got %v", errs)
				}
				if ingredient.Name != scenario.igName {
					t.Errorf("expected %s, got %s", scenario.igName, ingredient.Name)
				}
				if ingredient.Stock != scenario.igStock {
					t.Errorf("expected %f, got %f", scenario.igStock, ingredient.Stock)
				}
			}
		})
	}
}
