package queries

import (
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
)

var getUserByIdQuery = `
SELECT id, nick, cid
FROM tasteit_user 
WHERE id=$1`

func GetUser(id uuid.UUID) (*tables.User, error) {
	db := getDb()

	var user tables.User
	err := pgxscan.Get(ctx, db, &user, getUserByIdQuery, id)
	return &user, err
}

var getUserByNameQuery = `
SELECT id, nick, cid
FROM tasteit_user
WHERE cid=$1;
`

func GetUserByCID(name string) (*tables.User, error) {
	db := getDb()

	var user tables.User
	err := pgxscan.Get(ctx, db, &user, getUserByNameQuery, name)
	return &user, err
}
