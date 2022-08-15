package models

import (
	"github.com/google/uuid"
)

type DetailedRecipeBookJson struct {
	ID         uuid.UUID              `json:"id"`
	Name       string                 `json:"name"`
	UniqueName string                 `json:"uniqueName"`
	UploadedBy Owner                  `json:"uploadedBy"`
	Author     string                 `json:"author"`
	Recipes    []RecipeBookRecipeJson `json:"recipes"`
	Image      *ImageJson             `json:"image"`
}

type RecipeBookRecipeJson struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	UniqueName string    `json:"uniqueName"`
	Author     string    `json:"author"`
}
