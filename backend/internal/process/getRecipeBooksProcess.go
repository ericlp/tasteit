package process

import (
	"github.com/ericlp/tasteit/backend/internal/db/queries"
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/ericlp/tasteit/backend/internal/models"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
)

type RecipeBooksJson struct {
	RecipeBooks []ShortRecipeBookJson `json:"recipeBooks"`
}

type ShortRecipeBookJson struct {
	ID         uuid.UUID    `json:"id"`
	Name       string       `json:"name"`
	UniqueName string       `json:"uniqueName"`
	Author     string       `json:"author"`
	ImageLink  string       `json:"imageLink"`
	UploadedBy models.Owner `json:"uploadedBy"`
}

func toShortRecipeBookJson(recipeBook *tables.RecipeBook, owner *tables.Owner, imageUrl string) ShortRecipeBookJson {
	return ShortRecipeBookJson{
		ID:         recipeBook.ID,
		Name:       recipeBook.Name,
		UniqueName: recipeBook.UniqueName,
		Author:     recipeBook.Author,
		ImageLink:  imageUrl,
		UploadedBy: models.Owner{
			Id:     owner.ID,
			Name:   owner.Name,
			IsUser: owner.IsUser,
		},
	}
}

func GetRecipeBooks() (*RecipeBooksJson, error) {
	recipeBooks, err := queries.GetNonDeletedRecipeBooks()
	if err != nil {
		return nil, err
	}

	if recipeBooks == nil {
		recipeBooks = make([]*tables.RecipeBook, 0)
	}

	shortRecipeBooks := make([]ShortRecipeBookJson, 0)
	for _, book := range recipeBooks {
		image, err := queries.GetImageForRecipeBook(book.ID)

		imageUrl := ""
		if err != nil {
			if pgxscan.NotFound(err) == false {
				return nil, err
			}
		} else {
			imageUrl = imageNameToPath(image.ID, image.Name)
		}

		owner, err := queries.GetOwner(book.OwnedBy)
		if err != nil {
			return nil, err
		}

		shortRecipeBooks = append(shortRecipeBooks, toShortRecipeBookJson(book, owner, imageUrl))
	}

	return &RecipeBooksJson{
		RecipeBooks: shortRecipeBooks,
	}, nil
}
