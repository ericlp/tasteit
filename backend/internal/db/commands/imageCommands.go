package commands

import (
	"github.com/georgysavva/scany/pgxscan"
)

var createImageCommand = `INSERT INTO image(name) VALUES ($1) RETURNING id`

func CreateImage(name string) (uint64, error) {
	db := getDb()

	var id uint64
	err := pgxscan.Get(ctx, db, &id, createImageCommand, name)

	return id, err
}