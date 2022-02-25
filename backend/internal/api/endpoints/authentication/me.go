package authentication

import (
	"github.com/ericlp/tasteit/backend/internal/common"
	"github.com/ericlp/tasteit/backend/internal/process"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Me(c *gin.Context) {
	sessionData, err := readSession(c)
	if err != nil {
		log.Printf("Failed to read user from session: %v\n", err)
		c.JSON(http.StatusUnauthorized, common.Error(common.ResponseNotAuthorized))
		return
	}

	user, err := process.GetUserJson(sessionData.UserID)
	if err != nil {
		log.Printf("Failed to get user from db: %v\n", err)
		c.JSON(http.StatusInternalServerError, common.Error(common.ResponseInvalidUserId))
		return
	}

	c.JSON(http.StatusOK, common.Success(user))
}
