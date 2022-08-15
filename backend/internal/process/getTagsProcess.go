package process

import (
	"github.com/ericlp/tasteit/backend/internal/db/queries"
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/ericlp/tasteit/backend/internal/models"
)

type TagsJson struct {
	Tags []models.TagJson `json:"tags"`
}

func GetTags() (*TagsJson, error) {
	tags, err := queries.GetAllTags()
	if err != nil {
		return nil, err
	}

	if tags == nil {
		tags = make([]*tables.Tag, 0)
	}

	tagJsons := make([]models.TagJson, 0)
	for _, tag := range tags {
		owner, err := queries.GetOwner(tag.OwnedBy)
		if err != nil {
			return nil, err
		}

		recipeCount, err := queries.CountRecipesWithTag(&tag.ID)
		if err != nil {
			return nil, err
		}

		tagJsons = append(
			tagJsons, models.TagJson{
				ID:          tag.ID,
				Name:        tag.Name,
				Description: tag.Description,
				Color: models.ColorJson{
					R: &tag.ColorRed,
					G: &tag.ColorGreen,
					B: &tag.ColorBlue,
				},
				RecipeCount: recipeCount,
				Author: models.Owner{
					Id:     owner.ID,
					Name:   owner.Name,
					IsUser: owner.IsUser,
				},
			},
		)
	}

	return &TagsJson{
		Tags: tagJsons,
	}, nil
}
