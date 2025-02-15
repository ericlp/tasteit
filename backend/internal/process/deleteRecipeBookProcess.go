package process

import (
	"fmt"
	"github.com/ericlp/tasteit/backend/internal/db/commands"
	"github.com/ericlp/tasteit/backend/internal/db/tables"
)

func DeleteRecipeBook(recipeBook *tables.RecipeBook) error {
	deletedName := fmt.Sprintf("%s_%s_deleted", recipeBook.Name, recipeBook.ID)
	deletedUniqueName := fmt.Sprintf("%s_%s_deleted", recipeBook.UniqueName, recipeBook.ID)

	err := commands.RecipeBookSetDeleted(deletedName, deletedUniqueName, recipeBook.ID)
	return err
}
