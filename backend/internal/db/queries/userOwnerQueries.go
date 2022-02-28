package queries

import (
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/georgysavva/scany/pgxscan"
)

var getUserOwnerByUserIDQuery = `
SELECT tasteit_user_id, owner_id 
FROM user_owner
WHERE tasteit_user_id=$1 AND owner_id=$2`

func GetOwnerByUserAndOwner(user *tables.User, owner *tables.Owner) (*tables.UserOwner, error) {
	db := getDb()

	var userOwner tables.UserOwner
	err := pgxscan.Get(ctx, db, &userOwner, getUserOwnerByUserIDQuery, user.ID, owner.ID)
	return &userOwner, err
}
