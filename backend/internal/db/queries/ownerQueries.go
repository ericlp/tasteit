package queries

import (
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
)

var getOwnerByIdQuery = `
SELECT id, name
FROM owner		 
WHERE id=$1`

func GetOwner(id uuid.UUID) (*tables.Owner, error) {
	db := getDb()

	var owner tables.Owner
	err := pgxscan.Get(ctx, db, &owner, getOwnerByIdQuery, id)
	return &owner, err
}

var getOwnerByNameQuery = `
SELECT id, name
FROM owner
WHERE name=$1;
`

func GetOwnerByName(name string) (*tables.Owner, error) {
	db := getDb()

	var owner tables.Owner
	err := pgxscan.Get(ctx, db, &owner, getOwnerByNameQuery, name)
	return &owner, err
}

var getOwnersByUserIDQuery = `
SELECT id, name
FROM user_owner JOIN owner ON user_owner.owner_id = owner.id
WHERE tasteit_user_id=$1`

func GetOwnersByUser(id uuid.UUID) ([]*tables.Owner, error) {
	db := getDb()

	var owners []*tables.Owner
	err := pgxscan.Select(ctx, db, &owners, getOwnersByUserIDQuery, id)
	return owners, err
}
