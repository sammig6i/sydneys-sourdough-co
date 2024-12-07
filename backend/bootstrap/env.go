package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	ContextTimeout      int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost              string `mapstructure:"DB_HOST"`
	DBPort              string `mapstructure:"DB_PORT"`
	DBUser              string `mapstructure:"DB_USER"`
	DBPass              string `mapstructure:"DB_PASSWORD"`
	DBName              string `mapstructure:"DB_NAME"`
	DatabaseURL         string `mapstructure:"DB_URL"`
	EmbeddingServiceURL string `mapstructure:"EMBEDDING_SERVICE_URL"`
	BackendPort         string `mapstructure:"BACKEND_PORT"`
	Environment         string `mapstructure:"ENVIRONMENT"`
}

func NewEnv() *Env {
	env := Env{}

	viper.SetConfigFile("/app/.env")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: Error reading config file: %v", err)
		viper.SetConfigFile("/app/.env.local")
		if err := viper.ReadInConfig(); err != nil {
			log.Printf("Warning: Error reading local config file: %v", err)
		}
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Printf("Warning: Error unmarshaling config: %v", err)
	}

	return &env
}
