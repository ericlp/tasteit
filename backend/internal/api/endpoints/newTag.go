package endpoints

import (
	"errors"
	"github.com/ericlp/tasteit/backend/internal/common"
	"github.com/ericlp/tasteit/backend/internal/models"
	"github.com/ericlp/tasteit/backend/internal/process"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func NewTag(c *gin.Context) {
	tagJson, err := validateTag(c)
	if err != nil {
		log.Printf("Failed to validate new tag %v\n", err)
		c.JSON(http.StatusBadRequest, common.Error(common.ResponseInvalidJson))
		return
	}

	tag, err := process.CreateNewTag(tagJson)
	if err != nil {
		if errors.Is(err, common.ErrNameTaken) {
			c.JSON(
				http.StatusUnprocessableEntity,
				common.Error(common.ResponseTagNameTaken),
			)
			return
		}
		log.Printf("Failed creating tag %v\n", err)
		c.JSON(
			http.StatusInternalServerError,
			common.Error(common.ResponseFailedToCreateTag),
		)
		return
	}

	c.JSON(http.StatusCreated, common.Success(tag.ID))
}

func validateTag(c *gin.Context) (*models.NewTagJson, error) {
	var tag models.NewTagJson
	err := c.ShouldBindJSON(&tag)

	return &tag, err
}
