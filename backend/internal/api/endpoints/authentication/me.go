package authentication

import (
	"github.com/ericlp/tasteit/backend/internal/common"
	"github.com/ericlp/tasteit/backend/internal/models"
	"github.com/ericlp/tasteit/backend/internal/process"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type MeJson struct {
	User   models.User    `json:"user"`
	Owners []models.Owner `json:"owners"`
}

func Me(c *gin.Context) {
	sessionData, err := readSession(c)
	if err != nil {
		log.Printf("Failed to read user from session: %v\n", err)
		c.JSON(http.StatusUnauthorized, common.Error(common.ResponseNotAuthorized))
		return
	}

	owners, err := process.GetOwnersJson(sessionData.User.Id)
	if err != nil {
		log.Printf("Failed to get user's owners from db: %v\n", err)
		c.JSON(http.StatusInternalServerError, common.Error(common.ResponseInvalidUserId))
		return
	}

	c.JSON(http.StatusOK, common.Success(&MeJson{
		User:   sessionData.User,
		Owners: owners,
	}))
}
