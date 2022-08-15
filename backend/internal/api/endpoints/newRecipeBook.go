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

type NewRecipeBookResponse struct {
	UniqueName string `json:"uniqueName"`
}

func NewRecipeBook(c *gin.Context) {
	recipeBook, err := validateNewRecipeBook(c)
	if err != nil {
		log.Printf("Failed to validate new recipe book json: %v\n", err)
		c.JSON(http.StatusBadRequest, common.Error(common.ResponseInvalidJson))
		return
	}

	uniqueName, err := process.CreateNewRecipeBook(recipeBook)
	if err != nil {
		if errors.Is(err, common.ErrNameTaken) {
			log.Printf("Tried to create duplicate recipebook")
			c.JSON(
				http.StatusUnprocessableEntity,
				common.Error(common.ResponseRecipeBookNameExists),
			)
			return
		}

		log.Printf("Failed to create new recipebook: %v\n", err)
		c.JSON(
			http.StatusInternalServerError,
			common.Error(common.ResponseFailedToCreateRecipeBook),
		)
		return
	}

	c.JSON(
		http.StatusCreated, common.Success(
			NewRecipeBookResponse{
				UniqueName: uniqueName,
			},
		),
	)
}

func validateNewRecipeBook(c *gin.Context) (*models.NewRecipeBookJson, error) {
	var recipeBook models.NewRecipeBookJson
	err := c.ShouldBindJSON(&recipeBook)
	return &recipeBook, err
}
