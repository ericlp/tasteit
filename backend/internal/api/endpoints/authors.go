package endpoints

import (
	"github.com/ericlp/tasteit/backend/internal/common"
	"github.com/ericlp/tasteit/backend/internal/process"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Authors(c *gin.Context) {
	recipes, err := process.GetAllAuthors()
	if err != nil {
		log.Printf("Error: Failed to retrieve recipes due to %s\n", err)
		c.JSON(
			http.StatusInternalServerError,
			common.Error(common.ResponseFailedToRetrieveRecipes),
		)
		return
	}

	c.JSON(http.StatusOK, common.Success(recipes))
}
