package process

import (
	"errors"
	common2 "github.com/viddem/vrecipes/backend/internal/common"
	"github.com/viddem/vrecipes/backend/internal/db/commands"
	dbModels "github.com/viddem/vrecipes/backend/internal/db/models"
	"github.com/viddem/vrecipes/backend/internal/db/queries"
	"github.com/viddem/vrecipes/backend/internal/models"
	"gorm.io/gorm"
	"strings"
)

func GetOrCreateIngredient(ingredientName string) (*dbModels.Ingredient, error) {
	ingredientName = strings.ToLower(strings.TrimSpace(ingredientName))
	ingredient, err := queries.GetIngredient(ingredientName)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Ingredient doesn't exist, create a new one
			ingredient = &dbModels.Ingredient{
				Name: ingredientName,
			}
			err := commands.CreateIngredient(ingredient)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return ingredient, nil
}

func GetOrCreateUnit(unitName string) (*dbModels.Unit, error) {
	unitName = strings.ToLower(unitName)
	unit, err := queries.GetUnit(unitName)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Ingredient doesn't exist, create a new one
			unit = &dbModels.Unit{
				Name: unitName,
			}
			err := commands.CreateUnit(unit)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return unit, nil
}

func CreateRecipeIngredient(ingredientName string, unitName string, amount float32, recipe *dbModels.Recipe) (*dbModels.RecipeIngredient, error) {
	ingredient, err := GetOrCreateIngredient(ingredientName)
	if err != nil {
		return nil, err
	}

	unit, err := GetOrCreateUnit(unitName)
	if err != nil {
		return nil, err
	}

	recipeIngredient := dbModels.RecipeIngredient{
		Recipe:     *recipe,
		Ingredient: *ingredient,
		Unit:       *unit,
		Amount:     amount,
	}

	err = commands.CreateRecipeIngredient(&recipeIngredient)
	return &recipeIngredient, err
}

func CreateRecipeStep(step string, number uint16, recipe *dbModels.Recipe) (*dbModels.RecipeStep, error) {
	recipeStep := dbModels.RecipeStep{
		Recipe: *recipe,
		Number: number,
		Step:   step,
	}
	err := commands.CreateRecipeStep(&recipeStep)
	if err != nil {
		return &recipeStep, err
	}

	return &recipeStep, nil
}

func CreateRecipeImage(imagePath string, recipeId uint64) (*dbModels.RecipeImage, error) {
	imageId, err := commands.CreateImage(&dbModels.Image{
		Name: imagePath,
	})

	if err != nil {
		return nil, err
	}

	return connectImageToRecipe(imageId, recipeId)
}

func connectImageToRecipe(imageId uint64, recipeId uint64) (*dbModels.RecipeImage, error) {
	recipeImage := dbModels.RecipeImage{
		ImageID: imageId,
		RecipeID: recipeId,
	}

	err := commands.CreateRecipeImage(&recipeImage)

	return &recipeImage, err
}

func CreateRecipe(name, description string, ovenTemp, estimatedTime int) (*dbModels.Recipe, error) {
	uniqueName, err := generateUniqueName(name)
	if err != nil {
		return &dbModels.Recipe{}, err
	}

	recipe := dbModels.Recipe{
		Name:          name,
		UniqueName:    uniqueName,
		Description:   description,
		OvenTemp:      ovenTemp,
		EstimatedTime: estimatedTime,
	}

	_, err = commands.CreateRecipe(&recipe)

	return &recipe, err
}

func CreateNewRecipe(recipeJson *models.NewRecipeJson) (string, error) {
	recipe, err := CreateRecipe(recipeJson.Name, recipeJson.Description, recipeJson.OvenTemperature, recipeJson.CookingTime)
	if err != nil {
		return "", err
	}

	for _, ingredient := range recipeJson.Ingredients {
		_, err := CreateRecipeIngredient(ingredient.Name, ingredient.Unit, ingredient.Amount, recipe)
		if err != nil {
			return "", err
		}
	}

	for _, step := range recipeJson.Steps {
		_, err := CreateRecipeStep(step.Step, step.Number, recipe)
		if err != nil {
			return "", err
		}
	}

	for _, image := range recipeJson.Images {
		_, err := connectImageToRecipe(image.ID, recipe.ID)
		if err != nil {
			return "", err
		}
	}

	return recipe.UniqueName, nil
}

func generateUniqueName(name string) (string, error) {
	uniqueName := strings.ReplaceAll(strings.ToLower(name), " ", "_")
	_, err := queries.GetRecipeByName(uniqueName)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return uniqueName, nil
		}
		return "", err
	}
	return uniqueName, common2.ErrRowAlreadyExists
}