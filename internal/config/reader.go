package config

import (
	"log"

	dotenv "github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Service Service
	Server  Server
}

type Service struct {
	PathSms         string `envconfig:"PATHSMS"`
	MmsAddress      string `envconfig:"MMSADDRESS"`
	PathVoiceCall   string `envconfig:"PATHVOICECALL"`
	PathEmail       string `envconfig:"PATHEMAIL"`
	PathBilling     string `envconfig:"PATHBILLING"`
	SupportAddress  string `envconfig:"SUPPORTADDRESS"`
	IncidentAddress string `envconfig:"INCIDENTADDRESS"`
}

type Server struct {
	Port string `envconfig:"SERVER_PORT"`
}

func Read() *Config {
	err := dotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config

	envconfig.MustProcess("", &cfg)

	return &cfg
}
