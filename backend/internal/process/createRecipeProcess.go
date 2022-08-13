package process

import (
	common2 "github.com/ericlp/tasteit2/backend/internal/common"
	"github.com/ericlp/tasteit2/backend/internal/db/commands"
	"github.com/ericlp/tasteit2/backend/internal/db/queries"
	"github.com/ericlp/tasteit2/backend/internal/db/tables"
	"github.com/ericlp/tasteit2/backend/internal/models"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"strings"
)

func CreateRecipe(
	name string, userId uuid.UUID,
) (*tables.Recipe, error) {
	uniqueName, err := generateUniqueName(name)
	if err != nil {
		return nil, err
	}
	recipe, err := commands.CreateRecipe(
		name,
		uniqueName,
		"",
		0,
		0,
		userId,
	)
	return recipe, err
}

func CreateNewRecipe(
	recipeJson *models.NewRecipeJson,
	user *tables.User,
) (string, error) {
	recipe, err := CreateRecipe(recipeJson.Name, user.ID)
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
