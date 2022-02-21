package commands

import (
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/georgysavva/scany/pgxscan"
)

var createUserCommand = `
INSERT INTO tasteit_user(name)
VALUES(					  $1)
RETURNING id, name
`

func CreateUser(name string) (*tables.User, error) {
	db := getDb()

	var user tables.User
	err := pgxscan.Get(ctx, db, &user, createUserCommand, name)
	return &user, err
}
