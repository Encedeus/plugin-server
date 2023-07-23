package config

import (
	"errors"
	"fmt"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const DefaultLocation = "./"

var Config Configuration

type Configuration struct {
	Credentials CredentialConfiguration
	Server      ServerConfiguration   `hcl:"server,block"`
	DB          DatabaseConfiguration `hcl:"database,block"`
	Auth        AuthConfiguration     `hcl:"auth,block"`
	SMTP        SMTPConfiguration     `hcl:"smtp,block"`
	CDN         CDNConfiguration      `hcl:"cdn,block"`
}

type CredentialConfiguration struct {
	Email         string
	EmailPassword string
}

type ServerConfiguration struct {
	Host string `hcl:"host"`
	Port int    `hcl:"port"`
}

type DatabaseConfiguration struct {
	Host     string `hcl:"host"`
	Port     int    `hcl:"port"`
	User     string `hcl:"user"`
	DBName   string `hcl:"name"`
	Password string `hcl:"password"`
}

type AuthConfiguration struct {
	JWTSecretAccess  string `hcl:"jwt_secret_access"`
	JWTSecretRefresh string `hcl:"jwt_secret_refresh"`
}

type SMTPConfiguration struct {
	Host string `hcl:"host"`
	Port int    `hcl:"port"`
}

type CDNConfiguration struct {
	Directory string `hcl:"dir"`
}

func (s *ServerConfiguration) URI() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

func InitConfig() {
	if _, err := os.Stat(fmt.Sprintf("%s/config.hcl", DefaultLocation)); errors.Is(err, os.ErrNotExist) {
		log.Fatal("Hcl config file does not exist")
	}

	err := hclsimple.DecodeFile(fmt.Sprintf("%s/config.hcl", DefaultLocation), nil, &Config)
	if err != nil {
		log.Fatalf("Failed to load hcl configuration file: %v", err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// store credentials in an ignored file because this is a public repo
	Config.Credentials = CredentialConfiguration{os.Getenv("EMAIL_NAME"), os.Getenv("EMAIL_PASS")}
}
