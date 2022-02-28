package authentication

import (
	"fmt"
	"github.com/ericlp/tasteit/backend/internal/common"
	"github.com/ericlp/tasteit/backend/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

var gammaConfig *oauth2.Config

func init() {
	registerOnInit(func() {
		envVars := common.GetEnvVars()

		gammaConfig = &oauth2.Config{
			ClientID:     envVars.GammaClientID,
			ClientSecret: envVars.GammaSecret,
			Endpoint: oauth2.Endpoint{
				AuthURL:  envVars.GammaAuthorizationUri,
				TokenURL: envVars.GammaTokenUri,
			},
			RedirectURL: envVars.GammaRedirectUri,
			Scopes:      []string{},
		}
	})
}

func GammaInitAuth(c *gin.Context) {
	initAuth(c, gammaConfig)
}

func GammaCallback(c *gin.Context) {
	token := handleCallback(c, gammaConfig)
	if token == nil {
		return
	}

	user, err := GammaMeRequest(token.AccessToken)
	if err != nil {
		log.Printf("Failed to perform user request: %v\n", err)
		renewAuth(c)
		return
	}

	// user.Nick may be wrong as it is used
	err = setSession(c, user, token)
	if err != nil {
		log.Printf("Failed to save sessionData: %v\n", err)
		abort(c)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, "/")
	return

}

func GammaMeRequest(accessToken string) (*models.GammaMe, error) {
	var user models.GammaMe
	_, err := common.GetRequest(common.GetEnvVars().GammaMeUri, map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", accessToken)}, &user)
	return &user, err
}
