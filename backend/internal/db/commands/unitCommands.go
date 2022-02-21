package commands

import (
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/georgysavva/scany/pgxscan"
)

var createUnitCommand = `
INSERT INTO unit 
VALUES($1)
returning name
`

func CreateUnit(name string) (*tables.Unit, error) {
	db := getDb()

	var unit tables.Unit
	err := pgxscan.Get(ctx, db, &unit, createUnitCommand, name)
	return &unit, err
}
