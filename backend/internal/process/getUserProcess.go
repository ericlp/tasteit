package process

import (
	"github.com/ericlp/tasteit/backend/internal/db/queries"
	"github.com/ericlp/tasteit/backend/internal/models"
	"github.com/google/uuid"
)

func GetUserJson(id uuid.UUID) (*models.User, error) {
	user, err := queries.GetUser(id)
	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:   user.ID,
		Nick: user.Nick,
		CID:  user.CID,
	}, nil
}
