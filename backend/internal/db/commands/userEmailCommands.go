package commands

import (
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
)

var createUserEmailCommand = `
INSERT INTO user_email(user_id, email, provider)
VALUES(				   $1,		$2,	   $3)
RETURNING user_id, email, provider
`

func CreateUserEmail(userId uuid.UUID, email, provider string) (*tables.UserEmail, error) {
	db := getDb()

	var userEmail tables.UserEmail
	err := pgxscan.Get(ctx, db, &userEmail, createUserEmailCommand, userId, email, provider)
	return &userEmail, err
}
