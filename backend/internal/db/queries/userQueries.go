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

var getUserByCIDQuery = `
SELECT id, nick, cid
FROM tasteit_user
WHERE cid=$1;
`

func GetUserByCID(cid string) (*tables.User, error) {
	db := getDb()

	var user tables.User
	err := pgxscan.Get(ctx, db, &user, getUserByCIDQuery, cid)
	return &user, err
}

var GetAllOwnersWithRecipeQuery = `
SELECT DISTINCT owner.id, owner.name
FROM owner
INNER JOIN recipe ON recipe.owned_by=owner.id
`

func GetAllUsersWithRecipe() ([]tables.Owner, error) {
	db := getDb()

	var owners []tables.Owner
	err := pgxscan.Select(ctx, db, &owners, GetAllOwnersWithRecipeQuery)
	return owners, err
}
