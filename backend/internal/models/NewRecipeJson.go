package models

import "github.com/google/uuid"

type NewRecipeJson struct {
	Name    string    `json:"name" binding:"required"`
	OwnerId uuid.UUID `json:"ownerId" binding:"required"`
}
