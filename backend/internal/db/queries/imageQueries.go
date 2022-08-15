package queries

import (
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
)

var getImageByIdQuery = `SELECT id, name FROM image WHERE id=$1`

func GetImageById(id uuid.UUID) (*tables.Image, error) {
	db := getDb()

	var image tables.Image
	err := pgxscan.Get(ctx, db, &image, getImageByIdQuery, id)
	return &image, err
}
