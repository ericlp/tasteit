package endpoints

import (
	"errors"
	"github.com/ericlp/tasteit/backend/internal/common"
	"github.com/ericlp/tasteit/backend/internal/process"
	"github.com/gin-gonic/gin"
	"log"
)

func RecipeBook(c *gin.Context) {
	uniqueName := c.Param("uniqueName")
	detailedRecipeBook, err := process.GetRecipeBook(uniqueName)
	if err != nil {
		if errors.Is(err, common.ErrNoSuchRecipeBook) {
			c.JSON(404, common.Error(common.ResponseRecipeBookNotFound))
		}
		log.Printf("Error: Failed to retrieve recipebook %s, due to error: %v\n", uniqueName, err)
		c.JSON(500, common.Error(common.ResponseFailedToRetrieveRecipeBook))
		return
	}

	c.JSON(200, common.Success(detailedRecipeBook))
}
