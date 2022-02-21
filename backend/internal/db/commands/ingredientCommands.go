package commands

import (
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/georgysavva/scany/pgxscan"
)

var createIngredientCommand = `
INSERT INTO ingredient 
VALUES ($1)
RETURNING name
`

func CreateIngredient(name string) (*tables.Ingredient, error) {
	db := getDb()

	var ingredient tables.Ingredient
	err := pgxscan.Get(ctx, db, &ingredient, createIngredientCommand, name)
	return &ingredient, err
}
