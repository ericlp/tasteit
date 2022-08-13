package authentication

import (
	"fmt"
	"github.com/ericlp/tasteit2/backend/internal/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

type gammaMeResponse struct {
	GammaId        uuid.UUID `json:"id"`
	Cid            string    `json:"cid"`
	Nick           string    `json:"nick"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Language       string    `json:"language"`
	AvatarUrl      string    `json:"avatarUrl"`
	Gdpr           bool      `json:"gdpr"`
	UserAgreement  bool      `json:"userAgreement"`
	AccountLocked  bool      `json:"accountLocked"`
	AcceptanceYear int       `json:"acceptanceYear"`
	Authorities    []struct {
		Id        uuid.UUID `json:"id"`
		Authority string    `json:"authority"`
	} `json:"authorities"`
	Activated             bool   `json:"activated"`
	Enabled               bool   `json:"enabled"`
	Username              string `json:"username"`
	CredentialsNonExpired bool   `json:"credentialsNonExpired"`
	AccountNonExpired     bool   `json:"accountNonExpired"`
	AccountNonLocked      bool   `json:"accountNonLocked"`
	Groups                []struct {
		Id              uuid.UUID `json:"id"`
		BecomesActive   int64     `json:"becomesActive"`
		BecomesInactive int64     `json:"becomesInactive"`
		Description     struct {
			Sv string `json:"sv"`
			En string `json:"en"`
		} `json:"description"`
		Email   string `json:"email"`
		Purpose struct {
			Sv string `json:"sv"`
			En string `json:"en"`
		} `json:"function"`
		Name       string `json:"name"`
		PrettyName string `json:"prettyName"`
		AvatarURL  string `json:"avatarURL"`
		SuperGroup struct {
			Id         uuid.UUID `json:"id"`
			Name       string    `json:"name"`
			PrettyName string    `json:"prettyName"`
			Type       string    `json:"type"`
			Email      string    `json:"email"`
		} `json:"superGroup"`
		Active bool `json:"active"`
	} `json:"groups"`
}

const providerGamma = "Gamma"

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
	err = setSession(c, user.Nick, user.Email, providerGamma, token)
	if err != nil {
		log.Printf("Failed to save sessionData: %v\n", err)
		abort(c)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, "/")
	return

}

func GammaMeRequest(accessToken string) (*gammaMeResponse, error) {
	var user gammaMeResponse
	_, err := common.GetRequest(common.GetEnvVars().GammaMeUri, map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", accessToken)}, &user)
	return &user, err
}
