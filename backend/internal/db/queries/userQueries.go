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

func GetUserByName(cid string) (*tables.User, error) {
	db := getDb()

	var user tables.User
	err := pgxscan.Get(ctx, db, &user, getUserByCIDQuery, cid)
	return &user, err
}

var GetAllUsersWithRecipeQuery = `
SELECT DISTINCT tasteit_user.id, tasteit_user.name
FROm tasteit_user
INNER JOIN recipe ON recipe.owned_by=tasteit_user.id
`

func GetAllUsersWithRecipe() ([]tables.User, error) {
	db := getDb()

	var users []tables.User
	err := pgxscan.Select(ctx, db, &users, GetAllUsersWithRecipeQuery)
	return users, err
}
