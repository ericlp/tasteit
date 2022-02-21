package endpoints

import (
	"github.com/ericlp/tasteit/backend/internal/common"
	"github.com/ericlp/tasteit/backend/internal/process"
	"github.com/gin-gonic/gin"
	"log"
)

func Recipes(c *gin.Context) {
	recipes, err := process.GetRecipes()
	if err != nil {
		log.Printf("Error: Failed to retrieve recipes due to %s\n", err)
		c.JSON(500, common.Error(common.ResponseFailedToRetrieveRecipes))
		return
	}

	c.JSON(200, common.Success(recipes))
}
