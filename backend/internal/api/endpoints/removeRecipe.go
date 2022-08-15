package endpoints

import (
	"github.com/ericlp/tasteit/backend/internal/common"
	"github.com/ericlp/tasteit/backend/internal/process"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RemoveRecipe(c *gin.Context) {
	recipe, err := validateRecipeId(c)
	if err != nil {
		log.Printf("Failed to validate recipe id: %v\n", err)
		return
	}

	err = validateOwnerAuthorized(c, recipe.OwnedBy)
	if err != nil {
		log.Printf("Failed to authorize user: %v\n", err)
		c.JSON(http.StatusForbidden, common.Error(common.ResponseIncorrectUser))
		return
	}

	err = process.DeleteRecipe(recipe)
	if err != nil {
		log.Printf("Failed to delete recipe: %v\n", err)
		c.JSON(
			http.StatusInternalServerError,
			common.Error(common.ResponseFailedToDeleteRecipe),
		)
		return
	}

	c.JSON(http.StatusOK, common.Success("Recipe deleted"))
}
