package process

import (
	"github.com/ericlp/tasteit/backend/internal/common"
	"github.com/ericlp/tasteit/backend/internal/db/commands"
	"github.com/ericlp/tasteit/backend/internal/db/queries"
	"github.com/ericlp/tasteit/backend/internal/models"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"strings"
)

func CreateNewRecipeBook(
	recipeBookJson *models.NewRecipeBookJson,
) (string, error) {
	uniqueName, err := generateUniqueBookName(recipeBookJson.Name)
	if err != nil {
		return "", err
	}

	recipeBook, err := commands.CreateRecipeBook(
		recipeBookJson.Name,
		uniqueName,
		recipeBookJson.OwnerId,
	)
	if err != nil {
		return "", err
	}

	return recipeBook.UniqueName, nil
}

func generateUniqueBookName(name string) (string, error) {
	lowerCase := strings.ToLower(name)
	uniqueName := strings.ReplaceAll(lowerCase, " ", "_")

	_, err := queries.GetRecipeBookByName(uniqueName)
	if err != nil {
		if pgxscan.NotFound(err) {
			return uniqueName, nil
		}
		return "", err
	}

	return uniqueName, common.ErrNameTaken
}

func createRecipeBookRecipes(
	recipeBookId uuid.UUID,
	recipes []uuid.UUID,
) error {
	for _, recipe := range recipes {
		_, err := commands.CreateRecipeBookRecipe(recipeBookId, recipe)
		if err != nil {
			return err
		}
	}
	return nil
}

func connectImagesToRecipeBook(
	recipeBookId uuid.UUID,
	imageIds []uuid.UUID,
) error {
	for _, imageId := range imageIds {
		_, err := commands.CreateRecipeBookImage(recipeBookId, imageId)
		if err != nil {
			return err
		}
	}

	return nil
}
