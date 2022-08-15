package commands

import (
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/ericlp/tasteit/backend/internal/models"
	"github.com/georgysavva/scany/pgxscan"
)

var createUserCommand = `
INSERT INTO tasteit_user(nick, cid)
VALUES(					  $1,  $2)
RETURNING id, nick, cid
`

func CreateUser(gammaUser *models.GammaMe) (*tables.User, error) {
	db := getDb()

	var user tables.User
	err := pgxscan.Get(ctx, db, &user, createUserCommand, gammaUser.Nick, gammaUser.Cid)

	if err != nil {
		return nil, err
	}

	_, err = CreateDefaultUserOwner(&user)
	return &user, err
}
