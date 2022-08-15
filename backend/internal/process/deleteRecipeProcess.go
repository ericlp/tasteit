package process

import (
	"fmt"
	"github.com/ericlp/tasteit/backend/internal/db/commands"
	"github.com/ericlp/tasteit/backend/internal/db/tables"
)

func DeleteRecipe(recipe *tables.Recipe) error {
	deletedName := fmt.Sprintf("%s_%s_deleted", recipe.Name, recipe.ID)
	deletedUniqueName := fmt.Sprintf("%s_%s_deleted", recipe.UniqueName, recipe.ID)

	return commands.RecipeSetDeleted(deletedName, deletedUniqueName, recipe.ID)
}
