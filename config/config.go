package config

import (
	"errors"
	"fmt"
	"github.com/hashicorp/hcl/v2/hclsimple"
	_ "github.com/joho/godotenv"
	"log"
	"os"
)

// const DefaultLocation = "/etc/encedeus"
const DefaultLocation = "./"

type Configuration struct {
	Server     ServerConfiguration   `hcl:"server,block"`
	DB         DatabaseConfiguration `hcl:"database,block"`
	Auth       AuthConfiguration     `hcl:"auth,block"`
	Storage    StorageConfiguration  `hcl:"storage,block"`
	SMTP       SMTPConfiguration     `hcl:"smtp,block"`
	Validation ValidationConfig      `hcl:"validation,block"`
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

type StorageConfiguration struct {
	Directory string `hcl:"dir"`
}

type SMTPConfiguration struct {
	Host     string `hcl:"host"`
	Port     int    `hcl:"port"`
	Address  string `hcl:"address"`
	Password string `hcl:"password"`
}

type ValidationConfig struct {
	MaxEmailLen int `hcl:"max_email_len"`

	MaxNameLen int `hcl:"max_name_len"`
	MinNameLen int `hcl:"min_name_len"`

	MaxPassLen int `hcl:"max_pass_len"`
	MinPassLen int `hcl:"min_pass_len"`

	MaxPluginNameLen int `hcl:"max_plugin_name_len"`
	MinPluginNameLen int `hcl:"min_plugin_name_len"`

	MaxReleaseNameLen int `hcl:"max_release_name_len"`
	MinReleaseNameLen int `hcl:"min_release_name_len"`
}

func (s *ServerConfiguration) URI() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

var Config Configuration

func InitConfig() {
	if _, err := os.Stat(fmt.Sprintf("%s/config.hcl", DefaultLocation)); errors.Is(err, os.ErrNotExist) {
		log.Fatal("Config file does not exist")
	}

	err := hclsimple.DecodeFile(fmt.Sprintf("%s/config.hcl", DefaultLocation), nil, &Config)
	if err != nil {
		log.Fatalf("Failed to load configuration file: %v", err)
	}
}
