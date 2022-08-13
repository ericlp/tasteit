package models

import "github.com/google/uuid"

type Owner struct {
	Id     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	IsUser bool      `json:"isUser"`
}
