package endpoints

import (
	"github.com/ericlp/tasteit/backend/internal/common"
	"github.com/ericlp/tasteit/backend/internal/process"
	"github.com/gin-gonic/gin"
	"log"
)

func Tags(c *gin.Context) {
	tags, err := process.GetTags()
	if err != nil {
		log.Printf("Error: Failed to retrieve tags due to %s\n", err)
		c.JSON(500, common.Error(common.ResponseFailedToRetrieveTags))
		return
	}

	c.JSON(200, common.Success(tags))
}
