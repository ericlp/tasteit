package process

import (
	"github.com/ericlp/tasteit/backend/internal/db/commands"
	"github.com/ericlp/tasteit/backend/internal/db/tables"
)

func DeleteTag(tag *tables.Tag) error {
	err := commands.DeleteRecipeTagsByTagId(tag.ID)
	if err != nil {
		return err
	}

	return commands.DeleteTag(tag.ID)
}
