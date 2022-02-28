package commands

import (
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
)

var createTagCommand = `
INSERT INTO tag(name, description, color_red, color_green, color_blue, owned_by)
		 VALUES($1,   $2,		   $3, 		  $4, 		   $5, 		   $6)
RETURNING id, description, color_red, color_green, color_blue, owned_by
`

func CreateTag(
	name, description string,
	colorRed, colorGreen, colorBlue uint8,
	OwnedBy uuid.UUID,
) (*tables.Tag, error) {
	db := getDb()

	var tag tables.Tag
	err := pgxscan.Get(
		ctx,
		db,
		&tag,
		createTagCommand,
		name,
		description,
		colorRed,
		colorGreen,
		colorBlue,
		OwnedBy,
	)
	return &tag, err
}

var deleteTagCommand = `
DELETE
FROM tag
WHERE id=$1
`

func DeleteTag(id uuid.UUID) error {
	db := getDb()

	_, err := db.Exec(ctx, deleteTagCommand, id)
	return err
}

var updateTagCommand = `
UPDATE tag
SET name=$1,
	description=$2,
	color_red=$3,
	color_green=$4,
	color_blue=$5
WHERE id = $6
`

func UpdateTag(
	name, description string,
	red, green, blue uint8,
	id uuid.UUID,
) error {
	db := getDb()

	_, err := db.Exec(
		ctx,
		updateTagCommand,
		name,
		description,
		red,
		green,
		blue,
		id,
	)

	return err
}
