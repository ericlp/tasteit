package tables

import (
	"github.com/ericlp/tasteit/backend/internal/common"
	"github.com/google/uuid"
)

// An ingredient in a recipe
type RecipeIngredient struct {
	ID             uuid.UUID
	RecipeID       uuid.UUID
	IngredientName string
	UnitName       string
	Amount         float32
	Number         int
}

func (_ RecipeIngredient) StructName() string {
	return "Recipe Ingredient"
}

func (recIng *RecipeIngredient) Equals(other *RecipeIngredient) bool {
	return recIng.ID == other.ID &&
		recIng.RecipeID == other.RecipeID &&
		recIng.IngredientName == other.IngredientName &&
		recIng.UnitName == other.UnitName &&
		common.FloatsAreSame(recIng.Amount, other.Amount)
}
