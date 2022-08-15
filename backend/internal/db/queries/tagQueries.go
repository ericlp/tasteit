package queries

import (
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
)

var getTagByNameQuery = `
SELECT id, name, description, color_red, color_green, color_blue, owned_by
FROM tag
WHERE name=$1
`

func GetTagByName(name string) (*tables.Tag, error) {
	db := getDb()

	var tag tables.Tag
	err := pgxscan.Get(ctx, db, &tag, getTagByNameQuery, name)
	return &tag, err
}

var getTagByIdQuery = `
SELECT id, name, description, color_red, color_green, color_blue, owned_by
FROM tag
WHERE id=$1
`

func GetTagById(id uuid.UUID) (*tables.Tag, error) {
	db := getDb()

	var tag tables.Tag
	err := pgxscan.Get(ctx, db, &tag, getTagByIdQuery, id)
	return &tag, err
}

var getAllTagsQuery = `
SELECT id, name, description, color_red, color_green, color_blue, owned_by
FROM tag
`

func GetAllTags() ([]*tables.Tag, error) {
	db := getDb()

	var tags []*tables.Tag
	err := pgxscan.Select(ctx, db, &tags, getAllTagsQuery)

	return tags, err
}
