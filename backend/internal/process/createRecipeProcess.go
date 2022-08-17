package process

import (
	common2 "github.com/ericlp/tasteit/backend/internal/common"
	"github.com/ericlp/tasteit/backend/internal/db/commands"
	"github.com/ericlp/tasteit/backend/internal/db/queries"
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/ericlp/tasteit/backend/internal/models"
	"github.com/georgysavva/scany/pgxscan"
	"strings"
)

func CreateRecipe(
	newRecipe *models.NewRecipeJson,
) (*tables.Recipe, error) {
	uniqueName, err := generateUniqueName(newRecipe.Name)
	if err != nil {
		return nil, err
	}
	recipe, err := commands.CreateRecipe(
		newRecipe.Name,
		uniqueName,
		"",
		0,
		0,
		0,
		newRecipe.OwnerId,
	)
	return recipe, err
}

func CreateNewRecipe(
	recipeJson *models.NewRecipeJson,
) (string, error) {
	recipe, err := CreateRecipe(recipeJson)
	if err != nil {
		return "", err
	}

	return recipe.UniqueName, nil
}

func generateUniqueName(name string) (string, error) {
	lowerCase := strings.ToLower(name)
	uniqueName := strings.ReplaceAll(lowerCase, " ", "_")

	_, err := queries.GetRecipeByName(uniqueName)
	if err != nil {
		if pgxscan.NotFound(err) {
			return uniqueName, nil
		}
		return "", err
	}
	return uniqueName, common2.ErrNameTaken
}
