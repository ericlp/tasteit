package queries

import (
	"github.com/ericlp/tasteit2/backend/internal/db/tables"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
)

var getEmailsForUserQuery = `
SELECT user_id, email, provider
FROM user_email
WHERE user_id=$1
`

func GetEmailsForUser(id uuid.UUID) ([]*tables.UserEmail, error) {
	db := getDb()

	var emails []*tables.UserEmail
	err := pgxscan.Select(ctx, db, &emails, getEmailsForUserQuery, id)
	return emails, err
}
