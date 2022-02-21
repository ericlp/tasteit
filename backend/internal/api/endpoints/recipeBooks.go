package endpoints

import (
	"github.com/ericlp/tasteit/backend/internal/common"
	"github.com/ericlp/tasteit/backend/internal/process"
	"github.com/gin-gonic/gin"
	"log"
)

func RecipeBooks(c *gin.Context) {
	recipeBooks, err := process.GetRecipeBooks()
	if err != nil {
		log.Printf("Failed to retrieve recipe books: %v\n", err)
		c.JSON(500, common.Error(common.ResponseFailedToRetrieveRecipeBooks))
		return
	}

	c.JSON(200, common.Success(recipeBooks))
}
