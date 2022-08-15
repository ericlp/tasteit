package process

import (
	"github.com/ericlp/tasteit/backend/internal/common"
	"github.com/ericlp/tasteit/backend/internal/db/commands"
	"github.com/ericlp/tasteit/backend/internal/db/queries"
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/ericlp/tasteit/backend/internal/models"
	"github.com/georgysavva/scany/pgxscan"
)

func CreateNewTag(tagJson *models.NewTagJson) (*tables.Tag, error) {
	_, err := queries.GetTagByName(tagJson.Name)
	if err != nil {
		if pgxscan.NotFound(err) == false {
			return nil, err
		}
	} else {
		return nil, common.ErrNameTaken
	}

	tag, err := commands.CreateTag(tagJson.Name, tagJson.Description, *tagJson.Color.R, *tagJson.Color.G, *tagJson.Color.B, tagJson.OwnerId)
	return tag, err
}
