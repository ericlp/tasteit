package models

import (
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/google/uuid"
)

type DetailedRecipeJson struct {
	ID              uuid.UUID              `json:"id"`
	UniqueName      string                 `json:"uniqueName"`
	Name            string                 `json:"name"`
	Description     string                 `json:"description"`
	OvenTemperature int                    `json:"ovenTemperature"`
	EstimatedTime   int                    `json:"estimatedTime"`
	Steps           []RecipeStepJson       `json:"steps"`
	Ingredients     []RecipeIngredientJson `json:"ingredients"`
	Images          []ImageJson            `json:"images"`
	Author          tables.Owner           `json:"author"`
	Tags            []TagJson              `json:"tags"`
}

type RecipeStepJson struct {
	Number      uint16 `json:"number"`
	Description string `json:"description"`
	IsHeading   bool   `json:"isHeading"`
}

type RecipeIngredientJson struct {
	Number    int     `json:"number"`
	Name      string  `json:"name"`
	Unit      string  `json:"unit"`
	Amount    float32 `json:"amount"`
	IsHeading bool    `json:"isHeading"`
}

type ImageJson struct {
	Path string    `json:"url"`
	ID   uuid.UUID `json:"id"`
}
