package commands

import (
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
)

var createRecipeCommand = `
INSERT INTO recipe(name, unique_name, description, oven_temp, estimated_time, deleted, portions, owned_by)
		   VALUES ($1,   $2,          $3,          $4,        $5,             $6,	   $7,       $8)
RETURNING id, name, unique_name, description, oven_temp, estimated_time, deleted, portions, owned_by`

func CreateRecipe(name, uniqueName, description string, ovenTemp, estimatedTime, portions int, OwnedBy uuid.UUID) (*tables.Recipe, error) {
	db := getDb()

	var recipe tables.Recipe
	err := pgxscan.Get(ctx, db, &recipe, createRecipeCommand, name, uniqueName, description, ovenTemp, estimatedTime, portions, false, OwnedBy)
	return &recipe, err
}

var updateRecipeCommand = `
UPDATE recipe 
SET name=$1,
	unique_name=$2,
	description=$3,
	oven_temp=$4,
	estimated_time=$5,
	portions=$6
WHERE id=$7
`

func UpdateRecipe(name, uniqueName, description string, ovenTemp, estimatedTime, portions int, recipeId uuid.UUID) error {
	db := getDb()

	_, err := db.Exec(ctx, updateRecipeCommand, name, uniqueName, description,
		ovenTemp, estimatedTime, portions, recipeId)
	return err
}

var recipeSetDeletedCommand = `
UPDATE recipe
SET deleted=true,
	name=$1,
	unique_name=$2
WHERE id=$3
`

func RecipeSetDeleted(name, uniqueName string, id uuid.UUID) error {
	db := getDb()

	_, err := db.Exec(ctx, recipeSetDeletedCommand, name, uniqueName, id)
	return err
}
