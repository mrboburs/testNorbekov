package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	// "github.com/spf13/viper"
)

type Configs struct {
	ServiceHost string
	HTTPPort    string
	PhotoPath   string
	DBHost      string
	DBPort      string
	DBUsername  string
	DBName      string
	DBPassword  string
	DBSSLMode   string
}

func InitConfig() (cfg *Configs, err error) {

	fmt.Println("not here")

	if err := godotenv.Load(); err != nil {
		return cfg, fmt.Errorf("error loading env variables: %s", err.Error())
	}

	cfg = &Configs{

		ServiceHost: os.Getenv("DEPLOY"),
		HTTPPort:    os.Getenv("PORT"),
		PhotoPath:   os.Getenv("PHOTO_PATH"),
		DBHost:      os.Getenv("POSTGRES_HOST"),
		DBPort:      os.Getenv("POSTGRES_PORT"),
		DBUsername:  os.Getenv("POSTGRES_USER"),
		DBName:      os.Getenv("POSTGRES_DB"),
		DBSSLMode:   os.Getenv("POSTGRES_SSLMODE"),
		DBPassword:  os.Getenv("POSTGRES_PASSWORD"),
	}
	return
}
