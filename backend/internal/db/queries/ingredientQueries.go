package queries

import (
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/georgysavva/scany/pgxscan"
)

var getIngredientByNameQuery = `SELECT name FROM ingredient WHERE name=$1`

func GetIngredient(name string) (*tables.Ingredient, error) {
	db := getDb()

	var ingredient tables.Ingredient
	err := pgxscan.Get(ctx, db, &ingredient, getIngredientByNameQuery, name)
	return &ingredient, err
}
