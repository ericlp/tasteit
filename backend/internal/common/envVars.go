package common

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type envVars struct {
	DbUser     string
	DbPassword string
	DbName     string
	DbHost     string
	ResetDb    bool

	ImageFolder string
	Secret      string
	GinMode     string
	Port        uint16

	GammaAuthorizationUri string
	GammaRedirectUri      string
	GammaTokenUri         string
	GammaMeUri            string
	GammaSecret           string
	GammaClientID         string
	GammaLogoutURL        string
}

var ginModes = []string{
	"debug",
	"release",
}

var vars envVars

func GetEnvVars() *envVars {
	return &vars
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Unable to load .env file")
	} else {
		log.Println("Loaded environment variables from .env file")
	}

	loadEnvVars()
}

func loadEnvVars() {
	vars = envVars{
		DbUser:     loadNonEmptyString("db_user"),
		DbPassword: loadNonEmptyString("db_password"),
		DbName:     loadNonEmptyString("db_name"),
		DbHost:     loadNonEmptyString("db_host"),
		ResetDb:    loadBool("reset_db"),

		Secret:      loadNonEmptyString("secret"),
		GinMode:     loadGinMode("GIN_MODE"),
		Port:        loadUint16("PORT"),
		ImageFolder: loadNonEmptyString("image_folder"),

		GammaAuthorizationUri: loadNonEmptyString("GAMMA_AUTHORIZATION_URI"),
		GammaRedirectUri:      loadNonEmptyString("GAMMA_REDIRECT_URI"),
		GammaTokenUri:         loadNonEmptyString("GAMMA_TOKEN_URI"),
		GammaMeUri:            loadNonEmptyString("GAMMA_ME_URI"),
		GammaSecret:           loadNonEmptyString("GAMMA_SECRET"),
		GammaClientID:         loadNonEmptyString("GAMMA_CLIENT_ID"),
		GammaLogoutURL:        loadNonEmptyString("GAMMA_LOGOUT_URL"),
	}
}

func loadNonEmptyString(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Environment variable '%s' is not set or empty\n", key)
	}

	return val
}

func loadUint16(key string) uint16 {
	val := loadNonEmptyString(key)
	num, err := strconv.ParseUint(val, 10, 16)
	if err != nil {
		log.Fatalf("Environment variable '%s' is not a valid uint16: %v\n", key, err)
	}

	return uint16(num)
}

func loadBool(key string) bool {
	val := loadNonEmptyString(key)
	b, err := strconv.ParseBool(val)
	if err != nil {
		log.Fatalf("Environment variable '%s' is not a valid boolean: %v\n", key, err)
	}

	return b
}

func loadGinMode(key string) string {
	val := loadNonEmptyString(key)
	for _, mode := range ginModes {
		if mode == val {
			return val
		}
	}

	log.Fatalf("Invalid gin mode '%s', must be one of: %+v\n", val, ginModes)
	return ""
}
