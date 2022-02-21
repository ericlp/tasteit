package queries

import (
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/georgysavva/scany/pgxscan"
)

var getUnitQuery = `SELECT name FROM unit WHERE name=$1`

func GetUnit(name string) (*tables.Unit, error) {
	db := getDb()

	var unit tables.Unit
	err := pgxscan.Get(ctx, db, &unit, getUnitQuery, name)
	return &unit, err
}
