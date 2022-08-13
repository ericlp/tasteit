package models

import "github.com/google/uuid"

type GammaGroup struct {
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
}

type GammaMe struct {
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
	Activated             bool         `json:"activated"`
	Enabled               bool         `json:"enabled"`
	Username              string       `json:"username"`
	CredentialsNonExpired bool         `json:"credentialsNonExpired"`
	AccountNonExpired     bool         `json:"accountNonExpired"`
	AccountNonLocked      bool         `json:"accountNonLocked"`
	Groups                []GammaGroup `json:"groups"`
}
