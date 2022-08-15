package endpoints

import (
	"github.com/ericlp/tasteit/backend/internal/common"
	"github.com/ericlp/tasteit/backend/internal/process"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Tags(c *gin.Context) {
	tags, err := process.GetTags()
	if err != nil {
		log.Printf("Error: Failed to retrieve tags due to %s\n", err)
		c.JSON(
			http.StatusInternalServerError,
			common.Error(common.ResponseFailedToRetrieveTags),
		)
		return
	}

	c.JSON(http.StatusOK, common.Success(tags))
}
