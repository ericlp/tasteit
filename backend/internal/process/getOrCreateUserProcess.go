package process

import (
	"github.com/ericlp/tasteit2/backend/internal/db/commands"
	"github.com/ericlp/tasteit2/backend/internal/db/queries"
	"github.com/ericlp/tasteit2/backend/internal/db/tables"
	"github.com/georgysavva/scany/pgxscan"
)

func GetOrCreateUser(name, email, provider string) (*tables.User, error) {
	user, err := queries.GetUserByEmail(email)
	if err == nil {
		return user, nil
	}

	if pgxscan.NotFound(err) == false {
		return nil, err
	}

	user, err = commands.CreateUser(name)
	if err != nil {
		return nil, err
	}

	_, err = commands.CreateUserEmail(user.ID, email, provider)
	return user, err
}
