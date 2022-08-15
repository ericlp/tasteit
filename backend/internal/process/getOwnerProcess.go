package process

import (
	"github.com/ericlp/tasteit/backend/internal/db/queries"
	"github.com/ericlp/tasteit/backend/internal/models"
	"github.com/google/uuid"
)

func GetOwnersJson(id uuid.UUID) ([]models.Owner, error) {
	owners, err := queries.GetOwnersByUser(id)
	if err != nil {
		return nil, err
	}

	ownersJson := make([]models.Owner, 0)
	for _, owner := range owners {
		ownersJson = append(ownersJson, models.Owner{Name: owner.Name, Id: owner.ID, IsUser: owner.IsUser})
	}

	return ownersJson, nil
}
