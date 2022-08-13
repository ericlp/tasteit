package queries

import (
	"github.com/ericlp/tasteit2/backend/internal/db/tables"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
)

var getUserByIdQuery = `
SELECT id, name
FROM tasteit_user 
WHERE id=$1`

func GetUser(id uuid.UUID) (*tables.User, error) {
	db := getDb()

	var user tables.User
	err := pgxscan.Get(ctx, db, &user, getUserByIdQuery, id)
	return &user, err
}

var getUserByEmailQuery = `
SELECT id, name
FROM tasteit_user
INNER JOIN user_email ON user_email.user_id=tasteit_user.id
WHERE email=$1;
`

func GetUserByEmail(email string) (*tables.User, error) {
	db := getDb()

	var user tables.User
	err := pgxscan.Get(ctx, db, &user, getUserByEmailQuery, email)
	return &user, err
}

var GetAllUsersWithRecipeQuery = `
SELECT DISTINCT tasteit_user.id, tasteit_user.name
FROm tasteit_user
INNER JOIN recipe ON recipe.created_by=tasteit_user.id
`

func GetAllUsersWithRecipe() ([]tables.User, error) {
	db := getDb()

	var users []tables.User
	err := pgxscan.Select(ctx, db, &users, GetAllUsersWithRecipeQuery)
	return users, err
}
