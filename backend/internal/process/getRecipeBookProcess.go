package process

import (
	"github.com/ericlp/tasteit/backend/internal/common"
	"github.com/ericlp/tasteit/backend/internal/db/queries"
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/ericlp/tasteit/backend/internal/models"
	"github.com/georgysavva/scany/pgxscan"
)

func GetRecipeBook(uniqueName string) (*models.DetailedRecipeBookJson, error) {
	recipeBook, err := queries.GetRecipeBookByName(uniqueName)
	if err != nil {
		if pgxscan.NotFound(err) {
			return nil, common.ErrNoSuchRecipeBook
		}
		return nil, err
	}

	if recipeBook.Deleted {
		return nil, common.ErrNoSuchRecipeBook
	}

	recipes, err := queries.GetRecipesForRecipeBook(recipeBook.ID)
	if err != nil && !pgxscan.NotFound(err) {
		return nil, err
	}

	var imageJson *models.ImageJson = nil
	image, err := queries.GetImageForRecipeBook(recipeBook.ID)
	if err != nil {
		if !pgxscan.NotFound(err) {
			return nil, err
		}
	} else {
		imageJson = &models.ImageJson{
			Path: imageNameToPath(image.ID, image.Name),
			ID:   image.ID,
		}
	}

	owner, err := queries.GetOwner(recipeBook.OwnedBy)
	if err != nil {
		return nil, err
	}

	recipeJsons, err := RecipesToJson(recipes)
	if err != nil {
		return nil, err
	}

	return &models.DetailedRecipeBookJson{
		ID:         recipeBook.ID,
		Name:       recipeBook.Name,
		UniqueName: recipeBook.UniqueName,
		UploadedBy: models.Owner{
			Id:     owner.ID,
			Name:   owner.Name,
			IsUser: owner.IsUser,
		},
		Author:  recipeBook.Author,
		Recipes: recipeJsons,
		Image:   imageJson,
	}, nil
}

func RecipesToJson(recipes []*tables.Recipe) (
	[]models.RecipeBookRecipeJson,
	error,
) {
	recipeJsons := make([]models.RecipeBookRecipeJson, 0)
	for _, recipe := range recipes {
		owner, err := queries.GetOwner(recipe.OwnedBy)
		if err != nil {
			return nil, err
		}

		recipeJsons = append(
			recipeJsons, models.RecipeBookRecipeJson{
				Name:       recipe.Name,
				UniqueName: recipe.UniqueName,
				Author:     owner.Name,
				ID:         recipe.ID,
			},
		)
	}

	return recipeJsons, nil
}
