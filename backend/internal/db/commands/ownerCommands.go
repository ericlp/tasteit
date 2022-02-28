package commands

import (
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/ericlp/tasteit/backend/internal/models"
	"github.com/georgysavva/scany/pgxscan"
)

var createDefaultOwnerCommand = `
INSERT INTO owner(name, is_user)
VALUES(			  $1, TRUE)
RETURNING id, name
`

var createOwnerCommand = `
INSERT INTO owner(name, is_user)
VALUES(			  $1, FALSE)
RETURNING id, name
`

var createUserOwnerCommand = `
INSERT INTO user_owner(	owner_id, 	tasteit_user_id)
VALUES(				$1, 		$2)
RETURNING owner_id, tasteit_user_id
`

func CreateDefaultUserOwner(user *tables.User) (*tables.Owner, error) {
	db := getDb()

	var defaultOwner tables.Owner
	err := pgxscan.Get(ctx, db, &defaultOwner, createDefaultOwnerCommand, user.Nick)
	if err != nil {
		return nil, err
	}

	var defaultUserOwner tables.UserOwner
	err = pgxscan.Get(ctx, db, &defaultUserOwner, createUserOwnerCommand, defaultOwner.ID, user.ID)

	return &defaultOwner, err
}

// Is used to create an owner for a supergroup, e.g. digit, snit etc..
func CreateOwner(gammaGroup *models.GammaGroup) (*tables.Owner, error) {
	db := getDb()

	var owner tables.Owner
	err := pgxscan.Get(ctx, db, &owner, createOwnerCommand, gammaGroup.SuperGroup.Name)
	if err != nil {
		return nil, err
	}

	return &owner, err
}

func CreateUserOwner(user *tables.User, owner *tables.Owner) (*tables.UserOwner, error) {
	db := getDb()

	var userOwner tables.UserOwner
	err := pgxscan.Get(ctx, db, &userOwner, createUserOwnerCommand, owner.ID, user.ID)

	return &userOwner, err
}
