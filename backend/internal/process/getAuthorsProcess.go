package process

import (
	"github.com/ericlp/tasteit/backend/internal/db/queries"
	"github.com/ericlp/tasteit/backend/internal/models"
)

type AuthorsJson struct {
	Authors []models.Owner `json:"authors"`
}

func GetAllAuthors() (*AuthorsJson, error) {
	authors, err := queries.GetAllUsersWithRecipe()
	if err != nil {
		return nil, err
	}

	authorsJson := make([]models.Owner, 0)
	for _, author := range authors {
		authorsJson = append(authorsJson, models.Owner{
			Id:     author.ID,
			Name:   author.Name,
			IsUser: author.IsUser,
		})
	}

	return &AuthorsJson{Authors: authorsJson}, nil
}
