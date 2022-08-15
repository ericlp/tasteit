package process

import (
	"github.com/ericlp/tasteit/backend/internal/common"
	"github.com/ericlp/tasteit/backend/internal/db/commands"
	"github.com/ericlp/tasteit/backend/internal/db/queries"
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/ericlp/tasteit/backend/internal/models"
	"github.com/georgysavva/scany/pgxscan"
)

func EditTag(oldTag *tables.Tag, newTag *models.NewTagJson) error {
	changed := false

	if oldTag.Name != newTag.Name {
		_, err := queries.GetTagByName(newTag.Name)
		if err != nil {
			if !pgxscan.NotFound(err) {
				return err
			}
		} else {
			return common.ErrNameTaken
		}
		changed = true
	}

	changed = changed ||
		newTag.Description != oldTag.Description ||
		*newTag.Color.R != oldTag.ColorRed ||
		*newTag.Color.G != oldTag.ColorGreen ||
		*newTag.Color.B != oldTag.ColorBlue

	if changed {
		return commands.UpdateTag(
			newTag.Name,
			newTag.Description,
			*newTag.Color.R,
			*newTag.Color.G,
			*newTag.Color.B,
			oldTag.ID,
		)
	}

	return nil
}
