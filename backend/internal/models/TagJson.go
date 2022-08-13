package models

import (
	"github.com/ericlp/tasteit2/backend/internal/db/tables"
	"github.com/google/uuid"
)

type TagJson struct {
	ID          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Color       ColorJson   `json:"color"`
	RecipeCount uint64      `json:"recipeCount"`
	Author      tables.User `json:"author"`
}
