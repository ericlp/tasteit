package process

import (
	"github.com/ericlp/tasteit/backend/internal/db/commands"
	"github.com/ericlp/tasteit/backend/internal/db/queries"
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/ericlp/tasteit/backend/internal/models"
	"github.com/georgysavva/scany/pgxscan"
)

func getOrCreateUser(gammaUser *models.GammaMe) (*tables.User, error) {
	user, err := queries.GetUserByCID(gammaUser.Cid)
	if err == nil {
		return user, err
	}

	if pgxscan.NotFound(err) == false {
		return nil, err
	}

	user, err = commands.CreateUser(gammaUser)
	if err != nil {
		return nil, err
	}

	return user, err
}

func setupDefaultUser(user *tables.User) (*tables.Owner, error) {
	_, err := queries.GetOwnerByName(user.Nick)
	if err == nil {
		return nil, err
	}

	if pgxscan.NotFound(err) == false {
		return nil, err
	}

	owner, err := commands.CreateDefaultUserOwner(user)
	if err != nil {
		return nil, err
	}

	return owner, err
}

func GetOrSetupUser(gammaUser *models.GammaMe) (*tables.User, error) {
	user, err := getOrCreateUser(gammaUser)
	if err != nil {
		return nil, err
	}

	_, err = setupDefaultUser(user)
	if err != nil {
		return nil, err
	}

	for _, group := range gammaUser.Groups {
		if group.Active {
			// Add group if it doesn't exist
			owner, err := queries.GetOwnerByName(group.Name)
			if err != nil {
				if pgxscan.NotFound(err) {
					owner, err = commands.CreateOwner(&group)
					if err != nil {
						return nil, err
					}
				} else {
					return nil, err
				}
			}

			// Add user to group if not part of group
			_, err = queries.GetOwnerByUserAndOwner(user, owner)
			if err != nil {
				if pgxscan.NotFound(err) {
					_, err = commands.CreateUserOwner(user, owner)
					if err != nil {
						return nil, err
					}
				} else {
					return nil, err
				}
			}
		}
	}

	return user, err
}
