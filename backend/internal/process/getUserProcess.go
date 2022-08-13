package process

import (
	"github.com/ericlp/tasteit2/backend/internal/db/queries"
	"github.com/ericlp/tasteit2/backend/internal/models"
	"github.com/google/uuid"
)

func GetUserJson(id uuid.UUID) (*models.User, error) {
	user, err := queries.GetUser(id)
	if err != nil {
		return nil, err
	}

	emails, err := queries.GetEmailsForUser(user.ID)
	if err != nil {
		return nil, err
	}

	var emailJsons []models.UserEmail
	for _, email := range emails {
		emailJsons = append(emailJsons, models.UserEmail{
			Email:    email.Email,
			Provider: email.Provider,
		})
	}

	return &models.User{
		Id:     user.ID,
		Name:   user.Name,
		Emails: emailJsons,
	}, nil
}
