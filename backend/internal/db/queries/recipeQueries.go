package queries

import (
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
)

var getRecipeByNameQuery = `SELECT id, name, unique_name, description, oven_temp, estimated_time, deleted, owned_by, portions, portions_suffix
FROM recipe
WHERE unique_name=$1`

func GetRecipeByName(uniqueName string) (*tables.Recipe, error) {
	db := getDb()

	var recipe tables.Recipe
	err := pgxscan.Get(ctx, db, &recipe, getRecipeByNameQuery, uniqueName)
	return &recipe, err
}

var getRecipeByIdQuery = `SELECT id, name, unique_name, description, oven_temp, estimated_time, deleted, owned_by, portions, portions_suffix
FROM recipe
WHERE id=$1`

func GetRecipeById(id uuid.UUID) (*tables.Recipe, error) {
	db := getDb()

	var recipe tables.Recipe
	err := pgxscan.Get(ctx, db, &recipe, getRecipeByIdQuery, id)
	return &recipe, err
}

var getNonDeletedRecipesQuery = `SELECT id, name, unique_name, description, oven_temp, estimated_time, deleted, owned_by, portions, portions_suffix
FROM recipe
WHERE deleted=false`

func GetNonDeletedRecipes() ([]*tables.Recipe, error) {
	db := getDb()

	var recipes []*tables.Recipe
	err := pgxscan.Select(ctx, db, &recipes, getNonDeletedRecipesQuery)

	return recipes, err
}

var getRecipesForRecipeBookQuery = `
SELECT recipe.id, name, unique_name, description, oven_temp, estimated_time, deleted, owned_by, portions, portions_suffix
FROM recipe_book_recipe 
	JOIN recipe ON recipe_book_recipe.recipe_id = recipe.id
WHERE recipe_book_recipe.recipe_book_id = $1
AND recipe.deleted = false;
`

func GetRecipesForRecipeBook(recipeBookId uuid.UUID) ([]*tables.Recipe, error) {
	db := getDb()

	var recipes []*tables.Recipe
	err := pgxscan.Select(ctx, db, &recipes, getRecipesForRecipeBookQuery, recipeBookId)

	return recipes, err
}
