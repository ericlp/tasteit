package process

import (
	"github.com/ericlp/tasteit/backend/internal/db/queries"
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/ericlp/tasteit/backend/internal/models"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
)

type RecipesJson struct {
	Recipes []ShortRecipeJson `json:"recipes"`
}

type ShortRecipeJson struct {
	ID                  uuid.UUID        `json:"id"`
	Name                string           `json:"name"`
	UniqueName          string           `json:"uniqueName"`
	ImageLink           string           `json:"imageLink"`
	Author              models.Owner     `json:"author"`
	Tags                []models.TagJson `json:"tags"`
	EstimatedTime       int              `json:"estimatedTime"`
	NumberOfIngredients int              `json:"numberOfIngredients"`
}

func toShortRecipeJson(
	recipe *tables.Recipe,
	owner *tables.Owner,
	imageUrl string,
	tags []models.TagJson,
	numberOfIngredients int,
) ShortRecipeJson {
	return ShortRecipeJson{
		ID:         recipe.ID,
		Name:       recipe.Name,
		UniqueName: recipe.UniqueName,
		ImageLink:  imageUrl,
		Author: models.Owner{
			Id:     owner.ID,
			Name:   owner.Name,
			IsUser: owner.IsUser,
		},
		Tags:                tags,
		EstimatedTime:       recipe.EstimatedTime,
		NumberOfIngredients: numberOfIngredients,
	}
}

func GetRecipes() (*RecipesJson, error) {
	recipes, err := queries.GetNonDeletedRecipes()
	if err != nil {
		return nil, err
	}

	if recipes == nil {
		recipes = make([]*tables.Recipe, 0)
	}

	shortRecipes := make([]ShortRecipeJson, 0)
	for _, recipe := range recipes {
		image, err := queries.GetMainImageForRecipe(recipe.ID)

		imageUrl := ""
		if err != nil {
			if pgxscan.NotFound(err) == false {
				return nil, err
			}
		} else {
			imageUrl = imageNameToPath(image.ID, image.Name)
		}

		owner, err := queries.GetOwner(recipe.OwnedBy)
		if err != nil {
			return nil, err
		}

		recipeTags, err := queries.GetTagsForRecipe(&recipe.ID)
		if err != nil {
			return nil, err
		}

		tags, err := recipeTagsToTagJsons(recipeTags)
		if err != nil {
			return nil, err
		}

		ingredientsCount, err := queries.GetNumberOfIngredientsForRecipe(recipe.ID)
		if err != nil {
			return nil, err
		}

		shortRecipes = append(
			shortRecipes,
			toShortRecipeJson(recipe, owner, imageUrl, tags, ingredientsCount),
		)
	}

	return &RecipesJson{
		Recipes: shortRecipes,
	}, nil
}

func recipeTagsToTagJsons(recipeTags []*tables.RecipeTag) (
	[]models.TagJson,
	error,
) {
	tagJson := make([]models.TagJson, 0)
	for _, recipeTag := range recipeTags {
		tag, err := queries.GetTagById(recipeTag.TagId)
		if err != nil {
			return nil, err
		}

		owner, err := queries.GetOwner(tag.OwnedBy)
		if err != nil {
			return nil, err
		}

		recipesCount, err := queries.CountRecipesWithTag(&tag.ID)
		if err != nil {
			return nil, err
		}

		tagJson = append(
			tagJson, models.TagJson{
				ID:          tag.ID,
				Name:        tag.Name,
				Description: tag.Description,
				Color: models.ColorJson{
					R: &tag.ColorRed,
					G: &tag.ColorGreen,
					B: &tag.ColorBlue,
				},
				RecipeCount: recipesCount,
				Author: models.Owner{
					Id:     owner.ID,
					Name:   owner.Name,
					IsUser: owner.IsUser,
				},
			},
		)
	}

	return tagJson, nil
}
