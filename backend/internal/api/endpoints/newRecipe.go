package endpoints

import (
	"errors"
	"github.com/ericlp/tasteit/backend/internal/common"
	"github.com/ericlp/tasteit/backend/internal/db/queries"
	"github.com/ericlp/tasteit/backend/internal/db/tables"
	"github.com/ericlp/tasteit/backend/internal/models"
	"github.com/ericlp/tasteit/backend/internal/process"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

var ErrNoUserInContext = errors.New("no userID could be extracted from the context")
var ErrInvalidUserIdInContext = errors.New("the userID in the context was of an invalid type")

type NewRecipeJson struct {
	UniqueName string `json:"uniqueName"`
}

func NewRecipe(c *gin.Context) {
	recipeJson, err := validateNewRecipe(c)
	if err != nil {
		log.Printf("Failed to validate new recipe json: %v\n", err)
		c.JSON(http.StatusBadRequest, common.Error(common.ResponseInvalidJson))
		return
	}

	uniqueName, err := process.CreateNewRecipe(recipeJson)
	if err != nil {
		if errors.Is(err, common.ErrNameTaken) {
			log.Printf("Tried to create duplicate recipe")
			c.JSON(
				http.StatusUnprocessableEntity,
				common.Error(common.ResponseRecipeNameExist),
			)
			return
		}

		log.Printf("Failed to create new recipe: %v\n", err)
		c.JSON(
			http.StatusInternalServerError,
			common.Error(common.ResponseFailedToCreateRecipe),
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		common.Success(NewRecipeJson{UniqueName: uniqueName}),
	)
}

func validateNewRecipe(c *gin.Context) (*models.NewRecipeJson, error) {
	var recipe models.NewRecipeJson
	err := c.ShouldBindJSON(&recipe)
	return &recipe, err
}

var ErrUserNotAuthorized = errors.New("user not authorized")

func validateOwnerAuthorized(c *gin.Context, userId uuid.UUID) error {
	user, err := getSessionUser(c)
	if err != nil {
		return err
	}

	owners, err := queries.GetOwnersByUser(user.ID)
	if err != nil {
		return err
	}

	isAuthorized := false
	for _, owner := range owners {
		if owner.ID == userId {
			isAuthorized = true
		}
	}

	if isAuthorized {
		return nil

	}
	return ErrUserNotAuthorized
}

func getSessionUser(c *gin.Context) (*tables.User, error) {
	userId, exists := c.Get("userId")
	if !exists {
		return nil, ErrNoUserInContext
	}

	id, ok := userId.(uuid.UUID)
	if !ok {
		log.Printf("Failed to cast %s to UUID", userId)
		return nil, ErrInvalidUserIdInContext
	}

	return queries.GetUser(id)
}
